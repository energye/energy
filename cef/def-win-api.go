package cef

import "github.com/energye/energy/v2/types"

func WinEndPaint(Handle types.HWND, PS TPaintStruct) types.Integer {

	return EndPaint(Handle, PS)
}

func WinPaint(Handle types.HWND, PS TPaintStruct) types.HDC {

	return Paint(Handle, PS)
}

func WinArc(DC types.HDC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength types.Integer) types.LongBool {

	return Arc(DC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength)
}

func WinAngleChord(DC types.HDC, x1, y1, x2, y2, angle1, angle2 types.Integer) types.LongBool {

	return AngleChord(DC, x1, y1, x2, y2, angle1, angle2)
}

func WinCallNextHookEx(hhk HHOOK, ncode types.Integer, WParam types.WPARAM, LParam types.LPARAM) types.Integer {

	return CallNextHookEx(hhk, ncode, WParam, LParam)
}

func WinCallWindowProc(lpPrevWndFunc TFarProc, Handle types.HWND, Msg UINT, WParam types.WPARAM, LParam types.LPARAM) types.Integer {

	return CallWindowProc(lpPrevWndFunc, Handle, Msg, WParam, LParam)
}

func WinBitBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Rop DWORD) types.LongBool {

	return BitBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, Rop)
}

func WinCreateBitmap(Width, Height types.Integer, Planes, BitCount Longint, BitmapBits Pointer) HBITMAP {

	return CreateBitmap(Width, Height, Planes, BitCount, BitmapBits)
}

func WinCreateBrushIndirect(LogBrush TLogBrush) HBRUSH {

	return CreateBrushIndirect(LogBrush)
}

func WinCreateBrushWithRadialGradient(LogBrush TLogRadialGradient) HBRUSH {

	return CreateBrushWithRadialGradient(LogBrush)
}

func WinCreateCaret(Handle types.HWND, Bitmap HBITMAP, width, Height types.Integer) types.LongBool {

	return CreateCaret(Handle, Bitmap, width, Height)
}

func WinCreateCompatibleBitmap(DC types.HDC, Width, Height types.Integer) HBITMAP {

	return CreateCompatibleBitmap(DC, Width, Height)
}

func WinCreateCompatibleDC(DC types.HDC) types.HDC {

	return CreateCompatibleDC(DC)
}

func WinCreateDIBitmap(DC types.HDC, InfoHeader TBitmapInfoHeader, dwUsage DWORD, InitBits PChar, InitInfo TBitmapInfo, wUsage UINT) HBITMAP {

	return CreateDIBitmap(DC, InfoHeader, dwUsage, InitBits, InitInfo, wUsage)
}

func WinCreateDIBSection(DC types.HDC, BitmapInfo tagBitmapInfo, Usage UINT, Bits Pointer, SectionHandle THandle, Offset DWORD) HBITMAP {

	return CreateDIBSection(DC, BitmapInfo, Usage, Bits, SectionHandle, Offset)
}

func WinCreateEllipticRgn(X1, Y1, X2, Y2 types.Integer) HRGN {

	return CreateEllipticRgn(X1, Y1, X2, Y2)
}

func WinCreateFontIndirect(LogFont TLogFont) HFONT {

	return CreateFontIndirect(LogFont)
}

func WinCreateFontIndirectEx(LogFont TLogFont, LongFontName PChar) HFONT {

	return CreateFontIndirectEx(LogFont, StrPas(LongFontName))
}

func WinCreateIconIndirect(IconInfo PIconInfo) HICON {

	return CreateIconIndirect(IconInfo)
}

func WinCreatePalette(LogPalette TLogPalette) HPalette {

	return CreatePalette(LogPalette)
}

func WinCreatePatternBrush(ABitmap HBITMAP) HBRUSH {

	return CreatePatternBrush(ABitmap)
}

func WinCreatePenIndirect(LogPen TLogPen) HPEN {

	return CreatePenIndirect(LogPen)
}

