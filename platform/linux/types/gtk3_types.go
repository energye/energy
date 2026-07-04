//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import (
	"github.com/energye/lcl/types"
	"unsafe"
)

type GDouble = float64
type GBoolean = int32
type PGdkWindow = uintptr
type GInt8 = int8
type GInt16 = int16
type GUint32 = uint32
type GUint = uint32
type GInt = int32
type GPChar = uintptr
type GUInt16 = uint16
type GUInt8 = uint8
type GULong = uint
type PGDouble = uintptr
type GPointer = uintptr
type PGdkDevice = uintptr
type PGtkWidget = uintptr
type PEventKey = uintptr
type PEventButton = uintptr
type PEventCrossing = uintptr
type PEventConfigure = uintptr
type PContext = uintptr
type PDragContext = uintptr
type PSelectionData = uintptr

// PGdkEventFocus = TGdkEventFocus
type PGdkEventFocus uintptr

type GdkRGBA struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

func (m *PGdkEventFocus) Get() *TGdkEventFocus {
	return (*TGdkEventFocus)(unsafe.Pointer(m))
}

type TGdkEventFocus struct {
	Type      TGdkEventType
	Window    PGdkWindow
	SendEvent GInt8
	In        GInt16
}

// const

const (
	ColorSchemePreferLight  = 0
	ColorSchemePreferDark   = 1
	ColorSchemeHighContrast = 2
)

type StyleProviderPriority = uint

const (
	STYLE_PROVIDER_PRIORITY_FALLBACK    StyleProviderPriority = 1
	STYLE_PROVIDER_PRIORITY_THEME       StyleProviderPriority = 200
	STYLE_PROVIDER_PRIORITY_SETTINGS    StyleProviderPriority = 400
	STYLE_PROVIDER_PRIORITY_APPLICATION StyleProviderPriority = 600
	STYLE_PROVIDER_PRIORITY_USER        StyleProviderPriority = 800
)

type DragAction = uint

const (
	ACTION_DEFAULT DragAction = 1 << 0 // 1
	ACTION_COPY    DragAction = 1 << 1 // 2
	ACTION_MOVE    DragAction = 1 << 2 // 4
	ACTION_LINK    DragAction = 1 << 3 // 8
	ACTION_PRIVATE DragAction = 1 << 4 // 16
	ACTION_ASK     DragAction = 1 << 5 // 32
)

type TAtom = uintptr

// Selections
const (
	SELECTION_PRIMARY       TAtom = 1
	SELECTION_SECONDARY     TAtom = 2
	SELECTION_CLIPBOARD     TAtom = 69
	TARGET_BITMAP           TAtom = 5
	TARGET_COLORMAP         TAtom = 7
	TARGET_DRAWABLE         TAtom = 17
	TARGET_PIXMAP           TAtom = 20
	TARGET_STRING           TAtom = 31
	SELECTION_TYPE_ATOM     TAtom = 4
	SELECTION_TYPE_BITMAP   TAtom = 5
	SELECTION_TYPE_COLORMAP TAtom = 7
	SELECTION_TYPE_DRAWABLE TAtom = 17
	SELECTION_TYPE_INTEGER  TAtom = 19
	SELECTION_TYPE_PIXMAP   TAtom = 20
	SELECTION_TYPE_WINDOW   TAtom = 33
	SELECTION_TYPE_STRING   TAtom = 31
)

// EventType is a representation of GDK's GdkEventType.
// Do not confuse these event types with the signals that GTK+ widgets emit
type EventType = int

const (
	EVENT_NOTHING             EventType = 0
	EVENT_DELETE              EventType = 1
	EVENT_DESTROY             EventType = 2
	EVENT_EXPOSE              EventType = 3
	EVENT_MOTION_NOTIFY       EventType = 4
	EVENT_BUTTON_PRESS        EventType = 5
	EVENT_2BUTTON_PRESS       EventType = 6
	EVENT_DOUBLE_BUTTON_PRESS EventType = 6 // 和 2BUTTON_PRESS 相同
	EVENT_3BUTTON_PRESS       EventType = 7
	EVENT_TRIPLE_BUTTON_PRESS EventType = 7 // 和 3BUTTON_PRESS 相同
	EVENT_BUTTON_RELEASE      EventType = 8
	EVENT_KEY_PRESS           EventType = 9
	EVENT_KEY_RELEASE         EventType = 10
	EVENT_ENTER_NOTIFY        EventType = 11
	EVENT_LEAVE_NOTIFY        EventType = 12
	EVENT_FOCUS_CHANGE        EventType = 13
	EVENT_CONFIGURE           EventType = 14
	EVENT_MAP                 EventType = 15
	EVENT_UNMAP               EventType = 16
	EVENT_PROPERTY_NOTIFY     EventType = 17
	EVENT_SELECTION_CLEAR     EventType = 18
	EVENT_SELECTION_REQUEST   EventType = 19
	EVENT_SELECTION_NOTIFY    EventType = 20
	EVENT_PROXIMITY_IN        EventType = 21
	EVENT_PROXIMITY_OUT       EventType = 22
	EVENT_DRAG_ENTER          EventType = 23
	EVENT_DRAG_LEAVE          EventType = 24
	EVENT_DRAG_MOTION         EventType = 25
	EVENT_DRAG_STATUS         EventType = 26
	EVENT_DROP_START          EventType = 27
	EVENT_DROP_FINISHED       EventType = 28
	EVENT_CLIENT_EVENT        EventType = 29
	EVENT_VISIBILITY_NOTIFY   EventType = 30
	EVENT_SCROLL              EventType = 31
	EVENT_WINDOW_STATE        EventType = 32
	EVENT_SETTING             EventType = 33
	EVENT_OWNER_CHANGE        EventType = 34
	EVENT_GRAB_BROKEN         EventType = 35
	EVENT_DAMAGE              EventType = 36
	EVENT_TOUCH_BEGIN         EventType = 37
	EVENT_TOUCH_UPDATE        EventType = 38
	EVENT_TOUCH_END           EventType = 39
	EVENT_TOUCH_CANCEL        EventType = 40
	EVENT_LAST                EventType = 41
)

