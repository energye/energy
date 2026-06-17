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

type TContextMenuKind int32

const (
	CmkCommand TContextMenuKind = iota
	CmkSub
	CmkSeparator
)

var gContextMenuCommandId int32 = 10000

func NextContextMenuCommandId() int32 {
	gContextMenuCommandId++
	return gContextMenuCommandId
}

// TContextMenuItem 右键菜单
type TContextMenuItem struct {
	Add   func(text string, kind TContextMenuKind) (*TContextMenuItem, int32)
	Clear func()
}
