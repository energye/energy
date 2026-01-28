//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import (
	"bytes"
	"encoding/json"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	gEvaluateScriptEventID     int
	gNextEvaluateScriptEventID = func() int {
		gEvaluateScriptEventID++
		return gEvaluateScriptEventID
	}
	gEvaluateScriptEventCallback = sync.Map{}
)

// global browser id
var globalBrowserID uint32

// return next browser id
func getNextBrowserID() uint32 {
	atomic.AddUint32(&globalBrowserID, 1)
	return globalBrowserID
}

type IWebview interface {
	lcl.ICustomPanel
	// SetWindow 设置webview的窗口实例，并初始化相关回调函数
	// window - 窗口接口实例，用于承载webview内容
	SetWindow(window window.IWindow)
	// 设置当前 browser 的自定义资源加载方试
	SetLocalLoad(localLoad application.LocalLoad)
	LocalLoadResource() *application.LocalLoadResource
	// 更新当前 browser 配置
	UpdateBrowserOptions()
	SetParent(window lcl.IWinControl)
	CreateBrowser()
	BrowserId() uint32
	SendMessage(payload []byte)
	Close()
	SetDefaultURL(url string)
	LoadURL(url string)
	ExecuteScript(javaScript string)
	ExecuteScriptCallback(script string, callback TOnEvaluateScriptCallback)
	SetWidth(v int32)
	SetHeight(v int32)
	SetBoundsRect(value types.TRect)
	SetBounds(left int32, top int32, width int32, height int32)
	SetOnBrowserAfterCreated(fn lcl.TNotifyEvent)
	SetOnResourceRequest(fn TOnResourceRequestEvent)
	SetOnProcessMessage(fn TOnProcessMessageEvent)
	SetOnLoadChange(fn TOnLoadChangeEvent)
	SetOnContextMenu(fn TOnContextMenuEvent)
	SetOnContextMenuCommand(fn TOnContextMenuCommandEvent)
	SetOnPopupWindow(fn TOnPopupWindowEvent)
}

type TEnergyWebview struct {
	localLoad *application.LocalLoadResource
}

func (m *TEnergyWebview) SetLocalLoad(localLoad application.LocalLoad) {
	m.localLoad = application.NewLocalLoadResource(&localLoad)
}

func (m *TEnergyWebview) LocalLoadResource() *application.LocalLoadResource {
	return m.localLoad
}

func (m *TWebview) createEnergyJavasScript() {
	jsCode := &bytes.Buffer{}
	var envJS = func(json string) {
		jsCode.WriteString(`window.energy.setOptionsEnv(`)
		jsCode.WriteString(json)
		jsCode.WriteString(`);`)
	}
	optionsJSON, err := json.Marshal(gApplication.Options)
	if err == nil {
		envJS(string(optionsJSON))
	}
	browser := make(map[string]any)
	browser["id"] = m.BrowserId()
	env := make(map[string]any)
	env["frameWidth"] = frameWidth
	env["frameHeight"] = frameHeight
	env["frameCorner"] = frameCorner
	env["os"] = runtime.GOOS
	env["browser"] = browser
	envJSON, err := json.Marshal(env)
	if err == nil {
		envJS(string(envJSON))
	}
	m.ExecuteScript(jsCode.String())
	m.ExecuteScript(`window.energy.drag.setup();`)
}

type TLoadChange int32

const (
	LcStart TLoadChange = iota
	LcLoading
	LcFinish
)

type TOnProcessMessageEvent func(message string)
type TOnResourceRequestEvent func(url, path, method string, header map[string]string) (resource string, ok bool)
type TOnLoadChangeEvent func(url, title string, load TLoadChange)
type TOnContextMenuEvent func(contextMenu *TContextMenuItem)
type TOnContextMenuCommandEvent func(commandId int32)
type TOnPopupWindowEvent func(targetURL string) bool
type TOnEvaluateScriptCallback func(result string, err string)
