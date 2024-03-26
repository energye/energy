package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/colors"
)

// 简单介绍下Lazarus中控件的布局方式
func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(700)
	mainForm.SetHeight(500)

	// 为方便演示，这里使用一个TPageControl作为多个布局的演示
	// Lazarus中所有可视与非可视组件的owner都可设置为TForm的实例，
	// 本示例中，TForm的实例即mainForm，当owner设置为mainForm时
	// 在TForm销毁前会自动Free相关组件
	pgc := lcl.NewPageControl(mainForm)

	// 如需让一个可视控件显示到某个控件上即可使用SetParent，
	// 这里需要注意并非所有控件都可使用SetParent，Lazarus中一般能接受子控件的只有
	// TWinControl，不过这里没有做区分。打个比方如TButton、TImage、TLabel是不可以作
	// 为父控件使用的，能作为父控件的常用的有TForm、TPanel、TPageControl、TTabSheet。
	pgc.SetParent(mainForm)
	// 这里将TPageControl设置为整个窗口客户区大小，并自动调整
	pgc.SetAlign(types.AlClient)

	//
	sheet := lcl.NewTabSheet(mainForm)
	// Lazarus中有几个控件的设置Parent属性有些不同，这里需要使用SetPageControl才行。
	// TTabSheet控件默认的Align为alClient，即填充父控件客户区域
	sheet.SetPageControl(pgc)
	sheet.SetCaption("顶-左-客户区")

	// 此时的pnl仅Height属性生效
	pnl := lcl.NewPanel(mainForm)
	pnl.SetCaption("顶")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClRed)
	pnl.SetParent(sheet)
	pnl.SetHeight(100)
	pnl.SetAlign(types.AlTop)

	// 此时的pnl仅Width属性生效
	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("左")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClGreen)
	pnl.SetParent(sheet)
	pnl.SetWidth(100)
	pnl.SetAlign(types.AlLeft)

	// 此时的pnl无法手动调整大小
	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("客户区")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlClient)

	// --------------------------------------------------------------------

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("顶-客户区-底")

	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("顶")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClRed)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlTop)

	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("客户区")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClGreen)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlClient)

	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("底")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlBottom)

	//--------------------------------------------------------------------

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("顶-客户区(左|-|右)-底")

	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("顶")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClRed)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlTop)

	ppnl := lcl.NewPanel(mainForm)
	ppnl.SetCaption("客户区")
	ppnl.SetParentBackground(false)
	ppnl.SetColor(colors.ClGreen)
	ppnl.SetParent(sheet)
	ppnl.SetAlign(types.AlClient)

	// 上个panel作为父控件
	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("左")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClAqua)
	pnl.SetParent(ppnl)
	pnl.SetAlign(types.AlLeft)

	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("右")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClAzure)
	pnl.SetParent(ppnl)
	pnl.SetAlign(types.AlRight)

	pnl = lcl.NewPanel(mainForm)
	pnl.SetCaption("底")
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlBottom)

	//----------------------------------Anchors----------------------------------

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("Anchors")

	pnl = lcl.NewPanel(mainForm)
	pnl.SetParentBackground(false)
	//pnl.SetColor(colors.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlClient)

	w := pnl.Width()
	h := pnl.Height()

	fmt.Println(w, h)

	btn := lcl.NewButton(mainForm)
	btn.SetParent(pnl)
	btn.SetCaption("左(Left)")
	btn.SetLeft(10)

	// lcl下使用ClientWidth或者ClientHeight
	// vcl下建议使用Width或者Height
	// 原因估计是两套组件对于某些方面的处理不同
	btn = lcl.NewButton(mainForm)
	btn.SetParent(pnl)
	btn.SetCaption("右(Right)")
	btn.SetLeft(w - btn.Width() - 10)
	btn.SetAnchors(types.NewSet(types.AkTop, types.AkRight))

	// 中
	btn = lcl.NewButton(mainForm)
	btn.SetParent(pnl)
	btn.SetCaption("中(Center)")
	btn.SetLeft(w - btn.Width() - 10)
	btn.AnchorHorizontalCenterTo(pnl)
	btn.AnchorVerticalCenterTo(pnl)

	btn = lcl.NewButton(mainForm)
	btn.SetParent(pnl)
	btn.SetCaption("左下(Left-Bottom)")
	btn.SetLeft(10)
	btn.SetTop(h - btn.Height() - 10)
	btn.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom))

	btn = lcl.NewButton(mainForm)
	btn.SetParent(pnl)
	btn.SetCaption("右下(Right-Bottom)")
	btn.SetLeft(w - btn.Width() - 10)
	btn.SetTop(h - btn.Height() - 10)
	btn.SetAnchors(types.NewSet(types.AkRight, types.AkBottom))

	//----------------------------------Margins----------------------------------

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("BorderSpacing")

	ppnl = lcl.NewPanel(mainForm)
	ppnl.SetParent(sheet)
	ppnl.SetParentBackground(false)
	ppnl.SetColor(colors.ClRed)
	ppnl.SetAlign(types.AlClient)

	pnl = lcl.NewPanel(mainForm)
	pnl.SetParent(ppnl)
	pnl.SetParentBackground(false)
	pnl.SetColor(colors.ClGreen)

	pnl.SetAlign(types.AlClient)

	m := pnl.BorderSpacing()
	m.SetLeft(20)
	m.SetTop(30)
	m.SetBottom(40)
	m.SetRight(50)

	//----------------------------------OnAlignPosition----------------------------------

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("Align = alCustom = OnAlignPosition")

	pnl = lcl.NewPanel(mainForm)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlClient)
	// 子控件如果有设置为AlCustom的，则会触发这个事件
	pnl.SetOnAlignPosition(onCustomAlignPosiion)

	pnl2 := lcl.NewPanel(mainForm)
	pnl2.SetParent(pnl)
	pnl2.SetAlign(types.AlCustom)
	pnl2.SetBounds(10, 10, 300, 300)
	// 子控件如果有设置为AlCustom的，则会触发这个事件
	pnl2.SetOnAlignPosition(onCustomAlignPosiion)

	btn = lcl.NewButton(mainForm)
	btn.SetParent(pnl2)
	btn.SetAlign(types.AlCustom) // 自定义
	btn.SetCaption("按钮。")

	lcl.Application.Run()
}

// sender 调用此事件的控件
// control 被调整的控件
// newLeft, newTop, newWidth, newHeight 保存被调整的控件原始位置和大小
// alignRect 保存对齐的矩形范围
// alignInfo 对齐信息
func onCustomAlignPosiion(sender lcl.IWinControl, control lcl.IControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo) {
	*newLeft = (alignRect.Width() - *newWidth) / 2
	*newTop = (alignRect.Height() - *newHeight) / 2
	fmt.Println(*newLeft, *newTop)
}
