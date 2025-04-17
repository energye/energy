//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Browser 实例
//	直接在 browser 使用 browser host 功能函数

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/cef/ipc/argument"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// ICefBrowser
type ICefBrowser struct {
	base           TCefBaseRefCounted
	instance       unsafe.Pointer
	mainFrame      *ICefFrame
	requestContext *ICefRequestContext
	windowHandle   types.HWND
	idFrames       map[string]*ICefFrame
	nameFrames     map[string]*ICefFrame
}

type frameNamesPtr struct {
	Name  uintptr
	Value uintptr
}

type FrameNames struct {
	Name  string
	Value string
}

// Instance 实例
func (m *ICefBrowser) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefBrowser) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// BrowserId 获取浏览器ID 每一个窗口唯一ID
func (m *ICefBrowser) BrowserId() int32 {
	if !m.IsValid() {
		return 0
	}
	return m.Identifier()
}

// FrameId 获取FrameID 一个窗口中可以有多Frame
func (m *ICefBrowser) FrameId() string {
	if !m.IsValid() {
		return ""
	}
	return m.MainFrame().Identifier()
}

// Identifier ICefBrowser ID
func (m *ICefBrowser) Identifier() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_GetIdentifier).Call(m.Instance())
	return int32(r1)
}

// HostWindowHandle 获取窗口句柄, Browser窗口句柄
func (m *ICefBrowser) HostWindowHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	if m.windowHandle == 0 {
		m.windowHandle, _, _ = imports.Proc(def.CEFBrowser_GetHostWindowHandle).Call(m.Instance())
	}
	return m.windowHandle
}

// CloseBrowser 关闭浏览器，同时关闭窗口
func (m *ICefBrowser) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_CloseBrowser).Call(m.Instance(), api.PascalBool(forceClose))
}

// TryCloseBrowser 尝试关闭浏览器，同时尝试关闭窗口
func (m *ICefBrowser) TryCloseBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_TryCloseBrowser).Call(m.Instance())
	return api.GoBool(r1)
}

// SetFocus 设置焦点
func (m *ICefBrowser) SetFocus(focus bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SetFocus).Call(m.Instance(), api.PascalBool(focus))
}

// GetZoomLevel 获取缩放级别
func (m *ICefBrowser) GetZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_GetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetZoomLevel 设置缩放级别
func (m *ICefBrowser) SetZoomLevel(zoomLevel float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&zoomLevel)))
}

// RunFileDialog FileDialog
//
//	 打开文件、文件夹、多选文件、保存
//		在回调函数中获取最终选择结果
//		如果回调函数为 nil 不会弹出窗口
func (m *ICefBrowser) RunFileDialog(mode FileDialogMode, title, defaultFilePath string, acceptFilters lcl.IStrings, callback *ICefRunFileDialogCallback) {
	if !m.IsValid() {
		return
	}
	if acceptFilters == nil {
		acceptFilters = lcl.NewStringList()
		defer acceptFilters.Free()
	}
	imports.Proc(def.CEFBrowser_RunFileDialog).Call(m.Instance(), uintptr(mode), api.PascalStr(title), api.PascalStr(defaultFilePath), acceptFilters.Instance(), callback.Instance())
}

// StartDownload 开始下载
func (m *ICefBrowser) StartDownload(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_StartDownload).Call(m.Instance(), api.PascalStr(url))
}

// DownloadImage 开始下载图片
func (m *ICefBrowser) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool, callback *ICefDownloadImageCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_DownloadImage).Call(m.Instance(), api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache), callback.Instance())
}

// Print
func (m *ICefBrowser) Print() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_Print).Call(m.Instance())
}

// MainFrame 获取当前窗口的主Frame
func (m *ICefBrowser) MainFrame() *ICefFrame {
	if !m.IsValid() {
		return nil
	}
	if m.mainFrame == nil || m.mainFrame.instance == nil || !m.mainFrame.IsValid() {
		var result uintptr
		imports.Proc(def.CEFBrowser_GetMainFrame).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.mainFrame = &ICefFrame{instance: unsafe.Pointer(result)}
	}
	return m.mainFrame
}

// GetFocusedFrame 获取当前窗口有焦点的Frame
func (m *ICefBrowser) GetFocusedFrame() *ICefFrame {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFBrowser_GetFocusedFrame).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFrame{instance: unsafe.Pointer(result)}
	}
	return nil
}

