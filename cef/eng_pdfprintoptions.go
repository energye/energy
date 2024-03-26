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

// IPDFPrintOptions Parent: IObject
//
//	The TPDFPrintOptions properties are used to fill the TCefPdfPrintSettings record which is used in the TChromiumCore.PrintToPDF call.
type IPDFPrintOptions interface {
	IObject
	// Landscape
	//  Set to true for landscape mode or false for portrait mode.
	Landscape() bool // property
	// SetLandscape Set Landscape
	SetLandscape(AValue bool) // property
	// PrintBackground
	//  Set to true to print background graphics.
	PrintBackground() bool // property
	// SetPrintBackground Set PrintBackground
	SetPrintBackground(AValue bool) // property
	// PreferCSSPageSize
	//  Set to true to prefer page size as defined by css. Defaults to false,
	//  in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize() bool // property
	// SetPreferCSSPageSize Set PreferCSSPageSize
	SetPreferCSSPageSize(AValue bool) // property
	// PageRanges
	//  Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed
	//  in the document order, not in the order specified, and no more than once.
	//  Defaults to empty string, which implies the entire document is printed.
	//  The page numbers are quietly capped to actual page count of the document,
	//  and ranges beyond the end of the document are ignored. If this results in
	//  no pages to print, an error is reported. It is an error to specify a range
	//  with start greater than end.
	PageRanges() string // property
	// SetPageRanges Set PageRanges
	SetPageRanges(AValue string) // property
	// DisplayHeaderFooter
	//  Set to true to display the header and/or footer. Modify
	//  HeaderTemplate and/or FooterTemplate to customize the display.
	DisplayHeaderFooter() bool // property
	// SetDisplayHeaderFooter Set DisplayHeaderFooter
	SetDisplayHeaderFooter(AValue bool) // property
	// HeaderTemplate
	//  HTML template for the print header. Only displayed if
	//  DisplayHeaderFooter is true. Should be valid HTML markup with
	//  the following classes used to inject printing values into them:
	//  <code>
	//  - date: formatted print date
	//  - title: document title
	//  - url: document location
	//  - pageNumber: current page number
	//  - totalPages: total pages in the document
	//  </code>
	//  For example, "<span class=title></span>" would generate a span containing
	//  the title.
	HeaderTemplate() string // property
	// SetHeaderTemplate Set HeaderTemplate
	SetHeaderTemplate(AValue string) // property
	// FooterTemplate
	//  HTML template for the print footer. Only displayed if
	//  DisplayHeaderFooter is true. Uses the same format as
	//  HeaderTemplate.
	FooterTemplate() string // property
	// SetFooterTemplate Set FooterTemplate
	SetFooterTemplate(AValue string) // property
	// GenerateTaggedPDF
	//  Set to true to generate tagged(accessible) PDF.
	GenerateTaggedPDF() bool // property
	// SetGenerateTaggedPDF Set GenerateTaggedPDF
	SetGenerateTaggedPDF(AValue bool) // property
	// Scale
	//  The percentage to scale the PDF by before printing(e.g. .5 is 50%).
	//  If this value is less than or equal to zero the default value of 1.0
	//  will be used.
	Scale() (resultFloat64 float64) // property
	// SetScale Set Scale
	SetScale(AValue float64) // property
	// ScalePct
	//  The percentage value to scale the PDF by before printing(e.g. 50 is 50%).
	ScalePct() (resultFloat64 float64) // property
	// SetScalePct Set ScalePct
	SetScalePct(AValue float64) // property
	// PaperWidthInch
	//  Output paper width in inches. If either of these values is less than or
	//  equal to zero then the default paper size(letter, 8.5 x 11 inches) will
	//  be used.
	PaperWidthInch() (resultFloat64 float64) // property
	// SetPaperWidthInch Set PaperWidthInch
	SetPaperWidthInch(AValue float64) // property
	// PaperHeightInch
	//  Output paper height in inches. If either of these values is less than or
	//  equal to zero then the default paper size(letter, 8.5 x 11 inches) will
	//  be used.
	PaperHeightInch() (resultFloat64 float64) // property
	// SetPaperHeightInch Set PaperHeightInch
	SetPaperHeightInch(AValue float64) // property
	// PaperWidthMM
	//  Output paper width in mm.
	PaperWidthMM() (resultFloat64 float64) // property
	// SetPaperWidthMM Set PaperWidthMM
	SetPaperWidthMM(AValue float64) // property
	// PaperHeightMM
	//  Output paper height in mm.
	PaperHeightMM() (resultFloat64 float64) // property
	// SetPaperHeightMM Set PaperHeightMM
	SetPaperHeightMM(AValue float64) // property
	// MarginType
	//  Margin type.
	MarginType() TCefPdfPrintMarginType // property
	// SetMarginType Set MarginType
	SetMarginType(AValue TCefPdfPrintMarginType) // property
	// MarginTopInch
	//  Top margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginTopInch() (resultFloat64 float64) // property
	// SetMarginTopInch Set MarginTopInch
	SetMarginTopInch(AValue float64) // property
	// MarginRightInch
	//  Right margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginRightInch() (resultFloat64 float64) // property
	// SetMarginRightInch Set MarginRightInch
	SetMarginRightInch(AValue float64) // property
	// MarginBottomInch
	//  Bottom margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginBottomInch() (resultFloat64 float64) // property
	// SetMarginBottomInch Set MarginBottomInch
	SetMarginBottomInch(AValue float64) // property
	// MarginLeftInch
	//  Left margin in inches. Only used if MarginType is set to
	//  PDF_PRINT_MARGIN_CUSTOM.
	MarginLeftInch() (resultFloat64 float64) // property
	// SetMarginLeftInch Set MarginLeftInch
	SetMarginLeftInch(AValue float64) // property
	// MarginTopMM
	//  Top margin in mm.
	MarginTopMM() (resultFloat64 float64) // property
	// SetMarginTopMM Set MarginTopMM
	SetMarginTopMM(AValue float64) // property
	// MarginRightMM
	//  Right margin in mm.
	MarginRightMM() (resultFloat64 float64) // property
	// SetMarginRightMM Set MarginRightMM
	SetMarginRightMM(AValue float64) // property
	// MarginBottomMM
	//  Bottom margin in mm.
	MarginBottomMM() (resultFloat64 float64) // property
	// SetMarginBottomMM Set MarginBottomMM
	SetMarginBottomMM(AValue float64) // property
	// MarginLeftMM
	//  Left margin in mm.
	MarginLeftMM() (resultFloat64 float64) // property
	// SetMarginLeftMM Set MarginLeftMM
	SetMarginLeftMM(AValue float64) // property
	// CopyToSettings
	//  Copy the fields of this class to the TCefPdfPrintSettings parameter.
	CopyToSettings(aSettings *TCefPdfPrintSettings) // procedure
}

