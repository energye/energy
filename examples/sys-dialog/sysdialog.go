package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"strings"
	//_ "net/http/pprof"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//创建应用
	var app = cef.NewApplication()
	app.SetUseMockKeyChain(true)
	// 这是使用LCL dialog和CEF dialog示例
	// 两者区别
	//   1. LCL 有更多的方式, CEF 仅有限几种方式
	//   2. LCL 仅适用于LCL窗口, CEF 适用LCL和VF窗口
	// 以下通过 ExternalMessagePump 和 MultiThreadedMessageLoop 区分当前所使用的窗口类型
	//app.SetExternalMessagePump(false)
	//app.SetMultiThreadedMessageLoop(false)
	if app.IsMessageLoop() { // IsMessageLoop VF window
		// CEF dialog 示例
		cef.BrowserWindow.Config.Url = "http://localhost:22022/sysdialog_cef.html"
	} else {
		// LCL dialog 示例
		cef.BrowserWindow.Config.Url = "http://localhost:22022/sysdialog.html"
	}
	cef.BrowserWindow.Config.Title = "Energy - dialog"

	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	// 在浏览器窗口初始化回调中注册IPC事件
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 系统消息提示框目前仅能在LCL窗口组件下使用
		// LCL 各种系统组件需要在UI线程中执行, 但ipc.on非UI线程
		// 所以需要使用 QueueAsyncCall 包裹在UI线程中执行
		if window.IsLCL() { // LCL window from
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			replaceDialog := lcl.NewReplaceDialog(bw)
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

			dlPageSetupDialog := lcl.NewPageSetupDialog(bw)

			dlPrinterSetupDialog := lcl.NewPrinterSetupDialog(bw)

			findDialog := lcl.NewFindDialog(bw)
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

			dlSelDirdlg := lcl.NewSelectDirectoryDialog(bw)

			var (
				dlPicSave *lcl.TSavePictureDialog
				dlPicOpen *lcl.TOpenPictureDialog
			)
			// VF 窗口，在linux不能使用它，还不知道是什么原因...
			if !common.IsLinux() && !app.IsMessageLoop() {
				dlPicSave = lcl.NewSavePictureDialog(bw)
				dlPicOpen = lcl.NewOpenPictureDialog(bw)
			}

			dlColor := lcl.NewColorDialog(bw)

			dlFont := lcl.NewFontDialog(bw)

			dlSave := lcl.NewSaveDialog(bw)
			dlSave.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
			dlSave.SetOptions(dlSave.Options().Include(types.OfShowHelp))
			dlSave.SetTitle("保存")

			dlgOpen := lcl.NewOpenDialog(bw)
			dlgOpen.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
			//    dlgOpen.SetInitialDir()
			//	dlgOpen.SetFilterIndex()
			dlgOpen.SetOptions(dlgOpen.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
			dlgOpen.SetTitle("打开")

			ipc.On("showDialog", func(t int) {
				window.RunOnMainThread(func() {
					fmt.Println("call-1-DMainThreadId:", api.DMainThreadId(), api.DCurrentThreadId())
					//cef.QueueAsyncCall(func(id int) {
					switch t {
					case 1: // InputComboEx
						fmt.Println(lcl.InputComboEx("选择", "请选择一项："+strings.Repeat(" ", 50), []string{"第一项", "第二项", "第三项", "第四项"}, false))
					case 2: // InputCombo
						fmt.Println(lcl.InputCombo("选择", "请选择一项："+strings.Repeat(" ", 50), []string{"第一项", "第二项", "第三项", "第四项"}))
					case 3: // PasswordBox
						fmt.Println(lcl.PasswordBox("输入", "请输入密码："))
					case 4: // PageSetup Dialog
						dlPageSetupDialog.Execute()
					case 5: // PrinterSetup Dialog
						dlPrinterSetupDialog.Execute()
					case 6: // InputQuery
						s := "default"
						if lcl.InputQuery("标题", "提示", &s) {
							fmt.Println("结果：", s)
						}
					case 7: // InputBox
						s := lcl.InputBox("标题", "提示", "默认值")
						if s != "" {
							fmt.Println("结果：", s)
						}
					case 8: // ReplaceDialog
						replaceDialog.Execute()
					case 9: // FindDialog
						findDialog.Execute()
					case 10: // SelectDirectory1
						if ok, dir := lcl.SelectDirectory1(0); ok {
							fmt.Println("选择的目录为：", dir)
						}
					case 11: // SelectDirectory2
						if ok, dir := lcl.SelectDirectory2("标题了", "C:/", true); ok {
							fmt.Println("选择的目录为：", dir)
						}
					case 12: // Select Directory Dialog
						if dlSelDirdlg.Execute() {
							fmt.Println("Name: ", dlSelDirdlg.FileName())
						}
					case 13: // SavePic Dialog
						if dlPicSave != nil && dlPicSave.Execute() {
							fmt.Println("Name: ", dlPicSave.FileName())
						}
					case 14: // OpenPic Dialog
						if dlPicOpen != nil && dlPicOpen.Execute() {
							fmt.Println("Name: ", dlPicOpen.FileName())
						}
					case 15: // Color Dialog
						if dlColor.Execute() {
							fmt.Println("Color: ", dlColor.Color())
						}
					case 16: // Font Dialog
						if dlFont.Execute() {
							fmt.Println("Name: ", dlFont.Font().Name())
						}
					case 17: // Save Dialog
						if dlSave.Execute() {
							fmt.Println("filename: ", dlSave.FileName())
						}
					case 18: // Open Dialog
						if dlgOpen.Execute() {
							fmt.Println("filename: ", dlgOpen.FileName())
						}
					}
				})
			})
		} else { // VF window
			// 可配合 Chromium().SetOnFileDialog 回调函数使用
			//window.Chromium().SetOnFileDialog(func(sender lcl.IObject, browser *cef.ICefBrowser, mode consts.FileDialogMode, title, defaultFilePath string, acceptFilters *lcl.TStrings, callback *cef.ICefFileDialogCallback) bool {
			//	fmt.Println("Chromium SetOnFileDialog", mode, title, defaultFilePath, "acceptFilters:", acceptFilters.Count())
			//	acceptFilters.Add(".png")
			//	callback.Cont([]string{"/file/to/path/file.xx"}) // 设置选择的文件
			//	return true
			//})
			// 定义 dialog 回调函数
			callback := cef.RunFileDialogCallbackRef.New()
			callback.SetOnFileDialogDismissed(func(filePaths *lcl.TStrings) {
				for i := 0; i < int(filePaths.Count()); i++ {
					path := filePaths.Strings(int32(i))
					fmt.Println(path)
				}
			})
			ipc.On("showDialog", func(t int) {
				switch t {
				case 1:
					window.Chromium().Browser().RunFileDialog(consts.FILE_DIALOG_OPEN_FOLDER, "打开文件夹", "", nil, callback)
				case 2:
					acceptFilters := lcl.NewStringList()
					acceptFilters.Add(".png")
					window.Chromium().Browser().RunFileDialog(consts.FILE_DIALOG_SAVE, "保存图片", "", acceptFilters, callback)
				case 3:
					acceptFilters := lcl.NewStringList()
					acceptFilters.Add(".png")
					window.Chromium().Browser().RunFileDialog(consts.FILE_DIALOG_OPEN, "打开图片", "", acceptFilters, callback)
				case 4:
					window.Chromium().Browser().RunFileDialog(consts.FILE_DIALOG_SAVE, "打开", "", nil, callback)
				case 5:
					window.Chromium().Browser().RunFileDialog(consts.FILE_DIALOG_OPEN, "保存", "", nil, callback)
				}
			})
		}
	})

	//运行应用
	cef.Run(app)
}
