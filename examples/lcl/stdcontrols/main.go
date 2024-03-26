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
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(500)
	mainForm.SetHeight(700)

	var top int32 = 40
	// TButton
	btn := lcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetLeft(10)
	btn.SetTop(top)
	btn.SetCaption("按钮1")
	btn.SetOnClick(func(lcl.IObject) {
		fmt.Println("按钮1单击")
	})

	top += btn.Height() + 5

	// TEdit
	edit := lcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(10)
	edit.SetTop(top)
	edit.SetTextHint("提示")
	//edit.SetText("文字")
	//	edit.SetReadOnly(true)
	edit.SetOnChange(func(lcl.IObject) {
		fmt.Println("文字改变了")
	})

	top += edit.Height() + 5
	// TEdit Password
	edit = lcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(10)
	edit.SetTop(top)
	edit.SetText("文字")
	edit.SetPasswordChar('*')

	top += edit.Height() + 5

	// TLabel
	lbl := lcl.NewLabel(mainForm)
	lbl.SetParent(mainForm)
	lbl.SetLeft(10)
	lbl.SetTop(top)
	lbl.SetCaption("标签1")
	lbl.Font().SetColor(255)

	top += lbl.Height() + 5

	// TCheckBox
	chk := lcl.NewCheckBox(mainForm)
	chk.SetParent(mainForm)
	chk.SetLeft(10)
	chk.SetTop(top)
	chk.SetCaption("选择框1")
	chk.SetOnClick(func(lcl.IObject) {
		fmt.Println("checked: ", chk.Checked())
	})

	// TStatusBar
	stat := lcl.NewStatusBar(mainForm)
	stat.SetParent(mainForm)
	// 不知道从哪个版开始，默认变成了true了
	stat.SetSimplePanel(false)
	//stat.SetSizeGrip(true) // 右下角是否有可调的
	spnl := lcl.AsStatusPanel(stat.Panels().Add())
	spnl.SetText("第一个")
	spnl.SetWidth(80)

	spnl = lcl.AsStatusPanel(stat.Panels().Add())
	spnl.SetText("第二个")
	spnl.SetWidth(80)

	// TToolBar
	tlbar := lcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)
	tlbar.SetShowCaptions(true)

	// 倒过来创建
	tlbtn := lcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetCaption("2")
	tlbtn.SetStyle(types.TbsDropDown)

	tlbtn = lcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetStyle(types.TbsSeparator)

	tlbtn = lcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetCaption("1")

	top += chk.Height() + 5
	// TRadioButton
	rd := lcl.NewRadioButton(mainForm)
	rd.SetParent(mainForm)
	rd.SetLeft(10)
	rd.SetTop(top)
	rd.SetCaption("选项1")

	var left int32 = rd.Left() + rd.Width() + 5

	rd = lcl.NewRadioButton(mainForm)
	rd.SetParent(mainForm)
	rd.SetLeft(left)
	rd.SetTop(top)
	rd.SetCaption("选项2")

	top += rd.Height() + 5
	// TMemo
	mmo := lcl.NewMemo(mainForm)
	mmo.SetParent(mainForm)
	mmo.SetBounds(10, top, 167, 50)
	//    mmo.Text()
	mmo.Lines().Add("1")
	mmo.Lines().Add("2")

	top += mmo.Height() + 5
	// TComboBox
	cb := lcl.NewComboBox(mainForm)
	cb.SetParent(mainForm)
	cb.SetLeft(10)
	cb.SetTop(top)
	cb.SetStyle(types.CsDropDownList)
	cb.Items().Add("1")
	cb.Items().Add("2")
	cb.Items().Add("3")
	cb.SetItemIndex(0)
	cb.SetOnChange(func(lcl.IObject) {
		if cb.ItemIndex() != -1 {
			fmt.Println(cb.Items().Strings(cb.ItemIndex()))
		}
	})

	// TListBox
	top += cb.Height() + 5
	lst := lcl.NewListBox(mainForm)
	lst.SetParent(mainForm)
	lst.SetBounds(10, top, 167, 50)
	lst.Items().Add("1")
	lst.Items().Add("2")
	lst.Items().Add("3")

	// TPanel
	top += lst.Height() + 5
	pnl := lcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetCaption("fff")
	//    pnl.SetShowCaption(false)
	pnl.SetBounds(10, top, 167, 50)

	// color
	top += pnl.Height() + 5
	clr := lcl.NewColorBox(mainForm)
	clr.SetParent(mainForm)
	clr.SetLeft(10)
	clr.SetTop(top)
	clr.SetOnChange(func(lcl.IObject) {
		if clr.ItemIndex() != -1 {
			lbl.Font().SetColor(clr.Selected())
		}
	})

	// TPageControl
	top += clr.Height() + 5
	pgc := lcl.NewPageControl(mainForm)
	pgc.SetParent(mainForm)
	pgc.SetBounds(10, top, 167, 100)
	pgc.SetOnChange(func(lcl.IObject) {
		fmt.Println("当前索引:", pgc.ActivePageIndex())
	})

	sheet := lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("一")
	btn = lcl.NewButton(mainForm)
	btn.SetParent(sheet)
	btn.SetLeft(10)
	btn.SetTop(10)
	btn.SetCaption("按钮1")

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("二")
	btn = lcl.NewButton(mainForm)
	btn.SetParent(sheet)
	btn.SetLeft(10)
	btn.SetTop(10)
	btn.SetCaption("按钮2")

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("三")
	btn = lcl.NewButton(mainForm)
	btn.SetParent(sheet)
	btn.SetLeft(10)
	btn.SetTop(10)
	btn.SetCaption("按钮3")

	// TImage
	top += pgc.Height() + 5
	img := lcl.NewImage(mainForm)
	img.SetBounds(10, top, 167, 97)
	img.SetParent(mainForm)
	img.Picture().LoadFromFile("E:\\SWT\\generate-tool\\examples\\lcl\\stdcontrols\\1.jpg")
	img.SetStretch(true)
	img.SetProportional(true)

	left = 210
	top = 10
	// TTrackBar
	trkbar := lcl.NewTrackBar(mainForm)
	trkbar.SetParent(mainForm)
	trkbar.SetBounds(left, top, 167, 20)
	trkbar.SetMax(100)
	trkbar.SetMin(0)
	trkbar.SetPosition(50)

	// TProgressBar
	top += trkbar.Height() + 10
	prgbar := lcl.NewProgressBar(mainForm)
	prgbar.SetParent(mainForm)
	prgbar.SetBounds(left, top, 10, 167)
	prgbar.SetMax(100)
	prgbar.SetMin(0)
	prgbar.SetPosition(1)
	prgbar.SetOrientation(types.PbVertical)

	trkbar.SetOnChange(func(lcl.IObject) {
		prgbar.SetPosition(trkbar.Position())
	})

	top += prgbar.Height() + 10

	dtp := lcl.NewDateTimePicker(mainForm)
	dtp.SetParent(mainForm)
	dtp.SetBounds(left, top, 167, 25)
	//dtp.SetDateMode(types.DmUpDown)
	dtp.SetKind(types.DtkDateTime)
	dtp.SetOptions(dtp.Options().Include(types.DtpoFlatButton))
	//dtp.SetShowMonthNames(true)
	dtp.SetTimeFormat(types.Tf12)
	//dtp.SetArrowShape(types.AsYetAnotherShape)
	dtp.SetDateDisplayOrder(types.DdoMDY)
	dtp.SetDateSeparator("-")
	dtp.SetTimeSeparator(".")
	dtp.SetHideDateTimeParts(dtp.HideDateTimeParts().Include(types.DtpYear))

	//top += dtp.Height() + 10
	//
	//mdtp := lcl.NewMonthCalendar(mainForm)
	//mdtp.SetParent(mainForm)
	//mdtp.SetBounds(left, top, 250, 250)
	//mdtp.SetOnClick(func(lcl.IObject) {
	//	fmt.Println(mdtp.Date())
	//})
	//
	//top += mdtp.Height() + 10
	//dtp.SetDateTime(time.Now().Add(time.Hour * 48))
	//dtp.SetDate(time.Now().AddDate(1, 0, 0))
	//
	//fmt.Println("time: ", mdtp.Date(), dtp.DateTime())
	//
	//btn = lcl.NewButton(mainForm)
	//btn.SetParent(mainForm)
	//btn.SetLeft(left)
	//btn.SetTop(top)
	//btn.SetCaption("改变日期")
	//btn.SetOnClick(func(lcl.IObject) {
	//	mdtp.SetDate(time.Now().AddDate(7777, 1, 23))
	//})
	//
	//top += btn.Height() + 10
	//spinedit := lcl.NewSpinEdit(mainForm)
	//spinedit.SetParent(mainForm)
	//spinedit.SetLeft(left)
	//spinedit.SetTop(top)
	//spinedit.SetWidth(100)
	//spinedit.SetMaxValue(10000)
	//spinedit.SetMinValue(50)
	//spinedit.SetValue(100)
	//spinedit.SetOnChange(func(sender lcl.IObject) {
	//	fmt.Println(spinedit.Value())
	//})

	// run
	lcl.Application.Run()
}
