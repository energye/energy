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

// ICefDomNode
type ICefDomNode struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefDomNode) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDomNode) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDomNode) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDomNode) GetNodeType() consts.TCefDomNodeType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetNodeType).Call(m.Instance())
	return consts.TCefDomNodeType(r1)
}

func (m *ICefDomNode) IsText() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_IsText).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) IsElement() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_IsElement).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) IsEditable() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_IsEditable).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) IsFormControlElement() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_IsFormControlElement).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) GetFormControlElementType() consts.TCefDomFormControlType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetFormControlElementType).Call(m.Instance())
	return consts.TCefDomFormControlType(r1)
}

func (m *ICefDomNode) IsSame() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_IsSame).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) GetName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomNode) GetValue() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetValue).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomNode) SetValue(value string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_SetValue).Call(m.Instance(), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefDomNode) GetAsMarkup() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetAsMarkup).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomNode) GetDocument() *ICefDomDocument {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetDocument).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomDocument{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomNode) GetParent() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetParent).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomNode) GetPreviousSibling() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetPreviousSibling).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomNode) GetNextSibling() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetNextSibling).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomNode) HasChildren() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_HasChildren).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) GetFirstChild() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetFirstChild).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomNode) GetLastChild() *ICefDomNode {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetLastChild).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomNode{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefDomNode) GetElementTagName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetElementTagName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomNode) HasElementAttributes() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_HasElementAttributes).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefDomNode) HasElementAttribute(attrName string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_HasElementAttribute).Call(m.Instance(), api.PascalStr(attrName))
	return api.GoBool(r1)
}

func (m *ICefDomNode) GetElementAttribute(attrName string) string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetElementAttribute).Call(m.Instance(), api.PascalStr(attrName))
	return api.GoStr(r1)
}

func (m *ICefDomNode) GetElementAttributes() []string {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefDomNode_GetElementAttributes).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	attrs := lcl.AsStrings(result)
	if attrs != nil && attrs.IsValid() {
		var attrList = make([]string, attrs.Count())
		for i := 0; i < len(attrList); i++ {
			attrList[i] = attrs.Strings(int32(i))
		}
		return attrList
	}
	return nil
}

func (m *ICefDomNode) SetElementAttribute(attrName, value string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefDomNode_SetElementAttribute).Call(m.Instance(), api.PascalStr(attrName), api.PascalStr(value))
	return api.GoBool(r1)
}

func (m *ICefDomNode) GetElementInnerText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefDomNode_GetElementInnerText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefDomNode) GetElementBounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDomNode_GetElementBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}
