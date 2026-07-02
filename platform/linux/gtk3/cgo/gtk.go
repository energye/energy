package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

func Init(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		for i, arg := range *args {
			cstr := C.CString(arg)
			C.set_string(argv, C.int(i), (*C.gchar)(cstr))
		}

		C.gtk_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))

		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			cstr := C.get_string(argv, C.int(i))
			unhandled[i] = GoString(cstr)
			C.free(unsafe.Pointer(cstr))
		}
		*args = unhandled
	} else {
		C.gtk_init(nil, nil)
	}
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

func marshalModifierType(p uintptr) (any, error) {
	c := C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))
	return ModifierType(c), nil
}

// castWidget takes a native GtkWidget and casts it to the appropriate Go struct.
func castWidget(c *C.GtkWidget) IWidget {
	wdt := new(Widget)
	wdt.Object = ToGoObject(unsafe.Pointer(c))
	return wdt
}

func marshalAlign(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Align(c), nil
}

func marshalStateFlags(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return StateFlags(c), nil
}

func marshalIconSize(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return IconSize(c), nil
}

func marshalImageType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ImageType(c), nil
}

func marshalEntryIconPosition(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EntryIconPosition(c), nil
}

func marshalOrientation(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Orientation(c), nil
}

func marshalPackType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PackType(c), nil
}

func marshalJustification(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Justification(c), nil
}

func marshalSizeRequestMode(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SizeRequestMode(c), nil
}

func MainDoEvent(event *Event) {
	C.gtk_main_do_event(event.native())
}
