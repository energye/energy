//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package winapi energy - lcl window api
package winapi

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
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

type HCursor struct {
	instance unsafe.Pointer
}

func (m *HCursor) Free() {
	m.instance = nil
}

func CreateRectRgn(X1, Y1, X2, Y2 int32) *types.HRGN {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateRectRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return types.NewHRGN(r1)
}

func SetRectRgn(aRGN *types.HRGN, X1, Y1, X2, Y2 int32) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetRectRgn).Call(aRGN.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return api.GoBool(r1)
}

func DeleteObject(aRGN *types.HRGN) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DeleteObject).Call(aRGN.Instance())
	return api.GoBool(r1)
}

func CombineRgn(dest, src1, src2 *types.HRGN, fnCombineMode consts.RNGFnCombineMode) int32 {
	r1, _, _ := imports.Proc(def.CEF_Win_CombineRgn).Call(dest.Instance(), src1.Instance(), src2.Instance(), uintptr(fnCombineMode))
	return int32(r1)
}

func PtInRegion(RGN *types.HRGN, X, Y int32) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_PtInRegion).Call(RGN.Instance(), uintptr(X), uintptr(Y))
	return api.GoBool(r1)
}

func ScreenToClient(handle types.HWND, p *types.Point) int32 {
	r1, _, _ := imports.Proc(def.CEF_Win_ScreenToClient).Call(uintptr(handle), uintptr(unsafe.Pointer(p)))
	return int32(r1)
}

func ClientToScreen(handle types.HWND, p *types.Point) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ClientToScreen).Call(uintptr(handle), uintptr(unsafe.Pointer(p)))
	return api.GoBool(r1)
}

func DefWindowProc(handle types.HWND, msg types.UINT, wParam types.WPARAM, lParam types.LPARAM) types.LResult {
	r1, _, _ := imports.Proc(def.CEF_Win_DefWindowProc).Call(uintptr(handle), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return types.LResult(r1)
}

func DefSubclassProc(handle types.HWND, msg types.UINT, wParam types.WPARAM, lParam types.LPARAM) types.LResult {
	r1, _, _ := imports.Proc(def.CEF_Win_DefSubclassProc).Call(uintptr(handle), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return types.LResult(r1)
}

func CreateRoundRectRgn(_para1, _para2, _para3, _para4, _para5, _para6 types.LongInt) *types.HRGN {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateRoundRectRgn).Call(_para1.ToPtr(), _para2.ToPtr(), _para3.ToPtr(), _para4.ToPtr(), _para5.ToPtr(), _para6.ToPtr())
	return types.NewHRGN(r1)
}

func SetWindowRgn(handle types.HWND, hRgn *types.HRGN, bRedraw bool) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowRgn).Call(uintptr(handle), hRgn.Instance(), api.PascalBool(bRedraw))
	return types.LongInt(r1)
}

func SetCursor(hCursor *HCursor) *HCursor {
	r1, _, _ := imports.Proc(def.CEF_Win_SetCursor).Call(uintptr(hCursor.instance))
	return &HCursor{
		instance: unsafe.Pointer(r1),
	}
}

func LoadCursor(handle types.HWND, lpCursorName int32) *HCursor {
	r1, _, _ := imports.Proc(def.CEF_Win_LoadCursor).Call(uintptr(handle), uintptr(lpCursorName))
	return &HCursor{
		instance: unsafe.Pointer(r1),
	}
}

func OnPaint(handle types.HWND) {
	imports.Proc(def.CEF_Win_OnPaint).Call(uintptr(handle))
}

//func SetDraggableRegions(aRGN *types.HRGN, regions []cef.TCefDraggableRegion) {
//	/*
//		//SetDraggableRegions 代码实现
//		draggableRegion := WinCreateRectRgn(0, 0, 0, 0)
//		WinSetRectRgn(draggableRegion, 0, 0, 0, 0)
//		for i := 0; i < regions.RegionsCount(); i++ {
//			region := regions.Region(i)
//			creRGN := WinCreateRectRgn(region.Bounds.X, region.Bounds.Y, region.Bounds.X+region.Bounds.Width, region.Bounds.Y+region.Bounds.Height)
//			if region.Draggable {
//				WinCombineRgn(draggableRegion, draggableRegion, creRGN, consts.RGN_OR)
//			} else {
//				WinCombineRgn(draggableRegion, draggableRegion, creRGN, consts.RGN_DIFF)
//			}
//			WinDeleteObject(creRGN)
//		}
//		fmt.Println("Check PtInRegion：", WinPtInRegion(draggableRegion, 50, 50))
//	*/
//	imports.Proc(def.CEF_Win_SetDraggableRegions).Call(aRGN.Instance(), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
//}

func EndPaint(Handle types.HWND, PS *types.TagPaintStruct) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_EndPaint).Call(uintptr(Handle), uintptr(unsafe.Pointer(PS)))
	return types.Integer(r1)
}

func BeginPaint(Handle types.HWND, PS *types.TagPaintStruct) types.HDC {
	r1, _, _ := imports.Proc(def.CEF_Win_BeginPaint).Call(uintptr(Handle), uintptr(unsafe.Pointer(PS)))
	return types.HDC(r1)
}

func Arc(DC types.HDC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_Arc).Call(uintptr(DC), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom), uintptr(Angle16Deg), uintptr(Angle16DegLength))
	return bool(api.GoBool(r1))
}

