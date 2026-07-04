//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type LevelBar struct {
	Widget
}

func AsLevelBar(ptr unsafe.Pointer) *LevelBar {
	if ptr == nil {
		return nil
	}
	m := new(LevelBar)
	m.instance = ptr
	return m
}

func NewLevelBar() *LevelBar {
	r := gtk3.SysCall("gtk_level_bar_new")
	if r == 0 {
		return nil
	}
	return AsLevelBar(unsafe.Pointer(r))
}

func (m *LevelBar) SetValue(value float64) {
	registerGtkFloatFuncs()
	gtkLevelBarSetValue(m.Instance(), value)
}

func (m *LevelBar) GetValue() float64 {
	registerGtkFloatFuncs()
	return gtkLevelBarGetValue(m.Instance())
}

func (m *LevelBar) SetMinValue(value float64) {
	registerGtkFloatFuncs()
	gtkLevelBarSetMinValue(m.Instance(), value)
}

func (m *LevelBar) GetMinValue() float64 {
	registerGtkFloatFuncs()
	return gtkLevelBarGetMinValue(m.Instance())
}

func (m *LevelBar) SetMaxValue(value float64) {
	registerGtkFloatFuncs()
	gtkLevelBarSetMaxValue(m.Instance(), value)
}

func (m *LevelBar) GetMaxValue() float64 {
	registerGtkFloatFuncs()
	return gtkLevelBarGetMaxValue(m.Instance())
}

func (m *LevelBar) SetMode(mode LevelBarMode) {
	gtk3.SysCall("gtk_level_bar_set_mode", m.Instance(), uintptr(mode))
}

func (m *LevelBar) GetMode() LevelBarMode {
	r := gtk3.SysCall("gtk_level_bar_get_mode", m.Instance())
	return LevelBarMode(r)
}

func (m *LevelBar) SetOnOffsetChanged(fn TOffsetChangedEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnOffsetChanged, callback.C_trampoline_3_void, fn, 0)
}
