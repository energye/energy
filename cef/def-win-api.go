package cef

import (
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

func WinEndPaint(Handle types.HWND, PS types.TagPaintStruct) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_EndPaint).Call(uintptr(Handle), uintptr(unsafe.Pointer(&PS)))
	return types.Integer(r1)
}

func WinBeginPaint(Handle types.HWND, PS types.TagPaintStruct) types.HDC {
	r1, _, _ := imports.Proc(internale_CEF_Win_BeginPaint).Call(uintptr(Handle), uintptr(unsafe.Pointer(&PS)))
	return types.HDC(r1)
}

func WinArc(DC types.HDC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_Arc).Call(uintptr(DC), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom), uintptr(Angle16Deg), uintptr(Angle16DegLength))
	return types.LongBool(api.GoBool(r1))
}

func WinAngleChord(DC types.HDC, x1, y1, x2, y2, angle1, angle2 types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_AngleChord).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(angle1), uintptr(angle2))
	return types.LongBool(api.GoBool(r1))
}

func WinCallNextHookEx(hhk types.HOOK, ncode types.Integer, WParam types.WPARAM, LParam types.LPARAM) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_CallNextHookEx).Call(uintptr(hhk), uintptr(ncode), uintptr(WParam), uintptr(LParam))
	return types.Integer(r1)
}

func WinCallWindowProc(lpPrevWndFunc types.TFarProc, Handle types.HWND, Msg types.UINT, WParam types.WPARAM, LParam types.LPARAM) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_CallWindowProc).Call(uintptr(lpPrevWndFunc), uintptr(Handle), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return types.Integer(r1)
}

func WinBitBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Rop types.DWORD) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_BitBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Rop))
	return types.LongBool(api.GoBool(r1))
}

func WinCreateBitmap(Width, Height types.Integer, Planes, BitCount types.LongInt, BitmapBits types.Pointer) types.HBITMAP {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateBitmap).Call(uintptr(Width), uintptr(Height), uintptr(Planes), uintptr(BitCount), uintptr(BitmapBits))
	return types.HBITMAP(r1)
}

func WinCreateBrushIndirect(LogBrush types.TagLogBrush) types.HBRUSH {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateBrushIndirect).Call(uintptr(unsafe.Pointer(&LogBrush)))
	return types.HBITMAP(r1)
}

func WinCreateBrushWithRadialGradient(LogBrush types.TLogRadialGradient) types.HBRUSH {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateBrushWithRadialGradient).Call(uintptr(unsafe.Pointer(&LogBrush)))
	return types.HBITMAP(r1)
}

func WinCreateCaret(Handle types.HWND, Bitmap types.HBITMAP, width, Height types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateCaret).Call(uintptr(Handle), uintptr(Bitmap), uintptr(width), uintptr(Height))
	return types.LongBool(api.GoBool(r1))
}

func WinCreateCompatibleBitmap(DC types.HDC, Width, Height types.Integer) types.HBITMAP {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateCompatibleBitmap).Call(uintptr(DC), uintptr(Width), uintptr(Height))
	return types.HBITMAP(r1)
}

func WinCreateCompatibleDC(DC types.HDC) types.HDC {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateCompatibleDC).Call(uintptr(DC))
	return types.HDC(r1)
}

func WinCreateDIBitmap(DC types.HDC, InfoHeader types.TagBitmapInfoHeader, dwUsage types.DWORD, InitBits types.PChar, InitInfo types.TagBitmapInfo, wUsage types.UINT) types.HBITMAP {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateDIBitmap).Call(uintptr(DC), uintptr(unsafe.Pointer(&InfoHeader)), uintptr(dwUsage), InitBits.ToPtr(), uintptr(unsafe.Pointer(&InitInfo)), uintptr(wUsage))
	return types.HBITMAP(r1)
}

func WinCreateDIBSection(DC types.HDC, BitmapInfo types.TagBitmapInfo, Usage types.UINT, Bits types.Pointer, SectionHandle types.THandle, Offset types.DWORD) types.HBITMAP {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateDIBSection).Call(uintptr(DC), uintptr(unsafe.Pointer(&BitmapInfo)), uintptr(Usage), uintptr(Bits), uintptr(SectionHandle), uintptr(Offset))
	return types.HBITMAP(r1)
}

func WinCreateEllipticRgn(X1, Y1, X2, Y2 types.Integer) *types.HRGN {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateEllipticRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return types.NewHRGN(r1)
}

func WinCreateFontIndirect(LogFont types.LogFontA) types.HFONT {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateFontIndirect).Call(uintptr(unsafe.Pointer(&LogFont)))
	return types.HFONT(r1)
}

func WinCreateFontIndirectEx(LogFont types.LogFontA, LongFontName types.PChar) types.HFONT {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateFontIndirectEx).Call(uintptr(unsafe.Pointer(&LogFont)), LongFontName.ToPtr())
	return types.HFONT(r1)
}

