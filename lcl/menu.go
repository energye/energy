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

// IMenu Parent: ILCLComponent
type IMenu interface {
	ILCLComponent
	Handle() HMENU                                             // property
	Parent() IComponent                                        // property
	SetParent(AValue IComponent)                               // property
	ShortcutHandled() bool                                     // property
	SetShortcutHandled(AValue bool)                            // property
	BidiMode() TBiDiMode                                       // property
	SetBidiMode(AValue TBiDiMode)                              // property
	ParentBidiMode() bool                                      // property
	SetParentBidiMode(AValue bool)                             // property
	Items() IMenuItem                                          // property
	Images() ICustomImageList                                  // property
	SetImages(AValue ICustomImageList)                         // property
	ImagesWidth() int32                                        // property
	SetImagesWidth(AValue int32)                               // property
	OwnerDraw() bool                                           // property
	SetOwnerDraw(AValue bool)                                  // property
	FindItem(AValue uint32, Kind TFindItemKind) IMenuItem      // function
	GetHelpContext(AValue uint32, ByCommand bool) THelpContext // function
	IsShortcut(Message *TLMKey) bool                           // function
	HandleAllocated() bool                                     // function
	IsRightToLeft() bool                                       // function
	UseRightToLeftAlignment() bool                             // function
	UseRightToLeftReading() bool                               // function
	DispatchCommand(ACommand Word) bool                        // function
	DestroyHandle()                                            // procedure
	HandleNeeded()                                             // procedure
	SetOnDrawItem(fn TMenuDrawItemEvent)                       // property event
	SetOnMeasureItem(fn TMenuMeasureItemEvent)                 // property event
}

// TMenu Parent: TLCLComponent
type TMenu struct {
	TLCLComponent
	drawItemPtr    uintptr
	measureItemPtr uintptr
}

func NewMenu(AOwner IComponent) IMenu {
	r1 := LCL().SysCallN(3669, GetObjectUintptr(AOwner))
	return AsMenu(r1)
}

func (m *TMenu) Handle() HMENU {
	r1 := LCL().SysCallN(3674, m.Instance())
	return HMENU(r1)
}

func (m *TMenu) Parent() IComponent {
	r1 := LCL().SysCallN(3683, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TMenu) SetParent(AValue IComponent) {
	LCL().SysCallN(3683, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenu) ShortcutHandled() bool {
	r1 := LCL().SysCallN(3687, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetShortcutHandled(AValue bool) {
	LCL().SysCallN(3687, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) BidiMode() TBiDiMode {
	r1 := LCL().SysCallN(3667, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TMenu) SetBidiMode(AValue TBiDiMode) {
	LCL().SysCallN(3667, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenu) ParentBidiMode() bool {
	r1 := LCL().SysCallN(3684, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetParentBidiMode(AValue bool) {
	LCL().SysCallN(3684, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) Items() IMenuItem {
	r1 := LCL().SysCallN(3681, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenu) Images() ICustomImageList {
	r1 := LCL().SysCallN(3677, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TMenu) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(3677, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenu) ImagesWidth() int32 {
	r1 := LCL().SysCallN(3678, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenu) SetImagesWidth(AValue int32) {
	LCL().SysCallN(3678, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenu) OwnerDraw() bool {
	r1 := LCL().SysCallN(3682, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetOwnerDraw(AValue bool) {
	LCL().SysCallN(3682, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) FindItem(AValue uint32, Kind TFindItemKind) IMenuItem {
	r1 := LCL().SysCallN(3672, m.Instance(), uintptr(AValue), uintptr(Kind))
	return AsMenuItem(r1)
}

func (m *TMenu) GetHelpContext(AValue uint32, ByCommand bool) THelpContext {
	r1 := LCL().SysCallN(3673, m.Instance(), uintptr(AValue), PascalBool(ByCommand))
	return THelpContext(r1)
}

func (m *TMenu) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(3680, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TMenu) HandleAllocated() bool {
	r1 := LCL().SysCallN(3675, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) IsRightToLeft() bool {
	r1 := LCL().SysCallN(3679, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) UseRightToLeftAlignment() bool {
	r1 := LCL().SysCallN(3688, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) UseRightToLeftReading() bool {
	r1 := LCL().SysCallN(3689, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) DispatchCommand(ACommand Word) bool {
	r1 := LCL().SysCallN(3671, m.Instance(), uintptr(ACommand))
	return GoBool(r1)
}

func MenuClass() TClass {
	ret := LCL().SysCallN(3668)
	return TClass(ret)
}

func (m *TMenu) DestroyHandle() {
	LCL().SysCallN(3670, m.Instance())
}

func (m *TMenu) HandleNeeded() {
	LCL().SysCallN(3676, m.Instance())
}

func (m *TMenu) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3685, m.Instance(), m.drawItemPtr)
}

func (m *TMenu) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3686, m.Instance(), m.measureItemPtr)
}
