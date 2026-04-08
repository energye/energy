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
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

// DragContext is a representation of GDK's GdkDragContext.
type DragContext struct {
	Object
}

func AsDragContext(ptr unsafe.Pointer) IDragContext {
	if ptr == nil {
		return nil
	}
	m := new(DragContext)
	m.instance = ptr
	return m
}

func (m *DragContext) ListTargets() IList {
	r := gdk3.SysCall("gdk_drag_context_list_targets", m.Instance())
	return AsList(unsafe.Pointer(r))
}

func (m *DragContext) Finish(success bool, del bool, time uint) {
	gdk3.SysCall("gtk_drag_finish", m.Instance(), ToCBool(success), ToCBool(del), uintptr(time))
}

func (m *DragContext) Status(actions DragAction, time uint) {
	gdk3.SysCall("gdk_drag_status", m.Instance(), uintptr(actions), uintptr(time))
}
