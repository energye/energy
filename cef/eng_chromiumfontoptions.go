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

// IChromiumFontOptions Parent: IPersistent
//
//	The TChromiumFontOptions properties are used to fill the TCefBrowserSettings record which is used during the browser creation.
type IChromiumFontOptions interface {
	IPersistent
	// StandardFontFamily
	//  Standard font family name.
	StandardFontFamily() string // property
	// SetStandardFontFamily Set StandardFontFamily
	SetStandardFontFamily(AValue string) // property
	// FixedFontFamily
	//  Fixed font family name.
	FixedFontFamily() string // property
	// SetFixedFontFamily Set FixedFontFamily
	SetFixedFontFamily(AValue string) // property
	// SerifFontFamily
	//  Serif font family name.
	SerifFontFamily() string // property
	// SetSerifFontFamily Set SerifFontFamily
	SetSerifFontFamily(AValue string) // property
	// SansSerifFontFamily
	//  SansSerif font family name.
	SansSerifFontFamily() string // property
	// SetSansSerifFontFamily Set SansSerifFontFamily
	SetSansSerifFontFamily(AValue string) // property
	// CursiveFontFamily
	//  Cursive font family name.
	CursiveFontFamily() string // property
	// SetCursiveFontFamily Set CursiveFontFamily
	SetCursiveFontFamily(AValue string) // property
	// FantasyFontFamily
	//  Fantasy font family name.
	FantasyFontFamily() string // property
	// SetFantasyFontFamily Set FantasyFontFamily
	SetFantasyFontFamily(AValue string) // property
	// DefaultFontSize
	//  Default font size.
	DefaultFontSize() int32 // property
	// SetDefaultFontSize Set DefaultFontSize
	SetDefaultFontSize(AValue int32) // property
	// DefaultFixedFontSize
	//  Default fixed font size.
	DefaultFixedFontSize() int32 // property
	// SetDefaultFixedFontSize Set DefaultFixedFontSize
	SetDefaultFixedFontSize(AValue int32) // property
	// MinimumFontSize
	//  Minimum font size.
	MinimumFontSize() int32 // property
	// SetMinimumFontSize Set MinimumFontSize
	SetMinimumFontSize(AValue int32) // property
	// MinimumLogicalFontSize
	//  Minimum logical font size.
	MinimumLogicalFontSize() int32 // property
	// SetMinimumLogicalFontSize Set MinimumLogicalFontSize
	SetMinimumLogicalFontSize(AValue int32) // property
	// RemoteFonts
	//  Controls the loading of fonts from remote sources. Also configurable using
	//  the "disable-remote-fonts" command-line switch.
	RemoteFonts() TCefState // property
	// SetRemoteFonts Set RemoteFonts
	SetRemoteFonts(AValue TCefState) // property
}

// TChromiumFontOptions Parent: TPersistent
//
//	The TChromiumFontOptions properties are used to fill the TCefBrowserSettings record which is used during the browser creation.
type TChromiumFontOptions struct {
	TPersistent
}

func NewChromiumFontOptions() IChromiumFontOptions {
	r1 := CEF().SysCallN(2084)
	return AsChromiumFontOptions(r1)
}

func (m *TChromiumFontOptions) StandardFontFamily() string {
	r1 := CEF().SysCallN(2095, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetStandardFontFamily(AValue string) {
	CEF().SysCallN(2095, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) FixedFontFamily() string {
	r1 := CEF().SysCallN(2089, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetFixedFontFamily(AValue string) {
	CEF().SysCallN(2089, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) SerifFontFamily() string {
	r1 := CEF().SysCallN(2094, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetSerifFontFamily(AValue string) {
	CEF().SysCallN(2094, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) SansSerifFontFamily() string {
	r1 := CEF().SysCallN(2093, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetSansSerifFontFamily(AValue string) {
	CEF().SysCallN(2093, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) CursiveFontFamily() string {
	r1 := CEF().SysCallN(2085, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetCursiveFontFamily(AValue string) {
	CEF().SysCallN(2085, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) FantasyFontFamily() string {
	r1 := CEF().SysCallN(2088, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumFontOptions) SetFantasyFontFamily(AValue string) {
	CEF().SysCallN(2088, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumFontOptions) DefaultFontSize() int32 {
	r1 := CEF().SysCallN(2087, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetDefaultFontSize(AValue int32) {
	CEF().SysCallN(2087, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) DefaultFixedFontSize() int32 {
	r1 := CEF().SysCallN(2086, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetDefaultFixedFontSize(AValue int32) {
	CEF().SysCallN(2086, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) MinimumFontSize() int32 {
	r1 := CEF().SysCallN(2090, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetMinimumFontSize(AValue int32) {
	CEF().SysCallN(2090, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) MinimumLogicalFontSize() int32 {
	r1 := CEF().SysCallN(2091, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumFontOptions) SetMinimumLogicalFontSize(AValue int32) {
	CEF().SysCallN(2091, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumFontOptions) RemoteFonts() TCefState {
	r1 := CEF().SysCallN(2092, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumFontOptions) SetRemoteFonts(AValue TCefState) {
	CEF().SysCallN(2092, 1, m.Instance(), uintptr(AValue))
}

func ChromiumFontOptionsClass() TClass {
	ret := CEF().SysCallN(2083)
	return TClass(ret)
}
