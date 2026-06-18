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
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/core"
	internalIPC "github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/logger"
	"strings"
	"unsafe"
)

func (m *Application) initDefaultEvent() {
	logger.Debug("Application.initDefaultEvent")
	m.SetOnContextCreated(m.applicationOnContextCreated)
	m.SetOnProcessMessageReceived(m.applicationOnProcessMessageReceived)
	m.SetOnWebKitInitialized(m.applicationOnWebKitInitialized)
	m.SetOnRegCustomSchemes(m.applicationOnRegCustomSchemes)
}

func (m *Application) applicationOnContextCreated(browser cef.ICefBrowser, frame cef.ICefFrame, context cef.ICefv8Context) {
	logger.Debug("Application.OnContextCreated")
	m.postMessage = makePostMessageObject(browser, frame, context)
	frame.ExecuteJavaScript(string(internalIPC.JSIPC), "", 0)
}

func (m *Application) applicationOnProcessMessageReceived(browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, handled *bool) {
	name := message.GetName()
	logger.Debug("Application.OnProcessMessageReceived name:", name)
	defer func() {
		message.Release()
	}()
	if m.postMessage != nil && name == core.PostMessageName {
		callback, ok := m.postMessage.eventCallbacks[core.RenderProcessMessageName]
		if !ok {
			return
		}
		args := message.GetArgumentList()
		dataBin := args.GetBinary(0)
		v8ctx := frame.GetV8Context()
		defer func() {
			dataBin.Release()
			args.Release()
			v8ctx.Release()
		}()
		messageDataBytes := make([]byte, int(dataBin.GetSize()))
		dataBin.GetData(uintptr(unsafe.Pointer(&messageDataBytes[0])), dataBin.GetSize(), 0)
		messageData := string(messageDataBytes)
		if v8ctx.Enter() {
			callFuncArgs := cef.NewCefv8ValueArray(0, 0)
			defer callFuncArgs.Free()
			callFuncArgs.Add(cef.V8ValueRef.NewString(messageData))
			callback.ExecuteFunctionWithContext(v8ctx, nil, callFuncArgs)
			v8ctx.Exit()
		}
	}
}

func (m *Application) applicationOnWebKitInitialized() {
	logger.Debug("Application.OnWebKitInitialized")
}

func (m *Application) applicationOnRegCustomSchemes(registrar cef.ICefSchemeRegistrarRef) {
	logger.Debug("Application.OnRegCustomSchemes")
	gApp := application.GApplication
	if gApp == nil || gApp.LocalLoad == nil {
		return
	}
	switch strings.ToUpper(gApp.LocalLoad.Scheme) {
	case "HTTP", "HTTPS", "FILE", "FTP", "ABOUT", "DATA":
		return
	}
	registrar.AddCustomScheme(gApp.LocalLoad.Scheme,
		cefTypes.CEF_SCHEME_OPTION_STANDARD|cefTypes.CEF_SCHEME_OPTION_CORS_ENABLED|cefTypes.CEF_SCHEME_OPTION_SECURE|cefTypes.CEF_SCHEME_OPTION_FETCH_ENABLED)
}
