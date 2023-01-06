package core

import (
	"github.com/goexl/simaqian"
	"github.com/pangum/logging"
)

// Logger 仅作为引用少写代码用途
type Logger simaqian.Logger

func newLogger(logger *logging.Logger) Logger {
	return logger
}