// TPDFPrintOptions Parent: TObject
//
//	The TPDFPrintOptions properties are used to fill the TCefPdfPrintSettings record which is used in the TChromiumCore.PrintToPDF call.
type TPDFPrintOptions struct {
	TObject
}

func NewPDFPrintOptions() IPDFPrintOptions {
	r1 := CEF().SysCallN(2190)
	return AsPDFPrintOptions(r1)
}

func (m *TPDFPrintOptions) Landscape() bool {
	r1 := CEF().SysCallN(2195, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetLandscape(AValue bool) {
	CEF().SysCallN(2195, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) PrintBackground() bool {
	r1 := CEF().SysCallN(2211, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetPrintBackground(AValue bool) {
	CEF().SysCallN(2211, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) PreferCSSPageSize() bool {
	r1 := CEF().SysCallN(2210, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetPreferCSSPageSize(AValue bool) {
	CEF().SysCallN(2210, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) PageRanges() string {
	r1 := CEF().SysCallN(2205, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPDFPrintOptions) SetPageRanges(AValue string) {
	CEF().SysCallN(2205, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPDFPrintOptions) DisplayHeaderFooter() bool {
	r1 := CEF().SysCallN(2191, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetDisplayHeaderFooter(AValue bool) {
	CEF().SysCallN(2191, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) HeaderTemplate() string {
	r1 := CEF().SysCallN(2194, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPDFPrintOptions) SetHeaderTemplate(AValue string) {
	CEF().SysCallN(2194, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPDFPrintOptions) FooterTemplate() string {
	r1 := CEF().SysCallN(2192, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPDFPrintOptions) SetFooterTemplate(AValue string) {
	CEF().SysCallN(2192, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPDFPrintOptions) GenerateTaggedPDF() bool {
	r1 := CEF().SysCallN(2193, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPDFPrintOptions) SetGenerateTaggedPDF(AValue bool) {
	CEF().SysCallN(2193, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPDFPrintOptions) Scale() (resultFloat64 float64) {
	CEF().SysCallN(2212, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetScale(AValue float64) {
	CEF().SysCallN(2212, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) ScalePct() (resultFloat64 float64) {
	CEF().SysCallN(2213, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetScalePct(AValue float64) {
	CEF().SysCallN(2213, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperWidthInch() (resultFloat64 float64) {
	CEF().SysCallN(2208, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthInch(AValue float64) {
	CEF().SysCallN(2208, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperHeightInch() (resultFloat64 float64) {
	CEF().SysCallN(2206, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightInch(AValue float64) {
	CEF().SysCallN(2206, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperWidthMM() (resultFloat64 float64) {
	CEF().SysCallN(2209, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperWidthMM(AValue float64) {
	CEF().SysCallN(2209, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) PaperHeightMM() (resultFloat64 float64) {
	CEF().SysCallN(2207, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetPaperHeightMM(AValue float64) {
	CEF().SysCallN(2207, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginType() TCefPdfPrintMarginType {
	r1 := CEF().SysCallN(2204, 0, m.Instance(), 0)
	return TCefPdfPrintMarginType(r1)
}

func (m *TPDFPrintOptions) SetMarginType(AValue TCefPdfPrintMarginType) {
	CEF().SysCallN(2204, 1, m.Instance(), uintptr(AValue))
}

func (m *TPDFPrintOptions) MarginTopInch() (resultFloat64 float64) {
	CEF().SysCallN(2202, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopInch(AValue float64) {
	CEF().SysCallN(2202, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginRightInch() (resultFloat64 float64) {
	CEF().SysCallN(2200, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightInch(AValue float64) {
	CEF().SysCallN(2200, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginBottomInch() (resultFloat64 float64) {
	CEF().SysCallN(2196, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomInch(AValue float64) {
	CEF().SysCallN(2196, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginLeftInch() (resultFloat64 float64) {
	CEF().SysCallN(2198, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftInch(AValue float64) {
	CEF().SysCallN(2198, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginTopMM() (resultFloat64 float64) {
	CEF().SysCallN(2203, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginTopMM(AValue float64) {
	CEF().SysCallN(2203, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginRightMM() (resultFloat64 float64) {
	CEF().SysCallN(2201, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginRightMM(AValue float64) {
	CEF().SysCallN(2201, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginBottomMM() (resultFloat64 float64) {
	CEF().SysCallN(2197, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginBottomMM(AValue float64) {
	CEF().SysCallN(2197, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TPDFPrintOptions) MarginLeftMM() (resultFloat64 float64) {
	CEF().SysCallN(2199, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TPDFPrintOptions) SetMarginLeftMM(AValue float64) {
	CEF().SysCallN(2199, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func PDFPrintOptionsClass() TClass {
	ret := CEF().SysCallN(2188)
	return TClass(ret)
}

func (m *TPDFPrintOptions) CopyToSettings(aSettings *TCefPdfPrintSettings) {
	var result0 tCefPdfPrintSettings
	CEF().SysCallN(2189, m.Instance(), uintptr(unsafePointer(&result0)))
	*aSettings = *(result0.Convert())
}
