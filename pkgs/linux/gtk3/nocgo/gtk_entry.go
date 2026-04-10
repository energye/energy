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
	"github.com/energye/energy/v3/pkgs/linux/callback"
	. "github.com/energye/energy/v3/pkgs/linux/types"
	"unsafe"
)

// Layout is a representation of GTK's GtkLayout.
type Entry struct {
	Widget
}

func AsEntry(ptr unsafe.Pointer) IEntry {
	if ptr == nil {
		return nil
	}
	m := new(Entry)
	m.instance = ptr
	return m
}

// NewEntry is a wrapper around gtk_entry_new().
func NewEntry() IEntry {
	r := gtk3.SysCall("gtk_entry_new")
	return AsEntry(unsafe.Pointer(r))
}

// SetText is a wrapper around gtk_entry_set_text().
func (m *Entry) SetText(text string) {
	gtk3.SysCall("gtk_entry_set_text", m.Instance(), CStr(text))
}

// GetText is a wrapper around gtk_entry_get_text().
func (m *Entry) GetText() string {
	c := gtk3.SysCall("gtk_entry_get_text", m.Instance())
	return GoStr(c)
}

// GetTextLength is a wrapper around gtk_entry_get_text_length().
func (m *Entry) GetTextLength() uint16 {
	c := gtk3.SysCall("gtk_entry_get_text_length", m.Instance())
	return uint16(c)
}

func (m *Entry) SetOnChanged(fn TTextChangedEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnChanged, fn, 0)
	return signalHandlerID
}

func (m *Entry) SetOnCommit(fn TTextCommitEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnActivate, fn, 0)
	return signalHandlerID
}

func (m *Entry) SetOnKeyPress(fn TTextKeyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnKeyPressEvent,
		fn, 0)
	return signalHandlerID
}

func (m *Entry) SetOnKeyRelease(fn TTextKeyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnKeyReleaseEvent, fn, 0)
	return signalHandlerID
}
