package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "ns_window.h"

extern GoArguments* doOnWindowDelegateEvent(TCallbackContext *CContext);

*/
import "C"
import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

//export doOnWindowDelegateEvent
func doOnWindowDelegateEvent(cContext *C.TCallbackContext) *C.GoArguments {
	ctx := CCallbackContextToGo(cContext)
	result := doDispatchEvent(ctx)
	if result != nil {
		return result.ToOC()
	} else {
		return nil
	}
}

func createEventCallback() C.TEventCallback {
	return (C.TEventCallback)(C.doOnWindowDelegateEvent)
}

type NSWindowDelegate struct {
	NSObject
}

func AsNSWindowDelegate(ptr unsafe.Pointer) INSWindowDelegate {
	if ptr == nil {
		return nil
	}
	m := new(NSWindowDelegate)
	m.SetInstance(ptr)
	return m
}

func NewWindowDelegate(window INSWindow) INSWindowDelegate {
	if window == nil {
		return nil
	}
	windowEventCallback := createEventCallback()
	windowDelegate := C.CreateWindowDelegate(unsafe.Pointer(window.Instance()), windowEventCallback)
	return AsNSWindowDelegate(unsafe.Pointer(windowDelegate))
}