func WinCreateIconIndirect(IconInfo types.ICONINFO) types.HICON {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreateIconIndirect).Call(uintptr(unsafe.Pointer(&IconInfo)))
	return types.HFONT(r1)
}

func WinCreatePalette(LogPalette types.TagLogPalette) types.HPALETTE {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreatePalette).Call(uintptr(unsafe.Pointer(&LogPalette)))
	return types.HPALETTE(r1)
}

func WinCreatePatternBrush(ABitmap types.HBITMAP) types.HBRUSH {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreatePatternBrush).Call(uintptr(ABitmap))
	return types.HBRUSH(r1)
}

func WinCreatePenIndirect(LogPen types.TagLogPen) types.HPEN {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreatePenIndirect).Call(uintptr(unsafe.Pointer(&LogPen)))
	return types.HPEN(r1)
}

func WinCreatePolygonRgn(Points types.Point, NumPts types.Integer, FillMode types.Integer) *types.HRGN {
	r1, _, _ := imports.Proc(internale_CEF_Win_CreatePolygonRgn).Call(uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), uintptr(FillMode))
	return types.NewHRGN(r1)
}

func WinDeleteCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(internale_CEF_Win_DeleteCriticalSection).Call(uintptr(CritSection))
}

func WinDeleteDC(hDC types.HDC) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DeleteDC).Call(uintptr(hDC))
	return types.LongBool(api.GoBool(r1))
}

func WinDestroyCaret(Handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DestroyCaret).Call(uintptr(Handle))
	return types.LongBool(api.GoBool(r1))
}

func WinDestroyCursor(Handle types.HCURSOR) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DestroyCursor).Call(uintptr(Handle))
	return types.LongBool(api.GoBool(r1))
}

func WinDestroyIcon(Handle types.HICON) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DestroyIcon).Call(uintptr(Handle))
	return types.LongBool(api.GoBool(r1))
}

func WinDrawFrameControl(DC types.HDC, Rect types.Rect, uType, uState types.Cardinal) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DrawFrameControl).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(uType), uintptr(uState))
	return types.LongBool(api.GoBool(r1))
}

func WinDrawFocusRect(DC types.HDC, Rect types.Rect) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DrawFocusRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)))
	return types.LongBool(api.GoBool(r1))
}

func WinDrawEdge(DC types.HDC, Rect types.Rect, edge types.Cardinal, grfFlags types.Cardinal) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_DrawEdge).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(edge), uintptr(grfFlags))
	return types.LongBool(api.GoBool(r1))
}

func WinDrawText(DC types.HDC, Str types.PChar, Count types.Integer, Rect types.Rect, Flags types.Cardinal) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_DrawText).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(unsafe.Pointer(&Rect)), uintptr(Flags))
	return types.Integer(r1)
}

func WinEnableScrollBar(Wnd types.HWND, wSBflags, wArrows types.Cardinal) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_EnableScrollBar).Call(uintptr(Wnd), uintptr(wSBflags), uintptr(wArrows))
	return types.LongBool(api.GoBool(r1))
}

func WinEnableWindow(hWnd types.HWND, bEnable types.LongBool) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_EnableWindow).Call(uintptr(hWnd), bEnable.ToPtr())
	return types.LongBool(api.GoBool(r1))
}

func WinEnterCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(internale_CEF_Win_EnterCriticalSection).Call(CritSection.ToPtr())
}

func WinEnumDisplayMonitors(hdc types.HDC, lprcClip types.Rect, lpfnEnum types.MonitorEnumProc, dwData types.LPARAM) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_EnumDisplayMonitors).Call(uintptr(hdc), uintptr(unsafe.Pointer(&lprcClip)), uintptr(lpfnEnum), uintptr(dwData))
	return types.LongBool(api.GoBool(r1))
}

func WinEnumFontFamilies(DC types.HDC, Family types.PChar, EnumFontFamProc types.FontEnumProc, LParam types.LPARAM) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_EnumFontFamilies).Call(uintptr(DC), Family.ToPtr(), uintptr(EnumFontFamProc), uintptr(LParam))
	return types.LongInt(r1)
}

func WinEnumFontFamiliesEx(DC types.HDC, lpLogFont types.LogFontA, Callback types.FontEnumExProc, LParam types.LPARAM, Flags types.DWORD) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_EnumFontFamiliesEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&lpLogFont)), uintptr(Callback), uintptr(LParam), uintptr(Flags))
	return types.LongInt(r1)
}

func WinEllipse(DC types.HDC, x1, y1, x2, y2 types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_Ellipse).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return types.LongBool(api.GoBool(r1))
}

func WinEqualRgn(Rgn1 *types.HRGN, Rgn2 *types.HRGN) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_EqualRgn).Call(Rgn1.Instance(), Rgn2.Instance())
	return types.LongBool(api.GoBool(r1))
}

func WinExcludeClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_ExcludeClipRect).Call(uintptr(dc), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom))
	return types.Integer(r1)
}

func WinExtCreatePen(dwPenStyle, dwWidth types.DWORD, lplb types.TagLogBrush, dwStyleCount types.DWORD, lpStyle types.DWORD) types.HPEN {
	r1, _, _ := imports.Proc(internale_CEF_Win_ExtCreatePen).Call(uintptr(dwPenStyle), uintptr(dwWidth), uintptr(unsafe.Pointer(&lplb)), uintptr(dwStyleCount), uintptr(lpStyle))
	return types.HPEN(r1)
}

func WinExtTextOut(DC types.HDC, X, Y types.Integer, Options types.LongInt, Rect types.Rect, Str types.PChar, Count types.LongInt, Dx types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ExtTextOut).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(Options), uintptr(unsafe.Pointer(&Rect)), Str.ToPtr(), uintptr(Count), uintptr(Dx))
	return types.LongBool(api.GoBool(r1))
}

func WinExtSelectClipRGN(dc types.HDC, rgn *types.HRGN, Mode types.LongInt) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_ExtSelectClipRGN).Call(uintptr(dc), rgn.Instance(), uintptr(Mode))
	return types.Integer(r1)
}

func WinFillRect(DC types.HDC, Rect types.Rect, Brush types.HBRUSH) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_FillRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(Brush))
	return types.LongBool(api.GoBool(r1))
}

func WinFillRgn(DC types.HDC, RegionHnd *types.HRGN, hbr types.HBRUSH) types.BOOL {
	r1, _, _ := imports.Proc(internale_CEF_Win_FillRgn).Call(uintptr(DC), RegionHnd.Instance(), uintptr(hbr))
	return types.BOOL(api.GoBool(r1))
}

func WinFloodFill(DC types.HDC, X, Y types.Integer, Color types.TGraphicsColor, FillStyle types.TGraphicsFillStyle, Brush types.HBRUSH) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_FloodFill).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(Color), uintptr(FillStyle), uintptr(Brush))
	return types.LongBool(api.GoBool(r1))
}

func WinFrameRect(DC types.HDC, Rect types.Rect, hBr types.HBRUSH) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_FrameRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(hBr))
	return types.Integer(r1)
}

func WinGetActiveWindow() types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetActiveWindow).Call()
	return types.HWND(r1)
}

func WinGetBitmapBits(Bitmap types.HBITMAP, Count types.LongInt, Bits types.Pointer) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetBitmapBits).Call(uintptr(Bitmap), uintptr(Count), uintptr(Bits))
	return types.LongInt(r1)
}

func WinGetBkColor(DC types.HDC) types.TColorRef {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetBkColor).Call(uintptr(DC))
	return types.TColorRef(r1)
}

func WinGetCapture() types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetCapture).Call()
	return types.HWND(r1)
}

func WinGetCaretPos(lpPoint types.Point) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetCaretPos).Call(uintptr(unsafe.Pointer(&lpPoint)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetClientRect(handle types.HWND, Rect types.Rect) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetClientRect).Call(uintptr(handle), uintptr(unsafe.Pointer(&Rect)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetClipBox(DC types.HDC, lpRect types.Rect) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetClipBox).Call(uintptr(DC), uintptr(unsafe.Pointer(&lpRect)))
	return types.LongInt(r1)
}

func WinGetClipRGN(DC types.HDC, RGN *types.HRGN) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetClipRGN).Call(uintptr(DC), RGN.Instance())
	return types.LongInt(r1)
}

func WinGetCurrentObject(DC types.HDC, uObjectType types.UINT) types.HGDIOBJ {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetCurrentObject).Call(uintptr(DC), uintptr(uObjectType))
	return types.HGDIOBJ(r1)
}

func WinGetCursorPos(lpPoint types.Point) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetCursorPos).Call(uintptr(unsafe.Pointer(&lpPoint)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetDC(hWnd types.HWND) types.HDC {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetDC).Call(uintptr(hWnd))
	return types.HDC(r1)
}

func WinGetDeviceCaps(DC types.HDC, Index types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetDeviceCaps).Call(uintptr(DC), uintptr(Index))
	return types.Integer(r1)
}

func WinGetDIBits(DC types.HDC, Bitmap types.HBITMAP, StartScan, NumScans types.UINT, Bits types.Pointer, BitInfo types.TagBitmapInfo, Usage types.UINT) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetDIBits).Call(uintptr(DC), uintptr(Bitmap), uintptr(StartScan), uintptr(NumScans), uintptr(Bits), uintptr(unsafe.Pointer(&BitInfo)), uintptr(Usage))
	return types.Integer(r1)
}

