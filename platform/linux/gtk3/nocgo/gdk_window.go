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

// WindowGetWidth is a wrapper around gdk_window_get_width().
func (m *GdkWindow) WindowGetWidth() (width int) {
	r := gdk3.SysCall("gdk_window_get_width", m.Instance())
	return int(r)
}

// WindowGetHeight is a wrapper around gdk_window_get_height().
func (m *GdkWindow) WindowGetHeight() (height int) {
	r := gdk3.SysCall("gdk_window_get_height", m.Instance())
	return int(r)
}

// GetRootOrigin is a wrapper around gdk_window_get_root_origin().
func (m *GdkWindow) GetRootOrigin() (x int, y int) {
	var cX, cY int32
	gdk3.SysCall("gdk_window_get_root_origin", m.Instance(), uintptr(unsafe.Pointer(&cX)), uintptr(unsafe.Pointer(&cY)))
	return int(cX), int(cY)
}

// GetOrigin is a wrapper around gdk_window_get_origin().
func (m *GdkWindow) GetOrigin() (x int, y int) {
	var cX, cY int32
	gdk3.SysCall("gdk_window_get_origin", m.Instance(), uintptr(unsafe.Pointer(&cX)), uintptr(unsafe.Pointer(&cY)))
	return int(cX), int(cY)
}

// GetDevicePosition is a wrapper around gdk_window_get_device_position().
func (m *GdkWindow) GetDevicePosition(device uintptr) (*GdkWindow, int, int, uint) {
	var x, y int32
	var mt uint32
	r := gdk3.SysCall("gdk_window_get_device_position", m.Instance(), device, uintptr(unsafe.Pointer(&x)), uintptr(unsafe.Pointer(&y)), uintptr(unsafe.Pointer(&mt)))
	var rw *GdkWindow
	if r != 0 {
		rw = &GdkWindow{}
		rw.instance = unsafe.Pointer(r)
	}
	return rw, int(x), int(y), uint(mt)
}

// SetOverrideRedirect is a wrapper around gdk_window_set_override_redirect().
func (m *GdkWindow) SetOverrideRedirect(overrideRedirect bool) {
	gdk3.SysCall("gdk_window_set_override_redirect", m.Instance(), ToCBool(overrideRedirect))
}

func (m *GdkWindow) GetState() int32 {
	state := gdk3.SysCall("gdk_window_get_state", m.Instance())
	return int32(state)
}
