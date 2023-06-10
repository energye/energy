package cef

import "github.com/energye/energy/v2/types"

func WinEndPaint(Handle types.HWND, PS types.TagPaintStruct) types.Integer {

	return EndPaint(Handle, PS)
}

func WinPaint(Handle types.HWND, PS types.TagPaintStruct) types.HDC {

	return Paint(Handle, PS)
}

func WinArc(DC types.HDC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength types.Integer) types.LongBool {

	return Arc(DC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength)
}

func WinAngleChord(DC types.HDC, x1, y1, x2, y2, angle1, angle2 types.Integer) types.LongBool {

	return AngleChord(DC, x1, y1, x2, y2, angle1, angle2)
}

func WinCallNextHookEx(hhk types.HOOK, ncode types.Integer, WParam types.WPARAM, LParam types.LPARAM) types.Integer {

	return CallNextHookEx(hhk, ncode, WParam, LParam)
}

func WinCallWindowProc(lpPrevWndFunc types.TFarProc, Handle types.HWND, Msg types.UINT, WParam types.WPARAM, LParam types.LPARAM) types.Integer {

	return CallWindowProc(lpPrevWndFunc, Handle, Msg, WParam, LParam)
}

func WinBitBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Rop types.DWORD) types.LongBool {

	return BitBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, Rop)
}

func WinCreateBitmap(Width, Height types.Integer, Planes, BitCount types.LongInt, BitmapBits types.Pointer) types.HBITMAP {

	return CreateBitmap(Width, Height, Planes, BitCount, BitmapBits)
}

func WinCreateBrushIndirect(LogBrush types.TagLogBrush) types.HBRUSH {

	return CreateBrushIndirect(LogBrush)
}

func WinCreateBrushWithRadialGradient(LogBrush types.TLogRadialGradient) types.HBRUSH {

	return CreateBrushWithRadialGradient(LogBrush)
}

func WinCreateCaret(Handle types.HWND, Bitmap types.HBITMAP, width, Height types.Integer) types.LongBool {

	return CreateCaret(Handle, Bitmap, width, Height)
}

func WinCreateCompatibleBitmap(DC types.HDC, Width, Height types.Integer) types.HBITMAP {

	return CreateCompatibleBitmap(DC, Width, Height)
}

func WinCreateCompatibleDC(DC types.HDC) types.HDC {

	return CreateCompatibleDC(DC)
}

func WinCreateDIBitmap(DC types.HDC, InfoHeader types.TagBitmapInfoHeader, dwUsage types.DWORD, InitBits types.PChar, InitInfo types.TagBitmapInfo, wUsage types.UINT) types.HBITMAP {

	return CreateDIBitmap(DC, InfoHeader, dwUsage, InitBits, InitInfo, wUsage)
}

func WinCreateDIBSection(DC types.HDC, BitmapInfo types.TagBitmapInfo, Usage types.UINT, Bits types.Pointer, SectionHandle types.THandle, Offset types.DWORD) types.HBITMAP {

	return CreateDIBSection(DC, BitmapInfo, Usage, Bits, SectionHandle, Offset)
}

func WinCreateEllipticRgn(X1, Y1, X2, Y2 types.Integer) HRGN {

	return CreateEllipticRgn(X1, Y1, X2, Y2)
}

func WinCreateFontIndirect(LogFont types.LogFontA) types.HFONT {

	return CreateFontIndirect(LogFont)
}

func WinCreateFontIndirectEx(LogFont types.LogFontA, LongFontName types.PChar) types.HFONT {

	return CreateFontIndirectEx(LogFont, StrPas(LongFontName))
}

func WinCreateIconIndirect(IconInfo types.ICONINFO) types.HICON {

	return CreateIconIndirect(IconInfo)
}

func WinCreatePalette(LogPalette types.TagLogPalette) types.HPALETTE {

	return CreatePalette(LogPalette)
}

func WinCreatePatternBrush(ABitmap types.HBITMAP) types.HBRUSH {

	return CreatePatternBrush(ABitmap)
}

