//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

func WinCreateRectRgn(X1, Y1, X2, Y2 int32) *HRGN {
	r1, _, _ := common.Proc(internale_CEF_Win_CreateRectRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return &HRGN{
		instance: unsafe.Pointer(r1),
	}
}

func WinSetRectRgn(aRGN *HRGN, X1, Y1, X2, Y2 int32) bool {
	r1, _, _ := common.Proc(internale_CEF_Win_SetRectRgn).Call(uintptr(aRGN.instance), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return api.GoBool(r1)
}

func WinDeleteObject(aRGN *HRGN) bool {
	r1, _, _ := common.Proc(internale_CEF_Win_DeleteObject).Call(uintptr(aRGN.instance))
	return api.GoBool(r1)
}

func WinCombineRgn(dest, src1, src2 *HRGN, fnCombineMode consts.RNGFnCombineMode) int32 {
	r1, _, _ := common.Proc(internale_CEF_Win_CombineRgn).Call(uintptr(dest.instance), uintptr(src1.instance), uintptr(src2.instance), uintptr(fnCombineMode))
	return int32(r1)
}

func WinPtInRegion(RGN *HRGN, X, Y int32) bool {
	r1, _, _ := common.Proc(internale_CEF_Win_PtInRegion).Call(uintptr(RGN.instance), uintptr(X), uintptr(Y))
	return api.GoBool(r1)
}

func WinScreenToClient(handle types.HWND, p *types.TPoint) int32 {
	r1, _, _ := common.Proc(internale_CEF_Win_ScreenToClient).Call(handle, uintptr(unsafe.Pointer(p)))
	return int32(r1)
}

func WinClientToScreen(handle types.HWND, p *types.TPoint) bool {
	r1, _, _ := common.Proc(internale_CEF_Win_ClientToScreen).Call(handle, uintptr(unsafe.Pointer(p)))
	return api.GoBool(r1)
}

func WinDefWindowProc(handle types.HWND, msg types.UINT, wParam types.WPARAM, lParam types.LPARAM) types.LRESULT {
	r1, _, _ := common.Proc(internale_CEF_Win_DefWindowProc).Call(handle, uintptr(msg), wParam, lParam)
	return types.LRESULT(r1)
}

func WinDefSubclassProc(handle types.HWND, msg types.UINT, wParam types.WPARAM, lParam types.LPARAM) types.LRESULT {
	r1, _, _ := common.Proc(internale_CEF_Win_DefSubclassProc).Call(handle, uintptr(msg), wParam, lParam)
	return types.LRESULT(r1)
}

func WinOnPaint(handle types.HWND) {
	common.Proc(internale_CEF_Win_OnPaint).Call(handle)
}

func WinSetDraggableRegions(aRGN *HRGN, regions []TCefDraggableRegion) {
	/*
		//SetDraggableRegions 代码实现
		draggableRegion := WinCreateRectRgn(0, 0, 0, 0)
		WinSetRectRgn(draggableRegion, 0, 0, 0, 0)
		for i := 0; i < regions.RegionsCount(); i++ {
			region := regions.Region(i)
			creRGN := WinCreateRectRgn(region.Bounds.X, region.Bounds.Y, region.Bounds.X+region.Bounds.Width, region.Bounds.Y+region.Bounds.Height)
			if region.Draggable {
				WinCombineRgn(draggableRegion, draggableRegion, creRGN, consts.RGN_OR)
			} else {
				WinCombineRgn(draggableRegion, draggableRegion, creRGN, consts.RGN_DIFF)
			}
			WinDeleteObject(creRGN)
		}
		fmt.Println("Check PtInRegion：", WinPtInRegion(draggableRegion, 50, 50))
	*/
	common.Proc(internale_CEF_Win_SetDraggableRegions).Call(uintptr(aRGN.instance), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}
