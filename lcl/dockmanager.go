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

// IDockManager Is Abstract Class Parent: IPersistent
type IDockManager interface {
	IPersistent
	GetDockEdge(ADockObject IDragDockObject) bool                                  // function
	AutoFreeByControl() bool                                                       // function
	IsEnabledControl(Control IControl) bool                                        // function
	CanBeDoubleDocked() bool                                                       // function
	BeginUpdate()                                                                  // procedure
	EndUpdate()                                                                    // procedure
	GetControlBounds(Control IControl, OutControlBounds *TRect)                    // procedure Is Abstract
	InsertControl(ADockObject IDragDockObject)                                     // procedure
	InsertControl1(Control IControl, InsertAt TAlign, DropCtl IControl)            // procedure Is Abstract
	LoadFromStream(Stream IStream)                                                 // procedure Is Abstract
	PaintSite(DC HDC)                                                              // procedure
	MessageHandler(Sender IControl, Message *TLMessage)                            // procedure
	PositionDockRect(ADockObject IDragDockObject)                                  // procedure
	PositionDockRect1(Client, DropCtl IControl, DropAlign TAlign, DockRect *TRect) // procedure Is Abstract
	RemoveControl(Control IControl)                                                // procedure Is Abstract
	ResetBounds(Force bool)                                                        // procedure Is Abstract
	SaveToStream(Stream IStream)                                                   // procedure Is Abstract
	SetReplacingControl(Control IControl)                                          // procedure
}

// TDockManager Is Abstract Class Parent: TPersistent
type TDockManager struct {
	TPersistent
}

func (m *TDockManager) GetDockEdge(ADockObject IDragDockObject) bool {
	r1 := LCL().SysCallN(2630, m.Instance(), GetObjectUintptr(ADockObject))
	return GoBool(r1)
}

func (m *TDockManager) AutoFreeByControl() bool {
	r1 := LCL().SysCallN(2624, m.Instance())
	return GoBool(r1)
}

func (m *TDockManager) IsEnabledControl(Control IControl) bool {
	r1 := LCL().SysCallN(2633, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TDockManager) CanBeDoubleDocked() bool {
	r1 := LCL().SysCallN(2626, m.Instance())
	return GoBool(r1)
}

func DockManagerClass() TClass {
	ret := LCL().SysCallN(2627)
	return TClass(ret)
}

func (m *TDockManager) BeginUpdate() {
	LCL().SysCallN(2625, m.Instance())
}

func (m *TDockManager) EndUpdate() {
	LCL().SysCallN(2628, m.Instance())
}

func (m *TDockManager) GetControlBounds(Control IControl, OutControlBounds *TRect) {
	var result1 uintptr
	LCL().SysCallN(2629, m.Instance(), GetObjectUintptr(Control), uintptr(unsafePointer(&result1)))
	*OutControlBounds = *(*TRect)(getPointer(result1))
}

func (m *TDockManager) InsertControl(ADockObject IDragDockObject) {
	LCL().SysCallN(2631, m.Instance(), GetObjectUintptr(ADockObject))
}

func (m *TDockManager) InsertControl1(Control IControl, InsertAt TAlign, DropCtl IControl) {
	LCL().SysCallN(2632, m.Instance(), GetObjectUintptr(Control), uintptr(InsertAt), GetObjectUintptr(DropCtl))
}

func (m *TDockManager) LoadFromStream(Stream IStream) {
	LCL().SysCallN(2634, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TDockManager) PaintSite(DC HDC) {
	LCL().SysCallN(2636, m.Instance(), uintptr(DC))
}

func (m *TDockManager) MessageHandler(Sender IControl, Message *TLMessage) {
	var result1 uintptr
	LCL().SysCallN(2635, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
}

func (m *TDockManager) PositionDockRect(ADockObject IDragDockObject) {
	LCL().SysCallN(2637, m.Instance(), GetObjectUintptr(ADockObject))
}

func (m *TDockManager) PositionDockRect1(Client, DropCtl IControl, DropAlign TAlign, DockRect *TRect) {
	var result2 uintptr
	LCL().SysCallN(2638, m.Instance(), GetObjectUintptr(Client), GetObjectUintptr(DropCtl), uintptr(DropAlign), uintptr(unsafePointer(&result2)))
	*DockRect = *(*TRect)(getPointer(result2))
}

func (m *TDockManager) RemoveControl(Control IControl) {
	LCL().SysCallN(2639, m.Instance(), GetObjectUintptr(Control))
}

func (m *TDockManager) ResetBounds(Force bool) {
	LCL().SysCallN(2640, m.Instance(), PascalBool(Force))
}

func (m *TDockManager) SaveToStream(Stream IStream) {
	LCL().SysCallN(2641, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TDockManager) SetReplacingControl(Control IControl) {
	LCL().SysCallN(2642, m.Instance(), GetObjectUintptr(Control))
}
