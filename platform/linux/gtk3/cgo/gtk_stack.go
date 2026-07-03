package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "gtk_header_bar.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// ─────────────────────────────────────────────
// Stack
// ─────────────────────────────────────────────

type Stack struct {
	Container
}

func (v *Stack) native() *C.GtkStack {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkStack(unsafe.Pointer(v.GObject))
}

func wrapStack(obj *Object) *Stack {
	return &Stack{Container{Widget{InitiallyUnowned{obj}}}}
}

func AsStack(ptr unsafe.Pointer) *Stack {
	return wrapStack(ToGoObject(ptr))
}

func NewStack() *Stack {
	c := C.gtk_stack_new()
	if c == nil {
		return nil
	}
	return wrapStack(ToGoObject(unsafe.Pointer(c)))
}

func (v *Stack) AddNamed(child IWidget, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_add_named(v.native(), GtkWidget(child), (*C.gchar)(cstr))
}

func (v *Stack) AddTitled(child IWidget, name, title string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_stack_add_titled(v.native(), GtkWidget(child), (*C.gchar)(cName), (*C.gchar)(cTitle))
}

func (v *Stack) SetVisibleChild(child IWidget) {
	C.gtk_stack_set_visible_child(v.native(), GtkWidget(child))
}

func (v *Stack) GetVisibleChild() IWidget {
	c := C.gtk_stack_get_visible_child(v.native())
	if c == nil {
		return nil
	}
	return AsWidget(unsafe.Pointer(c))
}

func (v *Stack) SetVisibleChildName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_name(v.native(), (*C.gchar)(cstr))
}

func (v *Stack) GetVisibleChildName() string {
	c := C.gtk_stack_get_visible_child_name(v.native())
	return C.GoString((*C.char)(c))
}

func (v *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_full(v.native(), (*C.gchar)(cstr), C.GtkStackTransitionType(transition))
}

func (v *Stack) SetHomogeneous(homogeneous bool) {
	C.gtk_stack_set_homogeneous(v.native(), CBool(homogeneous))
}

func (v *Stack) GetHomogeneous() bool {
	return GoBool(C.gtk_stack_get_homogeneous(v.native()))
}

func (v *Stack) SetTransitionDuration(duration uint) {
	C.gtk_stack_set_transition_duration(v.native(), C.guint(duration))
}

func (v *Stack) GetTransitionDuration() uint {
	return uint(C.gtk_stack_get_transition_duration(v.native()))
}

func (v *Stack) SetTransitionType(transition StackTransitionType) {
	C.gtk_stack_set_transition_type(v.native(), C.GtkStackTransitionType(transition))
}

func (v *Stack) GetTransitionType() StackTransitionType {
	return StackTransitionType(C.gtk_stack_get_transition_type(v.native()))
}

// ─────────────────────────────────────────────
// StackSwitcher
// ─────────────────────────────────────────────

type StackSwitcher struct {
	Box
}

func (v *StackSwitcher) native() *C.GtkStackSwitcher {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkStackSwitcher(unsafe.Pointer(v.GObject))
}

func wrapStackSwitcher(obj *Object) *StackSwitcher {
	return &StackSwitcher{Box{Container{Widget{InitiallyUnowned{obj}}}}}
}

func AsStackSwitcher(ptr unsafe.Pointer) *StackSwitcher {
	return wrapStackSwitcher(ToGoObject(ptr))
}

func NewStackSwitcher() *StackSwitcher {
	c := C.gtk_stack_switcher_new()
	if c == nil {
		return nil
	}
	return wrapStackSwitcher(ToGoObject(unsafe.Pointer(c)))
}

func (v *StackSwitcher) SetStack(stack IStack) {
	s := stack.(*Stack)
	C.gtk_stack_switcher_set_stack(v.native(), s.native())
}

func (v *StackSwitcher) GetStack() IStack {
	c := C.gtk_stack_switcher_get_stack(v.native())
	if c == nil {
		return nil
	}
	return wrapStack(ToGoObject(unsafe.Pointer(c)))
}
