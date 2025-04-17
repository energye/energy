//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//CEF v8 异常

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefV8Exception
type ICefV8Exception struct {
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefV8Exception) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8Exception) IsValid() bool {
	return m != nil && m.instance != nil
}

func (m *ICefV8Exception) Message() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_Message).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefV8Exception) SourceLine() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_SourceLine).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefV8Exception) ScriptResourceName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_ScriptResourceName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefV8Exception) LineNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_LineNumber).Call(m.Instance())
	return int32(r1)
}

func (m *ICefV8Exception) StartPosition() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_StartPosition).Call(m.Instance())
	return int32(r1)
}

func (m *ICefV8Exception) EndPosition() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_EndPosition).Call(m.Instance())
	return int32(r1)
}

func (m *ICefV8Exception) StartColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_StartColumn).Call(m.Instance())
	return int32(r1)
}

func (m *ICefV8Exception) EndColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Exception_EndColumn).Call(m.Instance())
	return int32(r1)
}
