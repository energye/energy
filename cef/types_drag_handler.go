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
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefDragHandler
type ICefDragHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DragHandlerRef -> ICefDragHandler
var DragHandlerRef dragHandler

type dragHandler uintptr

func (*dragHandler) New() *ICefDragHandler {
	var result uintptr
	imports.Proc(def.CefDragHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDragHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefDragHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDragHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDragHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDragHandler) SetOnDragEnter(fn onDragEnter) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDragHandler_OnDragEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDragHandler) SetOnDraggableRegionsChanged(fn onDraggableRegionsChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDragHandler_OnDraggableRegionsChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onDragEnter func(browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations) bool
type onDraggableRegionsChanged func(browser *ICefBrowser, frame *ICefFrame, regions *TCefDraggableRegions)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onDragEnter:
			browse := &ICefBrowser{instance: getPtr(0)}
			dragData := &ICefDragData{instance: getPtr(1)}
			mask := consts.TCefDragOperations(getVal(2))
			result := (*bool)(getPtr(3))
			*result = fn.(onDragEnter)(browse, dragData, mask)
		case onDraggableRegionsChanged:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			regionsCount := int32(getVal(2))
			regions := NewCefDraggableRegions()
			var region TCefDraggableRegion
			var regionSize = unsafe.Sizeof(region)
			for i := 0; i < int(regionsCount); i++ {
				region = *(*TCefDraggableRegion)(common.GetParamPtr(getVal(3), i*int(regionSize)))
				regions.Append(region)
			}
			fn.(onDraggableRegionsChanged)(browse, frame, regions)
		default:
			return false
		}
		return true
	})
}
