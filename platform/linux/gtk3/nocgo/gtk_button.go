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

// Button is a representation of GTK's GtkButton.
type Button struct {
	Bin
}

func AsButton(ptr unsafe.Pointer) IButton {
	if ptr == nil {
		return nil
	}
	m := new(Button)
	m.instance = ptr
	return m
}

// NewButton is a wrapper around gtk_button_new().
func NewButton() IButton {
	r := gtk3.SysCall("gtk_button_new")
	if r == 0 {
		return nil
	}
	return AsButton(unsafe.Pointer(r))
}

// NewButtonWithLabel is a wrapper around gtk_button_new_with_label().
func NewButtonWithLabel(label string) IButton {
	r := gtk3.SysCall("gtk_button_new_with_label", CStr(label))
	if r == 0 {
		return nil
	}
	return AsButton(unsafe.Pointer(r))
}

// NewButtonWithMnemonic is a wrapper around gtk_button_new_with_mnemonic().
func NewButtonWithMnemonic(label string) IButton {
	r := gtk3.SysCall("gtk_button_new_with_mnemonic", CStr(label))
	if r == 0 {
		return nil
	}
	return AsButton(unsafe.Pointer(r))
}

// Clicked is a wrapper around gtk_button_clicked().
func (m *Button) Clicked() {
	gtk3.SysCall("gtk_button_clicked", m.Instance())
}

// SetRelief is a wrapper around gtk_button_set_relief().
func (m *Button) SetRelief(newStyle ReliefStyle) {
	gtk3.SysCall("gtk_button_set_relief", m.Instance(), uintptr(newStyle))
}

// GetRelief is a wrapper around gtk_button_get_relief().
func (m *Button) GetRelief() ReliefStyle {
	r := gtk3.SysCall("gtk_button_get_relief", m.Instance())
	return ReliefStyle(r)
}

// SetLabel is a wrapper around gtk_button_set_label().
func (m *Button) SetLabel(label string) {
	gtk3.SysCall("gtk_button_set_label", m.Instance(), CStr(label))
}

// GetLabel is a wrapper around gtk_button_get_label().
func (m *Button) GetLabel() (string, error) {
	r := gtk3.SysCall("gtk_button_get_label", m.Instance())
	return GoStr(r), nil
}

// SetUseUnderline is a wrapper around gtk_button_set_use_underline().
func (m *Button) SetUseUnderline(useUnderline bool) {
	gtk3.SysCall("gtk_button_set_use_underline", m.Instance(), ToCBool(useUnderline))
}

// GetUseUnderline is a wrapper around gtk_button_get_use_underline().
func (m *Button) GetUseUnderline() bool {
	r := gtk3.SysCall("gtk_button_get_use_underline", m.Instance())
	return ToGoBool(r)
}

// SetImage is a wrapper around gtk_button_set_image().
func (m *Button) SetImage(image IWidget) {
	var widget uintptr
	if image != nil {
		widget = image.Instance()
	}
	gtk3.SysCall("gtk_button_set_image", m.Instance(), widget)
}

// GetImage is a wrapper around gtk_button_get_image().
func (m *Button) GetImage() IWidget {
	r := gtk3.SysCall("gtk_button_get_image", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

// SetImagePosition is a wrapper around gtk_button_set_image_position().
func (m *Button) SetImagePosition(position PositionType) {
	gtk3.SysCall("gtk_button_set_image_position", m.Instance(), uintptr(position))
}

// GetImagePosition is a wrapper around gtk_button_get_image_position().
func (m *Button) GetImagePosition() PositionType {
	r := gtk3.SysCall("gtk_button_get_image_position", m.Instance())
	return PositionType(r)
}

// SetAlwaysShowImage is a wrapper around gtk_button_set_always_show_image().
func (m *Button) SetAlwaysShowImage(alwaysShow bool) {
	gtk3.SysCall("gtk_button_set_always_show_image", m.Instance(), ToCBool(alwaysShow))
}

// GetAlwaysShowImage is a wrapper around gtk_button_get_always_show_image().
func (m *Button) GetAlwaysShowImage() bool {
	r := gtk3.SysCall("gtk_button_get_always_show_image", m.Instance())
	return ToGoBool(r)
}

// GetEventWindow is a wrapper around gtk_button_get_event_window().
func (m *Button) GetEventWindow() (*GdkWindow, error) {
	r := gtk3.SysCall("gtk_button_get_event_window", m.Instance())
	if r == 0 {
		return nil, errNilPtr
	}
	return &GdkWindow{Object{instance: unsafe.Pointer(r)}}, nil
}

// SetOnClick is a callback for the "clicked" signal.
func (m *Button) SetOnClick(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnClicked, callback.C_trampoline_2_void, fn, 0)
}

// SetOnEnter is a callback for the "enter-notify-event" signal.
func (m *Button) SetOnEnter(fn TLeaveEnterNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnEnterNotifyEvent, callback.C_trampoline_3_gboolean, fn, 0)
}

// SetOnLeave is a callback for the "leave-notify-event" signal.
func (m *Button) SetOnLeave(fn TLeaveEnterNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnLeaveNotifyEvent, callback.C_trampoline_3_gboolean, fn, 0)
}
