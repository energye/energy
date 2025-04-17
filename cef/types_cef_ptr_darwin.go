//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package cef

import (
	. "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCefWindowInfo
type tCefWindowInfoPtr struct {
	WindowName                 UIntptr //TCefString
	X                          UIntptr //Integer
	Y                          UIntptr //Integer
	Width                      UIntptr //Integer
	Height                     UIntptr //Integer
	Hidden                     UIntptr //Integer
	ParentView                 UIntptr //TCefWindowHandle
	WindowlessRenderingEnabled UIntptr //Integer
	SharedTextureEnabled       UIntptr //Integer
	ExternalBeginFrameEnabled  UIntptr //Integer
	View                       UIntptr //TCefWindowHandle
}

// Convert 转换为结构
func (m *tCefWindowInfoPtr) convert() *TCefWindowInfo {
	getPtr := func(ptr uintptr) unsafe.Pointer {
		return unsafe.Pointer(ptr)
	}
	getInteger := func(ptr uintptr) Integer {
		if ptr == 0 {
			return 0
		}
		return *(*Integer)(getPtr(ptr))
	}
	return &TCefWindowInfo{
		instance:                   m,
		WindowName:                 TCefString(api.GoStr(m.WindowName.ToPtr())),        // TCefString
		X:                          getInteger(m.X.ToPtr()),                            // Integer
		Y:                          getInteger(m.Y.ToPtr()),                            // Integer
		Width:                      getInteger(m.Width.ToPtr()),                        // Integer
		Height:                     getInteger(m.Height.ToPtr()),                       // Integer
		Hidden:                     getInteger(m.Hidden.ToPtr()),                       // Integer
		ParentView:                 *(*TCefWindowHandle)(getPtr(m.ParentView.ToPtr())), // TCefWindowHandle
		WindowlessRenderingEnabled: getInteger(m.WindowlessRenderingEnabled.ToPtr()),   // Integer
		SharedTextureEnabled:       getInteger(m.SharedTextureEnabled.ToPtr()),         // Integer
		ExternalBeginFrameEnabled:  getInteger(m.ExternalBeginFrameEnabled.ToPtr()),    // Integer
		View:                       *(*TCefWindowHandle)(getPtr(m.View.ToPtr())),       // TCefWindowHandle
	}
}
