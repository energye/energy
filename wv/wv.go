//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// Package wv Webview2 实现模块
// 包含所有 Webview2 API 实现
package wv

import (
	"github.com/energye/energy/v2/lcl"
	"unsafe"
)

// IObject 根对象
type IObject = lcl.IObject
type TObject = lcl.TObject

// IComponent 根组件
type IComponent = lcl.IComponent
type TComponent = lcl.TComponent

// IWinControl 根Control
type IWinControl = lcl.IWinControl
type TWinControl = lcl.TWinControl

// IUnknown 根接口
//
//	UUID: '{00000000-0000-0000-C000-000000000046}'
type IUnknown = lcl.IUnknown
type Unknown = lcl.Unknown

type unsafePointer = unsafe.Pointer

// GetInstance As操作的简化。
//
// Simplification of As operation.
//
//go:noinline
func GetInstance(value interface{}) unsafePointer {
	return lcl.GetInstance(value)
}

// SetObjectInstance 设置对你指针实例值, 用于外部组件创建
func SetObjectInstance(object interface{}, instance unsafePointer) {
	if object == nil {
		return
	}
	switch object.(type) {
	case IObject:
		lcl.SetObjectInstance(object.(IObject), instance)
	case IUnknown:
		lcl.SetUnknownInstance(object.(IUnknown), instance)
	}
}

// GetObjectUintptr 获取对象指针地址值
func GetObjectUintptr(object interface{}) uintptr {
	if object == nil {
		return 0
	}
	switch object.(type) {
	case IObject:
		return lcl.GetObjectUintptr(object.(IObject))
	case IUnknown:
		return lcl.GetUnknownUintptr(object.(IUnknown))
	case IWVBrowserEvents:
		return object.(IWVBrowserEvents).Instance()
	}
	return 0
}

// TWVCustomSchemeInfo
// Record with all the information to create a TCoreWebView2CustomSchemeRegistration instance to register a custom scheme.
type TWVCustomSchemeInfo struct {
	SchemeName            string // The name of the custom scheme to register.
	TreatAsSecure         bool   // Whether the sites with this scheme will be treated as a Secure Context like an HTTPS site.
	AllowedDomains        string // Comma separated list of origins that are allowed to issue requests with the custom scheme, such as XHRs and subresource requests that have an Origin header.
	HasAuthorityComponent bool   // Set this property to true if the URIs with this custom scheme will have an authority component (a host for custom schemes).
}

// tWVCustomSchemeInfoPtr = ^TWVCustomSchemeInfo
type tWVCustomSchemeInfoPtr struct {
	SchemeName            uintptr //string
	TreatAsSecure         uintptr //bool
	AllowedDomains        uintptr //string
	HasAuthorityComponent uintptr //bool
}

// OleVariant OLE TODO no impl
type OleVariant struct {
	instance unsafe.Pointer
}

// COREWEBVIEW2_COLOR
//
//	A value representing RGBA color (Red, Green, Blue, Alpha) for WebView2.
//	Each component takes a value from 0 to 255, with 0 being no intensity and 255 being the highest intensity.
//	<para><see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_color">See the Globals article.</see></para>
type COREWEBVIEW2_COLOR struct {
	// Specifies the intensity of the Alpha ie. opacity value. 0 is transparent, 255 is opaque.
	A byte
	// Specifies the intensity of the Red color.
	R byte
	// Specifies the intensity of the Green color.
	G byte
	// Specifies the intensity of the Blue color.
	B byte
}
