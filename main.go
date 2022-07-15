package main

import (
	"github.com/pangum/pangu"
)

func main() {
	panic(pangu.New(
		pangu.Named(`cli`),
		pangu.Banner(`Sczx Cli`, pangu.BannerTypeAscii),
	).Run(newBootstrap))
}