// ButtonType constants
type ButtonType = uint

const (
	BUTTON_PRIMARY   ButtonType = 1
	BUTTON_MIDDLE    ButtonType = 2
	BUTTON_SECONDARY ButtonType = 3
)

type CrossingMode = int

const (
	CROSSING_NORMAL CrossingMode = iota
	CROSSING_GRAB
	CROSSING_UNGRAB
	CROSSING_GTK_GRAB
	CROSSING_GTK_UNGRAB
	CROSSING_STATE_CHANGED
	CROSSING_TOUCH_BEGIN
	CROSSING_TOUCH_END
	CROSSING_DEVICE_SWITCH
)

type NotifyType = int

const (
	NOTIFY_ANCESTOR NotifyType = iota
	NOTIFY_VIRTUAL
	NOTIFY_INFERIOR
	NOTIFY_NONLINEAR
	NOTIFY_NONLINEAR_VIRTUAL
	NOTIFY_UNKNOWN
)

type TGdkEventType = int32

const (
	TGdkEventTypeMinValue   = -0x7FFFFFFF
	GDK_NOTHING             = -1
	GDK_DELETE              = 0
	GDK_DESTROY             = 1
	GDK_EXPOSE              = 2
	GDK_MOTION_NOTIFY       = 3
	GDK_BUTTON_PRESS        = 4
	GDK_DOUBLE_BUTTON_PRESS = 5
	GDK_2BUTTON_PRESS       = 5
	GDK_TRIPLE_BUTTON_PRESS = 6
	GDK_3BUTTON_PRESS       = 6
	GDK_BUTTON_RELEASE      = 7
	GDK_KEY_PRESS           = 8
	GDK_KEY_RELEASE         = 9
	GDK_ENTER_NOTIFY        = 10
	GDK_LEAVE_NOTIFY        = 11
	GDK_FOCUS_CHANGE        = 12
	GDK_CONFIGURE           = 13
	GDK_MAP                 = 14
	GDK_UNMAP               = 15
	GDK_PROPERTY_NOTIFY     = 16
	GDK_SELECTION_CLEAR     = 17
	GDK_SELECTION_REQUEST   = 18
	GDK_SELECTION_NOTIFY    = 19
	GDK_PROXIMITY_IN        = 20
	GDK_PROXIMITY_OUT       = 21
	GDK_DRAG_ENTER          = 22
	GDK_DRAG_LEAVE          = 23
	GDK_DRAG_MOTION_        = 24
	GDK_DRAG_STATUS_        = 25
	GDK_DROP_START          = 26
	GDK_DROP_FINISHED       = 27
	GDK_CLIENT_EVENT        = 28
	GDK_VISIBILITY_NOTIFY   = 29
	GDK_SCROLL              = 31
	GDK_WINDOW_STATE        = 32
	GDK_SETTING             = 33
	GDK_OWNER_CHANGE        = 34
	GDK_GRAB_BROKEN         = 35
	GDK_DAMAGE              = 36
	GDK_TOUCH_BEGIN         = 37
	GDK_TOUCH_UPDATE        = 38
	GDK_TOUCH_END           = 39
	GDK_TOUCH_CANCEL        = 40
	GDK_TOUCHPAD_SWIPE      = 41
	GDK_TOUCHPAD_PINCH      = 42
	GDK_PAD_BUTTON_PRESS    = 43
	GDK_PAD_BUTTON_RELEASE  = 44
	GDK_PAD_RING            = 45
	GDK_PAD_STRIP           = 46
	GDK_PAD_GROUP_MODE      = 47
	GDK_EVENT_LAST          = 48
	TGdkEventTypeMaxValue   = 0x7FFFFFFF
)

type TGdkCrossingMode = int32

