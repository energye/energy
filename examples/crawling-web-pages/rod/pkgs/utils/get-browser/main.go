// Package main ...
package main

import (
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/launcher"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
