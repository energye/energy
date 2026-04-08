package cgo

import "C"
import "unsafe"
import (
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
)

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <stdlib.h>
#include <gtk/gtk.h>
#include <gtk_header_bar.go.h>
*/
import "C"

// HeaderBar is a representation of GtkHeaderBar
type HeaderBar struct {
	Container
}

// native returns a pointer to the underlying GtkHeaderBar.
func (v *HeaderBar) native() *C.GtkHeaderBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkHeaderBar(p)
}

func wrapHeaderBar(obj *Object) *HeaderBar {
	if obj == nil {
		return nil
	}

	return &HeaderBar{Container{Widget{InitiallyUnowned{obj}}}}
}

// NewHeaderBar is a wrapper around gtk_header_bar_new().
func NewHeaderBar() *HeaderBar {
	c := C.gtk_header_bar_new()
	if c == nil {
		return nil
	}
	return wrapHeaderBar(ToGoObject(unsafe.Pointer(c)))
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (v *HeaderBar) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (v *HeaderBar) GetTitle() string {
	cstr := C.gtk_header_bar_get_title(v.native())
	return C.GoString((*C.char)(cstr))
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (v *HeaderBar) SetSubtitle(subtitle string) {
	cstr := C.CString(subtitle)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_subtitle(v.native(), (*C.gchar)(cstr))
}

// GetSubtitle is a wrapper around gtk_header_bar_get_subtitle().
func (v *HeaderBar) GetSubtitle() string {
	cstr := C.gtk_header_bar_get_subtitle(v.native())
	return C.GoString((*C.char)(cstr))
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (v *HeaderBar) SetCustomTitle(titleWidget IWidget) {
	C.gtk_header_bar_set_custom_title(v.native(), GtkWidget(titleWidget))
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (v *HeaderBar) PackStart(child IWidget) {
	C.gtk_header_bar_pack_start(v.native(), GtkWidget(child))
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (v *HeaderBar) PackEnd(child IWidget) {
	C.gtk_header_bar_pack_end(v.native(), GtkWidget(child))
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (v *HeaderBar) SetShowCloseButton(setting bool) {
	C.gtk_header_bar_set_show_close_button(v.native(), CBool(setting))
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (v *HeaderBar) GetShowCloseButton() bool {
	c := C.gtk_header_bar_get_show_close_button(v.native())
	return GoBool(c)
}
