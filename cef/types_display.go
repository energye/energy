// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// ICefDisplay
type ICefDisplay struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DisplayRef -> ICefDisplay
var DisplayRef display

type display struct {
	alls *ICefDisplayArray
}

func (m *display) Primary() *ICefDisplay {
	var result uintptr
	imports.Proc(def.CEFDisplayRef_Primary).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDisplay{instance: getInstance(result)}
	}
	return nil
}

func (m *display) NearestPoint(point TCefPoint, inputPixelCoords bool) *ICefDisplay {
	var result uintptr
	imports.Proc(def.CEFDisplayRef_NearestPoint).Call(uintptr(unsafe.Pointer(&point)), api.PascalBool(inputPixelCoords), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDisplay{instance: getInstance(result)}
	}
	return nil
}

func (m *display) MatchingBounds(point TCefRect, inputPixelCoords bool) *ICefDisplay {
	var result uintptr
	imports.Proc(def.CEFDisplayRef_MatchingBounds).Call(uintptr(unsafe.Pointer(&point)), api.PascalBool(inputPixelCoords), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDisplay{instance: getInstance(result)}
	}
	return nil
}

func (m *display) GetCount() uint32 {
	r1, _, _ := imports.Proc(def.CEFDisplayRef_GetCount).Call()
	return uint32(r1)
}

func (m *display) GetAlls() *ICefDisplayArray {
	if m.alls == nil {
		var result uintptr
		r1, _, _ := imports.Proc(def.CEFDisplayRef_GetAlls).Call(uintptr(unsafe.Pointer(&result)))
		if r1 != 0 && result != 0 {
			m.alls = &ICefDisplayArray{instance: getInstance(result), count: m.GetCount()}
		}
	}
	return m.alls
}

func (m *display) ScreenPointToPixels(screenPoint *types.TPoint) (point types.TPoint) {
	imports.Proc(def.CEFDisplayRef_ScreenPointToPixels).Call(uintptr(unsafe.Pointer(screenPoint)), uintptr(unsafe.Pointer(&point)))
	return
}

func (m *display) ScreenPointFromPixels(pixelsPoint *types.TPoint) (point types.TPoint) {
	imports.Proc(def.CEFDisplayRef_ScreenPointFromPixels).Call(uintptr(unsafe.Pointer(pixelsPoint)), uintptr(unsafe.Pointer(&point)))
	return
}

func (m *display) ScreenRectToPixels(screenRect *types.TRect) (rect types.TRect) {
	imports.Proc(def.CEFDisplayRef_ScreenRectToPixels).Call(uintptr(unsafe.Pointer(screenRect)), uintptr(unsafe.Pointer(&rect)))
	return
}

func (m *display) ScreenRectFromPixels(pixelsRect *types.TRect) (rect types.TRect) {
	imports.Proc(def.CEFDisplayRef_ScreenRectFromPixels).Call(uintptr(unsafe.Pointer(pixelsRect)), uintptr(unsafe.Pointer(&rect)))
	return
}

func (m *ICefDisplay) ID() (id int64) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CEFDisplay_ID).Call(m.Instance(), uintptr(unsafe.Pointer(&id)))
	return
}

func (m *ICefDisplay) DeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.CEFDisplay_DeviceScaleFactor).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) Rotation() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFDisplay_Rotation).Call(m.Instance())
	return int32(r1)
}

func (m *ICefDisplay) Bounds() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFDisplay_Bounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

func (m *ICefDisplay) WorkArea() (workArea TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFDisplay_WorkArea).Call(m.Instance(), uintptr(unsafe.Pointer(&workArea)))
	return
}

func (m *ICefDisplay) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefDisplay) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDisplay) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// ICefDisplayArray
//
//	[]ICefDisplayArray
type ICefDisplayArray struct {
	instance     unsafe.Pointer
	binaryValues []*ICefDisplayArray
	count        uint32
}

func (m *ICefDisplayArray) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefDisplayArray) Free() {
	if m.instance != nil {
		m.instance = nil
	}
}

func (m *ICefDisplayArray) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDisplayArray) Get(index uint32) *ICefDisplay {
	if !m.IsValid() {
		return nil
	}
	if index < m.count {
		var result uintptr
		r1, _, _ := imports.Proc(def.CEFDisplayRef_Get).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
		if r1 != 0 && result != 0 {
			return &ICefDisplay{instance: getInstance(result)}
		}
	}
	return nil
}

func (m *ICefDisplayArray) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	return m.count
}
