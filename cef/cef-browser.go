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
package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

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

// GetBrowserId 获取浏览器ID 每一个窗口唯一ID
func (m *ICefBrowser) GetBrowserId() int32 {
	return m.Identifier()
}

// GetFrameId 获取FrameID 一个窗口中可以有多Frame
func (m *ICefBrowser) GetFrameId() int64 {
	return m.MainFrame().Identifier()
}

// Identifier ICefBrowser ID
func (m *ICefBrowser) Identifier() int32 {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetIdentifier).Call(m.Instance())
	return int32(r1)
}

// HostWindowHandle 获取窗口句柄, Browser窗口句柄
func (m *ICefBrowser) HostWindowHandle() types.HWND {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetHostWindowHandle).Call(m.Instance())
	return r1
}

// CloseBrowser 关闭浏览器，同时关闭窗口
func (m *ICefBrowser) CloseBrowser(forceClose bool) {
	imports.Proc(internale_CEFBrowser_CloseBrowser).Call(m.Instance(), api.PascalBool(forceClose))
}

// TryCloseBrowser 尝试关闭浏览器，同时尝试关闭窗口
func (m *ICefBrowser) TryCloseBrowser() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_TryCloseBrowser).Call(m.Instance())
	return api.GoBool(r1)
}

// SetFocus 设置焦点
func (m *ICefBrowser) SetFocus(focus bool) {
	imports.Proc(internale_CEFBrowser_SetFocus).Call(m.Instance(), api.PascalBool(focus))
}

// GetZoomLevel 获取缩放级别
func (m *ICefBrowser) GetZoomLevel() (result float64) {
	imports.Proc(internale_CEFBrowser_GetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetZoomLevel 设置缩放级别
func (m *ICefBrowser) SetZoomLevel(zoomLevel float64) {
	imports.Proc(internale_CEFBrowser_SetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&zoomLevel)))
}

// RunFileDialog 运行文件选择窗口, 可配合在下载文件事件中使用
func (m *ICefBrowser) RunFileDialog(mode int32, title, defaultFilePath string, acceptFilters *lcl.TStrings) {
	imports.Proc(internale_CEFBrowser_RunFileDialog).Call(m.Instance(), uintptr(mode), api.PascalStr(title), api.PascalStr(defaultFilePath), acceptFilters.Instance())
}

// StartDownload 开始下载
func (m *ICefBrowser) StartDownload(url string) {
	imports.Proc(internale_CEFBrowser_StartDownload).Call(m.Instance(), api.PascalStr(url))
}

// DownloadImage 开始下载图片
func (m *ICefBrowser) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	imports.Proc(internale_CEFBrowser_DownloadImage).Call(m.Instance(), api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache))
}

// Print
func (m *ICefBrowser) Print() {
	imports.Proc(internale_CEFBrowser_Print).Call(m.Instance())
}

// MainFrame 获取当前窗口的主Frame
func (m *ICefBrowser) MainFrame() *ICefFrame {
	if m == nil {
		return nil
	}
	if m.mainFrame == nil {
		var result uintptr
		imports.Proc(internale_CEFBrowser_GetMainFrame).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
		m.mainFrame = &ICefFrame{instance: unsafe.Pointer(result)}
	}
	return m.mainFrame
}

// GetFocusedFrame 获取当前窗口有焦点的Frame
func (m *ICefBrowser) GetFocusedFrame() *ICefFrame {
	if m == nil {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CEFBrowser_GetFocusedFrame).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefFrame{instance: unsafe.Pointer(result)}
}

// GetFrameById 根据FrameId获取Frame对象
func (m *ICefBrowser) GetFrameById(frameId int64) *ICefFrame {
	if m == nil {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CEFBrowser_GetFrameById).Call(m.Instance(), uintptr(unsafe.Pointer(&frameId)), uintptr(unsafe.Pointer(&result)))
	return &ICefFrame{instance: unsafe.Pointer(result)}
}

