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
#cgo LDFLAGS: -lX11
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
#include <gtk/gtk.h>
#include <X11/Xlib.h>
#include <stdio.h>

// Go-exported trampolines (defined in Go code below)
extern int goBridgeX11ErrorHandler(Display*, XErrorEvent*);
extern int goBridgeXIOErrorHandler(Display*);

static int cX11ErrorHandler(Display* display, XErrorEvent* error) {
	return goBridgeX11ErrorHandler(display, error);
}

static int cXIOErrorHandler(Display* display) {
	return goBridgeXIOErrorHandler(display);
}

static void installX11Handlers() {
	XSetErrorHandler(cX11ErrorHandler);
	XSetIOErrorHandler(cXIOErrorHandler);
}

// Use the default X11 visual on a GtkWindow before it is realized.
// GTK+ > 3.15.1 uses an X11 visual optimized for OpenGL which breaks CEF.
// This finds the default X11 visual and applies it to the widget.
// Ref: https://github.com/cztomczak/cefcapi
static void useDefaultX11Visual(GtkWidget* widget) {
	GdkScreen* screen = gdk_screen_get_default();
	GList* visuals = gdk_screen_list_visuals(screen);

	if (GDK_IS_X11_SCREEN(screen)) {
		Display* xdisplay = gdk_x11_display_get_xdisplay(gdk_screen_get_display(screen));
		int screen_number = gdk_x11_screen_get_screen_number(screen);
		Visual* default_xvisual = DefaultVisual(xdisplay, screen_number);

		if (default_xvisual != NULL) {
			GList* cursor = visuals;
			while (cursor != NULL) {
				GdkVisual* gdk_visual = GDK_X11_VISUAL(cursor->data);
				Visual* xvisual = gdk_x11_visual_get_xvisual(gdk_visual);
				if (xvisual != NULL && default_xvisual->visualid == xvisual->visualid) {
					gtk_widget_set_visual(widget, gdk_visual);
					break;
				}
				cursor = cursor->next;
			}
		}
	}

	g_list_free(visuals);
}

// Flush the X11 display to ensure the underlying X window is created.
static void flushDisplay(GtkWidget* widget) {
	GdkWindow* gdk_window = gtk_widget_get_window(widget);
	if (gdk_window == NULL) return;
	GdkDisplay* display = gdk_window_get_display(gdk_window);
	gdk_display_flush(display);
}
*/
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

var (
	x11ErrFn func() bool
	x11IOFn  func() bool
)

// WindowX11ID returns the X11 Window XID for a realized GTK window.
// CEF on Linux requires the native XID, not a GtkWidget pointer.
func WindowX11ID(win IWindow) uintptr {
	w := (*C.GtkWidget)(unsafe.Pointer(win.Instance()))
	if C.gtk_widget_get_realized(w) == 0 {
		C.gtk_widget_realize(w)
	}
	gdkWin := C.gtk_widget_get_window(w)
	if gdkWin == nil {
		return 0
	}
	return uintptr(C.gdk_x11_window_get_xid(gdkWin))
}

//export goBridgeX11ErrorHandler
func goBridgeX11ErrorHandler(display *C.Display, error *C.XErrorEvent) C.int {
	_ = display
	_ = error
	if fn := x11ErrFn; fn != nil && fn() {
		return 1
	}
	return 0
}

//export goBridgeXIOErrorHandler
func goBridgeXIOErrorHandler(display *C.Display) C.int {
	_ = display
	if fn := x11IOFn; fn != nil && fn() {
		return 1
	}
	return 0
}

// SetX11ErrorHandlers registers X11 error handler callbacks and installs them.
//
//	onError   — called on X11 protocol errors; return true if handled.
//	onIOError — called on X11 I/O errors (e.g. lost connection); return true if handled.
//
// Passing nil for either removes the callback (and the handler returns 0 / ignores).
func SetX11ErrorHandlers(onError, onIOError func() bool) {
	x11ErrFn = onError
	x11IOFn = onIOError
	C.installX11Handlers()
}

// UseDefaultX11VisualForGtk overrides the GTK window's visual with the default
// X11 visual. Must be called after the window is created but before it is
// shown/realized. GTK+ > 3.15.1 uses an OpenGL-optimized visual that breaks CEF.
// See: https://github.com/cztomczak/cefcapi
func UseDefaultX11VisualForGtk(win IWindow) {
	C.useDefaultX11Visual((*C.GtkWidget)(unsafe.Pointer(win.Instance())))
}

// FlushDisplay synchronizes the X11 display for a realized GTK window.
// Uses gdk_display_sync to wait for all pending X11 requests to be processed,
// ensuring correct ordering of window creation and focus management.
func FlushDisplay(win IWindow) {
	C.flushDisplay((*C.GtkWidget)(unsafe.Pointer(win.Instance())))
}
