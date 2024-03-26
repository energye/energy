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

// IPageSetupDialog Parent: ICustomPrinterSetupDialog
type IPageSetupDialog interface {
	ICustomPrinterSetupDialog
	AttachTo() ICustomForm                     // property
	SetAttachTo(AValue ICustomForm)            // property
	PageWidth() int32                          // property
	SetPageWidth(AValue int32)                 // property
	PageHeight() int32                         // property
	SetPageHeight(AValue int32)                // property
	MarginLeft() int32                         // property
	SetMarginLeft(AValue int32)                // property
	MarginTop() int32                          // property
	SetMarginTop(AValue int32)                 // property
	MarginRight() int32                        // property
	SetMarginRight(AValue int32)               // property
	MarginBottom() int32                       // property
	SetMarginBottom(AValue int32)              // property
	MinMarginLeft() int32                      // property
	SetMinMarginLeft(AValue int32)             // property
	MinMarginTop() int32                       // property
	SetMinMarginTop(AValue int32)              // property
	MinMarginRight() int32                     // property
	SetMinMarginRight(AValue int32)            // property
	MinMarginBottom() int32                    // property
	SetMinMarginBottom(AValue int32)           // property
	Options() TPageSetupDialogOptions          // property
	SetOptions(AValue TPageSetupDialogOptions) // property
	Units() TPageMeasureUnits                  // property
	SetUnits(AValue TPageMeasureUnits)         // property
	SetOnDialogResult(fn TDialogResultEvent)   // property event
}

// TPageSetupDialog Parent: TCustomPrinterSetupDialog
type TPageSetupDialog struct {
	TCustomPrinterSetupDialog
	dialogResultPtr uintptr
}

func NewPageSetupDialog(TheOwner IComponent) IPageSetupDialog {
	r1 := LCL().SysCallN(3781, GetObjectUintptr(TheOwner))
	return AsPageSetupDialog(r1)
}

func (m *TPageSetupDialog) AttachTo() ICustomForm {
	r1 := LCL().SysCallN(3779, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPageSetupDialog) SetAttachTo(AValue ICustomForm) {
	LCL().SysCallN(3779, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPageSetupDialog) PageWidth() int32 {
	r1 := LCL().SysCallN(3792, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetPageWidth(AValue int32) {
	LCL().SysCallN(3792, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) PageHeight() int32 {
	r1 := LCL().SysCallN(3791, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetPageHeight(AValue int32) {
	LCL().SysCallN(3791, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginLeft() int32 {
	r1 := LCL().SysCallN(3783, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginLeft(AValue int32) {
	LCL().SysCallN(3783, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginTop() int32 {
	r1 := LCL().SysCallN(3785, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginTop(AValue int32) {
	LCL().SysCallN(3785, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginRight() int32 {
	r1 := LCL().SysCallN(3784, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginRight(AValue int32) {
	LCL().SysCallN(3784, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginBottom() int32 {
	r1 := LCL().SysCallN(3782, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginBottom(AValue int32) {
	LCL().SysCallN(3782, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginLeft() int32 {
	r1 := LCL().SysCallN(3787, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginLeft(AValue int32) {
	LCL().SysCallN(3787, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginTop() int32 {
	r1 := LCL().SysCallN(3789, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginTop(AValue int32) {
	LCL().SysCallN(3789, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginRight() int32 {
	r1 := LCL().SysCallN(3788, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginRight(AValue int32) {
	LCL().SysCallN(3788, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginBottom() int32 {
	r1 := LCL().SysCallN(3786, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginBottom(AValue int32) {
	LCL().SysCallN(3786, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) Options() TPageSetupDialogOptions {
	r1 := LCL().SysCallN(3790, 0, m.Instance(), 0)
	return TPageSetupDialogOptions(r1)
}

func (m *TPageSetupDialog) SetOptions(AValue TPageSetupDialogOptions) {
	LCL().SysCallN(3790, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) Units() TPageMeasureUnits {
	r1 := LCL().SysCallN(3794, 0, m.Instance(), 0)
	return TPageMeasureUnits(r1)
}

func (m *TPageSetupDialog) SetUnits(AValue TPageMeasureUnits) {
	LCL().SysCallN(3794, 1, m.Instance(), uintptr(AValue))
}

func PageSetupDialogClass() TClass {
	ret := LCL().SysCallN(3780)
	return TClass(ret)
}

func (m *TPageSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3793, m.Instance(), m.dialogResultPtr)
}
