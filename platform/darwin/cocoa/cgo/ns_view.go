package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "cocoa.h"
#include "ns_view.h"
*/
import "C"
import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

type NSView struct {
	NSResponder
	config TItemBase
}

func AsNSView(ptr unsafe.Pointer) INSView {
	if ptr == nil {
		return nil
	}
	m := new(NSView)
	m.SetInstance(ptr)
	return m
}

func NewNSView(config TItemBase) INSView {
	if config.Identifier == "" {
		config.Identifier = nextSerialNumber("NSView")
	}
	cView := C.NewCustomView()
	m := new(NSView)
	m.instance = Pointer(cView)
	m.config = config
	return m
}

func (m *NSView) Identifier() string {
	return m.config.Identifier
}
