//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https//www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type TCEFBitmapBitBuffer struct {
	instance unsafe.Pointer
}

func NewBitmapBitBuffer(width, height int32) *TCEFBitmapBitBuffer {
	var result uintptr
	imports.SysCallN(def.BitmapBitBuffer_Free, uintptr(unsafe.Pointer(&result)))
	if result > 0 {
		return &TCEFBitmapBitBuffer{instance: getInstance(result)}
	}
	return nil
}

func (m *TCEFBitmapBitBuffer) Free() {
	if m.instance != nil {
		imports.SysCallN(def.BitmapBitBuffer_Free, m.Instance())
		m.instance = nil
	}
}

func (m *TCEFBitmapBitBuffer) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *TCEFBitmapBitBuffer) IsValid() bool {
	return m.instance != nil
}

func (m *TCEFBitmapBitBuffer) UpdateSize(width, height int32) {
	imports.SysCallN(def.BitmapBitBuffer_UpdateSize, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TCEFBitmapBitBuffer) Width() int32 {
	return int32(imports.SysCallN(def.BitmapBitBuffer_Width, m.Instance()))
}

func (m *TCEFBitmapBitBuffer) Height() int32 {
	return int32(imports.SysCallN(def.BitmapBitBuffer_Height, m.Instance()))
}

func (m *TCEFBitmapBitBuffer) BufferLength() int32 {
	return int32(imports.SysCallN(def.BitmapBitBuffer_BufferLength, m.Instance()))
}

func (m *TCEFBitmapBitBuffer) Empty() bool {
	return api.GoBool(imports.SysCallN(def.BitmapBitBuffer_Empty, m.Instance()))
}

// Scanline
//
//	return PByte = byte pointer
func (m *TCEFBitmapBitBuffer) Scanline(i int32) uintptr {
	return imports.SysCallN(def.BitmapBitBuffer_Scanline, m.Instance(), uintptr(i))
}

func (m *TCEFBitmapBitBuffer) ScanlineSize() int32 {
	return int32(imports.SysCallN(def.BitmapBitBuffer_ScanlineSize, m.Instance()))
}

func (m *TCEFBitmapBitBuffer) BufferScanlineSize() int32 {
	return int32(imports.SysCallN(def.BitmapBitBuffer_BufferScanlineSize, m.Instance()))
}

// BufferBits
//
//	Bits Pointer
func (m *TCEFBitmapBitBuffer) BufferBits() uintptr {
	return imports.SysCallN(def.BitmapBitBuffer_BufferScanlineSize, m.Instance())
}
