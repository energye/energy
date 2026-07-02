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

// EventBox is a representation of GTK's GtkEventBox.
type EventBox struct {
	Bin
}

func AsEventBox(ptr unsafe.Pointer) IEventBox {
	if ptr == nil {
		return nil
	}
	m := new(EventBox)
	m.instance = ptr
	return m
}

// NewEventBox is a wrapper around gtk_event_box_new().
func NewEventBox() IEventBox {
	r := gtk3.SysCall("gtk_event_box_new")
	if r == 0 {
		return nil
	}
	return AsEventBox(unsafe.Pointer(r))
}

// SetAboveChild is a wrapper around gtk_event_box_set_above_child().
func (m *EventBox) SetAboveChild(aboveChild bool) {
	gtk3.SysCall("gtk_event_box_set_above_child", m.Instance(), ToCBool(aboveChild))
}

// GetAboveChild is a wrapper around gtk_event_box_get_above_child().
func (m *EventBox) GetAboveChild() bool {
	r := gtk3.SysCall("gtk_event_box_get_above_child", m.Instance())
	return ToGoBool(r)
}

// SetVisibleWindow is a wrapper around gtk_event_box_set_visible_window().
func (m *EventBox) SetVisibleWindow(visibleWindow bool) {
	gtk3.SysCall("gtk_event_box_set_visible_window", m.Instance(), ToCBool(visibleWindow))
}

// GetVisibleWindow is a wrapper around gtk_event_box_get_visible_window().
func (m *EventBox) GetVisibleWindow() bool {
	r := gtk3.SysCall("gtk_event_box_get_visible_window", m.Instance())
	return ToGoBool(r)
}

// SetOnClick is a callback for the "button-press-event" signal.
func (m *EventBox) SetOnClick(fn TButtonPressEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnButtonPressEvent, callback.C_trampoline_3_gboolean, fn, 0)
}

// SetOnEnter is a callback for the "enter-notify-event" signal.
func (m *EventBox) SetOnEnter(fn TLeaveEnterNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnEnterNotifyEvent, callback.C_trampoline_3_gboolean, fn, 0)
}

// SetOnLeave is a callback for the "leave-notify-event" signal.
func (m *EventBox) SetOnLeave(fn TLeaveEnterNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnLeaveNotifyEvent, callback.C_trampoline_3_gboolean, fn, 0)
}
