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

// ICefRenderHandler
type ICefRenderHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// RenderHandlerRef -> ICefRenderHandler
var RenderHandlerRef renderSpanHandler

type renderSpanHandler uintptr

func (*renderSpanHandler) New() *ICefRenderHandler {
	var result uintptr
	imports.Proc(def.CefRenderHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRenderHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefRenderHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRenderHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRenderHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRenderHandler) SetGetAccessibilityHandler(fn renderHandlerGetAccessibilityHandler) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_GetAccessibilityHandler).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetGetRootScreenRect(fn renderHandlerGetRootScreenRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_GetRootScreenRect).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetGetViewRect(fn renderHandlerGetViewRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_GetViewRect).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetGetScreenPoint(fn renderHandlerGetScreenPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_GetScreenPoint).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetGetScreenInfo(fn renderHandlerGetScreenInfo) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_GetScreenInfo).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnPopupShow(fn renderHandlerOnPopupShow) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnPopupShow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnPopupSize(fn renderHandlerOnPopupSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnPopupSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnPaint(fn renderHandlerOnPaint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnPaint).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnAcceleratedPaint(fn renderHandlerOnAcceleratedPaint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnAcceleratedPaint).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetGetTouchHandleSize(fn renderHandlerGetTouchHandleSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_GetTouchHandleSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnTouchHandleStateChanged(fn renderHandlerOnTouchHandleStateChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnTouchHandleStateChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnStartDragging(fn renderHandlerOnStartDragging) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnStartDragging).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnUpdateDragCursor(fn renderHandlerOnUpdateDragCursor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnUpdateDragCursor).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnScrollOffsetChanged(fn renderHandlerOnScrollOffsetChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnScrollOffsetChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnIMECompositionRangeChanged(fn renderHandlerOnIMECompositionRangeChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnIMECompositionRangeChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnTextSelectionChanged(fn renderHandlerOnTextSelectionChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnTextSelectionChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRenderHandler) SetOnVirtualKeyboardRequested(fn renderHandlerOnVirtualKeyboardRequested) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRenderHandler_OnVirtualKeyboardRequested).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type renderHandlerGetAccessibilityHandler func(accessibilityHandler *ICefAccessibilityHandler)
type renderHandlerGetRootScreenRect func(browser *ICefBrowser) (rect *TCefRect, result bool)
type renderHandlerGetViewRect func(browser *ICefBrowser) *TCefRect
type renderHandlerGetScreenPoint func(browser *ICefBrowser, viewX, viewY int32) (screenX, screenY int32, result bool)
type renderHandlerGetScreenInfo func(browser *ICefBrowser) (screenInfo *TCefScreenInfo, result bool)
type renderHandlerOnPopupShow func(browser *ICefBrowser, show bool)
type renderHandlerOnPopupSize func(browser *ICefBrowser, rect *TCefRect)
type renderHandlerOnPaint func(browser *ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *TCefRectArray, buffer uintptr, width, height int32)
type renderHandlerOnAcceleratedPaint func(browser *ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *TCefRectArray, info TCefAcceleratedPaintInfo)
type renderHandlerGetTouchHandleSize func(browser *ICefBrowser, orientation consts.TCefHorizontalAlignment) *TCefSize
type renderHandlerOnTouchHandleStateChanged func(browser *ICefBrowser, state *TCefTouchHandleState)
type renderHandlerOnStartDragging func(browser *ICefBrowser, dragData *ICefDragData, allowedOps consts.TCefDragOperations, x, y int32) bool
type renderHandlerOnUpdateDragCursor func(browser *ICefBrowser, operation consts.TCefDragOperation)
type renderHandlerOnScrollOffsetChanged func(browser *ICefBrowser, x, y float64)
type renderHandlerOnIMECompositionRangeChanged func(browser *ICefBrowser, selectedRange *TCefRange, characterBoundsCount uint32, characterBounds *TCefRect)
type renderHandlerOnTextSelectionChanged func(browser *ICefBrowser, selectedText string, selectedRange *TCefRange)
type renderHandlerOnVirtualKeyboardRequested func(browser *ICefBrowser, inputMode consts.TCefTextInputMode)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case renderHandlerGetAccessibilityHandler:
			accessibilityHandler := &ICefAccessibilityHandler{instance: getPtr(0)}
			fn.(renderHandlerGetAccessibilityHandler)(accessibilityHandler)
		case renderHandlerGetRootScreenRect:
			browser := &ICefBrowser{instance: getPtr(0)}
			rectPtr := (*TCefRect)(getPtr(1))
			resultPtr := (*bool)(getPtr(2))
			rect, result := fn.(renderHandlerGetRootScreenRect)(browser)
			if rect != nil {
				*rectPtr = *rect
			}
			*resultPtr = result
		case renderHandlerGetViewRect:
			browser := &ICefBrowser{instance: getPtr(0)}
			rectPtr := (*TCefRect)(getPtr(1))
			rect := fn.(renderHandlerGetViewRect)(browser)
			if rect != nil {
				*rectPtr = *rect
			}
		case renderHandlerGetScreenPoint:
			browser := &ICefBrowser{instance: getPtr(0)}
			viewX, viewY := int32(getVal(1)), int32(getVal(2))
			screenXPtr, screenYPtr, resultPtr := (*int32)(getPtr(3)), (*int32)(getPtr(4)), (*bool)(getPtr(5))
			screenX, screenY, result := fn.(renderHandlerGetScreenPoint)(browser, viewX, viewY)
			*screenXPtr, *screenYPtr, *resultPtr = screenX, screenY, result
		case renderHandlerGetScreenInfo:
			browser := &ICefBrowser{instance: getPtr(0)}
			screenInfoPtr := (*TCefScreenInfo)(getPtr(1))
			resultPtr := (*bool)(getPtr(2))
			screenInfo, result := fn.(renderHandlerGetScreenInfo)(browser)
			if screenInfo != nil {
				*screenInfoPtr = *screenInfo
			}
			*resultPtr = result
		case renderHandlerOnPopupShow:
			browser := &ICefBrowser{instance: getPtr(0)}
			show := api.GoBool(getVal(1))
			fn.(renderHandlerOnPopupShow)(browser, show)
		case renderHandlerOnPopupSize:
			browser := &ICefBrowser{instance: getPtr(0)}
			rect := (*TCefRect)(getPtr(1))
			fn.(renderHandlerOnPopupSize)(browser, rect)
		case renderHandlerOnPaint:
			browser := &ICefBrowser{instance: getPtr(0)}
			kind := consts.TCefPaintElementType(getVal(1))
			dirtyRectsCount := uint32(getVal(2))
			dirtyRectsPtr := getVal(3)
			buffer := getVal(4)
			width, height := int32(getVal(5)), int32(getVal(6))
			fn.(renderHandlerOnPaint)(browser, kind, NewTCefRectArray(dirtyRectsPtr, dirtyRectsCount), buffer, width, height)
		case renderHandlerOnAcceleratedPaint:
			browser := &ICefBrowser{instance: getPtr(0)}
			kind := consts.TCefPaintElementType(getVal(1))
			dirtyRectsCount := uint32(getVal(2))
			dirtyRectsPtr := getVal(3)
			info := *(*TCefAcceleratedPaintInfo)(getPtr(4))
			fn.(renderHandlerOnAcceleratedPaint)(browser, kind, NewTCefRectArray(dirtyRectsPtr, dirtyRectsCount), info)
		case renderHandlerGetTouchHandleSize:
			browser := &ICefBrowser{instance: getPtr(0)}
			orientation := consts.TCefHorizontalAlignment(getVal(1))
			sizePtr := (*TCefSize)(getPtr(2))
			size := fn.(renderHandlerGetTouchHandleSize)(browser, orientation)
			if size != nil {
				*sizePtr = *size
			}
		case renderHandlerOnTouchHandleStateChanged:
			browser := &ICefBrowser{instance: getPtr(0)}
			statePtr := (*tCefTouchHandleStatePtr)(getPtr(1))
			state := statePtr.convert()
			fn.(renderHandlerOnTouchHandleStateChanged)(browser, state)
		case renderHandlerOnStartDragging:
			browser := &ICefBrowser{instance: getPtr(0)}
			dragData := &ICefDragData{instance: getPtr(1)}
			allowedOps := consts.TCefDragOperations(getVal(2))
			x, y := int32(getVal(3)), int32(getVal(4))
			resultPtr := (*bool)(getPtr(5))
			*resultPtr = fn.(renderHandlerOnStartDragging)(browser, dragData, allowedOps, x, y)
		case renderHandlerOnUpdateDragCursor:
			browser := &ICefBrowser{instance: getPtr(0)}
			operation := consts.TCefDragOperation(getVal(1))
			fn.(renderHandlerOnUpdateDragCursor)(browser, operation)
		case renderHandlerOnScrollOffsetChanged:
			browser := &ICefBrowser{instance: getPtr(0)}
			x, y := *(*float64)(getPtr(1)), *(*float64)(getPtr(2))
			fn.(renderHandlerOnScrollOffsetChanged)(browser, x, y)
		case renderHandlerOnIMECompositionRangeChanged:
			browser := &ICefBrowser{instance: getPtr(0)}
			rng := (*TCefRange)(getPtr(1))
			characterBoundsCount := uint32(getVal(2))
			characterBounds := (*TCefRect)(getPtr(3))
			fn.(renderHandlerOnIMECompositionRangeChanged)(browser, rng, characterBoundsCount, characterBounds)
		case renderHandlerOnTextSelectionChanged:
			browser := &ICefBrowser{instance: getPtr(0)}
			selectedText := api.GoStr(getVal(1))
			selectedRange := (*TCefRange)(getPtr(2))
			fn.(renderHandlerOnTextSelectionChanged)(browser, selectedText, selectedRange)
		case renderHandlerOnVirtualKeyboardRequested:
			browser := &ICefBrowser{instance: getPtr(0)}
			inputMode := consts.TCefTextInputMode(getVal(1))
			fn.(renderHandlerOnVirtualKeyboardRequested)(browser, inputMode)
		default:
			return false
		}
		return true
	})
}
