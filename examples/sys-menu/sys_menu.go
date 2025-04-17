package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	//创建应用
	cefApp := cef.NewApplication()
	cefApp.SetUseMockKeyChain(true)
	//cefApp.SetRootCache(rootCache)
	//cefApp.SetCache(filepath.Join(rootCache, "cache"))
	//主窗口的配置
	//
	cef.BrowserWindow.Config.Url = "fs://energy"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	cef.BrowserWindow.Config.Title = "SysMenu"
	//cef.BrowserWindow.Config.EnableHideCaption = true
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		Home:       "sys_menu.html",
		ResRootDir: "resources",
		FS:         demoCommon.ResourcesFS(),
	}.Build())
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 在窗口初始化时重置窗口布局 默认仅有CEFWindowParent
		// 在这里重新指定 CEFWindowParent 的父组件, 默认是主窗口
		// 仅 lcl 窗口
		if window.IsLCL() {
			// 主窗口
			//bw := window.AsLCLBrowserWindow().BrowserWindow()
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			// 拿到CEFWindowParent
			//windowParent := bw.WindowParent()
			// 恢复四角定位， 因为默认整个窗口自动调整大小
			//windowParent.RevertCustomAnchors()

			// 系统菜单
			// 开始创建菜单
			createSysMenu(bw)

			// html 内容
			// 重新设置父组件
			// 创建contentBox panel组件
			//contentBox := lcl.NewPanel(bw)
			//contentBox.SetParent(bw)            //这一步不能少
			//contentBox.SetAlign(types.AlClient) // 根据客户区自动适应大小, 这块是panel的宽高根据主窗口自动调整大小

			//重新指定 CEFWindowParent 的父组件
			//windowParent.SetParent(contentBox)
			// 重新设置四角定位，自动调整大小
			//windowParent.DefaultAnchors()
		}
	})
	//运行应用
	cef.Run(cefApp)
}

// 创建菜单
func createSysMenu(owner lcl.IComponent) {
	//  菜单通过 MainMenu popupMenu、menuitem组合
	// 先创建 主菜单, 有很一些事件，也可以自己绘制菜单效果通过 canvas
	mainMenu := lcl.NewMainMenu(owner)
	mainMenu.SetOnMeasureItem(func(sender lcl.IObject, aCanvas *lcl.TCanvas, width, height *int32) {
	})
	// 创建一级菜单
	fileClassA := lcl.NewMenuItem(owner)
	fileClassA.SetCaption("文件(&F)") //菜单名称 alt + f
	aboutClassA := lcl.NewMenuItem(owner)
	aboutClassA.SetCaption("关于(&A)")

	// 把一及菜单添加到主菜单
	mainMenu.Items().Add(fileClassA)
	mainMenu.Items().Add(aboutClassA)

	var createMenuItem = func(label, shortCut string, click func(lcl.IObject)) (result *lcl.TMenuItem) {
		result = lcl.NewMenuItem(owner)
		result.SetCaption(label)               //菜单项显示的文字
		result.SetShortCutFromString(shortCut) // 快捷键
		result.SetOnClick(click)               // 触发事件，回调函数
		return
	}

	// 给一级菜单添加菜单项
	createItem := createMenuItem("新建(&N)", "Ctrl+N", func(lcl.IObject) {
		fmt.Println("单击了新建")
	})
	fileClassA.Add(createItem) // 把创建好的菜单项添加到 第一个菜单中
	openItem := createMenuItem("打开(&O)", "Ctrl+O", func(lcl.IObject) {
		fmt.Println("单击了打开")
	})
	fileClassA.Add(openItem) // 把创建好的菜单项添加到 第一个菜单中
	// 分割线
	separate := lcl.NewMenuItem(owner)
	separate.SetCaption("-")
	fileClassA.Add(separate) // 把创建好的菜单项添加到 第一个菜单中
	//二级菜单
	twoLevelMenu := lcl.NewMenuItem(owner)
	twoLevelMenu.SetCaption("二级菜单")

	// 给二级菜单添加菜单
	twoLevelSubOneMenu := createMenuItem("二级子菜单一", "", func(object lcl.IObject) {
		lcl.ShowMessage("点击了 二级子菜单一, 提示框样式不好看，是因为没生成 syso 主题文件")
	})
	twoLevelMenu.Add(twoLevelSubOneMenu)
	twoLevelSubTwoMenu := createMenuItem("二级子菜单二", "", func(object lcl.IObject) {
		lcl.ShowMessage("点击了 二级子菜单二")
	})
	twoLevelMenu.Add(twoLevelSubTwoMenu)
	fileClassA.Add(twoLevelMenu) // 把创建好的菜单项添加到 第一个菜单中
	// 分割线
	separate = lcl.NewMenuItem(owner)
	separate.SetCaption("-")
	fileClassA.Add(separate)

	// 退出
	exit := createMenuItem("退出(&E)", "Ctrl+Q", func(object lcl.IObject) {
		//window.CloseBrowserWindow()
	})
	fileClassA.Add(exit)
	// help
	help := createMenuItem("帮助(&H)", "", func(object lcl.IObject) {
		if lcl.Application.MessageBox("你确定需要帮助吗？", "消息框标题", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == types.IdOK {
			lcl.ShowMessage("你点击了“是")
		}
	})
	aboutClassA.Add(help)

	// Mac 偏好设置和关于菜单
	if common.IsDarwin() {
		// https://wiki.lazarus.freepascal.org/Mac_Preferences_and_About_Menu
		// 动态添加的，静态好像是通过设计器将顶级的菜单标题设置为应用程序名，但动态的就是另一种方式
		appMenu := lcl.NewMenuItem(owner)
		// 动态添加的，设置一个Unicode Apple logo char
		appMenu.SetCaption(types.AppleLogoChar)
		subItem := lcl.NewMenuItem(owner)

		subItem.SetCaption("关于")
		subItem.SetOnClick(func(sender lcl.IObject) {
			lcl.ShowMessage("About")
		})
		appMenu.Add(subItem)

		subItem = lcl.NewMenuItem(owner)
		subItem.SetCaption("-")
		appMenu.Add(subItem)

		subItem = lcl.NewMenuItem(owner)
		subItem.SetCaption("首选项...")
		subItem.SetShortCutFromString("Meta+,")
		subItem.SetOnClick(func(sender lcl.IObject) {
			lcl.ShowMessage("Preferences")
		})
		appMenu.Add(subItem)
		// 添加
		mainMenu.Items().Insert(0, appMenu)
	}
}
