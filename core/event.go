//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package core

type TLoadChange int32

const (
	LcStart TLoadChange = iota
	LcLoading
	LcFinish
)

type TDragType int32

const (
	DragTypeNo   TDragType = iota // file list
	DragTypeFile                  // file list
	DragTypeData                  // file bytes
)

type TDragData struct {
	Type      TDragType
	Data      []byte
	Filenames []string
}

type TOnProcessMessageEvent func(message string)
type TOnResourceRequestEvent func(url, path, method string, header map[string]string) (resource string, ok bool)
type TOnLoadChangeEvent func(url, title string, load TLoadChange)
type TOnContextMenuEvent func(contextMenu *TContextMenuItem)
type TOnContextMenuCommandEvent func(commandId int32, handle *bool)
type TOnPopupWindowEvent func(targetURL string) bool
type TOnEvaluateScriptCallbackEvent func(result string, err string)
type TOnDragEnterEvent func(type_ TDragType, x, y int32)
type TOnDragLeaveEvent func()
type TOnDragOverEvent func(data *TDragData, x, y int32)
type TOnThemeChange func(isDark bool)