func AngleChord(DC types.HDC, x1, y1, x2, y2, angle1, angle2 types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_AngleChord).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(angle1), uintptr(angle2))
	return bool(api.GoBool(r1))
}

func CallNextHookEx(hhk types.HOOK, ncode types.Integer, WParam types.WPARAM, LParam types.LPARAM) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_CallNextHookEx).Call(uintptr(hhk), uintptr(ncode), uintptr(WParam), uintptr(LParam))
	return types.Integer(r1)
}

func CallWindowProc(lpPrevWndFunc types.TFarProc, Handle types.HWND, Msg types.UINT, WParam types.WPARAM, LParam types.LPARAM) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_CallWindowProc).Call(uintptr(lpPrevWndFunc), uintptr(Handle), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return types.Integer(r1)
}

func BitBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Rop types.DWORD) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_BitBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Rop))
	return bool(api.GoBool(r1))
}

func CreateBitmap(Width, Height types.Integer, Planes, BitCount types.LongInt, BitmapBits types.Pointer) types.HBITMAP {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateBitmap).Call(uintptr(Width), uintptr(Height), uintptr(Planes), uintptr(BitCount), uintptr(BitmapBits))
	return types.HBITMAP(r1)
}

func CreateBrushIndirect(LogBrush types.TagLogBrush) types.HBRUSH {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateBrushIndirect).Call(uintptr(unsafe.Pointer(&LogBrush)))
	return types.HBITMAP(r1)
}

func CreateBrushWithRadialGradient(LogBrush types.TLogRadialGradient) types.HBRUSH {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateBrushWithRadialGradient).Call(uintptr(unsafe.Pointer(&LogBrush)))
	return types.HBITMAP(r1)
}

func CreateCaret(Handle types.HWND, Bitmap types.HBITMAP, width, Height types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateCaret).Call(uintptr(Handle), uintptr(Bitmap), uintptr(width), uintptr(Height))
	return bool(api.GoBool(r1))
}

func CreateCompatibleBitmap(DC types.HDC, Width, Height types.Integer) types.HBITMAP {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateCompatibleBitmap).Call(uintptr(DC), uintptr(Width), uintptr(Height))
	return types.HBITMAP(r1)
}

func CreateCompatibleDC(DC types.HDC) types.HDC {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateCompatibleDC).Call(uintptr(DC))
	return types.HDC(r1)
}

func CreateDIBitmap(DC types.HDC, InfoHeader types.TagBitmapInfoHeader, dwUsage types.DWORD, InitBits types.PChar, InitInfo types.TagBitmapInfo, wUsage types.UINT) types.HBITMAP {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateDIBitmap).Call(uintptr(DC), uintptr(unsafe.Pointer(&InfoHeader)), uintptr(dwUsage), InitBits.ToPtr(), uintptr(unsafe.Pointer(&InitInfo)), uintptr(wUsage))
	return types.HBITMAP(r1)
}

func CreateDIBSection(DC types.HDC, BitmapInfo types.TagBitmapInfo, Usage types.UINT, Bits types.Pointer, SectionHandle types.THandle, Offset types.DWORD) types.HBITMAP {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateDIBSection).Call(uintptr(DC), uintptr(unsafe.Pointer(&BitmapInfo)), uintptr(Usage), uintptr(Bits), uintptr(SectionHandle), uintptr(Offset))
	return types.HBITMAP(r1)
}

func CreateEllipticRgn(X1, Y1, X2, Y2 types.Integer) *types.HRGN {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateEllipticRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return types.NewHRGN(r1)
}

func CreateFontIndirect(LogFont types.LogFontA) types.HFONT {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateFontIndirect).Call(uintptr(unsafe.Pointer(&LogFont)))
	return types.HFONT(r1)
}

func CreateFontIndirectEx(LogFont types.LogFontA, LongFontName types.PChar) types.HFONT {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateFontIndirectEx).Call(uintptr(unsafe.Pointer(&LogFont)), LongFontName.ToPtr())
	return types.HFONT(r1)
}

func CreateIconIndirect(IconInfo types.ICONINFO) types.HICON {
	r1, _, _ := imports.Proc(def.CEF_Win_CreateIconIndirect).Call(uintptr(unsafe.Pointer(&IconInfo)))
	return types.HFONT(r1)
}

func CreatePalette(LogPalette types.TagLogPalette) types.HPALETTE {
	r1, _, _ := imports.Proc(def.CEF_Win_CreatePalette).Call(uintptr(unsafe.Pointer(&LogPalette)))
	return types.HPALETTE(r1)
}

func CreatePatternBrush(ABitmap types.HBITMAP) types.HBRUSH {
	r1, _, _ := imports.Proc(def.CEF_Win_CreatePatternBrush).Call(uintptr(ABitmap))
	return types.HBRUSH(r1)
}

func CreatePenIndirect(LogPen types.TagLogPen) types.HPEN {
	r1, _, _ := imports.Proc(def.CEF_Win_CreatePenIndirect).Call(uintptr(unsafe.Pointer(&LogPen)))
	return types.HPEN(r1)
}

