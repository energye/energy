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

// EventSignalName 信号名
type EventSignalName = string

const (
	EsnClicked               EventSignalName = "clicked"
	EsnChanged               EventSignalName = "changed"
	EsnActivate              EventSignalName = "activate"
	EsnKeyPressEvent         EventSignalName = "key-press-event"
	EsnKeyReleaseEvent       EventSignalName = "key-release-event"
	EsnButtonPressEvent      EventSignalName = "button-press-event"
	EsnEnterNotifyEvent      EventSignalName = "enter-notify-event"
	EsnLeaveNotifyEvent      EventSignalName = "leave-notify-event"
	EsnConfigureEvent        EventSignalName = "configure-event"
	EsnMapEvent              EventSignalName = "map"
	EsnDrawEvent             EventSignalName = "draw"
	EsnDragDataReceivedEvent EventSignalName = "drag-data-received"
	EsnDragDropEvent         EventSignalName = "drag-drop"
	EsnDragMotionEvent       EventSignalName = "drag-motion"
	EsnDragLeaveEvent        EventSignalName = "drag-leave"
	EsnDragDataDeleteEvent   EventSignalName = "drag-data-delete"
	EsnDragBeginEvent        EventSignalName = "drag-begin"
	EsnDragEndEvent          EventSignalName = "drag-end"
)

// 信号事件类型
// Note: 参数和返回值除了基础类型，其它所有都用指针(uintptr)表示

type TNotifyEvent func(sender PGtkWidget, userData GPointer)
type TTextChangedEvent func(sender PGtkWidget, userData GPointer)
type TTextCommitEvent func(sender PGtkWidget, userData GPointer)
type TTextKeyEvent func(sender PGtkWidget, event PEventKey, userData GPointer) bool
type TButtonPressEvent func(sender PGtkWidget, event PEventButton, userData GPointer) bool
type TLeaveEnterNotifyEvent func(sender PGtkWidget, event PEventCrossing, userData GPointer) bool
type TConfigureEvent func(sender PGtkWidget, event PEventConfigure, userData GPointer) bool
type TMapEvent func(sender PGtkWidget, userData GPointer)
type TDrawEvent func(sender PGtkWidget, cr PContext, userData GPointer) bool

// type drag event

type TDragDataReceivedEvent func(sender PGtkWidget, context PDragContext, x, y int, data PSelectionData, info uint, time uint, userData GPointer)
type TDragDropEvent func(sender PGtkWidget, context PDragContext, x, y int, time uint, userData GPointer) bool
type TDragMotionEvent func(sender PGtkWidget, context PDragContext, x, y int, time uint, userData GPointer) bool
type TDragLeaveEvent func(sender PGtkWidget, context PDragContext, time uint, userData GPointer)
type TDragDataDeleteOrBeginOrEndEvent func(sender PGtkWidget, context PDragContext, userData GPointer)
