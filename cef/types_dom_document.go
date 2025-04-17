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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefDomDocument
type ICefDomDocument struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefDomDocument) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDomDocument) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDomDocument) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDomDocument) GetDocType() consts.TCefDomDocumentType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetDocType).Call(m.Instance())
	return consts.TCefDomDocumentType(r1)
}

func (m *ICefDomDocument) GetDocument() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomDocument_GetDocument).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomDocument) GetBody() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomDocument_GetBody).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomDocument) GetHead() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomDocument_GetHead).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomDocument) GetTitle() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetTitle).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomDocument) GetElementById(id string) *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomDocument_GetElementById).Call(m.Instance(), api.PascalStr(id), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomDocument) GetFocusedNode() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomDocument_GetFocusedNode).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomDocument) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_HasSelection).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomDocument) GetSelectionStartOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetSelectionStartOffset).Call(m.Instance())
	return int32(r1)
}

func (m *ICefDomDocument) GetSelectionEndOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetSelectionEndOffset).Call(m.Instance())
	return int32(r1)
}

func (m *ICefDomDocument) GetSelectionAsMarkup() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetSelectionAsMarkup).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomDocument) GetSelectionAsText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetSelectionAsText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomDocument) GetBaseUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetBaseUrl).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomDocument) GetCompleteUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomDocument_GetCompleteUrl).Call(m.Instance())
	return api.GoStr(r1)
}
