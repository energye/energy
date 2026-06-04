//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux

package lcl

import (
	"github.com/energye/energy/v3/platform/linux/webkit2gtk"
	"github.com/energye/lcl/lcl"
	wvLinux "github.com/energye/wv/linux"
	"unsafe"
)

func isWebview(control lcl.IWinControl) IMenuEditing {
	if control == nil {
		return nil
	}
	clsName := control.ClassName()
	if clsName == "TWkWebviewParent" {
		windowParent := wvLinux.AsWkWebviewParent(control)
		if windowParent == nil {
			return nil
		}
		wvWebview := windowParent.Webview()
		if wvWebview == nil {
			return nil
		}
		wvData := wvWebview.WebView()
		if wvData == 0 {
			return nil
		}
		webview := webkit2gtk.AsWebkit2(unsafe.Pointer(wvData))
		return webview
	}
	return nil
}
