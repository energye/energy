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

type GdkEventConfigure struct {
	Type      TGdkEventType
	Window    PGdkWindow
	SendEvent GInt8
	X         GInt
	Y         GInt
	Width     GInt
	Height    GInt
}

// EventConfigure is a representation of GDK's GdkEventConfigure.
type EventConfigure struct {
	Event
	ec *GdkEventConfigure
}

func AsEventConfigure(ptr unsafe.Pointer) IEventConfigure {
	if ptr == nil {
		return nil
	}
	m := new(EventConfigure)
	m.instance = ptr
	m.ec = (*GdkEventConfigure)(ptr)
	return m
}

func (m *EventConfigure) X() int {
	return int(m.ec.X)
}

func (m *EventConfigure) Y() int {
	return int(m.ec.Y)
}

func (m *EventConfigure) Width() int {
	return int(m.ec.Width)
}

func (m *EventConfigure) Height() int {
	return int(m.ec.Height)
}

func (m *EventConfigure) Type() EventType {
	return EventType(m.ec.Type)
}