const (
	TGdkCrossingModeMinValue   = -0x7FFFFFFF
	GDK_CROSSING_NORMAL        = 0
	GDK_CROSSING_GRAB          = 1
	GDK_CROSSING_UNGRAB        = 2
	GDK_CROSSING_GTK_GRAB      = 3
	GDK_CROSSING_GTK_UNGRAB    = 4
	GDK_CROSSING_STATE_CHANGED = 5
	GDK_CROSSING_TOUCH_BEGIN   = 6
	GDK_CROSSING_TOUCH_END     = 7
	GDK_CROSSING_DEVICE_SWITCH = 8
	TGdkCrossingModeMaxValue   = 0x7FFFFFFF
)

type TGdkNotifyType = int32

const (
	TGdkNotifyTypeMinValue       = -0x7FFFFFFF
	GDK_NOTIFY_ANCESTOR          = 0
	GDK_NOTIFY_VIRTUAL           = 1
	GDK_NOTIFY_INFERIOR          = 2
	GDK_NOTIFY_NONLINEAR         = 3
	GDK_NOTIFY_NONLINEAR_VIRTUAL = 4
	GDK_NOTIFY_UNKNOWN           = 5
	TGdkNotifyTypeMaxValue       = 0x7FFFFFFF
)

type TGdkModifierTypeIdx = int32

const (
	TGdkModifierTypeIdxMinValue   = 0
	GDK_SHIFT_MASK                = 0
	GDK_LOCK_MASK                 = 1
	GDK_CONTROL_MASK              = 2
	GDK_MOD1_MASK                 = 3
	GDK_MOD2_MASK                 = 4
	GDK_MOD3_MASK                 = 5
	GDK_MOD4_MASK                 = 6
	GDK_MOD5_MASK                 = 7
	GDK_BUTTON1_MASK              = 8
	GDK_BUTTON2_MASK              = 9
	GDK_BUTTON3_MASK              = 10
	GDK_BUTTON4_MASK              = 11
	GDK_BUTTON5_MASK              = 12
	GDK_MODIFIER_RESERVED_13_MASK = 13
	GDK_MODIFIER_RESERVED_14_MASK = 14
	GDK_MODIFIER_RESERVED_15_MASK = 15
	GDK_MODIFIER_RESERVED_16_MASK = 16
	GDK_MODIFIER_RESERVED_17_MASK = 17
	GDK_MODIFIER_RESERVED_18_MASK = 18
	GDK_MODIFIER_RESERVED_19_MASK = 19
	GDK_MODIFIER_RESERVED_20_MASK = 20
	GDK_MODIFIER_RESERVED_21_MASK = 21
	GDK_MODIFIER_RESERVED_22_MASK = 22
	GDK_MODIFIER_RESERVED_23_MASK = 23
	GDK_MODIFIER_RESERVED_24_MASK = 24
	GDK_MODIFIER_RESERVED_25_MASK = 25
	GDK_SUPER_MASK                = 26
	GDK_HYPER_MASK                = 27
	GDK_META_MASK                 = 28
	GDK_MODIFIER_RESERVED_29_MASK = 29
	GDK_RELEASE_MASK              = 30
	TGdkModifierTypeIdxMaxValue   = 31
)

// TGdkModifierType :  TGdkModifierTypeIdx
type TGdkModifierType = types.TSet

// Status is a representation of Cairo's cairo_status_t.
type Status int

const (
	STATUS_SUCCESS Status = iota
	STATUS_NO_MEMORY
	STATUS_INVALID_RESTORE
	STATUS_INVALID_POP_GROUP
	STATUS_NO_CURRENT_POINT
	STATUS_INVALID_MATRIX
	STATUS_INVALID_STATUS
	STATUS_NULL_POINTER
	STATUS_INVALID_STRING
	STATUS_INVALID_PATH_DATA
	STATUS_READ_ERROR
	STATUS_WRITE_ERROR
	STATUS_SURFACE_FINISHED
	STATUS_SURFACE_TYPE_MISMATCH
	STATUS_PATTERN_TYPE_MISMATCH
	STATUS_INVALID_CONTENT
	STATUS_INVALID_FORMAT
	STATUS_INVALID_VISUAL
	STATUS_FILE_NOT_FOUND
	STATUS_INVALID_DASH
	STATUS_INVALID_DSC_COMMENT
	STATUS_INVALID_INDEX
	STATUS_CLIP_NOT_REPRESENTABLE
	STATUS_TEMP_FILE_ERROR
	STATUS_INVALID_STRIDE
	STATUS_FONT_TYPE_MISMATCH
	STATUS_USER_FONT_IMMUTABLE
	STATUS_USER_FONT_ERROR
	STATUS_NEGATIVE_COUNT
	STATUS_INVALID_CLUSTERS
	STATUS_INVALID_SLANT
	STATUS_INVALID_WEIGHT
	STATUS_INVALID_SIZE
	STATUS_USER_FONT_NOT_IMPLEMENTED
	STATUS_DEVICE_TYPE_MISMATCH
	STATUS_DEVICE_ERROR
	// STATUS_INVALID_MESH_CONSTRUCTION Status = C.CAIRO_STATUS_INVALID_MESH_CONSTRUCTION (since 1.12)
	// STATUS_DEVICE_FINISHED           Status = C.CAIRO_STATUS_DEVICE_FINISHED (since 1.12)
)

