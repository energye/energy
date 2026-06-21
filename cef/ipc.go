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
	"github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/logger"
	lclTypes "github.com/energye/lcl/types"
	"unsafe"
)

type tPostMessage struct {
	energy                               cef.ICefv8Value
	funcHandle                           cef.IEngV8Handler
	postMessageFunc                      cef.ICefv8Value
	addEventListenerFunc                 cef.ICefv8Value
	postMessageWithAdditionalObjectsFunc cef.ICefv8Value
	eventCallbacks                       map[string]cef.ICefv8Value
}

func makePostMessageObject(browser cef.ICefBrowser, frame cef.ICefFrame, context cef.ICefv8Context) *tPostMessage {
	m := &tPostMessage{eventCallbacks: make(map[string]cef.ICefv8Value)}
	m.energy = cef.V8ValueRef.NewObject(nil, nil)

	m.funcHandle = cef.NewEngV8Handler()
	m.funcHandle.SetOnV8Execute(m.funcHandleOnV8Execute)
	intfFuncHandle := cef.AsEngV8Handler(m.funcHandle.AsIntfV8Handler())

	m.postMessageFunc = cef.V8ValueRef.NewFunction(internalPostMessageName, intfFuncHandle)
	m.addEventListenerFunc = cef.V8ValueRef.NewFunction(internalAddEventListenerName, intfFuncHandle)
	m.postMessageWithAdditionalObjectsFunc = cef.V8ValueRef.NewFunction(internalPostMessageWithAdditionalObjectsName, intfFuncHandle)

	m.energy.SetValueByKey(internalPostMessageName, m.postMessageFunc, types.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energy.SetValueByKey(internalAddEventListenerName, m.addEventListenerFunc, types.V8_PROPERTY_ATTRIBUTE_READONLY)
	m.energy.SetValueByKey(internalPostMessageWithAdditionalObjectsName, m.postMessageWithAdditionalObjectsFunc, types.V8_PROPERTY_ATTRIBUTE_READONLY)

	context.GetGlobal().GetValueByKey("chrome").
		SetValueByKey("energy", m.energy, types.V8_PROPERTY_ATTRIBUTE_READONLY)
	return m
}

func (m *tPostMessage) funcHandleOnV8Execute(name string, object cef.ICefv8Value, arguments cef.ICefv8ValueArray,
	retval *cef.ICefv8Value, exception *string) bool {
	defer arguments.Free()
	size := arguments.Count()
	logger.Debug("PostMessageHandle.OnV8Execute name:", name, "args-size:", size)
	if name == internalPostMessageName {
		if size == 1 {
			if messageData := m.getPostMessageData(arguments); messageData != "" {
				v8Context := cef.V8ContextRef.Current()
				frame := v8Context.GetFrame()
				defer func() {
					frame.Release()
					v8Context.Release()
				}()
				sendBrowserProcessMessage(frame, internalPostMessageName, []byte(messageData), nil)
				//*retval = cef.V8ValueRef.NewUndefined()
				return true
			}
		}
	} else if name == internalAddEventListenerName {
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
				if callbackName == internalRenderProcessMessageName {
					m.eventCallbacks[callbackName] = cef.V8ValueRef.UnWrap(callbackV8Value.Wrap())
					//*retval = cef.V8ValueRef.NewUndefined()
					return true
				}
			}
		}
	} else if name == internalPostMessageWithAdditionalObjectsName {
		if size == 2 {
			messageData := m.getPostMessageData(arguments)
			objectsFileData := m.getPostMessageWithAdditionalObjectsFileData(arguments)
			if messageData != "" {
				v8Context := cef.V8ContextRef.Current()
				frame := v8Context.GetFrame()
				defer func() {
					frame.Release()
					v8Context.Release()
				}()
				var (
					objectsFileDataBytes []byte
					err                  error
				)
				if len(objectsFileData) > 0 {
					objectsFileDataBytes, err = json.Marshal(objectsFileData)
					if err != nil {
						return false
					}
				}
				sendBrowserProcessMessage(frame, internalPostMessageName, []byte(messageData),
					objectsFileDataBytes)
			}
			//*retval = cef.V8ValueRef.NewUndefined()
			return true
		}
	}
	return false
}

func (m *tPostMessage) getPostMessageData(arguments cef.ICefv8ValueArray) string {
	messageV8Value := arguments.Get(0)
	defer func() {
		messageV8Value.Release()
	}()
	if messageV8Value.IsString() {
		return messageV8Value.GetStringValue()
	}
	return ""
}

type TObjectFile struct {
	Name         string `json:"name"`
	LastModified int64  `json:"last_modified"`
	Size         uint32 `json:"size"`
}

func (m *tPostMessage) getPostMessageWithAdditionalObjectsFileData(arguments cef.ICefv8ValueArray) []TObjectFile {
	objectsV8Value := arguments.Get(1)
	defer func() {
		objectsV8Value.Release()
	}()
	if !objectsV8Value.IsArray() {
		return nil
	}
	var files []TObjectFile
	for i := int32(0); i < objectsV8Value.GetArrayLength(); i++ {
		fileItemV8Value := objectsV8Value.GetValueByIndex(i)
		fileName := fileItemV8Value.GetValueByKey("name")
		fileLastModified := fileItemV8Value.GetValueByKey("lastModified")
		fileSize := fileItemV8Value.GetValueByKey("size")
		files = append(files, TObjectFile{
			Name:         fileName.GetStringValue(),
			LastModified: lclTypes.TDateTime(fileLastModified.GetDoubleValue()).ToTime().UnixMilli(),
			Size:         fileSize.GetUIntValue(),
		})
		fileItemV8Value.Release()
	}
	return files
}

func sendBrowserProcessMessage(frame cef.ICefFrame, name string, data []byte, objects []byte) {
	processMessage := cef.ProcessMessageRef.New(name)
	messageArgumentList := processMessage.GetArgumentList()
	if len(data) > 0 {
		dataPtr := uintptr(unsafe.Pointer(&data[0]))
		dataBin := cef.BinaryValueRef.New(dataPtr, uint32(len(data)))
		defer dataBin.Release()
		messageArgumentList.SetBinary(0, dataBin)
	}
	if len(objects) > 0 {
		objectsPtr := uintptr(unsafe.Pointer(&objects[0]))
		objectsBin := cef.BinaryValueRef.New(objectsPtr, uint32(len(objects)))
		defer objectsBin.Release()
		messageArgumentList.SetBinary(1, objectsBin)
	}
	frame.SendProcessMessage(types.PID_BROWSER, processMessage)
	messageArgumentList.Clear()
	messageArgumentList.Release()
	processMessage.Release()
}
