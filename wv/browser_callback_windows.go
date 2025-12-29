// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build windows

package wv

import (
	"github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/lcl/lcl"
	wv "github.com/energye/wv/windows"
)

type TOnWebResourceRequestedEvent func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebResourceRequestedEventArgs) bool
type TOnWebResourceResponseReceivedEvent func(sender lcl.IObject, webview wv.ICoreWebView2, args wv.ICoreWebView2WebResourceResponseReceivedEventArgs) bool

type TOnWindowResize func(ht string)

type TOnWindowDrag func(message ipc.ProcessMessage)
