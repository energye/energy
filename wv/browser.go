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
	"github.com/energye/energy/v3/core"
	"github.com/energye/lcl/lcl"
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

// IWebview alias -> core.IBrowser
type IWebview interface {
	lcl.IWinControl
	core.IBrowser
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
