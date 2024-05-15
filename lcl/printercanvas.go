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

// IPrinterCanvas Parent: ICanvas
type IPrinterCanvas interface {
	ICanvas
	Printer() IPrinter                         // property
	Title() string                             // property
	SetTitle(AValue string)                    // property
	PageHeight() int32                         // property
	PageWidth() int32                          // property
	PaperWidth() int32                         // property
	SetPaperWidth(AValue int32)                // property
	PaperHeight() int32                        // property
	SetPaperHeight(AValue int32)               // property
	PageNumber() int32                         // property
	TopMargin() int32                          // property
	SetTopMargin(AValue int32)                 // property
	LeftMargin() int32                         // property
	SetLeftMargin(AValue int32)                // property
	BottomMargin() int32                       // property
	SetBottomMargin(AValue int32)              // property
	RightMargin() int32                        // property
	SetRightMargin(AValue int32)               // property
	Orientation() TPrinterOrientation          // property
	SetOrientation(AValue TPrinterOrientation) // property
	XDPI() int32                               // property
	SetXDPI(AValue int32)                      // property
	YDPI() int32                               // property
	SetYDPI(AValue int32)                      // property
	BeginDoc()                                 // procedure
	NewPage()                                  // procedure
	BeginPage()                                // procedure
	EndPage()                                  // procedure
	EndDoc()                                   // procedure
}

// TPrinterCanvas Parent: TCanvas
type TPrinterCanvas struct {
	TCanvas
}

func NewPrinterCanvas(APrinter IPrinter) IPrinterCanvas {
	r1 := LCL().SysCallN(4586, GetObjectUintptr(APrinter))
	return AsPrinterCanvas(r1)
}

func (m *TPrinterCanvas) Printer() IPrinter {
	r1 := LCL().SysCallN(4597, m.Instance())
	return AsPrinter(r1)
}

func (m *TPrinterCanvas) Title() string {
	r1 := LCL().SysCallN(4599, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinterCanvas) SetTitle(AValue string) {
	LCL().SysCallN(4599, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinterCanvas) PageHeight() int32 {
	r1 := LCL().SysCallN(4592, m.Instance())
	return int32(r1)
}

func (m *TPrinterCanvas) PageWidth() int32 {
	r1 := LCL().SysCallN(4594, m.Instance())
	return int32(r1)
}

func (m *TPrinterCanvas) PaperWidth() int32 {
	r1 := LCL().SysCallN(4596, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetPaperWidth(AValue int32) {
	LCL().SysCallN(4596, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) PaperHeight() int32 {
	r1 := LCL().SysCallN(4595, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetPaperHeight(AValue int32) {
	LCL().SysCallN(4595, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) PageNumber() int32 {
	r1 := LCL().SysCallN(4593, m.Instance())
	return int32(r1)
}

func (m *TPrinterCanvas) TopMargin() int32 {
	r1 := LCL().SysCallN(4600, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetTopMargin(AValue int32) {
	LCL().SysCallN(4600, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) LeftMargin() int32 {
	r1 := LCL().SysCallN(4589, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetLeftMargin(AValue int32) {
	LCL().SysCallN(4589, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) BottomMargin() int32 {
	r1 := LCL().SysCallN(4584, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetBottomMargin(AValue int32) {
	LCL().SysCallN(4584, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) RightMargin() int32 {
	r1 := LCL().SysCallN(4598, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetRightMargin(AValue int32) {
	LCL().SysCallN(4598, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) Orientation() TPrinterOrientation {
	r1 := LCL().SysCallN(4591, 0, m.Instance(), 0)
	return TPrinterOrientation(r1)
}

func (m *TPrinterCanvas) SetOrientation(AValue TPrinterOrientation) {
	LCL().SysCallN(4591, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) XDPI() int32 {
	r1 := LCL().SysCallN(4601, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetXDPI(AValue int32) {
	LCL().SysCallN(4601, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) YDPI() int32 {
	r1 := LCL().SysCallN(4602, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetYDPI(AValue int32) {
	LCL().SysCallN(4602, 1, m.Instance(), uintptr(AValue))
}

func PrinterCanvasClass() TClass {
	ret := LCL().SysCallN(4585)
	return TClass(ret)
}

func (m *TPrinterCanvas) BeginDoc() {
	LCL().SysCallN(4582, m.Instance())
}

func (m *TPrinterCanvas) NewPage() {
	LCL().SysCallN(4590, m.Instance())
}

func (m *TPrinterCanvas) BeginPage() {
	LCL().SysCallN(4583, m.Instance())
}

func (m *TPrinterCanvas) EndPage() {
	LCL().SysCallN(4588, m.Instance())
}

func (m *TPrinterCanvas) EndDoc() {
	LCL().SysCallN(4587, m.Instance())
}
