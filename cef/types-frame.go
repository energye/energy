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
	"github.com/energye/energy/cef/ipc/argument"
	"github.com/energye/energy/cef/ipc/target"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/json"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Instance 实例
func (m *ICefFrame) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// IsValid 该Frame是否有效
func (m *ICefFrame) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CEFFrame_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// Undo 撤销操作
func (m *ICefFrame) Undo() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Undo).Call(m.Instance())
}

// Redo 恢复
func (m *ICefFrame) Redo() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Redo).Call(m.Instance())
}

// Cut 剪切
func (m *ICefFrame) Cut() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Cut).Call(m.Instance())
}

// Copy 复制
func (m *ICefFrame) Copy() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Copy).Call(m.Instance())
}

// Paste 粘贴
func (m *ICefFrame) Paste() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Paste).Call(m.Instance())
}

// Del 删除
func (m *ICefFrame) Del() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Del).Call(m.Instance())
}

// SelectAll 选择所有
func (m *ICefFrame) SelectAll() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_SelectAll).Call(m.Instance())
}

// ViewSource 显示源码
func (m *ICefFrame) ViewSource() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_ViewSource).Call(m.Instance())
}

// LoadUrl 加载URL
func (m *ICefFrame) LoadUrl(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_LoadUrl).Call(m.Instance(), api.PascalStr(url))
}

// ExecuteJavaScript 执行JS
func (m *ICefFrame) ExecuteJavaScript(code, scriptUrl string, startLine int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_ExecuteJavaScript).Call(m.Instance(), api.PascalStr(code), api.PascalStr(scriptUrl), uintptr(startLine))
}

// IsMain 是否为主Frame
func (m *ICefFrame) IsMain() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CEFFrame_IsMain).Call(m.Instance())
	return api.GoBool(r1)
}

// IsFocused 是否已获取焦点
func (m *ICefFrame) IsFocused() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CEFFrame_IsFocused).Call(m.Instance())
	return api.GoBool(r1)
}

// SendProcessMessage 发送进程消息
func (m *ICefFrame) SendProcessMessage(targetProcess CefProcessId, message *ICefProcessMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_SendProcessMessage).Call(m.Instance(), targetProcess.ToPtr(), message.Instance())
	message.Free()
}

// SendProcessMessageForJSONBytes 发送进程消息
func (m *ICefFrame) SendProcessMessageForJSONBytes(messageName string, targetProcess CefProcessId, data []byte) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_SendProcessMessageForJSONBytes).Call(m.Instance(), api.PascalStr(messageName), targetProcess.ToPtr(), uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))))
}

// SendProcessMessageForV8Value 发送进程消息
func (m *ICefFrame) SendProcessMessageForV8Value(messageName string, targetProcess CefProcessId, arguments *ICefV8Value) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_SendProcessMessageForV8Value).Call(m.Instance(), api.PascalStr(messageName), targetProcess.ToPtr(), arguments.Instance())
}

// EmitRender IPC 发送进程 消息
//
// messageId != 0 是带有回调函数消息
func (m *ICefFrame) EmitRender(messageId int32, eventName string, target target.ITarget, data ...any) bool {
	if !m.IsValid() {
		return false
	}
	message := &argument.List{Id: messageId, EventName: eventName}
	if len(data) > 0 {
		argumentJSONArray := json.NewJSONArray(nil)
		for _, result := range data {
			switch result.(type) {
			case error:
				argumentJSONArray.Add(result.(error).Error())
			default:
				argumentJSONArray.Add(result)
			}
		}
		message.Data = argumentJSONArray.Data()
	}
	m.SendProcessMessageForJSONBytes(internalIPCGoExecuteJSEvent, PID_RENDER, message.Bytes())
	message.Reset()
	return true
}

func (m *ICefFrame) LoadRequest(request *ICefRequest) {
	if !m.IsValid() || request == nil {
		return
	}
	imports.Proc(internale_CEFFrame_LoadRequest).Call(m.Instance(), request.Instance())
}

func (m *ICefFrame) Browser() *ICefBrowser {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CEFFrame_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBrowser{instance: unsafe.Pointer(result)}
}

func (m *ICefFrame) V8Context() *ICefV8Context {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CEFFrame_GetV8Context).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Context{instance: unsafe.Pointer(result)}
}

func (m *ICefFrame) VisitDom(visitor *ICefDomVisitor) {
	if !m.IsValid() || !visitor.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_VisitDom).Call(m.Instance(), visitor.Instance())
}

func (m *ICefFrame) Identifier() (result int64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CEFFrame_Identifier).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefFrame) Name() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CEFFrame_Name).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefFrame) Url() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CEFFrame_Url).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefFrame) Parent() *ICefFrame {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CEFFrame_Parent).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefFrame{instance: unsafe.Pointer(result)}
}

func (m *ICefFrame) Free() {
	if !m.IsValid() {
		return
	}
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
