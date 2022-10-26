//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TCEFWindowParent struct {
	BaseWinControl
}

func NewCEFWindowParent(owner lcl.IComponent) *TCEFWindowParent {
	m := new(TCEFWindowParent)
	m.procName = "CEFWindow"
	m.instance = Create(m.procName, lcl.CheckPtr(owner))
	m.ptr = unsafe.Pointer(m.instance)
	return m
}

func (m *TCEFWindowParent) Handle() types.HWND {
	return GetHandle(m.procName, m.instance)
}

func (m *TCEFWindowParent) UpdateSize() {
	_CEFWindow_UpdateSize(m.instance)
}

func (m *TCEFWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_WindowParent
}

func (m *TCEFWindowParent) SetChromium(chromium *TCEFChromium, tag int32) {
}

func (m *TCEFWindowParent) HandleAllocated() bool {
	return api.DBoolToGoBool(HandleAllocated(m.procName, m.instance))
}

func (m *TCEFWindowParent) CreateHandle() {
	CreateHandle(m.procName, m.instance)
}

func (m *TCEFWindowParent) DestroyChildWindow() bool {
	return api.DBoolToGoBool(DestroyChildWindow(m.procName, m.instance))
}

func (m *TCEFWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	_CEFWindow_OnEnter(m.instance, fn)
}

func (m *TCEFWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	_CEFWindow_OnExit(m.instance, fn)
}
