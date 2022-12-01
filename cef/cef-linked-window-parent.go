package cef

import (
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TCEFLinkedWindowParent struct {
	BaseWinControl
}

func NewCEFLinkedWindowParent(owner lcl.IComponent) *TCEFLinkedWindowParent {
	m := new(TCEFLinkedWindowParent)
	m.procName = "CEFLinkedWindow"
	m.instance = _Create(m.procName, lcl.CheckPtr(owner))
	m.ptr = unsafe.Pointer(m.instance)
	return m
}

func (m *TCEFLinkedWindowParent) Handle() types.HWND {
	return _GetHandle(m.procName, m.instance)
}

func (m *TCEFLinkedWindowParent) UpdateSize() {
	_CEFLinkedWindow_UpdateSize(m.instance)
}

func (m *TCEFLinkedWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_LinkedWindowParent
}

func (m *TCEFLinkedWindowParent) SetChromium(chromium IChromium, tag int32) {
	_CEFLinkedWindow_SetChromium(m.instance, chromium, tag)
}

func (m *TCEFLinkedWindowParent) HandleAllocated() bool {
	return api.DBoolToGoBool(_HandleAllocated(m.procName, m.instance))
}

func (m *TCEFLinkedWindowParent) CreateHandle() {
	_CreateHandle(m.procName, m.instance)
}

func (m *TCEFLinkedWindowParent) DestroyChildWindow() bool {
	return api.DBoolToGoBool(_DestroyChildWindow(m.procName, m.instance))
}

func (m *TCEFLinkedWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	_CEFLinkedWindow_OnEnter(m.instance, fn)
}

func (m *TCEFLinkedWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	_CEFLinkedWindow_OnExit(m.instance, fn)
}

func (m *TCEFLinkedWindowParent) Free() {
	_Free(m.procName, m.instance)
}
