package cgo

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

// Event is a representation of GDK's GdkEvent.
type Event struct {
	GdkEvent *C.GdkEvent
}

func (m *Event) Instance() uintptr {
	return uintptr(unsafe.Pointer(m.GdkEvent))
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return v.GdkEvent
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *Event) Free() {
	C.gdk_event_free(v.native())
}

func (v *Event) ScanCode() int {
	return int(C.gdk_event_get_scancode(v.native()))
}

func (v *Event) ToEvent() *Event {
	return v
}

// EventKey is a representation of GDK's GdkEventKey.
type EventKey struct {
	*Event
}

func NewEventKey() *EventKey {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventKey{}))
	ev := Event{ee}
	return &EventKey{&ev}
}

func AsEventKey(p unsafe.Pointer) *EventKey {
	return &EventKey{Event: &Event{GdkEvent: (*C.GdkEvent)(p)}}
}

// NewEventKeyFromEvent returns an EventKey from an Event.
//
// Using widget.Connect() for a key related signal such as
// "key-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventKey. EventKeyNewFromEvent provides a means of creating
// an EventKey from the Event.
func NewEventKeyFromEvent(event *Event) *EventKey {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventKey{&ev}
}

// Native returns a pointer to the underlying GdkEventKey.
func (v *EventKey) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventKey) native() *C.GdkEventKey {
	return (*C.GdkEventKey)(unsafe.Pointer(v.Event.native()))
}

func (v *EventKey) KeyVal() uint {
	c := v.native().keyval
	return uint(c)
}

func (v *EventKey) HardwareKeyCode() uint16 {
	c := v.native().hardware_keycode
	return uint16(c)
}

func (v *EventKey) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventKey) State() uint {
	c := v.native().state
	return uint(c)
}
