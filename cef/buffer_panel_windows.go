//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

// CEF TBufferPanel

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type onIMECommitTextEvent func(sender lcl.IObject, text string, replacementRange TCefRange, relativeCursorPos int32)
type onIMESetCompositionEvent func(sender lcl.IObject, text string, underlines *TCefCompositionUnderlineArray, replacementRange, selectionRange TCefRange)
type onHandledMessageEvent func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, handled *bool)

func (m *TBufferPanel) SetOnIMECancelComposition(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnIMECancelComposition, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnIMECommitText(fn onIMECommitTextEvent) {
	imports.SysCallN(def.BufferPanel_SetOnIMECommitText, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnIMESetComposition(fn onIMESetCompositionEvent) {
	imports.SysCallN(def.BufferPanel_SetOnIMESetComposition, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnCustomTouch(fn onHandledMessageEvent) {
	imports.SysCallN(def.BufferPanel_SetOnCustomTouch, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnPointerDown(fn onHandledMessageEvent) {
	imports.SysCallN(def.BufferPanel_SetOnPointerDown, m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *TBufferPanel) SetOnPointerUp(fn onHandledMessageEvent) {
	imports.SysCallN(def.BufferPanel_SetOnPointerUp, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnPointerUpdate(fn onHandledMessageEvent) {
	imports.SysCallN(def.BufferPanel_SetOnPointerUpdate, m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onIMECommitTextEvent:
			fn.(onIMECommitTextEvent)(lcl.AsObject(getPtr(0)), api.GoStr(getVal(1)), *(*TCefRange)(getPtr(2)), int32(getVal(3)))
		case onIMESetCompositionEvent:
			underlines := &TCefCompositionUnderlineArray{
				count:  int(int32(getVal(3))),
				ptr:    *(*uintptr)(getPtr(2)),
				sizeOf: unsafe.Sizeof(TCefCompositionUnderline{}),
			}
			replacementRange := *(*TCefRange)(getPtr(4))
			selectionRange := *(*TCefRange)(getPtr(5))
			fn.(onIMESetCompositionEvent)(lcl.AsObject(getPtr(0)), api.GoStr(getVal(1)), underlines, replacementRange, selectionRange)
		case onHandledMessageEvent:
			message := (*types.TMessage)(getPtr(1))
			lResultPtr := (*types.LRESULT)(getPtr(2))
			fn.(onHandledMessageEvent)(lcl.AsObject(getVal(0)), message, lResultPtr, (*bool)(getPtr(3)))
		default:
			return false
		}
		return true
	})
}
