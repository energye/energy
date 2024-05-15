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

// ICefMediaSource Parent: ICefBaseRefCounted
//
//	Represents a source from which media can be routed. Instances of this object are retrieved via ICefMediaRouter.GetSource. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_source_t))</a>
type ICefMediaSource interface {
	ICefBaseRefCounted
	// GetId
	//  Returns the ID (media source URN or URL) for this source.
	GetId() string // function
	// IsCastSource
	//  Returns true (1) if this source outputs its content via Cast.
	IsCastSource() bool // function
	// IsDialSource
	//  Returns true (1) if this source outputs its content via DIAL.
	IsDialSource() bool // function
}

// TCefMediaSource Parent: TCefBaseRefCounted
//
//	Represents a source from which media can be routed. Instances of this object are retrieved via ICefMediaRouter.GetSource. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_source_t))</a>
type TCefMediaSource struct {
	TCefBaseRefCounted
}

// MediaSourceRef -> ICefMediaSource
var MediaSourceRef mediaSource

// mediaSource TCefMediaSource Ref
type mediaSource uintptr

func (m *mediaSource) UnWrap(data uintptr) ICefMediaSource {
	var resultCefMediaSource uintptr
	CEF().SysCallN(1077, uintptr(data), uintptr(unsafePointer(&resultCefMediaSource)))
	return AsCefMediaSource(resultCefMediaSource)
}

func (m *TCefMediaSource) GetId() string {
	r1 := CEF().SysCallN(1074, m.Instance())
	return GoStr(r1)
}

func (m *TCefMediaSource) IsCastSource() bool {
	r1 := CEF().SysCallN(1075, m.Instance())
	return GoBool(r1)
}

func (m *TCefMediaSource) IsDialSource() bool {
	r1 := CEF().SysCallN(1076, m.Instance())
	return GoBool(r1)
}
