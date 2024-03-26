package main

import (
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/colors"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(500)
	mainForm.SetHeight(400)

	var itemHeight int32 = 30
	listbox := lcl.NewListBox(mainForm)
	listbox.SetParent(mainForm)
	listbox.SetStyle(types.LbOwnerDrawFixed)
	listbox.SetAlign(types.AlClient)
	listbox.Items().Add("第一项")
	listbox.Items().Add("第二项")
	listbox.Items().Add("第三项")
	listbox.Items().Add("第四项")
	listbox.Items().Add("第五项")
	listbox.Items().Add("第六项")
	listbox.Items().Add("第七项")
	listbox.Items().Add("第八项")
	listbox.SetItemHeight(itemHeight)
	listbox.SetOnDrawItem(func(control lcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
		canvas := listbox.Canvas()
		s := listbox.Items().Strings(index)
		fw := canvas.TextWidth(s)
		fh := canvas.TextHeight(s)
		font := lcl.AsFont(canvas.Font())
		brush := lcl.AsBrush(canvas.Brush())
		pen := lcl.AsPen(canvas.Pen())
		font.SetColor(colors.ClBlack)
		brush.SetColor(colors.ClBtnFace)
		canvas.FillRect(&aRect)
		brush.SetColor(0x00FFF7F7)
		pen.SetColor(colors.ClSkyblue)
		canvas.Rectangle1(aRect.Left+1, aRect.Top+1, aRect.Right-1, aRect.Bottom-1)
		canvas.Rectangle1(aRect.Left, aRect.Top, aRect.Right, aRect.Bottom)
		if state.In(types.OdSelected) {
			brush.SetColor(0x00FFB2B5)
			canvas.Rectangle1(aRect.Left+1, aRect.Top+1, aRect.Right-1, aRect.Bottom-1)
			font.SetColor(colors.ClBlue)
			if state.In(types.OdFocused) {
				canvas.DrawFocusRect(&aRect)
			}
		}
		canvas.TextOut(aRect.Left+(aRect.Right-fw)/2, aRect.Top+(itemHeight-fh)/2, s)
	})
	lcl.Application.Run()
}
