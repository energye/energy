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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefAudioHandler
type ICefAudioHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// TCefAudioParameters
// include/internal/cef_types.h (cef_audio_parameters_t)
type TCefAudioParameters struct {
	channelLayout   consts.TCefChannelLayout
	sampleRate      int32
	framesPerBuffer int32
}

// AudioHandlerRef -> ICefAudioHandler
var AudioHandlerRef audioHandler

type audioHandler uintptr

func (*audioHandler) New() *ICefAudioHandler {
	var result uintptr
	imports.Proc(def.CefAudioHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefAudioHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefAudioHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefAudioHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefAudioHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefAudioHandler) SetOnGetAudioParameters(fn onGetAudioParameters) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAudioHandler_OnGetAudioParameters).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefAudioHandler) SetOnAudioStreamStarted(fn onAudioStreamStarted) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAudioHandler_OnAudioStreamStarted).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefAudioHandler) SetOnAudioStreamPacket(fn onAudioStreamPacket) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAudioHandler_OnAudioStreamPacket).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefAudioHandler) SetOnAudioStreamStopped(fn onAudioStreamStopped) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAudioHandler_OnAudioStreamStopped).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefAudioHandler) SetOnAudioStreamError(fn onAudioStreamError) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefAudioHandler_OnAudioStreamError).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onGetAudioParameters func(browser *ICefBrowser, params *TCefAudioParameters) bool
type onAudioStreamStarted func(browser *ICefBrowser, params *TCefAudioParameters, channels int32)
type onAudioStreamPacket func(browser *ICefBrowser, data *uintptr, frames int32, pts int64)
type onAudioStreamStopped func(browser *ICefBrowser)
type onAudioStreamError func(browser *ICefBrowser, message string)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onGetAudioParameters:
			browse := &ICefBrowser{instance: getPtr(0)}
			params := (*TCefAudioParameters)(getPtr(1))
			result := (*bool)(getPtr(2))
			*result = fn.(onGetAudioParameters)(browse, params)
		case onAudioStreamStarted:
			browse := &ICefBrowser{instance: getPtr(0)}
			params := (*TCefAudioParameters)(getPtr(1))
			fn.(onAudioStreamStarted)(browse, params, int32(getVal(2)))
		case onAudioStreamPacket:
			browse := &ICefBrowser{instance: getPtr(0)}
			data := (*uintptr)(getPtr(1))
			frames := int32(getVal(2))
			pts := *(*int64)(getPtr(3))
			fn.(onAudioStreamPacket)(browse, data, frames, pts)
		case onAudioStreamStopped:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onAudioStreamStopped)(browse)
		case onAudioStreamError:
			browse := &ICefBrowser{instance: getPtr(0)}
			fn.(onAudioStreamError)(browse, api.GoStr(getVal(1)))
		default:
			return false
		}
		return true
	})
}
