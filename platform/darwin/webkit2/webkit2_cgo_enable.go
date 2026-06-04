//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package webkit2

import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/energy/v3/platform/darwin/webkit2/cgo"
	"unsafe"
)

func AsWkWebView(ptr unsafe.Pointer) IWkWebView {
	return cgo.AsWkWebView(ptr)
}
