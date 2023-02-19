//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	t "github.com/energye/energy/types"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

const (
	S_OK           = 0x00000000
	S_FALSE        = 0x00000001
	E_UNEXPECTED   = 0x8000FFFF
	E_NOTIMPL      = 0x80004001
	E_OUTOFMEMORY  = 0x8007000E
	E_INVALIDARG   = 0x80070057
	E_NOINTERFACE  = 0x80004002
	E_POINTER      = 0x80004003
	E_HANDLE       = 0x80070006
	E_ABORT        = 0x80004004
	E_FAIL         = 0x80004005
	E_ACCESSDENIED = 0x80070005
	E_PENDING      = 0x8000000A
)

// Win32 Predefined cursor constants
const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_ICON        = 32641
	IDC_SIZE        = 32640
)

const (
	FALSE = 0
	TRUE  = 1
)

const (
	WMSZ_LEFT        = 1
	WMSZ_RIGHT       = 2
	WMSZ_TOP         = 3
	WMSZ_TOPLEFT     = 4
	WMSZ_TOPRIGHT    = 5
	WMSZ_BOTTOM      = 6
	WMSZ_BOTTOMLEFT  = 7
	WMSZ_BOTTOMRIGHT = 8
)

type HRGN struct {
	instance unsafe.Pointer
}

func (m *HRGN) Free() {
	m.instance = nil
}

type HCursor struct {
	instance unsafe.Pointer
}

func (m *HCursor) Free() {
	m.instance = nil
}

func WinCreateRectRgn(X1, Y1, X2, Y2 int32) *HRGN {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateRectRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return &HRGN{
		instance: unsafe.Pointer(r1),
	}
}

func WinSetRectRgn(aRGN *HRGN, X1, Y1, X2, Y2 int32) bool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetRectRgn).Call(uintptr(aRGN.instance), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return api.GoBool(r1)
}

func WinDeleteObject(aRGN *HRGN) bool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DeleteObject).Call(uintptr(aRGN.instance))
	return api.GoBool(r1)
}

func WinCombineRgn(dest, src1, src2 *HRGN, fnCombineMode consts.RNGFnCombineMode) int32 {
	r1, _, _ := imports.Proc(internale_CEF_Win_CombineRgn).Call(uintptr(dest.instance), uintptr(src1.instance), uintptr(src2.instance), uintptr(fnCombineMode))
	return int32(r1)
}

func WinPtInRegion(RGN *HRGN, X, Y int32) bool {
	r1, _, _ := imports.Proc(internale_CEF_Win_PtInRegion).Call(uintptr(RGN.instance), uintptr(X), uintptr(Y))
	return api.GoBool(r1)
}

func WinScreenToClient(handle types.HWND, p *types.TPoint) int32 {
	r1, _, _ := imports.Proc(internale_CEF_Win_ScreenToClient).Call(handle, uintptr(unsafe.Pointer(p)))
	return int32(r1)
}

func WinClientToScreen(handle types.HWND, p *types.TPoint) bool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ClientToScreen).Call(handle, uintptr(unsafe.Pointer(p)))
	return api.GoBool(r1)
}

func WinDefWindowProc(handle types.HWND, msg types.UINT, wParam types.WPARAM, lParam types.LPARAM) types.LRESULT {
	r1, _, _ := imports.Proc(internale_CEF_Win_DefWindowProc).Call(handle, uintptr(msg), wParam, lParam)
	return types.LRESULT(r1)
}

func WinDefSubclassProc(handle types.HWND, msg types.UINT, wParam types.WPARAM, lParam types.LPARAM) types.LRESULT {
	r1, _, _ := imports.Proc(internale_CEF_Win_DefSubclassProc).Call(handle, uintptr(msg), wParam, lParam)
	return types.LRESULT(r1)
}

func WinCreateRoundRectRgn(_para1, _para2, _para3, _para4, _para5, _para6 t.LongInt) *HRGN {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateRoundRectRgn).Call(_para1.ToPtr(), _para2.ToPtr(), _para3.ToPtr(), _para4.ToPtr(), _para5.ToPtr(), _para6.ToPtr())
	return &HRGN{
		instance: unsafe.Pointer(r1),
	}
}

func WinSetWindowRgn(handle types.HWND, hRgn *HRGN, bRedraw bool) t.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetWindowRgn).Call(handle, uintptr(hRgn.instance), api.PascalBool(bRedraw))
	return t.LongInt(r1)
}

func WinSetCursor(hCursor *HCursor) *HCursor {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetCursor).Call(uintptr(hCursor.instance))
	return &HCursor{
		instance: unsafe.Pointer(r1),
	}
}

func WinLoadCursor(handle types.HWND, lpCursorName int32) *HCursor {
	r1, _, _ := imports.Proc(internale_CEF_Win_LoadCursor).Call(handle, uintptr(lpCursorName))
	return &HCursor{
		instance: unsafe.Pointer(r1),
	}
}

func WinOnPaint(handle types.HWND) {
	imports.Proc(internale_CEF_Win_OnPaint).Call(handle)
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
	imports.Proc(internale_CEF_Win_SetDraggableRegions).Call(uintptr(aRGN.instance), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

type (
	BOOL    int32
	HRESULT int32
)

func LOBYTE(w uint16) byte {
	return byte(w)
}

func HIBYTE(w uint16) byte {
	return byte(w >> 8 & 0xff)
}

func LOWORD(dw uint32) uint16 {
	return uint16(dw & 0xFFFF)
}

func HIWORD(dw uint32) uint16 {
	return uint16(dw >> 16 & 0xffff)
}

func GET_X_LPARAM(lp uintptr) int32 {
	return int32(int16(LOWORD(uint32(lp))))
}

func GET_Y_LPARAM(lp uintptr) int32 {
	return int32(int16(HIWORD(uint32(lp))))
}
