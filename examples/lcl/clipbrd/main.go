//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)
import _ "github.com/energye/energy/v2/examples/syso"

type TMainForm struct {
	*lcl.TForm
	btn  lcl.IButton
	btn2 lcl.IButton
	btn3 lcl.IButton
	btn4 lcl.IButton
	btn5 lcl.IButton
	img  lcl.IImage
}

var mainForm *TMainForm

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&mainForm)
	lcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	f.btn = lcl.NewButton(f)
	f.btn.SetParent(f)
	f.btn.SetCaption("GetText")
	f.btn.SetOnClick(f.onBtnClick)

	f.btn2 = lcl.NewButton(f)
	f.btn2.SetParent(f)
	f.btn2.SetTop(40)
	f.btn2.SetWidth(120)
	f.btn2.SetCaption("GetTextBuffer")
	f.btn2.SetOnClick(f.onBtn2Click)

	f.btn3 = lcl.NewButton(f)
	f.btn3.SetParent(f)
	f.btn3.SetTop(80)
	f.btn3.SetWidth(120)
	f.btn3.SetCaption("SetText")
	f.btn3.SetOnClick(f.onBtn3Click)

	f.btn4 = lcl.NewButton(f)
	f.btn4.SetParent(f)
	f.btn4.SetTop(f.btn3.Top() + 50)
	f.btn4.SetWidth(120)
	f.btn4.SetCaption("SetBmp")
	f.btn4.SetOnClick(f.onBtn4Click)

	f.btn5 = lcl.NewButton(f)
	f.btn5.SetParent(f)
	f.btn5.SetTop(f.btn4.Top() + 50)
	f.btn5.SetWidth(120)
	f.btn5.SetCaption("GetBmp")
	f.btn5.SetOnClick(f.onBtn5Click)

	f.img = lcl.NewImage(f)
	f.img.SetParent(f)
	f.img.SetTop(f.btn5.Top() + 30)
	f.img.SetWidth(200)
	f.img.SetHeight(200)
}

func (f *TMainForm) onBtnClick(sender lcl.IObject) {
	str := lcl.Clipboard.AsText()
	fmt.Println("len1:", len(str), str)
}

func (f *TMainForm) onBtn2Click(sender lcl.IObject) {
	str := ""
	lcl.Clipboard.GetTextBuf(&str, 1000) // buff不够长
	fmt.Println("len2:", len(str), str)
}

func (f *TMainForm) onBtn3Click(sender lcl.IObject) {
	lcl.Clipboard.SetAsText("energy-lcl")
}

func (f *TMainForm) onBtn4Click(sender lcl.IObject) {
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	// only bmp
	mem.LoadFromFile("bg.bmp")
	mem.SetPosition(0)

	// 注册自定义的格式
	//lcl.RegisterClipboardFormat()

	// 预定义格式
	format := lcl.PredefinedClipboardFormat(types.PcfBitmap)
	fmt.Println("format:", format)

	if !lcl.Clipboard.SetFormat(format, mem) {
		lcl.ShowMessage("设置格式失败")
	} else {
		lcl.ShowMessage("设置成功")
	}
}

func (f *TMainForm) onBtn5Click(sender lcl.IObject) {
	if !lcl.Clipboard.HasPictureFormat() {
		return
	}
	bmpFormat := lcl.Clipboard.FindPictureFormatID()
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	if lcl.Clipboard.GetFormat(bmpFormat, mem) {
		mem.SetPosition(0)
		f.img.Picture().LoadFromStream(mem)
	}
}
