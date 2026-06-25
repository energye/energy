//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package core

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
)

type Browser interface {
}

type WindowParent interface {
}

type IBrowser interface {
	lcl.IComponent
	// SetWindow sets the window instance for webview and initializes relevant callback functions
	// window - Window interface instance used to host webview content
	SetWindow(window window.IWindow)
	// SetLocalLoad sets custom resource loading mode for current browser
	SetLocalLoad(localLoad application.LocalLoad)
	LocalLoadResource() *application.LocalLoadResource
	// UpdateBrowserOptions updates configuration of the current browser instance
	UpdateBrowserOptions()
	SetParent(window lcl.IWinControl)
	CreateBrowser()
	BrowserId() uint32
	WindowParent() WindowParent
	Browser() Browser
	SendMessage(payload []byte)
	Close()
	SetDefaultURL(url string)
	LoadURL(url string)
	ExecuteScript(javaScript string)
	ExecuteScriptCallback(script string, callback TOnEvaluateScriptCallbackEvent)
	SetOnBrowserAfterCreated(fn lcl.TNotifyEvent)
	SetOnResourceRequest(fn TOnResourceRequestEvent)
	SetOnProcessMessage(fn TOnProcessMessageEvent)
	SetOnLoadChange(fn TOnLoadChangeEvent)
	SetOnContextMenu(fn TOnContextMenuEvent)
	SetOnContextMenuCommand(fn TOnContextMenuCommandEvent)
	SetOnPopupWindow(fn TOnPopupWindowEvent)
	SetOnDragEnter(fn TOnDragEnterEvent)
	SetOnDragLeave(fn TOnDragLeaveEvent)
	SetOnDragOver(fn TOnDragOverEvent)
}
