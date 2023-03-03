//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// chromium event 默认事件实现
package cef

import (
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types/messages"
)

// browserProcessMessageReceived 主进程消息 - 默认实现
func browserProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
	fmt.Println("browserProcessMessageReceived", message.Name(), message.ArgumentList().Size())
	return false
}

// chromiumOnAfterCreate 事件处理函数返回true将不继续执行
func chromiumOnAfterCreate(browser *ICefBrowser) bool {
	if common.IsWindows() {
		rtl.SendMessage(browser.HostWindowHandle(), messages.WM_SETICON, 1, lcl.Application.Icon().Handle())
	}
	return false
}

// chromiumOnBeforeBrowser
func chromiumOnBeforeBrowser(browser *ICefBrowser, frame *ICefFrame) {
	if BrowserWindow.popupWindow != nil {
		if browser.Identifier() < BrowserWindow.GetNextWindowNum() {
			return
		}
	}
	BrowserWindow.setOrIncNextWindowNum(browser.Identifier() + 1)
	if !consts.IsMessageLoop {
		QueueAsyncCall(func(id int) {
			BrowserWindow.createNextLCLPopupWindow()
		})
	}
}

// chromiumOnBeforeClose - chromium 关闭之前
func chromiumOnBeforeClose(browser *ICefBrowser) {
	//if ipc.IPC.Render() != nil && !SingleProcess && processName != PT_DEVTOOLS {
	//	ipc.IPC.Render().Close()
	//}
}

// chromiumOnFrameDetached
func chromiumOnFrameDetached(browser *ICefBrowser, frame *ICefFrame) {
}

var (
	refreshId, forcedRefreshId                                 consts.MenuId
	devToolsId, viewSourceId, closeBrowserId                   consts.MenuId
	backId, forwardId, printId                                 consts.MenuId
	undoId, redoId, cutId, copyId, pasteId, delId, selectAllId consts.MenuId
	imageUrlId, copyImageId, imageSaveId, aUrlId               consts.MenuId
)

