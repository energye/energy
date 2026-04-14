//go:build !cgo

package cocoa

import (
	"github.com/energye/energy/v3/platform/darwin/cocoa/nocgo"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

var NSApp = nocgo.AsNSApp()

func AsNSWindow(ptr unsafe.Pointer) INSWindow {
	return nocgo.AsNSWindow(ptr)
}

func AsNSVisualEffectView(ptr unsafe.Pointer) INSVisualEffectView {
	return nocgo.AsNSVisualEffectView(ptr)
}

func NewWindowDelegate(window INSWindow) INSWindowDelegate {
	return nocgo.NewWindowDelegate(window)
}

func NewToolBar(window INSWindow, delegate INSWindowDelegate, config ToolbarConfiguration) {
	nocgo.NewToolBar(window, delegate, config)
}

func WrapNSDraggingInfo(data unsafe.Pointer) INSDraggingInfo {
	return nocgo.WrapNSDraggingInfo(data)
}

func WrapNSPasteboard(data unsafe.Pointer) INSPasteboard {
	return nocgo.WrapNSPasteboard(data)
}
