package gtk3

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// Requisition is a representation of GTK's GtkRequisition
type Requisition struct {
	requisition *C.GtkRequisition
	Width,
	Height int
}

func (v *Requisition) native() *C.GtkRequisition {
	if v == nil {
		return nil
	}
	v.requisition.width = C.int(v.Width)
	v.requisition.height = C.int(v.Height)
	return v.requisition
}

// Native returns a pointer to the underlying GtkRequisition.
func (v *Requisition) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func wrapRequisition(requisition *C.GtkRequisition) *Requisition {
	if requisition == nil {
		return nil
	}
	return &Requisition{requisition, int(requisition.width), int(requisition.height)}
}

// requisitionFromNative that handle finalizer.
func requisitionFromNative(requisitionNative *C.GtkRequisition) *Requisition {
	requisition := wrapRequisition(requisitionNative)
	if requisition == nil {
		return nil
	}
	return requisition
}

// NewRequisition is a wrapper around gtk_requisition_new().
func NewRequisition() *Requisition {
	c := C.gtk_requisition_new()
	if c == nil {
		return nil
	}
	return requisitionFromNative(c)
}

// Free is a wrapper around gtk_requisition_free().
func (v *Requisition) Free() {
	C.gtk_requisition_free(v.native())
}

// Copy is a wrapper around gtk_requisition_copy().
func (v *Requisition) Copy() *Requisition {
	c := C.gtk_requisition_copy(v.native())
	if c == nil {
		return nil
	}
	return requisitionFromNative(c)
}
