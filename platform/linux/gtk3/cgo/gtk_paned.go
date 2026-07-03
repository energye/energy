package cgo

// #include <gtk/gtk.h>
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Paned struct {
	Container
}

func (v *Paned) native() *C.GtkPaned {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkPaned)(unsafe.Pointer(v.GObject))
}

func wrapPaned(obj *Object) *Paned {
	return &Paned{Container{Widget{InitiallyUnowned{obj}}}}
}

func NewPaned(orientation Orientation) *Paned {
	c := C.gtk_paned_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil
	}
	return wrapPaned(ToGoObject(unsafe.Pointer(c)))
}

func (v *Paned) Add1(child IWidget) {
	C.gtk_paned_add1(v.native(), GtkWidget(child))
}

func (v *Paned) Add2(child IWidget) {
	C.gtk_paned_add2(v.native(), GtkWidget(child))
}

func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(v.native(), C.gint(position))
}

func (v *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(v.native()))
}

func (v *Paned) SetWideHandle(wide bool) {
	C.gtk_paned_set_wide_handle(v.native(), CBool(wide))
}

func (v *Paned) GetWideHandle() bool {
	return GoBool(C.gtk_paned_get_wide_handle(v.native()))
}
