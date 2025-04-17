//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
)

// The TChromiumFontOptions properties are used to fill the TCefBrowserSettings record which is used during the browser creation.
type TChromiumFontOptions struct {
	chromium IChromium
}

func NewChromiumFontOptions(chromium IChromium) *TChromiumFontOptions {
	return &TChromiumFontOptions{
		chromium: chromium,
	}
}

// Standard font family name.
func (m *TChromiumFontOptions) StandardFontFamily() string {
	if !m.chromium.IsValid() {
		return ""
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_StandardFontFamily).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return api.GoStr(result)
}

func (m *TChromiumFontOptions) SetStandardFontFamily(value string) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_StandardFontFamily).Call(1, m.chromium.Instance(), api.PascalStr(value), 0)
}

// Fixed font family name.
func (m *TChromiumFontOptions) FixedFontFamily() string {
	if !m.chromium.IsValid() {
		return ""
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_FixedFontFamily).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return api.GoStr(result)
}

func (m *TChromiumFontOptions) SetFixedFontFamily(value string) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_FixedFontFamily).Call(1, m.chromium.Instance(), api.PascalStr(value), 0)
}

// Serif font family name.
func (m *TChromiumFontOptions) SerifFontFamily() string {
	if !m.chromium.IsValid() {
		return ""
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_SerifFontFamily).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return api.GoStr(result)
}

func (m *TChromiumFontOptions) SetSerifFontFamily(value string) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_SerifFontFamily).Call(1, m.chromium.Instance(), api.PascalStr(value), 0)
}

// SansSerif font family name.
func (m *TChromiumFontOptions) SansSerifFontFamily() string {
	if !m.chromium.IsValid() {
		return ""
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_SansSerifFontFamily).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return api.GoStr(result)
}

func (m *TChromiumFontOptions) SetSansSerifFontFamily(value string) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_SansSerifFontFamily).Call(1, m.chromium.Instance(), api.PascalStr(value), 0)
}

// Cursive font family name.
func (m *TChromiumFontOptions) CursiveFontFamily() string {
	if !m.chromium.IsValid() {
		return ""
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_CursiveFontFamily).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return api.GoStr(result)
}

func (m *TChromiumFontOptions) SetCursiveFontFamily(value string) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_CursiveFontFamily).Call(1, m.chromium.Instance(), api.PascalStr(value), 0)
}

// Fantasy font family name.
func (m *TChromiumFontOptions) FantasyFontFamily() string {
	if !m.chromium.IsValid() {
		return ""
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_FantasyFontFamily).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return api.GoStr(result)
}

func (m *TChromiumFontOptions) SetFantasyFontFamily(value string) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_FantasyFontFamily).Call(1, m.chromium.Instance(), api.PascalStr(value), 0)
}

// Default font size.
func (m *TChromiumFontOptions) DefaultFontSize() int32 {
	if !m.chromium.IsValid() {
		return 0
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_DefaultFontSize).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return int32(result)
}

func (m *TChromiumFontOptions) SetDefaultFontSize(value int32) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_DefaultFontSize).Call(1, m.chromium.Instance(), uintptr(value), 0)
}

// Default fixed font size.
func (m *TChromiumFontOptions) DefaultFixedFontSize() int32 {
	if !m.chromium.IsValid() {
		return 0
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_DefaultFixedFontSize).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return int32(result)
}

func (m *TChromiumFontOptions) SetDefaultFixedFontSize(value int32) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_DefaultFixedFontSize).Call(1, m.chromium.Instance(), uintptr(value), 0)
}

// Minimum font size.
func (m *TChromiumFontOptions) MinimumFontSize() int32 {
	if !m.chromium.IsValid() {
		return 0
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_MinimumFontSize).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return int32(result)
}

func (m *TChromiumFontOptions) SetMinimumFontSize(value int32) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_MinimumFontSize).Call(1, m.chromium.Instance(), uintptr(value), 0)
}

// Minimum logical font size.
func (m *TChromiumFontOptions) MinimumLogicalFontSize() int32 {
	if !m.chromium.IsValid() {
		return 0
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_MinimumLogicalFontSize).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return int32(result)
}

func (m *TChromiumFontOptions) SetMinimumLogicalFontSize(value int32) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_MinimumLogicalFontSize).Call(1, m.chromium.Instance(), uintptr(value), 0)
}

// Controls the loading of fonts from remote sources. Also configurable using
// the "disable-remote-fonts" command-line switch.
func (m *TChromiumFontOptions) RemoteFonts() consts.TCefState {
	if !m.chromium.IsValid() {
		return 0
	}
	var result uintptr
	imports.Proc(def.ChromiumFontOptions_RemoteFonts).Call(0, m.chromium.Instance(), 0, uintptr(unsafePointer(&result)))
	return consts.TCefState(result)
}

func (m *TChromiumFontOptions) SetRemoteFonts(value consts.TCefState) {
	if !m.chromium.IsValid() {
		return
	}
	imports.Proc(def.ChromiumFontOptions_RemoteFonts).Call(1, m.chromium.Instance(), uintptr(value), 0)
}
