// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package nocgo

import (
	"github.com/energye/energy/v3/platform/linux"
	"github.com/energye/lcl/api/imports"
	"unsafe"
)

type Bin struct {
	Container
}

type Object struct {
	instance unsafe.Pointer
}

func (m *Object) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *Object) SetInstance(ptr unsafe.Pointer) {
	m.instance = ptr
}

// Ref is a wrapper around g_object_ref().
func (m *Object) Ref() {
	gobject2_0.SysCall("g_object_ref", m.Instance())
}

// Unref is a wrapper around g_object_unref().
func (m *Object) Unref() {
	gobject2_0.SysCall("g_object_unref", m.Instance())
}

// RefSink is a wrapper around g_object_ref_sink().
func (m *Object) RefSink() {
	gobject2_0.SysCall("g_object_ref_sink", m.Instance())
}

// IsFloating is a wrapper around g_object_is_floating().
func (m *Object) IsFloating() bool {
	r := gobject2_0.SysCall("g_object_is_floating", m.Instance())
	return ToGoBool(r)
}

// ForceFloating is a wrapper around g_object_force_floating().
func (m *Object) ForceFloating() {
	gobject2_0.SysCall("g_object_force_floating", m.Instance())
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (m *Object) StopEmission(s string) {
	cstr := CStr(s)
	gobject2_0.SysCall("g_signal_stop_emission_by_name", m.Instance(), cstr)
}

// IsA is a wrapper around g_type_is_a().
func (m *Object) IsA(typ uintptr) bool {
	return GTypeIsA(GTypeFormInstance(m.Instance()), typ)
}

// SetData is a wrapper around g_object_set_data().
func (m *Object) SetData(key string, value unsafe.Pointer) {
	cKey := CStr(key)
	gobject2_0.SysCall("g_object_set_data", m.Instance(), cKey, uintptr(value))
}

// GetData is a wrapper around g_object_get_data().
func (m *Object) GetData(key string) unsafe.Pointer {
	cKey := CStr(key)
	r := gobject2_0.SysCall("g_object_get_data", m.Instance(), cKey)
	return unsafe.Pointer(r)
}

func ucharString(guchar uintptr) string {
	if guchar == 0 {
		return ""
	}

	var strlen uintptr
	for ptr := guchar; *(*byte)(unsafe.Pointer(ptr)) != 0; ptr++ {
		strlen++
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(guchar)), strlen))
}

func GSignalConnectData(object uintptr, detailedSignal string, cHandler uintptr, data uintptr) uintptr {
	//gobject2_0.SysCall("g_signal_connect_data")
	//var gSignalConnect func(object uintptr, detailedSignal *byte, cHandler uintptr, data uintptr) uintptr
	//purego.RegisterLibFunc(&gSignalConnect, libgobject, "g_signal_connect_data")

	return 0
}

func GTypeFromName(name uintptr) uintptr {
	return gobject2_0.SysCall("g_type_from_name", name)
}

func GTypeIsA(type_id, is_a_type uintptr) bool {
	r := gobject2_0.SysCall("g_type_is_a", type_id, is_a_type)
	return ToGoBool(r)
}

func GTypeFormInstance(widget uintptr) uintptr {
	return GTypeFromName(gobject2_0.SysCall("g_type_name_from_instance", widget))
}

var (
	gTypeString uintptr
)

type GValue struct {
	GType uintptr
	Data  [2]uintptr
}

func NewStringGValue(value string) *GValue {
	if gTypeString == 0 {
		return nil
	}
	gValueSize := GValue{}
	gobject2_0.SysCall("g_value_init", uintptr(unsafe.Pointer(&gValueSize)), gTypeString)
	gobject2_0.SysCall("g_value_set_string", uintptr(unsafe.Pointer(&gValueSize)), CStr(value))
	return &gValueSize
}

func (m *GValue) Free() {
	gobject2_0.SysCall("g_value_unset", uintptr(unsafe.Pointer(m)))
}

var gobject2_0 *linux.DnyLibrary

func init() {
	gobject2_0 = linux.LibLoad(linux.Libgobject2_0)
	gobject2_0.Table = []*imports.Table{
		imports.NewTable("g_object_ref", 0),
		imports.NewTable("g_object_unref", 0),
		imports.NewTable("g_object_ref_sink", 0),
		imports.NewTable("g_object_is_floating", 0),
		imports.NewTable("g_object_force_floating", 0),
		imports.NewTable("g_signal_stop_emission_by_name", 0),
		imports.NewTable("g_signal_connect_data", 0),
		imports.NewTable("g_type_from_name", 0),
		imports.NewTable("g_type_is_a", 0),
		imports.NewTable("g_type_name_from_instance", 0),
		imports.NewTable("g_type_name", 0),
		imports.NewTable("g_type_depth", 0),
		imports.NewTable("g_type_parent", 0),
		imports.NewTable("g_type_check_is_value_type", 0),
		imports.NewTable("g_object_set_data", 0),
		imports.NewTable("g_object_get_data", 0),
		imports.NewTable("g_new0", 0),
		// GValue
		imports.NewTable("g_value_init", 0),
		imports.NewTable("g_value_set_string", 0),
		imports.NewTable("g_value_unset", 0),
	}
	gobject2_0.SetLibClose()
	gobject2_0.MapperIndex()

	gTypeString = GTypeFromName(CStr("gchararray"))
}
