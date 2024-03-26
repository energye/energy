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

// ICustomTaskDialog Parent: IComponent
type ICustomTaskDialog interface {
	IComponent
	Button() ITaskDialogButtonItem                    // property
	SetButton(AValue ITaskDialogButtonItem)           // property
	Buttons() ITaskDialogButtons                      // property
	SetButtons(AValue ITaskDialogButtons)             // property
	Caption() string                                  // property
	SetCaption(AValue string)                         // property
	CommonButtons() TTaskDialogCommonButtons          // property
	SetCommonButtons(AValue TTaskDialogCommonButtons) // property
	DefaultButton() TTaskDialogCommonButton           // property
	SetDefaultButton(AValue TTaskDialogCommonButton)  // property
	ExpandButtonCaption() string                      // property
	SetExpandButtonCaption(AValue string)             // property
	ExpandedText() string                             // property
	SetExpandedText(AValue string)                    // property
	Flags() TTaskDialogFlags                          // property
	SetFlags(AValue TTaskDialogFlags)                 // property
	FooterIcon() TTaskDialogIcon                      // property
	SetFooterIcon(AValue TTaskDialogIcon)             // property
	FooterText() string                               // property
	SetFooterText(AValue string)                      // property
	MainIcon() TTaskDialogIcon                        // property
	SetMainIcon(AValue TTaskDialogIcon)               // property
	ModalResult() TModalResult                        // property
	SetModalResult(AValue TModalResult)               // property
	RadioButton() ITaskDialogRadioButtonItem          // property
	RadioButtons() ITaskDialogButtons                 // property
	SetRadioButtons(AValue ITaskDialogButtons)        // property
	Text() string                                     // property
	SetText(AValue string)                            // property
	Title() string                                    // property
	SetTitle(AValue string)                           // property
	VerificationText() string                         // property
	SetVerificationText(AValue string)                // property
	Width() int32                                     // property
	SetWidth(AValue int32)                            // property
	Execute() bool                                    // function
	Execute1(ParentWnd HWND) bool                     // function
	SetOnButtonClicked(fn TTaskDlgClickEvent)         // property event
}

// TCustomTaskDialog Parent: TComponent
type TCustomTaskDialog struct {
	TComponent
	buttonClickedPtr uintptr
}

func NewCustomTaskDialog(AOwner IComponent) ICustomTaskDialog {
	r1 := LCL().SysCallN(2140, GetObjectUintptr(AOwner))
	return AsCustomTaskDialog(r1)
}

func (m *TCustomTaskDialog) Button() ITaskDialogButtonItem {
	r1 := LCL().SysCallN(2135, 0, m.Instance(), 0)
	return AsTaskDialogButtonItem(r1)
}

func (m *TCustomTaskDialog) SetButton(AValue ITaskDialogButtonItem) {
	LCL().SysCallN(2135, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Buttons() ITaskDialogButtons {
	r1 := LCL().SysCallN(2136, 0, m.Instance(), 0)
	return AsTaskDialogButtons(r1)
}

func (m *TCustomTaskDialog) SetButtons(AValue ITaskDialogButtons) {
	LCL().SysCallN(2136, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Caption() string {
	r1 := LCL().SysCallN(2137, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetCaption(AValue string) {
	LCL().SysCallN(2137, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) CommonButtons() TTaskDialogCommonButtons {
	r1 := LCL().SysCallN(2139, 0, m.Instance(), 0)
	return TTaskDialogCommonButtons(r1)
}

func (m *TCustomTaskDialog) SetCommonButtons(AValue TTaskDialogCommonButtons) {
	LCL().SysCallN(2139, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) DefaultButton() TTaskDialogCommonButton {
	r1 := LCL().SysCallN(2141, 0, m.Instance(), 0)
	return TTaskDialogCommonButton(r1)
}

func (m *TCustomTaskDialog) SetDefaultButton(AValue TTaskDialogCommonButton) {
	LCL().SysCallN(2141, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) ExpandButtonCaption() string {
	r1 := LCL().SysCallN(2144, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetExpandButtonCaption(AValue string) {
	LCL().SysCallN(2144, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) ExpandedText() string {
	r1 := LCL().SysCallN(2145, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetExpandedText(AValue string) {
	LCL().SysCallN(2145, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Flags() TTaskDialogFlags {
	r1 := LCL().SysCallN(2146, 0, m.Instance(), 0)
	return TTaskDialogFlags(r1)
}

func (m *TCustomTaskDialog) SetFlags(AValue TTaskDialogFlags) {
	LCL().SysCallN(2146, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) FooterIcon() TTaskDialogIcon {
	r1 := LCL().SysCallN(2147, 0, m.Instance(), 0)
	return TTaskDialogIcon(r1)
}

func (m *TCustomTaskDialog) SetFooterIcon(AValue TTaskDialogIcon) {
	LCL().SysCallN(2147, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) FooterText() string {
	r1 := LCL().SysCallN(2148, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetFooterText(AValue string) {
	LCL().SysCallN(2148, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) MainIcon() TTaskDialogIcon {
	r1 := LCL().SysCallN(2149, 0, m.Instance(), 0)
	return TTaskDialogIcon(r1)
}

func (m *TCustomTaskDialog) SetMainIcon(AValue TTaskDialogIcon) {
	LCL().SysCallN(2149, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) ModalResult() TModalResult {
	r1 := LCL().SysCallN(2150, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomTaskDialog) SetModalResult(AValue TModalResult) {
	LCL().SysCallN(2150, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) RadioButton() ITaskDialogRadioButtonItem {
	r1 := LCL().SysCallN(2151, m.Instance())
	return AsTaskDialogRadioButtonItem(r1)
}

func (m *TCustomTaskDialog) RadioButtons() ITaskDialogButtons {
	r1 := LCL().SysCallN(2152, 0, m.Instance(), 0)
	return AsTaskDialogButtons(r1)
}

func (m *TCustomTaskDialog) SetRadioButtons(AValue ITaskDialogButtons) {
	LCL().SysCallN(2152, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Text() string {
	r1 := LCL().SysCallN(2154, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetText(AValue string) {
	LCL().SysCallN(2154, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Title() string {
	r1 := LCL().SysCallN(2155, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetTitle(AValue string) {
	LCL().SysCallN(2155, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) VerificationText() string {
	r1 := LCL().SysCallN(2156, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetVerificationText(AValue string) {
	LCL().SysCallN(2156, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Width() int32 {
	r1 := LCL().SysCallN(2157, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTaskDialog) SetWidth(AValue int32) {
	LCL().SysCallN(2157, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) Execute() bool {
	r1 := LCL().SysCallN(2142, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTaskDialog) Execute1(ParentWnd HWND) bool {
	r1 := LCL().SysCallN(2143, m.Instance(), uintptr(ParentWnd))
	return GoBool(r1)
}

func CustomTaskDialogClass() TClass {
	ret := LCL().SysCallN(2138)
	return TClass(ret)
}

func (m *TCustomTaskDialog) SetOnButtonClicked(fn TTaskDlgClickEvent) {
	if m.buttonClickedPtr != 0 {
		RemoveEventElement(m.buttonClickedPtr)
	}
	m.buttonClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2153, m.Instance(), m.buttonClickedPtr)
}