// GetFrameById 根据FrameId获取Frame对象
func (m *ICefBrowser) GetFrameById(frameId string) *ICefFrame {
	if !m.IsValid() {
		return nil
	}
	if m.idFrames == nil {
		m.idFrames = make(map[string]*ICefFrame)
	}
	if frame, ok := m.idFrames[frameId]; ok {
		if frame.instance != nil && frame.IsValid() {
			return frame
		}
		delete(m.idFrames, frameId)
	}
	var result uintptr
	imports.Proc(def.CEFBrowser_GetFrameById).Call(m.Instance(), api.PascalStr(frameId), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		frame := &ICefFrame{instance: unsafe.Pointer(result)}
		m.idFrames[frameId] = frame
		return frame
	}
	return nil
}

// GetFrameByName 根据FrameName获取Frame对象
func (m *ICefBrowser) GetFrameByName(frameName string) *ICefFrame {
	if !m.IsValid() {
		return nil
	}
	if m.nameFrames == nil {
		m.nameFrames = make(map[string]*ICefFrame)
	}
	if frame, ok := m.nameFrames[frameName]; ok {
		if frame.instance != nil && frame.IsValid() {
			return frame
		}
		delete(m.nameFrames, frameName)
	}
	var result uintptr
	imports.Proc(def.CEFBrowser_GetFrameByName).Call(m.Instance(), api.PascalStr(frameName), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		frame := &ICefFrame{instance: unsafe.Pointer(result)}
		m.nameFrames[frameName] = frame
		return frame
	}
	return nil
}

// PrintToPdf
func (m *ICefBrowser) PrintToPdf(path string, settings TCefPdfPrintSettings, callback *ICefPdfPrintCallback) {
	if !m.IsValid() {
		return
	}
	settingsPtr := settings.ToPtr()
	imports.Proc(def.CEFBrowser_PrintToPdf).Call(m.Instance(), api.PascalStr(path), uintptr(unsafe.Pointer(settingsPtr)), callback.Instance())
}

// SendDevToolsMessage 发送开发者工具消息
func (m *ICefBrowser) SendDevToolsMessage(message string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendDevToolsMessage).Call(m.Instance(), api.PascalStr(message))
}

// ExecuteDevToolsMethod 执行开发者工具方法
func (m *ICefBrowser) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	if !m.IsValid() || dictionaryValue == nil {
		return
	}
	imports.Proc(def.CEFBrowser_ExecuteDevToolsMethod).Call(m.Instance(), uintptr(messageId), api.PascalStr(method), dictionaryValue.Instance())
}

// SendKeyEvent 发送模拟键盘事件
func (m *ICefBrowser) SendKeyEvent(event *TCefKeyEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendKeyEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)))
}

// SetAudioMuted
func (m *ICefBrowser) SetAudioMuted(mute bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SetAudioMuted).Call(m.Instance(), api.PascalBool(mute))
}

// IsAudioMuted
func (m *ICefBrowser) IsAudioMuted() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_IsAudioMuted).Call(m.Instance())
	return api.GoBool(r1)
}

// SetAutoResizeEnabled 设置启用自动调整大小
func (m *ICefBrowser) SetAutoResizeEnabled(enabled bool, minSize, maxSize TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SetAutoResizeEnabled).Call(m.Instance(), api.PascalBool(enabled), uintptr(unsafe.Pointer(&minSize)), uintptr(unsafe.Pointer(&maxSize)))
}

// SetAccessibilityState 设置可访问性状态
func (m *ICefBrowser) SetAccessibilityState(accessibilityState TCefState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SetAccessibilityState).Call(m.Instance(), uintptr(accessibilityState))
}

// NotifyMoveOrResizeStarted 通用移动或大小开始
func (m *ICefBrowser) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_NotifyMoveOrResizeStarted).Call(m.Instance())
}

// NotifyScreenInfoChanged 通知屏幕改变
func (m *ICefBrowser) NotifyScreenInfoChanged() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_NotifyScreenInfoChanged).Call(m.Instance())
}

// SendCaptureLostEvent 发送失去标题事件
func (m *ICefBrowser) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendCaptureLostEvent).Call(m.Instance())
}

// SendTouchEvent 发送触摸事件
func (m *ICefBrowser) SendTouchEvent(event *TCefTouchEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendTouchEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)))
}