func WinGetDoubleClickTime() types.UINT {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetDoubleClickTime).Call()
	return types.UINT(r1)
}

func WinGetFocus() types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetFocus).Call()
	return types.HWND(r1)
}

func WinGetFontLanguageInfo(DC types.HDC) types.DWORD {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetFontLanguageInfo).Call(uintptr(DC))
	return types.DWORD(r1)
}

func WinGetForegroundWindow() types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetForegroundWindow).Call()
	return types.HWND(r1)
}

func WinGetIconInfo(AIcon types.HICON, AIconInfo types.ICONINFO) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetIconInfo).Call(uintptr(AIcon), uintptr(unsafe.Pointer(&AIconInfo)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetKeyState(nVirtKey types.Integer) types.Smallint {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetKeyState).Call(uintptr(nVirtKey))
	return types.Smallint(r1)
}

func WinGetMapMode(DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetMapMode).Call(uintptr(DC))
	return types.Integer(r1)
}

func WinGetMonitorInfo(hMonitor types.HMONITOR, lpmi types.TagMonitorInfo) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetMonitorInfo).Call(uintptr(hMonitor), uintptr(unsafe.Pointer(&lpmi)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetDpiForMonitor(hmonitor types.HMONITOR, dpiType types.MONITOR_DPI_TYPE, dpiX *types.UINT, dpiY *types.UINT) types.HRESULT { // out
	var (
		outDpiX types.UINT
		outDpiY types.UINT
	)
	r1, _, _ := imports.Proc(internale_CEF_Win_GetDpiForMonitor).Call(uintptr(hmonitor), uintptr(dpiType), uintptr(unsafe.Pointer(&outDpiY)), uintptr(outDpiY))
	*dpiX = outDpiX
	*dpiY = outDpiY
	return types.HRESULT(r1)
}

func WinGetObject(GDIObject types.HGDIOBJ, BufSize types.Integer, Buf types.Pointer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetObject).Call(uintptr(GDIObject), uintptr(BufSize), uintptr(Buf))
	return types.Integer(r1)
}

func WinGetParent(Handle types.HWND) types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetParent).Call(uintptr(Handle))
	return types.HWND(r1)
}

func WinGetProp(Handle types.HWND, Str types.PChar) types.Pointer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetProp).Call(uintptr(Handle), Str.ToPtr())
	return types.HWND(r1)
}

func WinGetRgnBox(RGN *types.HRGN, lpRect types.Rect) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetRgnBox).Call(RGN.Instance(), uintptr(unsafe.Pointer(&lpRect)))
	return types.LongInt(r1)
}

func WinGetROP2(DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetROP2).Call(uintptr(DC))
	return types.Integer(r1)
}

func WinGetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo types.TagScrollInfo) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetScrollInfo).Call(uintptr(Handle), uintptr(SBStyle), uintptr(unsafe.Pointer(&ScrollInfo)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetStockObject(Value types.Integer) types.THandle {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetStockObject).Call(uintptr(Value))
	return types.THandle(r1)
}

func WinGetSysColor(nIndex types.Integer) types.DWORD {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetSysColor).Call(uintptr(nIndex))
	return types.DWORD(r1)
}

func WinGetSysColorBrush(nIndex types.Integer) types.HBRUSH {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetSysColorBrush).Call(uintptr(nIndex))
	return types.HBRUSH(r1)
}

func WinGetSystemMetrics(nIndex types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetSystemMetrics).Call(uintptr(nIndex))
	return types.Integer(r1)
}

func WinGetTextColor(DC types.HDC) types.TColorRef {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetTextColor).Call(uintptr(DC))
	return types.TColorRef(r1)
}

func WinGetTextExtentExPoint(DC types.HDC, Str types.PChar, Count, MaxWidth types.Integer, MaxCount, PartialWidths types.Integer, Size types.Size) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetTextExtentExPoint).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(MaxWidth), uintptr(MaxCount), uintptr(PartialWidths), uintptr(unsafe.Pointer(&Size)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetTextExtentPoint(DC types.HDC, Str types.PChar, Count types.Integer, Size types.Size) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetTextExtentPoint).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(unsafe.Pointer(&Size)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetTextExtentPoint32(DC types.HDC, Str types.PChar, Count types.Integer, Size types.Size) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetTextExtentPoint32).Call(uintptr(DC), Str.ToPtr(), uintptr(Count), uintptr(unsafe.Pointer(&Size)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetTextMetrics(DC types.HDC, TM types.TagTextMetricA) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetTextMetrics).Call(uintptr(DC), uintptr(unsafe.Pointer(&TM)))
	return types.LongBool(api.GoBool(r1))
}

func WinGetViewPortExtEx(DC types.HDC, Size types.Size) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetViewPortExtEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&Size)))
	return types.Integer(r1)
}

func WinGetViewPortOrgEx(DC types.HDC, P types.Point) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetViewPortOrgEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&P)))
	return types.Integer(r1)
}