func CreatePolygonRgn(points []types.Point, NumPts types.Integer, FillMode types.Integer) *types.HRGN {
	r1, _, _ := imports.Proc(def.CEF_Win_CreatePolygonRgn).Call(uintptr(unsafe.Pointer(&points[0])), uintptr(NumPts), uintptr(FillMode))
	return types.NewHRGN(r1)
}

func DeleteCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(def.CEF_Win_DeleteCriticalSection).Call(uintptr(CritSection))
}

func DeleteDC(hDC types.HDC) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DeleteDC).Call(uintptr(hDC))
	return bool(api.GoBool(r1))
}

func DestroyCaret(Handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DestroyCaret).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func DestroyCursor(Handle types.HCURSOR) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DestroyCursor).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func DestroyIcon(Handle types.HICON) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DestroyIcon).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func DrawFrameControl(DC types.HDC, Rect types.Rect, uType, uState types.Cardinal) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DrawFrameControl).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(uType), uintptr(uState))
	return bool(api.GoBool(r1))
}

func DrawFocusRect(DC types.HDC, Rect types.Rect) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DrawFocusRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)))
	return bool(api.GoBool(r1))
}

func DrawEdge(DC types.HDC, Rect types.Rect, edge types.Cardinal, grfFlags types.Cardinal) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_DrawEdge).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(edge), uintptr(grfFlags))
	return bool(api.GoBool(r1))
}

func DrawText(DC types.HDC, Str types.PChar, Count types.Integer, Rect types.Rect, Flags types.Cardinal) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_DrawText).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(unsafe.Pointer(&Rect)), uintptr(Flags))
	return types.Integer(r1)
}

func EnableScrollBar(Wnd types.HWND, wSBflags, wArrows types.Cardinal) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_EnableScrollBar).Call(uintptr(Wnd), uintptr(wSBflags), uintptr(wArrows))
	return bool(api.GoBool(r1))
}

func EnableWindow(hWnd types.HWND, bEnable bool) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_EnableWindow).Call(uintptr(hWnd), api.PascalBool(bEnable))
	return bool(api.GoBool(r1))
}

func EnterCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(def.CEF_Win_EnterCriticalSection).Call(CritSection.ToPtr())
}

func EnumDisplayMonitors(hdc types.HDC, lprcClip *types.Rect, callback *EnumDisplayMonitorsCallback, dwData types.LPARAM) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_EnumDisplayMonitors).Call(uintptr(hdc), uintptr(unsafe.Pointer(lprcClip)), callback.instance, uintptr(dwData))
	return bool(api.GoBool(r1))
}

func EnumFontFamilies(DC types.HDC, Family types.PChar, callback *EnumFontFamiliesCallback, LParam types.LPARAM) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_EnumFontFamilies).Call(uintptr(DC), Family.ToPtr(), callback.instance, uintptr(LParam))
	return types.LongInt(r1)
}

func EnumFontFamiliesEx(DC types.HDC, lpLogFont types.LogFontA, callback *EnumFontFamiliesExCallback, LParam types.LPARAM, Flags types.DWORD) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_EnumFontFamiliesEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&lpLogFont)), callback.instance, uintptr(LParam), uintptr(Flags))
	return types.LongInt(r1)
}

func Ellipse(DC types.HDC, x1, y1, x2, y2 types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_Ellipse).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return bool(api.GoBool(r1))
}

func EqualRgn(Rgn1 *types.HRGN, Rgn2 *types.HRGN) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_EqualRgn).Call(Rgn1.Instance(), Rgn2.Instance())
	return bool(api.GoBool(r1))
}

func ExcludeClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_ExcludeClipRect).Call(uintptr(dc), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom))
	return types.Integer(r1)
}

func ExtCreatePen(dwPenStyle, dwWidth types.DWORD, lplb types.TagLogBrush, dwStyleCount types.DWORD, lpStyle types.DWORD) types.HPEN {
	r1, _, _ := imports.Proc(def.CEF_Win_ExtCreatePen).Call(uintptr(dwPenStyle), uintptr(dwWidth), uintptr(unsafe.Pointer(&lplb)), uintptr(dwStyleCount), uintptr(lpStyle))
	return types.HPEN(r1)
}

func ExtTextOut(DC types.HDC, X, Y types.Integer, Options types.LongInt, Rect types.Rect, Str types.PChar, Count types.LongInt, Dx types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ExtTextOut).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(Options), uintptr(unsafe.Pointer(&Rect)), Str.ToPtr(), uintptr(Count), uintptr(Dx))
	return bool(api.GoBool(r1))
}

func ExtSelectClipRGN(dc types.HDC, rgn *types.HRGN, Mode types.LongInt) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_ExtSelectClipRGN).Call(uintptr(dc), rgn.Instance(), uintptr(Mode))
	return types.Integer(r1)
}

func FillRect(DC types.HDC, Rect types.Rect, Brush types.HBRUSH) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_FillRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(Brush))
	return bool(api.GoBool(r1))
}

func FillRgn(DC types.HDC, RegionHnd *types.HRGN, hbr types.HBRUSH) types.BOOL {
	r1, _, _ := imports.Proc(def.CEF_Win_FillRgn).Call(uintptr(DC), RegionHnd.Instance(), uintptr(hbr))
	return types.BOOL(api.GoBool(r1))
}

