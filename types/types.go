//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package types CEF & Go type mapping
package types

import (
	"github.com/energye/golcl/lcl/api"
	lcltypes "github.com/energye/golcl/lcl/types"
	"unsafe"
)

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Int int
type UInt8 uint8
type UInt16 uint16
type UInt32 uint32
type UInt64 uint64
type UInt uint
type UIntptr uintptr
type String string
type Boolean bool
type Float32 float32
type Float64 float64
type HDWP uintptr
type Single = Float32
type PChar = String
type Char = byte
type AnsiChar = Char

// 32-bit ARGB color value, not premultiplied. The color components are always  in a known order. Equivalent to the SkColor type.
type TCefColor = UInt32
type Integer = Int32
type LongInt = Int32
type LongPtr = UIntptr
type LongWord = UInt32
type NativeUInt = UInt32
type TCefString = String
type Cardinal = UInt32
type LongBool = Boolean
type BOOL = LongBool
type DWORD = Cardinal
type TCefWindowHandle = HWND
type HWND = UIntptr
type WPARAM = UIntptr
type LPARAM = UIntptr
type HDC = UIntptr
type UINT = LongWord
type Pointer = UIntptr
type QWord = UIntptr
type Word = UInt16
type ULONG_PTR = QWord
type THandle = QWord
type HGDIOBJ = UIntptr
type HPALETTE = UIntptr
type LResult = UIntptr
type COLORREF = Cardinal
type TColorRef = COLORREF
type HMENU = UIntptr
type PtrInt = UIntptr
type PtrUInt = QWord
type HBITMAP = UIntptr
type HICON = UIntptr
type HMONITOR = UIntptr
type TCriticalSection = PtrUInt
type HOOK = QWord
type TFarProc = Pointer
type HBRUSH = UIntptr
type HPEN = UIntptr
type HKEY = UIntptr
type HCURSOR = HICON
type TGraphicsColor = Int32
type Smallint = Int16
type HFONT = UIntptr
type HRESULT = Int32
type TCefStringList = Pointer
type TGraphicsFillStyle = Int32
type TCefSharedTextureHandle = THandle

const (
	FsSurface TGraphicsFillStyle = iota // fill till the color (it fills all except this color)
	FsBorder                            // fill this color (it fills only connected pixels of this color)
)

type HRGN struct {
	instance unsafe.Pointer
}

type TagEnumLogFontA struct {
	ElfLogFont  *LogFontA
	ElfFullName []AnsiChar // len = 64
	ElfStyle    []AnsiChar // len = 32
}

type TagEnumLogFontAPtr struct {
	ElfLogFont  uintptr //*LogFontA
	LfFaceName  uintptr //string // len = 32
	ElfFullName uintptr //string // len = 64
	ElfStyle    uintptr //string // len = 32
}

type TagEnumLogFontExA struct {
	ElfLogFont  *LogFontA
	ElfFullName []AnsiChar // len = 64
	ElfStyle    []AnsiChar // len = 32
	ElfScript   []AnsiChar // len = 32
}

type TagEnumLogFontExAPtr struct {
	ElfLogFont  uintptr //*LogFontA
	LfFaceName  uintptr //[]AnsiChar // len = 32
	ElfFullName uintptr //[]AnsiChar // len = 64
	ElfStyle    uintptr //[]AnsiChar // len = 32
	ElfScript   uintptr //[]AnsiChar // len = 32
}

type Point struct {
	X int32
	Y int32
}

type Size struct {
	X int32
	Y int32
}

type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type TNewTextMetricEx struct {
	Ntmentm           TNewTextMetric
	NtmeFontSignature TFontSignature
}

type TNewTextMetricExPtr struct {
	Ntmentm           uintptr //TNewTextMetric
	NtmeFontSignature uintptr //TFontSignature
}

type TFontSignature struct {
	FsUsb []DWORD // len = 4
	FsCsb []DWORD // len = 2
}

type TFontSignaturePtr struct {
	FsUsb uintptr // []DWORD // len = 4
	FsCsb uintptr // []DWORD // len = 2
}

type TNewTextMetric struct {
	TmHeight           LongInt
	TmAscent           LongInt
	TmDescent          LongInt
	TmInternalLeading  LongInt
	TmExternalLeading  LongInt
	TmAveCharWidth     LongInt
	TmMaxCharWidth     LongInt
	TmWeight           LongInt
	TmOverhang         LongInt
	TmDigitizedAspectX LongInt
	TmDigitizedAspectY LongInt
	TmFirstChar        AnsiChar
	TmLastChar         AnsiChar
	TmDefaultChar      AnsiChar
	TmBreakChar        AnsiChar
	TmItalic           byte
	TmUnderlined       byte
	TmStruckOut        byte
	TmPitchAndFamily   byte
	TmCharSet          byte
	NtmFlags           DWORD
	NtmSizeEM          UINT
	NtmCellHeight      UINT
	NtmAvgWidth        UINT
}

type TagTextMetricA struct {
	TmHeight           LongInt
	TmAscent           LongInt
	TmDescent          LongInt
	TmInternalLeading  LongInt
	TmExternalLeading  LongInt
	TmAveCharWidth     LongInt
	TmMaxCharWidth     LongInt
	TmWeight           LongInt
	TmOverhang         LongInt
	TmDigitizedAspectX LongInt
	TmDigitizedAspectY LongInt
	TmFirstChar        AnsiChar
	TmLastChar         AnsiChar
	TmDefaultChar      AnsiChar
	TmBreakChar        AnsiChar
	TmItalic           byte
	TmUnderlined       byte
	TmStruckOut        byte
	TmPitchAndFamily   byte
	TmCharSet          byte
}

type TagLogPen struct {
	LopnStyle LongWord
	LopnWidth Point
	LopnColor TColorRef
}

