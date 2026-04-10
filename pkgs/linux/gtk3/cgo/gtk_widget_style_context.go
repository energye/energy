package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/linux/types"
	"unsafe"
)

// StyleContext is a representation of GTK's GtkStyleContext.
type StyleContext struct {
	*Object
}

// native returns a pointer to the underlying GtkStyleContext.
func (v *StyleContext) native() *C.GtkStyleContext {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStyleContext(p)
}

func wrapStyleContext(obj *Object) *StyleContext {
	if obj == nil {
		return nil
	}

	return &StyleContext{obj}
}

func (v *StyleContext) AddClass(class_name string) {
	cstr := C.CString(class_name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_style_context_add_class(v.native(), (*C.gchar)(cstr))
}

func (v *StyleContext) RemoveClass(class_name string) {
	cstr := C.CString(class_name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_style_context_remove_class(v.native(), (*C.gchar)(cstr))
}

func fromNativeStyleContext(c *C.GtkStyleContext) *StyleContext {
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapStyleContext(obj)
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (v *Widget) GetStyleContext() IStyleContext {
	return fromNativeStyleContext(C.gtk_widget_get_style_context(v.native()))
}

// GetParent is a wrapper around gtk_style_context_get_parent().
func (v *StyleContext) GetParent() *StyleContext {
	return fromNativeStyleContext(C.gtk_style_context_get_parent(v.native()))
}

// GetState is a wrapper around gtk_style_context_get_state().
func (v *StyleContext) GetState() StateFlags {
	return StateFlags(C.gtk_style_context_get_state(v.native()))
}

// GetColor is a wrapper around gtk_style_context_get_color().
func (v *StyleContext) GetColor(state StateFlags) *RGBA {
	gdkColor := NewRGBA()
	C.gtk_style_context_get_color(v.native(), C.GtkStateFlags(state), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor
}

// LookupColor is a wrapper around gtk_style_context_lookup_color().
func (v *StyleContext) LookupColor(colorName string) (*RGBA, bool) {
	cstr := (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(cstr))
	gdkColor := NewRGBA()
	ret := C.gtk_style_context_lookup_color(v.native(), cstr, (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor, GoBool(ret)
}

// StyleContextResetWidgets is a wrapper around gtk_style_context_reset_widgets().
func StyleContextResetWidgets(v *Screen) {
	C.gtk_style_context_reset_widgets((*C.GdkScreen)(unsafe.Pointer(v.Native())))
}

// Restore is a wrapper around gtk_style_context_restore().
func (v *StyleContext) Restore() {
	C.gtk_style_context_restore(v.native())
}

// Save is a wrapper around gtk_style_context_save().
func (v *StyleContext) Save() {
	C.gtk_style_context_save(v.native())
}

// SetParent is a wrapper around gtk_style_context_set_parent().
func (v *StyleContext) SetParent(p *StyleContext) {
	C.gtk_style_context_set_parent(v.native(), p.native())
}

// HasClass is a wrapper around gtk_style_context_has_class().
func (v *StyleContext) HasClass(className string) bool {
	cstr := C.CString(className)
	defer C.free(unsafe.Pointer(cstr))

	return GoBool(C.gtk_style_context_has_class(v.native(), (*C.gchar)(cstr)))
}

// SetScreen is a wrapper around gtk_style_context_set_screen().
func (v *StyleContext) SetScreen(s *Screen) {
	C.gtk_style_context_set_screen(v.native(), (*C.GdkScreen)(unsafe.Pointer(s.Native())))
}

// SetState is a wrapper around gtk_style_context_set_state().
func (v *StyleContext) SetState(state StateFlags) {
	C.gtk_style_context_set_state(v.native(), C.GtkStateFlags(state))
}

type iStyleProvider interface {
	toStyleProvider() *C.GtkStyleProvider
}

// AddProvider is a wrapper around gtk_style_context_add_provider().
func (v *StyleContext) AddProvider(provider IStyleProvider, prio uint) {
	C.gtk_style_context_add_provider(v.native(), provider.(iStyleProvider).toStyleProvider(), C.guint(prio))
}

// AddProviderForScreen is a wrapper around gtk_style_context_add_provider_for_screen().
func AddProviderForScreen(s *Screen, provider IStyleProvider, prio uint) {
	C.gtk_style_context_add_provider_for_screen((*C.GdkScreen)(unsafe.Pointer(s.Native())), provider.(iStyleProvider).toStyleProvider(), C.guint(prio))
}

// RemoveProvider is a wrapper around gtk_style_context_remove_provider().
func (v *StyleContext) RemoveProvider(provider IStyleProvider) {
	C.gtk_style_context_remove_provider(v.native(), provider.(iStyleProvider).toStyleProvider())
}

// RemoveProviderForScreen is a wrapper around gtk_style_context_remove_provider_for_screen().
func RemoveProviderForScreen(s *Screen, provider IStyleProvider) {
	C.gtk_style_context_remove_provider_for_screen((*C.GdkScreen)(unsafe.Pointer(s.Native())), provider.(iStyleProvider).toStyleProvider())
}