// SendMouseWheelEvent 发送鼠标滚轮事件
func (m *ICefBrowser) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendMouseWheelEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), uintptr(deltaX), uintptr(deltaY))
}

// SendMouseMoveEvent 发送鼠标移动事件
func (m *ICefBrowser) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendMouseMoveEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), api.PascalBool(mouseLeave))
}

// SendMouseClickEvent 发送鼠标点击事件
func (m *ICefBrowser) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_SendMouseClickEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), uintptr(type_), api.PascalBool(mouseUp), uintptr(clickCount))
}

// ViewSource 显示网页源码
func (m *ICefBrowser) ViewSource(currentWindow IBrowserWindow) {
	if !m.IsValid() {
		return
	}
	m.createBrowserViewSource(currentWindow)
}

// ShowDevTools 显示开发者工具
func (m *ICefBrowser) ShowDevTools(currentWindow IBrowserWindow, currentChromium ICEFChromiumBrowser) {
	if !m.IsValid() {
		return
	}
	if currentWindow != nil {
		m.createBrowserDevTools(currentWindow, currentChromium)
	}
}

// CloseDevTools 关闭开发者工具
func (m *ICefBrowser) CloseDevTools() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_CloseDevTools).Call(m.Instance())
}

// HasDevTools
func (m *ICefBrowser) HasView() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_HasView).Call(m.Instance())
	return api.GoBool(r1)
}

// HasDevTools
func (m *ICefBrowser) HasDevTools() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_HasDevTools).Call(m.Instance())
	return api.GoBool(r1)
}

// CanGoBack 是否允许后退
func (m *ICefBrowser) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_CanGoBack).Call(m.Instance())
	return api.GoBool(r1)
}

// GoBack 历史-后退
func (m *ICefBrowser) GoBack() {
	if !m.IsValid() {
		return
	}
	if m.CanGoBack() {
		imports.Proc(def.CEFBrowser_GoBack).Call(m.Instance())
	}
}

// CanGoForward 是否允许前进
func (m *ICefBrowser) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_CanGoForward).Call(m.Instance())
	return api.GoBool(r1)
}

// GoForward 历史-前进
func (m *ICefBrowser) GoForward() {
	if !m.IsValid() {
		return
	}
	if m.CanGoForward() {
		imports.Proc(def.CEFBrowser_GoForward).Call(m.Instance())
	}
}

// IsLoading 是否正在加载
func (m *ICefBrowser) IsLoading() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_IsLoading).Call(m.Instance())
	return api.GoBool(r1)
}

// Reload 重新加载
func (m *ICefBrowser) Reload() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_Reload).Call(m.Instance())
}

// ReloadIgnoreCache 重新加载忽略缓存
func (m *ICefBrowser) ReloadIgnoreCache() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_ReloadIgnoreCache).Call(m.Instance())
}

// StopLoad 停止加载
func (m *ICefBrowser) StopLoad() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_StopLoad).Call(m.Instance())
}

// FrameCount 获取当前窗口Frame数量
func (m *ICefBrowser) FrameCount() int {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_FrameCount).Call(m.Instance())
	return int(r1)
}

// GetFrameNames 获取当前窗口所有Frame名称
func (m *ICefBrowser) GetFrameNames() []*FrameNames {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	var resultSize int32
	imports.Proc(def.CEFBrowser_GetFrameNames).Call(m.Instance(), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&resultSize)))
	frameNames := make([]*FrameNames, resultSize, resultSize)
	for i := 0; i < int(resultSize); i++ {
		fnsPtr := (*frameNamesPtr)(unsafe.Pointer(common.GetParamOf(i, result)))
		frameNames[i] = &FrameNames{
			Name:  api.GoStr(fnsPtr.Name),
			Value: api.GoStr(fnsPtr.Value),
		}
	}
	result = 0
	resultSize = 0
	return frameNames
}

func (m *ICefBrowser) GetFrameIdentifiers() *lcl.TStrings {
	var result uintptr
	imports.Proc(def.CEFBrowser_GetFrameIdentifiers).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return lcl.AsStrings(result)
}

// Find 检索页面文本
func (m *ICefBrowser) Find(searchText string, forward, matchCase, findNext bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_Find).Call(m.Instance(), api.PascalStr(searchText), api.PascalBool(forward), api.PascalBool(matchCase), api.PascalBool(findNext))
}

