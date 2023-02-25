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

// ICefFrame
// Html <iframe></iframe>
type ICefFrame struct {
	instance unsafe.Pointer
	Browser  *ICefBrowser
	Name     string
	Url      string
	Id       int64
}

type cefFrame struct {
	Name       uintptr
	Url        uintptr
	Identifier uintptr
}

// TCEFFrame Frame 集合
type TCEFFrame map[int64]*ICefFrame

func (m TCEFFrame) GetByFrameId(frameId int64) *ICefFrame {
	if m != nil {
		if frame, ok := m[frameId]; ok {
			return frame
		}
	}
	return nil
}

// Instance 实例
func (m *ICefFrame) Instance() uintptr {
	return uintptr(m.instance)
}

// ---- CEF Frame for frameId ----

// Undo 撤销操作
func (m *ICefFrame) Undo() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_Undo).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// Redo 恢复
func (m *ICefFrame) Redo() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_Redo).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// Cut 剪切
func (m *ICefFrame) Cut() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_Cut).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// Copy 复制
func (m *ICefFrame) Copy() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_Copy).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// Paste 粘贴
func (m *ICefFrame) Paste() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_Paste).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// Del 删除
func (m *ICefFrame) Del() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_Del).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// SelectAll 选择所有
func (m *ICefFrame) SelectAll() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_SelectAll).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// ViewSource 显示源码
func (m *ICefFrame) ViewSource() {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_ViewSource).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
}

// LoadUrl 加载URL
func (m *ICefFrame) LoadUrl(url string) {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_LoadUrl).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)), api.PascalStr(url))
}

// ExecuteJavaScript 执行JS
func (m *ICefFrame) ExecuteJavaScript(code, scriptUrl string, startLine int32) {
	var frameId = m.Id
	imports.Proc(internale_CEFFrame_ExecuteJavaScript).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)), api.PascalStr(code), api.PascalStr(scriptUrl), uintptr(startLine))
}

// IsValid 该Frame是否有效
func (m *ICefFrame) IsValid() bool {
	var frameId = m.Id
	r1, _, _ := imports.Proc(internale_CEFFrame_IsValid).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
	return api.GoBool(r1)
}

// IsMain 是否为主Frame
func (m *ICefFrame) IsMain() bool {
	var frameId = m.Id
	r1, _, _ := imports.Proc(internale_CEFFrame_IsMain).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
	return api.GoBool(r1)
}

// IsFocused 是否已获取焦点
func (m *ICefFrame) IsFocused() bool {
	var frameId = m.Id
	r1, _, _ := imports.Proc(internale_CEFFrame_IsFocused).Call(uintptr(m.Browser.Identifier()), uintptr(unsafe.Pointer(&frameId)))
	return api.GoBool(r1)
}

// SendProcessMessageByIPC 发送进程消息
func (m *ICefFrame) SendProcessMessageByIPC(targetProcess CefProcessId, processMessage *ipc.ICefProcessMessage) ProcessMessageError {
	if processMessage == nil || processMessage.Name == "" || processMessage.ArgumentList == nil {
		return PMErr_REQUIRED_PARAMS_IS_NULL
	} else if ipc.InternalIPCNameCheck(processMessage.Name) {
		return PMErr_NAME_CANNOT_USED
	}
	data := processMessage.ArgumentList.Package()
	r1 := _CEFFrame_SendProcessMessageByIPC(m.Browser.Identifier(), m.Id, processMessage.Name, targetProcess, int32(processMessage.ArgumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return ProcessMessageError(r1)
}

func _CEFFrame_SendProcessMessageByIPC(browseId int32, frameId int64, name string, targetProcess CefProcessId, itemLength int32, data, dataLen uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFFrame_SendProcessMessageByIPC).Call(uintptr(browseId), uintptr(unsafe.Pointer(&frameId)), api.PascalStr(name), uintptr(targetProcess), uintptr(itemLength), data, dataLen)
	return r1
}

// ---- CEF Frame for ptr ----

//func (m *ICefFrame)  IsValid: Boolean;
//func (m *ICefFrame)  Undo;
//func (m *ICefFrame)  Redo;
//func (m *ICefFrame)  Cut;
//func (m *ICefFrame)  Copy;
//func (m *ICefFrame)  Paste;
//func (m *ICefFrame)  Del;
//func (m *ICefFrame)  SelectAll;
//func (m *ICefFrame)  ViewSource;
//func (m *ICefFrame)  GetSource(const visitor: ICefStringVisitor);
//func (m *ICefFrame)  GetSourceProc(const proc: TCefStringVisitorProc);
//func (m *ICefFrame)  GetText(const visitor: ICefStringVisitor);
//func (m *ICefFrame)  GetTextProc(const proc: TCefStringVisitorProc);
//func (m *ICefFrame)  LoadRequest(const request: ICefRequest);
//func (m *ICefFrame)  LoadUrl(const url: ustring);
//func (m *ICefFrame)  ExecuteJavaScript(const code, scriptUrl: ustring; startLine: Integer);
//func (m *ICefFrame)  IsMain: Boolean;
//func (m *ICefFrame)  IsFocused: Boolean;
//func (m *ICefFrame)  GetName: ustring;
//func (m *ICefFrame)  GetIdentifier: Int64;
//func (m *ICefFrame)  GetParent: ICefFrame;
//func (m *ICefFrame)  GetUrl: ustring;
//func (m *ICefFrame)  GetBrowser: ICefBrowser;
//func (m *ICefFrame)  GetV8Context: ICefv8Context;
//func (m *ICefFrame)  VisitDom(const visitor: ICefDomVisitor);
//func (m *ICefFrame)  VisitDomProc(const proc: TCefDomVisitorProc);
//func (m *ICefFrame)  CreateUrlRequest(const request: ICefRequest; const client: ICefUrlrequestClient): ICefUrlRequest;

func (m *ICefFrame) SendProcessMessage(targetProcess CefProcessId, message *ICefProcessMessage) {
	imports.Proc(internale_CEFFrame_SendProcessMessage).Call(m.Instance(), targetProcess.ToPtr(), message.Instance())
}
