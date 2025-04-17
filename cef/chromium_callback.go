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
	"github.com/cyber-xxm/energy/v2/cef/i18n"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
)

// 事件处理函数返回true将不继续执行
func chromiumOnAfterCreate(window IBrowserWindow, browser *ICefBrowser) bool {
	if common.IsWindows() {
		rtl.SendMessage(browser.HostWindowHandle(), messages.WM_SETICON, 1, lcl.Application.Icon().Handle())
	}
	// 浏览器创建完之后
	main := BrowserWindow.MainWindow()
	if browser.Identifier() == main.Id() && main.IsLCL() {
		//更新windowParent位置和大小
		wp := main.AsLCLBrowserWindow().WindowParent()
		if wp.Align() != types.AlClient {
			rect := wp.BoundsRect()
			rect.Left, rect.Top = wp.point()
			rect.SetSize(wp.size())
			wp.SetBoundsRect(rect)
		}
		wp.UpdateSize()
	}
	// 当前应用是LCL窗口预先创建下一个window
	if !application.IsMessageLoop() && window.Chromium().Config().EnableWindowPopup() {
		RunOnMainThread(func() {
			BrowserWindow.createNextLCLPopupWindow()
		})
	}
	if !common.IsWindows() { // LCL: Linux, MacOS 需要在这里加载URL
		localLoadRes.loadDefaultURL(window, browser) // TODO
	}
	// 方式二 本地资源加载处理器
	localLoadRes.getSchemeHandlerFactory(window, browser) // TODO
	return false
}

// chromiumOnBeforeBrowser
func chromiumOnBeforeBrowser(window IBrowserWindow, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) {
	// 辅助工具不具有浏览器窗口特性
	if window.WindowType() == consts.WT_DEV_TOOLS || window.WindowType() == consts.WT_VIEW_SOURCE ||
		window.WindowProperty().WindowType == consts.WT_DEV_TOOLS || window.WindowProperty().WindowType == consts.WT_VIEW_SOURCE {
		return
	}
	BrowserWindow.PutWindowInfo(browser, window)
	// 方式二 本地资源加载处理器
	//localLoadRes.getSchemeHandlerFactory(window, browser) // TODO
}

// chromium 关闭之前
func chromiumOnBeforeClose(m IBrowserWindow, browser *ICefBrowser) {
	// 移除当前关闭所在的集合窗口维护
	BrowserWindow.removeWindowInfo(browser.Identifier())
}

var (
	refreshId, forcedRefreshId                                 consts.MenuId
	devToolsId, viewSourceId, closeBrowserId                   consts.MenuId
	backId, forwardId, printId                                 consts.MenuId
	undoId, redoId, cutId, copyId, pasteId, delId, selectAllId consts.MenuId
	imageUrlId, copyImageId, imageSaveId, aUrlId               consts.MenuId
)