func FloodFill(DC types.HDC, X, Y types.Integer, Color types.TGraphicsColor, FillStyle types.TGraphicsFillStyle, Brush types.HBRUSH) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_FloodFill).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(Color), uintptr(FillStyle), uintptr(Brush))
	return bool(api.GoBool(r1))
}

func FrameRect(DC types.HDC, Rect types.Rect, hBr types.HBRUSH) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_FrameRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(hBr))
	return types.Integer(r1)
}

func GetActiveWindow() types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetActiveWindow).Call()
	return types.HWND(r1)
}

func GetBitmapBits(Bitmap types.HBITMAP, Count types.LongInt, Bits types.Pointer) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetBitmapBits).Call(uintptr(Bitmap), uintptr(Count), uintptr(Bits))
	return types.LongInt(r1)
}

func GetBkColor(DC types.HDC) types.TColorRef {
	r1, _, _ := imports.Proc(def.CEF_Win_GetBkColor).Call(uintptr(DC))
	return types.TColorRef(r1)
}

func GetCapture() types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetCapture).Call()
	return types.HWND(r1)
}

func GetCaretPos(lpPoint *types.Point) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetCaretPos).Call(uintptr(unsafe.Pointer(lpPoint)))
	return bool(api.GoBool(r1))
}

func GetClientRect(handle types.HWND, Rect *types.Rect) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClientRect).Call(uintptr(handle), uintptr(unsafe.Pointer(Rect)))
	return bool(api.GoBool(r1))
}

func GetClipBox(DC types.HDC, lpRect *types.Rect) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClipBox).Call(uintptr(DC), uintptr(unsafe.Pointer(lpRect)))
	return types.LongInt(r1)
}

func GetClipRGN(DC types.HDC, RGN *types.HRGN) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetClipRGN).Call(uintptr(DC), RGN.Instance())
	return types.LongInt(r1)
}

func GetCurrentObject(DC types.HDC, uObjectType types.UINT) types.HGDIOBJ {
	r1, _, _ := imports.Proc(def.CEF_Win_GetCurrentObject).Call(uintptr(DC), uintptr(uObjectType))
	return types.HGDIOBJ(r1)
}

func GetCursorPos(lpPoint types.Point) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetCursorPos).Call(uintptr(unsafe.Pointer(&lpPoint)))
	return bool(api.GoBool(r1))
}

func GetDC(hWnd types.HWND) types.HDC {
	r1, _, _ := imports.Proc(def.CEF_Win_GetDC).Call(uintptr(hWnd))
	return types.HDC(r1)
}

func GetDeviceCaps(DC types.HDC, Index types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetDeviceCaps).Call(uintptr(DC), uintptr(Index))
	return types.Integer(r1)
}

func GetDIBits(DC types.HDC, Bitmap types.HBITMAP, StartScan, NumScans types.UINT, Bits types.Pointer, BitInfo types.TagBitmapInfo, Usage types.UINT) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetDIBits).Call(uintptr(DC), uintptr(Bitmap), uintptr(StartScan), uintptr(NumScans), uintptr(Bits), uintptr(unsafe.Pointer(&BitInfo)), uintptr(Usage))
	return types.Integer(r1)
}

func GetDoubleClickTime() types.UINT {
	r1, _, _ := imports.Proc(def.CEF_Win_GetDoubleClickTime).Call()
	return types.UINT(r1)
}

func GetFocus() types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetFocus).Call()
	return types.HWND(r1)
}

func GetFontLanguageInfo(DC types.HDC) types.DWORD {
	r1, _, _ := imports.Proc(def.CEF_Win_GetFontLanguageInfo).Call(uintptr(DC))
	return types.DWORD(r1)
}

func GetForegroundWindow() types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetForegroundWindow).Call()
	return types.HWND(r1)
}

func GetIconInfo(AIcon types.HICON, AIconInfo types.ICONINFO) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetIconInfo).Call(uintptr(AIcon), uintptr(unsafe.Pointer(&AIconInfo)))
	return bool(api.GoBool(r1))
}

func GetKeyState(nVirtKey types.Integer) types.Smallint {
	r1, _, _ := imports.Proc(def.CEF_Win_GetKeyState).Call(uintptr(nVirtKey))
	return types.Smallint(r1)
}

func GetMapMode(DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetMapMode).Call(uintptr(DC))
	return types.Integer(r1)
}

func GetMonitorInfo(hMonitor types.HMONITOR, lpmi *types.TagMonitorInfo) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetMonitorInfo).Call(uintptr(hMonitor), uintptr(unsafe.Pointer(lpmi)))
	return api.GoBool(r1)
}

func GetDpiForMonitor(hmonitor types.HMONITOR, dpiType MONITOR_DPI_TYPE, dpiX *types.UINT, dpiY *types.UINT) types.HRESULT { // out
	var (
		outDpiX uint32
		outDpiY uint32
	)
	r1, _, _ := imports.Proc(def.CEF_Win_GetDpiForMonitor).Call(uintptr(hmonitor), uintptr(dpiType), uintptr(unsafe.Pointer(&outDpiX)), uintptr(unsafe.Pointer(&outDpiY)))
	*dpiX = types.UINT(outDpiX)
	*dpiY = types.UINT(outDpiY)
	return types.HRESULT(r1)
}

func GetObject(GDIObject types.HGDIOBJ, BufSize types.Integer, Buf types.Pointer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetObject).Call(uintptr(GDIObject), uintptr(BufSize), uintptr(Buf))
	return types.Integer(r1)
}

