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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// ICEFChromiumBrowser
//
//	CEFChromium浏览器接口
type ICEFChromiumBrowser interface {
	SetCreateBrowserExtraInfo(windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) //
	CreateBrowser()                                                                                           // 创建浏览器
	Chromium() IChromium                                                                                      // 返回 chromium
	WindowParent() ICEFWindowParent                                                                           // 返回 chromium window 组件
	IsCreated() bool                                                                                          // 创建浏览器是否成功
}

// TCEFChromiumBrowser
//
//	CEFChromium浏览器包装结构
type TCEFChromiumBrowser struct {
	chromium     IChromium        // chromium
	windowParent ICEFWindowParent // chromium window 组件
	isCreated    bool             // chromium browser is created
	createTimer  *lcl.TTimer      // loop check and create chromium browser
	windowName   string
	context      *ICefRequestContext
	extraInfo    *ICefDictionaryValue
}

// NewChromiumBrowser
//
//	初始创建一个 chromium 浏览器
func NewChromiumBrowser(owner lcl.IWinControl, config *TCefChromiumConfig) ICEFChromiumBrowser {
	var m = new(TCEFChromiumBrowser)
	m.chromium = NewChromium(owner, config)
	m.windowParent = NewCEFWindowParent(owner)
	m.windowParent.SetParent(owner)
	m.windowParent.SetChromium(m.chromium, 0)
	m.windowParent.SetWidth(100)
	m.windowParent.SetHeight(100)
	m.windowParent.SetAlign(types.AlNone)
	m.createTimer = lcl.NewTimer(owner)
	m.createTimer.SetInterval(200)
	m.createTimer.SetOnTimer(m.checkAndCreateBrowser)
	return m
}

// checkAndCreateBrowser
//
//	创建浏览器
//	创建时如果未创建成功, 使用定时器创建直到成功
func (m *TCEFChromiumBrowser) checkAndCreateBrowser(sender lcl.IObject) {
	if m.isCreated || m.chromium == nil || m.createTimer == nil {
		return
	}
	m.createTimer.SetEnabled(false)
	if m.isCreated { // 成功创建 释放定时器
		m.createTimer.Free()
		m.createTimer = nil
		return
	}
	m.chromium.Initialized()
	m.isCreated = m.chromium.CreateBrowser(m.windowParent, m.windowName, m.context, m.extraInfo)
	if !m.isCreated {
		m.createTimer.SetEnabled(true)
	} else {
		m.windowParent.UpdateSize()
	}
}

func (m *TCEFChromiumBrowser) SetCreateBrowserExtraInfo(windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}

// CreateBrowser
//
//	创建浏览器
//	创建时如果未创建成功, 使用定时任务创建直到成功
func (m *TCEFChromiumBrowser) CreateBrowser() {
	m.checkAndCreateBrowser(nil)
}

// Chromium
//
//	返回 chromium
func (m *TCEFChromiumBrowser) Chromium() IChromium {
	return m.chromium
}

// WindowParent
//
//	返回 chromium window 组件
func (m *TCEFChromiumBrowser) WindowParent() ICEFWindowParent {
	return m.windowParent
}

// IsCreated
//
//	创建浏览器是否成功
func (m *TCEFChromiumBrowser) IsCreated() bool {
	return m.isCreated
}
