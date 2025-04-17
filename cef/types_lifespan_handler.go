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

// ICefLifeSpanHandler
type ICefLifeSpanHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// LifeSpanHandlerRef -> ICefLifeSpanHandler
var LifeSpanHandlerRef lifeSpanHandler

type lifeSpanHandler uintptr

func (*lifeSpanHandler) New() *ICefLifeSpanHandler {
	var result uintptr
	imports.Proc(def.CefLifeSpanHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLifeSpanHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefLifeSpanHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefLifeSpanHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefLifeSpanHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefLifeSpanHandler) SetOnBeforePopup(fn lifeSpanHandlerOnBeforePopup) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLifeSpanHandler_OnBeforePopup).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLifeSpanHandler) SetOnAfterCreated(fn lifeSpanHandlerOnAfterCreated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLifeSpanHandler_OnAfterCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLifeSpanHandler) SetOnDoClose(fn lifeSpanHandlerOnDoClose) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLifeSpanHandler_DoClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefLifeSpanHandler) SetOnBeforeClose(fn lifeSpanHandlerOnBeforeClose) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefLifeSpanHandler_OnBeforeClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type lifeSpanHandlerOnBeforePopup func(browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, resultClient *ICefClient, settings *TCefBrowserSettings, resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool
type lifeSpanHandlerOnAfterCreated func(browser *ICefBrowser)
type lifeSpanHandlerOnDoClose func(browser *ICefBrowser) bool
type lifeSpanHandlerOnBeforeClose func(browser *ICefBrowser)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case lifeSpanHandlerOnBeforePopup:
			var (
				browse             = &ICefBrowser{instance: getPtr(0)}
				frame              = &ICefFrame{instance: getPtr(1)}
				beforePInfoPtr     = (*beforePopupInfoPtr)(getPtr(2))
				popupFeaturesPtr   = (*tCefPopupFeaturesPtr)(getPtr(3))
				windowInfoPtr      = (*tCefWindowInfoPtr)(getPtr(4))
				resultClientPtr    = (*uintptr)(getPtr(5))
				browserSettingsPtr = (*tCefBrowserSettingsPtr)(getPtr(6))
				resultExtraInfoPtr = (*uintptr)(getPtr(7)) // CEF49 = nil
				noJavascriptAccess = (*bool)(getPtr(8))
				result             = (*bool)(getPtr(9))
			)
			beforePopupInfo := beforePInfoPtr.convert()
			popupFeatures := popupFeaturesPtr.convert()
			windowInfo := windowInfoPtr.convert()
			resultClient := &ICefClient{}
			browserSettings := browserSettingsPtr.convert()
			resultExtraInfo := &ICefDictionaryValue{}
			*result = fn.(lifeSpanHandlerOnBeforePopup)(browse, frame, beforePopupInfo, popupFeatures, windowInfo, resultClient, browserSettings, resultExtraInfo, noJavascriptAccess)
			windowInfo.setInstanceValue()
			if resultClient.IsValid() {
				*resultClientPtr = resultClient.Instance()
			}
			browserSettings.setInstanceValue()
			if resultExtraInfo.IsValid() && *resultExtraInfoPtr != 0 {
				*resultExtraInfoPtr = resultExtraInfo.Instance()
			}
		case lifeSpanHandlerOnAfterCreated:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(lifeSpanHandlerOnAfterCreated)(browse)
		case lifeSpanHandlerOnDoClose:
			browse := &ICefBrowser{instance: getPtr(0)}
			result := (*bool)(getPtr(1))
			*result = fn.(lifeSpanHandlerOnDoClose)(browse)
		case lifeSpanHandlerOnBeforeClose:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(lifeSpanHandlerOnBeforeClose)(browse)
		default:
			return false
		}
		return true
	})
}
