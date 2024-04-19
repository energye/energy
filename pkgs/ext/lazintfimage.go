//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ext

import (
	"github.com/energye/energy/v2/common/imports"
	"unsafe"
)

type DataDescription = int32

const (
	Init_BPP32_B8G8R8_BIO_TTB DataDescription = iota
	Init_BPP32_B8G8R8_M1_BIO_TTB
	Init_BPP32_B8G8R8A8_BIO_TTB
	Init_BPP32_B8G8R8A8_M1_BIO_TTB
)

type TLazIntfImage struct {
	instance unsafe.Pointer
}

type TColor struct {
	Red, Green, Blue, Alpha uint16
}

func NewLazIntfImage(width, height int32) *TLazIntfImage {
	m := &TLazIntfImage{}
	r1, _, _ := imports.LibLCLExt().Proc(LazIntfImage_Create).Call(uintptr(width), uintptr(height))
	m.instance = unsafe.Pointer(r1)
	return m
}

// Instance
//
// 返回对象实例指针。
func (m *TLazIntfImage) Instance() uintptr {
	return uintptr(m.instance)
}

// IsValid
//
// 检测地址是否为空。
func (m *TLazIntfImage) IsValid() bool {
	return m.instance != nil
}

func (m *TLazIntfImage) DataDescription(dataDescription DataDescription, width, height int32) {
	if !m.IsValid() {
		return
	}
	imports.LibLCLExt().Proc(LazIntfImage_DataDescription).Call(m.Instance(), uintptr(dataDescription), uintptr(width), uintptr(height))
}

func (m *TLazIntfImage) Colors(x, y int32, color TColor) {
	if !m.IsValid() {
		return
	}
	imports.LibLCLExt().Proc(LazIntfImage_Colors).Call(m.Instance(), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(&color)))
}

func (m *TLazIntfImage) Free() {
	if !m.IsValid() {
		return
	}
	imports.LibLCLExt().Proc(LazIntfImage_Free).Call(m.Instance())
	m.instance = nil
}