func WinCreatePenIndirect(LogPen types.TagLogPen) types.HPEN {

	return CreatePenIndirect(LogPen)
}

func WinCreatePolygonRgn(Points types.Point, NumPts types.Integer, FillMode types.Integer) HRGN {

	return CreatePolygonRgn(Points, NumPts, FillMode)
}

func WinDeleteCriticalSection(CritSection types.TCriticalSection) {

	DeleteCriticalSection(CritSection)
}

func WinDeleteDC(hDC types.HDC) types.LongBool {

	return DeleteDC(hDC)
}

func WinDestroyCaret(Handle types.HWND) types.LongBool {

	return DestroyCaret(Handle)
}

func WinDestroyCursor(Handle types.HCURSOR) types.LongBool {

	return DestroyCursor(Handle)
}

func WinDestroyIcon(Handle types.HICON) types.LongBool {

	return DestroyIcon(Handle)
}

func WinDrawFrameControl(DC types.HDC, Rect types.Rect, uType, uState types.Cardinal) types.LongBool {

	return DrawFrameControl(DC, Rect, uType, uState)
}

func WinDrawFocusRect(DC types.HDC, Rect types.Rect) types.LongBool {

	return DrawFocusRect(DC, Rect)
}

func WinDrawEdge(DC types.HDC, Rect types.Rect, edge types.Cardinal, grfFlags types.Cardinal) types.LongBool {

	return DrawEdge(DC, Rect, edge, grfFlags)
}

func WinDrawText(DC types.HDC, Str types.PChar, Count types.Integer, Rect types.Rect, Flags types.Cardinal) types.Integer {

	return DrawText(DC, Str, Count, Rect, Flags)
}

func WinEnableScrollBar(Wnd types.HWND, wSBflags, wArrows types.Cardinal) types.LongBool {

	return EnableScrollBar(Wnd, wSBflags, wArrows)
}

func WinEnableWindow(hWnd types.HWND, bEnable types.LongBool) types.LongBool {

	return EnableWindow(hWnd, bEnable)
}

func WinEnterCriticalSection(CritSection types.TCriticalSection) {

	EnterCriticalSection(CritSection)
}

func WinEnumDisplayMonitors(hdc types.HDC, lprcClip types.Rect, lpfnEnum MonitorEnumProc, dwData types.LPARAM) types.LongBool {

	return EnumDisplayMonitors(hdc, lprcClip, lpfnEnum, dwData)
}

func WinEnumFontFamilies(DC types.HDC, Family types.PChar, EnumFontFamProc FontEnumProc, LParam types.LPARAM) types.LongInt {

	return EnumFontFamilies(DC, Family, EnumFontFamProc, LParam)
}

func WinEnumFontFamiliesEx(DC types.HDC, lpLogFont types.LogFontA, Callback FontEnumExProc, LParam types.LPARAM, Flags types.DWORD) types.LongInt {

	return EnumFontFamiliesEx(DC, lpLogFont, Callback, LParam, Flags)
}

func WinEllipse(DC types.HDC, x1, y1, x2, y2 types.Integer) types.LongBool {

	return Ellipse(DC, x1, y1, x2, y2)
}

func WinEqualRgn(Rgn1 HRGN, Rgn2 HRGN) types.LongBool {

	return EqualRgn(Rgn1, Rgn2)
}

func WinExcludeClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {

	return ExcludeClipRect(dc, Left, Top, Right, Bottom)
}

func WinExtCreatePen(dwPenStyle, dwWidth types.DWORD, lplb types.TagLogBrush, dwStyleCount types.DWORD, lpStyle types.DWORD) types.HPEN {

	return ExtCreatePen(dwPenStyle, dwWidth, lplb, dwStyleCount, lpStyle)
}

func WinExtTextOut(DC types.HDC, X, Y types.Integer, Options types.LongInt, Rect types.Rect, Str types.PChar, Count types.LongInt, Dx types.Integer) types.LongBool {

	return ExtTextOut(DC, X, Y, Options, Rect, Str, Count, Dx)
}

