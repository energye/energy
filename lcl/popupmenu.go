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

// IPopupMenu Parent: IMenu
type IPopupMenu interface {
	IMenu
	PopupComponent() IComponent          // property
	SetPopupComponent(AValue IComponent) // property
	PopupPoint() (resultPoint TPoint)    // property
	Alignment() TPopupAlignment          // property
	SetAlignment(AValue TPopupAlignment) // property
	AutoPopup() bool                     // property
	SetAutoPopup(AValue bool)            // property
	HelpContext() THelpContext           // property
	SetHelpContext(AValue THelpContext)  // property
	TrackButton() TTrackButton           // property
	SetTrackButton(AValue TTrackButton)  // property
	PopUp()                              // procedure
	PopUp1(X, Y int32)                   // procedure
	Close()                              // procedure
	SetOnPopup(fn TNotifyEvent)          // property event
	SetOnClose(fn TNotifyEvent)          // property event
}

// TPopupMenu Parent: TMenu
type TPopupMenu struct {
	TMenu
	popupPtr uintptr
	closePtr uintptr
}

func NewPopupMenu(AOwner IComponent) IPopupMenu {
	r1 := LCL().SysCallN(4559, GetObjectUintptr(AOwner))
	return AsPopupMenu(r1)
}

func (m *TPopupMenu) PopupComponent() IComponent {
	r1 := LCL().SysCallN(4563, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TPopupMenu) SetPopupComponent(AValue IComponent) {
	LCL().SysCallN(4563, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPopupMenu) PopupPoint() (resultPoint TPoint) {
	LCL().SysCallN(4564, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TPopupMenu) Alignment() TPopupAlignment {
	r1 := LCL().SysCallN(4555, 0, m.Instance(), 0)
	return TPopupAlignment(r1)
}

func (m *TPopupMenu) SetAlignment(AValue TPopupAlignment) {
	LCL().SysCallN(4555, 1, m.Instance(), uintptr(AValue))
}

func (m *TPopupMenu) AutoPopup() bool {
	r1 := LCL().SysCallN(4556, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPopupMenu) SetAutoPopup(AValue bool) {
	LCL().SysCallN(4556, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPopupMenu) HelpContext() THelpContext {
	r1 := LCL().SysCallN(4560, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TPopupMenu) SetHelpContext(AValue THelpContext) {
	LCL().SysCallN(4560, 1, m.Instance(), uintptr(AValue))
}

func (m *TPopupMenu) TrackButton() TTrackButton {
	r1 := LCL().SysCallN(4567, 0, m.Instance(), 0)
	return TTrackButton(r1)
}

func (m *TPopupMenu) SetTrackButton(AValue TTrackButton) {
	LCL().SysCallN(4567, 1, m.Instance(), uintptr(AValue))
}

func PopupMenuClass() TClass {
	ret := LCL().SysCallN(4557)
	return TClass(ret)
}

func (m *TPopupMenu) PopUp() {
	LCL().SysCallN(4561, m.Instance())
}

func (m *TPopupMenu) PopUp1(X, Y int32) {
	LCL().SysCallN(4562, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TPopupMenu) Close() {
	LCL().SysCallN(4558, m.Instance())
}

func (m *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
	if m.popupPtr != 0 {
		RemoveEventElement(m.popupPtr)
	}
	m.popupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4566, m.Instance(), m.popupPtr)
}

func (m *TPopupMenu) SetOnClose(fn TNotifyEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4565, m.Instance(), m.closePtr)
}
