//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefPrintSettings Parent: ICefBaseRefCounted
//
//	Interface representing print settings.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_settings_capi.h">CEF source file: /include/capi/cef_print_settings_capi.h (cef_print_settings_t))</a>
type ICefPrintSettings interface {
	ICefBaseRefCounted
	// SetPageRanges
	//  Set the page ranges.
	SetPageRanges(ranges TRangeArray)
	// GetPageRanges
	//  Retrieve the page ranges.
	GetPageRanges() TRangeArray
	// IsValid
	//  Returns true (1) if this object is valid. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsReadOnly
	//  Returns true (1) if the values of this object are read-only. Some APIs may expose read-only objects.
	IsReadOnly() bool // function
	// IsLandscape
	//  Returns true (1) if the orientation is landscape.
	IsLandscape() bool // function
	// GetDeviceName
	//  Get the device name.
	GetDeviceName() string // function
	// GetDpi
	//  Get the DPI (dots per inch).
	GetDpi() int32 // function
	// GetPageRangesCount
	//  Returns the number of page ranges that currently exist.
	GetPageRangesCount() NativeUInt // function
	// IsSelectionOnly
	//  Returns true (1) if only the selection will be printed.
	IsSelectionOnly() bool // function
	// WillCollate
	//  Returns true (1) if pages will be collated.
	WillCollate() bool // function
	// GetColorModel
	//  Get the color model.
	GetColorModel() TCefColorModel // function
	// GetCopies
	//  Get the number of copies.
	GetCopies() int32 // function
	// GetDuplexMode
	//  Get the duplex mode.
	GetDuplexMode() TCefDuplexMode // function
	// SetOrientation
	//  Set the page orientation.
	SetOrientation(landscape bool) // procedure
	// SetPrinterPrintableArea
	//  Set the printer printable area in device units. Some platforms already provide flipped area. Set |landscape_needs_flip| to false (0) on those platforms to avoid double flipping.
	SetPrinterPrintableArea(physicalSizeDeviceUnits *TCefSize, printableAreaDeviceUnits *TCefRect, landscapeNeedsFlip bool) // procedure
	// SetDeviceName
	//  Set the device name.
	SetDeviceName(name string) // procedure
	// SetDpi
	//  Set the DPI (dots per inch).
	SetDpi(dpi int32) // procedure
	// SetSelectionOnly
	//  Set whether only the selection will be printed.
	SetSelectionOnly(selectionOnly bool) // procedure
	// SetCollate
	//  Set whether pages will be collated.
	SetCollate(collate bool) // procedure
	// SetColorModel
	//  Set the color model.
	SetColorModel(model TCefColorModel) // procedure
	// SetCopies
	//  Set the number of copies.
	SetCopies(copies int32) // procedure
	// SetDuplexMode
	//  Set the duplex mode.
	SetDuplexMode(mode TCefDuplexMode) // procedure
}

// TCefPrintSettings Parent: TCefBaseRefCounted
//
//	Interface representing print settings.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_settings_capi.h">CEF source file: /include/capi/cef_print_settings_capi.h (cef_print_settings_t))</a>
type TCefPrintSettings struct {
	TCefBaseRefCounted
}

// PrintSettingsRef -> ICefPrintSettings
var PrintSettingsRef printSettings

// printSettings TCefPrintSettings Ref
type printSettings uintptr

func (m *printSettings) New() ICefPrintSettings {
	var resultCefPrintSettings uintptr
	CEF().SysCallN(1229, uintptr(unsafePointer(&resultCefPrintSettings)))
	return AsCefPrintSettings(resultCefPrintSettings)
}

func (m *printSettings) UnWrap(data uintptr) ICefPrintSettings {
	var resultCefPrintSettings uintptr
	CEF().SysCallN(1239, uintptr(data), uintptr(unsafePointer(&resultCefPrintSettings)))
	return AsCefPrintSettings(resultCefPrintSettings)
}

func (m *TCefPrintSettings) IsValid() bool {
	r1 := CEF().SysCallN(1228, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) IsReadOnly() bool {
	r1 := CEF().SysCallN(1226, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) IsLandscape() bool {
	r1 := CEF().SysCallN(1225, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) GetDeviceName() string {
	r1 := CEF().SysCallN(1221, m.Instance())
	return GoStr(r1)
}

func (m *TCefPrintSettings) GetDpi() int32 {
	r1 := CEF().SysCallN(1222, m.Instance())
	return int32(r1)
}

func (m *TCefPrintSettings) GetPageRangesCount() NativeUInt {
	r1 := CEF().SysCallN(1224, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefPrintSettings) IsSelectionOnly() bool {
	r1 := CEF().SysCallN(1227, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) WillCollate() bool {
	r1 := CEF().SysCallN(1240, m.Instance())
	return GoBool(r1)
}

func (m *TCefPrintSettings) GetColorModel() TCefColorModel {
	r1 := CEF().SysCallN(1219, m.Instance())
	return TCefColorModel(r1)
}

func (m *TCefPrintSettings) GetCopies() int32 {
	r1 := CEF().SysCallN(1220, m.Instance())
	return int32(r1)
}

func (m *TCefPrintSettings) GetDuplexMode() TCefDuplexMode {
	r1 := CEF().SysCallN(1223, m.Instance())
	return TCefDuplexMode(r1)
}

func (m *TCefPrintSettings) SetOrientation(landscape bool) {
	CEF().SysCallN(1236, m.Instance(), PascalBool(landscape))
}

func (m *TCefPrintSettings) SetPrinterPrintableArea(physicalSizeDeviceUnits *TCefSize, printableAreaDeviceUnits *TCefRect, landscapeNeedsFlip bool) {
	CEF().SysCallN(1237, m.Instance(), uintptr(unsafePointer(physicalSizeDeviceUnits)), uintptr(unsafePointer(printableAreaDeviceUnits)), PascalBool(landscapeNeedsFlip))
}

func (m *TCefPrintSettings) SetDeviceName(name string) {
	CEF().SysCallN(1233, m.Instance(), PascalStr(name))
}

func (m *TCefPrintSettings) SetDpi(dpi int32) {
	CEF().SysCallN(1234, m.Instance(), uintptr(dpi))
}

func (m *TCefPrintSettings) SetSelectionOnly(selectionOnly bool) {
	CEF().SysCallN(1238, m.Instance(), PascalBool(selectionOnly))
}

func (m *TCefPrintSettings) SetCollate(collate bool) {
	CEF().SysCallN(1230, m.Instance(), PascalBool(collate))
}

func (m *TCefPrintSettings) SetColorModel(model TCefColorModel) {
	CEF().SysCallN(1231, m.Instance(), uintptr(model))
}

func (m *TCefPrintSettings) SetCopies(copies int32) {
	CEF().SysCallN(1232, m.Instance(), uintptr(copies))
}

func (m *TCefPrintSettings) SetDuplexMode(mode TCefDuplexMode) {
	CEF().SysCallN(1235, m.Instance(), uintptr(mode))
}