func WinGetWindowExtEx(DC types.HDC, Size types.Size) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetWindowExtEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&Size)))
	return types.Integer(r1)
}

func WinGetWindowLong(Handle types.HWND, int types.Integer) types.PtrInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetWindowLong).Call(uintptr(Handle), uintptr(int))
	return types.PtrInt(r1)
}

func WinGetWindowRect(Handle types.HWND, Rect types.Rect) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetWindowRect).Call(uintptr(Handle), uintptr(unsafe.Pointer(&Rect)))
	return types.Integer(r1)
}

func WinGetWindowSize(Handle types.HWND, Width, Height types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GetWindowSize).Call(uintptr(Handle), uintptr(Width), uintptr(Height))
	return types.LongBool(api.GoBool(r1))
}

func WinGetWindowOrgEx(dc types.HDC, P types.Point) types.Integer { // because of delphi compatibility
	r1, _, _ := imports.Proc(internale_CEF_Win_GetWindowOrgEx).Call(uintptr(dc), uintptr(unsafe.Pointer(&P)))
	return types.Integer(r1)
}

func WinGradientFill(DC types.HDC, Vertices types.TagTriVertex, NumVertices types.LongInt, Meshes types.Pointer, NumMeshes types.LongInt, Mode types.LongInt) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_GradientFill).Call(uintptr(DC), uintptr(unsafe.Pointer(&Vertices)), uintptr(NumVertices), uintptr(Meshes), uintptr(NumMeshes), uintptr(Mode))
	return types.LongBool(api.GoBool(r1))
}

func WinHideCaret(hWnd types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_HideCaret).Call(uintptr(hWnd))
	return types.LongBool(api.GoBool(r1))
}

func WinInitializeCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(internale_CEF_Win_InitializeCriticalSection).Call(uintptr(CritSection))
}

func WinIntersectClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_IntersectClipRect).Call(uintptr(dc), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom))
	return types.Integer(r1)
}

func WinInvalidateRect(aHandle types.HWND, ARect types.Rect, bErase types.LongBool) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_InvalidateRect).Call(uintptr(aHandle), uintptr(unsafe.Pointer(&ARect)), api.PascalBool(bool(bErase)))
	return types.LongBool(api.GoBool(r1))
}

func WinInvalidateRgn(Handle types.HWND, Rgn *types.HRGN, Erase types.LongBool) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_InvalidateRgn).Call(uintptr(Handle), Rgn.Instance(), api.PascalBool(bool(Erase)))
	return types.LongBool(api.GoBool(r1))
}

func WinIsDBCSLeadByte(TestChar byte) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_IsDBCSLeadByte).Call(uintptr(TestChar))
	return types.LongBool(api.GoBool(r1))
}

func WinIsIconic(handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_IsIconic).Call(uintptr(handle))
	return types.LongBool(api.GoBool(r1))
}

func WinIsWindow(handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_IsWindow).Call(uintptr(handle))
	return types.LongBool(api.GoBool(r1))
}

func WinIsWindowEnabled(handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_IsWindowEnabled).Call(uintptr(handle))
	return types.LongBool(api.GoBool(r1))
}

func WinIsWindowVisible(handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_IsWindowVisible).Call(uintptr(handle))
	return types.LongBool(api.GoBool(r1))
}

func WinIsZoomed(handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_IsZoomed).Call(uintptr(handle))
	return types.LongBool(api.GoBool(r1))
}

func WinLeaveCriticalSection(CritSection types.TCriticalSection) {
	imports.Proc(internale_CEF_Win_LeaveCriticalSection).Call(uintptr(CritSection))
}

func WinLineTo(DC types.HDC, X, Y types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_LineTo).Call(uintptr(DC), uintptr(X), uintptr(Y))
	return types.LongBool(api.GoBool(r1))
}

func WinLoadBitmap(hInstance types.THandle, lpBitmapName types.PChar) types.HBITMAP {
	r1, _, _ := imports.Proc(internale_CEF_Win_LoadBitmap).Call(uintptr(hInstance), lpBitmapName.ToPtr())
	return types.HBITMAP(r1)
}

func WinLoadIcon(hInstance types.THandle, lpIconName types.PChar) types.HICON {
	r1, _, _ := imports.Proc(internale_CEF_Win_LoadIcon).Call(uintptr(hInstance), lpIconName.ToPtr())
	return types.HICON(r1)
}

func WinMaskBltRop(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask types.HBITMAP, XMask, YMask types.Integer, Rop types.DWORD) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_MaskBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Mask), uintptr(XMask), uintptr(YMask), uintptr(Rop))
	return types.LongBool(api.GoBool(r1))
}

func WinMaskBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask types.HBITMAP, XMask, YMask types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_MaskBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Mask), uintptr(XMask), uintptr(YMask))
	return types.LongBool(api.GoBool(r1))
}