func WinExtSelectClipRGN(dc types.HDC, rgn HRGN, Mode types.LongInt) types.Integer {

	return ExtSelectClipRGN(dc, rgn, Mode)
}

func WinFillRect(DC types.HDC, Rect types.Rect, Brush types.HBRUSH) types.LongBool {

	return FillRect(DC, Rect, Brush)
}

func WinFillRgn(DC types.HDC, RegionHnd HRGN, hbr types.HBRUSH) BOOL {

	return FillRgn(DC, RegionHnd, hbr)
}

func WinFloodFill(DC types.HDC, X, Y types.Integer, Color types.TGraphicsColor, FillStyle types.TGraphicsFillStyle, Brush types.HBRUSH) types.LongBool {

	return FloodFill(DC, X, Y, Color, FillStyle, Brush)
}

func WinFrameRect(DC types.HDC, ARect types.Rect, hBr types.HBRUSH) types.Integer {

	return FrameRect(DC, ARect, hBr)
}

func WinGetActiveWindow() types.HWND {

	return GetActiveWindow()
}

func WinGetBitmapBits(Bitmap types.HBITMAP, Count types.LongInt, Bits types.Pointer) types.LongInt {

	return GetBitmapBits(Bitmap, Count, Bits)
}

func WinGetBkColor(DC types.HDC) types.TColorRef {

	return GetBkColor(DC)
}

func WinGetCapture() types.HWND {

	return GetCapture()
}

func WinGetCaretPos(lpPoint types.Point) types.LongBool {

	return GetCaretPos(lpPoint)
}

func WinGetClientRect(handle types.HWND, Rect types.Rect) types.LongBool {

	return GetClientRect(handle, Rect)
}

func WinGetClipBox(DC types.HDC, lpRect types.Rect) types.LongInt {

	return GetClipBox(DC, lpRect)
}

func WinGetClipRGN(DC types.HDC, RGN HRGN) types.LongInt {

	return GetClipRGN(DC, RGN)
}

func WinGetCurrentObject(DC types.HDC, uObjectType types.UINT) types.HGDIOBJ {

	return GetCurrentObject(DC, uObjectType)
}

func WinGetCursorPos(lpPoint types.Point) types.LongBool {

	return GetCursorPos(lpPoint)
}

func WinGetDC(hWnd types.HWND) types.HDC {

	return GetDC(hWnd)
}

func WinGetDeviceCaps(DC types.HDC, Index types.Integer) types.Integer {

	return GetDeviceCaps(DC, Index)
}

func WinGetDIBits(DC types.HDC, Bitmap types.HBITMAP, StartScan, NumScans types.UINT, Bits types.Pointer,

	BitInfo types.TagBitmapInfo, Usage types.UINT) types.Integer {

	return GetDIBits(DC, Bitmap, StartScan, NumScans, Bits, BitInfo, Usage)
}

func WinGetDoubleClickTime() types.UINT {

	return GetDoubleClickTime()
}

func WinGetFocus() types.HWND {

	return GetFocus()
}

func WinGetFontLanguageInfo(DC types.HDC) types.DWORD {

	return GetFontLanguageInfo(DC)
}

func WinGetForegroundWindow() types.HWND {

	return GetForegroundWindow()
}

func WinGetIconInfo(AIcon types.HICON, AIconInfo types.ICONINFO) types.LongBool {

	return GetIconInfo(AIcon, AIconInfo)
}

func WinGetKeyState(nVirtKey types.Integer) types.Smallint {

	return GetKeyState(nVirtKey)
}

func WinGetMapMode(DC types.HDC) types.Integer {

	return GetMapMode(DC)
}

func WinGetMonitorInfo(hMonitor types.HMONITOR, lpmi types.TagMonitorInfo) types.LongBool {

	return GetMonitorInfo(hMonitor, lpmi)
}

func WinGetDpiForMonitor(hmonitor types.HMONITOR, dpiType types.MONITOR_DPI_TYPE, dpiX types.UINT, dpiY types.UINT) HRESULT { // out

	return GetDpiForMonitor(hmonitor, dpiType, dpiX, dpiY)
}

