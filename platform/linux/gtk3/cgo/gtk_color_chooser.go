package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// ColorChooserDialog is a representation of GTK's GtkColorChooserDialog.
type ColorChooserDialog struct {
	Dialog
}

func (v *ColorChooserDialog) native() *C.GtkColorChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkColorChooserDialog(unsafe.Pointer(v.GObject))
}

func wrapColorChooserDialog(obj *Object) *ColorChooserDialog {
	return &ColorChooserDialog{Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// NewColorChooserDialog is a wrapper around gtk_color_chooser_dialog_new().
func NewColorChooserDialog(title string, parent IWindow) *ColorChooserDialog {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow
	if parent != nil {
		w = (*C.GtkWindow)(unsafe.Pointer(parent.Instance()))
	}
	c := C.gtk_color_chooser_dialog_new((*C.gchar)(cstr), w)
	if c == nil {
		return nil
	}
	return wrapColorChooserDialog(ToGoObject(unsafe.Pointer(c)))
}

// GetUseAlpha is a wrapper around gtk_color_chooser_get_use_alpha().
func (v *ColorChooserDialog) GetUseAlpha() bool {
	return GoBool(C.gtk_color_chooser_get_use_alpha(C.toGtkColorChooser(unsafe.Pointer(v.GObject))))
}

// SetUseAlpha is a wrapper around gtk_color_chooser_set_use_alpha().
func (v *ColorChooserDialog) SetUseAlpha(useAlpha bool) {
	C.gtk_color_chooser_set_use_alpha(C.toGtkColorChooser(unsafe.Pointer(v.GObject)), CBool(useAlpha))
}

func (v *ColorChooserDialog) GetRGBA() GdkRGBA {
	var c C.GdkRGBA
	C.gtk_color_chooser_get_rgba(C.toGtkColorChooser(unsafe.Pointer(v.GObject)), &c)
	return GdkRGBA{Red: float64(c.red), Green: float64(c.green), Blue: float64(c.blue), Alpha: float64(c.alpha)}
}

// SetRGBA is a wrapper around gtk_color_chooser_set_rgba().
func (v *ColorChooserDialog) SetRGBA(rgba GdkRGBA) {
	c := C.GdkRGBA{
		red:   C.gdouble(rgba.Red),
		green: C.gdouble(rgba.Green),
		blue:  C.gdouble(rgba.Blue),
		alpha: C.gdouble(rgba.Alpha),
	}
	C.gtk_color_chooser_set_rgba(C.toGtkColorChooser(unsafe.Pointer(v.GObject)), &c)
}

// AddPalette is a wrapper around gtk_color_chooser_add_palette().
func (v *ColorChooserDialog) AddPalette(orientation Orientation, colorsPerLine int, colors []GdkRGBA) {
	n := len(colors)
	if n == 0 {
		return
	}
	cColors := make([]C.GdkRGBA, n)
	for i, col := range colors {
		cColors[i] = C.GdkRGBA{
			red:   C.gdouble(col.Red),
			green: C.gdouble(col.Green),
			blue:  C.gdouble(col.Blue),
			alpha: C.gdouble(col.Alpha),
		}
	}
	C.gtk_color_chooser_add_palette(C.toGtkColorChooser(unsafe.Pointer(v.GObject)),
		C.GtkOrientation(orientation), C.gint(colorsPerLine), C.gint(n), &cColors[0])
}
