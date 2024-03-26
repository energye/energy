// govcl project main.go
// go.exe build -i -ldflags="-H windowsgui"
package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/lcl/rtl/version"
	"github.com/energye/energy/v2/pkgs/win"
	"github.com/energye/energy/v2/types"
)

var (
	mainForm lcl.IForm
	trayicon lcl.ITrayIcon
)

func main() {
	inits.Init(nil, nil)
	// 异常捕获
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
			lcl.ShowMessage(err.(error).Error())
		}
	}()

	fmt.Println("InheritsFrom：", lcl.Application.Is().Object())
	fmt.Println("InheritsFrom：", lcl.Application.Is().Component())
	fmt.Println("InheritsFrom：", lcl.Application.Is().Control())
	fmt.Println(rtl.LibAbout())

	guid := rtl.CreateGUID()
	fmt.Println("guid:", guid)
	guidstr := rtl.GUIDToString(guid)
	fmt.Println("guidToStr:", guidstr)
	fmt.Println("strToGUID: ", rtl.StringToGUID(guidstr))

	fmt.Println("main")

	lcl.Application.Initialize()
	lcl.Application.SetOnException(func(sender lcl.IObject, e lcl.IException) {
		fmt.Println("exception.")
	})

	lcl.Application.SetTitle("Hello World! 系统信息：" + version.OSVersion.ToString())
	lcl.Application.SetMainFormOnTaskBar(true)
	// 窗口自动根据系统绽放，默认为false
	//lcl.Application.SetScaled(true)

	mainForm = lcl.Application.CreateForm()
	mainForm.SetWidth(800)
	mainForm.SetHeight(600)
	mainForm.SetOnClose(func(Sender lcl.IObject, Action *types.TCloseAction) {
		fmt.Println("close")
	})

	// 窗口大小约束
	mainForm.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		*minWidth = 800
		*minHeight = 600
		*maxWidth = 800
		*maxHeight = 600
	})

	mainForm.SetOnDestroy(func(sender lcl.IObject) {
		fmt.Println("Form Destroy.")
	})

	fmt.Println("MainForm ClientRect: ", mainForm.ClientRect())
	filename := lcl.Application.ExeName()
	fmt.Println("application.ExeName: ", filename)
	fmt.Println("path: ", rtl.ExtractFilePath(filename))
	fmt.Println("fileExists: ", rtl.FileExists(filename))

	mainForm.SetOnCloseQuery(func(Sender lcl.IObject, CanClose *bool) {
		*CanClose = lcl.MessageDlg("是否退出?", types.MtInformation, types.MbYes, types.MbNo) == types.MrYes
		fmt.Println("OnCloseQuery")
	})

	mainForm.SetCaption(lcl.Application.Title())
	mainForm.EnabledMaximize(false)
	mainForm.SetDoubleBuffered(true)
	//mainForm.SetPosition(types.PoScreenCenter)
	//mainForm.ScreenCenter()
	mainForm.WorkAreaCenter()
	mainForm.SetKeyPreview(true)
	mainForm.SetOnKeyDown(func(Sender lcl.IObject, Key *types.Char, Shift types.TShiftState) {
		fmt.Println(Shift.In(types.SsCtrl))
		fmt.Println("key:", *Key)
	})

	mainForm.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("Button:", button == types.MbLeft, ", X:", x, ", y:", y)
		fmt.Println("OnMouseDown")
	})

	chk := lcl.NewCheckBox(mainForm)
	chk.SetParent(mainForm)
	chk.SetChecked(true)
	chk.SetCaption("测试")
	chk.SetLeft(1)
	chk.SetTop(60)
	chk.SetOnClick(func(lcl.IObject) {
		fmt.Println("chk.Checked=", chk.Checked())
	})

	// action
	action := lcl.NewAction(mainForm)
	action.SetCaption("action1")
	action.SetOnUpdate(func(sender lcl.IObject) {
		lcl.AsAction(sender).SetEnabled(chk.Checked())
	})
	action.SetOnExecute(func(lcl.IObject) {
		fmt.Println("action execute")
	})
	btn := lcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetBounds(250, 30, 90, 25)
	btn.SetCaption("action")
	btn.SetAction(action)

	trayicon = lcl.NewTrayIcon(mainForm)
	trayicon.SetIcon(lcl.Application.Icon()) //不设置会自动使用Application.Icon

	trayicon.SetHint(mainForm.Caption())
	trayicon.SetVisible(true)
	trayicon.SetOnClick(func(lcl.IObject) {
		trayicon.SetBalloonTitle("test")
		trayicon.SetBalloonTimeout(10000)
		trayicon.SetBalloonHint("我是提示正文啦")
		trayicon.ShowBalloonHint()
		fmt.Println("TrayIcon Click.")
	})

	// img
	img := lcl.NewImage(mainForm)
	img.SetBounds(132, 30, 156, 97)
	img.SetParent(mainForm)
	img.Picture().LoadFromFile("./imgs/1.jpg")
	//img.SetStretch(true)
	img.SetProportional(true)

	// linklabel
	linklbl := lcl.NewLinkLabel(mainForm)
	linklbl.SetAlign(types.AlBottom)
	linklbl.SetCaption("<a href=\"https://www.github.com/energye/energy\">Go Energy测试链接</a>")
	linklbl.SetParent(mainForm)
	linklbl.SetOnLinkClick(func(sender lcl.IObject, link string, linktype types.TSysLinkType) {
		fmt.Println("link label: ", link, ", type: ", linktype)
		rtl.SysOpen(link)
	})

	// menu
	mainMenu := lcl.NewMainMenu(mainForm)
	item := lcl.NewMenuItem(mainForm)
	item.SetCaption("File(&F)")
	mainMenu.Items().Add(item)

	item2 := lcl.NewMenuItem(mainForm)
	item2.SetCaption("MemoryStreamTest")
	item2.SetOnClick(func(lcl.IObject) {
		mem := lcl.NewMemoryStream()
		defer mem.Free()
		mem.Write([]byte("测试"))
		mem.SaveToFile("test.txt")

		mem.SetPosition(0)
		bs := mem.Read(int32(mem.Size()))
		fmt.Println(", bs:", bs, ", str:", string(bs))
	})
	item.Add(item2)

	item2 = lcl.NewMenuItem(mainForm)
	item2.SetCaption("Exit(&E)")
	//item2.SetShortCutFromString("Ctrl+Q")
	item2.SetOnClick(func(lcl.IObject) {
		mainForm.Close()
	})
	item.Add(item2)

	//	mainForm.EnabledMinimize(false)
	//	mainForm.EnabledSystemMenu(false)

	button := lcl.NewButton(mainForm)

	button.SetCaption("消息")
	button.SetParent(mainForm)
	button.SetOnClick(func(lcl.IObject) {
		fmt.Println("button click")
		lcl.ShowMessage("这是一个消息")
		lcl.Application.MessageBox("Hello!", "Message", win.MB_YESNO+win.MB_ICONINFORMATION)
	})
	button.SetLeft(50)
	button.SetTop(50)
	button.SetAlign(types.AlRight)

	edit := lcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(1)
	edit.SetTop(30)
	edit.SetTextHint("测试")
	edit.SetOnChange(func(lcl.IObject) {
		fmt.Println("edit OnChange")
	})

	button2 := lcl.NewButton(mainForm)
	button2.SetParent(mainForm)
	button2.SetCaption("a")
	button2.SetWidth(100)
	button2.SetHeight(28)
	button2.SetOnClick(func(lcl.IObject) {
		fmt.Println("button2 click")

		edit.SetText("Hello!")
		fmt.Println("ScreenWidth:", lcl.Screen.Width(), ", ScreenHeight:", lcl.Screen.Height())
	})
	button2.SetAlign(types.AlTop)

	combo := lcl.NewComboBox(mainForm)
	combo.SetAlign(types.AlBottom)
	combo.SetParent(mainForm)
	combo.SetText("ffff")
	combo.Items().Add("1")
	combo.Items().Add("2")
	combo.SetItemIndex(0)
	combo.SetOnChange(func(lcl.IObject) {
		if combo.ItemIndex() != -1 {
			fmt.Println("combo Change: ", combo.Items().Strings(combo.ItemIndex()))
		}

	})

	page := lcl.NewPageControl(mainForm)
	page.SetParent(mainForm)
	page.SetAlign(types.AlBottom)
	sheet := lcl.NewTabSheet(mainForm)
	sheet.SetPageControl(page)
	sheet.SetCaption("第一页")

	// 需要先将TabSheet设置了父窗口，TListView才可用，不然就会报错
	lv1 := lcl.NewListView(mainForm)
	lv1.SetAlign(types.AlClient)
	lv1.SetParent(sheet)

	lv1.SetViewStyle(types.VsReport)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetGridLines(true)
	col := lcl.AsListColumn(lv1.Columns().Add())
	col.SetCaption("序号")
	col.SetWidth(100)
	// 强制柱头宽，即使被调整也会被还原
	col.SetMaxWidth(100)
	col.SetMinWidth(100)
	col = lcl.AsListColumn(lv1.Columns().Add())
	col.SetCaption("名称")
	col.SetWidth(200)
	col = lcl.AsListColumn(lv1.Columns().Add())
	col.SetCaption("内容")
	col.SetWidth(200)
	lv1.SetOnClick(func(lcl.IObject) {
		if lv1.ItemIndex() != -1 {
			item := lv1.Selected() // lv1.Items().Item(lv1.ItemIndex())
			fmt.Println(item.Caption(),
				item.SubItems().Strings(0),
				item.SubItems().Strings(1))
		}
	})

	lv1.Items().BeginUpdate()
	for i := 1; i <= 50; i++ {
		lstitem := lv1.Items().Add()
		lstitem.SetCaption(fmt.Sprintf("%d", i))
		lstitem.SubItems().Add(fmt.Sprintf("第%d", i))
		lstitem.SubItems().Add(fmt.Sprintf("内容%d", i))
	}
	lv1.Items().EndUpdate()

	sheet = lcl.NewTabSheet(mainForm)
	sheet.SetCaption("第二页")
	sheet.SetPageControl(page)

	// -----------TreeView 不同Node弹出不同菜单，两个右键例程不同使用

	tvpm1 := lcl.NewPopupMenu(mainForm)
	mItem := lcl.NewMenuItem(mainForm)
	mItem.SetCaption("第一种")
	tvpm1.Items().Add(mItem)

	tvpm2 := lcl.NewPopupMenu(mainForm)
	mItem = lcl.NewMenuItem(mainForm)
	mItem.SetCaption("第二种")
	tvpm2.Items().Add(mItem)

	tv1 := lcl.NewTreeView(mainForm)
	tv1.SetAutoExpand(true)
	tv1.SetParent(sheet)
	tv1.SetAlign(types.AlClient)
	//	tv1.SetRightClickSelect(true)
	tv1.SetOnClick(func(lcl.IObject) {
		if tv1.SelectionCount() > 0 {
			node := tv1.Selected()
			fmt.Println("text:", node.Text(), ", index:", node.Index())
		}
	})

	tv1.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbRight {
			node := tv1.GetNodeAt(x, y)
			if node != nil && node.IsValid() {
				// 自由决择是否选中
				node.SetSelected(true)
				// 根据Level来判断，这里只是做演示
				p := lcl.Mouse.CursorPos()
				switch node.Level() {
				case 0:
					tvpm1.PopUpInOverload1(p.X, p.Y)
				case 1:
					tvpm2.PopUpInOverload1(p.X, p.Y)
				}
				fmt.Println("node.Level():", node.Level(), ", text:", node.Text())
			}
		}
	})

	tv1.Items().BeginUpdate()
	node := tv1.Items().AddChild(nil, "首个")
	for i := 1; i <= 50; i++ {
		tv1.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	node = tv1.Items().AddChild(nil, "第二个")
	for i := 1; i <= 50; i++ {
		tv1.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	tv1.Items().EndUpdate()

	fmt.Println("Compoment Count:", mainForm.ComponentCount())
	//	mainForm.ScreenCenter()

	lbl := lcl.NewLabel(mainForm)
	lbl.SetCaption("标签")
	lbl.SetAlign(types.AlBottom)
	fmt.Println("InheritsFromControl:", mainForm.Is().Control())
	fmt.Println("InheritsFromWinControl:", mainForm.Is().WinControl())
	fmt.Println("InheritsFromComponent:", mainForm.Is().Component())
	fmt.Println("InheritsFromWinControl:", lbl.Is().WinControl())

	lcl.Application.Run()
}
