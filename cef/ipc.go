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
	"github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/logger"
	"unsafe"
)

type tPostMessage struct {
	energy               cef.ICefv8Value
	postMessageHandle    cef.IEngV8Handler
	postMessageFunc      cef.ICefv8Value
	addEventListenerFunc cef.ICefv8Value
	eventCallbacks       map[string]cef.ICefv8Value
}

func makePostMessageObject(browser cef.ICefBrowser, frame cef.ICefFrame, context cef.ICefv8Context) *tPostMessage {
	m := &tPostMessage{eventCallbacks: make(map[string]cef.ICefv8Value)}
	m.energy = cef.V8ValueRef.NewObject(nil, nil)

	m.postMessageHandle = cef.NewEngV8Handler()
	m.postMessageHandle.SetOnV8Execute(m.postMessageHandleOnV8Execute)
	m.postMessageFunc = cef.V8ValueRef.NewFunction("postMessage", cef.AsEngV8Handler(m.postMessageHandle.AsIntfV8Handler()))
	m.addEventListenerFunc = cef.V8ValueRef.NewFunction("addEventListener", cef.AsEngV8Handler(m.postMessageHandle.AsIntfV8Handler()))

	m.energy.SetValueByKey("postMessage", m.postMessageFunc, types.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energy.SetValueByKey("addEventListener", m.addEventListenerFunc, types.V8_PROPERTY_ATTRIBUTE_READONLY)

	context.GetGlobal().GetValueByKey("chrome").
		SetValueByKey("energy", m.energy, types.V8_PROPERTY_ATTRIBUTE_READONLY)
	return m
}

func (m *tPostMessage) postMessageHandleOnV8Execute(name string, object cef.ICefv8Value, arguments cef.ICefv8ValueArray,
	retval *cef.ICefv8Value, exception *string) bool {
	logger.Debug("PostMessageHandle.OnV8Execute name:", name)
	defer arguments.Free()
	size := arguments.Count()
	if name == core.PostMessageName {
		if size == 1 {
			messageV8Value := arguments.Get(0)
			defer func() {
				messageV8Value.Release()
			}()
			if messageV8Value.IsString() {
				message := messageV8Value.GetStringValue()
				v8Context := cef.V8ContextRef.Current()
				frame := v8Context.GetFrame()
				defer func() {
					frame.Release()
					v8Context.Release()
				}()
				m.sendBrowserProcessMessage(frame, core.PostMessageName, []byte(message))
			}
			//*retval = cef.V8ValueRef.NewUndefined()
		}
	} else if name == "addEventListener" {
		if size == 2 {
			nameV8Value := arguments.Get(0)
			callbackV8Value := arguments.Get(1)
			defer func() {
				nameV8Value.Release()
				callbackV8Value.Release()
			}()
			if nameV8Value.IsString() && callbackV8Value.IsFunction() {
				callbackName := nameV8Value.GetStringValue()
				logger.Debug("PostMessageHandle.OnV8Execute - addEventListener callbackName:", callbackName)
				if callbackName == core.RenderProcessMessageName {
					m.eventCallbacks[callbackName] = cef.V8ValueRef.UnWrap(callbackV8Value.Wrap())
					//*retval = cef.V8ValueRef.NewUndefined()
					return true
				}
			}
		}
	}
	return false
}

func (m *tPostMessage) triggerMessageEvent() {

}

func (m *tPostMessage) sendBrowserProcessMessage(frame cef.ICefFrame, name string, data []byte) {
	processMessage := cef.ProcessMessageRef.New(name)
	messageArgumentList := processMessage.GetArgumentList()
	var dataPtr = uintptr(0)
	if len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
		dataBin := cef.BinaryValueRef.New(dataPtr, uint32(len(data)))
		defer dataBin.Release()
		messageArgumentList.SetBinary(0, dataBin)
	}
	frame.SendProcessMessage(types.PID_BROWSER, processMessage)
	messageArgumentList.Clear()
	messageArgumentList.Release()
	processMessage.Release()
}