func WinCreatePolygonRgn(Points PPoint, NumPts types.Integer, FillMode types.Integer) HRGN {

	return CreatePolygonRgn(Points, NumPts, FillMode)
}

func WinDeleteCriticalSection(CritSection TCriticalSection) {

	DeleteCriticalSection(CritSection)
}

func WinDeleteDC(hDC types.HDC) types.LongBool {

	return DeleteDC(hDC)
}

func WinDestroyCaret(Handle types.HWND) types.LongBool {

	return DestroyCaret(Handle)
}

func WinDestroyCursor(Handle HCURSOR) types.LongBool {

	return DestroyCursor(Handle)
}

func WinDestroyIcon(Handle HICON) types.LongBool {

	return DestroyIcon(Handle)
}

func WinDrawFrameControl(DC types.HDC, Rect TRect, uType, uState Cardinal) types.LongBool {

	return DrawFrameControl(DC, Rect, uType, uState)
}

func WinDrawFocusRect(DC types.HDC, Rect TRect) types.LongBool {

	return DrawFocusRect(DC, Rect)
}

func WinDrawEdge(DC types.HDC, Rect TRect, edge Cardinal, grfFlags Cardinal) types.LongBool {

	return DrawEdge(DC, Rect, edge, grfFlags)
}

func WinDrawText(DC types.HDC, Str PChar, Count types.Integer, Rect TRect, Flags Cardinal) types.Integer {

	return DrawText(DC, Str, Count, Rect, Flags)
}

func WinEnableScrollBar(Wnd types.HWND, wSBflags, wArrows Cardinal) types.LongBool {

	return EnableScrollBar(Wnd, wSBflags, wArrows)
}

func WinEnableWindow(hWnd types.HWND, bEnable types.LongBool) types.LongBool {

	return EnableWindow(hWnd, bEnable)
}

func WinEnterCriticalSection(CritSection TCriticalSection) {

	EnterCriticalSection(CritSection)
}

func WinEnumDisplayMonitors(hdc types.HDC, lprcClip PRect, lpfnEnum MonitorEnumProc, dwData types.LPARAM) types.LongBool {

	return EnumDisplayMonitors(hdc, lprcClip, lpfnEnum, dwData)
}

func WinEnumFontFamilies(DC types.HDC, Family Pchar, EnumFontFamProc FontEnumProc, LParam types.LPARAM) longint {

	return EnumFontFamilies(DC, Family, EnumFontFamProc, LParam)
}