// WindowEdge is a representation of GDK's GdkWindowEdge
type WindowEdge int

const (
	GDK_WINDOW_EDGE_NORTH_WEST WindowEdge = 0 // 左上
	GDK_WINDOW_EDGE_NORTH      WindowEdge = 1 // 上
	GDK_WINDOW_EDGE_NORTH_EAST WindowEdge = 2 // 右上
	GDK_WINDOW_EDGE_WEST       WindowEdge = 3 // 左
	GDK_WINDOW_EDGE_EAST       WindowEdge = 4 // 右
	GDK_WINDOW_EDGE_SOUTH_WEST WindowEdge = 5 // 左下
	GDK_WINDOW_EDGE_SOUTH      WindowEdge = 6 // 下
	GDK_WINDOW_EDGE_SOUTH_EAST WindowEdge = 7 // 右下
)

type TGdkWMDecorationIdx = uint

const (
	TGdkWMDecorationIdxMinValue TGdkWMDecorationIdx = 0
	GDK_DECOR_ALL               TGdkWMDecorationIdx = 0
	GDK_DECOR_BORDER            TGdkWMDecorationIdx = 1
	GDK_DECOR_RESIZEH           TGdkWMDecorationIdx = 2
	GDK_DECOR_TITLE             TGdkWMDecorationIdx = 3
	GDK_DECOR_MENU              TGdkWMDecorationIdx = 4
	GDK_DECOR_MINIMIZE          TGdkWMDecorationIdx = 5
	GDK_DECOR_MAXIMIZE          TGdkWMDecorationIdx = 6
	TGdkWMDecorationIdxMaxValue TGdkWMDecorationIdx = 31
)

// TGdkWMDecoration : TGdkWMDecorationIdx
type TGdkWMDecoration = types.TSet

// ==================== GTK3 Enums (cgo/nocgo shared) ====================
// All enum types use int32 to match C's `int` type for 32/64 bit consistency.

// WindowType is a representation of GTK's GtkWindowType.
type WindowType = int32

const (
	WINDOW_TOPLEVEL WindowType = 0
	WINDOW_POPUP    WindowType = 1
)

// Gravity is a representation of GDK's GdkGravity.
type Gravity = int32

const (
	GDK_GRAVITY_NORTH_WEST Gravity = 1
	GDK_GRAVITY_NORTH      Gravity = 2
	GDK_GRAVITY_NORTH_EAST Gravity = 3
	GDK_GRAVITY_WEST       Gravity = 4
	GDK_GRAVITY_CENTER     Gravity = 5
	GDK_GRAVITY_EAST       Gravity = 6
	GDK_GRAVITY_SOUTH_WEST Gravity = 7
	GDK_GRAVITY_SOUTH      Gravity = 8
	GDK_GRAVITY_SOUTH_EAST Gravity = 9
	GDK_GRAVITY_STATIC     Gravity = 10
)

// WindowTypeHint is a representation of GDK's GdkWindowTypeHint.
type WindowTypeHint = int32

const (
	WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = 0
	WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = 1
	WINDOW_TYPE_HINT_MENU          WindowTypeHint = 2
	WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = 3
	WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = 4
	WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = 5
	WINDOW_TYPE_HINT_DOCK          WindowTypeHint = 6
	WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = 7
	WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = 8
	WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = 9
	WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = 10
	WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = 11
	WINDOW_TYPE_HINT_COMBO         WindowTypeHint = 12
	WINDOW_TYPE_HINT_DND           WindowTypeHint = 13
)

// ReliefStyle is a representation of GTK's GtkReliefStyle.
type ReliefStyle = int32

const (
	RELIEF_NORMAL ReliefStyle = 0
	RELIEF_HALF   ReliefStyle = 1
	RELIEF_NONE   ReliefStyle = 2
)

// PositionType is a representation of GTK's GtkPositionType.
type PositionType = int32

const (
	POS_LEFT   PositionType = 0
	POS_RIGHT  PositionType = 1
	POS_TOP    PositionType = 2
	POS_BOTTOM PositionType = 3
)

// Align is a representation of GTK's GtkAlign.
type Align = int32

const (
	ALIGN_FILL   Align = 0
	ALIGN_START  Align = 1
	ALIGN_END    Align = 2
	ALIGN_CENTER Align = 3
)

// StateFlags is a representation of GTK's GtkStateFlags.
type StateFlags = int32

const (
	STATE_FLAG_NORMAL       StateFlags = 0
	STATE_FLAG_ACTIVE       StateFlags = 1
	STATE_FLAG_PRELIGHT     StateFlags = 2
	STATE_FLAG_SELECTED     StateFlags = 4
	STATE_FLAG_INSENSITIVE  StateFlags = 8
	STATE_FLAG_INCONSISTENT StateFlags = 16
	STATE_FLAG_FOCUSED      StateFlags = 32
	STATE_FLAG_BACKDROP     StateFlags = 64
)

