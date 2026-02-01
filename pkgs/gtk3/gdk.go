package gtk3

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"unsafe"
)

// WindowEdge is a representation of GDK's GdkWindowEdge
type WindowEdge int

const (
	WINDOW_EDGE_NORTH_WEST WindowEdge = C.GDK_WINDOW_EDGE_NORTH_WEST
	WINDOW_EDGE_NORTH      WindowEdge = C.GDK_WINDOW_EDGE_NORTH
	WINDOW_EDGE_NORTH_EAST WindowEdge = C.GDK_WINDOW_EDGE_NORTH_EAST
	WINDOW_EDGE_WEST       WindowEdge = C.GDK_WINDOW_EDGE_WEST
	WINDOW_EDGE_EAST       WindowEdge = C.GDK_WINDOW_EDGE_EAST
	WINDOW_EDGE_SOUTH_WEST WindowEdge = C.GDK_WINDOW_EDGE_SOUTH_WEST
	WINDOW_EDGE_SOUTH      WindowEdge = C.GDK_WINDOW_EDGE_SOUTH
	WINDOW_EDGE_SOUTH_EAST WindowEdge = C.GDK_WINDOW_EDGE_SOUTH_EAST
)

const CURRENT_TIME = C.GDK_CURRENT_TIME

// ButtonType constants
type ButtonType uint

const (
	BUTTON_PRIMARY   ButtonType = C.GDK_BUTTON_PRIMARY
	BUTTON_MIDDLE    ButtonType = C.GDK_BUTTON_MIDDLE
	BUTTON_SECONDARY ButtonType = C.GDK_BUTTON_SECONDARY
)

// RGBA To create a GdkRGBA you have to use NewRGBA function.
type RGBA struct {
	rgba *C.GdkRGBA
}

func marshalRGBA(p uintptr) (any, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRGBA(unsafe.Pointer(c)), nil
}

func WrapRGBA(p unsafe.Pointer) *RGBA {
	return wrapRGBA((*C.GdkRGBA)(p))
}

func wrapRGBA(cRgba *C.GdkRGBA) *RGBA {
	if cRgba == nil {
		return nil
	}
	return &RGBA{cRgba}
}

func NewRGBA(values ...float64) *RGBA {
	cRgba := new(C.GdkRGBA)
	for i, value := range values {
		switch i {
		case 0:
			cRgba.red = C.gdouble(value)
		case 1:
			cRgba.green = C.gdouble(value)
		case 2:
			cRgba.blue = C.gdouble(value)
		case 3:
			cRgba.alpha = C.gdouble(value)
		}
	}
	return wrapRGBA(cRgba)
}

func (c *RGBA) Floats() []float64 {
	return []float64{
		float64(c.rgba.red),
		float64(c.rgba.green),
		float64(c.rgba.blue),
		float64(c.rgba.alpha)}
}

func (c *RGBA) Native() uintptr {
	return uintptr(unsafe.Pointer(c.rgba))
}

// SetColors sets all colors values in the RGBA.
func (c *RGBA) SetColors(r, g, b, a float64) {
	c.rgba.red = C.gdouble(r)
	c.rgba.green = C.gdouble(g)
	c.rgba.blue = C.gdouble(b)
	c.rgba.alpha = C.gdouble(a)
}

/*
GetRed
  The following methods (Get/Set) are made for
  more convenient use of the GdkRGBA object
*/
// GetRed get red value from the RGBA.
func (c *RGBA) GetRed() float64 {
	return float64(c.rgba.red)
}

// GetGreen get green value from the RGBA.
func (c *RGBA) GetGreen() float64 {
	return float64(c.rgba.green)
}

// GetBlue get blue value from the RGBA.
func (c *RGBA) GetBlue() float64 {
	return float64(c.rgba.blue)
}

// GetAlpha get alpha value from the RGBA.
func (c *RGBA) GetAlpha() float64 {
	return float64(c.rgba.alpha)
}

// SetRed set red value in the RGBA.
func (c *RGBA) SetRed(red float64) {
	c.rgba.red = C.gdouble(red)
}

// SetGreen set green value in the RGBA.
func (c *RGBA) SetGreen(green float64) {
	c.rgba.green = C.gdouble(green)
}

