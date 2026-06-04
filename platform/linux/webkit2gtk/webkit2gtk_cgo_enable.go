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

package webkit2gtk

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/energy/v3/platform/linux/webkit2gtk/cgo"
	wvTypes "github.com/energye/wv/types/linux"
	"unsafe"
)

func AsWebkit2(ptr unsafe.Pointer) IWebkit2 {
	return cgo.AsWebkit2(ptr)
}

func Webkit2Ver() wvTypes.TWebkit2Version {
	return cgo.Wkv
}
