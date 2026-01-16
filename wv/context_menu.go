//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

type TContextMenuKind int32

const (
	CmkCommand TContextMenuKind = iota
	CmkSub
	CmkSeparator
)

var gContextMenuCommandId int32 = 10000

func nextContextMenuCommandId() int32 {
	gContextMenuCommandId++
	return gContextMenuCommandId
}

// TContextMenuItem 右键菜单
type TContextMenuItem struct {
	add   func(text string, kind TContextMenuKind) (*TContextMenuItem, int32)
	clear func()
}

// Add 向上下文菜单中添加一个新的菜单项
func (m *TContextMenuItem) Add(text string, kind TContextMenuKind) (*TContextMenuItem, int32) {
	return m.add(text, kind)
}

func (m *TContextMenuItem) Clear() {
	m.clear()
}
