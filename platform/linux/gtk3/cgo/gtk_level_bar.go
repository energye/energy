package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type LevelBar struct {
	Widget
}

func (v *LevelBar) native() *C.GtkLevelBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkLevelBar(unsafe.Pointer(v.GObject))
}

func wrapLevelBar(obj *Object) *LevelBar {
	return &LevelBar{Widget{InitiallyUnowned{obj}}}
}

func NewLevelBar() *LevelBar {
	c := C.gtk_level_bar_new()
	if c == nil {
		return nil
	}
	return wrapLevelBar(ToGoObject(unsafe.Pointer(c)))
}

func (v *LevelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(v.native(), C.gdouble(value))
}

func (v *LevelBar) GetValue() float64 {
	return float64(C.gtk_level_bar_get_value(v.native()))
}

func (v *LevelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(v.native(), C.gdouble(value))
}

func (v *LevelBar) GetMinValue() float64 {
	return float64(C.gtk_level_bar_get_min_value(v.native()))
}

func (v *LevelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(v.native(), C.gdouble(value))
}

func (v *LevelBar) GetMaxValue() float64 {
	return float64(C.gtk_level_bar_get_max_value(v.native()))
}

func (v *LevelBar) SetMode(mode LevelBarMode) {
	C.gtk_level_bar_set_mode(v.native(), C.GtkLevelBarMode(mode))
}

func (v *LevelBar) GetMode() LevelBarMode {
	return LevelBarMode(C.gtk_level_bar_get_mode(v.native()))
}

func (m *LevelBar) SetOnOffsetChanged(fn TOffsetChangedEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnOffsetChanged, callback.C_trampoline_3_void, fn, 0)
}
