package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import "unsafe"

// EventButton is a representation of GDK's GdkEventButton.
type EventButton struct {
	*Event
}

func NewEventButton() *EventButton {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventButton{}))
	ev := Event{ee}
	return &EventButton{&ev}
}

func ToEventButton(p unsafe.Pointer) *EventButton {
	return &EventButton{&Event{GdkEvent: (*C.GdkEvent)(p)}}
}

// EventButtonNewFromEvent returns an EventButton from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventButton. EventButtonNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventButtonNewFromEvent(event *Event) *EventButton {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventButton{&ev}
}

// Native returns a pointer to the underlying GdkEventButton.
func (v *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(unsafe.Pointer(v.Event.native()))
}

func (v *EventButton) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *EventButton) Y() float64 {
	c := v.native().y
	return float64(c)
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (v *EventButton) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (v *EventButton) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}

func (v *EventButton) Button() ButtonType {
	c := v.native().button
	return ButtonType(c)
}

func (v *EventButton) State() uint {
	c := v.native().state
	return uint(c)
}

// Time returns the time of the event in milliseconds.
func (v *EventButton) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventButton) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventButton) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventButton) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}
