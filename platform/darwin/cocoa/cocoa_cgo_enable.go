//go:build cgo

package cocoa

import (
	"github.com/energye/energy/v3/platform/darwin/cocoa/cgo"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

var NSApp = cgo.AsNSApp()

func AsNSWindow(ptr unsafe.Pointer) INSWindow {
	return cgo.AsNSWindow(ptr)
}

func AsNSVisualEffectView(ptr unsafe.Pointer) INSVisualEffectView {
	return cgo.AsNSVisualEffectView(ptr)
}

func NewWindowDelegate(window INSWindow) INSWindowDelegate {
	return cgo.NewWindowDelegate(window)
}

func NewToolBar(window INSWindow, delegate INSWindowDelegate, config ToolbarConfiguration) {
	cgo.NewToolBar(window, delegate, config)
}
