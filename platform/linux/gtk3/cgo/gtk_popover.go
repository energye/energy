package cgo

// #include <gtk/gtk.h>
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Popover struct {
	Bin
}

func (v *Popover) native() *C.GtkPopover {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkPopover)(unsafe.Pointer(v.GObject))
}

func wrapPopover(obj *Object) *Popover {
	return &Popover{Bin{Container{Widget{InitiallyUnowned{obj}}}}}
}

func NewPopover() *Popover {
	c := C.gtk_popover_new(nil)
	if c == nil {
		return nil
	}
	return wrapPopover(ToGoObject(unsafe.Pointer(c)))
}

func (v *Popover) SetRelativeTo(widget IWidget) {
	C.gtk_popover_set_relative_to(v.native(), GtkWidget(widget))
}

func (v *Popover) GetRelativeTo() IWidget {
	c := C.gtk_popover_get_relative_to(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(ToGoObject(unsafe.Pointer(c)))
}

func (v *Popover) SetPosition(position PositionType) {
	C.gtk_popover_set_position(v.native(), C.GtkPositionType(position))
}

func (v *Popover) GetPosition() PositionType {
	return PositionType(C.gtk_popover_get_position(v.native()))
}

func (v *Popover) SetModal(modal bool) {
	C.gtk_popover_set_modal(v.native(), CBool(modal))
}

func (v *Popover) GetModal() bool {
	return GoBool(C.gtk_popover_get_modal(v.native()))
}

func (v *Popover) Popdown() {
	C.gtk_popover_popdown(v.native())
}

func (v *Popover) Popup() {
	C.gtk_popover_popup(v.native())
}
