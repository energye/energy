package cgo

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <gio/gio.h>
#include <gtk/gtk.h>
#include "gtk.go.h"
*/
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Button is a representation of GTK's GtkButton.
type Button struct {
	Bin
}

// native() returns a pointer to the underlying GtkButton.
func (v *Button) native() *C.GtkButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkButton(p)
}

func wrapButton(obj *Object) *Button {
	if obj == nil {
		return nil
	}

	return &Button{Bin{Container{Widget{InitiallyUnowned{obj}}}}}
}

// ButtonNew is a wrapper around gtk_button_new().
func NewButton() *Button {
	c := C.gtk_button_new()
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapButton(obj)
}

// ButtonNewWithLabel is a wrapper around gtk_button_new_with_label().
func NewButtonWithLabel(label string) *Button {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapButton(obj)
}

// Clicked is a wrapper around gtk_button_clicked().
func (v *Button) Clicked() {
	C.gtk_button_clicked(v.native())
}

// SetRelief is a wrapper around gtk_button_set_relief().
func (v *Button) SetRelief(newStyle ReliefStyle) {
	C.gtk_button_set_relief(v.native(), C.GtkReliefStyle(newStyle))
}

// GetRelief is a wrapper around gtk_button_get_relief().
func (v *Button) GetRelief() ReliefStyle {
	c := C.gtk_button_get_relief(v.native())
	return ReliefStyle(c)
}

// SetLabel is a wrapper around gtk_button_set_label().
func (v *Button) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_button_get_label().
func (v *Button) GetLabel() (string, error) {
	c := C.gtk_button_get_label(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return GoString(c), nil
}

// SetUseUnderline is a wrapper around gtk_button_set_use_underline().
func (v *Button) SetUseUnderline(useUnderline bool) {
	C.gtk_button_set_use_underline(v.native(), CBool(useUnderline))
}

// GetUseUnderline is a wrapper around gtk_button_get_use_underline().
func (v *Button) GetUseUnderline() bool {
	c := C.gtk_button_get_use_underline(v.native())
	return GoBool(c)
}

// SetImage is a wrapper around gtk_button_set_image().
func (v *Button) SetImage(image IWidget) {
	C.gtk_button_set_image(v.native(), GtkWidget(image))
}

// GetImage is a wrapper around gtk_button_get_image().
func (v *Button) GetImage() IWidget {
	c := C.gtk_button_get_image(v.native())
	if c == nil {
		return nil
	}
	return castWidget(c)
}

// SetImagePosition is a wrapper around gtk_button_set_image_position().
func (v *Button) SetImagePosition(position PositionType) {
	C.gtk_button_set_image_position(v.native(), C.GtkPositionType(position))
}

// GetImagePosition is a wrapper around gtk_button_get_image_position().
func (v *Button) GetImagePosition() PositionType {
	c := C.gtk_button_get_image_position(v.native())
	return PositionType(c)
}

// SetAlwaysShowImage is a wrapper around gtk_button_set_always_show_image().
func (v *Button) SetAlwaysShowImage(alwaysShow bool) {
	C.gtk_button_set_always_show_image(v.native(), CBool(alwaysShow))
}

// GetAlwaysShowImage is a wrapper around gtk_button_get_always_show_image().
func (v *Button) GetAlwaysShowImage() bool {
	c := C.gtk_button_get_always_show_image(v.native())
	return GoBool(c)
}

// GetEventWindow is a wrapper around gtk_button_get_event_window().
func (v *Button) GetEventWindow() (*Window, error) {
	c := C.gtk_button_get_event_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	w := &Window{}
	w.Object = ToGoObject(unsafe.Pointer(c))
	return w, nil
}

func (v *Button) SetOnLeave(fn TLeaveEnterNotifyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(v.Instance()), EsnLeaveNotifyEvent, "c_trampoline_3_gboolean",
		fn, nil)
	return signalHandlerID
}

func (v *Button) SetOnEnter(fn TLeaveEnterNotifyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(v.Instance()), EsnEnterNotifyEvent, "c_trampoline_3_gboolean",
		fn, nil)
	return signalHandlerID
}

func (v *Button) SetOnClick(fn TNotifyEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(unsafe.Pointer(v.Instance()), EsnClicked, "c_trampoline_2_void",
		fn, nil)
	return signalHandlerID
}