func WinGetObject(GDIObject types.HGDIOBJ, BufSize types.Integer, Buf types.Pointer) types.Integer {

	return GetObject(GDIObject, BufSize, Buf)
}

func WinGetParent(Handle types.HWND) types.HWND {

	return GetParent(Handle)
}

func WinGetProp(Handle types.HWND, Str types.PChar) types.Pointer {

	return GetProp(Handle, Str)
}

func WinGetRgnBox(RGN HRGN, lpRect types.Rect) types.LongInt {

	return GetRgnBox(RGN, lpRect)
}

func WinGetROP2(DC types.HDC) types.Integer {

	return GetROP2(DC)
}

func WinGetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo types.TagScrollInfo) types.LongBool {

	return GetScrollInfo(Handle, SBStyle, ScrollInfo)
}

func WinGetStockObject(Value types.Integer) types.THandle {

	return GetStockObject(Value)
}

func WinGetSysColor(nIndex types.Integer) types.DWORD {

	return GetSysColor(nIndex)
}

func WinGetSysColorBrush(nIndex types.Integer) types.HBRUSH {

	return GetSysColorBrush(nIndex)
}

func WinGetSystemMetrics(nIndex types.Integer) types.Integer {

	return GetSystemMetrics(nIndex)
}

func WinGetTextColor(DC types.HDC) types.TColorRef {

	return GetTextColor(DC)
}

func WinGetTextExtentExPoint(DC types.HDC, Str types.PChar, Count, MaxWidth types.Integer, MaxCount, PartialWidths types.Integer, Size types.Size) types.LongBool {

	return GetTextExtentExPoint(DC, Str, Count, MaxWidth, MaxCount, PartialWidths, Size)
}

func WinGetTextExtentPoint(DC types.HDC, Str types.PChar, Count types.Integer, Size types.Size) types.LongBool {

	return GetTextExtentPoint(DC, Str, Count, Size)
}

func WinGetTextExtentPoint32(DC types.HDC, Str types.PChar, Count types.Integer, Size types.Size) types.LongBool {

	return GetTextExtentPoint32(DC, Str, Count, Size)
}

func WinGetTextMetrics(DC types.HDC, TM types.TagTextMetricA) types.LongBool {

	return GetTextMetrics(DC, TM)
}

func WinGetViewPortExtEx(DC types.HDC, Size types.Size) types.Integer {

	return GetViewPortExtEx(DC, Size)
}

func WinGetViewPortOrgEx(DC types.HDC, P types.Point) types.Integer {

	return GetViewPortOrgEx(DC, P)
}

func WinGetWindowExtEx(DC types.HDC, Size types.Size) types.Integer {

	return GetWindowExtEx(DC, Size)
}

func WinGetWindowLong(Handle types.HWND, int types.Integer) types.PtrInt {

	return GetWindowLong(Handle, int)
}

func WinGetWindowRect(Handle types.HWND, Rect types.Rect) types.Integer {

	return GetWindowRect(Handle, Rect)
}

func WinGetWindowSize(Handle types.HWND, Width, Height types.Integer) types.LongBool {

	return GetWindowSize(Handle, Width, Height)
}

func WinGetWindowOrgEx(dc types.HDC, P types.Point) types.Integer { // because of delphi compatibility

	return GetWindowOrgEx(dc, P)
}

func WinGradientFill(DC types.HDC, Vertices types.TagTriVertex, NumVertices types.LongInt, Meshes types.Pointer, NumMeshes types.LongInt, Mode types.LongInt) types.LongBool {

	return GradientFill(DC, Vertices, NumVertices, Meshes, NumMeshes, Mode)
}

func WinHideCaret(hWnd types.HWND) types.LongBool {

	return HideCaret(hWnd)
}

func WinInitializeCriticalSection(CritSection types.TCriticalSection) {

	InitializeCriticalSection(CritSection)
}

func WinIntersectClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {

	return IntersectClipRect(dc, Left, Top, Right, Bottom)
}

