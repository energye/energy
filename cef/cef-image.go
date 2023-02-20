//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Image
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefImage 实例
type ICefImage struct {
	instance unsafe.Pointer
}

//func NewImage() *ICefImage {
//	r1, _, _ := imports.Proc(internale_CEFImage_New).Call()
//	return &ICefImage{
//		instance: unsafe.Pointer(r1),
//	}
//}

func (m *ICefImage) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefImage) AddPngFS(scaleFactor float32, filename string) bool {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CEFImage_AddPng).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(uint32(len(bytes))))
	return api.GoBool(r1)
}

func (m *ICefImage) AddPng(scaleFactor float32, png []byte) bool {
	r1, _, _ := imports.Proc(internale_CEFImage_AddPng).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&png[0])), uintptr(int32(len(png))))
	return api.GoBool(r1)
}

func (m *ICefImage) AddJpeg(scaleFactor float32, jpeg []byte) bool {
	r1, _, _ := imports.Proc(internale_CEFImage_AddJpeg).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&jpeg[0])), uintptr(int32(len(jpeg))))
	return api.GoBool(r1)
}

func (m *ICefImage) GetWidth() int32 {
	r1, _, _ := imports.Proc(internale_CEFImage_GetWidth).Call(uintptr(m.instance))
	return int32(r1)
}

func (m *ICefImage) GetHeight() int32 {
	r1, _, _ := imports.Proc(internale_CEFImage_GetHeight).Call(uintptr(m.instance))
	return int32(r1)
}
