package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "fs://energy"
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         resources,
	}.Build())
	//在主窗口初始化监听浏览器事件
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		var (
			menuId01           consts.MenuId
			menuId02           consts.MenuId
			menuId03           consts.MenuId
			menuId0301         consts.MenuId
			menuId0302         consts.MenuId
			menuIdCheck        consts.MenuId
			isMenuIdCheck      = true
			menuIdEnable       consts.MenuId
			isMenuIdEnable     = true
			menuIdEnableCtl    consts.MenuId
			menuIdRadio101     consts.MenuId
			menuIdRadio102     consts.MenuId
			menuIdRadio103     consts.MenuId
			radioDefault1Check consts.MenuId
			menuIdRadio201     consts.MenuId
			menuIdRadio202     consts.MenuId
			menuIdRadio203     consts.MenuId
			radioDefault2Check consts.MenuId
		)
		//右键弹出菜单
		event.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, model *cef.ICefMenuModel, window cef.IBrowserWindow) bool {
			//既然是自定义，当然要去除之前事先定义好的
			model.Clear()
			//开始创建菜单，每个菜单项都有自己的ID, 所以要先定义一个能保存这些菜单项的ID的变量
			fmt.Printf("%+v\n", params)
			//注意： 每个菜单项的ID有固定的 ID 生成函数
			//获取一个菜单项ID
			menuId01 = model.CefMis.NextCommandId()
			model.AddItem(menuId01, "菜单一 html 文字变红色")
			menuId02 = model.CefMis.NextCommandId()
			model.AddItem(menuId02, "菜单二 html 文字变绿色")
			menuId03 = model.CefMis.NextCommandId()
			menu03 := model.AddSubMenu(menuId03, "菜单三 带有子菜单")
			menuId0301 = model.CefMis.NextCommandId()
			menu03.AddItem(menuId0301, "菜单三的子菜单一 ")
			menuId0302 = model.CefMis.NextCommandId()
			menu03.AddItem(menuId0302, "菜单三的子菜单二")
			model.AddSeparator()
			//check
			menuIdCheck = model.CefMis.NextCommandId()
			model.AddCheckItem(menuIdCheck, "这是一个checkItem-好像就windows有效")
			model.SetChecked(menuIdCheck, isMenuIdCheck)
			//enable
			model.AddSeparator()
			menuIdEnable = model.CefMis.NextCommandId()
			if isMenuIdEnable {
				model.AddItem(menuIdEnable, "菜单-已启用")
				model.SetColor(menuIdEnable, consts.CEF_MENU_COLOR_TEXT, cef.CefColorSetARGB(0, 111, 12, 200))
			} else {
				model.AddItem(menuIdEnable, "菜单-已禁用")
			}
			model.SetEnabled(menuIdEnable, isMenuIdEnable)
			menuIdEnableCtl = model.CefMis.NextCommandId()
			model.AddItem(menuIdEnableCtl, "启用上面菜单")
			//为什么要用Visible而不是不创建这个菜单? 因为菜单项的ID是动态的啊。
			model.SetVisible(menuIdEnableCtl, !isMenuIdEnable)
			if !isMenuIdEnable {
				model.SetColor(menuIdEnableCtl, consts.CEF_MENU_COLOR_TEXT, cef.CefColorSetARGB(0, 222, 111, 0))
			}
			model.AddSeparator()
			//radio 1组
			menuIdRadio101 = model.CefMis.NextCommandId()
			menuIdRadio102 = model.CefMis.NextCommandId()
			menuIdRadio103 = model.CefMis.NextCommandId()
			model.AddRadioItem(menuIdRadio101, "单选按钮 1 1组", 1001)
			model.AddRadioItem(menuIdRadio102, "单选按钮 2 1组", 1001)
			model.AddRadioItem(menuIdRadio103, "单选按钮 3 1组", 1001)
			if radioDefault1Check == 0 {
				radioDefault1Check = menuIdRadio101
			}
			model.SetChecked(radioDefault1Check, true)
			model.AddSeparator()
			//radio 2组
			menuIdRadio201 = model.CefMis.NextCommandId()
			menuIdRadio202 = model.CefMis.NextCommandId()
			menuIdRadio203 = model.CefMis.NextCommandId()
			model.AddRadioItem(menuIdRadio201, "单选按钮 1 2组", 1002)
			model.AddRadioItem(menuIdRadio202, "单选按钮 2 2组", 1002)
			model.AddRadioItem(menuIdRadio203, "单选按钮 3 2组", 1002)
			if radioDefault2Check == 0 {
				radioDefault2Check = menuIdRadio201
			}
			model.SetChecked(radioDefault2Check, true)
			return true
		})
		//右键菜单项命令
		event.SetOnContextMenuCommand(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, menuId consts.MenuId, eventFlags uint32, window cef.IBrowserWindow) bool {
			fmt.Printf("params: %+v\n", params)
			fmt.Println("menuId: ", menuId, eventFlags)
			//在这里处理某个菜单项的点击事件所触发的命令，这里的命令对应着一个菜单项的ID
			var clickMenuId = 0
			switch menuId {
			case menuId01:
				clickMenuId = 1
			case menuId02:
				clickMenuId = 2
			case menuIdEnable:
				isMenuIdEnable = !isMenuIdEnable
			case menuIdCheck:
				isMenuIdCheck = !isMenuIdCheck
			case menuIdEnableCtl:
				isMenuIdEnable = true
			case menuIdRadio101, menuIdRadio102, menuIdRadio103:
				radioDefault1Check = menuId
			case menuIdRadio201, menuIdRadio202, menuIdRadio203:
				radioDefault2Check = menuId
			}
			ipc.Emit("menu", clickMenuId, fmt.Sprintf("菜单 %d 随便传点什么吧 但是，字符串参数字符串参数字符串参数字符串参数字符串参数字符串参数字符串参数.", menuId))
			//*result = true
			return true
		})
	})
	//运行应用
	cef.Run(cefApp)
}
