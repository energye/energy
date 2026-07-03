package cgo

// #include <gtk/gtk.h>
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Scale struct {
	Range
}

func (v *Scale) native() *C.GtkScale {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkScale)(unsafe.Pointer(v.GObject))
}

func wrapScale(obj *Object) *Scale {
	return &Scale{Range{Widget{InitiallyUnowned{obj}}}}
}

func NewHScale(adjustment IAdjustment) *Scale {
	var adj *C.GtkAdjustment
	if adjustment != nil {
		adj = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Instance()))
	}
	c := C.gtk_scale_new(C.GtkOrientation(ORIENTATION_HORIZONTAL), adj)
	if c == nil {
		return nil
	}
	return wrapScale(ToGoObject(unsafe.Pointer(c)))
}

func NewVScale(adjustment IAdjustment) *Scale {
	var adj *C.GtkAdjustment
	if adjustment != nil {
		adj = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Instance()))
	}
	c := C.gtk_scale_new(C.GtkOrientation(ORIENTATION_VERTICAL), adj)
	if c == nil {
		return nil
	}
	return wrapScale(ToGoObject(unsafe.Pointer(c)))
}

func (v *Scale) SetDigits(digits int) {
	C.gtk_scale_set_digits(v.native(), C.gint(digits))
}

func (v *Scale) GetDigits() int {
	return int(C.gtk_scale_get_digits(v.native()))
}

func (v *Scale) SetDrawValue(drawValue bool) {
	C.gtk_scale_set_draw_value(v.native(), CBool(drawValue))
}

func (v *Scale) GetDrawValue() bool {
	return GoBool(C.gtk_scale_get_draw_value(v.native()))
}

func (v *Scale) SetValuePos(pos PositionType) {
	C.gtk_scale_set_value_pos(v.native(), C.GtkPositionType(pos))
}

func (v *Scale) GetValuePos() PositionType {
	return PositionType(C.gtk_scale_get_value_pos(v.native()))
}