func GetParent(Handle types.HWND) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_GetParent).Call(uintptr(Handle))
	return types.HWND(r1)
}

func GetProp(Handle types.HWND, Str types.PChar) types.Pointer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetProp).Call(uintptr(Handle), Str.ToPtr())
	return types.HWND(r1)
}

func GetRgnBox(RGN *types.HRGN, lpRect types.Rect) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetRgnBox).Call(RGN.Instance(), uintptr(unsafe.Pointer(&lpRect)))
	return types.LongInt(r1)
}

func GetROP2(DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetROP2).Call(uintptr(DC))
	return types.Integer(r1)
}

func GetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo types.TagScrollInfo) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetScrollInfo).Call(uintptr(Handle), uintptr(SBStyle), uintptr(unsafe.Pointer(&ScrollInfo)))
	return bool(api.GoBool(r1))
}

func GetStockObject(Value types.Integer) types.THandle {
	r1, _, _ := imports.Proc(def.CEF_Win_GetStockObject).Call(uintptr(Value))
	return types.THandle(r1)
}

func GetSysColor(nIndex types.Integer) types.DWORD {
	r1, _, _ := imports.Proc(def.CEF_Win_GetSysColor).Call(uintptr(nIndex))
	return types.DWORD(r1)
}

func GetSysColorBrush(nIndex types.Integer) types.HBRUSH {
	r1, _, _ := imports.Proc(def.CEF_Win_GetSysColorBrush).Call(uintptr(nIndex))
	return types.HBRUSH(r1)
}

func GetSystemMetrics(nIndex types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetSystemMetrics).Call(uintptr(nIndex))
	return types.Integer(r1)
}

func GetTextColor(DC types.HDC) types.TColorRef {
	r1, _, _ := imports.Proc(def.CEF_Win_GetTextColor).Call(uintptr(DC))
	return types.TColorRef(r1)
}

func GetTextExtentExPoint(DC types.HDC, Str types.PChar, Count, MaxWidth types.Integer, MaxCount, PartialWidths types.Integer, Size *types.Size) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetTextExtentExPoint).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(MaxWidth), uintptr(MaxCount), uintptr(PartialWidths), uintptr(unsafe.Pointer(Size)))
	return bool(api.GoBool(r1))
}

func GetTextExtentPoint(DC types.HDC, Str types.PChar, Count types.Integer, Size *types.Size) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetTextExtentPoint).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(unsafe.Pointer(Size)))
	return bool(api.GoBool(r1))
}

func GetTextExtentPoint32(DC types.HDC, Str types.PChar, Count types.Integer, Size *types.Size) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetTextExtentPoint32).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(unsafe.Pointer(Size)))
	return bool(api.GoBool(r1))
}

func GetTextMetrics(DC types.HDC, TM *types.TagTextMetricA) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetTextMetrics).Call(uintptr(DC), uintptr(unsafe.Pointer(TM)))
	return bool(api.GoBool(r1))
}

func GetViewPortExtEx(DC types.HDC, Size *types.Size) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetViewPortExtEx).Call(uintptr(DC), uintptr(unsafe.Pointer(Size)))
	return types.Integer(r1)
}

func GetViewPortOrgEx(DC types.HDC, P *types.Point) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetViewPortOrgEx).Call(uintptr(DC), uintptr(unsafe.Pointer(P)))
	return types.Integer(r1)
}

func GetWindowExtEx(DC types.HDC, Size *types.Size) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowExtEx).Call(uintptr(DC), uintptr(unsafe.Pointer(Size)))
	return types.Integer(r1)
}

func GetWindowLong(Handle types.HWND, int types.Integer) types.PtrInt {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowLong).Call(uintptr(Handle), uintptr(int))
	return types.PtrInt(r1)
}

func GetWindowRect(Handle types.HWND, Rect *types.Rect) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowRect).Call(uintptr(Handle), uintptr(unsafe.Pointer(Rect)))
	return types.Integer(r1)
}

func GetWindowSize(Handle types.HWND, Width, Height *types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowSize).Call(uintptr(Handle), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
	return bool(api.GoBool(r1))
}

func GetWindowOrgEx(dc types.HDC, P *types.Point) types.Integer { // because of delphi compatibility
	r1, _, _ := imports.Proc(def.CEF_Win_GetWindowOrgEx).Call(uintptr(dc), uintptr(unsafe.Pointer(P)))
	return types.Integer(r1)
}

func GradientFill(DC types.HDC, Vertices types.TagTriVertex, NumVertices types.LongInt, Meshes types.Pointer, NumMeshes types.LongInt, Mode types.LongInt) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_GradientFill).Call(uintptr(DC), uintptr(unsafe.Pointer(&Vertices)), uintptr(NumVertices), uintptr(Meshes), uintptr(NumMeshes), uintptr(Mode))
	return bool(api.GoBool(r1))
}

func HideCaret(hWnd types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_HideCaret).Call(uintptr(hWnd))
	return bool(api.GoBool(r1))
}

func InitializeCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(def.CEF_Win_InitializeCriticalSection).Call(uintptr(CritSection))
}

func IntersectClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_IntersectClipRect).Call(uintptr(dc), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom))
	return types.Integer(r1)
}

