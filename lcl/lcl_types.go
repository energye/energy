//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// TRawImagePosition
//
//	Record describing a position in Raw Image data.
//	Byte is the byte offset, Bit the bit number/offset in that byte.
type TRawImagePosition struct {
	Byte PtrUInt
	Bit  Cardinal
}

// TColumnIndex int32
type TColumnIndex = int32

// TNodeArray array of PVirtualNode
// TODO no impl, 当前写法不正确
type TNodeArray = uintptr

// TColumnPosition Cardinal
type TColumnPosition = Cardinal

// TVTPaintInfo Record
type TVTPaintInfo struct {
	Canvas           ICanvas                 // the canvas to paint on
	PaintOptions     TVTInternalPaintOptions // a copy of the paint options passed to PaintTree
	Node             IVirtualNode            // the node to paint
	Column           TColumnIndex            // the node's column index to paint
	Position         TColumnPosition         // the column position of the node
	CellRect         TRect                   // the node cell
	ContentRect      TRect                   // the area of the cell used for the node's content
	NodeWidth        Integer                 // the actual node width
	Alignment        TAlignment              // how to align within the node rectangle
	CaptionAlignment TAlignment              // how to align text within the caption rectangle
	BidiMode         TBiDiMode               // directionality to be used for painting
	BrushOrigin      TPoint                  // the alignment for the brush used to draw dotted lines
	imageInfoPtr     unsafePointer           // [4]TVTImageInfo // info about each possible node image, array[TVTImageInfoIndex] of TVTImageInfo
}

// TVTImageInfo Record
//
//	For painting a node and its columns/cells a lot of information must be passed frequently around.
type TVTImageInfo struct {
	Index   Integer          // Index in the associated image list.
	XPos    Integer          // Horizontal position in the current target canvas.
	YPos    Integer          // Vertical position in the current target canvas.
	Ghosted Boolean          // Flag to indicate that the image must be drawn slightly lighter.
	Images  TCustomImageList // The image list to be used for painting.
}

// ImageInfo array[TVTImageInfoIndex] of TVTImageInfo
func (m *TVTPaintInfo) ImageInfo(index TVTImageInfoIndex) *TVTImageInfo {
	if index >= 0 && index <= 3 {
		var result uintptr
		api.LCLPreDef().SysCallN(api.VTImageInfoGet(), uintptr(m.imageInfoPtr), uintptr(index), uintptr(unsafePointer(&result)))
	}
	return nil
}

// TVTHeaderHitInfo Record
//
//	Structure used when info about a certain position in the header is needed.
type TVTHeaderHitInfo struct {
	X           Integer
	Y           Integer
	Button      TMouseButton
	Shift       TShiftState
	Column      TColumnIndex
	HitPosition TVTHeaderHitPositions
}

// THeaderPaintInfo Record
//
//	This structure carries all important information about header painting and is used in the advanced header painting.
type THeaderPaintInfo struct {
	instance        *tHeaderPaintInfo
	TargetCanvas    ICanvas
	Column          IVirtualTreeColumn
	PaintRectangle  *TRect
	TextRectangle   *TRect
	IsHoverIndex    Boolean
	IsDownIndex     Boolean
	IsEnabled       Boolean
	ShowHeaderGlyph Boolean
	ShowSortGlyph   Boolean
	ShowRightBorder Boolean
	DropMark        TVTDropMarkMode
	GlyphPos        *TPoint
	SortGlyphPos    *TPoint
}

func (m *THeaderPaintInfo) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	var setRectPtrVal = func(src, target *TRect) {
		target.Left = src.Left
		target.Top = src.Top
		target.Right = src.Right
		target.Bottom = src.Bottom
	}
	var setPointPtrVal = func(src, target *TPoint) {
		target.X = src.X
		target.Y = src.Y
	}
	*(*uintptr)(unsafePointer(m.instance.TargetCanvas)) = m.TargetCanvas.Instance()
	*(*uintptr)(unsafePointer(m.instance.Column)) = m.Column.Instance()
	paintRectangle := (*TRect)(unsafePointer(m.instance.PaintRectangle))
	textRectangleL := (*TRect)(unsafePointer(m.instance.TextRectangle))
	setRectPtrVal(m.PaintRectangle, paintRectangle)
	setRectPtrVal(m.TextRectangle, textRectangleL)
	*(*Boolean)(unsafePointer(m.instance.IsHoverIndex)) = m.IsHoverIndex
	*(*Boolean)(unsafePointer(m.instance.IsDownIndex)) = m.IsDownIndex
	*(*Boolean)(unsafePointer(m.instance.IsEnabled)) = m.IsEnabled
	*(*Boolean)(unsafePointer(m.instance.ShowHeaderGlyph)) = m.ShowHeaderGlyph
	*(*Boolean)(unsafePointer(m.instance.ShowSortGlyph)) = m.ShowSortGlyph
	*(*Boolean)(unsafePointer(m.instance.ShowRightBorder)) = m.ShowRightBorder
	*(*TVTDropMarkMode)(unsafePointer(m.instance.DropMark)) = m.DropMark
	glyphPos := (*TPoint)(unsafePointer(m.instance.GlyphPos))
	sortGlyphPos := (*TPoint)(unsafePointer(m.instance.SortGlyphPos))
	setPointPtrVal(m.GlyphPos, glyphPos)
	setPointPtrVal(m.SortGlyphPos, sortGlyphPos)
}

// THitInfo Record
//
//	Structure used when info about a certain position in the tree is needed.
type THitInfo struct {
	HitNode      IVirtualNode
	HitPositions THitPositions
	HitColumn    TColumnIndex
	HitPoint     TPoint
}

// THintInfo record
type THintInfo struct {
	instance        *tHintInfo
	HintControl     IControl
	HintWindowClass TWinControlClass
	HintPos         *TPoint // screen coordinates
	HintMaxWidth    int32
	HintColor       TColor
	CursorRect      *TRect
	CursorPos       *TPoint
	ReshowTimeout   int32
	HideTimeout     int32
	HintStr         string
	HintData        Pointer
}

func (m *THintInfo) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	var setRectPtrVal = func(src, target *TRect) {
		target.Left = src.Left
		target.Top = src.Top
		target.Right = src.Right
		target.Bottom = src.Bottom
	}
	var setPointPtrVal = func(src, target *TPoint) {
		target.X = src.X
		target.Y = src.Y
	}
	*(*uintptr)(unsafePointer(m.instance.HintControl)) = m.HintControl.Instance()
	*(*TWinControlClass)(unsafePointer(m.instance.HintWindowClass)) = m.HintWindowClass
	hintPos := (*TPoint)(unsafePointer(m.instance.HintPos))
	setPointPtrVal(m.HintPos, hintPos)
	*(*int32)(unsafePointer(m.instance.HintMaxWidth)) = m.HintMaxWidth
	*(*TColor)(unsafePointer(m.instance.HintColor)) = m.HintColor
	cursorRect := (*TRect)(unsafePointer(m.instance.CursorRect))
	setRectPtrVal(m.CursorRect, cursorRect)
	cursorPos := (*TPoint)(unsafePointer(m.instance.CursorPos))
	setPointPtrVal(m.CursorPos, cursorPos)
	*(*int32)(unsafePointer(m.instance.ReshowTimeout)) = m.ReshowTimeout
	*(*int32)(unsafePointer(m.instance.HideTimeout)) = m.HideTimeout
	m.instance.HintStr = api.PascalStr(m.HintStr)
	m.instance.HintData = m.HintData
}