// SetBlue set blue value in the RGBA.
func (c *RGBA) SetBlue(blue float64) {
	c.rgba.blue = C.gdouble(blue)
}

// SetAlpha set alpha value in the RGBA.
func (c *RGBA) SetAlpha(alpha float64) {
	c.rgba.alpha = C.gdouble(alpha)
}

// Parse is a representation of gdk_rgba_parse().
func (c *RGBA) Parse(spec string) bool {
	cstr := (*C.gchar)(C.CString(spec))
	defer C.free(unsafe.Pointer(cstr))
	return GoBool(C.gdk_rgba_parse(c.rgba, cstr))
}

// String is a representation of gdk_rgba_to_string().
func (c *RGBA) String() string {
	return C.GoString((*C.char)(C.gdk_rgba_to_string(c.rgba)))
}

// Free is a representation of gdk_rgba_free().
func (c *RGBA) Free() {
	C.gdk_rgba_free(c.rgba)
}

// Equal is a representation of gdk_rgba_equal().
func (c *RGBA) Equal(rgba *RGBA) bool {
	return GoBool(C.gdk_rgba_equal(
		C.gconstpointer(c.rgba),
		C.gconstpointer(rgba.rgba)))
}

// Hash is a representation of gdk_rgba_hash().
func (c *RGBA) Hash() uint {
	return uint(C.gdk_rgba_hash(C.gconstpointer(c.rgba)))
}

/*
 * GdkAtom
 */

// Atom is a representation of GDK's GdkAtom.
type Atom uintptr

// native returns the underlying GdkAtom.
func (v Atom) native() C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(uintptr(v)))
}

func (v Atom) Name() string {
	c := C.gdk_atom_name(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) Atom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), CBool(onlyIfExists))
	return Atom(uintptr(unsafe.Pointer(c)))
}

// Selections
const (
	SELECTION_PRIMARY       Atom = 1
	SELECTION_SECONDARY     Atom = 2
	SELECTION_CLIPBOARD     Atom = 69
	TARGET_BITMAP           Atom = 5
	TARGET_COLORMAP         Atom = 7
	TARGET_DRAWABLE         Atom = 17
	TARGET_PIXMAP           Atom = 20
	TARGET_STRING           Atom = 31
	SELECTION_TYPE_ATOM     Atom = 4
	SELECTION_TYPE_BITMAP   Atom = 5
	SELECTION_TYPE_COLORMAP Atom = 7
	SELECTION_TYPE_DRAWABLE Atom = 17
	SELECTION_TYPE_INTEGER  Atom = 19
	SELECTION_TYPE_PIXMAP   Atom = 20
	SELECTION_TYPE_WINDOW   Atom = 33
	SELECTION_TYPE_STRING   Atom = 31
)

// VisualType is a representation of GDK's GdkVisualType.
type VisualType int

const (
	VISUAL_STATIC_GRAY  VisualType = C.GDK_VISUAL_STATIC_GRAY
	VISUAL_GRAYSCALE    VisualType = C.GDK_VISUAL_GRAYSCALE
	VISUAL_STATIC_COLOR VisualType = C.GDK_VISUAL_STATIC_COLOR
	ISUAL_PSEUDO_COLOR  VisualType = C.GDK_VISUAL_PSEUDO_COLOR
	VISUAL_TRUE_COLOR   VisualType = C.GDK_VISUAL_TRUE_COLOR
	VISUAL_DIRECT_COLOR VisualType = C.GDK_VISUAL_DIRECT_COLOR
)

func marshalVisualType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return VisualType(c), nil
}

// EventType is a representation of GDK's GdkEventType.
// Do not confuse these event types with the signals that GTK+ widgets emit
type EventType int

func marshalEventType(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EventType(c), nil
}

