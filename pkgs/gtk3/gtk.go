package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
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

// WindowType is a representation of GTK's GtkWindowType.
type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = C.GTK_WINDOW_TOPLEVEL
	WINDOW_POPUP    WindowType = C.GTK_WINDOW_POPUP
)

type Gravity int

const (
	GDK_GRAVITY_NORTH_WEST = C.GDK_GRAVITY_NORTH_WEST
	GDK_GRAVITY_NORTH      = C.GDK_GRAVITY_NORTH
	GDK_GRAVITY_NORTH_EAST = C.GDK_GRAVITY_NORTH_EAST
	GDK_GRAVITY_WEST       = C.GDK_GRAVITY_WEST
	GDK_GRAVITY_CENTER     = C.GDK_GRAVITY_CENTER
	GDK_GRAVITY_EAST       = C.GDK_GRAVITY_EAST
	GDK_GRAVITY_SOUTH_WEST = C.GDK_GRAVITY_SOUTH_WEST
	GDK_GRAVITY_SOUTH      = C.GDK_GRAVITY_SOUTH
	GDK_GRAVITY_SOUTH_EAST = C.GDK_GRAVITY_SOUTH_EAST
	GDK_GRAVITY_STATIC     = C.GDK_GRAVITY_STATIC
)

// WindowTypeHint is a representation of GDK's GdkWindowTypeHint
type WindowTypeHint int

const (
	WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_NORMAL
	WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DIALOG
	WINDOW_TYPE_HINT_MENU          WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_MENU
	WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_TOOLBAR
	WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_SPLASHSCREEN
	WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_UTILITY
	WINDOW_TYPE_HINT_DOCK          WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DOCK
	WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DESKTOP
	WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU
	WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_POPUP_MENU
	WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_TOOLTIP
	WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_NOTIFICATION
	WINDOW_TYPE_HINT_COMBO         WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_COMBO
	WINDOW_TYPE_HINT_DND           WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DND
)

// ModifierType is a representation of GDK's GdkModifierType.
type ModifierType uint

const (
	SHIFT_MASK    ModifierType = C.GDK_SHIFT_MASK
	LOCK_MASK                  = C.GDK_LOCK_MASK
	CONTROL_MASK               = C.GDK_CONTROL_MASK
	MOD1_MASK                  = C.GDK_MOD1_MASK
	MOD2_MASK                  = C.GDK_MOD2_MASK
	MOD3_MASK                  = C.GDK_MOD3_MASK
	MOD4_MASK                  = C.GDK_MOD4_MASK
	MOD5_MASK                  = C.GDK_MOD5_MASK
	BUTTON1_MASK               = C.GDK_BUTTON1_MASK
	BUTTON2_MASK               = C.GDK_BUTTON2_MASK
	BUTTON3_MASK               = C.GDK_BUTTON3_MASK
	BUTTON4_MASK               = C.GDK_BUTTON4_MASK
	BUTTON5_MASK               = C.GDK_BUTTON5_MASK
	SUPER_MASK                 = C.GDK_SUPER_MASK
	HYPER_MASK                 = C.GDK_HYPER_MASK
	META_MASK                  = C.GDK_META_MASK
	RELEASE_MASK               = C.GDK_RELEASE_MASK
	MODIFIER_MASK              = C.GDK_MODIFIER_MASK
)

// ReliefStyle is a representation of GTK's GtkReliefStyle.
type ReliefStyle int

const (
	RELIEF_NORMAL ReliefStyle = C.GTK_RELIEF_NORMAL
	RELIEF_HALF   ReliefStyle = C.GTK_RELIEF_HALF
	RELIEF_NONE   ReliefStyle = C.GTK_RELIEF_NONE
)

// PositionType is a representation of GTK's GtkPositionType.
type PositionType int

const (
	POS_LEFT   PositionType = C.GTK_POS_LEFT
	POS_RIGHT  PositionType = C.GTK_POS_RIGHT
	POS_TOP    PositionType = C.GTK_POS_TOP
	POS_BOTTOM PositionType = C.GTK_POS_BOTTOM
)

func marshalModifierType(p uintptr) (interface{}, error) {
	c := C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))
	return ModifierType(c), nil
}

// castWidget takes a native GtkWidget and casts it to the appropriate Go struct.
func castWidget(c *C.GtkWidget) IWidget {
	wdt := new(Widget)
	wdt.Object = ToGoObject(unsafe.Pointer(c))
	return wdt
}

// Align is a representation of GTK's GtkAlign.
type Align int

const (
	ALIGN_FILL   Align = C.GTK_ALIGN_FILL
	ALIGN_START  Align = C.GTK_ALIGN_START
	ALIGN_END    Align = C.GTK_ALIGN_END
	ALIGN_CENTER Align = C.GTK_ALIGN_CENTER
)

func marshalAlign(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Align(c), nil
}

// StateFlags is a representation of GTK's GtkStateFlags.
type StateFlags int

const (
	STATE_FLAG_NORMAL       StateFlags = C.GTK_STATE_FLAG_NORMAL
	STATE_FLAG_ACTIVE       StateFlags = C.GTK_STATE_FLAG_ACTIVE
	STATE_FLAG_PRELIGHT     StateFlags = C.GTK_STATE_FLAG_PRELIGHT
	STATE_FLAG_SELECTED     StateFlags = C.GTK_STATE_FLAG_SELECTED
	STATE_FLAG_INSENSITIVE  StateFlags = C.GTK_STATE_FLAG_INSENSITIVE
	STATE_FLAG_INCONSISTENT StateFlags = C.GTK_STATE_FLAG_INCONSISTENT
	STATE_FLAG_FOCUSED      StateFlags = C.GTK_STATE_FLAG_FOCUSED
	STATE_FLAG_BACKDROP     StateFlags = C.GTK_STATE_FLAG_BACKDROP
)

func marshalStateFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return StateFlags(c), nil
}

// IconSize is a representation of GTK's GtkIconSize.
type IconSize int

const (
	ICON_SIZE_INVALID       IconSize = C.GTK_ICON_SIZE_INVALID
	ICON_SIZE_MENU          IconSize = C.GTK_ICON_SIZE_MENU
	ICON_SIZE_SMALL_TOOLBAR IconSize = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	ICON_SIZE_LARGE_TOOLBAR IconSize = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	ICON_SIZE_BUTTON        IconSize = C.GTK_ICON_SIZE_BUTTON
	ICON_SIZE_DND           IconSize = C.GTK_ICON_SIZE_DND
	ICON_SIZE_DIALOG        IconSize = C.GTK_ICON_SIZE_DIALOG
)

func marshalIconSize(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return IconSize(c), nil
}

// ImageType is a representation of GTK's GtkImageType.
type ImageType int

const (
	IMAGE_EMPTY     ImageType = C.GTK_IMAGE_EMPTY
	IMAGE_PIXBUF    ImageType = C.GTK_IMAGE_PIXBUF
	IMAGE_STOCK     ImageType = C.GTK_IMAGE_STOCK
	IMAGE_ICON_SET  ImageType = C.GTK_IMAGE_ICON_SET
	IMAGE_ANIMATION ImageType = C.GTK_IMAGE_ANIMATION
	IMAGE_ICON_NAME ImageType = C.GTK_IMAGE_ICON_NAME
	IMAGE_GICON     ImageType = C.GTK_IMAGE_GICON
)

func marshalImageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ImageType(c), nil
}

// EntryIconPosition is a representation of GTK's GtkEntryIconPosition.
type EntryIconPosition int

const (
	ENTRY_ICON_PRIMARY   EntryIconPosition = C.GTK_ENTRY_ICON_PRIMARY
	ENTRY_ICON_SECONDARY EntryIconPosition = C.GTK_ENTRY_ICON_SECONDARY
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EntryIconPosition(c), nil
}

// Orientation is a representation of GTK's GtkOrientation.
type Orientation int

const (
	ORIENTATION_HORIZONTAL Orientation = C.GTK_ORIENTATION_HORIZONTAL
	ORIENTATION_VERTICAL   Orientation = C.GTK_ORIENTATION_VERTICAL
)

func marshalOrientation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Orientation(c), nil
}

// PackType is a representation of GTK's GtkPackType.
type PackType int

const (
	PACK_START PackType = C.GTK_PACK_START
	PACK_END   PackType = C.GTK_PACK_END
)

func marshalPackType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PackType(c), nil
}

// Justification is a representation of GTK's GtkJustification.
type Justification int

const (
	JUSTIFY_LEFT   Justification = C.GTK_JUSTIFY_LEFT
	JUSTIFY_RIGHT  Justification = C.GTK_JUSTIFY_RIGHT
	JUSTIFY_CENTER Justification = C.GTK_JUSTIFY_CENTER
	JUSTIFY_FILL   Justification = C.GTK_JUSTIFY_FILL
)

func marshalJustification(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Justification(c), nil
}

// SizeRequestMode is a representation of GTK's GtkSizeRequestMode.
type SizeRequestMode int

const (
	SIZE_REQUEST_HEIGHT_FOR_WIDTH SizeRequestMode = C.GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH
	SIZE_REQUEST_WIDTH_FOR_HEIGHT SizeRequestMode = C.GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT
	SIZE_REQUEST_CONSTANT_SIZE    SizeRequestMode = C.GTK_SIZE_REQUEST_CONSTANT_SIZE
)

func marshalSizeRequestMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SizeRequestMode(c), nil
}

func MainDoEvent(event *Event) {
	C.gtk_main_do_event(event.native())
}

// PolicyType is a representation of GTK's GtkPolicyType.
type PolicyType int

const (
	POLICY_ALWAYS    PolicyType = C.GTK_POLICY_ALWAYS
	POLICY_AUTOMATIC PolicyType = C.GTK_POLICY_AUTOMATIC
	POLICY_NEVER     PolicyType = C.GTK_POLICY_NEVER
)

// CornerType is a representation of GTK's GtkCornerType.
type CornerType int

const (
	CORNER_TOP_LEFT     CornerType = C.GTK_CORNER_TOP_LEFT
	CORNER_BOTTOM_LEFT  CornerType = C.GTK_CORNER_BOTTOM_LEFT
	CORNER_TOP_RIGHT    CornerType = C.GTK_CORNER_TOP_RIGHT
	CORNER_BOTTOM_RIGHT CornerType = C.GTK_CORNER_BOTTOM_RIGHT
)

// ShadowType is a representation of GTK's GtkShadowType.
type ShadowType int

const (
	SHADOW_NONE       ShadowType = C.GTK_SHADOW_NONE
	SHADOW_IN         ShadowType = C.GTK_SHADOW_IN
	SHADOW_OUT        ShadowType = C.GTK_SHADOW_OUT
	SHADOW_ETCHED_IN  ShadowType = C.GTK_SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT ShadowType = C.GTK_SHADOW_ETCHED_OUT
)

// SensitivityType is a representation of GTK's GtkSensitivityType
type SensitivityType int

const (
	SENSITIVITY_AUTO SensitivityType = C.GTK_SENSITIVITY_AUTO
	SENSITIVITY_ON   SensitivityType = C.GTK_SENSITIVITY_ON
	SENSITIVITY_OFF  SensitivityType = C.GTK_SENSITIVITY_OFF
)