func WinEnumFontFamiliesEx(DC types.HDC, lpLogFontPLogFont, Callback FontEnumExProc, LParam types.LPARAM, Flags dword) longint {

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

func WinExtCreatePen(dwPenStyle, dwWidth DWord, lplb TLogBrush, dwStyleCount DWord, lpStyle PDWord) HPEN {

	return ExtCreatePen(dwPenStyle, dwWidth, lplb, dwStyleCount, lpStyle)
}

func WinExtTextOut(DC types.HDC, X, Y types.Integer, Options Longint, Rect PRect, Str PChar, Count Longint, Dx Ptypes.Integer) types.LongBool,  {

return ExtTextOut(DC, X, Y, Options, Rect, Str, Count, Dx)
}

func WinExtSelectClipRGN(dc types.HDC, rgn hrgn, Mode Longint) types.Integer,  {

return ExtSelectClipRGN(dc, rgn, Mode)
}

func WinFillRect(DC types.HDC, Rect TRect, Brush HBRUSH) types.LongBool {

	return FillRect(DC, Rect, Brush)
}

func WinFillRgn(DC types.HDC, RegionHnd HRGN, hbr HBRUSH) BOOL {

	return FillRgn(DC, RegionHnd, hbr)
}

func WinFloodFill(DC types.HDC, X, Y types.Integer, Color TGraphicsColor, FillStyle TGraphicsFillStyle, Brush HBRUSH) types.LongBool {

	return FloodFill(DC, X, Y, Color, FillStyle, Brush)
}

func WinFrameRect(DC types.HDC, ARect TRect, hBr HBRUSH) types.Integer {

	return FrameRect(DC, ARect, hBr)
}

func WinGetActiveWindow() types.HWND {

	return GetActiveWindow()
}

func WinGetBitmapBits(Bitmap HBITMAP, Count Longint, Bits Pointer) Longint {

	return GetBitmapBits(Bitmap, Count, Bits)
}

func WinGetBkColor(DC types.HDC) TColorRef {

	return GetBkColor(DC)
}

func WinGetCapture() types.HWND {

	return GetCapture()
}

func WinGetCaretPos(lpPoint TPoint) types.LongBool {

	return GetCaretPos(lpPoint)
}

func WinGetClientRect(handle types.HWND, Rect TRect) types.LongBool {

	return GetClientRect(handle, Rect)
}

func WinGetClipBox(DC types.HDC, lpRect PRect) Longint {

	return GetClipBox(DC, lpRect)
}

func WinGetClipRGN(DC types.HDC, RGN hRGN) Longint {

	return GetClipRGN(DC, RGN)
}

func WinGetCurrentObject(DC types.HDC, uObjectType UINT) HGDIOBJ {

	return GetCurrentObject(DC, uObjectType)
}

func WinGetCursorPos(lpPoint TPoint) types.LongBool {

	return GetCursorPos(lpPoint)
}

func WinGetDC(hWnd types.HWND) types.HDC {

	return GetDC(hWnd)
}

func WinGetDeviceCaps(DC types.HDC, Index types.Integer) types.Integer {

	return GetDeviceCaps(DC, Index)
}

func WinGetDIBits(DC types.HDC, Bitmap HBitmap, StartScan, NumScans UINT, Bits Pointer,

	BitInfo BitmapInfo, Usage UINT) types.Integer {

	return GetDIBits(DC, Bitmap, StartScan, NumScans, Bits, BitInfo, Usage)
}

func WinGetDoubleClickTime() UINT {

	return GetDoubleClickTime()
}

func WinGetFocus() types.HWND {

	return GetFocus()
}

func WinGetFontLanguageInfo(DC types.HDC) DWord {

	return GetFontLanguageInfo(DC)
}

func WinGetForegroundWindow types.HWND {

	return GetForegroundWindow()
}

func WinGetIconInfo(AIcon HICON, AIconInfo PIconInfo) types.LongBool {

	return GetIconInfo(AIcon, AIconInfo)
}

func WinGetKeyState(nVirtKey types.Integer) Smallint {

	return GetKeyState(nVirtKey)
}

func WinGetMapMode(DC types.HDC) types.Integer {

	return GetMapMode(DC)
}

func WinGetMonitorInfo(hMonitor HMONITOR, lpmi PMonitorInfo) types.LongBool {

	return GetMonitorInfo(hMonitor, lpmi)
}

func WinGetDpiForMonitor(hmonitor HMONITOR, dpiType TMonitorDpiType, out dpiX UINT, out dpiY UINT) HRESULT {

	return GetDpiForMonitor(hmonitor, dpiType, dpiX, dpiY)
}

func WinGetObject(GDIObject HGDIOBJ, BufSize types.Integer, Buf Pointer) types.Integer {

	return GetObject(GDIObject, BufSize, Buf)
}

func WinGetParent(Handle types.HWND) types.HWND {

	return GetParent(Handle)
}

func WinGetProp(Handle types.HWND, Str PChar) Pointer {

	return GetProp(Handle, Str)
}

func WinGetRgnBox(RGN HRGN, lpRect PRect) Longint {

	return GetRgnBox(RGN, lpRect)
}

func WinGetROP2(DC types.HDC) types.Integer {

	return GetROP2(DC)
}

func WinGetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo TScrollInfo) types.LongBool {

	return GetScrollInfo(Handle, SBStyle, ScrollInfo)
}

func WinGetStockObject(Value types.Integer) THandle {

	return GetStockObject(Value)
}

func WinGetSysColor(nIndex types.Integer) DWORD {

	return GetSysColor(nIndex)
}

