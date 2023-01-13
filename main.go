package main

import (
	"github.com/pangum/pangu"
)

func main() {
	panic(pangu.New(
		pangu.Named("jchu"),
		pangu.Banner("Jchu", pangu.BannerTypeAscii),
	).Run(newBootstrap))
}
