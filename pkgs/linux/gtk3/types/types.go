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
)

type GDouble = float64
type GBoolean = int32
type PGdkWindow = uintptr
type GInt8 = int8
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

// const

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
