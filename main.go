package main

import (
	"github.com/pangum/pangu"
)

func main() {
	panic(pangu.New(
		pangu.Named("cli"),
		pangu.Banner("Storezhang Cli", pangu.BannerTypeAscii),
	).Run(newBootstrap))
}
