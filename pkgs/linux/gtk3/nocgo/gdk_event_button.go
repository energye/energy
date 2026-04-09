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

type TGdkEventButton struct {
	Type      TGdkEventType
	Window    PGdkWindow
	SendEvent GInt8
	Time      GUint32
	X         GDouble
	Y         GDouble
	Axes      PGDouble
	State     GUint32
	Button    GUint
	Device    PGdkDevice
	XRoot     GDouble
	YRoot     GDouble
}

// EventButton is a representation of GDK's GdkEventButton.
type EventButton struct {
	Event
	eb *TGdkEventButton
}

func AsEventButton(ptr unsafe.Pointer) IEventButton {
	if ptr == nil {
		return nil
	}
	m := new(EventButton)
	m.instance = ptr
	m.eb = (*TGdkEventButton)(ptr)
	return m
}

func (m *EventButton) X() float64 {
	return m.eb.X
}

func (m *EventButton) Y() float64 {
	return m.eb.Y
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (m *EventButton) XRoot() float64 {
	return m.eb.XRoot
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (m *EventButton) YRoot() float64 {
	return m.eb.YRoot
}

func (m *EventButton) Button() ButtonType {
	return ButtonType(m.eb.Button)
}

func (m *EventButton) State() uint {
	return uint(m.eb.State)
}

// Time returns the time of the event in milliseconds.
func (m *EventButton) Time() uint32 {
	return uint32(m.eb.Time)
}

func (m *EventButton) Type() EventType {
	return EventType(m.eb.Type)
}