// chromiumOnBeforeContextMenu 右键菜单 - 默认实现
func chromiumOnBeforeContextMenu(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
	if !api.GoBool(BrowserWindow.Config.chromiumConfig.enableMenu) {
		model.Clear()
		return
	}
	if winInfo := BrowserWindow.GetWindowInfo(browser.Identifier()); winInfo != nil {
		//开发者工具和显示源代码不展示自定义默认菜单
		if winInfo.WindowType() == consts.WT_DEV_TOOLS || winInfo.WindowType() == consts.WT_VIEW_SOURCE {
			return
		}
	}
	undoVisible, undoEnabled := model.IsVisible(consts.MENU_ID_UNDO), model.IsEnabled(consts.MENU_ID_UNDO)
	redoVisible, redoEnabled := model.IsVisible(consts.MENU_ID_REDO), model.IsEnabled(consts.MENU_ID_REDO)
	cutVisible, cutEnabled := model.IsVisible(consts.MENU_ID_CUT), model.IsEnabled(consts.MENU_ID_CUT)
	copyVisible, copyEnabled := model.IsVisible(consts.MENU_ID_COPY), model.IsEnabled(consts.MENU_ID_COPY)
	pasteVisible, pasteEnabled := model.IsVisible(consts.MENU_ID_PASTE), model.IsEnabled(consts.MENU_ID_PASTE)
	selectAllVisible, selectAllEnabled := model.IsVisible(consts.MENU_ID_SELECT_ALL), model.IsEnabled(consts.MENU_ID_SELECT_ALL)
	model.Clear()
	if undoVisible {
		undoId = consts.MENU_ID_UNDO //model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   undoId,
			Text:        "撤销",
			Accelerator: "ctrl+z",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.GetFocusedFrame().Undo()
			},
		})
		model.SetEnabled(undoId, undoEnabled)
	}
	if redoVisible {
		redoId = consts.MENU_ID_REDO //model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   redoId,
			Text:        "恢复",
			Accelerator: "ctrl+shift+z",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.GetFocusedFrame().Redo()
			},
		})
		model.SetEnabled(redoId, redoEnabled)
	}
	if undoVisible && redoVisible {
		model.AddSeparator()
	}
	if cutVisible {
		cutId = consts.MENU_ID_CUT //model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   cutId,
			Text:        "剪切",
			Accelerator: "ctrl+x",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.GetFocusedFrame().Cut()
			},
		})
		model.SetEnabled(cutId, cutEnabled)
	}
	if copyVisible {
		copyId = consts.MENU_ID_COPY //model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   copyId,
			Text:        "复制",
			Accelerator: "ctrl+c",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.GetFocusedFrame().Copy()
			},
		})
		model.SetEnabled(copyId, copyEnabled)
	}
	if pasteVisible {
		pasteId = consts.MENU_ID_PASTE //model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   pasteId,
			Text:        "粘贴",
			Accelerator: "ctrl+v",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.GetFocusedFrame().Paste()
			},
		})
		model.SetEnabled(pasteId, pasteEnabled)
	}
	if selectAllVisible {
		selectAllId = consts.MENU_ID_SELECT_ALL //model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   selectAllId,
			Text:        "全选",
			Accelerator: "ctrl+a",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.GetFocusedFrame().SelectAll()
			},
		})
		model.SetEnabled(pasteId, selectAllEnabled)
	}
	if cutVisible && copyVisible && pasteVisible && selectAllVisible {
		model.AddSeparator()
	}
	//A标签和图片 链接
	if params.TypeFlags == 5 && params.MediaType == consts.CM_MEDIATYPE_NONE { //a=5
		aUrlId = model.CefMis.NextCommandId()
		model.AddItem(aUrlId, "复制链接")
	}
	if params.TypeFlags == 9 && params.MediaType == consts.CM_MEDIATYPE_IMAGE { // image=9
		//copyImageId = model.CefMis.NextCommandId()
		//model.AddItem(copyImageId, "复制图片")
		imageUrlId = model.CefMis.NextCommandId()
		model.AddItem(imageUrlId, "复制图片链接")
		imageSaveId = model.CefMis.NextCommandId()
		model.AddItem(imageSaveId, "图片另存为")
	}
	if (params.TypeFlags == 5 && params.MediaType == consts.CM_MEDIATYPE_NONE) || params.TypeFlags == 9 && params.MediaType == consts.CM_MEDIATYPE_IMAGE {
		model.AddSeparator()
	}
	//A标签和图片 链接
	backId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   backId,
		Text:        "返回",
		Accelerator: "alt+" + string(rune(37)),
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			if browser.CanGoBack() {
				browser.GoBack()
			}
		},
	})
	forwardId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   forwardId,
		Text:        "前进",
		Accelerator: "alt+" + string(rune(39)),
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			if browser.CanGoForward() {
				browser.GoForward()
			}
		},
	})
	model.AddSeparator()
	printId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   printId,
		Text:        "打印",
		Accelerator: "ctrl+p",
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			browser.Print()
		},
	})
	model.AddSeparator()
	closeBrowserId = model.CefMis.NextCommandId()
	model.AddItem(closeBrowserId, "关闭页面")
	model.AddSeparator()
	refreshId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   refreshId,
		Text:        "刷新",
		Accelerator: "ctrl+r",
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			browser.Reload()
		},
	})
	forcedRefreshId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   forcedRefreshId,
		Text:        "强制刷新",
		Accelerator: "shift+ctrl+r",
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			browser.ReloadIgnoreCache()
		},
	})
	model.AddSeparator()
	if api.GoBool(BrowserWindow.Config.chromiumConfig.enableViewSource) {
		viewSourceId = model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   viewSourceId,
			Text:        "查看网面源代码",
			Accelerator: "ctrl+u",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.ViewSource()
			},
		})
	}
	if api.GoBool(BrowserWindow.Config.chromiumConfig.enableDevTools) {
		devToolsId = model.CefMis.NextCommandId()
		model.AddItem(devToolsId, "开发者工具(F12)")
	}
	if browser.CanGoBack() {
		model.SetEnabled(backId, true)
	} else {
		model.SetEnabled(backId, false)
	}
	if browser.CanGoForward() {
		model.SetEnabled(forwardId, true)
	} else {
		model.SetEnabled(forwardId, false)
	}
}

// chromiumOnContextMenuCommand 右键菜单 - 默认实现
func chromiumOnContextMenuCommand(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("OnContextMenuCommand Error:", err)
		}
	}()
	*result = true
	if commandId == backId {
		browser.GoBack()
	} else if commandId == forwardId {
		browser.GoForward()
	} else if commandId == printId {
		browser.Print()
	} else if commandId == closeBrowserId {
		winInfo := BrowserWindow.GetWindowInfo(browser.Identifier())
		winInfo.CloseBrowserWindow()
	} else if commandId == refreshId {
		browser.Reload()
	} else if commandId == forcedRefreshId {
		browser.ReloadIgnoreCache()
	} else if commandId == viewSourceId {
		if api.GoBool(BrowserWindow.Config.chromiumConfig.enableViewSource) {
			browser.ViewSource()
		}
	} else if commandId == devToolsId {
		if api.GoBool(BrowserWindow.Config.chromiumConfig.enableDevTools) {
			browser.ShowDevTools()
		}
	} else if commandId == aUrlId {
		lcl.Clipboard.SetAsText(params.LinkUrl)
	} else if commandId == copyImageId {
		frame.Copy()
	} else if commandId == imageUrlId {
		lcl.Clipboard.SetAsText(params.SourceUrl)
	} else if commandId == imageSaveId {
		browser.StartDownload(params.SourceUrl)
	}
	*result = true
}
