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
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type GdkEventKey struct {
	Type            TGdkEventType
	Window          PGdkWindow
	SendEvent       GInt8
	Time            GUint32
	State           GUint32
	Keyval          GUint
	Length          GInt
	String          GPChar
	HardwareKeycode GUInt16
	Group           GUInt8
	IsModifier      GUInt8
}

// Event is a representation of GDK's GdkEvent.
type Event struct {
	instance unsafe.Pointer
}

func AsEvent(ptr unsafe.Pointer) IEvent {
	if ptr == nil {
		return nil
	}
	m := new(Event)
	m.instance = ptr
	return m
}

func (m *Event) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *Event) Free() {
	gdk3.SysCall("gdk_event_free", m.Instance())
}

func (m *Event) ScanCode() int {
	return int(gdk3.SysCall("gdk_event_get_scancode", m.Instance()))
}

type EventKey struct {
	Event
	ek *GdkEventKey
}

func AsEventKey(p unsafe.Pointer) IEventKey {
	if p == nil {
		return nil
	}
	m := &EventKey{}
	m.instance = p
	m.ek = (*GdkEventKey)(p)
	return m
}
func (m *EventKey) KeyVal() uint {
	return uint(m.ek.Keyval)
}

func (m *EventKey) HardwareKeyCode() uint16 {
	return uint16(m.ek.HardwareKeycode)
}

func (m *EventKey) Type() EventType {
	return EventType(m.ek.Type)
}

func (m *EventKey) State() uint {
	return uint(m.ek.State)
}