func InvalidateRect(aHandle types.HWND, ARect types.Rect, bErase bool) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_InvalidateRect).Call(uintptr(aHandle), uintptr(unsafe.Pointer(&ARect)), api.PascalBool(bool(bErase)))
	return bool(api.GoBool(r1))
}

func InvalidateRgn(Handle types.HWND, Rgn *types.HRGN, Erase bool) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_InvalidateRgn).Call(uintptr(Handle), Rgn.Instance(), api.PascalBool(bool(Erase)))
	return bool(api.GoBool(r1))
}

func IsDBCSLeadByte(TestChar byte) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_IsDBCSLeadByte).Call(uintptr(TestChar))
	return bool(api.GoBool(r1))
}

func IsIconic(handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_IsIconic).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsWindow(handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_IsWindow).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsWindowEnabled(handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_IsWindowEnabled).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsWindowVisible(handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_IsWindowVisible).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsZoomed(handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_IsZoomed).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func LeaveCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(def.CEF_Win_LeaveCriticalSection).Call(uintptr(CritSection))
}

func LineTo(DC types.HDC, X, Y types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_LineTo).Call(uintptr(DC), uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func LoadBitmap(hInstance types.THandle, lpBitmapName types.PChar) types.HBITMAP {
	r1, _, _ := imports.Proc(def.CEF_Win_LoadBitmap).Call(uintptr(hInstance), lpBitmapName.ToPtr())
	return types.HBITMAP(r1)
}

func LoadIcon(hInstance types.THandle, lpIconName types.PChar) types.HICON {
	r1, _, _ := imports.Proc(def.CEF_Win_LoadIcon).Call(uintptr(hInstance), lpIconName.ToPtr())
	return types.HICON(r1)
}

func MaskBltRop(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask types.HBITMAP, XMask, YMask types.Integer, Rop types.DWORD) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_MaskBltRop).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Mask), uintptr(XMask), uintptr(YMask), uintptr(Rop))
	return bool(api.GoBool(r1))
}

func MaskBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask types.HBITMAP, XMask, YMask types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_MaskBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Mask), uintptr(XMask), uintptr(YMask))
	return bool(api.GoBool(r1))
}

func MessageBox(hWnd types.HWND, lpText, lpCaption types.PChar, uType types.Cardinal) types.Integer { //def MB_OK
	r1, _, _ := imports.Proc(def.CEF_Win_MessageBox).Call(uintptr(hWnd), lpText.ToPtr(), lpCaption.ToPtr(), uintptr(uType))
	return types.Integer(r1)
}

func MonitorFromPoint(ptScreenCoords types.Point, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := imports.Proc(def.CEF_Win_MonitorFromPoint).Call(uintptr(unsafe.Pointer(&ptScreenCoords)), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func MonitorFromRect(lprcScreenCoords types.Rect, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := imports.Proc(def.CEF_Win_MonitorFromRect).Call(uintptr(unsafe.Pointer(&lprcScreenCoords)), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func MonitorFromWindow(hWnd types.HWND, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := imports.Proc(def.CEF_Win_MonitorFromWindow).Call(uintptr(hWnd), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func MoveToEx(DC types.HDC, X, Y types.Integer, OldPoint types.Point) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_MoveToEx).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(&OldPoint)))
	return bool(api.GoBool(r1))
}

func OffsetRgn(RGN *types.HRGN, nXOffset, nYOffset types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_OffsetRgn).Call(RGN.Instance(), uintptr(nXOffset), uintptr(nYOffset))
	return types.Integer(r1)
}

func PaintRgn(DC types.HDC, RGN *types.HRGN) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_PaintRgn).Call(uintptr(DC), RGN.Instance())
	return bool(api.GoBool(r1))
}

func Pie(DC types.HDC, x1, y1, x2, y2, sx, sy, ex, ey types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_Pie).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(sx), uintptr(sy), uintptr(ex), uintptr(ey))
	return bool(api.GoBool(r1))
}

func PolyBezier(DC types.HDC, Points types.Point, NumPts types.Integer, Filled, Continuous bool) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_PolyBezier).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), api.PascalBool(bool(Filled)), api.PascalBool(bool(Continuous)))
	return bool(api.GoBool(r1))
}

func Polygon(DC types.HDC, Points types.Point, NumPts types.Integer, Winding bool) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_Polygon).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), api.PascalBool(bool(Winding)))
	return bool(api.GoBool(r1))
}

func Polyline(DC types.HDC, Points types.Point, NumPts types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_Polyline).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts))
	return bool(api.GoBool(r1))
}

func PostMessage(Handle types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_PostMessage).Call(uintptr(Handle), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return bool(api.GoBool(r1))
}

func RealizePalette(DC types.HDC) types.Cardinal {
	r1, _, _ := imports.Proc(def.CEF_Win_RealizePalette).Call(uintptr(DC))
	return types.Cardinal(r1)
}

func Rectangle(DC types.HDC, X1, Y1, X2, Y2 types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_Rectangle).Call(uintptr(DC), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return bool(api.GoBool(r1))
}

func RectInRegion(RGN *types.HRGN, ARect types.Rect) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_RectInRegion).Call(RGN.Instance(), uintptr(unsafe.Pointer(&ARect)))
	return bool(api.GoBool(r1))
}

