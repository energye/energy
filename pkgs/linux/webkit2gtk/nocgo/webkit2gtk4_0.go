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

package nocgo

import (
	"github.com/energye/energy/v3/pkgs/linux"
	"github.com/energye/energy/v3/pkgs/linux/callback"
	"github.com/energye/energy/v3/pkgs/linux/gtk3/nocgo"
	. "github.com/energye/energy/v3/pkgs/linux/types"
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
	webkit2gtk4_0.SysCall("webkit_web_view_set_background_color", m.Instance(), uintptr(unsafe.Pointer(&rgba)))
}

func (m *Webkit2) SetOnDragDataReceived(fn TDragDataReceivedEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDataReceivedEvent, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDrop(fn TDragDropEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDropEvent, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragMotion(fn TDragMotionEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragMotionEvent, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragLeave(fn TDragLeaveEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragLeaveEvent, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDataDelete(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDataDeleteEvent, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragBegin(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragBeginEvent, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragEnd(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragEndEvent, fn, 0)
	return signalHandlerID
}

var webkit2gtk4_0 *linux.DnyLibrary

func init() {
	tmpLibs := []string{linux.Libwebkit2gtk4_0, linux.Libwebkit2gtk4_0_37, linux.Libwebkit2gtk4_0_0}
	for _, lib := range tmpLibs {
		webkit2gtk4_0 = linux.LibLoad(lib)
		if webkit2gtk4_0 != nil {
			break
		}
	}
	webkit2gtk4_0.Table = []*imports.Table{
		imports.NewTable("webkit_web_view_set_background_color", 0),
	}
	webkit2gtk4_0.SetLibClose()
	webkit2gtk4_0.MapperIndex()
}
