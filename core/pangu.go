package core

import "github.com/pangum/pangu"

func init() {
	pangu.New().Dependencies(
		newLogger,
	)
}
