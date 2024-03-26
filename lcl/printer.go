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

// IPrinter Parent: IObject
type IPrinter interface {
	IObject
	PrinterIndex() int32                         // property
	SetPrinterIndex(AValue int32)                // property
	PrinterName() string                         // property
	PaperSize() IPaperSize                       // property
	Orientation() TPrinterOrientation            // property
	SetOrientation(AValue TPrinterOrientation)   // property
	PrinterState() TPrinterState                 // property
	Copies() int32                               // property
	SetCopies(AValue int32)                      // property
	Printers() IStrings                          // property
	FileName() string                            // property
	SetFileName(AValue string)                   // property
	Fonts() IStrings                             // property
	Canvas() ICanvas                             // property
	CanvasClass() uintptr                        // property
	SetCanvasClass(AValue uintptr)               // property
	PageHeight() int32                           // property
	PageWidth() int32                            // property
	PageNumber() int32                           // property
	Aborted() bool                               // property
	Printing() bool                              // property
	Title() string                               // property
	SetTitle(AValue string)                      // property
	PrinterType() TPrinterType                   // property
	CanPrint() bool                              // property
	CanRenderCopies() bool                       // property
	XDPI() int32                                 // property
	YDPI() int32                                 // property
	RawMode() bool                               // property
	SetRawMode(AValue bool)                      // property
	DefaultBinName() string                      // property
	BinName() string                             // property
	SetBinName(AValue string)                    // property
	SupportedBins() IStrings                     // property
	Write(Buffer []byte, OutWritten *int32) bool // function
	Write1(s string) bool                        // function
	Abort()                                      // procedure
	BeginDoc()                                   // procedure
	EndDoc()                                     // procedure
	NewPage()                                    // procedure
	BeginPage()                                  // procedure
	EndPage()                                    // procedure
	Refresh()                                    // procedure
	SetPrinter(aName string)                     // procedure
	RestoreDefaultBin()                          // procedure
}

// TPrinter Parent: TObject
type TPrinter struct {
	TObject
}

func NewPrinter() IPrinter {
	r1 := LCL().SysCallN(3976)
	return AsPrinter(r1)
}

func (m *TPrinter) PrinterIndex() int32 {
	r1 := LCL().SysCallN(3988, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetPrinterIndex(AValue int32) {
	LCL().SysCallN(3988, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterName() string {
	r1 := LCL().SysCallN(3989, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) PaperSize() IPaperSize {
	r1 := LCL().SysCallN(3987, m.Instance())
	return AsPaperSize(r1)
}

func (m *TPrinter) Orientation() TPrinterOrientation {
	r1 := LCL().SysCallN(3983, 0, m.Instance(), 0)
	return TPrinterOrientation(r1)
}

func (m *TPrinter) SetOrientation(AValue TPrinterOrientation) {
	LCL().SysCallN(3983, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterState() TPrinterState {
	r1 := LCL().SysCallN(3990, m.Instance())
	return TPrinterState(r1)
}

func (m *TPrinter) Copies() int32 {
	r1 := LCL().SysCallN(3975, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetCopies(AValue int32) {
	LCL().SysCallN(3975, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) Printers() IStrings {
	r1 := LCL().SysCallN(3992, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) FileName() string {
	r1 := LCL().SysCallN(3980, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetFileName(AValue string) {
	LCL().SysCallN(3980, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) Fonts() IStrings {
	r1 := LCL().SysCallN(3981, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Canvas() ICanvas {
	r1 := LCL().SysCallN(3972, m.Instance())
	return AsCanvas(r1)
}

func (m *TPrinter) CanvasClass() uintptr {
	r1 := LCL().SysCallN(3973, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TPrinter) SetCanvasClass(AValue uintptr) {
	LCL().SysCallN(3973, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PageHeight() int32 {
	r1 := LCL().SysCallN(3984, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageWidth() int32 {
	r1 := LCL().SysCallN(3986, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageNumber() int32 {
	r1 := LCL().SysCallN(3985, m.Instance())
	return int32(r1)
}

func (m *TPrinter) Aborted() bool {
	r1 := LCL().SysCallN(3966, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Printing() bool {
	r1 := LCL().SysCallN(3993, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Title() string {
	r1 := LCL().SysCallN(3999, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetTitle(AValue string) {
	LCL().SysCallN(3999, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) PrinterType() TPrinterType {
	r1 := LCL().SysCallN(3991, m.Instance())
	return TPrinterType(r1)
}

func (m *TPrinter) CanPrint() bool {
	r1 := LCL().SysCallN(3970, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) CanRenderCopies() bool {
	r1 := LCL().SysCallN(3971, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) XDPI() int32 {
	r1 := LCL().SysCallN(4002, m.Instance())
	return int32(r1)
}

func (m *TPrinter) YDPI() int32 {
	r1 := LCL().SysCallN(4003, m.Instance())
	return int32(r1)
}

func (m *TPrinter) RawMode() bool {
	r1 := LCL().SysCallN(3994, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPrinter) SetRawMode(AValue bool) {
	LCL().SysCallN(3994, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPrinter) DefaultBinName() string {
	r1 := LCL().SysCallN(3977, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) BinName() string {
	r1 := LCL().SysCallN(3969, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetBinName(AValue string) {
	LCL().SysCallN(3969, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) SupportedBins() IStrings {
	r1 := LCL().SysCallN(3998, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Write(Buffer []byte, OutWritten *int32) bool {
	var resultWritten uintptr
	r1 := LCL().SysCallN(4000, m.Instance(), uintptr(unsafe.Pointer(&Buffer[0])), uintptr(len(Buffer)), uintptr(unsafe.Pointer(&resultWritten)))
	*OutWritten = int32(resultWritten)
	return GoBool(r1)
}

func (m *TPrinter) Write1(s string) bool {
	r1 := LCL().SysCallN(4001, m.Instance(), PascalStr(s))
	return GoBool(r1)
}

func PrinterClass() TClass {
	ret := LCL().SysCallN(3974)
	return TClass(ret)
}

func (m *TPrinter) Abort() {
	LCL().SysCallN(3965, m.Instance())
}

func (m *TPrinter) BeginDoc() {
	LCL().SysCallN(3967, m.Instance())
}

func (m *TPrinter) EndDoc() {
	LCL().SysCallN(3978, m.Instance())
}

func (m *TPrinter) NewPage() {
	LCL().SysCallN(3982, m.Instance())
}

func (m *TPrinter) BeginPage() {
	LCL().SysCallN(3968, m.Instance())
}

func (m *TPrinter) EndPage() {
	LCL().SysCallN(3979, m.Instance())
}

func (m *TPrinter) Refresh() {
	LCL().SysCallN(3995, m.Instance())
}

func (m *TPrinter) SetPrinter(aName string) {
	LCL().SysCallN(3997, m.Instance(), PascalStr(aName))
}

func (m *TPrinter) RestoreDefaultBin() {
	LCL().SysCallN(3996, m.Instance())
}
