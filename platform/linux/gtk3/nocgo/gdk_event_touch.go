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

// TGdkEventTouch mirrors the C GdkEventTouch struct layout.
type TGdkEventTouch struct {
	Type             int32
	Window           uintptr
	SendEvent        int8
	_                [3]byte // padding
	Time             uint32
	X                float64
	Y                float64
	Axes             uintptr
	State            uint32
	EmulatingPointer int32
	Device           uintptr
	XRoot            float64
	YRoot            float64
}

// EventTouch is a representation of GDK's GdkEventTouch.
type EventTouch struct {
	Event
	et *TGdkEventTouch
}

func AsEventTouch(ptr unsafe.Pointer) *EventTouch {
	if ptr == nil {
		return nil
	}
	m := new(EventTouch)
	m.instance = ptr
	m.et = (*TGdkEventTouch)(ptr)
	return m
}

func EventTouchNewFromEvent(event *Event) *EventTouch {
	return AsEventTouch(event.instance)
}

// Type returns the event type.
func (m *EventTouch) Type() EventType {
	return EventType(m.et.Type)
}

// Time returns the time of the event in milliseconds.
func (m *EventTouch) Time() uint32 {
	return m.et.Time
}

// X returns the x coordinate.
func (m *EventTouch) X() float64 {
	return m.et.X
}

// Y returns the y coordinate.
func (m *EventTouch) Y() float64 {
	return m.et.Y
}

// State returns the modifier key state.
func (m *EventTouch) State() uint {
	return uint(m.et.State)
}

// EmulatingPointer returns whether the touch event is emulating a pointer event.
func (m *EventTouch) EmulatingPointer() uint {
	return uint(m.et.EmulatingPointer)
}

// XRoot returns the x coordinate relative to the root of the screen.
func (m *EventTouch) XRoot() float64 {
	return m.et.XRoot
}

// YRoot returns the y coordinate relative to the root of the screen.
func (m *EventTouch) YRoot() float64 {
	return m.et.YRoot
}
