package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// EventCrossing is a representation of GDK's GdkEventCrossing.
type EventCrossing struct {
	*Event
}

func NewEventCrossing() *EventCrossing {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventCrossing{}))
	ev := Event{ee}
	return &EventCrossing{&ev}
}

func AsEventCrossing(p unsafe.Pointer) IEventCrossing {
	return &EventCrossing{&Event{GdkEvent: (*C.GdkEvent)(p)}}
}

// EventCrossingNewFromEvent returns an EventCrossing from an Event.
//
// Using widget.Connect() for a key related signal such as
// "enter-notify-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventCrossing. EventCrossingNewFromEvent provides a means of creating
// an EventCrossing from the Event.
func EventCrossingNewFromEvent(event *Event) *EventCrossing {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventCrossing{&ev}
}

// Native returns a pointer to the underlying GdkEventCrossing.
func (v *EventCrossing) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventCrossing) native() *C.GdkEventCrossing {
	return (*C.GdkEventCrossing)(unsafe.Pointer(v.Event.native()))
}

func (v *EventCrossing) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *EventCrossing) Y() float64 {
	c := v.native().y
	return float64(c)
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (v *EventCrossing) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (v *EventCrossing) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}

func (v *EventCrossing) State() uint {
	c := v.native().state
	return uint(c)
}

// Time returns the time of the event in milliseconds.
func (v *EventCrossing) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventCrossing) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventCrossing) Mode() CrossingMode {
	c := v.native().mode
	return CrossingMode(c)
}

func (v *EventCrossing) Detail() NotifyType {
	c := v.native().detail
	return NotifyType(c)
}

func (v *EventCrossing) Focus() bool {
	c := v.native().focus
	return GoBool(c)
}
