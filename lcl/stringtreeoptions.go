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

// IStringTreeOptions Parent: ICustomStringTreeOptions
type IStringTreeOptions interface {
	ICustomStringTreeOptions
	AnimationOptions() TVTAnimationOptions          // property
	SetAnimationOptions(AValue TVTAnimationOptions) // property
	AutoOptions() TVTAutoOptions                    // property
	SetAutoOptions(AValue TVTAutoOptions)           // property
	ExportMode() TVTExportMode                      // property
	SetExportMode(AValue TVTExportMode)             // property
	MiscOptions() TVTMiscOptions                    // property
	SetMiscOptions(AValue TVTMiscOptions)           // property
	PaintOptions() TVTPaintOptions                  // property
	SetPaintOptions(AValue TVTPaintOptions)         // property
	SelectionOptions() TVTSelectionOptions          // property
	SetSelectionOptions(AValue TVTSelectionOptions) // property
	StringOptions() TVTStringOptions                // property
	SetStringOptions(AValue TVTStringOptions)       // property
}

// TStringTreeOptions Parent: TCustomStringTreeOptions
type TStringTreeOptions struct {
	TCustomStringTreeOptions
}

func NewStringTreeOptions(AOwner IBaseVirtualTree) IStringTreeOptions {
	r1 := LCL().SysCallN(5239, GetObjectUintptr(AOwner))
	return AsStringTreeOptions(r1)
}

func (m *TStringTreeOptions) AnimationOptions() TVTAnimationOptions {
	r1 := LCL().SysCallN(5236, 0, m.Instance(), 0)
	return TVTAnimationOptions(r1)
}

func (m *TStringTreeOptions) SetAnimationOptions(AValue TVTAnimationOptions) {
	LCL().SysCallN(5236, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) AutoOptions() TVTAutoOptions {
	r1 := LCL().SysCallN(5237, 0, m.Instance(), 0)
	return TVTAutoOptions(r1)
}

func (m *TStringTreeOptions) SetAutoOptions(AValue TVTAutoOptions) {
	LCL().SysCallN(5237, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) ExportMode() TVTExportMode {
	r1 := LCL().SysCallN(5240, 0, m.Instance(), 0)
	return TVTExportMode(r1)
}

func (m *TStringTreeOptions) SetExportMode(AValue TVTExportMode) {
	LCL().SysCallN(5240, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) MiscOptions() TVTMiscOptions {
	r1 := LCL().SysCallN(5241, 0, m.Instance(), 0)
	return TVTMiscOptions(r1)
}

func (m *TStringTreeOptions) SetMiscOptions(AValue TVTMiscOptions) {
	LCL().SysCallN(5241, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) PaintOptions() TVTPaintOptions {
	r1 := LCL().SysCallN(5242, 0, m.Instance(), 0)
	return TVTPaintOptions(r1)
}

func (m *TStringTreeOptions) SetPaintOptions(AValue TVTPaintOptions) {
	LCL().SysCallN(5242, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) SelectionOptions() TVTSelectionOptions {
	r1 := LCL().SysCallN(5243, 0, m.Instance(), 0)
	return TVTSelectionOptions(r1)
}

func (m *TStringTreeOptions) SetSelectionOptions(AValue TVTSelectionOptions) {
	LCL().SysCallN(5243, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) StringOptions() TVTStringOptions {
	r1 := LCL().SysCallN(5244, 0, m.Instance(), 0)
	return TVTStringOptions(r1)
}

func (m *TStringTreeOptions) SetStringOptions(AValue TVTStringOptions) {
	LCL().SysCallN(5244, 1, m.Instance(), uintptr(AValue))
}

func StringTreeOptionsClass() TClass {
	ret := LCL().SysCallN(5238)
	return TClass(ret)
}
