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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ExtensionHandlerRef -> TCustomExtensionHandler
var ExtensionHandlerRef extensionHandler

type extensionHandler uintptr

func (*extensionHandler) NewForChromium(chromium IChromium) *TCustomExtensionHandler {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefExtensionHandlerRef_CreateForChromium).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCustomExtensionHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*extensionHandler) New() *ICefExtensionHandler {
	var result uintptr
	imports.Proc(def.CefExtensionHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefExtensionHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *TCustomExtensionHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCustomExtensionHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
func (m *TCustomExtensionHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Instance 实例
func (m *ICefExtensionHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefExtensionHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
func (m *ICefExtensionHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefExtensionHandler) SetOnExtensionLoadFailed(fn onExtensionLoadFailed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_OnExtensionLoadFailed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetOnExtensionLoaded(fn onExtensionLoaded) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_OnExtensionLoaded).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetOnExtensionUnloaded(fn onExtensionUnloaded) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_OnExtensionUnloaded).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetOnBeforeBackgroundBrowser(fn onBeforeBackgroundBrowser) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_OnBeforeBackgroundBrowser).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetOnBeforeBrowser(fn onBeforeBrowser) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_OnBeforeBrowser).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetGetActiveBrowser(fn getActiveBrowser) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_GetActiveBrowser).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetCanAccessBrowser(fn canAccessBrowser) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_CanAccessBrowser).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefExtensionHandler) SetGetExtensionResource(fn getExtensionResource) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefExtensionHandler_GetExtensionResource).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onExtensionLoadFailed func(result consts.TCefErrorCode)
type onExtensionLoaded func(extension *ICefExtension)
type onExtensionUnloaded func(extension *ICefExtension)
type onBeforeBackgroundBrowser func(extension *ICefExtension, url string) (resultClient *ICefClient, resultSettings *TCefBrowserSettings, result bool)
type onBeforeBrowser func(extension *ICefExtension, browser, activeBrowser *ICefBrowser, index int32, url string, active bool, resultClient *ICefClient, resultSettings *TCefBrowserSettings) bool
type getActiveBrowser func(extension *ICefExtension, browser *ICefBrowser, includeIncognito bool, resultBrowser *ICefBrowser)
type canAccessBrowser func(extension *ICefExtension, browser *ICefBrowser, includeIncognito bool, targetBrowser *ICefBrowser) bool
type getExtensionResource func(extension *ICefExtension, browser *ICefBrowser, file string, callback *ICefGetExtensionResourceCallback) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case onExtensionLoadFailed:
			fn.(onExtensionLoadFailed)(consts.TCefErrorCode(getVal(0)))
		case onExtensionLoaded:
			fn.(onExtensionLoaded)(&ICefExtension{instance: unsafe.Pointer(getVal(0))})
		case onExtensionUnloaded:
			fn.(onExtensionUnloaded)(&ICefExtension{instance: unsafe.Pointer(getVal(0))})
		case onBeforeBackgroundBrowser:
			extension := &ICefExtension{instance: unsafe.Pointer(getVal(0))}
			url := api.GoStr(getVal(1))
			clientPtr := (*uintptr)(unsafe.Pointer(getVal(2)))
			// TODO TCefBrowserSettings
			//resultSettingsPtr := (*uintptr)(unsafe.Pointer(getVal(3)))
			resultPtr := (*bool)(unsafe.Pointer(getVal(4)))
			client, resultSettings, result := fn.(onBeforeBackgroundBrowser)(extension, url)
			if client.instance != nil && client.IsValid() {
				*clientPtr = client.Instance()
			}
			if resultSettings != nil {
				//*resultSettingsPtr = resultSettings
			}
			*resultPtr = result
			//*resultSettingsPtr = resultSettings.ToPtr()
		case onBeforeBrowser:
			extension := &ICefExtension{instance: unsafe.Pointer(getVal(0))}
			browse, activeBrowser := &ICefBrowser{instance: unsafe.Pointer(getVal(1))}, &ICefBrowser{instance: unsafe.Pointer(getVal(2))}
			index := int32(getVal(3))
			url := api.GoStr(getVal(4))
			active := api.GoBool(getVal(5))
			//windowInfoPtr:=(*uintptr)(unsafe.Pointer(getVal(6)))
			resultClientPtr := (*uintptr)(unsafe.Pointer(getVal(7)))
			resultClient := &ICefClient{}
			// TODO TCefBrowserSettings
			//resultSettingsPtr := (*uintptr)(unsafe.Pointer(getVal(8)))
			resultSettings := &TCefBrowserSettings{}
			result := (*bool)(unsafe.Pointer(getVal(9)))
			*result = fn.(onBeforeBrowser)(extension, browse, activeBrowser, index, url, active, resultClient, resultSettings)
			if resultClient.instance != nil {
				*resultClientPtr = resultClient.Instance()
			}
		//*resultSettingsPtr = resultSettings.ToPtr()
		case getActiveBrowser:
			extension := &ICefExtension{instance: unsafe.Pointer(getVal(0))}
			browse := &ICefBrowser{instance: unsafe.Pointer(getVal(1))}
			includeIncognito := api.GoBool(getVal(2))
			resultBrowserPtr := (*uintptr)(unsafe.Pointer(getVal(3)))
			resultBrowser := &ICefBrowser{}
			fn.(getActiveBrowser)(extension, browse, includeIncognito, resultBrowser)
			if resultBrowser.instance != nil {
				*resultBrowserPtr = resultBrowser.Instance()
			}
		case canAccessBrowser:
			extension := &ICefExtension{instance: unsafe.Pointer(getVal(0))}
			browse := &ICefBrowser{instance: unsafe.Pointer(getVal(1))}
			includeIncognito := api.GoBool(getVal(2))
			targetBrowser := &ICefBrowser{instance: unsafe.Pointer(getVal(3))}
			result := (*bool)(unsafe.Pointer(getVal(4)))
			*result = fn.(canAccessBrowser)(extension, browse, includeIncognito, targetBrowser)
		case getExtensionResource:
			extension := &ICefExtension{instance: unsafe.Pointer(getVal(0))}
			browse := &ICefBrowser{instance: unsafe.Pointer(getVal(1))}
			file := api.GoStr(getVal(2))
			callback := &ICefGetExtensionResourceCallback{instance: unsafe.Pointer(getVal(3))}
			result := (*bool)(unsafe.Pointer(getVal(4)))
			*result = fn.(getExtensionResource)(extension, browse, file, callback)
		default:
			return false
		}
		return true
	})
}
