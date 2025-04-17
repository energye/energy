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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefPostData
type ICefPostData struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// PostDataRef -> ICefPostData
var PostDataRef postData

// postData
type postData uintptr

func (m *postData) New() *ICefPostData {
	var result uintptr
	imports.Proc(def.CefPostDataRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefPostData{instance: unsafe.Pointer(result)}
}

func (m *postData) UnWrap(data *ICefPostData) *ICefPostData {
	var result uintptr
	imports.Proc(def.CefPostDataRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// Instance 实例
func (m *ICefPostData) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPostData) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefPostData) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPostData) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefPostData_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPostData) HasExcludedElements() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefPostData_HasExcludedElements).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefPostData) GetElementCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefPostData_GetElementCount).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefPostData) GetElements() *TCefPostDataElementArray {
	if !m.IsValid() {
		return nil
	}
	var elementsCount = m.GetElementCount()
	if elementsCount <= 0 {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefPostData_GetElements).Call(m.Instance(), uintptr(elementsCount), uintptr(unsafe.Pointer(&result)))
	return &TCefPostDataElementArray{instance: unsafe.Pointer(result), postDataElement: result, postDataElementLength: elementsCount}
}

func (m *ICefPostData) RemoveElement(postDataElement *ICefPostDataElement) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefPostData_RemoveElement).Call(m.Instance(), postDataElement.Instance())
	return api.GoBool(r1)
}

func (m *ICefPostData) AddElement(postDataElement *ICefPostDataElement) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefPostData_AddElement).Call(m.Instance(), postDataElement.Instance())
	return api.GoBool(r1)
}

func (m *ICefPostData) RemoveElements() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPostData_RemoveElements).Call(m.Instance())
}
