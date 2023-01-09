package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/logging"
	"github.com/xuri/excelize/v2"
)

const (
	sheetCreated    = "已去重"
	sheetDuplicated = "重复的"
	sheetEmpty      = "空"
	sheetDefault    = "Sheet1"
)

type (
	// Duplicate 重复
	Duplicate struct {
		logger *logging.Logger
	}

	duplicateCounter struct {
		create    int
		duplicate int
		empty     int
	}

	duplicateInit struct {
		create    bool
		duplicate bool
		empty     bool
	}
)

func newDuplicate(logger *logging.Logger) *Duplicate {
	return &Duplicate{
		logger: logger,
	}
}

func (d *Duplicate) Removal(ins []string, out string, headers []string, sheets ...string) (err error) {
	output := excelize.NewFile()
	// 读取所有行内容
	if nil == sheets {
		sheets = []string{sheetDefault}
	}

	start := 1
	if 0 != len(headers) {
		start = 2
	}
	contents := make(map[string]bool)
	counter := &duplicateCounter{
		create:    start,
		duplicate: start,
		empty:     start,
	}
	init := &duplicateInit{
		create:    false,
		duplicate: false,
	}
	for _, in := range ins {
		var input *excelize.File
		if input, err = excelize.OpenFile(in); nil != err {
			continue
		}

		if err = d.read(input, output, headers, contents, counter, init, sheets...); nil != err {
			continue
		}
		if err = input.Close(); nil != err {
			continue
		}
	}

	// 删除默认表
	if err = output.DeleteSheet(sheetDefault); nil == err {
		err = output.SaveAs(out)
	}

	return
}

func (d *Duplicate) read(
	in *excelize.File, out *excelize.File,
	headers []string, contents map[string]bool,
	counter *duplicateCounter, init *duplicateInit,
	sheets ...string,
) (err error) {
	for _, sheet := range sheets {
		if err = d.sheet(in, out, headers, contents, counter, init, sheet); nil != err {
			continue
		}
	}

	return
}

func (d *Duplicate) sheet(
	in *excelize.File, out *excelize.File,
	headers []string, contents map[string]bool,
	counter *duplicateCounter, init *duplicateInit,
	sheet string,
) (err error) {
	var rows *excelize.Rows
	if rows, err = in.Rows(sheet); nil != err {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	var columns []string
	for rows.Next() {
		if columns, err = rows.Columns(); nil != err || reflect.DeepEqual(columns, headers) {
			continue
		}

		content := columns[3]
		if "" == strings.TrimSpace(content) {
			err = d.empty(out, sheet, content, &counter.empty, &init.empty, headers, columns...)
		} else if _, ok := contents[content]; !ok {
			err = d.new(out, sheet, content, &counter.create, &init.create, headers, columns...)
			contents[content] = true
		} else {
			err = d.old(out, sheet, content, &counter.duplicate, &init.duplicate, headers, columns...)
		}

		fields := gox.Fields[any]{
			field.New("sheet", sheet),
			field.New("contents", columns),
		}
		if nil != err {
			d.logger.Warn("写入数据遇到错误", fields.Connect(field.Error(err))...)
		} else {
			d.logger.Debug("写入数据成功", fields...)
		}
	}

	return
}

func (d *Duplicate) new(
	out *excelize.File,
	sheet string, content string,
	row *int, init *bool,
	headers []string,
	columns ...string,
) (err error) {
	if !*init {
		err = d.init(out, sheetCreated, init, headers...)
	}
	if nil != err {
		return
	}

	fields := gox.Fields[any]{
		field.New("sheet", sheet),
		field.New("content", content),
		field.New("contents", columns),
	}
	for index := 0; index < len(columns); index++ {
		err = out.SetCellValue(sheetCreated, d.indexToColumnName(index, *row), columns[index])
		if nil != err {
			d.logger.Warn("写入非重复数据遇到错误", fields.Connect(field.Error(err))...)
		} else {
			d.logger.Debug("写入非重复数据成功", fields...)
		}
	}
	// 如果写入都没有错误，行数往下移
	if nil == err {
		*row++
	}

	return
}

func (d *Duplicate) old(
	out *excelize.File,
	sheet string, content string,
	row *int, init *bool,
	headers []string,
	columns ...string,
) (err error) {
	if !*init {
		err = d.init(out, sheetDuplicated, init, headers...)
	}
	if nil != err {
		return
	}

	fields := gox.Fields[any]{
		field.New("sheet", sheet),
		field.New("content", content),
		field.New("contents", columns),
	}
	for index := 0; index < len(columns); index++ {
		err = out.SetCellDefault(sheetDuplicated, d.indexToColumnName(index, *row), columns[index])
		if nil != err {
			d.logger.Warn("写入重复数据遇到错误", fields.Connect(field.Error(err))...)
		} else {
			d.logger.Debug("写入重复数据成功", fields...)
		}
	}
	// 如果写入都没有错误，行数往下移
	if nil == err {
		*row++
	}

	return
}

func (d *Duplicate) empty(
	out *excelize.File,
	sheet string, content string,
	row *int, init *bool,
	headers []string,
	columns ...string,
) (err error) {
	if !*init {
		err = d.init(out, sheetEmpty, init, headers...)
	}
	if nil != err {
		return
	}

	fields := gox.Fields[any]{
		field.New("sheet", sheet),
		field.New("content", content),
		field.New("contents", columns),
	}
	for index := 0; index < len(columns); index++ {
		err = out.SetCellDefault(sheetEmpty, d.indexToColumnName(index, *row), columns[index])
		if nil != err {
			d.logger.Warn("写入重复数据遇到错误", fields.Connect(field.Error(err))...)
		} else {
			d.logger.Debug("写入重复数据成功", fields...)
		}
	}
	// 如果写入都没有错误，行数往下移
	if nil == err {
		*row++
	}

	return
}

func (d *Duplicate) init(out *excelize.File, sheet string, init *bool, headers ...string) (err error) {
	if _, ne := out.NewSheet(sheet); nil != ne {
		err = ne
	} else if err = d.writeHeaders(out, sheet, headers...); nil == err {
		*init = true
	}

	return
}

func (d *Duplicate) writeHeaders(out *excelize.File, sheet string, headers ...string) (err error) {
	fields := gox.Fields[any]{
		field.New("sheet", sheet),
		field.New("headers", headers),
	}
	for index := 0; index < len(headers); index++ {
		err = out.SetCellDefault(sheet, d.indexToColumnName(index, 1), headers[index])
		if nil != err {
			d.logger.Warn("写入表头错误", fields.Connect(field.Error(err))...)
		} else {
			d.logger.Debug("写入表头成功", fields...)
		}
	}

	return
}

func (d *Duplicate) indexToColumnName(index int, row int) (name string) {
	code := index + 65
	ascii := rune(code)
	name = fmt.Sprintf("%s%d", string(ascii), row)

	return
}
