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

// IVTColors Parent: IPersistent
// class to collect all switchable colors into one place
type IVTColors interface {
	IPersistent
	BackGroundColor() TColor                        // property
	HeaderFontColor() TColor                        // property
	NodeFontColor() TColor                          // property
	BorderColor() TColor                            // property
	SetBorderColor(AValue TColor)                   // property
	DisabledColor() TColor                          // property
	SetDisabledColor(AValue TColor)                 // property
	DropMarkColor() TColor                          // property
	SetDropMarkColor(AValue TColor)                 // property
	DropTargetColor() TColor                        // property
	SetDropTargetColor(AValue TColor)               // property
	DropTargetBorderColor() TColor                  // property
	SetDropTargetBorderColor(AValue TColor)         // property
	FocusedSelectionColor() TColor                  // property
	SetFocusedSelectionColor(AValue TColor)         // property
	FocusedSelectionBorderColor() TColor            // property
	SetFocusedSelectionBorderColor(AValue TColor)   // property
	GridLineColor() TColor                          // property
	SetGridLineColor(AValue TColor)                 // property
	HeaderHotColor() TColor                         // property
	SetHeaderHotColor(AValue TColor)                // property
	HotColor() TColor                               // property
	SetHotColor(AValue TColor)                      // property
	SelectionRectangleBlendColor() TColor           // property
	SetSelectionRectangleBlendColor(AValue TColor)  // property
	SelectionRectangleBorderColor() TColor          // property
	SetSelectionRectangleBorderColor(AValue TColor) // property
	SelectionTextColor() TColor                     // property
	SetSelectionTextColor(AValue TColor)            // property
	TreeLineColor() TColor                          // property
	SetTreeLineColor(AValue TColor)                 // property
	UnfocusedColor() TColor                         // property
	SetUnfocusedColor(AValue TColor)                // property
	UnfocusedSelectionColor() TColor                // property
	SetUnfocusedSelectionColor(AValue TColor)       // property
	UnfocusedSelectionBorderColor() TColor          // property
	SetUnfocusedSelectionBorderColor(AValue TColor) // property
}

// TVTColors Parent: TPersistent
// class to collect all switchable colors into one place
type TVTColors struct {
	TPersistent
}

func NewVTColors(AOwner IBaseVirtualTree) IVTColors {
	r1 := LCL().SysCallN(5814, GetObjectUintptr(AOwner))
	return AsVTColors(r1)
}

func (m *TVTColors) BackGroundColor() TColor {
	r1 := LCL().SysCallN(5811, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) HeaderFontColor() TColor {
	r1 := LCL().SysCallN(5822, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) NodeFontColor() TColor {
	r1 := LCL().SysCallN(5825, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) BorderColor() TColor {
	r1 := LCL().SysCallN(5812, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetBorderColor(AValue TColor) {
	LCL().SysCallN(5812, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DisabledColor() TColor {
	r1 := LCL().SysCallN(5815, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDisabledColor(AValue TColor) {
	LCL().SysCallN(5815, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropMarkColor() TColor {
	r1 := LCL().SysCallN(5816, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropMarkColor(AValue TColor) {
	LCL().SysCallN(5816, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropTargetColor() TColor {
	r1 := LCL().SysCallN(5818, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropTargetColor(AValue TColor) {
	LCL().SysCallN(5818, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropTargetBorderColor() TColor {
	r1 := LCL().SysCallN(5817, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropTargetBorderColor(AValue TColor) {
	LCL().SysCallN(5817, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) FocusedSelectionColor() TColor {
	r1 := LCL().SysCallN(5820, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetFocusedSelectionColor(AValue TColor) {
	LCL().SysCallN(5820, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) FocusedSelectionBorderColor() TColor {
	r1 := LCL().SysCallN(5819, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetFocusedSelectionBorderColor(AValue TColor) {
	LCL().SysCallN(5819, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) GridLineColor() TColor {
	r1 := LCL().SysCallN(5821, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetGridLineColor(AValue TColor) {
	LCL().SysCallN(5821, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) HeaderHotColor() TColor {
	r1 := LCL().SysCallN(5823, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetHeaderHotColor(AValue TColor) {
	LCL().SysCallN(5823, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) HotColor() TColor {
	r1 := LCL().SysCallN(5824, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetHotColor(AValue TColor) {
	LCL().SysCallN(5824, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionRectangleBlendColor() TColor {
	r1 := LCL().SysCallN(5826, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionRectangleBlendColor(AValue TColor) {
	LCL().SysCallN(5826, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionRectangleBorderColor() TColor {
	r1 := LCL().SysCallN(5827, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionRectangleBorderColor(AValue TColor) {
	LCL().SysCallN(5827, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionTextColor() TColor {
	r1 := LCL().SysCallN(5828, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionTextColor(AValue TColor) {
	LCL().SysCallN(5828, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) TreeLineColor() TColor {
	r1 := LCL().SysCallN(5829, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetTreeLineColor(AValue TColor) {
	LCL().SysCallN(5829, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedColor() TColor {
	r1 := LCL().SysCallN(5830, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedColor(AValue TColor) {
	LCL().SysCallN(5830, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedSelectionColor() TColor {
	r1 := LCL().SysCallN(5832, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedSelectionColor(AValue TColor) {
	LCL().SysCallN(5832, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedSelectionBorderColor() TColor {
	r1 := LCL().SysCallN(5831, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedSelectionBorderColor(AValue TColor) {
	LCL().SysCallN(5831, 1, m.Instance(), uintptr(AValue))
}

func VTColorsClass() TClass {
	ret := LCL().SysCallN(5813)
	return TClass(ret)
}
