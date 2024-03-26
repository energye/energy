package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"strings"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	initComponents()
	lcl.Application.Run()
}

func initComponents() {
	mainForm := lcl.Application.CreateForm()
	// 先禁用对齐
	mainForm.DisableAlign()
	defer mainForm.EnableAlign()

	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(500)

	dlgOpen := lcl.NewOpenDialog(mainForm)
	dlgOpen.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	//    dlgOpen.SetInitialDir()
	//	dlgOpen.SetFilterIndex()

	dlgOpen.SetOptions(dlgOpen.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	dlgOpen.SetTitle("打开")

	btn := lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Open Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlgOpen.Execute() {
			fmt.Println("filename: ", dlgOpen.FileName())
		}
	})

	dlSave := lcl.NewSaveDialog(mainForm)
	dlSave.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	dlSave.SetOptions(dlSave.Options().Include(types.OfShowHelp))
	dlSave.SetTitle("保存")

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Save Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlSave.Execute() {
			fmt.Println("filename: ", dlSave.FileName())
		}
	})

	dlFont := lcl.NewFontDialog(mainForm)

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Font Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlFont.Execute() {
			fmt.Println("Name: ", dlFont.Font().Name())
		}
	})

	dlColor := lcl.NewColorDialog(mainForm)
	dlColor.SetOnClose(func(sender lcl.IObject) {
		fmt.Println("color close")
	})
	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Color Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlColor.Execute() {
			fmt.Println("Color: ", dlColor.Color())
		}
	})

	dlPicOpen := lcl.NewOpenPictureDialog(mainForm)
	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("OpenPic Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlPicOpen.Execute() {
			fmt.Println("Name: ", dlPicOpen.FileName())
		}
	})

	dlPicSave := lcl.NewSavePictureDialog(mainForm)
	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SavePic Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlPicSave.Execute() {
			fmt.Println("Name: ", dlPicSave.FileName())
		}
	})

	dlSelDirdlg := lcl.NewSelectDirectoryDialog(mainForm)
	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Select Directory Dialog")
	btn.SetOnClick(func(lcl.IObject) {
		if dlSelDirdlg.Execute() {
			fmt.Println("Name: ", dlSelDirdlg.FileName())
		}
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SelectDirectory1")
	btn.SetOnClick(func(lcl.IObject) {
		if ok, dir := lcl.SelectDirectory1(0); ok {
			fmt.Println("选择的目录为：", dir)
		}
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SelectDirectory2")
	btn.SetOnClick(func(lcl.IObject) {
		if ok, dir := lcl.SelectDirectory2("标题了", "C:/", true); ok {
			fmt.Println("选择的目录为：", dir)
		}
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("FindDialog")
	findDialog := lcl.NewFindDialog(mainForm)
	findDialog.SetOnFind(func(sender lcl.IObject) {
		fmt.Println("FindText: ", findDialog.FindText())
		opt := findDialog.Options()
		/*
			FrDown = iota + 0
			FrFindNext
			FrHideMatchCase
			FrHideWholeWord
			FrHideUpDown
			FrMatchCase
			FrDisableMatchCase
			FrDisableUpDown
			FrDisableWholeWord
			FrReplace
			FrReplaceAll
			FrWholeWord
			FrShowHelp
		*/
		if opt.In(types.FrDown) {
			fmt.Println("向下")
		} else {
			fmt.Println("向上")
		}
		if opt.In(types.FrFindNext) {
			fmt.Println("查找下一个")
		}
		if opt.In(types.FrMatchCase) {
			fmt.Println("区分大小写")
		}
	})
	btn.SetOnClick(func(lcl.IObject) {
		findDialog.Execute()
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("ReplaceDialog")
	replaceDialog := lcl.NewReplaceDialog(mainForm)
	replaceDialog.SetOnFind(func(sender lcl.IObject) {
		fmt.Println("FindText:", replaceDialog.FindText(), ", Relpace: ", replaceDialog.ReplaceText())
		opt := replaceDialog.Options()
		/*
			FrDown = iota + 0
			FrFindNext
			FrHideMatchCase
			FrHideWholeWord
			FrHideUpDown
			FrMatchCase
			FrDisableMatchCase
			FrDisableUpDown
			FrDisableWholeWord
			FrReplace
			FrReplaceAll
			FrWholeWord
			FrShowHelp
		*/
		if opt.In(types.FrDown) {
			fmt.Println("向下")
		} else {
			fmt.Println("向上")
		}
		if opt.In(types.FrFindNext) {
			fmt.Println("查找下一个")
		}
		if opt.In(types.FrMatchCase) {
			fmt.Println("区分大小写")
		}
	})

	replaceDialog.SetOnReplace(func(sender lcl.IObject) {
		opt := replaceDialog.Options()
		if opt.In(types.FrReplaceAll) {
			fmt.Println("替换全部")
		}
		if opt.In(types.FrReplace) {
			fmt.Println("替换一次")
		}
		fmt.Println("替换字符：", replaceDialog.ReplaceText())
	})

	btn.SetOnClick(func(lcl.IObject) {
		replaceDialog.Execute()
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("InputBox")
	btn.SetOnClick(func(lcl.IObject) {
		s := lcl.InputBox("标题", "提示", "默认值")
		if s != "" {
			fmt.Println("结果：", s)
		}
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("InpuQuery")
	btn.SetOnClick(func(lcl.IObject) {
		s := "default"
		if lcl.InputQuery("标题", "提示", &s) {
			fmt.Println("结果：", s)
		}
	})

	dlPrinterSetupDialog := lcl.NewPrinterSetupDialog(mainForm)
	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("PrinterSetupDialog")
	btn.SetOnClick(func(lcl.IObject) {
		dlPrinterSetupDialog.Execute()
	})

	dlPageSetupDialog := lcl.NewPageSetupDialog(mainForm)
	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("PageSetupDialog")
	btn.SetOnClick(func(lcl.IObject) {
		dlPageSetupDialog.Execute()
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("PasswordBox")
	btn.SetOnClick(func(lcl.IObject) {
		fmt.Println(lcl.PasswordBox("输入", "请输入密码："))
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("InputCombo")
	// +strings.Repeat(" ", 50) 是因为显示的窗口大小会根据`aPrompt`这个计算宽度
	btn.SetOnClick(func(lcl.IObject) {
		fmt.Println(lcl.InputCombo("选择", "请选择一项："+strings.Repeat(" ", 50), []string{"第一项", "第二项", "第三项", "第四项"}))
	})

	btn = lcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("InputComboEx")
	btn.SetOnClick(func(lcl.IObject) {
		fmt.Println(lcl.InputComboEx("选择", "请选择一项："+strings.Repeat(" ", 50), []string{"第一项", "第二项", "第三项", "第四项"}, false))
	})
}