// IconSize is a representation of GTK's GtkIconSize.
type IconSize = int32

const (
	ICON_SIZE_INVALID       IconSize = 0
	ICON_SIZE_MENU          IconSize = 1
	ICON_SIZE_SMALL_TOOLBAR IconSize = 2
	ICON_SIZE_LARGE_TOOLBAR IconSize = 3
	ICON_SIZE_BUTTON        IconSize = 4
	ICON_SIZE_DND           IconSize = 5
	ICON_SIZE_DIALOG        IconSize = 6
)

// ImageType is a representation of GTK's GtkImageType.
type ImageType = int32

const (
	IMAGE_EMPTY     ImageType = 0
	IMAGE_PIXBUF    ImageType = 1
	IMAGE_STOCK     ImageType = 2
	IMAGE_ICON_SET  ImageType = 3
	IMAGE_ANIMATION ImageType = 4
	IMAGE_ICON_NAME ImageType = 5
	IMAGE_GICON     ImageType = 6
)

// EntryIconPosition is a representation of GTK's GtkEntryIconPosition.
type EntryIconPosition = int32

const (
	ENTRY_ICON_PRIMARY   EntryIconPosition = 0
	ENTRY_ICON_SECONDARY EntryIconPosition = 1
)

// Orientation is a representation of GTK's GtkOrientation.
type Orientation = int32

const (
	ORIENTATION_HORIZONTAL Orientation = 0
	ORIENTATION_VERTICAL   Orientation = 1
)

// PackType is a representation of GTK's GtkPackType.
type PackType = int32

const (
	PACK_START PackType = 0
	PACK_END   PackType = 1
)

// Justification is a representation of GTK's GtkJustification.
type Justification = int32

const (
	JUSTIFY_LEFT   Justification = 0
	JUSTIFY_RIGHT  Justification = 1
	JUSTIFY_CENTER Justification = 2
	JUSTIFY_FILL   Justification = 3
)

// SizeRequestMode is a representation of GTK's GtkSizeRequestMode.
type SizeRequestMode = int32

const (
	SIZE_REQUEST_HEIGHT_FOR_WIDTH SizeRequestMode = 0
	SIZE_REQUEST_WIDTH_FOR_HEIGHT SizeRequestMode = 1
	SIZE_REQUEST_CONSTANT_SIZE    SizeRequestMode = 2
)

// PolicyType is a representation of GTK's GtkPolicyType.
type PolicyType = int32

const (
	POLICY_ALWAYS    PolicyType = 0
	POLICY_AUTOMATIC PolicyType = 1
	POLICY_NEVER     PolicyType = 2
)

// CornerType is a representation of GTK's GtkCornerType.
type CornerType = int32

const (
	CORNER_TOP_LEFT     CornerType = 0
	CORNER_BOTTOM_LEFT  CornerType = 1
	CORNER_TOP_RIGHT    CornerType = 2
	CORNER_BOTTOM_RIGHT CornerType = 3
)

// ShadowType is a representation of GTK's GtkShadowType.
type ShadowType = int32

const (
	SHADOW_NONE       ShadowType = 0
	SHADOW_IN         ShadowType = 1
	SHADOW_OUT        ShadowType = 2
	SHADOW_ETCHED_IN  ShadowType = 3
	SHADOW_ETCHED_OUT ShadowType = 4
)

// SensitivityType is a representation of GTK's GtkSensitivityType.
type SensitivityType = int32

const (
	SENSITIVITY_AUTO SensitivityType = 0
	SENSITIVITY_ON   SensitivityType = 1
	SENSITIVITY_OFF  SensitivityType = 2
)

// StackTransitionType is a representation of GTK's GtkStackTransitionType.
type StackTransitionType = int32

const (
	STACK_TRANSITION_TYPE_NONE            StackTransitionType = 0
	STACK_TRANSITION_TYPE_CROSSFADE       StackTransitionType = 1
	STACK_TRANSITION_TYPE_SLIDE_RIGHT     StackTransitionType = 2
	STACK_TRANSITION_TYPE_SLIDE_LEFT      StackTransitionType = 3
	STACK_TRANSITION_TYPE_SLIDE_UP        StackTransitionType = 4
	STACK_TRANSITION_TYPE_SLIDE_DOWN      StackTransitionType = 5
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = 6
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN   StackTransitionType = 7
	STACK_TRANSITION_TYPE_OVER_UP         StackTransitionType = 8
	STACK_TRANSITION_TYPE_OVER_DOWN       StackTransitionType = 9
	STACK_TRANSITION_TYPE_OVER_LEFT       StackTransitionType = 10
	STACK_TRANSITION_TYPE_OVER_RIGHT      StackTransitionType = 11
	STACK_TRANSITION_TYPE_UNDER_UP        StackTransitionType = 12
	STACK_TRANSITION_TYPE_UNDER_DOWN      StackTransitionType = 13
	STACK_TRANSITION_TYPE_UNDER_LEFT      StackTransitionType = 14
	STACK_TRANSITION_TYPE_UNDER_RIGHT     StackTransitionType = 15
	STACK_TRANSITION_TYPE_OVER_UP_DOWN    StackTransitionType = 16
	STACK_TRANSITION_TYPE_OVER_DOWN_UP    StackTransitionType = 17
	STACK_TRANSITION_TYPE_ROTATE_LEFT     StackTransitionType = 18
	STACK_TRANSITION_TYPE_ROTATE_RIGHT    StackTransitionType = 19
	STACK_TRANSITION_TYPE_ROTATE_LEFT_RIGHT StackTransitionType = 20
)

