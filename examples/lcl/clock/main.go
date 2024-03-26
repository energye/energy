package main

import (
	"github.com/energye/energy/v2/api/exception"
	"github.com/energye/energy/v2/examples/lcl/clock/form"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
)

func main() {
	inits.Init(nil, nil)
	exception.SetOnException(func(message string) {

	})
	lcl.Application.SetOnException(func(sender lcl.IObject, e lcl.IException) {

	})
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&form.Form1)
	lcl.Application.Run()

}
