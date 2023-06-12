//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy - lcl window api - test

package winapi

import (
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/energy/inits"
	"testing"
)

func TestWinApiCallback(t *testing.T) {
	inits.Init(nil, nil)
	callback := NewEnumDisplayMonitorsCallback()
	callback.Callback(func(hMonitor types.HMONITOR, hdcMonitor types.HDC, lprcMonitor types.Rect, dwData types.LPARAM) types.LongBool {
		t.Log("EnumDisplayMonitorsCallback.Callback:", hMonitor, hdcMonitor, lprcMonitor, dwData)
		return true
	})
	r0 := WinEnumDisplayMonitors(0, nil, callback, 0)
	t.Log("R0:", r0)

	dc := WinGetDC(0)
	t.Log("dc:", dc)
	familiesCallback := NewEnumFontFamiliesCallback()
	familiesCallback.Callback(func(ELogFont *types.TagEnumLogFontA, Metric *types.TNewTextMetric, FontType types.LongInt, Data types.LPARAM) types.LongInt {
		t.Log("EnumFontFamiliesCallback.Callback:", ELogFont, Metric)
		return 2
	})
	r1 := WinEnumFontFamilies(dc, "", familiesCallback, 0)
	t.Log("r1:", r1)

	dc = WinGetDC(0)
	t.Log("dc:", dc)
	familiesExCallback := NewEnumFontFamiliesExCallback()
	familiesExCallback.Callback(func(ELogFont *types.TagEnumLogFontExA, Metric *types.TNewTextMetricEx, FontType types.LongInt, Data types.LPARAM) types.LongInt {
		t.Log("EnumFontFamiliesExCallback.Callback:", ELogFont, Metric)
		return 2
	})
	r2 := WinEnumFontFamiliesEx(dc, types.LogFontA{}, familiesExCallback, 0, 0)
	t.Log("r2:", r2)
}
