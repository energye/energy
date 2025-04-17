//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCefDownloadImageCallback
// include/capi/cef_browser_capi.h (cef_download_image_callback_t)
type ICefDownloadImageCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DownloadImageCallbackRef -> ICefDownloadImageCallback
var DownloadImageCallbackRef downloadImageCallback

type downloadImageCallback uintptr

func (*downloadImageCallback) New() *ICefDownloadImageCallback {
	var result uintptr
	imports.Proc(def.DownloadImageCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDownloadImageCallback{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefDownloadImageCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDownloadImageCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDownloadImageCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDownloadImageCallback) SetOnDownloadImageFinished(fn onDownloadImageFinished) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.DownloadImageCallback_OnDownloadImageFinished).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onDownloadImageFinished func(imageUrl string, httpStatusCode int32, image *ICefImage)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onDownloadImageFinished:
			fn.(onDownloadImageFinished)(api.GoStr(getVal(0)), int32(getVal(1)), &ICefImage{instance: getPtr(2)})
		default:
			return false
		}
		return true
	})
}
