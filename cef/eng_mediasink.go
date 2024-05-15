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

// ICefMediaSink Parent: ICefBaseRefCounted
//
//	Represents a sink to which media can be routed. Instances of this object are retrieved via ICefMediaObserver.OnSinks. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_t))</a>
type ICefMediaSink interface {
	ICefBaseRefCounted
	// GetId
	//  Returns the ID for this sink.
	GetId() string // function
	// GetName
	//  Returns the name of this sink.
	GetName() string // function
	// GetIconType
	//  Returns the icon type for this sink.
	GetIconType() TCefMediaSinkIconType // function
	// IsCastSink
	//  Returns true (1) if this sink accepts content via Cast.
	IsCastSink() bool // function
	// IsDialSink
	//  Returns true (1) if this sink accepts content via DIAL.
	IsDialSink() bool // function
	// IsCompatibleWith
	//  Returns true (1) if this sink is compatible with |source|.
	IsCompatibleWith(source ICefMediaSource) bool // function
	// GetDeviceInfo
	//  Asynchronously retrieves device info.
	GetDeviceInfo(callback ICefMediaSinkDeviceInfoCallback) // procedure
}

// TCefMediaSink Parent: TCefBaseRefCounted
//
//	Represents a sink to which media can be routed. Instances of this object are retrieved via ICefMediaObserver.OnSinks. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_t))</a>
type TCefMediaSink struct {
	TCefBaseRefCounted
}

// MediaSinkRef -> ICefMediaSink
var MediaSinkRef mediaSink

// mediaSink TCefMediaSink Ref
type mediaSink uintptr

func (m *mediaSink) UnWrap(data uintptr) ICefMediaSink {
	var resultCefMediaSink uintptr
	CEF().SysCallN(1073, uintptr(data), uintptr(unsafePointer(&resultCefMediaSink)))
	return AsCefMediaSink(resultCefMediaSink)
}

func (m *TCefMediaSink) GetId() string {
	r1 := CEF().SysCallN(1068, m.Instance())
	return GoStr(r1)
}

func (m *TCefMediaSink) GetName() string {
	r1 := CEF().SysCallN(1069, m.Instance())
	return GoStr(r1)
}

func (m *TCefMediaSink) GetIconType() TCefMediaSinkIconType {
	r1 := CEF().SysCallN(1067, m.Instance())
	return TCefMediaSinkIconType(r1)
}

func (m *TCefMediaSink) IsCastSink() bool {
	r1 := CEF().SysCallN(1070, m.Instance())
	return GoBool(r1)
}

func (m *TCefMediaSink) IsDialSink() bool {
	r1 := CEF().SysCallN(1072, m.Instance())
	return GoBool(r1)
}

func (m *TCefMediaSink) IsCompatibleWith(source ICefMediaSource) bool {
	r1 := CEF().SysCallN(1071, m.Instance(), GetObjectUintptr(source))
	return GoBool(r1)
}

func (m *TCefMediaSink) GetDeviceInfo(callback ICefMediaSinkDeviceInfoCallback) {
	CEF().SysCallN(1066, m.Instance(), GetObjectUintptr(callback))
}