// LevelBarMode is a representation of GTK's GtkLevelBarMode.
type LevelBarMode = int32

const (
	LEVEL_BAR_MODE_CONTINUOUS LevelBarMode = 0
	LEVEL_BAR_MODE_DISCRETE   LevelBarMode = 1
)

// PositionType is a representation of GTK's GtkPositionType (for Scale value-pos).
type ScaleValuePos = PositionType

// RevealerTransitionType is a representation of GTK's GtkRevealerTransitionType.
type RevealerTransitionType = int32

const (
	REVEALER_TRANSITION_TYPE_NONE     RevealerTransitionType = 0
	REVEALER_TRANSITION_TYPE_CROSSFADE RevealerTransitionType = 1
	REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType = 2
	REVEALER_TRANSITION_TYPE_SLIDE_LEFT RevealerTransitionType = 3
	REVEALER_TRANSITION_TYPE_SLIDE_UP RevealerTransitionType = 4
	REVEALER_TRANSITION_TYPE_SLIDE_DOWN RevealerTransitionType = 5
)

// ArrowType is a representation of GTK's GtkArrowType.
type ArrowType = int32

const (
	ARROW_UP    ArrowType = 0
	ARROW_DOWN  ArrowType = 1
	ARROW_LEFT  ArrowType = 2
	ARROW_RIGHT ArrowType = 3
	ARROW_NONE  ArrowType = 4
)

// DestDefaults is a representation of GTK's GtkDestDefaults.
type DestDefaults = int32

const (
	DEST_DEFAULT_MOTION    DestDefaults = 1
	DEST_DEFAULT_HIGHLIGHT DestDefaults = 2
	DEST_DEFAULT_DROP      DestDefaults = 4
	DEST_DEFAULT_ALL       DestDefaults = 7
)

// TargetFlags is a representation of GTK's GtkTargetFlags.
type TargetFlags = int32

const (
	TARGET_SAME_APP     TargetFlags = 1
	TARGET_SAME_WIDGET  TargetFlags = 2
	TARGET_OTHER_APP    TargetFlags = 4
	TARGET_OTHER_WIDGET TargetFlags = 8
)

// EllipsizeMode is a representation of Pango's PangoEllipsizeMode.
type EllipsizeMode = int32

const (
	ELLIPSIZE_NONE   EllipsizeMode = 0
	ELLIPSIZE_START  EllipsizeMode = 1
	ELLIPSIZE_MIDDLE EllipsizeMode = 2
	ELLIPSIZE_END    EllipsizeMode = 3
)

// EventMask is a representation of GDK's GdkEventMask.
type EventMask = int32

const (
	EXPOSURE_MASK            EventMask = 1 << 1
	POINTER_MOTION_MASK      EventMask = 1 << 2
	POINTER_MOTION_HINT_MASK EventMask = 1 << 3
	BUTTON_MOTION_MASK       EventMask = 1 << 4
	BUTTON1_MOTION_MASK      EventMask = 1 << 5
	BUTTON2_MOTION_MASK      EventMask = 1 << 6
	BUTTON3_MOTION_MASK      EventMask = 1 << 7
	BUTTON_PRESS_MASK        EventMask = 1 << 8
	BUTTON_RELEASE_MASK      EventMask = 1 << 9
	KEY_PRESS_MASK           EventMask = 1 << 10
	KEY_RELEASE_MASK         EventMask = 1 << 11
	ENTER_NOTIFY_MASK        EventMask = 1 << 12
	LEAVE_NOTIFY_MASK        EventMask = 1 << 13
	FOCUS_CHANGE_MASK        EventMask = 1 << 14
	STRUCTURE_MASK           EventMask = 1 << 15
	PROPERTY_CHANGE_MASK     EventMask = 1 << 16
	VISIBILITY_NOTIFY_MASK   EventMask = 1 << 17
	PROXIMITY_IN_MASK        EventMask = 1 << 18
	PROXIMITY_OUT_MASK       EventMask = 1 << 19
	SUBSTRUCTURE_MASK        EventMask = 1 << 20
	SCROLL_MASK              EventMask = 1 << 21
	TOUCH_MASK               EventMask = 1 << 22
	SMOOTH_SCROLL_MASK       EventMask = 1 << 23
	TOUCHPAD_GESTURE_MASK    EventMask = 1 << 24
	TABLET_PAD_MASK          EventMask = 1 << 25
	ALL_EVENTS_MASK          EventMask = 0x3FFFFFE
)

