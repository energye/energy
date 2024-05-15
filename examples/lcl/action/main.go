package main

import (
	"github.com/energye/energy/v2/examples/lcl/action/src"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
)

func main() {
	lcl.DEBUG = true
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&src.MainForm)
	lcl.Application.Run()
}
