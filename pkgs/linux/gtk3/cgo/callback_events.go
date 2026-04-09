//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cgo

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
*/
import "C"
import (
	"unsafe"
)

type EventSignalName = string

const (
	EsnClicked               EventSignalName = "clicked"
	EsnChanged               EventSignalName = "changed"
	EsnActivate              EventSignalName = "activate"
	EsnKeyPressEvent         EventSignalName = "key-press-event"
	EsnKeyReleaseEvent       EventSignalName = "key-release-event"
	EsnButtonPressEvent      EventSignalName = "button-press-event"
	EsnEnterNotifyEvent      EventSignalName = "enter-notify-event"
	EsnLeaveNotifyEvent      EventSignalName = "leave-notify-event"
	EsnConfigureEvent        EventSignalName = "configure-event"
	EsnMapEvent              EventSignalName = "map"
	EsnDrawEvent             EventSignalName = "draw"
	EsnDragDataReceivedEvent EventSignalName = "drag-data-received"
	EsnDragDropEvent         EventSignalName = "drag-drop"
	EsnDragMotionEvent       EventSignalName = "drag-motion"
	EsnDragLeaveEvent        EventSignalName = "drag-leave"
	EsnDragDataDeleteEvent   EventSignalName = "drag-data-delete"
	EsnDragBeginEvent        EventSignalName = "drag-begin"
	EsnDragEndEvent          EventSignalName = "drag-end"
)

type TNotifyEvent func(sender *Widget)

type TTextChangedEvent func(sender *Widget, text string)

type TTextCommitEvent func(sender *Widget, text string)

type TTextKeyEvent func(sender *Widget, key *EventKey) bool

type TButtonPressEvent func(sender *Widget, event *EventButton)

type TLeaveEnterNotifyEvent func(sender *Widget, event *EventCrossing)

type TConfigureEvent func(sender *Widget, event *EventConfigure) bool

type TMapEvent func(sender *Widget)

type TDrawEvent func(sender *Widget, cr *Context) bool

// type drag event

type TDragDataReceivedEvent func(sender *Widget, context *DragContext, x, y int, data *SelectionData, info uint, time uint)
type TDragDropEvent func(sender *Widget, context *DragContext, x, y int, time uint) bool
type TDragMotionEvent func(sender *Widget, context *DragContext, x, y int, time uint) bool
type TDragLeaveEvent func(sender *Widget, context *DragContext, time uint)
type TDragDataDeleteOrBeginOrEndEvent func(sender *Widget, context *DragContext)

// make event

type CallbackContext struct {
	widget unsafe.Pointer
	input  any
	result any
}

type Callback struct {
	cb func(ctx *CallbackContext)
}

func MakeButtonPressEvent(cb TButtonPressEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			eventPtr := ctx.input.(unsafe.Pointer)
			event := ToEventButton(eventPtr)
			cb(wrapWidget(ToGoObject(ctx.widget)), event)
		},
	}
}

func MakeLeaveEnterNotifyEvent(cb TLeaveEnterNotifyEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			eventPtr := ctx.input.(unsafe.Pointer)
			event := ToEventCrossing(eventPtr)
			cb(wrapWidget(ToGoObject(ctx.widget)), event)
		},
	}
}

func MakeNotifyEvent(cb TNotifyEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			cb(wrapWidget(ToGoObject(ctx.widget)))
		},
	}
}

func MakeTextChangedEvent(cb TTextChangedEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			text := C.gtk_entry_get_text((*C.GtkEntry)(ctx.widget))
			cb(wrapWidget(ToGoObject(ctx.widget)), C.GoString(text))
		},
	}
}

func MakeTextCommitEvent(cb TTextCommitEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			text := C.gtk_entry_get_text((*C.GtkEntry)(ctx.widget))
			cb(wrapWidget(ToGoObject(ctx.widget)), C.GoString(text))
		},
	}
}

func MakeTextKeyEvent(cb TTextKeyEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			keyPtr := ctx.input.(unsafe.Pointer)
			key := AsKeyEvent(keyPtr)
			result := cb(wrapWidget(ToGoObject(ctx.widget)), key)
			ctx.result = result
		},
	}
}

func MakeConfigureEvent(cb TConfigureEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			eventPtr := ctx.input.(unsafe.Pointer)
			event := ToEventConfigure(eventPtr)
			result := cb(wrapWidget(ToGoObject(ctx.widget)), event)
			ctx.result = result
		},
	}
}

func MakeMapEvent(cb TMapEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			cb(wrapWidget(ToGoObject(ctx.widget)))
		},
	}
}

func MakeDrawEvent(cb TDrawEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			crPtr := ctx.input.(unsafe.Pointer)
			cr := WrapContext(uintptr(crPtr))
			result := cb(wrapWidget(ToGoObject(ctx.widget)), cr)
			ctx.result = result
		},
	}
}

func MakeDragDataReceivedEvent(cb TDragDataReceivedEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			args := ctx.input.([]any)
			context := AsDragContext(args[0].(unsafePointer))
			x, y := args[1].(int), args[2].(int)
			data := AsSelectionData(args[3].(unsafe.Pointer))
			info := args[4].(uint)
			time := args[5].(uint)
			cb(wrapWidget(ToGoObject(ctx.widget)), context, x, y, data, info, time)
		},
	}
}

func MakeDragDropEvent(cb TDragDropEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			args := ctx.input.([]any)
			context := AsDragContext(args[0].(unsafePointer))
			x, y := args[1].(int), args[2].(int)
			time := args[3].(uint)
			result := cb(wrapWidget(ToGoObject(ctx.widget)), context, x, y, time)
			ctx.result = result
		},
	}
}

func MakeDragMotionEvent(cb TDragMotionEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			args := ctx.input.([]any)
			context := AsDragContext(args[0].(unsafePointer))
			x, y := args[1].(int), args[2].(int)
			time := args[3].(uint)
			result := cb(wrapWidget(ToGoObject(ctx.widget)), context, x, y, time)
			ctx.result = result
		},
	}
}

func MakeDragLeaveEvent(cb TDragLeaveEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			args := ctx.input.([]any)
			context := AsDragContext(args[0].(unsafePointer))
			time := args[1].(uint)
			cb(wrapWidget(ToGoObject(ctx.widget)), context, time)
		},
	}
}

func MakeDragDataDeleteOrBeginOrEndEvent(cb TDragDataDeleteOrBeginOrEndEvent) *Callback {
	return &Callback{
		cb: func(ctx *CallbackContext) {
			context := AsDragContext(ctx.input.(unsafePointer))
			cb(wrapWidget(ToGoObject(ctx.widget)), context)
		},
	}
}