// ModifierType is a representation of GDK's GdkModifierType.
type ModifierType uint

const (
	SHIFT_MASK    ModifierType = 1 << 0
	LOCK_MASK     ModifierType = 1 << 1
	CONTROL_MASK  ModifierType = 1 << 2
	MOD1_MASK     ModifierType = 1 << 3
	MOD2_MASK     ModifierType = 1 << 4
	MOD3_MASK     ModifierType = 1 << 5
	MOD4_MASK     ModifierType = 1 << 6
	MOD5_MASK     ModifierType = 1 << 7
	BUTTON1_MASK  ModifierType = 1 << 8
	BUTTON2_MASK  ModifierType = 1 << 9
	BUTTON3_MASK  ModifierType = 1 << 10
	BUTTON4_MASK  ModifierType = 1 << 11
	BUTTON5_MASK  ModifierType = 1 << 12
	SUPER_MASK    ModifierType = 1 << 26
	HYPER_MASK    ModifierType = 1 << 27
	META_MASK     ModifierType = 1 << 28
	RELEASE_MASK  ModifierType = 1 << 30
	MODIFIER_MASK ModifierType = 0x5c001fff
)

// VisualType is a representation of GDK's GdkVisualType.
type VisualType int

const (
	VISUAL_STATIC_GRAY  VisualType = 0
	VISUAL_GRAYSCALE    VisualType = 1
	VISUAL_STATIC_COLOR VisualType = 2
	VISUAL_PSEUDO_COLOR VisualType = 3
	VISUAL_TRUE_COLOR   VisualType = 4
	VISUAL_DIRECT_COLOR VisualType = 5
)

// ISUAL_PSEUDO_COLOR is a deprecated alias for VISUAL_PSEUDO_COLOR (typo in original cgo code).
const ISUAL_PSEUDO_COLOR = VISUAL_PSEUDO_COLOR

// CURRENT_TIME is the GDK_CURRENT_TIME constant.
const CURRENT_TIME = 0

// FillRule is a representation of Cairo's cairo_fill_rule_t.
type FillRule int

const (
	FILL_RULE_WINDING  FillRule = 0
	FILL_RULE_EVEN_ODD FillRule = 1
)

// LineCap is a representation of Cairo's cairo_line_cap_t.
type LineCap int

const (
	LINE_CAP_BUTT   LineCap = 0
	LINE_CAP_ROUND  LineCap = 1
	LINE_CAP_SQUARE LineCap = 2
)

// LineJoin is a representation of Cairo's cairo_line_join_t.
type LineJoin int

const (
	LINE_JOIN_MITER LineJoin = 0
	LINE_JOIN_ROUND LineJoin = 1
	LINE_JOIN_BEVEL LineJoin = 2
)

// Operator is a representation of Cairo's cairo_operator_t.
type Operator int

const (
	OPERATOR_CLEAR          Operator = 0
	OPERATOR_SOURCE         Operator = 1
	OPERATOR_OVER           Operator = 2
	OPERATOR_IN             Operator = 3
	OPERATOR_OUT            Operator = 4
	OPERATOR_ATOP           Operator = 5
	OPERATOR_DEST           Operator = 6
	OPERATOR_DEST_OVER      Operator = 7
	OPERATOR_DEST_IN        Operator = 8
	OPERATOR_DEST_OUT       Operator = 9
	OPERATOR_DEST_ATOP      Operator = 10
	OPERATOR_XOR            Operator = 11
	OPERATOR_ADD            Operator = 12
	OPERATOR_SATURATE       Operator = 13
	OPERATOR_MULTIPLY       Operator = 14
	OPERATOR_SCREEN         Operator = 15
	OPERATOR_OVERLAY        Operator = 16
	OPERATOR_DARKEN         Operator = 17
	OPERATOR_LIGHTEN        Operator = 18
	OPERATOR_COLOR_DODGE    Operator = 19
	OPERATOR_COLOR_BURN     Operator = 20
	OPERATOR_HARD_LIGHT     Operator = 21
	OPERATOR_SOFT_LIGHT     Operator = 22
	OPERATOR_DIFFERENCE     Operator = 23
	OPERATOR_EXCLUSION      Operator = 24
	OPERATOR_HSL_HUE        Operator = 25
	OPERATOR_HSL_SATURATION Operator = 26
	OPERATOR_HSL_COLOR      Operator = 27
	OPERATOR_HSL_LUMINOSITY Operator = 28
)

// Priority is the enumerated type for GLib priority event sources.
type Priority int