const (
	EVENT_NOTHING             EventType = C.GDK_NOTHING
	EVENT_DELETE              EventType = C.GDK_DELETE
	EVENT_DESTROY             EventType = C.GDK_DESTROY
	EVENT_EXPOSE              EventType = C.GDK_EXPOSE
	EVENT_MOTION_NOTIFY       EventType = C.GDK_MOTION_NOTIFY
	EVENT_BUTTON_PRESS        EventType = C.GDK_BUTTON_PRESS
	EVENT_2BUTTON_PRESS       EventType = C.GDK_2BUTTON_PRESS
	EVENT_DOUBLE_BUTTON_PRESS EventType = C.GDK_DOUBLE_BUTTON_PRESS
	EVENT_3BUTTON_PRESS       EventType = C.GDK_3BUTTON_PRESS
	EVENT_TRIPLE_BUTTON_PRESS EventType = C.GDK_TRIPLE_BUTTON_PRESS
	EVENT_BUTTON_RELEASE      EventType = C.GDK_BUTTON_RELEASE
	EVENT_KEY_PRESS           EventType = C.GDK_KEY_PRESS
	EVENT_KEY_RELEASE         EventType = C.GDK_KEY_RELEASE
	EVENT_ENTER_NOTIFY        EventType = C.GDK_ENTER_NOTIFY
	EVENT_LEAVE_NOTIFY        EventType = C.GDK_LEAVE_NOTIFY
	EVENT_FOCUS_CHANGE        EventType = C.GDK_FOCUS_CHANGE
	EVENT_CONFIGURE           EventType = C.GDK_CONFIGURE
	EVENT_MAP                 EventType = C.GDK_MAP
	EVENT_UNMAP               EventType = C.GDK_UNMAP
	EVENT_PROPERTY_NOTIFY     EventType = C.GDK_PROPERTY_NOTIFY
	EVENT_SELECTION_CLEAR     EventType = C.GDK_SELECTION_CLEAR
	EVENT_SELECTION_REQUEST   EventType = C.GDK_SELECTION_REQUEST
	EVENT_SELECTION_NOTIFY    EventType = C.GDK_SELECTION_NOTIFY
	EVENT_PROXIMITY_IN        EventType = C.GDK_PROXIMITY_IN
	EVENT_PROXIMITY_OUT       EventType = C.GDK_PROXIMITY_OUT
	EVENT_DRAG_ENTER          EventType = C.GDK_DRAG_ENTER
	EVENT_DRAG_LEAVE          EventType = C.GDK_DRAG_LEAVE
	EVENT_DRAG_MOTION         EventType = C.GDK_DRAG_MOTION
	EVENT_DRAG_STATUS         EventType = C.GDK_DRAG_STATUS
	EVENT_DROP_START          EventType = C.GDK_DROP_START
	EVENT_DROP_FINISHED       EventType = C.GDK_DROP_FINISHED
	EVENT_CLIENT_EVENT        EventType = C.GDK_CLIENT_EVENT
	EVENT_VISIBILITY_NOTIFY   EventType = C.GDK_VISIBILITY_NOTIFY
	EVENT_SCROLL              EventType = C.GDK_SCROLL
	EVENT_WINDOW_STATE        EventType = C.GDK_WINDOW_STATE
	EVENT_SETTING             EventType = C.GDK_SETTING
	EVENT_OWNER_CHANGE        EventType = C.GDK_OWNER_CHANGE
	EVENT_GRAB_BROKEN         EventType = C.GDK_GRAB_BROKEN
	EVENT_DAMAGE              EventType = C.GDK_DAMAGE
	EVENT_TOUCH_BEGIN         EventType = C.GDK_TOUCH_BEGIN
	EVENT_TOUCH_UPDATE        EventType = C.GDK_TOUCH_UPDATE
	EVENT_TOUCH_END           EventType = C.GDK_TOUCH_END
	EVENT_TOUCH_CANCEL        EventType = C.GDK_TOUCH_CANCEL
	EVENT_LAST                EventType = C.GDK_EVENT_LAST
)

// added by terrak
// EventMask is a representation of GDK's GdkEventMask.
type EventMask int