// StopFinding 停止加载
func (m *ICefBrowser) StopFinding(clearSelection bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_StopFinding).Call(m.Instance(), api.PascalBool(clearSelection))
}

// GetRequestContext -> ICefBrowserHost -> ICefRequestContext
func (m *ICefBrowser) GetRequestContext() *ICefRequestContext {
	if !m.IsValid() {
		return nil
	}
	if m.requestContext != nil && m.requestContext.IsValid() {
		return m.requestContext
	}
	var result uintptr
	imports.Proc(def.CEFBrowser_GetRequestContext).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		m.requestContext = &ICefRequestContext{instance: unsafe.Pointer(result)}
		return m.requestContext
	}
	return nil
}

// SendProcessMessage 发送进程消息
//
//	仅支持 CEF49
func (m *ICefBrowser) SendProcessMessage(targetProcess CefProcessId, message *ICefProcessMessage) {
	if application.Is49() {
		if !m.IsValid() {
			return
		}
		imports.Proc(def.CEFBrowser_SendProcessMessage).Call(m.Instance(), targetProcess.ToPtr(), message.Instance())
		message.Free()
	}
}

// SendProcessMessageForJSONBytes 发送进程消息
//
//	仅支持 CEF49
func (m *ICefBrowser) SendProcessMessageForJSONBytes(messageName string, targetProcess CefProcessId, data []byte) {
	if application.Is49() {
		if !m.IsValid() {
			return
		}
		imports.Proc(def.CEFBrowser_SendProcessMessageForJSONBytes).Call(m.Instance(), api.PascalStr(messageName), targetProcess.ToPtr(), uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))))
	}
}

// SendProcessMessageForV8Value 发送进程消息
//
//	仅支持 CEF49
func (m *ICefBrowser) SendProcessMessageForV8Value(messageName string, targetProcess CefProcessId, arguments *ICefV8Value) {
	if application.Is49() {
		if !m.IsValid() {
			return
		}
		imports.Proc(def.CEFBrowser_SendProcessMessageForV8Value).Call(m.Instance(), api.PascalStr(messageName), targetProcess.ToPtr(), arguments.Instance())
	}
}

// EmitRender IPC 发送进程 消息
//
// messageId != 0 是带有回调函数消息
//
//	仅支持 CEF49
func (m *ICefBrowser) EmitRender(messageId int32, eventName string, target target.ITarget, data ...interface{}) bool {
	if !application.Is49() {
		return false
	}
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
	m.SendProcessMessageForJSONBytes(internalIPCGoEmit, PID_RENDER, message.Bytes())
	message.Reset()
	return true
}

func (m *ICefBrowser) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_IsFullscreen).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefBrowser) ExitFullscreen(willCauseResize bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_ExitFullscreen).Call(m.Instance(), api.PascalBool(willCauseResize))
}

func (m *ICefBrowser) CanExecuteChromeCommand(commandId int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_CanExecuteChromeCommand).Call(m.Instance(), uintptr(commandId))
}

func (m *ICefBrowser) ExecuteChromeCommand(commandId int32, disposition TCefWindowOpenDisposition) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowser_ExecuteChromeCommand).Call(m.Instance(), uintptr(commandId), uintptr(disposition))
}

func (m *ICefBrowser) IsRenderProcessUnresponsive() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_IsRenderProcessUnresponsive).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefBrowser) GetRuntimeStyle() TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFBrowser_GetRuntimeStyle).Call(m.Instance())
	return TCefRuntimeStyle(r1)
}

func (m *ICefBrowser) Free() {
	if !m.IsValid() {
		return
	}
	if m.idFrames != nil {
		for _, frame := range m.idFrames {
			frame.Free()
		}
		m.idFrames = nil
	}
	if m.nameFrames != nil {
		for _, frame := range m.nameFrames {
			frame.Free()
		}
		m.nameFrames = nil
	}
	if m.requestContext != nil {
		m.requestContext.Free()
		m.requestContext = nil
	}
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// BrowserRef -> ICefBrowser
var BrowserRef browser

// browser
type browser uintptr

func (*browser) UnWrap(data *ICefBrowser) *ICefBrowser {
	var result uintptr
	imports.Proc(def.CEFBrowserRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	//data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}
