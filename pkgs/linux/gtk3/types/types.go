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
