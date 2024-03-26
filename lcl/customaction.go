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

// ICustomAction Parent: IContainedAction
type ICustomAction interface {
	IContainedAction
	AutoCheck() bool                            // property
	SetAutoCheck(AValue bool)                   // property
	Caption() string                            // property
	SetCaption(AValue string)                   // property
	Checked() bool                              // property
	SetChecked(AValue bool)                     // property
	Grayed() bool                               // property
	SetGrayed(AValue bool)                      // property
	DisableIfNoHandler() bool                   // property
	SetDisableIfNoHandler(AValue bool)          // property
	Enabled() bool                              // property
	SetEnabled(AValue bool)                     // property
	GroupIndex() int32                          // property
	SetGroupIndex(AValue int32)                 // property
	HelpContext() THelpContext                  // property
	SetHelpContext(AValue THelpContext)         // property
	HelpKeyword() string                        // property
	SetHelpKeyword(AValue string)               // property
	HelpType() THelpType                        // property
	SetHelpType(AValue THelpType)               // property
	Hint() string                               // property
	SetHint(AValue string)                      // property
	ImageIndex() TImageIndex                    // property
	SetImageIndex(AValue TImageIndex)           // property
	SecondaryShortCuts() IShortCutList          // property
	SetSecondaryShortCuts(AValue IShortCutList) // property
	ShortCut() TShortCut                        // property
	SetShortCut(AValue TShortCut)               // property
	Visible() bool                              // property
	SetVisible(AValue bool)                     // property
	SetOnHint(fn THintEvent)                    // property event
}

// TCustomAction Parent: TContainedAction
type TCustomAction struct {
	TContainedAction
	hintPtr uintptr
}

func NewCustomAction(AOwner IComponent) ICustomAction {
	r1 := LCL().SysCallN(1067, GetObjectUintptr(AOwner))
	return AsCustomAction(r1)
}

func (m *TCustomAction) AutoCheck() bool {
	r1 := LCL().SysCallN(1063, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetAutoCheck(AValue bool) {
	LCL().SysCallN(1063, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) Caption() string {
	r1 := LCL().SysCallN(1064, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAction) SetCaption(AValue string) {
	LCL().SysCallN(1064, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAction) Checked() bool {
	r1 := LCL().SysCallN(1065, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetChecked(AValue bool) {
	LCL().SysCallN(1065, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) Grayed() bool {
	r1 := LCL().SysCallN(1070, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetGrayed(AValue bool) {
	LCL().SysCallN(1070, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) DisableIfNoHandler() bool {
	r1 := LCL().SysCallN(1068, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetDisableIfNoHandler(AValue bool) {
	LCL().SysCallN(1068, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) Enabled() bool {
	r1 := LCL().SysCallN(1069, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetEnabled(AValue bool) {
	LCL().SysCallN(1069, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) GroupIndex() int32 {
	r1 := LCL().SysCallN(1071, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAction) SetGroupIndex(AValue int32) {
	LCL().SysCallN(1071, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) HelpContext() THelpContext {
	r1 := LCL().SysCallN(1072, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TCustomAction) SetHelpContext(AValue THelpContext) {
	LCL().SysCallN(1072, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) HelpKeyword() string {
	r1 := LCL().SysCallN(1073, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAction) SetHelpKeyword(AValue string) {
	LCL().SysCallN(1073, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAction) HelpType() THelpType {
	r1 := LCL().SysCallN(1074, 0, m.Instance(), 0)
	return THelpType(r1)
}

func (m *TCustomAction) SetHelpType(AValue THelpType) {
	LCL().SysCallN(1074, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) Hint() string {
	r1 := LCL().SysCallN(1075, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAction) SetHint(AValue string) {
	LCL().SysCallN(1075, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAction) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(1076, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomAction) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(1076, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) SecondaryShortCuts() IShortCutList {
	r1 := LCL().SysCallN(1077, 0, m.Instance(), 0)
	return AsShortCutList(r1)
}

func (m *TCustomAction) SetSecondaryShortCuts(AValue IShortCutList) {
	LCL().SysCallN(1077, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomAction) ShortCut() TShortCut {
	r1 := LCL().SysCallN(1079, 0, m.Instance(), 0)
	return TShortCut(r1)
}

func (m *TCustomAction) SetShortCut(AValue TShortCut) {
	LCL().SysCallN(1079, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) Visible() bool {
	r1 := LCL().SysCallN(1080, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetVisible(AValue bool) {
	LCL().SysCallN(1080, 1, m.Instance(), PascalBool(AValue))
}

func CustomActionClass() TClass {
	ret := LCL().SysCallN(1066)
	return TClass(ret)
}

func (m *TCustomAction) SetOnHint(fn THintEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1078, m.Instance(), m.hintPtr)
}