const (
	EXPOSURE_MASK            EventMask = C.GDK_EXPOSURE_MASK
	POINTER_MOTION_MASK      EventMask = C.GDK_POINTER_MOTION_MASK
	POINTER_MOTION_HINT_MASK EventMask = C.GDK_POINTER_MOTION_HINT_MASK
	BUTTON_MOTION_MASK       EventMask = C.GDK_BUTTON_MOTION_MASK
	BUTTON1_MOTION_MASK      EventMask = C.GDK_BUTTON1_MOTION_MASK
	BUTTON2_MOTION_MASK      EventMask = C.GDK_BUTTON2_MOTION_MASK
	BUTTON3_MOTION_MASK      EventMask = C.GDK_BUTTON3_MOTION_MASK
	BUTTON_PRESS_MASK        EventMask = C.GDK_BUTTON_PRESS_MASK
	BUTTON_RELEASE_MASK      EventMask = C.GDK_BUTTON_RELEASE_MASK
	KEY_PRESS_MASK           EventMask = C.GDK_KEY_PRESS_MASK
	KEY_RELEASE_MASK         EventMask = C.GDK_KEY_RELEASE_MASK
	ENTER_NOTIFY_MASK        EventMask = C.GDK_ENTER_NOTIFY_MASK
	LEAVE_NOTIFY_MASK        EventMask = C.GDK_LEAVE_NOTIFY_MASK
	FOCUS_CHANGE_MASK        EventMask = C.GDK_FOCUS_CHANGE_MASK
	STRUCTURE_MASK           EventMask = C.GDK_STRUCTURE_MASK
	PROPERTY_CHANGE_MASK     EventMask = C.GDK_PROPERTY_CHANGE_MASK
	VISIBILITY_NOTIFY_MASK   EventMask = C.GDK_VISIBILITY_NOTIFY_MASK
	PROXIMITY_IN_MASK        EventMask = C.GDK_PROXIMITY_IN_MASK
	PROXIMITY_OUT_MASK       EventMask = C.GDK_PROXIMITY_OUT_MASK
	SUBSTRUCTURE_MASK        EventMask = C.GDK_SUBSTRUCTURE_MASK
	SCROLL_MASK              EventMask = C.GDK_SCROLL_MASK
	TOUCH_MASK               EventMask = C.GDK_TOUCH_MASK
	SMOOTH_SCROLL_MASK       EventMask = C.GDK_SMOOTH_SCROLL_MASK
	ALL_EVENTS_MASK          EventMask = C.GDK_ALL_EVENTS_MASK
)

type CrossingMode int

const (
	CROSSING_NORMAL        CrossingMode = C.GDK_CROSSING_NORMAL
	CROSSING_GRAB          CrossingMode = C.GDK_CROSSING_GRAB
	CROSSING_UNGRAB        CrossingMode = C.GDK_CROSSING_UNGRAB
	CROSSING_GTK_GRAB      CrossingMode = C.GDK_CROSSING_GTK_GRAB
	CROSSING_GTK_UNGRAB    CrossingMode = C.GDK_CROSSING_GTK_UNGRAB
	CROSSING_STATE_CHANGED CrossingMode = C.GDK_CROSSING_STATE_CHANGED
	CROSSING_TOUCH_BEGIN   CrossingMode = C.GDK_CROSSING_TOUCH_BEGIN
	CROSSING_TOUCH_END     CrossingMode = C.GDK_CROSSING_TOUCH_END
	CROSSING_DEVICE_SWITCH CrossingMode = C.GDK_CROSSING_DEVICE_SWITCH
)

type NotifyType int

const (
	NOTIFY_ANCESTOR          NotifyType = C.GDK_NOTIFY_ANCESTOR
	NOTIFY_VIRTUAL           NotifyType = C.GDK_NOTIFY_VIRTUAL
	NOTIFY_INFERIOR          NotifyType = C.GDK_NOTIFY_INFERIOR
	NOTIFY_NONLINEAR         NotifyType = C.GDK_NOTIFY_NONLINEAR
	NOTIFY_NONLINEAR_VIRTUAL NotifyType = C.GDK_NOTIFY_NONLINEAR_VIRTUAL
	NOTIFY_UNKNOWN           NotifyType = C.GDK_NOTIFY_UNKNOWN
)

// DragAction is a representation of GDK's GdkDragAction.
type DragAction int

const (
	ACTION_DEFAULT DragAction = C.GDK_ACTION_DEFAULT
	ACTION_COPY    DragAction = C.GDK_ACTION_COPY
	ACTION_MOVE    DragAction = C.GDK_ACTION_MOVE
	ACTION_LINK    DragAction = C.GDK_ACTION_LINK
	ACTION_PRIVATE DragAction = C.GDK_ACTION_PRIVATE
	ACTION_ASK     DragAction = C.GDK_ACTION_ASK
)
