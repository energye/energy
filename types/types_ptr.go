//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// PATFlatTheme = ^TATFlatTheme
type PATFlatTheme struct {
	FontName               uintptr // string
	FontSize               uintptr // Integer
	FontStyles             uintptr // TFontStyles
	ColorFont              uintptr // TColor
	ColorFontDisabled      uintptr // TColor
	ColorFontListbox       uintptr // TColor
	ColorFontListboxSel    uintptr // TColor
	ColorFontOverlay       uintptr // TColor
	ColorBgPassive         uintptr // TColor
	ColorBgOver            uintptr // TColor
	ColorBgChecked         uintptr // TColor
	ColorBgDisabled        uintptr // TColor
	ColorBgListbox         uintptr // TColor
	ColorBgListboxSel      uintptr // TColor
	ColorBgListboxHottrack uintptr // TColor
	ColorBgOverlay         uintptr // TColor
	ColorArrows            uintptr // TColor
	ColorArrowsOver        uintptr // TColor
	ColorSeparators        uintptr // TColor
	ColorBorderPassive     uintptr // TColor
	ColorBorderOver        uintptr // TColor
	ColorBorderFocused     uintptr // TColor
	EnableColorBgOver      uintptr // Boolean
	MouseoverBorderWidth   uintptr // Integer
	PressedBorderWidth     uintptr // Integer
	PressedCaptionShiftY   uintptr // Integer
	PressedCaptionShiftX   uintptr // Integer
	BoldBorderWidth        uintptr // Integer
	ChoiceBorderWidth      uintptr // Integer
	ArrowSize              uintptr // Integer
	GapForAutoSize         uintptr // Integer
	TextOverlayPosition    uintptr // TATButtonOverlayPosition
	SeparatorOffset        uintptr // Integer
	XMarkWidth             uintptr // Integer
	XMarkOffsetLeft        uintptr // Integer
	XMarkOffsetRight       uintptr // Integer
	XMarkLineWidth         uintptr // Integer
	ScalePercents          uintptr // Integer
	ScaleFontPercents      uintptr // Integer
}

// PThemedElementDetails = ^TThemedElementDetails
type PThemedElementDetails struct {
	Element uintptr // TThemedElement
	Part    uintptr // int32
	State   uintptr // int32
}

// PPaperRect = ^TPaperRect
type PPaperRect struct {
	PhysicalRect uintptr //TRect
	WorkRect     uintptr //TRect
}

// PFontData = ^TFontData
type PFontData struct {
	Handle      uintptr //HFont
	Height      uintptr //Integer
	Pitch       uintptr //TFontPitch
	Style       uintptr //TFontStylesBase
	CharSet     uintptr //TFontCharSet
	Quality     uintptr //TFontQuality
	Name        uintptr //TFontDataName
	Orientation uintptr //Integer
}

// PControlBorderSpacingDefault = ^TControlBorderSpacingDefault
type PControlBorderSpacingDefault struct {
	Left   uintptr // int32
	Top    uintptr // int32
	Right  uintptr // int32
	Bottom uintptr // int32
	Around uintptr // int32
}

// PRegKeyInfo = ^TRegKeyInfo
type PRegKeyInfo struct {
	NumSubKeys   uintptr // int32
	MaxSubKeyLen uintptr // int32
	NumValues    uintptr // int32
	MaxValueLen  uintptr // int32
	MaxDataLen   uintptr // int32
	FileTime     uintptr // TDateTime float64
}

// PRegDataInfo = ^TRegDataInfo
type PRegDataInfo struct {
	RegData  uintptr // TRegDataType
	DataSize uintptr // int32
}

// PPrintParams = ^TPrintParams
type PPrintParams struct {
	JobTitle  uintptr // string       // print job title to be shown in system printing manager
	Margins   uintptr // TRectOffsets // margins in points
	SelStart  uintptr // int32
	SelLength uintptr // int32
}

// PRectOffsets = ^TRectOffsets
type PRectOffsets struct {
	Left   uintptr // float64
	Top    uintptr // float64
	Right  uintptr // float64
	Bottom uintptr // float64
}

// PTabStopList = ^TTabStopList
type PTabStopList struct {
	Count uintptr // int32
	Tabs  uintptr // []TTabStop
}

// PTabStop = ^TTabStop
type PTabStop struct {
	Offset uintptr // float64
	Align  uintptr // TTabAlignment // not used
}

// PParaRange = ^TParaRange
type PParaRange struct {
	Start      uintptr // int32 // the first character in the paragraph
	LengthNoBr uintptr // int32 // the length of the paragraph, excluding the line break character
	Length     uintptr // int32 // the length of the paragrpah, including the line break, if present the last line in the control doesn't contain a line break character, thus length = lengthNoBr
}

// PParaNumbering = ^TParaNumbering
type PParaNumbering struct {
	Style       uintptr // TParaNumStyle
	Indent      uintptr // float64
	CustomChar  uintptr // string
	NumberStart uintptr // int32 // used for pnNumber only
	SepChar     uintptr // string
	ForceNewNum uintptr // bool // if true and Style is pnNumber, NumberStart is used for the new numbering
}

// PParaMetric = ^TParaMetric
type PParaMetric struct {
	FirstLine   uintptr // float64 // in points
	TailIndent  uintptr // float64 // in points
	HeadIndent  uintptr // float64 // in points
	SpaceBefore uintptr // float64 // in points
	SpaceAfter  uintptr // float64 // in points
	LineSpacing uintptr // float64 // multiplier - matching CSS line-height by percentage/em note, that normal LineSpacing is 1.2, not 1.0
}

// PFontParams = ^TFontParams
type PFontParams struct {
	Name       uintptr // string
	Size       uintptr // int32
	Color      uintptr // TColor
	Style      uintptr // TFontStyles
	HasBkClr   uintptr // bool
	BkColor    uintptr // TColor
	VScriptPos uintptr // TVScriptPos
}
