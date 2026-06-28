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
	"encoding/json"
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	internalIPC "github.com/energye/energy/v3/internal/ipc"
	"github.com/energye/energy/v3/logger"
	"strings"
	"unsafe"
)

const (
	internalPostMessageName                      = "postMessage"
	internalAddEventListenerName                 = "addEventListener"
	internalPostMessageWithAdditionalObjectsName = "postMessageWithAdditionalObjects"
	internalRenderProcessMessageName             = "message"
	internalExecuteScriptName                    = "executeScript"
	internalExecuteScriptResultName              = "executeScriptResult"
)

func (m *Application) initDefaultEvent() {
	logger.Debug("Application.initDefaultEvent")
	m.ICefApplication.SetOnContextCreated(m.applicationOnContextCreated)
	m.ICefApplication.SetOnProcessMessageReceived(m.applicationOnProcessMessageReceived)
	m.ICefApplication.SetOnRegCustomSchemes(m.applicationOnRegCustomSchemes)
	m.ICefApplication.SetOnContextInitialized(m.applicationOnContextInitialized)
}

func (m *Application) SetOnContextCreated(fn cef.TOnContextCreatedEvent) {
	m.onContextCreated = fn
}

func (m *Application) SetOnProcessMessageReceived(fn cef.TOnProcessMessageReceivedEvent) {
	m.onProcessMessageReceived = fn
}

func (m *Application) SetOnRegCustomSchemes(fn cef.TOnRegisterCustomSchemesEvent) {
	m.onRegisterCustomSchemes = fn
}

func (m *Application) SetOnContextInitializedEvent(fn cef.TOnContextInitializedEvent) {
	m.onContextInitialized = fn
}

func (m *Application) applicationOnContextCreated(browser cef.ICefBrowser, frame cef.ICefFrame, context cef.ICefv8Context) {
	logger.Debug("Application.OnContextCreated")
	m.postMessage = makePostMessageObject(browser, frame, context)
	frame.ExecuteJavaScript(string(internalIPC.JSIPC), "", 0)
	frame.ExecuteJavaScript(string(internalIPC.JSDrag), "", 0)
	if m.onContextCreated != nil {
		m.onContextCreated(browser, frame, context)
	}
}

func (m *Application) applicationOnProcessMessageReceived(browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, handled *bool) {
	name := message.GetName()
	logger.Debug("Application.OnProcessMessageReceived name:", name)
	defer func() {
		message.Release()
	}()
	if m.postMessage != nil && name == internalPostMessageName {
		callback, ok := m.postMessage.eventCallbacks[internalRenderProcessMessageName]
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
	} else if name == internalExecuteScriptName {
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
		scriptMessage := tExecuteScriptMessage{}
		err := json.Unmarshal(messageDataBytes, &scriptMessage)
		if err != nil {
			return
		}
		executeScriptResult := tExecuteScriptResultMessage{Id: scriptMessage.Id}
		if v8ctx.Enter() {
			var (
				retval    cef.ICefv8Value
				exception cef.ICefV8Exception
			)

			defer func() {
				if !IsNil(exception) {
					retval.Release()
				}
				if !IsNil(exception) {
					exception.Release()
				}
			}()
			evalOK := v8ctx.Eval(scriptMessage.Script, "", 0, &retval, &exception)
			if evalOK {
				value := v8ValueToJSON(retval)
				result, err := json.Marshal(value)
				if err != nil {
					executeScriptResult.Error = err.Error()
				} else {
					executeScriptResult.Data = string(result)
				}
			}

			if !IsNil(exception) {
				executeScriptResult.Error = exception.GetMessage()
			}
			v8ctx.Exit()
		} else {
			executeScriptResult.Error = "Failed to enter V8Context"
		}
		data, _ := json.Marshal(executeScriptResult)
		sendBrowserProcessMessage(frame, internalExecuteScriptResultName, data, nil)
	} else {
		if m.onProcessMessageReceived != nil {
			m.onProcessMessageReceived(browser, frame, sourceProcess, message, handled)
		}
	}
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
	if m.onRegisterCustomSchemes != nil {
		m.onRegisterCustomSchemes(registrar)
	}
}

func (m *Application) applicationOnContextInitialized() {
	logger.Debug("Application.OnContextInitialized viewsWindows count: ", len(m.windowList), "ProcessType:", ProcessType(m.ProcessType()))
	if mainWindow, ok := m.windowList[0]; ok {
		if window, ok := mainWindow.(IViewsBrowser); ok {
			window.buildViewsBrowser(nil, window)
			window.CreateTopLevelWindow()
		}
	}
	if m.onContextInitialized != nil {
		m.onContextInitialized()
	}
}