func WinMessageBox(hWnd types.HWND, lpText, lpCaption types.PChar, uType types.Cardinal) types.Integer { //def MB_OK
	r1, _, _ := imports.Proc(internale_CEF_Win_MessageBox).Call(uintptr(hWnd), lpText.ToPtr(), lpCaption.ToPtr(), uintptr(uType))
	return types.Integer(r1)
}

func WinMonitorFromPoint(ptScreenCoords types.Point, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := imports.Proc(internale_CEF_Win_MonitorFromPoint).Call(uintptr(unsafe.Pointer(&ptScreenCoords)), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func WinMonitorFromRect(lprcScreenCoords types.Rect, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := imports.Proc(internale_CEF_Win_MonitorFromRect).Call(uintptr(unsafe.Pointer(&lprcScreenCoords)), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func WinMonitorFromWindow(hWnd types.HWND, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := imports.Proc(internale_CEF_Win_MonitorFromWindow).Call(uintptr(hWnd), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func WinMoveToEx(DC types.HDC, X, Y types.Integer, OldPoint types.Point) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_MoveToEx).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(&OldPoint)))
	return types.LongBool(api.GoBool(r1))
}

func WinOffsetRgn(RGN *types.HRGN, nXOffset, nYOffset types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_OffsetRgn).Call(RGN.Instance(), uintptr(nXOffset), uintptr(nYOffset))
	return types.Integer(r1)
}

func WinPaintRgn(DC types.HDC, RGN *types.HRGN) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_PaintRgn).Call(uintptr(DC), RGN.Instance())
	return types.LongBool(api.GoBool(r1))
}

func WinPie(DC types.HDC, x1, y1, x2, y2, sx, sy, ex, ey types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_Pie).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(sx), uintptr(sy), uintptr(ex), uintptr(ey))
	return types.LongBool(api.GoBool(r1))
}

func WinPolyBezier(DC types.HDC, Points types.Point, NumPts types.Integer, Filled, Continuous types.LongBool) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_PolyBezier).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), api.PascalBool(bool(Filled)), api.PascalBool(bool(Continuous)))
	return types.LongBool(api.GoBool(r1))
}

func WinPolygon(DC types.HDC, Points types.Point, NumPts types.Integer, Winding types.LongBool) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_Polygon).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), api.PascalBool(bool(Winding)))
	return types.LongBool(api.GoBool(r1))
}

func WinPolyline(DC types.HDC, Points types.Point, NumPts types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_Polyline).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts))
	return types.LongBool(api.GoBool(r1))
}

func WinPostMessage(Handle types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_PostMessage).Call(uintptr(Handle), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return types.LongBool(api.GoBool(r1))
}

func WinRealizePalette(DC types.HDC) types.Cardinal {
	r1, _, _ := imports.Proc(internale_CEF_Win_RealizePalette).Call(uintptr(DC))
	return types.Cardinal(r1)
}

func WinRectangle(DC types.HDC, X1, Y1, X2, Y2 types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_Rectangle).Call(uintptr(DC), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return types.LongBool(api.GoBool(r1))
}

func WinRectInRegion(RGN *types.HRGN, ARect types.Rect) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_RectInRegion).Call(RGN.Instance(), uintptr(unsafe.Pointer(&ARect)))
	return types.LongBool(api.GoBool(r1))
}

func WinRectVisible(DC types.HDC, ARect types.Rect) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_RectVisible).Call(uintptr(DC), uintptr(unsafe.Pointer(&ARect)))
	return types.LongBool(api.GoBool(r1))
}

func WinRedrawWindow(Wnd types.HWND, lprcUpdate types.Rect, hrgnUpdate *types.HRGN, flags types.UINT) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_RedrawWindow).Call(uintptr(Wnd), uintptr(unsafe.Pointer(&lprcUpdate)), hrgnUpdate.Instance(), uintptr(flags))
	return types.LongBool(api.GoBool(r1))
}

func WinReleaseCapture() types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ReleaseCapture).Call()
	return types.LongBool(api.GoBool(r1))
}

func WinReleaseDC(hWnd types.HWND, DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_ReleaseDC).Call(uintptr(hWnd), uintptr(DC))
	return types.Integer(r1)
}

func WinRemoveProp(Handle types.HWND, Str types.PChar) types.THandle {
	r1, _, _ := imports.Proc(internale_CEF_Win_RemoveProp).Call(uintptr(Handle), Str.ToPtr())
	return types.THandle(r1)
}

func WinRestoreDC(DC types.HDC, SavedDC types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_RestoreDC).Call(uintptr(DC), uintptr(SavedDC))
	return types.LongBool(api.GoBool(r1))
}

func WinRoundRect(DC types.HDC, X1, Y1, X2, Y2 types.Integer, RX, RY types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_RoundRect).Call(uintptr(DC), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(RX), uintptr(RY))
	return types.LongBool(api.GoBool(r1))
}

