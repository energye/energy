//go:build !cgo

package cocoa

import (
	"github.com/energye/energy/v3/pkgs/darwin/cocoa/nocgo"
	. "github.com/energye/energy/v3/pkgs/darwin/types"
	"unsafe"
)

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
