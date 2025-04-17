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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefCommandLine
type ICefCommandLine struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

var CommandLineRef commandLine

type commandLine uintptr

func (*commandLine) New() *ICefCommandLine {
	var result uintptr
	imports.Proc(def.CefCommandLineRef_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefCommandLine{
		instance: unsafe.Pointer(result),
	}
}

func (*commandLine) UnWrap(data *ICefCommandLine) *ICefCommandLine {
	var result uintptr
	imports.Proc(def.CefCommandLineRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

func (*commandLine) Global() *ICefCommandLine {
	var result uintptr
	imports.Proc(def.CefCommandLineRef_Global).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefCommandLine{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefCommandLine) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefCommandLine) IsValid() bool {
	r1, _, _ := imports.Proc(def.CefCommandLine_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefCommandLine) IsReadOnly() bool {
	r1, _, _ := imports.Proc(def.CefCommandLine_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefCommandLine) Copy() *ICefCommandLine {
	var result uintptr
	imports.Proc(def.CefCommandLine_Copy).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefCommandLine{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefCommandLine) InitFromArgv(argc int32, argv string) {
	imports.Proc(def.CefCommandLine_InitFromArgv).Call(m.Instance(), uintptr(argc), api.PascalStr(argv))
}

func (m *ICefCommandLine) InitFromString(commandLine string) {
	imports.Proc(def.CefCommandLine_InitFromString).Call(m.Instance(), api.PascalStr(commandLine))
}

func (m *ICefCommandLine) Reset() {
	imports.Proc(def.CefCommandLine_Reset).Call(m.Instance())
}

func (m *ICefCommandLine) GetCommandLineString() string {
	r1, _, _ := imports.Proc(def.CefCommandLine_GetCommandLineString).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefCommandLine) GetArgv() *lcl.TStrings {
	var result uintptr
	imports.Proc(def.CefCommandLine_GetArgv).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return lcl.AsStrings(result)
}

func (m *ICefCommandLine) GetProgram() string {
	r1, _, _ := imports.Proc(def.CefCommandLine_GetProgram).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefCommandLine) SetProgram(prog string) {
	imports.Proc(def.CefCommandLine_SetProgram).Call(m.Instance(), api.PascalStr(prog))
}

func (m *ICefCommandLine) HasSwitches() bool {
	r1, _, _ := imports.Proc(def.CefCommandLine_HasSwitches).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefCommandLine) HasSwitch(name string) bool {
	r1, _, _ := imports.Proc(def.CefCommandLine_HasSwitch).Call(m.Instance(), api.PascalStr(name))
	return api.GoBool(r1)
}

func (m *ICefCommandLine) GetSwitchValue(name string) string {
	r1, _, _ := imports.Proc(def.CefCommandLine_GetSwitchValue).Call(m.Instance(), api.PascalStr(name))
	return api.GoStr(r1)
}

func (m *ICefCommandLine) GetSwitches() *lcl.TStrings {
	var result uintptr
	imports.Proc(def.CefCommandLine_GetSwitches).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return lcl.AsStrings(result)
}

func (m *ICefCommandLine) GetSwitchesList() (*lcl.TStrings, *lcl.TStrings) {
	var keys uintptr
	var values uintptr
	imports.Proc(def.CefCommandLine_GetSwitchesList).Call(m.Instance(), uintptr(unsafe.Pointer(&keys)), uintptr(unsafe.Pointer(&values)))
	return lcl.AsStrings(keys), lcl.AsStrings(values)
}

func (m *ICefCommandLine) AppendSwitch(name string) {
	imports.Proc(def.CefCommandLine_AppendSwitch).Call(m.Instance(), api.PascalStr(name))
}

func (m *ICefCommandLine) AppendSwitchWithValue(name, value string) {
	imports.Proc(def.CefCommandLine_AppendSwitchWithValue).Call(m.Instance(), api.PascalStr(name), api.PascalStr(value))
}

func (m *ICefCommandLine) HasArguments() bool {
	r1, _, _ := imports.Proc(def.CefCommandLine_HasArguments).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefCommandLine) GetArguments() *lcl.TStrings {
	var result uintptr
	imports.Proc(def.CefCommandLine_GetArguments).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return lcl.AsStrings(result)
}

func (m *ICefCommandLine) AppendArgument(argument string) {
	imports.Proc(def.CefCommandLine_AppendArgument).Call(m.Instance(), api.PascalStr(argument))
}

func (m *ICefCommandLine) PrependWrapper(wrapper string) {
	imports.Proc(def.CefCommandLine_PrependWrapper).Call(m.Instance(), api.PascalStr(wrapper))
}

func (m *ICefCommandLine) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
