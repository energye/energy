//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefImage
type ICefImage struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// ImageRef -> ICefImage
var ImageRef image

// image
type image uintptr

func (m *image) New() *ICefImage {
	var result uintptr
	imports.Proc(def.CEFImage_New).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefImage) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefImage) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefImage) AddPngFS(scaleFactor float32, filename string) bool {
	if !m.IsValid() {
		return false
	}
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return false
	}
	return m.AddPng(scaleFactor, bytes)
}

func (m *ICefImage) AddPng(scaleFactor float32, png []byte) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFImage_AddPng).Call(m.Instance(), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&png[0])), uintptr(uint32(len(png))))
	return api.GoBool(r1)
}

func (m *ICefImage) AddJpegFS(scaleFactor float32, filename string) bool {
	if !m.IsValid() {
		return false
	}
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return false
	}
	return m.AddJpeg(scaleFactor, bytes)
}

func (m *ICefImage) AddJpeg(scaleFactor float32, jpeg []byte) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFImage_AddJpeg).Call(m.Instance(), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&jpeg[0])), uintptr(uint32(len(jpeg))))
	return api.GoBool(r1)
}

func (m *ICefImage) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFImage_GetWidth).Call(m.Instance())
	return int32(r1)
}

func (m *ICefImage) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFImage_GetHeight).Call(m.Instance())
	return int32(r1)
}

func (m *ICefImage) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
