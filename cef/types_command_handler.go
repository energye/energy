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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefCommandHandler
type ICefCommandHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// CommandHandlerRef -> ICefCommandHandler
var CommandHandlerRef commandHandler

type commandHandler uintptr

func (*commandHandler) New() *ICefCommandHandler {
	var result uintptr
	imports.Proc(def.CefCommandHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefCommandHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefCommandHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCommandHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefCommandHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefCommandHandler) SetOnChromeCommand(fn onChromeCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefCommandHandler_OnChromeCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onChromeCommand func(browser *ICefBrowser, commandId int32, disposition consts.TCefWindowOpenDisposition) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onChromeCommand:
			browse := &ICefBrowser{instance: getPtr(0)}
			params := (int32)(getVal(1))
			disposition := consts.TCefWindowOpenDisposition(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(onChromeCommand)(browse, params, disposition)
		default:
			return false
		}
		return true
	})
}
