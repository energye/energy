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

// ICefSchemeHandlerFactory
//
//	/include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)
type ICefSchemeHandlerFactory struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// SchemeHandlerFactoryRef -> ICefSchemeHandlerFactory
var SchemeHandlerFactoryRef schemeHandlerFactory

type schemeHandlerFactory uintptr

func (*schemeHandlerFactory) New() *ICefSchemeHandlerFactory {
	var result uintptr
	imports.Proc(def.SchemeHandlerFactoryRef_Create).Call(uintptr(0), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefSchemeHandlerFactory{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefSchemeHandlerFactory) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefSchemeHandlerFactory) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefSchemeHandlerFactory) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefSchemeHandlerFactory) SetNew(fn schemeHandlerFactoryNew) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.SchemeHandlerFactory_New).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type schemeHandlerFactoryNew func(browser *ICefBrowser, frame *ICefFrame, schemeName string, request *ICefRequest) *ICefResourceHandler

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case schemeHandlerFactoryNew:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			schemeName := api.GoStr(getVal(2))
			request := &ICefRequest{instance: getPtr(3)}
			resourceHandlerPtr := (*uintptr)(getPtr(4))
			resourceHandler := fn.(schemeHandlerFactoryNew)(browser, frame, schemeName, request)
			if resourceHandler != nil && resourceHandler.IsValid() {
				*resourceHandlerPtr = resourceHandler.Instance()
			}
		default:
			return false
		}
		return true
	})
}
