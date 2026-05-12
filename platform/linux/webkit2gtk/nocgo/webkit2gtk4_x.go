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
	"github.com/energye/energy/v3/platform/linux"
	"github.com/energye/energy/v3/platform/linux/callback"
	"github.com/energye/energy/v3/platform/linux/gtk3/nocgo"
	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/types/colors"
	"unsafe"
)

type Webkit2 struct {
	nocgo.Widget
}

func AsWebkit2(ptr unsafe.Pointer) IWebkit2 {
	if ptr == nil {
		return nil
	}
	m := new(Webkit2)
	m.SetInstance(ptr)
	return m
}

func (m *Webkit2) OpenDevTools() {
	settings := webkit2gtk4_x.SysCall("webkit_web_view_get_settings", m.Instance())
	if settings == 0 {
		return
	}
	enable := webkit2gtk4_x.SysCall("webkit_settings_get_enable_developer_extras", settings)
	if enable == 0 {
		return
	}
	inspector := webkit2gtk4_x.SysCall("webkit_web_view_get_inspector", m.Instance())
	if inspector == 0 {
		return
	}
	webkit2gtk4_x.SysCall("webkit_web_inspector_show", inspector)
}
func (m *Webkit2) SetBackgroundColor(color *colors.TARGB) {
	if color == nil {
		return
	}
	cR := float64(color.R) / 255.0
	cG := float64(color.G) / 255.0
	cB := float64(color.B) / 255.0
	cA := float64(color.A) / 255.0
	rgba := GdkRGBA{
		Red:   cR,
		Green: cG,
		Blue:  cB,
		Alpha: cA,
	}
	webkit2gtk4_x.SysCall("webkit_web_view_set_background_color", m.Instance(), uintptr(unsafe.Pointer(&rgba)))
}

func (m *Webkit2) SetOnDragDataReceived(fn TDragDataReceivedEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDataReceivedEvent,
		callback.C_trampoline_8_void_drag_data_received, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDrop(fn TDragDropEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDropEvent,
		callback.C_trampoline_6_gboolean_drag_drop_motion, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragMotion(fn TDragMotionEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragMotionEvent,
		callback.C_trampoline_6_gboolean_drag_drop_motion, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragLeave(fn TDragLeaveEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragLeaveEvent,
		callback.C_trampoline_4_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDataDelete(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDataDeleteEvent,
		callback.C_trampoline_3_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragBegin(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragBeginEvent,
		callback.C_trampoline_4_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragEnd(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragEndEvent,
		callback.C_trampoline_4_void, fn, 0)
	return signalHandlerID
}

var webkit2gtk4_x *linux.DnyLibrary

func init() {
	//tmpLibs := []string{linux.Libwebkit2gtk4_0_37, linux.Libwebkit2gtk4_0}
	tmpLibs := []string{linux.Libwebkit2gtk4_1_0, linux.Libwebkit2gtk4_0_37}
	for _, lib := range tmpLibs {
		webkit2gtk4_x = linux.LibLoad(lib)
		if webkit2gtk4_x != nil {
			break
		}
	}
	webkit2gtk4_x.Table = []*imports.Table{
		imports.NewTable("webkit_web_view_set_background_color", 0),
		imports.NewTable("webkit_web_view_get_settings", 0),
		imports.NewTable("webkit_settings_get_enable_developer_extras", 0),
		//imports.NewTable("webkit_settings_set_enable_developer_extras", 0),
		imports.NewTable("webkit_web_view_get_inspector", 0),
		imports.NewTable("webkit_web_inspector_show", 0),
	}
	webkit2gtk4_x.SetLibClose()
	webkit2gtk4_x.MapperIndex()
}