// GetFrameByName 根据FrameName获取Frame对象
func (m *ICefBrowser) GetFrameByName(frameName string) *ICefFrame {
	if m == nil {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CEFBrowser_GetFrameByName).Call(m.Instance(), api.PascalStr(frameName), uintptr(unsafe.Pointer(&result)))
	return &ICefFrame{instance: unsafe.Pointer(result)}
}

// PrintToPdf
func (m *ICefBrowser) PrintToPdf(path string, settings *CefPdfPrintSettings, callback *ICefPdfPrintCallback) {
	var settingsPtr uintptr = 0
	var setPtr *cefPdfPrintSettingsPtr
	if callback == nil {
		callback = PdfPrintCallbackRef.New()
	}
	if settings == nil {
		settings = &CefPdfPrintSettings{}
	}
	setPtr = settings.ToPtr()
	settingsPtr = uintptr(unsafe.Pointer(setPtr))
	imports.Proc(internale_CEFBrowser_PrintToPdf).Call(m.Instance(), api.PascalStr(path), settingsPtr, callback.Instance())
}

// ExecuteDevToolsMethod 执行开发者工具方法
func (m *ICefBrowser) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	if m == nil || dictionaryValue == nil {
		return
	}
	imports.Proc(internale_CEFBrowser_ExecuteDevToolsMethod).Call(m.Instance(), uintptr(messageId), api.PascalStr(method), dictionaryValue.Instance())
}

// SendKeyEvent 发送模拟键盘事件
func (m *ICefBrowser) SendKeyEvent(event *TCefKeyEvent) {
	imports.Proc(internale_CEFBrowser_SendKeyEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)))
}

// SetAudioMuted
func (m *ICefBrowser) SetAudioMuted(mute bool) {
	imports.Proc(internale_CEFBrowser_SetAudioMuted).Call(m.Instance(), api.PascalBool(mute))
}

// IsAudioMuted
func (m *ICefBrowser) IsAudioMuted() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_IsAudioMuted).Call(m.Instance())
	return api.GoBool(r1)
}

// SetAutoResizeEnabled 设置启用自动调整大小
func (m *ICefBrowser) SetAutoResizeEnabled(enabled bool, minSize, maxSize *TCefSize) {
	imports.Proc(internale_CEFBrowser_SetAutoResizeEnabled).Call(m.Instance(), api.PascalBool(enabled), uintptr(unsafe.Pointer(minSize)), uintptr(unsafe.Pointer(maxSize)))
}

// SetAccessibilityState 设置可访问性状态
func (m *ICefBrowser) SetAccessibilityState(accessibilityState TCefState) {
	imports.Proc(internale_CEFBrowser_SetAccessibilityState).Call(m.Instance(), uintptr(accessibilityState))
}

// NotifyMoveOrResizeStarted 通用移动或大小开始
func (m *ICefBrowser) NotifyMoveOrResizeStarted() {
	imports.Proc(internale_CEFBrowser_NotifyMoveOrResizeStarted).Call(m.Instance())
}

// NotifyScreenInfoChanged 通知屏幕改变
func (m *ICefBrowser) NotifyScreenInfoChanged() {
	imports.Proc(internale_CEFBrowser_NotifyScreenInfoChanged).Call(m.Instance())
}

// SendCaptureLostEvent 发送失去标题事件
func (m *ICefBrowser) SendCaptureLostEvent() {
	imports.Proc(internale_CEFBrowser_SendCaptureLostEvent).Call(m.Instance())
}

// SendTouchEvent 发送触摸事件
func (m *ICefBrowser) SendTouchEvent(event *TCefTouchEvent) {
	imports.Proc(internale_CEFBrowser_SendTouchEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)))
}

// SendMouseWheelEvent 发送鼠标滚轮事件
func (m *ICefBrowser) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	imports.Proc(internale_CEFBrowser_SendMouseWheelEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), uintptr(deltaX), uintptr(deltaY))
}

