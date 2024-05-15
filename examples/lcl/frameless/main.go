package main

import (
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)

var MainForm TMainForm

type TMainForm struct {
	lcl.TForm
}

func main() {
	lcl.DEBUG = true
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.Run()
}

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	m.SetCaption("frameless")
	m.SetPosition(types.PoScreenCenter)
	m.SetBorderStyleForFormBorderStyle(types.BsNone)
}
