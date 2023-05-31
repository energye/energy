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
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"sync"
	"unsafe"
)

// IChromium 组件接口
type IChromium interface {
	IChromiumProc
	IChromiumEvent
}

// TCEFChromium 组件
type TCEFChromium struct {
	*lcl.TComponent
	instance      unsafe.Pointer
	cfg           *tCefChromiumConfig
	browser       *ICefBrowser
	idBrowsers    map[int32]*ICefBrowser
	emitLock      *sync.Mutex
	browserHandle types.HWND
	widgetHandle  types.HWND
	renderHandle  types.HWND
	initialized   bool
}

// NewChromium 创建一个新的 TCEFChromium
func NewChromium(owner lcl.IComponent, config *tCefChromiumConfig) IChromium {
	m := new(TCEFChromium)
	if config != nil {
		m.cfg = config
	} else {
		m.cfg = NewChromiumConfig()
	}
	m.instance = unsafe.Pointer(_CEFChromium_Create(lcl.CheckPtr(owner), uintptr(unsafe.Pointer(m.cfg))))
	m.emitLock = new(sync.Mutex)
	m.initDefault()
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

// ExecuteJavaScript
// 执行JS代码
//
// code: js代码
//
// scriptURL: js脚本地址 默认about:blank
//
// startLine: js脚本启始执行行号
func (m *TCEFChromium) ExecuteJavaScript(code, scriptURL string, startLine int32) {
	_CEFChromium_ExecuteJavaScript(uintptr(m.instance), code, scriptURL, startLine)
}
