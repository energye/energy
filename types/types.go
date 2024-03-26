//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

import (
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
type TColor uint32 // TColor 常用值请见 types/colors 包

type PByte = uintptr
type HDWP = uintptr
type HMODULE = uintptr
type HINST = uintptr
type LPCWSTR = uintptr
type HRGN = uintptr
type LRESULT = uintptr
type HRSRC = uintptr
type HGLOBAL = uintptr
type TFNWndEnumProc = uintptr
type TXId = uint64
type ATOM = uint16
type TAtom = uint16
type SIZE_T = uintptr
type DWORD_PTR = uintptr
type TModalResult = int32
type THelpEventData = uintptr
type TTabOrder = int16
type PFNLVCOMPARE = uintptr
type PFNTVCOMPARE = uintptr
type Byte = uint8
type TFontCharset = uint8
type TSpacingSize = int32
type TClass = uintptr
type TThreadID = uintptr
type TClipboardFormat = uintptr
type Single = float32
type PChar = string
type Char = uint16 // Char Unicode 主要用于keymap, 参见types/keys包
type AnsiChar = Char

// type TCefColor = uint32
type Integer = int32
type LongInt = int32
type LongPtr = uintptr
type LongWord = uint32
type NativeUInt = uint32
type Cardinal = uint32
type HWND = uintptr
type WPARAM = uintptr
type LPARAM = uintptr
type HDC = uintptr
type Pointer = uintptr
type QWord = uintptr
type Word = uint16
type HGDIOBJ = uintptr
type HPALETTE = uintptr
type LResult = uintptr
type HMENU = uintptr
type PtrInt = uintptr
type HBITMAP = uintptr
type HICON = uintptr
type HMONITOR = uintptr
type HBRUSH = uintptr
type HPEN = uintptr
type HKEY = uintptr
type TGraphicsColor = int32
type SmallInt = int16
type HFONT = uintptr
type HRESULT = uintptr
type SizeInt = int
type DWORD = Cardinal
type ACCESS_MASK = DWORD
type REGSAM = ACCESS_MASK
type LongBool = Boolean
type BOOL = LongBool
type UINT = LongWord
type ULONG_PTR = QWord
type THandle = QWord
type HANDLE = THandle
type COLORREF = Cardinal
type TColorRef = COLORREF
type PtrUInt = QWord
type HCURSOR = HICON
type TFarProc = Pointer

// type TCefStringList = Pointer
type TCriticalSection = PtrUInt
type HOOK = QWord
type TCopyMode = LongInt
type TImageIndex = Integer
type TFontDataName = string
type HFont = HANDLE
type TFPResourceHandle = PtrUInt
type TFontCharSet = byte
type TResourceType = string
type HIMAGELIST = HANDLE
type TLCLHandle = PtrUInt

// Currency -922337203685477.5808到922337203685477.5807
// 10000倍
type Currency = int64

// TFPJPEGCompressionQuality = 1..100;
type TFPJPEGCompressionQuality = uint8
type TJPEGQualityRange = TFPJPEGCompressionQuality

// TDateTime Double
type TDateTime = Float64

// TDate TDateTime
type TDate = TDateTime

// TTime TDateTime
type TTime = TDateTime

type TDockZoneClass = uintptr
type TCollectionItemClass = uintptr
type TWinControlClass = uintptr
type TFPCustomImageClass = uintptr
type TGraphicClass = uintptr
type TStatusPanelClass = uintptr
type THeaderSectionClass = uintptr
type TListItemClass = uintptr
type TTreeNodeClass = uintptr

// TSet Pascal集合类型 set of xxx
type TSet uint32

// TUTF8Char
//
//	UTF-8 character is at most 6 bytes plus a #0
type TUTF8Char struct {
	Len     byte
	Content [7]byte
}

// TBrushPattern
//
//	array[0..PatternBitCount-1] of TPenPattern
//	PatternBitCount = sizeof(longword) * 8;
type TBrushPattern [32]uint32

// TOverlay = 0..14; // windows limitation
type TOverlay uint8

type TPoint struct {
	X int32
	Y int32
}

type TRect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type TSize struct {
	Cx int32
	Cy int32
}

func (u *TUTF8Char) ToString() string {
	if u.Len > 0 && u.Len < 7 {
		return string(u.Content[0:u.Len])
	}
	return ""
}

func (u *TUTF8Char) SetString(str string) {
	if str != "" {
		bs := []byte(str)
		u.Len = byte(len(bs))
		if u.Len > 6 {
			u.Len = 6
		}
		copy(u.Content[:], bs[:u.Len])
	}
}

//----------------------------------------------------------------------------------------------------------------------
// -- TRect

func Rect(left, top, right, bottom int32) TRect {
	return TRect{Left: left, Top: top, Right: right, Bottom: bottom}
}

func (r TRect) PtInRect(P TPoint) bool {
	return P.X >= r.Left && P.X < r.Right && P.Y >= r.Top && P.Y < r.Bottom
}

func (r TRect) Width() int32 {
	return r.Right - r.Left
}

func (r *TRect) SetWidth(val int32) {
	r.Right = r.Left + val
}

func (r TRect) Height() int32 {
	return r.Bottom - r.Top
}

func (r *TRect) SetHeight(val int32) {
	r.Bottom = r.Top + val
}

func (r TRect) IsEmpty() bool {
	return r.Right <= r.Left || r.Bottom <= r.Top
}

func (r *TRect) Empty() {
	r.Left = 0
	r.Top = 0
	r.Right = 0
	r.Bottom = 0
}

func (r TRect) Size() TSize {
	return TSize{r.Width(), r.Height()}
}

func (r *TRect) SetSize(w, h int32) {
	r.SetWidth(w)
	r.SetHeight(h)
}

func (r *TRect) Inflate(dx, dy int32) {
	r.Left += -dx
	r.Top += -dy
	r.Right += dx
	r.Bottom += dy
}

func (r TRect) Contains(aR TRect) bool {
	return r.Left <= aR.Left && r.Right >= aR.Right && r.Top <= aR.Top && r.Bottom >= aR.Bottom
}

func (r TRect) IntersectsWith(aR TRect) bool {
	return r.Left < aR.Right && r.Right > aR.Left && r.Top < aR.Bottom && r.Bottom > aR.Top
}

func (r TRect) CenterPoint() (ret TPoint) {
	ret.X = (r.Right-r.Left)/2 + r.Left
	ret.Y = (r.Bottom-r.Top)/2 + r.Top
	return
}

func (r *TRect) Scale(val float64) {
	r.Left = int32(float64(r.Left) * val)
	r.Top = int32(float64(r.Top) * val)
	r.Right = int32(float64(r.Right) * val)
	r.Bottom = int32(float64(r.Bottom) * val)
}

func (r *TRect) Scale2(val int) {
	r.Scale(float64(val))
}

// -- TPoint

func NewPoint(x, y int32) TPoint {
	return TPoint{X: x, Y: y}
}

func (p TPoint) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p *TPoint) Offset(dx, dy int32) {
	p.X += dx
	p.Y += dy
}

