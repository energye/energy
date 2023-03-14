// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// CEF Frame
package cef

import (
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefFrame) Instance() uintptr {
	return uintptr(m.instance)
}

// Undo 撤销操作
func (m *ICefFrame) Undo() {
	imports.Proc(internale_CEFFrame_Undo).Call(m.Instance())
}

// Redo 恢复
func (m *ICefFrame) Redo() {
	imports.Proc(internale_CEFFrame_Redo).Call(m.Instance())
}

// Cut 剪切
func (m *ICefFrame) Cut() {
	imports.Proc(internale_CEFFrame_Cut).Call(m.Instance())
}

// Copy 复制
func (m *ICefFrame) Copy() {
	imports.Proc(internale_CEFFrame_Copy).Call(m.Instance())
}

// Paste 粘贴
func (m *ICefFrame) Paste() {
	imports.Proc(internale_CEFFrame_Paste).Call(m.Instance())
}

// Del 删除
func (m *ICefFrame) Del() {
	imports.Proc(internale_CEFFrame_Del).Call(m.Instance())
}

// SelectAll 选择所有
func (m *ICefFrame) SelectAll() {
	imports.Proc(internale_CEFFrame_SelectAll).Call(m.Instance())
}

// ViewSource 显示源码
func (m *ICefFrame) ViewSource() {
	imports.Proc(internale_CEFFrame_ViewSource).Call(m.Instance())
}

// LoadUrl 加载URL
func (m *ICefFrame) LoadUrl(url string) {
	imports.Proc(internale_CEFFrame_LoadUrl).Call(m.Instance(), api.PascalStr(url))
}

// ExecuteJavaScript 执行JS
func (m *ICefFrame) ExecuteJavaScript(code, scriptUrl string, startLine int32) {
	imports.Proc(internale_CEFFrame_ExecuteJavaScript).Call(m.Instance(), api.PascalStr(code), api.PascalStr(scriptUrl), uintptr(startLine))
}

// IsValid 该Frame是否有效
func (m *ICefFrame) IsValid() bool {
	r1, _, _ := imports.Proc(internale_CEFFrame_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// IsMain 是否为主Frame
func (m *ICefFrame) IsMain() bool {
	r1, _, _ := imports.Proc(internale_CEFFrame_IsMain).Call(m.Instance())
	return api.GoBool(r1)
}

// IsFocused 是否已获取焦点
func (m *ICefFrame) IsFocused() bool {
	r1, _, _ := imports.Proc(internale_CEFFrame_IsFocused).Call(m.Instance())
	return api.GoBool(r1)
}

//// SendProcessMessageByIPC 发送进程消息
//func (m *ICefFrame) SendProcessMessageByIPC(targetProcess CefProcessId, processMessage *ipc.ICefProcessMessage) ProcessMessageError {
//	if processMessage == nil || processMessage.Name == "" || processMessage.ArgumentList == nil {
//		return PMErr_REQUIRED_PARAMS_IS_NULL
//	} else if ipc.InternalIPCNameCheck(processMessage.Name) {
//		return PMErr_NAME_CANNOT_USED
//	}
//	data := processMessage.ArgumentList.Package()
//	r1 := _CEFFrame_SendProcessMessageByIPC(m.Browser.Identifier(), m.Id, processMessage.Name, targetProcess, int32(processMessage.ArgumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
//	return ProcessMessageError(r1)
//}
//
//func _CEFFrame_SendProcessMessageByIPC(browseId int32, frameId int64, name string, targetProcess CefProcessId, itemLength int32, data, dataLen uintptr) uintptr {
//	r1, _, _ := imports.Proc(internale_CEFFrame_SendProcessMessageByIPC).Call(uintptr(browseId), uintptr(unsafe.Pointer(&frameId)), api.PascalStr(name), uintptr(targetProcess), uintptr(itemLength), data, dataLen)
//	return r1
//}

// SendProcessMessage 发送进程消息
func (m *ICefFrame) SendProcessMessage(targetProcess CefProcessId, message *ICefProcessMessage) {
	m.SendProcessMessageForIPC(targetProcess, message)
}

func (m *ICefFrame) SendProcessMessageForIPC(targetProcess CefProcessId, message ipc.ICefProcessMessageIPC) {
	imports.Proc(internale_CEFFrame_SendProcessMessage).Call(m.Instance(), targetProcess.ToPtr(), message.Instance())
}

func (m *ICefFrame) LoadRequest(request *ICefRequest) {
	if m == nil || request == nil {
		return
	}
	imports.Proc(internale_CEFFrame_LoadRequest).Call(m.Instance(), request.Instance())
}

func (m *ICefFrame) Browser() *ICefBrowser {
	var result uintptr
	imports.Proc(internale_CEFFrame_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBrowser{instance: unsafe.Pointer(result)}
}

func (m *ICefFrame) V8Context() *ICefV8Context {
	var result uintptr
	imports.Proc(internale_CEFFrame_GetV8Context).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Context{instance: unsafe.Pointer(result)}
}

func (m *ICefFrame) Identifier() int64 {
	var result uintptr
	imports.Proc(internale_CEFFrame_Identifier).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return int64(result)
}

func (m *ICefFrame) Name() string {
	r1, _, _ := imports.Proc(internale_CEFFrame_Name).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefFrame) Url() string {
	r1, _, _ := imports.Proc(internale_CEFFrame_Url).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefFrame) Parent() *ICefFrame {
	var result uintptr
	imports.Proc(internale_CEFFrame_Parent).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefFrame{instance: unsafe.Pointer(result)}
}

func (m *ICefFrame) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// FrameRef -> *ICefFrame
var FrameRef frameRef

// frameRef
type frameRef uintptr

func (m *frameRef) UnWrap(data *ICefFrame) *ICefFrame {
	var result uintptr
	imports.Proc(internale_CEFFrameRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
	return data
}