// SendMouseMoveEvent 发送鼠标移动事件
func (m *ICefBrowser) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	imports.Proc(internale_CEFBrowser_SendMouseMoveEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), api.PascalBool(mouseLeave))
}

// SendMouseClickEvent 发送鼠标点击事件
func (m *ICefBrowser) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	imports.Proc(internale_CEFBrowser_SendMouseClickEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), uintptr(type_), api.PascalBool(mouseUp), uintptr(clickCount))
}

// ViewSource 显示网页源码
func (m *ICefBrowser) ViewSource() {
	m.createBrowserViewSource()
}

// ShowDevTools 显示开发者工具
func (m *ICefBrowser) ShowDevTools() {
	if browserWinInfo := BrowserWindow.GetWindowInfo(m.Identifier()); browserWinInfo != nil {
		m.createBrowserDevTools(browserWinInfo)
	}
}

// CloseDevTools 关闭开发者工具
func (m *ICefBrowser) CloseDevTools() {
	imports.Proc(internale_CEFBrowser_CloseDevTools).Call(m.Instance())
}

// HasDevTools 判断启用开发者工具
func (m *ICefBrowser) HasDevTools() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_HasDevTools).Call(m.Instance())
	return api.GoBool(r1)
}

// CanGoBack 是否允许后退
func (m *ICefBrowser) CanGoBack() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_CanGoBack).Call(m.Instance())
	return api.GoBool(r1)
}

// GoBack 历史-后退
func (m *ICefBrowser) GoBack() {
	if m.CanGoBack() {
		imports.Proc(internale_CEFBrowser_GoBack).Call(m.Instance())
	}
}

// CanGoForward 是否允许前进
func (m *ICefBrowser) CanGoForward() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_CanGoForward).Call(m.Instance())
	return api.GoBool(r1)
}

// GoForward 历史-前进
func (m *ICefBrowser) GoForward() {
	if m.CanGoForward() {
		imports.Proc(internale_CEFBrowser_GoForward).Call(m.Instance())
	}
}

// IsLoading 是否正在加载
func (m *ICefBrowser) IsLoading() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_IsLoading).Call(m.Instance())
	return api.GoBool(r1)
}

// Reload 重新加载
func (m *ICefBrowser) Reload() {
	imports.Proc(internale_CEFBrowser_Reload).Call(m.Instance())
}

// ReloadIgnoreCache 重新加载忽略缓存
func (m *ICefBrowser) ReloadIgnoreCache() {
	imports.Proc(internale_CEFBrowser_ReloadIgnoreCache).Call(m.Instance())
}

// StopLoad 停止加载
func (m *ICefBrowser) StopLoad() {
	imports.Proc(internale_CEFBrowser_StopLoad).Call(m.Instance())
}

// FrameCount 获取当前窗口Frame数量
func (m *ICefBrowser) FrameCount() int {
	r1, _, _ := imports.Proc(internale_CEFBrowser_FrameCount).Call(m.Instance())
	return int(r1)
}

// GetFrameNames 获取当前窗口所有Frame名称
func (m *ICefBrowser) GetFrameNames() []*FrameNames {
	var result uintptr
	var resultSize int32
	imports.Proc(internale_CEFBrowser_GetFrameNames).Call(m.Instance(), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&resultSize)))
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

// Find 检索页面文本
func (m *ICefBrowser) Find(searchText string, forward, matchCase, findNext bool) {
	imports.Proc(internale_CEFBrowser_Find).Call(m.Instance(), api.PascalStr(searchText), api.PascalBool(forward), api.PascalBool(matchCase), api.PascalBool(findNext))
}

// StopFinding 停止加载
func (m *ICefBrowser) StopFinding(clearSelection bool) {
	imports.Proc(internale_CEFBrowser_StopFinding).Call(m.Instance(), api.PascalBool(clearSelection))
}

func (m *ICefBrowser) Free() {
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
	imports.Proc(internale_CEFBrowserRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
	return data
}
