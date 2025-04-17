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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

func (m *TBufferPanel) SetOnIMECancelComposition(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnIMECancelComposition, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnIMECommitText(fn bufferPanelOnIMECommitText) {
	imports.SysCallN(def.BufferPanel_SetOnIMECommitText, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnIMESetComposition(fn bufferPanelOnIMESetComposition) {
	imports.SysCallN(def.BufferPanel_SetOnIMESetComposition, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnCustomTouch(fn bufferPanelOnHandledMessage) {
	imports.SysCallN(def.BufferPanel_SetOnCustomTouch, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnPointerDown(fn bufferPanelOnHandledMessage) {
	imports.SysCallN(def.BufferPanel_SetOnPointerDown, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnPointerUp(fn bufferPanelOnHandledMessage) {
	imports.SysCallN(def.BufferPanel_SetOnPointerUp, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnPointerUpdate(fn bufferPanelOnHandledMessage) {
	imports.SysCallN(def.BufferPanel_SetOnPointerUpdate, m.Instance(), api.MakeEventDataPtr(fn))
}

type TCefCompositionUnderlineArray struct {
	instance unsafe.Pointer
	count    int
	size     uintptr
}

func (m *TCefCompositionUnderlineArray) Count() int {
	return m.count
}

func (m *TCefCompositionUnderlineArray) Get(index int) (compUnderLine TCefCompositionUnderline) {
	if m.instance == nil {
		return
	}
	if index >= 0 && index < m.count {
		compUnderLine = *(*TCefCompositionUnderline)(common.GetParamPtr(uintptr(m.instance), index*int(m.size)))
		//imports.SysCallN(def.BufferPanelCompositionUnderline_Get, uintptr(m.instance), uintptr(int32(index)), uintptr(unsafePointer(&compUnderLine)))
	}
	return
}

type bufferPanelOnIMECommitText func(sender lcl.IObject, text string, replacementRange TCefRange, relativeCursorPos int32)
type bufferPanelOnIMESetComposition func(sender lcl.IObject, text string, underlines *TCefCompositionUnderlineArray, replacementRange, selectionRange TCefRange)
type bufferPanelOnHandledMessage func(sender lcl.IObject, message *types.TMessage, lResult *types.LRESULT, handled *bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case bufferPanelOnIMECommitText:
			fn.(bufferPanelOnIMECommitText)(lcl.AsObject(getPtr(0)), api.GoStr(getVal(1)), *(*TCefRange)(getPtr(2)), int32(getVal(3)))
		case bufferPanelOnIMESetComposition:
			underlines := &TCefCompositionUnderlineArray{
				instance: getPtr(2),
				count:    int(int32(getVal(3))),
				size:     unsafe.Sizeof(TCefCompositionUnderline{}),
			}
			replacementRange := *(*TCefRange)(getPtr(4))
			selectionRange := *(*TCefRange)(getPtr(5))
			fn.(bufferPanelOnIMESetComposition)(lcl.AsObject(getPtr(0)), api.GoStr(getVal(1)), underlines, replacementRange, selectionRange)
		case bufferPanelOnHandledMessage:
			message := (*types.TMessage)(getPtr(1))
			lResultPtr := (*types.LRESULT)(getPtr(2))
			fn.(bufferPanelOnHandledMessage)(lcl.AsObject(getVal(0)), message, lResultPtr, (*bool)(getPtr(3)))
		default:
			return false
		}
		return true
	})
}
