package gtk3

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include <gtk/gtk.h>
// #include "gdk.go.h"
import "C"
import "unsafe"

// DragContext is a representation of GDK's GdkDragContext.
type DragContext struct {
	*Object
}

func ToDragContext(p unsafe.Pointer) *DragContext {
	obj := ToGoObject(p)
	return &DragContext{obj}
}

// native returns a pointer to the underlying GdkDragContext.
func (v *DragContext) native() *C.GdkDragContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDragContext(p)
}

// Native returns a pointer to the underlying GdkDragContext.
func (v *DragContext) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *DragContext) ListTargets() *List {
	clist := C.gdk_drag_context_list_targets(v.native())
	if clist == nil {
		return nil
	}
	return WrapList(uintptr(unsafe.Pointer(clist)))
}

func (v *DragContext) Finish(success bool, del bool, time uint) {
	C.gtk_drag_finish(v.native(), CBool(success), CBool(del), C.uint(time))
}

func (v *DragContext) DragStatus(actions DragAction, time uint) {
	C.gdk_drag_status(v.native(), C.GdkDragAction(actions), C.uint(time))
}
