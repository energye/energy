package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

func ToScrolledWindow(p unsafe.Pointer) *ScrolledWindow {
	return &ScrolledWindow{Bin{Container{Widget{InitiallyUnowned{ToGoObject(p)}}}}}
}

// native returns a pointer to the underlying GtkScrolledWindow.
func (v *ScrolledWindow) native() *C.GtkScrolledWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkScrolledWindow(p)
}

func wrapScrolledWindow(obj *Object) *ScrolledWindow {
	if obj == nil {
		return nil
	}
	return &ScrolledWindow{Bin{Container{Widget{InitiallyUnowned{obj}}}}}
}

// NewScrolledWindow is a wrapper around gtk_scrolled_window_new().
func NewScrolledWindow(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	c := C.gtk_scrolled_window_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapScrolledWindow(obj)
}

// GetPolicy() is a wrapper around gtk_scrolled_window_get_policy().
func (v *ScrolledWindow) GetPolicy() (hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	var hScrPol, vScrPol C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(v.native(), &hScrPol, &vScrPol)
	hScrollbarPolicy, vScrollbarPolicy = PolicyType(hScrPol), PolicyType(vScrPol)
	return
}

// SetPolicy() is a wrapper around gtk_scrolled_window_set_policy().
func (v *ScrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	C.gtk_scrolled_window_set_policy(v.native(),
		C.GtkPolicyType(hScrollbarPolicy),
		C.GtkPolicyType(vScrollbarPolicy))
}

// GetHAdjustment() is a wrapper around gtk_scrolled_window_get_hadjustment().
func (v *ScrolledWindow) GetHAdjustment() *Adjustment {
	c := C.gtk_scrolled_window_get_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetHAdjustment is a wrapper around gtk_scrolled_window_set_hadjustment().
func (v *ScrolledWindow) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(v.native(), adjustment.native())
}

// GetVAdjustment() is a wrapper around gtk_scrolled_window_get_vadjustment().
func (v *ScrolledWindow) GetVAdjustment() *Adjustment {
	c := C.gtk_scrolled_window_get_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetVAdjustment is a wrapper around gtk_scrolled_window_set_vadjustment().
func (v *ScrolledWindow) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(v.native(), adjustment.native())
}

// GetHScrollbar is a wrapper around gtk_scrolled_window_get_hscrollbar().
func (v *ScrolledWindow) GetHScrollbar() *Scrollbar {
	c := C.gtk_scrolled_window_get_hscrollbar(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapScrollbar(obj)
}

// GetVScrollbar is a wrapper around gtk_scrolled_window_get_vscrollbar().
func (v *ScrolledWindow) GetVScrollbar() *Scrollbar {
	c := C.gtk_scrolled_window_get_vscrollbar(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapScrollbar(obj)
}

// GetPlacement is a wrapper around gtk_scrolled_window_get_placement().
func (v *ScrolledWindow) GetPlacement() CornerType {
	c := C.gtk_scrolled_window_get_placement(v.native())
	return CornerType(c)
}

// SetPlacement is a wrapper around gtk_scrolled_window_set_placement().
func (v *ScrolledWindow) SetPlacement(windowPlacement CornerType) {
	C.gtk_scrolled_window_set_placement(v.native(), C.GtkCornerType(windowPlacement))
}

// UnsetPlacement is a wrapper around gtk_scrolled_window_unset_placement().
func (v *ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(v.native())
}

// GetShadowType is a wrapper around gtk_scrolled_window_get_shadow_type().
func (v *ScrolledWindow) GetShadowType() ShadowType {
	c := C.gtk_scrolled_window_get_shadow_type(v.native())
	return ShadowType(c)
}

// SetShadowType is a wrapper around gtk_scrolled_window_set_shadow_type().
func (v *ScrolledWindow) SetShadowType(t ShadowType) {
	C.gtk_scrolled_window_set_shadow_type(v.native(), C.GtkShadowType(t))
}

// GetKineticScrolling is a wrapper around gtk_scrolled_window_get_kinetic_scrolling().
func (v *ScrolledWindow) GetKineticScrolling() bool {
	c := C.gtk_scrolled_window_get_kinetic_scrolling(v.native())
	return GoBool(c)
}

// SetKineticScrolling is a wrapper around gtk_scrolled_window_set_kinetic_scrolling().
func (v *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	C.gtk_scrolled_window_set_kinetic_scrolling(v.native(), CBool(kineticScrolling))
}

// GetCaptureButtonPress is a wrapper around gtk_scrolled_window_get_capture_button_press().
func (v *ScrolledWindow) GetCaptureButtonPress() bool {
	c := C.gtk_scrolled_window_get_capture_button_press(v.native())
	return GoBool(c)
}

// SetCaptureButtonPress is a wrapper around gtk_scrolled_window_set_capture_button_press().
func (v *ScrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	C.gtk_scrolled_window_set_capture_button_press(v.native(), CBool(captureButtonPress))
}

// GetMinContentWidth is a wrapper around gtk_scrolled_window_get_min_content_width().
func (v *ScrolledWindow) GetMinContentWidth() int {
	c := C.gtk_scrolled_window_get_min_content_width(v.native())
	return int(c)
}

// SetMinContentWidth is a wrapper around gtk_scrolled_window_set_min_content_width().
func (v *ScrolledWindow) SetMinContentWidth(width int) {
	C.gtk_scrolled_window_set_min_content_width(v.native(), C.gint(width))
}

// GetMinContentHeight is a wrapper around gtk_scrolled_window_get_min_content_height().
func (v *ScrolledWindow) GetMinContentHeight() int {
	c := C.gtk_scrolled_window_get_min_content_height(v.native())
	return int(c)
}

// SetMinContentHeight is a wrapper around gtk_scrolled_window_set_min_content_height().
func (v *ScrolledWindow) SetMinContentHeight(width int) {
	C.gtk_scrolled_window_set_min_content_height(v.native(), C.gint(width))
}
