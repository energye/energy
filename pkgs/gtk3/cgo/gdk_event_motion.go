package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import "unsafe"

type EventMotion struct {
	*Event
}

func NewEventMotion() *EventMotion {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventMotion{}))
	ev := Event{ee}
	return &EventMotion{&ev}
}

// NewEventMotionFromEvent returns an EventMotion from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventMotion. EventMotionNewFromEvent provides a means of creating
// an EventKey from the Event.
func NewEventMotionFromEvent(event *Event) *EventMotion {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventMotion{&ev}
}

// Native returns a pointer to the underlying GdkEventMotion.
func (v *EventMotion) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventMotion) native() *C.GdkEventMotion {
	return (*C.GdkEventMotion)(unsafe.Pointer(v.Event.native()))
}

func (v *EventMotion) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventMotion) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

// Time returns the time of the event in milliseconds.
func (v *EventMotion) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventMotion) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

// A bit-mask representing the state of the modifier keys (e.g. Control, Shift
// and Alt) and the pointer buttons. See gdk.ModifierType constants.
func (v *EventMotion) State() ModifierType {
	c := v.native().state
	return ModifierType(c)
}

func (v *EventMotion) SetWindowRoot(root *GdkWindow) {
	v.native().window = root.native()
}

func (v *EventMotion) SetXY(x, y float64) {
	v.native().x = C.double(x)
	v.native().y = C.double(y)
}

func (v *EventMotion) SetXYRoot(x, y float64) {
	v.native().x_root = C.double(x)
	v.native().y_root = C.double(y)
}

func (v *EventMotion) SetTime(time uint32) {
	if time == 0 {
		v.native().time = CURRENT_TIME
	} else {
		v.native().time = C.guint32(time)
	}
}