func WinSaveDC(DC types.HDC) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SaveDC).Call(uintptr(DC))
	return types.Integer(r1)
}

func WinScrollWindowEx(hWnd types.HWND, dx, dy types.Integer, prcScroll, prcClip types.Rect, hrgnUpdate *types.HRGN, prcUpdate types.Rect, flags types.UINT) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ScrollWindowEx).Call(uintptr(hWnd), uintptr(dx), uintptr(dy), uintptr(unsafe.Pointer(&prcScroll)), uintptr(unsafe.Pointer(&prcClip)), hrgnUpdate.Instance(), uintptr(unsafe.Pointer(&prcUpdate)), uintptr(flags))
	return types.LongBool(api.GoBool(r1))
}

func WinSelectClipRGN(DC types.HDC, RGN *types.HRGN) types.LongInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_SelectClipRGN).Call(uintptr(DC), RGN.Instance())
	return types.LongInt(r1)
}

func WinSelectObject(DC types.HDC, GDIObj types.HGDIOBJ) types.HGDIOBJ {
	r1, _, _ := imports.Proc(internale_CEF_Win_SelectObject).Call(uintptr(DC), uintptr(GDIObj))
	return types.HGDIOBJ(r1)
}

func WinSelectPalette(DC types.HDC, Palette types.HPALETTE, ForceBackground types.LongBool) types.HPALETTE {
	r1, _, _ := imports.Proc(internale_CEF_Win_SelectPalette).Call(uintptr(DC), uintptr(Palette), api.PascalBool(bool(ForceBackground)))
	return types.HPALETTE(r1)
}

func WinSendMessage(HandleWnd types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LResult {
	r1, _, _ := imports.Proc(internale_CEF_Win_SendMessage).Call(uintptr(HandleWnd), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return types.LResult(r1)
}

func WinSetActiveWindow(Handle types.HWND) types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetActiveWindow).Call(uintptr(Handle))
	return types.HWND(r1)
}

func WinSetBkColor(DC types.HDC, Color types.TColorRef) types.TColorRef { //pbd
	r1, _, _ := imports.Proc(internale_CEF_Win_SetBkColor).Call(uintptr(DC), uintptr(Color))
	return types.TColorRef(r1)
}

func WinSetBkMode(DC types.HDC, bkMode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetBkMode).Call(uintptr(DC), uintptr(bkMode))
	return types.Integer(r1)
}

func WinSetCapture(AHandle types.HWND) types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetCapture).Call(uintptr(AHandle))
	return types.HWND(r1)
}

func WinSetCaretPos(X, Y types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetCaretPos).Call(uintptr(X), uintptr(Y))
	return types.LongBool(api.GoBool(r1))
}

func WinSetCaretPosEx(handle types.HWND, X, Y types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetCaretPosEx).Call(uintptr(handle), uintptr(X), uintptr(Y))
	return types.LongBool(api.GoBool(r1))
}

func WinSetCursorPos(X, Y types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetCursorPos).Call(uintptr(X), uintptr(Y))
	return types.LongBool(api.GoBool(r1))
}

func WinSetFocus(hWnd types.HWND) types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetFocus).Call(uintptr(hWnd))
	return types.HWND(r1)
}

func WinSetForegroundWindow(hWnd types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetForegroundWindow).Call(uintptr(hWnd))
	return types.LongBool(api.GoBool(r1))
}

func WinSetMapMode(DC types.HDC, fnMapMode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetMapMode).Call(uintptr(DC), uintptr(fnMapMode))
	return types.Integer(r1)
}

func WinSetMenu(AWindowHandle types.HWND, AMenuHandle types.HMENU) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetMenu).Call(uintptr(AWindowHandle), uintptr(AMenuHandle))
	return types.LongBool(api.GoBool(r1))
}

func WinSetParent(hWndChild types.HWND, hWndParent types.HWND) types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetParent).Call(uintptr(hWndChild), uintptr(hWndParent))
	return types.HWND(r1)
}

func WinSetProp(Handle types.HWND, Str types.PChar, Data types.Pointer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetProp).Call(uintptr(Handle), Str.ToPtr(), uintptr(Data))
	return types.LongBool(api.GoBool(r1))
}

func WinSetROP2(DC types.HDC, Mode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetROP2).Call(uintptr(DC), uintptr(Mode))
	return types.Integer(r1)
}

func WinSetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo types.TagScrollInfo, Redraw types.LongBool) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetScrollInfo).Call(uintptr(Handle), uintptr(SBStyle), uintptr(unsafe.Pointer(&ScrollInfo)), api.PascalBool(bool(Redraw)))
	return types.Integer(r1)
}

