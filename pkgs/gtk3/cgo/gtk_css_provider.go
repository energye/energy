package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"unsafe"
)

// CssProvider is a representation of GTK's GtkCssProvider.
type CssProvider struct {
	*Object
}

func (v *CssProvider) toStyleProvider() *C.GtkStyleProvider {
	if v == nil {
		return nil
	}
	return C.toGtkStyleProvider(unsafe.Pointer(v.native()))
}

// native returns a pointer to the underlying GtkCssProvider.
func (v *CssProvider) native() *C.GtkCssProvider {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCssProvider(p)
}

func wrapCssProvider(obj *Object) *CssProvider {
	if obj == nil {
		return nil
	}

	return &CssProvider{obj}
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func NewCssProvider() *CssProvider {
	c := C.gtk_css_provider_new()
	if c == nil {
		return nil
	}
	return wrapCssProvider(ToGoObject(unsafe.Pointer(c)))
}

// LoadFromPath is a wrapper around gtk_css_provider_load_from_path().
func (v *CssProvider) LoadFromPath(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_path(v.native(), (*C.gchar)(cpath), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(GoString(gerr.message))
	}
	return nil
}

// LoadFromData is a wrapper around gtk_css_provider_load_from_data().
func (v *CssProvider) LoadFromData(data string) error {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_data(v.native(), (*C.gchar)(unsafe.Pointer(cdata)), C.gssize(len(data)), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(GoString(gerr.message))
	}
	return nil
}

// ToString is a wrapper around gtk_css_provider_to_string().
func (v *CssProvider) ToString() (string, error) {
	c := C.gtk_css_provider_to_string(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString(c), nil
}

// CssProviderGetNamed is a wrapper around gtk_css_provider_get_named().
func CssProviderGetNamed(name string, variant string) *CssProvider {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvariant := C.CString(variant)
	defer C.free(unsafe.Pointer(cvariant))
	c := C.gtk_css_provider_get_named((*C.gchar)(cname), (*C.gchar)(cvariant))
	if c == nil {
		return nil
	}
	return wrapCssProvider(ToGoObject(unsafe.Pointer(c)))
}
