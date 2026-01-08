//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"sync/atomic"
)

// global browser id
var globalBrowserID uint32

// return next browser id
func getNextBrowserID() uint32 {
	atomic.AddUint32(&globalBrowserID, 1)
	return globalBrowserID
}

type IWebview interface {
	lcl.ICustomPanel
	SetWindow(window window.IWindow)
	SetBrowserOptions()
	SetParent(window lcl.IWinControl)
	CreateBrowser()
	BrowserId() uint32
	SendMessage(payload []byte)
	Close()
	SetDefaultURL(url string)
	LoadURL(url string)
	SetOnAfterCreated(fn lcl.TNotifyEvent)
	SetOnWindowClose(fn lcl.TCloseEvent)
	SetOnWindowShow(fn lcl.TNotifyEvent)
	SetOnWindowDestroy(fn lcl.TNotifyEvent)
	//SetAlign(v types.TAlign)
	SetWidth(v int32)
	SetHeight(v int32)
}

func TWebviewDesigner(owner lcl.IComponent) lcl.IPanel {
	m := lcl.NewPanel(owner)
	m.SetParentColor(true)
	m.SetParentDoubleBuffered(true)
	m.SetBevelInner(types.BvNone)
	m.SetBevelOuter(types.BvNone)
	return m
}

type TOnWindowResize func(ht string)
type TOnWindowDrag func(message ipc.ProcessMessage)
type TOnProcessMessageEvent func(message string)
type TOnResourceRequestEvent func(url, path, method string, header map[string]string) (resource string, ok bool)
