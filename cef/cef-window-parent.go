//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type ITCefWindow interface {
	lcl.IWinControl
	Type() TCefWindowHandleType
	SetChromium(chromium *TCEFChromium, tag int32)
	UpdateSize()
	HandleAllocated() bool
	CreateHandle()
	SetOnEnter(fn lcl.TNotifyEvent)
	SetOnExit(fn lcl.TNotifyEvent)
	DestroyChildWindow() bool
}

type TCEFWindowParent struct {
	BaseWinControl
}

type TCEFLinkedWindowParent struct {
	BaseWinControl
}

func NewCEFWindow(owner lcl.IComponent) ITCefWindow {
	if IsWindows() {
		return NewCEFWindowParent(owner)
	} else {
		return NewCEFLinkedWindowParent(owner)
	}
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

func (m *TCEFWindowParent) Type() TCefWindowHandleType {
	return Wht_WindowParent
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

func NewCEFLinkedWindowParent(owner lcl.IComponent) *TCEFLinkedWindowParent {
	m := new(TCEFLinkedWindowParent)
	m.procName = "CEFLinkedWindow"
	m.instance = Create(m.procName, lcl.CheckPtr(owner))
	m.ptr = unsafe.Pointer(m.instance)
	return m
}

func (m *TCEFLinkedWindowParent) Handle() types.HWND {
	return GetHandle(m.procName, m.instance)
}

func (m *TCEFLinkedWindowParent) UpdateSize() {
	_CEFLinkedWindow_UpdateSize(m.instance)
}

func (m *TCEFLinkedWindowParent) Type() TCefWindowHandleType {
	return Wht_LinkedWindowParent
}

func (m *TCEFLinkedWindowParent) SetChromium(chromium *TCEFChromium, tag int32) {
	_CEFLinkedWindow_SetChromium(m.instance, chromium, tag)
}

func (m *TCEFLinkedWindowParent) HandleAllocated() bool {
	return api.DBoolToGoBool(HandleAllocated(m.procName, m.instance))
}

func (m *TCEFLinkedWindowParent) CreateHandle() {
	CreateHandle(m.procName, m.instance)
}

func (m *TCEFLinkedWindowParent) DestroyChildWindow() bool {
	return api.DBoolToGoBool(DestroyChildWindow(m.procName, m.instance))
}

func (m *TCEFLinkedWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	_CEFLinkedWindow_OnEnter(m.instance, fn)
}

func (m *TCEFLinkedWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	_CEFLinkedWindow_OnExit(m.instance, fn)
}
