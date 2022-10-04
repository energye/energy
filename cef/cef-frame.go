//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

//type ICefFrame
type ICefFrame struct {
	Browser *ICefBrowser
	Name    string
	Url     string
	Id      int64
}

type cefFrame struct {
	Name       uintptr
	Url        uintptr
	Identifier uintptr
}

func (m *ICefFrame) Undo() {
	var frameId = m.Id
	Proc("CEFFrame_Undo").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) Redo() {
	var frameId = m.Id
	Proc("CEFFrame_Redo").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) Cut() {
	var frameId = m.Id
	Proc("CEFFrame_Cut").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) Copy() {
	var frameId = m.Id
	Proc("CEFFrame_Copy").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) Paste() {
	var frameId = m.Id
	Proc("CEFFrame_Paste").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) Del() {
	var frameId = m.Id
	Proc("CEFFrame_Del").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) SelectAll() {
	var frameId = m.Id
	Proc("CEFFrame_SelectAll").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) ViewSource() {
	var frameId = m.Id
	Proc("CEFFrame_ViewSource").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

func (m *ICefFrame) LoadUrl(url string) {
	var frameId = m.Id
	Proc("CEFFrame_LoadUrl").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)), api.GoStrToDStr(url))
}

func (m *ICefFrame) ExecuteJavaScript(code, scriptUrl string, startLine int32) {
	var frameId = m.Id
	Proc("CEFFrame_ExecuteJavaScript").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)), api.GoStrToDStr(code), api.GoStrToDStr(scriptUrl), uintptr(startLine))
}

func (m *ICefFrame) IsValid() bool {
	var frameId = m.Id
	r1, _, _ := Proc("CEFFrame_IsValid").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
	return api.DBoolToGoBool(r1)
}

func (m *ICefFrame) IsMain() bool {
	var frameId = m.Id
	r1, _, _ := Proc("CEFFrame_IsMain").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
	return api.DBoolToGoBool(r1)
}

func (m *ICefFrame) IsFocused() bool {
	var frameId = m.Id
	r1, _, _ := Proc("CEFFrame_IsFocused").Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
	return api.DBoolToGoBool(r1)
}

func (m *ICefFrame) SendProcessMessage(targetProcess CefProcessId, processMessage *ipc.ICefProcessMessage) ProcessMessageError {
	if processMessage == nil || processMessage.Name == "" || processMessage.ArgumentList == nil {
		return PMErr_REQUIRED_PARAMS_IS_NULL
	} else if ipc.InternalIPCNameCheck(processMessage.Name) {
		return PMErr_NAME_CANNOT_USED
	}
	data := processMessage.ArgumentList.Package()
	r1 := _CEFFrame_SendProcessMessage(m.Browser.Identifier(), m.Id, processMessage.Name, targetProcess, int32(processMessage.ArgumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return ProcessMessageError(r1)
}

func _CEFFrame_SendProcessMessage(browseId int32, frameId int64, name string, targetProcess CefProcessId, itemLength int32, data, dataLen uintptr) uintptr {
	r1, _, _ := Proc("CEFFrame_SendProcessMessage").Call(uintptr(browseId), uintptr(unsafe.Pointer(&frameId)), api.GoStrToDStr(name), uintptr(targetProcess), uintptr(itemLength), data, dataLen)
	return r1
}