const (
	PRIORITY_HIGH         Priority = -100
	PRIORITY_DEFAULT      Priority = 0
	PRIORITY_HIGH_IDLE    Priority = 100
	PRIORITY_DEFAULT_IDLE Priority = 200
	PRIORITY_LOW          Priority = 300
)

// SourceHandle is a handle returned by IdleAdd/IdleAddPriority.
type SourceHandle uint

// CompareDataFunc is a representation of GCompareDataFunc.
type CompareDataFunc func(a, b uintptr) int

// Type is a representation of GLib's GType.
type Type uint

const (
	TYPE_INVALID   Type = 0
	TYPE_NONE      Type = 4
	TYPE_INTERFACE Type = 8
	TYPE_CHAR      Type = 12
	TYPE_UCHAR     Type = 16
	TYPE_BOOLEAN   Type = 20
	TYPE_INT       Type = 24
	TYPE_UINT      Type = 28
	TYPE_LONG      Type = 32
	TYPE_ULONG     Type = 36
	TYPE_INT64     Type = 40
	TYPE_UINT64    Type = 44
	TYPE_ENUM      Type = 48
	TYPE_FLAGS     Type = 52
	TYPE_FLOAT     Type = 56
	TYPE_DOUBLE    Type = 60
	TYPE_STRING    Type = 64
	TYPE_POINTER   Type = 68
	TYPE_BOXED     Type = 72
	TYPE_PARAM     Type = 76
	TYPE_OBJECT    Type = 80
	TYPE_VARIANT   Type = 84
)

// DialogFlags is a representation of GTK's GtkDialogFlags.
type DialogFlags int

const (
	DIALOG_MODAL               DialogFlags = 1
	DIALOG_DESTROY_WITH_PARENT DialogFlags = 2
)

// MessageType is a representation of GTK's GtkMessageType.
type MessageType int

const (
	MESSAGE_INFO     MessageType = 0
	MESSAGE_WARNING  MessageType = 1
	MESSAGE_QUESTION MessageType = 2
	MESSAGE_ERROR    MessageType = 3
	MESSAGE_OTHER    MessageType = 4
)

// ButtonsType is a representation of GTK's GtkButtonsType.
type ButtonsType int

const (
	BUTTONS_NONE      ButtonsType = 0
	BUTTONS_OK        ButtonsType = 1
	BUTTONS_CLOSE     ButtonsType = 2
	BUTTONS_CANCEL    ButtonsType = 3
	BUTTONS_YES_NO    ButtonsType = 4
	BUTTONS_OK_CANCEL ButtonsType = 5
)

// WrapMode is a representation of GTK's GtkWrapMode.
type WrapMode int32

const (
	WRAP_NONE      WrapMode = 0
	WRAP_CHAR      WrapMode = 1
	WRAP_WORD      WrapMode = 2
	WRAP_WORD_CHAR WrapMode = 3
)

// TextWindowType is a representation of GTK's GtkTextWindowType.
type TextWindowType int32

const (
	TEXT_WINDOW_PRIVATE TextWindowType = 0
	TEXT_WINDOW_WIDGET  TextWindowType = 1
	TEXT_WINDOW_TEXT    TextWindowType = 2
	TEXT_WINDOW_LEFT    TextWindowType = 3
	TEXT_WINDOW_RIGHT   TextWindowType = 4
	TEXT_WINDOW_TOP     TextWindowType = 5
	TEXT_WINDOW_BOTTOM  TextWindowType = 6
)

// SelectionMode is a representation of GTK's GtkSelectionMode.
type SelectionMode int32

const (
	SELECTION_NONE     SelectionMode = 0
	SELECTION_SINGLE   SelectionMode = 1
	SELECTION_BROWSE   SelectionMode = 2
	SELECTION_MULTIPLE SelectionMode = 3
)

// FileChooserAction is a representation of GTK's GtkFileChooserAction.
type FileChooserAction int32

const (
	FILE_CHOOSER_ACTION_OPEN          FileChooserAction = 0
	FILE_CHOOSER_ACTION_SAVE          FileChooserAction = 1
	FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = 2
	FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = 3
)

// ResponseType is a representation of GTK's GtkResponseType.
type ResponseType int32

const (
	RESPONSE_NONE         ResponseType = -1
	RESPONSE_REJECT       ResponseType = -2
	RESPONSE_ACCEPT       ResponseType = -3
	RESPONSE_DELETE_EVENT ResponseType = -4
	RESPONSE_OK           ResponseType = -5
	RESPONSE_CANCEL       ResponseType = -6
	RESPONSE_YES          ResponseType = -7
	RESPONSE_NO           ResponseType = -8
	RESPONSE_CLOSE        ResponseType = -9
	RESPONSE_HELP         ResponseType = -11
)

// TreeViewColumnSizing is a representation of GTK's GtkTreeViewColumnSizing.
type TreeViewColumnSizing = int32

const (
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = 0
	TREE_VIEW_COLUMN_AUTOSIZE  TreeViewColumnSizing = 1
	TREE_VIEW_COLUMN_FIXED     TreeViewColumnSizing = 2
)
