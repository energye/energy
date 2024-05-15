//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IChangeLink Parent: IObject
type IChangeLink interface {
	IObject
	Sender() ICustomImageList                                      // property
	SetSender(AValue ICustomImageList)                             // property
	Change()                                                       // procedure
	SetOnChange(fn TNotifyEvent)                                   // property event
	SetOnDestroyResolutionHandle(fn TDestroyResolutionHandleEvent) // property event
}

// TChangeLink Parent: TObject
type TChangeLink struct {
	TObject
	changePtr                  uintptr
	destroyResolutionHandlePtr uintptr
}

func NewChangeLink() IChangeLink {
	r1 := LCL().SysCallN(562)
	return AsChangeLink(r1)
}

func (m *TChangeLink) Sender() ICustomImageList {
	r1 := LCL().SysCallN(563, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TChangeLink) SetSender(AValue ICustomImageList) {
	LCL().SysCallN(563, 1, m.Instance(), GetObjectUintptr(AValue))
}

func ChangeLinkClass() TClass {
	ret := LCL().SysCallN(561)
	return TClass(ret)
}

func (m *TChangeLink) Change() {
	LCL().SysCallN(560, m.Instance())
}

func (m *TChangeLink) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(564, m.Instance(), m.changePtr)
}

func (m *TChangeLink) SetOnDestroyResolutionHandle(fn TDestroyResolutionHandleEvent) {
	if m.destroyResolutionHandlePtr != 0 {
		RemoveEventElement(m.destroyResolutionHandlePtr)
	}
	m.destroyResolutionHandlePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(565, m.Instance(), m.destroyResolutionHandlePtr)
}
