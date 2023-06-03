package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"strings"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	var app = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/sysdialog.html"
	cef.BrowserWindow.Config.Title = "Energy - sysdialog"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 系统消息提示框目前仅能在LCL窗口组件下使用
		// LCL 各种系统组件需要在UI线程中执行, 但ipc.on非UI线程
		// 所以需要使用 QueueAsyncCall 包裹在UI线程中执行
		if window.IsLCL() {
			// window from
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

			dlPicSave := lcl.NewSavePictureDialog(bw)

			dlPicOpen := lcl.NewOpenPictureDialog(bw)

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
				cef.QueueAsyncCall(func(id int) {
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
						if dlPicSave.Execute() {
							fmt.Println("Name: ", dlPicSave.FileName())
						}
					case 14: // OpenPic Dialog
						if dlPicOpen.Execute() {
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
		}
	})

	//运行应用
	cef.Run(app)
}