func RectVisible(DC types.HDC, ARect types.Rect) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_RectVisible).Call(uintptr(DC), uintptr(unsafe.Pointer(&ARect)))
	return bool(api.GoBool(r1))
}

func RedrawWindow(Wnd types.HWND, lprcUpdate types.Rect, hrgnUpdate *types.HRGN, flags types.UINT) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_RedrawWindow).Call(uintptr(Wnd), uintptr(unsafe.Pointer(&lprcUpdate)), hrgnUpdate.Instance(), uintptr(flags))
	return bool(api.GoBool(r1))
}

func ReleaseCapture() bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ReleaseCapture).Call()
	return bool(api.GoBool(r1))
}

func ReleaseDC(hWnd types.HWND, DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_ReleaseDC).Call(uintptr(hWnd), uintptr(DC))
	return types.Integer(r1)
}

func RemoveProp(Handle types.HWND, Str types.PChar) types.THandle {
	r1, _, _ := imports.Proc(def.CEF_Win_RemoveProp).Call(uintptr(Handle), Str.ToPtr())
	return types.THandle(r1)
}

func RestoreDC(DC types.HDC, SavedDC types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_RestoreDC).Call(uintptr(DC), uintptr(SavedDC))
	return bool(api.GoBool(r1))
}

func RoundRect(DC types.HDC, X1, Y1, X2, Y2 types.Integer, RX, RY types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_RoundRect).Call(uintptr(DC), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(RX), uintptr(RY))
	return bool(api.GoBool(r1))
}

func SaveDC(DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SaveDC).Call(uintptr(DC))
	return types.Integer(r1)
}

func ScrollWindowEx(hWnd types.HWND, dx, dy types.Integer, prcScroll, prcClip types.Rect, hrgnUpdate *types.HRGN, prcUpdate types.Rect, flags types.UINT) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ScrollWindowEx).Call(uintptr(hWnd), uintptr(dx), uintptr(dy), uintptr(unsafe.Pointer(&prcScroll)), uintptr(unsafe.Pointer(&prcClip)), hrgnUpdate.Instance(), uintptr(unsafe.Pointer(&prcUpdate)), uintptr(flags))
	return bool(api.GoBool(r1))
}

func SelectClipRGN(DC types.HDC, RGN *types.HRGN) types.LongInt {
	r1, _, _ := imports.Proc(def.CEF_Win_SelectClipRGN).Call(uintptr(DC), RGN.Instance())
	return types.LongInt(r1)
}

func SelectObject(DC types.HDC, GDIObj types.HGDIOBJ) types.HGDIOBJ {
	r1, _, _ := imports.Proc(def.CEF_Win_SelectObject).Call(uintptr(DC), uintptr(GDIObj))
	return types.HGDIOBJ(r1)
}

func SelectPalette(DC types.HDC, Palette types.HPALETTE, ForceBackground bool) types.HPALETTE {
	r1, _, _ := imports.Proc(def.CEF_Win_SelectPalette).Call(uintptr(DC), uintptr(Palette), api.PascalBool(bool(ForceBackground)))
	return types.HPALETTE(r1)
}

func SendMessage(HandleWnd types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LResult {
	r1, _, _ := imports.Proc(def.CEF_Win_SendMessage).Call(uintptr(HandleWnd), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return types.LResult(r1)
}

func SetActiveWindow(Handle types.HWND) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_SetActiveWindow).Call(uintptr(Handle))
	return types.HWND(r1)
}

func SetBkColor(DC types.HDC, Color types.TColorRef) types.TColorRef { //pbd
	r1, _, _ := imports.Proc(def.CEF_Win_SetBkColor).Call(uintptr(DC), uintptr(Color))
	return types.TColorRef(r1)
}

func SetBkMode(DC types.HDC, bkMode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SetBkMode).Call(uintptr(DC), uintptr(bkMode))
	return types.Integer(r1)
}

func SetCapture(AHandle types.HWND) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_SetCapture).Call(uintptr(AHandle))
	return types.HWND(r1)
}

func SetCaretPos(X, Y types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetCaretPos).Call(uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func SetCaretPosEx(handle types.HWND, X, Y types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetCaretPosEx).Call(uintptr(handle), uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func SetCursorPos(X, Y types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetCursorPos).Call(uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func SetFocus(hWnd types.HWND) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_SetFocus).Call(uintptr(hWnd))
	return types.HWND(r1)
}

func SetForegroundWindow(hWnd types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetForegroundWindow).Call(uintptr(hWnd))
	return bool(api.GoBool(r1))
}

func SetMapMode(DC types.HDC, fnMapMode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SetMapMode).Call(uintptr(DC), uintptr(fnMapMode))
	return types.Integer(r1)
}

func SetMenu(AWindowHandle types.HWND, AMenuHandle types.HMENU) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetMenu).Call(uintptr(AWindowHandle), uintptr(AMenuHandle))
	return bool(api.GoBool(r1))
}

func SetParent(hWndChild types.HWND, hWndParent types.HWND) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_SetParent).Call(uintptr(hWndChild), uintptr(hWndParent))
	return types.HWND(r1)
}

func SetProp(Handle types.HWND, Str types.PChar, Data types.Pointer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetProp).Call(uintptr(Handle), Str.ToPtr(), uintptr(Data))
	return bool(api.GoBool(r1))
}

