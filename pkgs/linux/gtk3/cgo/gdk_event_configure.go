package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import "unsafe"

// EventConfigure is a representation of GDK's GdkEventConfigure.
type EventConfigure struct {
	*Event
}

func NewEventConfigure() *EventConfigure {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventConfigure{}))
	ev := Event{ee}
	return &EventConfigure{&ev}
}

func ToEventConfigure(p unsafe.Pointer) *EventConfigure {
	return &EventConfigure{&Event{GdkEvent: (*C.GdkEvent)(p)}}
}

// EventConfigureNewFromEvent returns an EventConfigure from an Event.
//
// Using widget.Connect() for the
// "configure-event" signal results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventConfigure. EventConfigureNewFromEvent provides a means of creating
// an EventConfigure from the Event.
func EventConfigureNewFromEvent(event *Event) *EventConfigure {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventConfigure{&ev}
}

// Native returns a pointer to the underlying GdkEventConfigure.
func (v *EventConfigure) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventConfigure) native() *C.GdkEventConfigure {
	return (*C.GdkEventConfigure)(unsafe.Pointer(v.Event.native()))
}

func (v *EventConfigure) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventConfigure) X() int {
	c := v.native().x
	return int(c)
}

func (v *EventConfigure) Y() int {
	c := v.native().y
	return int(c)
}

func (v *EventConfigure) Width() int {
	c := v.native().width
	return int(c)
}

func (v *EventConfigure) Height() int {
	c := v.native().height
	return int(c)
}