func WinInvalidateRect(aHandle types.HWND, ARect types.Rect, bErase types.LongBool) types.LongBool {

	return InvalidateRect(aHandle, ARect, bErase)
}

func WinInvalidateRgn(Handle types.HWND, Rgn HRGN, Erase types.LongBool) types.LongBool {

	return InvalidateRgn(Handle, Rgn, Erase)
}

func WinIsDBCSLeadByte(TestChar byte) types.LongBool {

	return IsDBCSLeadByte(TestChar)
}

func WinIsIconic(handle types.HWND) types.LongBool {

	return IsIconic(handle)
}

func WinIsWindow(handle types.HWND) types.LongBool {

	return IsWindow(handle)
}

func WinIsWindowEnabled(handle types.HWND) types.LongBool {

	return IsWindowEnabled(handle)
}

func WinIsWindowVisible(handle types.HWND) types.LongBool {

	return IsWindowVisible(handle)
}

func WinIsZoomed(handle types.HWND) types.LongBool {

	return IsZoomed(handle)
}

func WinLeaveCriticalSection(CritSection types.TCriticalSection) {

	LeaveCriticalSection(CritSection)
}

func WinLineTo(DC types.HDC, X, Y types.Integer) types.LongBool {

	return LineTo(DC, X, Y)
}

func WinLoadBitmap(hInstance types.THandle, lpBitmapName types.PChar) types.HBITMAP {

	return LoadBitmap(hInstance, lpBitmapName)
}

func WinLoadIcon(hInstance types.THandle, lpIconName types.PChar) types.HICON {

	return LoadIcon(hInstance, lpIconName)
}

func WinMaskBltRop(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask types.HBITMAP, XMask, YMask types.Integer, Rop types.DWORD) types.LongBool {

	return MaskBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, Mask, XMask, YMask, Rop)
}

func WinMaskBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask types.HBITMAP, XMask, YMask types.Integer) types.LongBool {

	return MaskBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, Mask, XMask, YMask)
}

func WinMessageBox(hWnd types.HWND, lpText, lpCaption types.PChar, uType types.Cardinal) types.Integer { //def MB_OK

	return MessageBox(hWnd, lpText, lpCaption, uType)
}

func WinMonitorFromPoint(ptScreenCoords types.Point, dwFlags types.DWORD) types.HMONITOR {

	return MonitorFromPoint(ptScreenCoords, dwFlags)
}

func WinMonitorFromRect(lprcScreenCoords types.Rect, dwFlags types.DWORD) types.HMONITOR {

	return MonitorFromRect(lprcScreenCoords, dwFlags)
}

func WinMonitorFromWindow(hWnd types.HWND, dwFlags types.DWORD) types.HMONITOR {

	return MonitorFromWindow(hWnd, dwFlags)
}

func WinMoveToEx(DC types.HDC, X, Y types.Integer, OldPoint types.Point) types.LongBool {

	return MoveToEx(DC, X, Y, OldPoint)
}

func WinOffsetRgn(RGN HRGN, nXOffset, nYOffset types.Integer) types.Integer {

	return OffsetRgn(RGN, nXOffset, nYOffset)
}

func WinPaintRgn(DC types.HDC, RGN HRGN) types.LongBool {

	return PaintRgn(DC, RGN)
}

func WinPie(DC types.HDC, x1, y1, x2, y2, sx, sy, ex, ey types.Integer) types.LongBool {

	return Pie(DC, x1, y1, x2, y2, sx, sy, ex, ey)
}

func WinPolyBezier(DC types.HDC, Points types.Point, NumPts types.Integer, Filled, Continuous types.LongBool) types.LongBool {

	return PolyBezier(DC, Points, NumPts, Filled, Continuous)
}

func WinPolygon(DC types.HDC, Points types.Point, NumPts types.Integer, Winding types.LongBool) types.LongBool {

	return Polygon(DC, Points, NumPts, Winding)
}

func WinPolyline(DC types.HDC, Points types.Point, NumPts types.Integer) types.LongBool {

	return Polyline(DC, Points, NumPts)
}

