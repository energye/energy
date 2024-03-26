package main

import (
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/colors"
)

var (
	mainForm lcl.IForm
	richEdit lcl.IRichEdit
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	initMainForm()
	initMainMenu()
	tlbar := lcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)

	richEdit = lcl.NewRichEdit(mainForm)
	richEdit.SetParent(mainForm)
	richEdit.SetAlign(types.AlClient)
	richEdit.Lines().Add("这是一段文字红色，粗体，斜體")
	richEdit.SetSelStart(6)
	richEdit.SetSelLength(2)
	richEdit.SelAttributes().SetColor(colors.ClRed)

	richEdit.SetSelStart(9)
	richEdit.SetSelLength(2)

	richEdit.SelAttributes().SetStyle(types.NewSet(types.FsBold))

	richEdit.SetSelStart(12)
	richEdit.SetSelLength(2)

	richEdit.SelAttributes().SetStyle(types.NewSet(types.FsItalic))

	richEdit.SetSelStart(15)

	initRichEditPopupMenu()

	stabar := lcl.NewStatusBar(mainForm)
	stabar.SetParent(mainForm)

	lcl.Application.Run()
}

func initMainForm() {
	mainForm = lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
}

func initMainMenu() {
	mainMenu := lcl.NewMainMenu(mainForm)

	item := lcl.NewMenuItem(mainForm)
	item.SetCaption("&File")
	mainMenu.Items().Add(item)

	item = lcl.NewMenuItem(mainForm)
	item.SetCaption("&Help")
	mainMenu.Items().Add(item)
}

func initRichEditPopupMenu() {
	pm := lcl.NewPopupMenu(mainForm)
	item := lcl.NewMenuItem(mainForm)
	item.SetCaption("&Clear")
	pm.Items().Add(item)

	richEdit.SetPopupMenu(pm)
}