func (p *TPoint) Scale(val float64) {
	p.X = int32(float64(p.X) * val)
	p.Y = int32(float64(p.Y) * val)
}

func (p *TPoint) Scale2(val int) {
	p.Scale(float64(val))
}

// TMsg
//
// Only Windows, tagMSG
type TMsg struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      TPoint
}

// TCursorInfo
type TCursorInfo struct {
	CbSize      uint32
	Flags       uint32
	HCursor     HCURSOR
	PtScreenPos TPoint
}

// TWndClass
type TWndClass struct {
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  LPCWSTR
	LpszClassName LPCWSTR
}

// NewSet
//
// 新建TSet，初始值为0，然后添加元素
//
// Create a new TSet, the initial value is 0, and then add elements.
func NewSet(opts ...int32) TSet {
	var s TSet
	return s.Include(opts...)
}

// Include
//
// 集合加法，val...中存储为位的索引，下标为0
//
// Set addition, stored as bit index in val..., subscript 0.
func (s TSet) Include(val ...int32) TSet {
	r := uint32(s)
	for _, v := range val {
		r |= 1 << uint8(v)
	}
	return TSet(r)
}

// Exclude
//
// 集合减法，val...中存储为位的索引，下标为0
//
// Set subtraction, stored as bit index in val..., subscript 0.
func (s TSet) Exclude(val ...int32) TSet {
	r := uint32(s)
	for _, v := range val {
		r &= ^(1 << uint8(v))
	}
	return TSet(r)
}