func WinPostMessage(Handle types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LongBool {

	return PostMessage(Handle, Msg, WParam, LParam)
}

func WinRealizePalette(DC types.HDC) types.Cardinal {

	return RealizePalette(DC)
}

func WinRectangle(DC types.HDC, X1, Y1, X2, Y2 types.Integer) types.LongBool {

	return Rectangle(DC, X1, Y1, X2, Y2)
}

func WinRectInRegion(RGN HRGN, ARect types.Rect) types.LongBool {

	return RectInRegion(RGN, ARect)
}

func WinRectVisible(DC types.HDC, ARect types.Rect) types.LongBool {

	return RectVisible(DC, ARect)
}

func WinRedrawWindow(Wnd types.HWND, lprcUpdate types.Rect, hrgnUpdate HRGN, flags types.UINT) types.LongBool {

	return RedrawWindow(Wnd, lprcUpdate, hrgnUpdate, flags)
}

func WinReleaseCapture() types.LongBool {

	return ReleaseCapture()
}

func WinReleaseDC(hWnd types.HWND, DC types.HDC) types.Integer {

	return ReleaseDC(hWnd, DC)
}

func WinRemoveProp(Handle types.HWND, Str types.PChar) types.THandle {

	return RemoveProp(Handle, Str)
}

func WinRestoreDC(DC types.HDC, SavedDC types.Integer) types.LongBool {

	return RestoreDC(DC, SavedDC)
}

func WinRoundRect(DC types.HDC, X1, Y1, X2, Y2 types.Integer, RX, RY types.Integer) types.LongBool {

	return RoundRect(DC, X1, Y1, X2, Y2, RX, RY)
}

func WinSaveDC(DC types.HDC) types.Integer {

	return SaveDC(DC)
}

func WinScrollWindowEx(hWnd types.HWND, dx, dy types.Integer, prcScroll, prcClip types.Rect, hrgnUpdate HRGN, prcUpdate types.Rect, flags types.UINT) types.LongBool {

	return ScrollWindowEx(hWnd, dx, dy, prcScroll, prcClip, hrgnUpdate, prcUpdate, flags)
}

func WinSelectClipRGN(DC types.HDC, RGN HRGN) types.LongInt {

	return SelectClipRGN(DC, RGN)
}

func WinSelectObject(DC types.HDC, GDIObj types.HGDIOBJ) types.HGDIOBJ {

	return SelectObject(DC, GDIObj)
}

func WinSelectPalette(DC types.HDC, Palette types.HPALETTE, ForceBackground types.LongBool) types.HPALETTE {

	return SelectPalette(DC, Palette, ForceBackground)
}

func WinSendMessage(HandleWnd types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LResult {

	return SendMessage(HandleWnd, Msg, WParam, LParam)
}

func WinSetActiveWindow(Handle types.HWND) types.HWND {

	return SetActiveWindow(Handle)
}

func WinSetBkColor(DC types.HDC, Color types.TColorRef) types.TColorRef { //pbd

	return SetBkColor(DC, Color)
}

func WinSetBkMode(DC types.HDC, bkMode types.Integer) types.Integer {

	return SetBkMode(DC, bkMode)
}

func WinSetCapture(AHandle types.HWND) types.HWND {

	return SetCapture(AHandle)
}

func WinSetCaretPos(X, Y types.Integer) types.LongBool {

	return SetCaretPos(X, Y)
}

func WinSetCaretPosEx(handle types.HWND, X, Y types.Integer) types.LongBool {

	return SetCaretPosEx(handle, X, Y)
}

func WinSetCursorPos(X, Y types.Integer) types.LongBool {

	return SetCursorPos(X, Y)
}

func WinSetFocus(hWnd types.HWND) types.HWND {

	return SetFocus(hWnd)
}

func WinSetForegroundWindow(hWnd types.HWND) types.LongBool {

	return SetForegroundWindow(hWnd)
}

func WinSetMapMode(DC types.HDC, fnMapMode types.Integer) types.Integer {

	return SetMapMode(DC, fnMapMode)
}

func WinSetMenu(AWindowHandle types.HWND, AMenuHandle types.HMENU) types.LongBool {

	return SetMenu(AWindowHandle, AMenuHandle)
}

func WinSetParent(hWndChild types.HWND, hWndParent types.HWND) types.HWND {

	return SetParent(hWndChild, hWndParent)
}

func WinSetProp(Handle types.HWND, Str types.PChar, Data types.Pointer) types.LongBool {

	return SetProp(Handle, Str, Data)
}

func WinSetROP2(DC types.HDC, Mode types.Integer) types.Integer {

	return SetROP2(DC, Mode)
}

func WinSetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo types.TagScrollInfo, Redraw types.LongBool) types.Integer {

	return SetScrollInfo(Handle, SBStyle, ScrollInfo, Redraw)
}

func WinSetStretchBltMode(DC types.HDC, StretchMode types.Integer) types.Integer {

	return SetStretchBltMode(DC, StretchMode)
}

func WinSetTextCharacterExtra(_hdc types.HDC, nCharExtra types.Integer) types.Integer {

	return SetTextCharacterExtra(_hdc, nCharExtra)
}

func WinSetTextColor(DC types.HDC, Color types.TColorRef) types.TColorRef {

	return SetTextColor(DC, Color)
}

func WinSetWindowLong(Handle types.HWND, Idx types.Integer, NewLong types.PtrInt) types.PtrInt {

	return SetWindowLong(Handle, Idx, NewLong)
}

func WinSetViewPortExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize types.Size) types.LongBool {

	return SetViewPortExtEx(DC, XExtent, YExtent, OldSize)
}

