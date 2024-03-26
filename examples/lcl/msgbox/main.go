package main

import (
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/win"
	"github.com/energye/energy/v2/types"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	lcl.ShowMessage("消息")
	if lcl.MessageDlg("消息", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes {
		lcl.ShowMessage("你点击了“是")
	}
	if lcl.Application.MessageBox("消息", "标题", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == types.IdOK {
		lcl.ShowMessage("你点击了“是")
	}

	lcl.Application.Run()
}
