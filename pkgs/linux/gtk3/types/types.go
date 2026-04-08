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
	GDK_ACTION_DEFAULT DragAction = 1 << 0 // 1
	GDK_ACTION_COPY    DragAction = 1 << 1 // 2
	GDK_ACTION_MOVE    DragAction = 1 << 2 // 4
	GDK_ACTION_LINK    DragAction = 1 << 3 // 8
	GDK_ACTION_PRIVATE DragAction = 1 << 4 // 16
	GDK_ACTION_ASK     DragAction = 1 << 5 // 32
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