func WinSetViewPortOrgEx(DC types.HDC, NewX, NewY types.Integer, OldPoint types.Point) types.LongBool {

	return SetViewPortOrgEx(DC, NewX, NewY, OldPoint)
}

func WinSetWindowExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize types.Size) types.LongBool {

	return SetWindowExtEx(DC, XExtent, YExtent, OldSize)
}

func WinSetWindowOrgEx(dc types.HDC, NewX, NewY types.Integer, OldPoint types.Point) types.LongBool {

	return SetWindowOrgEx(dc, NewX, NewY, OldPoint)
}

func WinSetWindowPos(hWnd types.HWND, hWndInsertAfter types.HWND, X, Y, cx, cy types.Integer, uFlags types.UINT) types.LongBool {

	return SetWindowPos(hWnd, hWndInsertAfter, X, Y, cx, cy, uFlags)
}

func WinShowCaret(hWnd types.HWND) types.LongBool {

	return ShowCaret(hWnd)
}

func WinShowScrollBar(Handle types.HWND, wBar types.Integer, bShow types.LongBool) types.LongBool {

	return ShowScrollBar(Handle, wBar, bShow)
}

func WinShowWindow(hWnd types.HWND, nCmdShow types.Integer) types.LongBool {

	return ShowWindow(hWnd, nCmdShow)
}

func WinStretchBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc, SrcWidth, SrcHeight types.Integer, Rop types.Cardinal) types.LongBool {

	return StretchBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, SrcWidth, SrcHeight, Rop)
}

func WinStretchDIBits(DC types.HDC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight types.Integer, Bits types.Pointer,

	BitsInfo types.TagBitmapInfo, Usage types.UINT, Rop types.DWORD) types.Integer {

	return StretchDIBits(DC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight, Bits, BitsInfo, Usage, Rop)
}

func WinSystemParametersInfo(uiAction types.DWORD, uiParam types.DWORD, pvParam types.Pointer, fWinIni types.DWORD) types.LongBool {

	return SystemParametersInfo(uiAction, uiParam, pvParam, fWinIni)
}

func WinTextOut(DC types.HDC, X, Y types.Integer, Str types.PChar, Count types.Integer) types.LongBool {

	return TextOut(DC, X, Y, Str, Count)
}

func WinUpdateWindow(Handle types.HWND) types.LongBool {

	return UpdateWindow(Handle)
}

func WinWindowFromPoint(Point types.Point) types.HWND {

	return WindowFromPoint(Point)
}