func WinSetStretchBltMode(DC types.HDC, StretchMode types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetStretchBltMode).Call(uintptr(DC), uintptr(StretchMode))
	return types.Integer(r1)
}

func WinSetTextCharacterExtra(_hdc types.HDC, nCharExtra types.Integer) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetTextCharacterExtra).Call(uintptr(_hdc), uintptr(nCharExtra))
	return types.Integer(r1)
}

func WinSetTextColor(DC types.HDC, Color types.TColorRef) types.TColorRef {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetTextColor).Call(uintptr(DC), uintptr(Color))
	return types.TColorRef(r1)
}

func WinSetWindowLong(Handle types.HWND, Idx types.Integer, NewLong types.PtrInt) types.PtrInt {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetWindowLong).Call(uintptr(Handle), uintptr(Idx), uintptr(NewLong))
	return types.PtrInt(r1)
}

func WinSetViewPortExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize types.Size) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetViewPortExtEx).Call(uintptr(DC), uintptr(XExtent), uintptr(YExtent), uintptr(unsafe.Pointer(&OldSize)))
	return types.LongBool(api.GoBool(r1))
}

func WinSetViewPortOrgEx(DC types.HDC, NewX, NewY types.Integer, OldPoint types.Point) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetViewPortOrgEx).Call(uintptr(DC), uintptr(NewX), uintptr(NewY), uintptr(unsafe.Pointer(&OldPoint)))
	return types.LongBool(api.GoBool(r1))
}

func WinSetWindowExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize types.Size) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetWindowExtEx).Call(uintptr(DC), uintptr(XExtent), uintptr(YExtent), uintptr(unsafe.Pointer(&OldSize)))
	return types.LongBool(api.GoBool(r1))
}

func WinSetWindowOrgEx(dc types.HDC, NewX, NewY types.Integer, OldPoint types.Point) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetWindowOrgEx).Call(uintptr(dc), uintptr(NewX), uintptr(NewY), uintptr(unsafe.Pointer(&OldPoint)))
	return types.LongBool(api.GoBool(r1))
}

func WinSetWindowPos(hWnd types.HWND, hWndInsertAfter types.HWND, X, Y, cx, cy types.Integer, uFlags types.UINT) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SetWindowPos).Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(X), uintptr(Y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return types.LongBool(api.GoBool(r1))
}

func WinShowCaret(hWnd types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ShowCaret).Call(uintptr(hWnd))
	return types.LongBool(api.GoBool(r1))
}

func WinShowScrollBar(Handle types.HWND, wBar types.Integer, bShow types.LongBool) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ShowScrollBar).Call(uintptr(Handle), uintptr(wBar), api.PascalBool(bool(bShow)))
	return types.LongBool(api.GoBool(r1))
}

func WinShowWindow(hWnd types.HWND, nCmdShow types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_ShowWindow).Call(uintptr(hWnd), uintptr(nCmdShow))
	return types.LongBool(api.GoBool(r1))
}

func WinStretchBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc, SrcWidth, SrcHeight types.Integer, Rop types.Cardinal) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_StretchBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(SrcWidth), uintptr(SrcHeight), uintptr(Rop))
	return types.LongBool(api.GoBool(r1))
}

func WinStretchDIBits(DC types.HDC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight types.Integer, Bits types.Pointer, BitsInfo types.TagBitmapInfo, Usage types.UINT, Rop types.DWORD) types.Integer {
	r1, _, _ := imports.Proc(internale_CEF_Win_StretchDIBits).Call(uintptr(DC), uintptr(DestX), uintptr(DestY), uintptr(DestWidth), uintptr(DestHeight), uintptr(SrcX), uintptr(SrcY), uintptr(SrcWidth), uintptr(SrcHeight), uintptr(Bits), uintptr(unsafe.Pointer(&BitsInfo)), uintptr(Usage), uintptr(Rop))
	return types.Integer(r1)
}

func WinSystemParametersInfo(uiAction types.DWORD, uiParam types.DWORD, pvParam types.Pointer, fWinIni types.DWORD) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_SystemParametersInfo).Call(uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni))
	return types.LongBool(api.GoBool(r1))
}

func WinTextOut(DC types.HDC, X, Y types.Integer, Str types.PChar, Count types.Integer) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_TextOut).Call(uintptr(DC), uintptr(X), uintptr(Y), Str.ToPtr(), uintptr(Count))
	return types.LongBool(api.GoBool(r1))
}

func WinUpdateWindow(Handle types.HWND) types.LongBool {
	r1, _, _ := imports.Proc(internale_CEF_Win_UpdateWindow).Call(uintptr(Handle))
	return types.LongBool(api.GoBool(r1))
}

func WinWindowFromPoint(Point types.Point) types.HWND {
	r1, _, _ := imports.Proc(internale_CEF_Win_WindowFromPoint).Call(uintptr(unsafe.Pointer(&Point)))
	return types.HWND(r1)
}
