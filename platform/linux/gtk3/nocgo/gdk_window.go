package nocgo

import (
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// GdkWindow is a representation of GDK's GdkWindow.
type GdkWindow struct {
	Object
}

func AsGdkWindow(ptr unsafe.Pointer) types.IGdkWindow {
	if ptr == nil {
		return nil
	}
	m := new(GdkWindow)
	m.instance = ptr
	return m
}

func (m *GdkWindow) SetDecorations(decorations types.TGdkWMDecoration) {
	gdk3.SysCall("gdk_window_set_decorations", m.Instance(), uintptr(decorations))
}
