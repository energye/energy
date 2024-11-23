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

import "unsafe"

// TODO no impl

// ICefMediaRouter
//
//	Supports discovery of and communication with media devices on the local
//	network via the Cast and DIAL protocols. The functions of this interface may
//	be called on any browser process thread unless otherwise indicated.
//	<para><see cref="uCEFTypes|TCefMediaRouter">Implements TCefMediaRouter</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_router_t)</see></para>
type ICefMediaRouter struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefMediaRouter) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaRouter) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaRouter) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// ICefMediaSink
//
//	Callback interface for ICefMediaSink.GetDeviceInfo. The functions of
//	this interface will be called on the browser process UI thread.
//	<para><see cref="uCEFTypes|TCefMediaSinkDeviceInfoCallback">Implements TCefMediaSinkDeviceInfoCallback</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_device_info_callback_t)</see></para>
type ICefMediaSink struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefMediaSink) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaSink) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaSink) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// TCefMediaSinkArray
//
//	of []ICefMediaSink
type TCefMediaSinkArray struct {
	instance  unsafe.Pointer
	mediaSink []*ICefMediaSink
	count     uint32
}

// Instance 实例
func (m *TCefMediaSinkArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefMediaSinkArray) Free() {
	if m.instance != nil {
		m.instance = nil
	}
}

func (m *TCefMediaSinkArray) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// ICefMediaRoute
//
//	Represents the route between a media source and sink. Instances of this
//	object are created via ICefMediaRouter.CreateRoute and retrieved via
//	ICefMediaObserver.OnRoutes. Contains the status and metadata of a
//	routing operation. The functions of this interface may be called on any
//	browser process thread unless otherwise indicated.
//	<para><see cref="uCEFTypes|TCefMediaRoute">Implements TCefMediaRoute</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_t)</see></para>
type ICefMediaRoute struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefMediaRoute) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaRoute) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaRoute) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// TCefMediaRouteArray
//
//	of []ICefMediaRoute
type TCefMediaRouteArray struct {
	instance   unsafe.Pointer
	mediaRoute []*ICefMediaRoute
	count      uint32
}

// Instance 实例
func (m *TCefMediaRouteArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefMediaRouteArray) Free() {
	if m.instance != nil {
		m.instance = nil
	}
}

func (m *TCefMediaRouteArray) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}
