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

type GdkEventCrossing struct {
	Type      TGdkEventType
	Window    PGdkWindow
	SendEvent GInt8
	Subwindow PGdkWindow
	Time      GUint32
	X         GDouble
	Y         GDouble
	XRoot     GDouble
	YRoot     GDouble
	Mode      TGdkCrossingMode
	Detail    TGdkNotifyType
	Focus     GBoolean
	State     TGdkModifierType
}

// EventCrossing is a representation of GDK's GdkEventCrossing.
type EventCrossing struct {
	Event
	ec *GdkEventCrossing
}

func AsEventCrossing(ptr unsafe.Pointer) IEventCrossing {
	if ptr == nil {
		return nil
	}
	m := new(EventCrossing)
	m.instance = ptr
	m.ec = (*GdkEventCrossing)(ptr)
	return m
}

func (m *EventCrossing) X() float64 {
	return m.ec.X
}

func (m *EventCrossing) Y() float64 {
	return m.ec.Y
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (m *EventCrossing) XRoot() float64 {
	return m.ec.XRoot
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (m *EventCrossing) YRoot() float64 {
	return m.ec.YRoot
}

func (m *EventCrossing) State() uint {
	return uint(m.ec.State)
}

// Time returns the time of the event in milliseconds.
func (m *EventCrossing) Time() uint32 {
	return m.ec.Time
}

func (m *EventCrossing) Type() EventType {
	return EventType(m.ec.Type)
}

func (m *EventCrossing) Mode() CrossingMode {
	return CrossingMode(m.ec.Mode)
}

func (m *EventCrossing) Detail() NotifyType {
	return NotifyType(m.ec.Detail)
}

func (m *EventCrossing) Focus() bool {
	return ToGoBool(uintptr(m.ec.Focus))
}