func SetROP2(DC types.HDC, Mode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SetROP2).Call(uintptr(DC), uintptr(Mode))
	return types.Integer(r1)
}

func SetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo types.TagScrollInfo, Redraw bool) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SetScrollInfo).Call(uintptr(Handle), uintptr(SBStyle), uintptr(unsafe.Pointer(&ScrollInfo)), api.PascalBool(bool(Redraw)))
	return types.Integer(r1)
}

func SetStretchBltMode(DC types.HDC, StretchMode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SetStretchBltMode).Call(uintptr(DC), uintptr(StretchMode))
	return types.Integer(r1)
}

func SetTextCharacterExtra(_hdc types.HDC, nCharExtra types.Integer) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_SetTextCharacterExtra).Call(uintptr(_hdc), uintptr(nCharExtra))
	return types.Integer(r1)
}

func SetTextColor(DC types.HDC, Color types.TColorRef) types.TColorRef {
	r1, _, _ := imports.Proc(def.CEF_Win_SetTextColor).Call(uintptr(DC), uintptr(Color))
	return types.TColorRef(r1)
}

func SetWindowLong(Handle types.HWND, Idx types.Integer, NewLong types.PtrInt) types.PtrInt {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowLong).Call(uintptr(Handle), uintptr(Idx), uintptr(NewLong))
	return types.PtrInt(r1)
}

func SetViewPortExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize types.Size) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetViewPortExtEx).Call(uintptr(DC), uintptr(XExtent), uintptr(YExtent), uintptr(unsafe.Pointer(&OldSize)))
	return bool(api.GoBool(r1))
}

func SetViewPortOrgEx(DC types.HDC, NewX, NewY types.Integer, OldPoint types.Point) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetViewPortOrgEx).Call(uintptr(DC), uintptr(NewX), uintptr(NewY), uintptr(unsafe.Pointer(&OldPoint)))
	return bool(api.GoBool(r1))
}

func SetWindowExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize types.Size) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowExtEx).Call(uintptr(DC), uintptr(XExtent), uintptr(YExtent), uintptr(unsafe.Pointer(&OldSize)))
	return bool(api.GoBool(r1))
}

func SetWindowOrgEx(dc types.HDC, NewX, NewY types.Integer, OldPoint types.Point) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowOrgEx).Call(uintptr(dc), uintptr(NewX), uintptr(NewY), uintptr(unsafe.Pointer(&OldPoint)))
	return bool(api.GoBool(r1))
}

func SetWindowPos(hWnd types.HWND, hWndInsertAfter types.HWND, X, Y, cx, cy types.Integer, uFlags types.UINT) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SetWindowPos).Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(X), uintptr(Y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return bool(api.GoBool(r1))
}

func ShowCaret(hWnd types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ShowCaret).Call(uintptr(hWnd))
	return bool(api.GoBool(r1))
}

func ShowScrollBar(Handle types.HWND, wBar types.Integer, bShow bool) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ShowScrollBar).Call(uintptr(Handle), uintptr(wBar), api.PascalBool(bool(bShow)))
	return bool(api.GoBool(r1))
}

func ShowWindow(hWnd types.HWND, nCmdShow types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_ShowWindow).Call(uintptr(hWnd), uintptr(nCmdShow))
	return bool(api.GoBool(r1))
}

func StretchBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc, SrcWidth, SrcHeight types.Integer, Rop types.Cardinal) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_StretchBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(SrcWidth), uintptr(SrcHeight), uintptr(Rop))
	return bool(api.GoBool(r1))
}

func StretchDIBits(DC types.HDC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight types.Integer, Bits types.Pointer, BitsInfo types.TagBitmapInfo, Usage types.UINT, Rop types.DWORD) types.Integer {
	r1, _, _ := imports.Proc(def.CEF_Win_StretchDIBits).Call(uintptr(DC), uintptr(DestX), uintptr(DestY), uintptr(DestWidth), uintptr(DestHeight), uintptr(SrcX), uintptr(SrcY), uintptr(SrcWidth), uintptr(SrcHeight), uintptr(Bits), uintptr(unsafe.Pointer(&BitsInfo)), uintptr(Usage), uintptr(Rop))
	return types.Integer(r1)
}

func SystemParametersInfo(uiAction types.DWORD, uiParam types.DWORD, pvParam types.Pointer, fWinIni types.DWORD) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_SystemParametersInfo).Call(uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni))
	return bool(api.GoBool(r1))
}

func TextOut(DC types.HDC, X, Y types.Integer, Str types.PChar, Count types.Integer) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_TextOut).Call(uintptr(DC), uintptr(X), uintptr(Y), Str.ToPtr(), uintptr(Count))
	return bool(api.GoBool(r1))
}

func UpdateWindow(Handle types.HWND) bool {
	r1, _, _ := imports.Proc(def.CEF_Win_UpdateWindow).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func WindowFromPoint(Point types.Point) types.HWND {
	r1, _, _ := imports.Proc(def.CEF_Win_WindowFromPoint).Call(uintptr(unsafe.Pointer(&Point)))
	return types.HWND(r1)
}

func ScalePercent() float32 {
	dc := GetDC(0)
	dpiX := GetDeviceCaps(dc, LOGPIXELSX)
	ReleaseDC(0, dc)
	return float32(dpiX) / 96.0
}