// chromiumOnBeforeContextMenu 右键菜单 - 默认实现
func chromiumOnBeforeContextMenu(currentWindow IBrowserWindow, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
	if !currentWindow.Chromium().Config().EnableMenu() {
		model.Clear()
		return
	}
	//开发者工具和显示源代码不展示自定义默认菜单
	if currentWindow.WindowType() == consts.WT_DEV_TOOLS || currentWindow.WindowType() == consts.WT_VIEW_SOURCE {
		return
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
			Text:        i18n.Resource("undo"),
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
			Text:        i18n.Resource("redo"),
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
			Text:        i18n.Resource("cut"),
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
			Text:        i18n.Resource("copy"),
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
			Text:        i18n.Resource("paste"),
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
			Text:        i18n.Resource("selectAll"),
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
	isLink := params.TypeFlags()&consts.CM_TYPEFLAG_LINK == consts.CM_TYPEFLAG_LINK
	isCopyLink := params.MediaType() == consts.CM_MEDIATYPE_NONE && isLink
	if isCopyLink {
		aUrlId = model.CefMis.NextCommandId()
		model.AddItem(aUrlId, i18n.Resource("copyLink"))
	}
	isCopyImage := params.MediaType() == consts.CM_MEDIATYPE_IMAGE && isLink
	if isCopyImage {
		//copyImageId = model.CefMis.NextCommandId()
		//model.AddItem(copyImageId, "复制图片")
		imageUrlId = model.CefMis.NextCommandId()
		model.AddItem(imageUrlId, i18n.Resource("copyImageLink"))
		imageSaveId = model.CefMis.NextCommandId()
		model.AddItem(imageSaveId, i18n.Resource("imageSaveAs"))
	}
	if isCopyLink || isCopyImage {
		model.AddSeparator()
	}
	backId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   backId,
		Text:        i18n.Resource("back"),
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
		Text:        i18n.Resource("forward"),
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
		Text:        i18n.Resource("print"),
		Accelerator: "ctrl+p",
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			browser.Print()
		},
	})
	model.AddSeparator()
	closeBrowserId = model.CefMis.NextCommandId()
	model.AddItem(closeBrowserId, i18n.Resource("closeBrowser"))
	model.AddSeparator()
	refreshId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   refreshId,
		Text:        i18n.Resource("refresh"),
		Accelerator: "ctrl+r",
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			browser.Reload()
		},
	})
	forcedRefreshId = model.CefMis.NextCommandId()
	model.AddMenuItem(&MenuItem{
		CommandId:   forcedRefreshId,
		Text:        i18n.Resource("forcedRefresh"),
		Accelerator: "shift+ctrl+r",
		Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
			browser.ReloadIgnoreCache()
		},
	})
	model.AddSeparator()
	if currentWindow.Chromium().Config().EnableViewSource() {
		viewSourceId = model.CefMis.NextCommandId()
		model.AddMenuItem(&MenuItem{
			CommandId:   viewSourceId,
			Text:        i18n.Resource("viewPageSource"),
			Accelerator: "ctrl+u",
			Callback: func(browser *ICefBrowser, commandId consts.MenuId, params *ICefContextMenuParams, menuType consts.TCefContextMenuType, eventFlags uint32, result *bool) {
				browser.ViewSource(currentWindow)
			},
		})
	}
	if currentWindow.Chromium().Config().EnableDevTools() {
		devToolsId = model.CefMis.NextCommandId()
		model.AddItem(devToolsId, i18n.Resource("devTools"))
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

// 右键菜单 - 默认实现
func chromiumOnContextMenuCommand(currentWindow IBrowserWindow, currentChromium ICEFChromiumBrowser, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32) bool {
	browserId := browser.Identifier()
	defer func() {
		if err := recover(); err != nil {
			logger.Error("OnContextMenuCommand Error:", err, "browserId:", browserId)
		}
	}()
	if commandId == backId {
		browser.GoBack()
	} else if commandId == forwardId {
		browser.GoForward()
	} else if commandId == printId {
		browser.Print()
	} else if commandId == closeBrowserId {
		currentWindow.CloseBrowserWindow()
	} else if commandId == refreshId {
		browser.Reload()
	} else if commandId == forcedRefreshId {
		browser.ReloadIgnoreCache()
	} else if commandId == viewSourceId {
		if currentChromium.Chromium().Config().EnableViewSource() {
			browser.ViewSource(currentWindow)
		}
	} else if commandId == devToolsId {
		if currentChromium.Chromium().Config().EnableDevTools() {
			browser.ShowDevTools(currentWindow, currentChromium)
		}
	} else if commandId == aUrlId {
		lcl.Clipboard.SetAsText(params.LinkUrl())
	} else if commandId == copyImageId {
		frame.Copy()
	} else if commandId == imageUrlId {
		lcl.Clipboard.SetAsText(params.SourceUrl())
	} else if commandId == imageSaveId {
		browser.StartDownload(params.SourceUrl())
	}
	return true
}
