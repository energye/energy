//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICEFBitmapBitBuffer Parent: IObject
//
//	Class that stores a copy of the raw bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
type ICEFBitmapBitBuffer interface {
	IObject
	// Width
	//  Image width.
	Width() int32 // property
	// Height
	//  Image height.
	Height() int32 // property
	// BufferLength
	//  Buffer length.
	BufferLength() int32 // property
	// Empty
	//  Returns true if the buffer is empty.
	Empty() bool // property
	// Scanline
	//  Returns a pointer to the first byte in of the Y scnaline.
	Scanline(y int32) PByte // property
	// ScanlineSize
	//  Returns the scanline size.
	ScanlineSize() int32 // property
	// BufferScanlineSize
	//  Returns the real buffer scanline size.
	BufferScanlineSize() int32 // property
	// BufferBits
	//  Returns a pointer to the buffer that stores the image.
	BufferBits() uintptr // property
	// UpdateSize
	//  Updates the image size.
	UpdateSize(aWidth, aHeight int32) // procedure
}

// TCEFBitmapBitBuffer Parent: TObject
//
//	Class that stores a copy of the raw bitmap buffer sent by CEF in the TChromiumCore.OnPaint event.
type TCEFBitmapBitBuffer struct {
	TObject
}

func NewCEFBitmapBitBuffer(aWidth, aHeight int32) ICEFBitmapBitBuffer {
	r1 := CEF().SysCallN(79, uintptr(aWidth), uintptr(aHeight))
	return AsCEFBitmapBitBuffer(r1)
}

func (m *TCEFBitmapBitBuffer) Width() int32 {
	r1 := CEF().SysCallN(85, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) Height() int32 {
	r1 := CEF().SysCallN(81, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) BufferLength() int32 {
	r1 := CEF().SysCallN(76, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) Empty() bool {
	r1 := CEF().SysCallN(80, m.Instance())
	return GoBool(r1)
}

func (m *TCEFBitmapBitBuffer) Scanline(y int32) PByte {
	r1 := CEF().SysCallN(82, m.Instance(), uintptr(y))
	return PByte(r1)
}

func (m *TCEFBitmapBitBuffer) ScanlineSize() int32 {
	r1 := CEF().SysCallN(83, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) BufferScanlineSize() int32 {
	r1 := CEF().SysCallN(77, m.Instance())
	return int32(r1)
}

func (m *TCEFBitmapBitBuffer) BufferBits() uintptr {
	r1 := CEF().SysCallN(75, m.Instance())
	return uintptr(r1)
}

func CEFBitmapBitBufferClass() TClass {
	ret := CEF().SysCallN(78)
	return TClass(ret)
}

func (m *TCEFBitmapBitBuffer) UpdateSize(aWidth, aHeight int32) {
	CEF().SysCallN(84, m.Instance(), uintptr(aWidth), uintptr(aHeight))
}
