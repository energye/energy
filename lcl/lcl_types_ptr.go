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

// TODO no impl
type tVTImageInfo struct {
}

// tVTPaintInfo Record
type tVTPaintInfo struct {
	Canvas           uintptr // ICanvas                 // the canvas to paint on
	PaintOptions     uintptr // TVTInternalPaintOptions // a copy of the paint options passed to PaintTree
	Node             uintptr // PVirtualNode            // the node to paint
	Column           uintptr // TColumnIndex            // the node's column index to paint
	Position         uintptr // TColumnPosition         // the column position of the node
	CellRect         uintptr // TRect                   // the node cell
	ContentRect      uintptr // TRect                   // the area of the cell used for the node's content
	NodeWidth        uintptr // Integer                 // the actual node width
	Alignment        uintptr // TAlignment              // how to align within the node rectangle
	CaptionAlignment uintptr // TAlignment              // how to align text within the caption rectangle
	BidiMode         uintptr // TBiDiMode               // directionality to be used for painting
	BrushOrigin      uintptr // TPoint                  // the alignment for the brush used to draw dotted lines
	imageInfoPtr     uintptr // unsafePointer           // [4]TVTImageInfo // info about each possible node image, array[TVTImageInfoIndex] of TVTImageInfo
}

func (m *tVTPaintInfo) Convert() *TVTPaintInfo {
	if m == nil {
		return nil
	}
	return &TVTPaintInfo{
		Canvas:           AsCanvas(*(*uintptr)(unsafePointer(m.Canvas))),
		PaintOptions:     *(*TVTInternalPaintOptions)(unsafePointer(m.PaintOptions)),
		Node:             AsVirtualNode(m.Node),
		Column:           *(*TColumnIndex)(unsafePointer(m.Column)),
		Position:         *(*TColumnPosition)(unsafePointer(m.Position)),
		CellRect:         *(*TRect)(unsafePointer(m.CellRect)),
		ContentRect:      *(*TRect)(unsafePointer(m.ContentRect)),
		NodeWidth:        *(*Integer)(unsafePointer(m.NodeWidth)),
		Alignment:        *(*TAlignment)(unsafePointer(m.Alignment)),
		CaptionAlignment: *(*TAlignment)(unsafePointer(m.CaptionAlignment)),
		BidiMode:         *(*TBiDiMode)(unsafePointer(m.BidiMode)),
		BrushOrigin:      *(*TPoint)(unsafePointer(m.BrushOrigin)),
		imageInfoPtr:     unsafePointer(m.imageInfoPtr),
	}
}

// tHeaderPaintInfo Record
type tHeaderPaintInfo struct {
	TargetCanvas    uintptr // ICanvas
	Column          uintptr // TVirtualTreeColumn
	PaintRectangle  uintptr // TRect
	TextRectangle   uintptr // TRect
	IsHoverIndex    uintptr // Boolean
	IsDownIndex     uintptr // Boolean
	IsEnabled       uintptr // Boolean
	ShowHeaderGlyph uintptr // Boolean
	ShowSortGlyph   uintptr // Boolean
	ShowRightBorder uintptr // Boolean
	DropMark        uintptr // TVTDropMarkMode
	GlyphPos        uintptr // TPoint
	SortGlyphPos    uintptr // TPoint
}

func (m *tHeaderPaintInfo) Convert() *THeaderPaintInfo {
	if m == nil {
		return nil
	}
	return &THeaderPaintInfo{
		instance:        m,
		TargetCanvas:    AsCanvas(*(*uintptr)(unsafePointer(m.TargetCanvas))),
		Column:          AsVirtualTreeColumn(*(*uintptr)(unsafePointer(m.Column))),
		PaintRectangle:  (*TRect)(unsafePointer(m.PaintRectangle)),
		TextRectangle:   (*TRect)(unsafePointer(m.TextRectangle)),
		IsHoverIndex:    *(*Boolean)(unsafePointer(m.IsHoverIndex)),
		IsDownIndex:     *(*Boolean)(unsafePointer(m.IsDownIndex)),
		IsEnabled:       *(*Boolean)(unsafePointer(m.IsEnabled)),
		ShowHeaderGlyph: *(*Boolean)(unsafePointer(m.ShowHeaderGlyph)),
		ShowSortGlyph:   *(*Boolean)(unsafePointer(m.ShowSortGlyph)),
		ShowRightBorder: *(*Boolean)(unsafePointer(m.ShowRightBorder)),
		DropMark:        *(*TVTDropMarkMode)(unsafePointer(m.DropMark)),
		GlyphPos:        (*TPoint)(unsafePointer(m.GlyphPos)),
		SortGlyphPos:    (*TPoint)(unsafePointer(m.SortGlyphPos)),
	}
}

// tHitInfo Record
type tHitInfo struct {
	HitNode      uintptr // PVirtualNode
	HitPositions uintptr // THitPositions
	HitColumn    uintptr // TColumnIndex
	HitPoint     uintptr // TPoint
}

func (m *tHitInfo) Convert() *THitInfo {
	if m == nil {
		return nil
	}
	return &THitInfo{
		HitNode:      AsVirtualNode(m.HitNode),
		HitPositions: *(*THitPositions)(unsafePointer(m.HitPositions)),
		HitColumn:    *(*TColumnIndex)(unsafePointer(m.HitColumn)),
		HitPoint:     *(*TPoint)(unsafePointer(m.HitPoint)),
	}
}

// tHintInfo = ^THintInfo,
type tHintInfo struct {
	HintControl     uintptr // TControl
	HintWindowClass uintptr // TWinControlClass
	HintPos         uintptr // TPoint // screen coordinates
	HintMaxWidth    uintptr // int32
	HintColor       uintptr // TColor
	CursorRect      uintptr // TRect
	CursorPos       uintptr // TPoint
	ReshowTimeout   uintptr // int32
	HideTimeout     uintptr // int32
	HintStr         uintptr // string
	HintData        uintptr // Pointer
}

func (m *tHintInfo) Convert() *THintInfo {
	if m == nil {
		return nil
	}
	return &THintInfo{
		instance:        m,
		HintControl:     AsControl(*(*uintptr)(unsafePointer(m.HintControl))),
		HintWindowClass: *(*TWinControlClass)(unsafePointer(m.HintWindowClass)),
		HintPos:         (*TPoint)(unsafePointer(m.HintPos)),
		HintMaxWidth:    *(*int32)(unsafePointer(m.HintMaxWidth)),
		HintColor:       *(*TColor)(unsafePointer(m.HintColor)),
		CursorRect:      (*TRect)(unsafePointer(m.CursorRect)),
		CursorPos:       (*TPoint)(unsafePointer(m.CursorPos)),
		ReshowTimeout:   *(*int32)(unsafePointer(m.ReshowTimeout)),
		HideTimeout:     *(*int32)(unsafePointer(m.HideTimeout)),
		HintStr:         api.GoStr(m.HintStr),
		HintData:        m.HintData,
	}
}
