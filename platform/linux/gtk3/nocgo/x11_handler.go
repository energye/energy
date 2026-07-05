//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !cgo

package nocgo

import (
	"github.com/ebitengine/purego"
	"github.com/energye/energy/v3/platform/linux"
	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"unsafe"
)

var (
	x11Lib   *linux.DnyLibrary
	x11ErrFn func() bool
	x11IOFn  func() bool
)

// Go callbacks with C-compatible signatures (uintptr ≈ pointer)
//
// XErrorHandler: int (*)(Display*, XErrorEvent*)
func x11ErrorBridge(display, errorEvent uintptr) int {
	_ = display
	_ = errorEvent
	if fn := x11ErrFn; fn != nil && fn() {
		return 1
	}
	return 0
}

// XIOErrorHandler: int (*)(Display*)
func x11IOErrorBridge(display uintptr) int {
	_ = display
	if fn := x11IOFn; fn != nil && fn() {
		return 1
	}
	return 0
}

func init() {
	x11Lib = linux.LibLoad("libX11.so.6")
	if x11Lib != nil {
		x11Lib.Table = []*imports.Table{
			imports.NewTable("XSetErrorHandler", 0),
			imports.NewTable("XSetIOErrorHandler", 0),
		}
		api.SetOnReleaseCallback(func() {
			x11Lib.Dll.Release()
		})
		x11Lib.MapperIndex()
	}

	// Additional GDK functions needed by FlushDisplay / UseDefaultX11VisualForGtk / WindowX11ID.
	gdk3.Table = append(gdk3.Table,
	 imports.NewTable("gdk_window_get_display", 0),
	 imports.NewTable("gdk_screen_list_visuals", 0),
	 imports.NewTable("gdk_x11_display_get_xdisplay", 0),
	 imports.NewTable("gdk_x11_screen_get_screen_number", 0),
	 imports.NewTable("gdk_x11_screen_get_xscreen", 0),
	 imports.NewTable("gdk_x11_visual_get_xvisual", 0),
	 imports.NewTable("gdk_x11_window_get_xid", 0),
	 imports.NewTable("g_list_free", 0),
	)
	gdk3.MapperIndex()

	gtk3.Table = append(gtk3.Table,
		imports.NewTable("gtk_widget_get_window", 0),
	)
	gtk3.MapperIndex()
}

// WindowX11ID returns the X11 Window XID for a realized GTK window.
func WindowX11ID(win IWindow) uintptr {
	gdkWindow := gtk3.SysCall("gtk_widget_get_window", win.Instance())
	if gdkWindow == 0 {
		// Try realizing first
		gtk3.SysCall("gtk_widget_realize", win.Instance())
		gdkWindow = gtk3.SysCall("gtk_widget_get_window", win.Instance())
		if gdkWindow == 0 {
			return 0
		}
	}
	// gdk_x11_window_get_xid is in libgdk-3.so.0 (GDK X11 backend)
	return gdk3.SysCall("gdk_x11_window_get_xid", gdkWindow)
}

// UseDefaultX11VisualForGtk — 和 cgo 的 C 实现完全一致：
//   gdk_x11_display_get_xdisplay → DefaultVisual → 遍历 GList 匹配 visualid
func UseDefaultX11VisualForGtk(win IWindow) {
	screen := gdk3.SysCall("gdk_screen_get_default")
	if screen == 0 {
		return
	}
	visuals := gdk3.SysCall("gdk_screen_list_visuals", screen)
	if visuals == 0 {
		return
	}
	defer gdk3.SysCall("g_list_free", visuals)

	// 检查 X11 backend（同 cgo 的 GDK_IS_X11_SCREEN）
	display := gdk3.SysCall("gdk_screen_get_display", screen)
	if display == 0 {
		return
	}
	xdisplay := gdk3.SysCall("gdk_x11_display_get_xdisplay", display)
	if xdisplay == 0 {
		return // 不是 X11 后端，跳过
	}

	// DefaultVisual(xdisplay, screen_number) = ScreenOfDisplay(xdisplay, screen_number)->root_visual
	// 通过 gdk_x11_screen_get_xscreen 取 X11 Screen*，动态计算 root_visual 偏移
	x11Screen := gdk3.SysCall("gdk_x11_screen_get_xscreen", screen)
	if x11Screen == 0 {
		return
	}
	// X11 Screen 结构体偏移计算（适配 32/64 位）：
	//   ext_data(ptrSize) display(ptrSize) root(ptrSize)
	//   width(4) height(4) mwidth(4) mheight(4) ndepths(4) = 20
	//   pad(to align ptrSize) depths(ptrSize) root_depth(4)
	//   pad(to align ptrSize) root_visual(ptrSize)
	pad := ptrSize - 4 // 64 位需要 4 字节填充，32 位不需要
	if ptrSize == 4 {
		pad = 0
	}
	rootVisualOffset := 3*ptrSize + 20 + pad + ptrSize + 4 + pad
	rootVisual := *(*uintptr)(unsafe.Pointer(x11Screen + rootVisualOffset))
	if rootVisual == 0 {
		return
	}
	defaultVisualID := *(*uintptr)(unsafe.Pointer(rootVisual + ptrSize))

	// 遍历 GList 匹配 visualid
	cursor := visuals
	for cursor != 0 {
		gdkVisual := *(*uintptr)(unsafe.Pointer(cursor))
		if gdkVisual != 0 {
			xvisual := gdk3.SysCall("gdk_x11_visual_get_xvisual", gdkVisual)
			if xvisual != 0 {
				xvID := *(*uintptr)(unsafe.Pointer(xvisual + ptrSize))
				if xvID == defaultVisualID {
					gtk3.SysCall("gtk_widget_set_visual", win.Instance(), gdkVisual)
					break
				}
			}
		}
		cursor = *(*uintptr)(unsafe.Pointer(cursor + ptrSize))
	}
}

// FlushDisplay flushes the X11 display for a realized GTK window.
func FlushDisplay(win IWindow) {
	gdkWindow := gtk3.SysCall("gtk_widget_get_window", win.Instance())
	if gdkWindow == 0 {
		return
	}
	display := gdk3.SysCall("gdk_window_get_display", gdkWindow)
	if display == 0 {
		return
	}
	gdk3.SysCall("gdk_display_flush", display)
}

// SetX11ErrorHandlers registers X11 error handler callbacks and installs them.
// Works without CGo by loading libX11.so.6 at runtime via purego.
func SetX11ErrorHandlers(onError, onIOError func() bool) {
	x11ErrFn = onError
	x11IOFn = onIOError

	if x11Lib == nil {
		return // libX11 not available
	}

	// Create C-compatible function pointers from Go callbacks
	errCb := purego.NewCallback(x11ErrorBridge)
	ioErrCb := purego.NewCallback(x11IOErrorBridge)

	x11Lib.SysCall("XSetErrorHandler", errCb)
	x11Lib.SysCall("XSetIOErrorHandler", ioErrCb)
}