func WinGetSysColorBrush(nIndex types.Integer) HBrush {

	return GetSysColorBrush(nIndex)
}

func WinGetSystemMetrics(nIndex types.Integer) types.Integer {

	return GetSystemMetrics(nIndex)
}

func WinGetTextColor(DC types.HDC) TColorRef {

	return GetTextColor(DC)
}

func WinGetTextExtentExPoint(DC types.HDC, Str PChar, Count, MaxWidth types.Integer, MaxCount, PartialWidths Ptypes.Integer, Size TSize) types.LongBool {

	return GetTextExtentExPoint(DC, Str, Count, MaxWidth, MaxCount, PartialWidths, Size)
}

func WinGetTextExtentPoint(DC types.HDC, Str PChar, Count types.Integer, Size TSize) types.LongBool {

	return GetTextExtentPoint(DC, Str, Count, Size)
}

func WinGetTextExtentPoint32(DC types.HDC, Str PChar, Count types.Integer, Size TSize) types.LongBool {

	return GetTextExtentPoint32(DC, Str, Count, Size)
}

func WinGetTextMetrics(DC types.HDC, TM TTextMetric) types.LongBool,  {

return GetTextMetrics(DC, TM)
}

func WinGetViewPortExtEx(DC types.HDC, Size PSize) types.Integer {

	return GetViewPortExtEx(DC, Size)
}

func WinGetViewPortOrgEx(DC types.HDC, P PPoint) types.Integer {

	return GetViewPortOrgEx(DC, P)
}

func WinGetWindowExtEx(DC types.HDC, Size PSize) types.Integer {

	return GetWindowExtEx(DC, Size)
}

func WinGetWindowLong(Handle types.HWND, int types.Integer) PtrInt {

	return GetWindowLong(Handle, int)
}

func WinGetWindowRect(Handle types.HWND, Rect TRect) types.Integer {

	return GetWindowRect(Handle, Rect)
}

func WinGetWindowSize(Handle types.HWND, Width, Height types.Integer) types.LongBool {

	return GetWindowSize(Handle, Width, Height)
}

func WinGetWindowOrgEx(dc types.HDC, P TPoint) types.Integer { // because of delphi compatibility

	return GetWindowOrgEx(dc, P)
}

func WinGradientFill(DC types.HDC, Vertices PTriVertex, NumVertices Longint, Meshes Pointer, NumMeshes Longint, Mode Longint) types.LongBool {

	return GradientFill(DC, Vertices, NumVertices, Meshes, NumMeshes, Mode)
}

func WinHideCaret(hWnd types.HWND) types.LongBool {

	return HideCaret(hWnd)
}

func WinInitializeCriticalSection(CritSection TCriticalSection) {

	InitializeCriticalSection(CritSection)
}

func WinIntersectClipRect(dc types.HDC, Left, Top, Right, Bottom types.Integer) types.Integer {

	return IntersectClipRect(dc, Left, Top, Right, Bottom)
}

func WinInvalidateRect(aHandle types.HWND, ARect pRect, bErase types.LongBool) types.LongBool {

	return InvalidateRect(aHandle, ARect, bErase)
}

func WinInvalidateRgn(Handle types.HWND, Rgn HRGN, Erase types.LongBool) types.LongBool {

	return InvalidateRgn(Handle, Rgn, Erase)
}

func WinIsDBCSLeadByte(TestChar Byte) types.LongBool {

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

func WinLeaveCriticalSection(CritSection TCriticalSection) {

	LeaveCriticalSection(CritSection)
}

func WinLineTo(DC types.HDC, X, Y types.Integer) types.LongBool {

	return LineTo(DC, X, Y)
}

func WinLoadBitmap(hInstance THandle, lpBitmapName PChar) HBitmap {

	return LoadBitmap(hInstance, lpBitmapName)
}

func WinLoadIcon(hInstance THandle, lpIconName PChar) HIcon {

	return LoadIcon(hInstance, lpIconName)
}

func WinMaskBltRop(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask HBITMAP, XMask, YMask types.Integer, Rop DWORD) types.LongBool {

	return MaskBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, Mask, XMask, YMask, Rop)
}

func WinMaskBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc types.Integer, Mask HBITMAP, XMask, YMask types.Integer) types.LongBool {

	return MaskBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, Mask, XMask, YMask)
}

func WinMessageBox(hWnd types.HWND, lpText, lpCaption PChar, uType Cardinal = MB_OK) types.Integer {

	return MessageBox(hWnd, lpText, lpCaption, uType)
}

func WinMonitorFromPoint(ptScreenCoords TPoint, dwFlags DWord) HMONITOR {

	return MonitorFromPoint(ptScreenCoords, dwFlags)
}

func WinMonitorFromRect(lprcScreenCoords PRect, dwFlags DWord) HMONITOR {

	return MonitorFromRect(lprcScreenCoords, dwFlags)
}

func WinMonitorFromWindow(hWnd types.HWND, dwFlags DWord) HMONITOR {

	return MonitorFromWindow(hWnd, dwFlags)
}

func WinMoveToEx(DC types.HDC, X, Y types.Integer, OldPoint PPoint) types.LongBool {

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

func WinPolyBezier(DC types.HDC, Points PPoint, NumPts types.Integer, Filled, Continuous types.LongBool) types.LongBool {

	return PolyBezier(DC, Points, NumPts, Filled, Continuous)
}

func WinPolygon(DC types.HDC, Points PPoint, NumPts types.Integer, Winding types.LongBool) types.LongBool {

	return Polygon(DC, Points, NumPts, Winding)
}

func WinPolyline(DC types.HDC, Points PPoint, NumPts types.Integer) types.LongBool {

	return Polyline(DC, Points, NumPts)
}

func WinPostMessage(Handle types.HWND, Msg Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LongBool {

	return PostMessage(Handle, Msg, WParam, LParam)
}

func WinRealizePalette(DC types.HDC) Cardinal {

	return RealizePalette(DC)
}

func WinRectangle(DC types.HDC, X1, Y1, X2, Y2 types.Integer) types.LongBool {

	return Rectangle(DC, X1, Y1, X2, Y2)
}

func WinRectInRegion(RGN HRGN, ARect TRect) types.LongBool {

	return RectInRegion(RGN, ARect)
}

func WinRectVisible(DC types.HDC, ARect TRect) types.LongBool {

	return RectVisible(DC, ARect)
}

func WinRedrawWindow(Wnd types.HWND, lprcUpdate PRECT, hrgnUpdate HRGN, flags UINT) types.LongBool {

	return RedrawWindow(Wnd, lprcUpdate, hrgnUpdate, flags)
}

func WinReleaseCapture() types.LongBool {

	return ReleaseCapture()
}

func WinReleaseDC(hWnd types.HWND, DC types.HDC) types.Integer {

	return ReleaseDC(hWnd, DC)
}

func WinRemoveProp(Handle types.HWND, Str PChar) THandle {

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

func WinScrollWindowEx(hWnd types.HWND, dx, dy types.Integer, prcScroll, prcClip PRect, hrgnUpdate HRGN, prcUpdate PRect, flags UINT) types.LongBool {

	return ScrollWindowEx(hWnd, dx, dy, prcScroll, prcClip, hrgnUpdate, prcUpdate, flags)
}

func WinSelectClipRGN(DC types.HDC, RGN HRGN) Longint {

	return SelectClipRGN(DC, RGN)
}

func WinSelectObject(DC types.HDC, GDIObj HGDIOBJ) HGDIOBJ {

	return SelectObject(DC, GDIObj)
}

func WinSelectPalette(DC types.HDC, Palette HPALETTE, ForceBackground types.LongBool) HPALETTE {

	return SelectPalette(DC, Palette, ForceBackground)
}

func WinSendMessage(HandleWnd types.HWND, Msg Cardinal, WParam types.WPARAM, LParam types.LPARAM) LResult {

	return SendMessage(HandleWnd, Msg, WParam, LParam)
}

func WinSetActiveWindow(Handle types.HWND) types.HWND, {

return SetActiveWindow(Handle)
}

func WinSetBkColor(DC types.HDC, Color TColorRef) TColorRef { //pbd

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

func WinSetMenu(AWindowHandle types.HWND, AMenuHandle HMENU) types.LongBool {

	return SetMenu(AWindowHandle, AMenuHandle)
}

func WinSetParent(hWndChild types.HWND, hWndParent types.HWND) types.HWND {

	return SetParent(hWndChild, hWndParent)
}

func WinSetProp(Handle types.HWND, Str PChar, Data Pointer) types.LongBool {

	return SetProp(Handle, Str, Data)
}

func WinSetROP2(DC types.HDC, Mode types.Integer) types.Integer {

	return SetROP2(DC, Mode)
}

func WinSetScrollInfo(Handle types.HWND, SBStyle types.Integer, ScrollInfo TScrollInfo, Redraw types.LongBool) types.Integer {

	return SetScrollInfo(Handle, SBStyle, ScrollInfo, Redraw)
}

func WinSetStretchBltMode(DC types.HDC, StretchMode types.Integer) types.Integer {

	return SetStretchBltMode(DC, StretchMode)
}

func WinSetTextCharacterExtra(_hdc types.HDC, nCharExtra types.Integer) types.Integer {

	return SetTextCharacterExtra(_hdc, nCharExtra)
}

func WinSetTextColor(DC types.HDC, Color TColorRef) TColorRef {

	return SetTextColor(DC, Color)
}

func WinSetWindowLong(Handle types.HWND, Idx types.Integer, NewLong PtrInt) PtrInt {

	return SetWindowLong(Handle, Idx, NewLong)
}

func WinSetViewPortExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize PSize) types.LongBool {

	return SetViewPortExtEx(DC, XExtent, YExtent, OldSize)
}

func WinSetViewPortOrgEx(DC types.HDC, NewX, NewY types.Integer, OldPoint PPoint) types.LongBool {

	return SetViewPortOrgEx(DC, NewX, NewY, OldPoint)
}

func WinSetWindowExtEx(DC types.HDC, XExtent, YExtent types.Integer, OldSize PSize) types.LongBool {

	return SetWindowExtEx(DC, XExtent, YExtent, OldSize)
}

func WinSetWindowOrgEx(dc types.HDC, NewX, NewY types.Integer, OldPoint PPoint) types.LongBool {

	return SetWindowOrgEx(dc, NewX, NewY, OldPoint)
}

func WinSetWindowPos(hWnd types.HWND, hWndInsertAfter types.HWND, X, Y, cx, cy types.Integer, uFlags UINT) types.LongBool {

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

func WinStretchBlt(DestDC types.HDC, X, Y, Width, Height types.Integer, SrcDC types.HDC, XSrc, YSrc, SrcWidth, SrcHeight types.Integer, Rop Cardinal) types.LongBool {

	return StretchBlt(DestDC, X, Y, Width, Height, SrcDC, XSrc, YSrc, SrcWidth, SrcHeight, Rop)
}

func WinStretchDIBits(DC types.HDC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight types.Integer, Bits Pointer,

	BitsInfo TBitmapInfo, Usage UINT, Rop DWORD) types.Integer {

	return StretchDIBits(DC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight, Bits, BitsInfo, Usage, Rop)
}

func WinSystemParametersInfo(uiAction DWord, uiParam DWord, pvParam Pointer, fWinIni DWord) types.LongBool {

	return SystemParametersInfo(uiAction, uiParam, pvParam, fWinIni)
}

func WinTextOut(DC types.HDC, X, Y types.Integer, Str Pchar, Count types.Integer) types.LongBool {

	return TextOut(DC, X, Y, Str, Count)
}

func WinUpdateWindow(Handle types.HWND) types.LongBool {

	return UpdateWindow(Handle)
}

func WinWindowFromPoint(Point TPoint) types.HWND {

	return WindowFromPoint(Point)
}