type TagPaletteEntry struct {
	PeRed   byte
	PeGreen byte
	PeBlue  byte
	PeFlags byte
}

type TagLogPalette struct {
	palVersion    Word
	palNumEntries Word
	palPalEntry   []TagPaletteEntry
}

type TagScrollInfo struct {
	CbSize    UINT
	FMask     UINT
	NMin      Integer
	NMax      Integer
	NPage     UInt
	NPos      Integer
	NTrackPos Integer
}

type TagTriVertex struct {
	X     LongInt
	Y     LongInt
	Red   Word
	Green Word
	Blue  Word
	Alpha Word
}

type TagPaintStruct struct {
	Hdc        HDC
	FErase     BOOL
	RcPaint    Rect
	FRestore   BOOL
	FIncUpdate BOOL
}

type TagLogBrush struct {
	LbStyle LongWord
	LbColor TColorRef
	LbHatch PtrUInt
}

type ICONINFO struct {
	FIcon    BOOL
	XHotspot DWORD
	YHotspot DWORD
	HbmMask  HBITMAP
	HbmColor HBITMAP
}

// non-winapi radial gradient log info
type TLogGradientStop struct {
	RadColorR   Word
	RadColorG   Word
	RadColorB   Word
	RadColorA   Word
	RadPosition float64 // must be in 0..1
}

type TLogRadialGradient struct {
	RadCenterX Integer
	RadCenterY Integer
	RadRadius  Integer
	RadFocalX  Integer
	RadFocalY  Integer
	RadStops   []TLogGradientStop
}

type TagBitmapInfoHeader struct { // use packed, this is the .bmp file format
	BiSize          DWORD
	BiWidth         LongInt
	BiHeight        LongInt
	BiPlanes        Word
	BiBitCount      Word
	BiCompression   DWORD
	BiSizeImage     DWORD
	BiXPelsPerMeter LongInt
	BiYPelsPerMeter LongInt
	BiClrUsed       DWORD
	BiClrImportant  DWORD
}
type TagRGBQuad struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

type TagBitmapInfo struct {
	BmiHeader TagBitmapInfoHeader
	BmiColors []TagRGBQuad
}

type TagMonitorInfo struct {
	CbSize    DWORD
	RcMonitor Rect
	RcWork    Rect
	DwFlags   DWORD
}

type LogFontA struct {
	LfHeight         LongInt
	LfWidth          LongInt
	LfEscapement     LongInt // angle, in tenths of degrees of each line of text
	LfOrientation    LongInt // angle, in tenths of degrees of each character's base line
	LfWeight         LongInt
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       []AnsiChar // len = 32
}

func (m Int8) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int16) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int32) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int64) ToPtr() uintptr {
	return uintptr(m)
}

func (m Int) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt8) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt16) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt32) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt64) ToPtr() uintptr {
	return uintptr(m)
}

func (m UInt) ToPtr() uintptr {
	return uintptr(m)
}

func (m UIntptr) ToPtr() uintptr {
	return uintptr(m)
}

// SetValue
//
// 给指针设置值, 仅基础类型, 字符串需直接赋值
func (m UIntptr) SetValue(value interface{}) {
	if m == 0 {
		return
	}
	switch value.(type) {
	case uintptr:
		*(*uintptr)(unsafe.Pointer(m)) = value.(uintptr)
	case int:
		*(*int)(unsafe.Pointer(m)) = value.(int)
	case int8:
		*(*int8)(unsafe.Pointer(m)) = value.(int8)
	case int16:
		*(*int16)(unsafe.Pointer(m)) = value.(int16)
	case int32:
		*(*int32)(unsafe.Pointer(m)) = value.(int32)
	case int64:
		*(*int64)(unsafe.Pointer(m)) = value.(int64)
	case uint:
		*(*uint)(unsafe.Pointer(m)) = value.(uint)
	case uint8:
		*(*uint8)(unsafe.Pointer(m)) = value.(uint8)
	case uint16:
		*(*uint16)(unsafe.Pointer(m)) = value.(uint16)
	case uint32:
		*(*uint32)(unsafe.Pointer(m)) = value.(uint32)
	case uint64:
		*(*uint64)(unsafe.Pointer(m)) = value.(uint64)
	case float32:
		//*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m)))) = value.(float32)
		*(*float32)(unsafe.Pointer(m)) = value.(float32)
	case float64:
		//*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m)))) = value.(float64)
		*(*float64)(unsafe.Pointer(m)) = value.(float64)
	case bool:
		*(*bool)(unsafe.Pointer(m)) = value.(bool)
	case string:

	}
}

func (m String) ToPtr() uintptr {
	return api.PascalStr(string(m))
}

func (m Boolean) ToPtr() uintptr {
	return api.PascalBool(bool(m))
}

func (m Float32) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

func (m Float64) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

func NewHRGN(instance uintptr) *HRGN {
	return &HRGN{instance: unsafe.Pointer(instance)}
}

func (m *HRGN) Free() {
	m.instance = nil
}

func (m *HRGN) Instance() uintptr {
	return uintptr(m.instance)
}

// TSet 定义和 LCL TSet 一样，方便使用
type TSet = lcltypes.TSet

// NewSet
//
// 新建TSet，初始值为0，然后添加元素
func NewSet(opts ...uint8) TSet {
	return lcltypes.NewSet(opts...)
}

// BroderDirectionAdjustment 边框方向调整集合
type BroderDirectionAdjustment = uint8

const (
	BdaTop = iota
	BdaTopRight
	BdaRight
	BdaBottomRight
	BdaBottom
	BdaBottomLeft
	BdaLeft
	BdaTopLeft
)

// BroderDirectionAdjustments SET
type BroderDirectionAdjustments = TSet
