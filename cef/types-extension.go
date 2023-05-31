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
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefExtension) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefExtension) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefExtension) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefExtension) GetIdentifier() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetIdentifier).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetPath() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetPath).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetManifest() *ICefDictionaryValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefExtension_GetManifest).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDictionaryValue{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefExtension) IsSame(that *ICefExtension) bool {
	if !m.IsValid() || !that.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefExtension_IsSame).Call(m.Instance(), that.Instance())
	return api.GoBool(r1)
}

func (m *ICefExtension) GetHandler() *ICefExtensionHandler {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefExtension_GetHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefExtensionHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefExtension) GetLoaderContext() *ICefRequestContext {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefExtension_GetLoaderContext).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestContext{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefExtension) IsLoaded() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefExtension_IsLoaded).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefExtension) unload() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefExtension_unload).Call(m.Instance())
}

func (m *ICefExtension) GetBrowserActionPopup() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetBrowserActionPopup).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetBrowserActionIcon() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetBrowserActionIcon).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetPageActionPopup() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetPageActionPopup).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetPageActionIcon() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetPageActionIcon).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetOptionsPage() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetOptionsPage).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetOptionsUIPage() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetOptionsUIPage).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetBackgroundPage() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetBackgroundPage).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefExtension) GetURL() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefExtension_GetURL).Call(m.Instance())
	return api.GoStr(r1)
}