// In
//
// 集合类型的判断，val表示位数，下标为0
//
// Judgment of the Set type, val represents the number of digits, and the subscript is 0.
func (s TSet) In(val int32) bool {
	if s&(1<<uint8(val)) != 0 {
		return true
	}
	return false
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
	LopnWidth TPoint
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
	RcPaint    TRect
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

type TFPColor struct {
	Red, Green, Blue, Alpha Word
}

type TagBitmapInfo struct {
	BmiHeader TagBitmapInfoHeader
	BmiColors []TagRGBQuad
}

type TagMonitorInfo struct {
	CbSize    DWORD
	RcMonitor TRect
	RcWork    TRect
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
	if m == "" {
		return 0
	}
	temp := []byte(m)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return uintptr(unsafe.Pointer(&utf8StrArr[0]))
}

func (m Boolean) ToPtr() uintptr {
	if m {
		return 1
	}
	return 0
}

func (m Float32) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

func (m Float64) ToPtr() uintptr {
	return uintptr(unsafe.Pointer(&m))
}

// TATFlatTheme Theme
type TATFlatTheme struct {
	FontName               string
	FontSize               Integer
	FontStyles             TFontStyles
	ColorFont              TColor
	ColorFontDisabled      TColor
	ColorFontListbox       TColor
	ColorFontListboxSel    TColor
	ColorFontOverlay       TColor
	ColorBgPassive         TColor
	ColorBgOver            TColor
	ColorBgChecked         TColor
	ColorBgDisabled        TColor
	ColorBgListbox         TColor
	ColorBgListboxSel      TColor
	ColorBgListboxHottrack TColor
	ColorBgOverlay         TColor
	ColorArrows            TColor
	ColorArrowsOver        TColor
	ColorSeparators        TColor
	ColorBorderPassive     TColor
	ColorBorderOver        TColor
	ColorBorderFocused     TColor
	EnableColorBgOver      Boolean
	MouseoverBorderWidth   Integer
	PressedBorderWidth     Integer
	PressedCaptionShiftY   Integer
	PressedCaptionShiftX   Integer
	BoldBorderWidth        Integer
	ChoiceBorderWidth      Integer
	ArrowSize              Integer
	GapForAutoSize         Integer
	TextOverlayPosition    TATButtonOverlayPosition
	SeparatorOffset        Integer
	XMarkWidth             Integer
	XMarkOffsetLeft        Integer
	XMarkOffsetRight       Integer
	XMarkLineWidth         Integer
	ScalePercents          Integer
	ScaleFontPercents      Integer
}

func (m *TATFlatTheme) DoScale(AValue Integer) Integer {
	return AValue * m.ScalePercents / 100
}

func (m *TATFlatTheme) DoScaleFont(AValue Integer) Integer {
	if m.ScaleFontPercents == 0 {
		return m.DoScale(AValue)
	} else {
		return AValue * m.ScaleFontPercents / 100
	}
}

// TScaledImageListResolution TODO record
type TScaledImageListResolution struct {
}

// TFontData record
type TFontData struct {
	Handle      HFont
	Height      Integer
	Pitch       TFontPitch
	Style       TFontStylesBase
	CharSet     TFontCharSet
	Quality     TFontQuality
	Name        TFontDataName
	Orientation Integer
}

// TPaperRect record
type TPaperRect struct {
	PhysicalRect TRect
	WorkRect     TRect
}

// TRawImage object
type TRawImage struct {
}

// TLCLTextMetric record
type TLCLTextMetric struct {
	Ascender  int32
	Descender int32
	Height    int32
}

// TThemedElementDetails record
type TThemedElementDetails struct {
	Element TThemedElement
	Part    int32
	State   int32
}

// TFontParams record
type TFontParams struct {
	Name       string
	Size       int32
	Color      TColor
	Style      TFontStyles
	HasBkClr   bool
	BkColor    TColor
	VScriptPos TVScriptPos
}

// TParaMetric record
type TParaMetric struct {
	FirstLine   float64 // in points
	TailIndent  float64 // in points
	HeadIndent  float64 // in points
	SpaceBefore float64 // in points
	SpaceAfter  float64 // in points
	LineSpacing float64 // multiplier - matching CSS line-height by percentage/em note, that normal LineSpacing is 1.2, not 1.0
}

// TParaNumbering record
type TParaNumbering struct {
	Style       TParaNumStyle
	Indent      float64
	CustomChar  string
	NumberStart int32 // used for pnNumber only
	SepChar     string
	ForceNewNum bool // if true and Style is pnNumber, NumberStart is used for the new numbering
}

// TParaRange record
type TParaRange struct {
	Start      int32 // the first character in the paragraph
	LengthNoBr int32 // the length of the paragraph, excluding the line break character
	Length     int32 // the length of the paragrpah, including the line break, if present the last line in the control doesn't contain a line break character, thus length = lengthNoBr
}

// TTabStop record
type TTabStop struct {
	Offset float64
	Align  TTabAlignment // not used
}

// TTabStopList record
type TTabStopList struct {
	Count int32
	Tabs  []TTabStop
}

// TRectOffsets record
type TRectOffsets struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

// TPrintParams record
type TPrintParams struct {
	JobTitle  string       // print job title to be shown in system printing manager
	Margins   TRectOffsets // margins in points
	SelStart  int32
	SelLength int32
}

// TRegDataInfo record
type TRegDataInfo struct {
	RegData  TRegDataType
	DataSize int32
}

// TRegKeyInfo record
type TRegKeyInfo struct {
	NumSubKeys   int32
	MaxSubKeyLen int32
	NumValues    int32
	MaxValueLen  int32
	MaxDataLen   int32
	FileTime     TDateTime
}

// TControlBorderSpacingDefault record
type TControlBorderSpacingDefault struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
	Around int32
}

// TWVWindowFeatures
// Record used by TCoreWebView2WindowFeatures.CopyToRecord to copy the windows featres
type TWVWindowFeatures struct {
	HasPosition             bool
	HasSize                 bool
	Left                    uint32
	Top                     uint32
	Width                   uint32
	Height                  uint32
	ShouldDisplayMenuBar    bool
	ShouldDisplayStatus     bool
	ShouldDisplayToolbar    bool
	ShouldDisplayScrollBars bool
}
