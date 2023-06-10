//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy - lcl window api - callback
// TODO : All functions of this API have not been fully tested yet

package cef

import (
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
)

type enumDisplayMonitorsProc func(hMonitor types.HMONITOR, hdcMonitor types.HDC, lprcMonitor types.Rect, dwData types.LPARAM) types.LongBool
type enumFontFamiliesProc func(ELogFont types.TagEnumLogFontA, Metric types.TNewTextMetric, FontType types.LongInt, Data types.LPARAM) types.LongInt
type enumFontFamiliesExProc func(ELogFont types.TagEnumLogFontExA, Metric types.TNewTextMetricEx, FontType types.LongInt, Data types.LPARAM) types.LongInt

type EnumDisplayMonitorsCallback struct {
	instance uintptr
}

type EnumFontFamiliesCallback struct {
	instance uintptr
}

type EnumFontFamiliesExCallback struct {
	instance uintptr
}

func NewEnumDisplayMonitorsCallback() *EnumDisplayMonitorsCallback {
	return &EnumDisplayMonitorsCallback{}
}

func (m *EnumDisplayMonitorsCallback) Callback(fn enumDisplayMonitorsProc) {
	if m.instance == 0 {
		m.instance = api.MakeEventDataPtr(fn)
	}
}

// Free 使用完需要将释放掉
func (m *EnumDisplayMonitorsCallback) Free() {
	if m.instance != 0 {
		api.RemoveEventElement(m.instance)
		imports.Proc(internale_CEF_Win_EnumDisplayMonitorsCallbackFree).Call()
	}
}

func NewEnumFontFamiliesCallback() *EnumFontFamiliesCallback {
	return &EnumFontFamiliesCallback{}
}

func (m *EnumFontFamiliesCallback) Callback(fn enumFontFamiliesProc) {
	if m.instance == 0 {
		m.instance = api.MakeEventDataPtr(fn)
	}
}

// Free 使用完需要将释放掉
func (m *EnumFontFamiliesCallback) Free() {
	if m.instance != 0 {
		api.RemoveEventElement(m.instance)
		imports.Proc(internale_CEF_Win_EnumFontFamiliesCallbackFree).Call()
	}
}

func NewEnumFontFamiliesExCallback() *EnumFontFamiliesExCallback {
	return &EnumFontFamiliesExCallback{}
}

func (m *EnumFontFamiliesExCallback) Callback(fn enumFontFamiliesExProc) {
	if m.instance == 0 {
		m.instance = api.MakeEventDataPtr(fn)
		imports.Proc(internale_CEF_Win_EnumFontFamiliesExCallbackFree).Call()
	}
}

// Free 使用完需要将释放掉
func (m *EnumFontFamiliesExCallback) Free() {
	if m.instance != 0 {
		api.RemoveEventElement(m.instance)
	}
}
