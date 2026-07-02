//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !cgo

package types

import (
	"github.com/energye/energy/v3/platform/linux"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"unsafe"
)

var gobjectLib *linux.DnyLibrary

func init() {
	gobjectLib = linux.LibLoad(linux.Libgobject2_0)
	gobjectLib.Table = []*imports.Table{
		imports.NewTable("strlen", 0),
		imports.NewTable("g_type_name", 0),
		imports.NewTable("g_type_depth", 0),
		imports.NewTable("g_type_parent", 0),
		imports.NewTable("g_type_check_is_value_type", 0),
		imports.NewTable("g_type_from_name", 0),
		imports.NewTable("g_type_is_a", 0),
	}
	gobjectLib.SetLibClose()
	gobjectLib.MapperIndex()
}

// IsValue checks whether the passed in type can be used for g_value_init().
func (t Type) IsValue() bool {
	r := gobjectLib.SysCall("g_type_check_is_value_type", uintptr(t))
	return r != 0
}

// Name is a wrapper around g_type_name().
func (t Type) Name() string {
	r := gobjectLib.SysCall("g_type_name", uintptr(t))
	if r == 0 {
		return ""
	}
	return _GoStr(r)
}

// Depth is a wrapper around g_type_depth().
func (t Type) Depth() uint {
	r := gobjectLib.SysCall("g_type_depth", uintptr(t))
	return uint(r)
}

// Parent is a wrapper around g_type_parent().
func (t Type) Parent() Type {
	r := gobjectLib.SysCall("g_type_parent", uintptr(t))
	return Type(r)
}

// IsA is a wrapper around g_type_is_a().
func (t Type) IsA(isAType Type) bool {
	r := gobjectLib.SysCall("g_type_is_a", uintptr(t), uintptr(isAType))
	return r != 0
}

// TypeFromName is a wrapper around g_type_from_name().
func TypeFromName(typeName string) Type {
	cstr := api.PasStr(typeName)
	r := gobjectLib.SysCall("g_type_from_name", cstr)
	return Type(r)
}

func _GoStr(cStrPtr uintptr) string {
	if cStrPtr == 0 {
		return ""
	}
	strLen := int(gobjectLib.SysCall("strlen", cStrPtr))
	if strLen == 0 {
		return ""
	}
	return string((*[1 << 30]byte)(unsafe.Pointer(cStrPtr))[:strLen:strLen])
}
