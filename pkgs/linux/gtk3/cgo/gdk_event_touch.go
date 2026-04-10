package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"

import (
	. "github.com/energye/energy/v3/pkgs/linux/types"
	"unsafe"
)

// EventTouch is a representation of GDK's GdkEventTouch
type EventTouch struct {
	*Event
}

func NewEventTouch() *EventTouch {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventTouch{}))
	ev := Event{ee}
	return &EventTouch{&ev}
}

func ToEventTouch(p unsafe.Pointer) *EventTouch {
	return &EventTouch{&Event{GdkEvent: (*C.GdkEvent)(p)}}
}

// EventTouchNewFromEvent returns an EventTouch from an Event.
//
// Using widget.Connect() for a key related signal such as
// "touch-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventTouch. EventTouchNewFromEvent provides a means of creating
// an EventTouch from the Event.
func EventTouchNewFromEvent(event *Event) *EventTouch {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventTouch{&ev}
}

// Native returns a pointer to the underlying GdkEventTouch.
func (v *EventTouch) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventTouch) native() *C.GdkEventTouch {
	return (*C.GdkEventTouch)(unsafe.Pointer(v.Event.native()))
}

func (v *EventTouch) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventTouch) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventTouch) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *EventTouch) Y() float64 {
	c := v.native().y
	return float64(c)
}

func (v *EventTouch) State() uint {
	c := v.native().state
	return uint(c)
}

func (v *EventTouch) EmulatingPointer() uint {
	c := v.native().emulating_pointer
	return uint(c)
}

func (v *EventTouch) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

func (v *EventTouch) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}
