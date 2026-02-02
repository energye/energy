package gtk3

import "C"
import (
	"unsafe"
)

type unsafePointer = unsafe.Pointer

func doOnEventHandler(widget, userData unsafePointer) {
	//println("doOnEventHandler userData:", uintptr(userData))
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget}
		cb.cb(context)
	}
}

func doOnKeyPressEventHandler(widget, event, userData unsafePointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
		result, ok := context.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doButtonPress(widget, event, userData unsafePointer) {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
	}
}

func doLeaveEnter(widget, event, userData unsafePointer) {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
	}
}

func doOnWindowConfigure(widget, event, userData unsafePointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
		result, ok := context.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doOnWindowDraw(widget, cr, userData unsafePointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: cr}
		cb.cb(context)
		result, ok := context.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doOnDragDataReceived(widget unsafePointer, context unsafePointer, x, y int, data unsafePointer,
	info uint, time uint, userData unsafePointer) {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		cbkCtx := &CallbackContext{widget: widget, input: []any{context, x, y, data, info, time}}
		cb.cb(cbkCtx)
	}
}

func doOnDragDrop(widget unsafePointer, context unsafePointer, x, y int, time uint, userData unsafePointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		cbkCtx := &CallbackContext{widget: widget, input: []any{context, x, y, time}}
		cb.cb(cbkCtx)
		result, ok := cbkCtx.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doOnDragMotion(widget unsafePointer, context unsafePointer, x, y int, time uint, userData unsafePointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		cbkCtx := &CallbackContext{widget: widget, input: []any{context, x, y, time}}
		cb.cb(cbkCtx)
		result, ok := cbkCtx.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doOnDragLeave(widget unsafePointer, context unsafePointer, time uint, userData unsafePointer) {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		cbkCtx := &CallbackContext{widget: widget, input: []any{context, time}}
		cb.cb(cbkCtx)
	}
}
