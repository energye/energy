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

// TGdkEventMotion mirrors the C GdkEventMotion struct layout.
type TGdkEventMotion struct {
	Type      int32
	Window    uintptr
	SendEvent int8
	_         [3]byte // padding
	Time      uint32
	X         float64
	Y         float64
	Axes      uintptr
	State     uint32
	IsHint    int16
	_         [2]byte // padding
	Device    uintptr
	XRoot     float64
	YRoot     float64
}

// EventMotion is a representation of GDK's GdkEventMotion.
type EventMotion struct {
	Event
	em *TGdkEventMotion
}

func AsEventMotion(ptr unsafe.Pointer) *EventMotion {
	if ptr == nil {
		return nil
	}
	m := new(EventMotion)
	m.instance = ptr
	m.em = (*TGdkEventMotion)(ptr)
	return m
}

func NewEventMotionFromEvent(event *Event) *EventMotion {
	return AsEventMotion(event.instance)
}

// MotionVal returns the x, y coordinates of the motion event.
func (m *EventMotion) MotionVal() (float64, float64) {
	return m.em.X, m.em.Y
}

// MotionValRoot returns the x_root, y_root coordinates.
func (m *EventMotion) MotionValRoot() (float64, float64) {
	return m.em.XRoot, m.em.YRoot
}

// Time returns the time of the event in milliseconds.
func (m *EventMotion) Time() uint32 {
	return m.em.Time
}

// Type returns the event type.
func (m *EventMotion) Type() EventType {
	return EventType(m.em.Type)
}

// State returns the modifier key state.
func (m *EventMotion) State() uint {
	return uint(m.em.State)
}

// SetXY sets the x, y coordinates.
func (m *EventMotion) SetXY(x, y float64) {
	m.em.X = x
	m.em.Y = y
}

// SetXYRoot sets the x_root, y_root coordinates.
func (m *EventMotion) SetXYRoot(x, y float64) {
	m.em.XRoot = x
	m.em.YRoot = y
}

// SetTime sets the time of the event.
func (m *EventMotion) SetTime(time uint32) {
	if time == 0 {
		m.em.Time = 0 // CURRENT_TIME
	} else {
		m.em.Time = time
	}
}
