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
	"unsafe"
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
	r1 := LCL().SysCallN(3917, GetObjectUintptr(AOwner))
	return AsPopupMenu(r1)
}

func (m *TPopupMenu) PopupComponent() IComponent {
	r1 := LCL().SysCallN(3921, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TPopupMenu) SetPopupComponent(AValue IComponent) {
	LCL().SysCallN(3921, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPopupMenu) PopupPoint() (resultPoint TPoint) {
	LCL().SysCallN(3922, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TPopupMenu) Alignment() TPopupAlignment {
	r1 := LCL().SysCallN(3913, 0, m.Instance(), 0)
	return TPopupAlignment(r1)
}

func (m *TPopupMenu) SetAlignment(AValue TPopupAlignment) {
	LCL().SysCallN(3913, 1, m.Instance(), uintptr(AValue))
}

func (m *TPopupMenu) AutoPopup() bool {
	r1 := LCL().SysCallN(3914, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPopupMenu) SetAutoPopup(AValue bool) {
	LCL().SysCallN(3914, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPopupMenu) HelpContext() THelpContext {
	r1 := LCL().SysCallN(3918, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TPopupMenu) SetHelpContext(AValue THelpContext) {
	LCL().SysCallN(3918, 1, m.Instance(), uintptr(AValue))
}

func (m *TPopupMenu) TrackButton() TTrackButton {
	r1 := LCL().SysCallN(3925, 0, m.Instance(), 0)
	return TTrackButton(r1)
}

func (m *TPopupMenu) SetTrackButton(AValue TTrackButton) {
	LCL().SysCallN(3925, 1, m.Instance(), uintptr(AValue))
}

func PopupMenuClass() TClass {
	ret := LCL().SysCallN(3915)
	return TClass(ret)
}

func (m *TPopupMenu) PopUp() {
	LCL().SysCallN(3919, m.Instance())
}

func (m *TPopupMenu) PopUp1(X, Y int32) {
	LCL().SysCallN(3920, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TPopupMenu) Close() {
	LCL().SysCallN(3916, m.Instance())
}

func (m *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
	if m.popupPtr != 0 {
		RemoveEventElement(m.popupPtr)
	}
	m.popupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3924, m.Instance(), m.popupPtr)
}

func (m *TPopupMenu) SetOnClose(fn TNotifyEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3923, m.Instance(), m.closePtr)
}
