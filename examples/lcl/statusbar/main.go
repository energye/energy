package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetShowHint(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetPosition(types.PoScreenCenter)

	form2 := lcl.Application.CreateForm()
	form2.SetCaption("Form2")
	stb := lcl.NewStatusBar(form2)
	stb.SetParent(form2)
	stb.SetSimpleText("状态条")

	// 窗口全局show提示
	mainForm.SetShowHint(true)

	button := lcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetCaption("按钮1")

	// 长提示，后面使用  | 分割
	button.SetHint("这是一个提示|后面可以是长提示")
	button.SetBounds(10, 20, 100, 30)
	button.SetOnClick(func(sender lcl.IObject) {
		form2.Show()
	})

	button = lcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetCaption("按钮2")
	button.SetHint("提示。好好好好好。。。。。。。。。。。。。")
	button.SetBounds(10, 60, 100, 30)

	statusbar := lcl.NewStatusBar(mainForm)
	statusbar.SetParent(mainForm)
	statusbar.SetName("statusbar")

	// 右下角出现可调整窗口三角形，默认显示
	//statusbar.SetSizeGrip(false)

	// 当开启后，菜单的或者按钮的提示会出现在panel上
	statusbar.SetAutoHint(true)

	// 一般配合AutoHint使用
	statusbar.SetSimplePanel(false)

	pnl := lcl.AsStatusPanel(statusbar.Panels().Add())
	pnl.SetText("pnl1")
	pnl.SetWidth(100)

	pnl = lcl.AsStatusPanel(statusbar.Panels().Add())
	pnl.SetText("pnl1")
	pnl.SetAlignment(types.TaCenter)
	pnl.SetWidth(100)

	pnl = lcl.AsStatusPanel(statusbar.Panels().Add())
	pnl.SetText("第二格")
	pnl.SetAlignment(types.TaRightJustify)
	pnl.SetWidth(100)

	var i int32
	for i = 0; i < statusbar.Panels().Count(); i++ {
		fmt.Println("text:", lcl.AsStatusPanel(statusbar.Panels().Items(i)).Text())
	}
	lcl.Application.Run()
}
