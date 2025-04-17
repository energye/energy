//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Chromium 组件

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// IChromium 组件接口
type IChromium interface {
	lcl.IObject
	IChromiumProc
	IChromiumEvent
}

// TCEFChromium 组件
type TCEFChromium struct {
	*lcl.TComponent
	instance      unsafe.Pointer
	options       *TChromiumOptions
	fontOptions   *TChromiumFontOptions
	config        *TCefChromiumConfig
	browser       *ICefBrowser
	idBrowsers    map[int32]*ICefBrowser
	browserHandle types.HWND
	widgetHandle  types.HWND
	renderHandle  types.HWND
	initialized   bool
	isClosing     bool
}

// NewChromium 创建一个新的 TCEFChromium
func NewChromium(owner lcl.IComponent, config *TCefChromiumConfig) IChromium {
	m := new(TCEFChromium)
	if config != nil {
		m.config = config
	} else {
		m.config = NewChromiumConfig()
	}
	r1, _, _ := imports.Proc(def.CEFChromium_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	m.initDefault()
	m.options = NewChromiumOptions(m)
	m.fontOptions = NewChromiumFontOptions(m)
	return m
}

// 默认的初始配置
func (m *TCEFChromium) initDefault() {
	//通过设置这些首选项，可以降低/避免WebRTC的IP泄漏
	m.SetWebRTCIPHandlingPolicy(HpDisableNonProxiedUDP)
	m.SetWebRTCMultipleRoutes(STATE_DISABLED)
	m.SetWebRTCNonproxiedUDP(STATE_DISABLED)
}

// Instance 组件实例指针
func (m *TCEFChromium) Instance() uintptr {
	if m == nil || m.instance == nil {
		return 0
	}
	return uintptr(m.instance)
}
