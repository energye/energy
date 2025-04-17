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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefMediaSink
//
//	Callback interface for ICefMediaSink.GetDeviceInfo. The func (m *ICefMediaRouter)s of
//	this interface will be called on the browser process UI thread.
//	<para><see cref="uCEFTypes|TCefMediaSinkDeviceInfoCallback">Implements TCefMediaSinkDeviceInfoCallback</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_device_info_callback_t)</see></para>
type ICefMediaSink struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// MediaSinkRef -> ICefMediaSink
var MediaSinkRef mediaSink

type mediaSink uintptr

func (*mediaSink) UnWrap(data *ICefMediaSink) *ICefMediaSink {
	var result uintptr
	imports.Proc(def.MediaSinkRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSink{instance: getInstance(result)}
	}
	return nil
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

// / Returns the ID for this sink.
func (m *ICefMediaSink) GetId() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.MediaSink_GetId).Call(m.Instance())
	return api.GoStr(r1)
}

// / Returns the name of this sink.
func (m *ICefMediaSink) GetName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.MediaSink_GetName).Call(m.Instance())
	return api.GoStr(r1)
}

// / Returns the icon type for this sink.
func (m *ICefMediaSink) GetIconType() consts.TCefMediaSinkIconType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.MediaSink_GetIconType).Call(m.Instance())
	return consts.TCefMediaSinkIconType(r1)
}

// / Asynchronously retrieves device info.
func (m *ICefMediaSink) GetDeviceInfo(callback *ICefMediaSinkDeviceInfoCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MediaSink_GetDeviceInfo).Call(m.Instance(), callback.Instance())
}

// / Returns true (1) if this sink accepts content via Cast.
func (m *ICefMediaSink) IsCastSink() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.MediaSink_IsCastSink).Call(m.Instance())
	return api.GoBool(r1)
}

// / Returns true (1) if this sink accepts content via DIAL.
func (m *ICefMediaSink) IsDialSink() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.MediaSink_IsDialSink).Call(m.Instance())
	return api.GoBool(r1)
}

// / Returns true (1) if this sink is compatible with |source|.
func (m *ICefMediaSink) IsCompatibleWith(source *ICefMediaSource) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.MediaSink_IsCompatibleWith).Call(m.Instance(), source.Instance())
	return api.GoBool(r1)
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

func (m *TCefMediaSinkArray) Get(index int) *ICefMediaSink {
	if !m.IsValid() && index >= 0 && index < int(m.count) {
		return nil
	}
	var result uintptr
	imports.Proc(def.MediaSinkArray_Get).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSink{instance: getInstance(result)}
	}
	return nil
}
