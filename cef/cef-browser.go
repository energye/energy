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

// type ICefBrowser
type ICefBrowser struct {
	instance unsafe.Pointer
	chromium unsafe.Pointer
	browseId int32
}

type frameNamesPtr struct {
	Name  uintptr
	Value uintptr
}

type FrameNames struct {
	Name  string
	Value string
}

// 获取浏览器ID 每一个窗口独占唯一ID
func (m *ICefBrowser) GetBrowserId() int32 {
	return m.browseId
}

// 获取FrameID 一个窗口中可以有多Frame
func (m *ICefBrowser) GetFrameId() int64 {
	if mainFrame := m.MainFrame(); mainFrame != nil {
		return mainFrame.Id
	}
	return 0
}

// 浏览器ID号
func (m *ICefBrowser) Identifier() int32 {
	return m.browseId
}

// HostWindowHandle 获取窗口句柄, Browser窗口句柄
func (m *ICefBrowser) HostWindowHandle() types.HWND {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetHostWindowHandle).Call(uintptr(m.browseId))
	return r1
}

// CloseBrowser 关闭浏览器，同时关闭窗口
func (m *ICefBrowser) CloseBrowser(forceClose bool) {
	imports.Proc(internale_CEFBrowser_CloseBrowser).Call(uintptr(m.browseId), api.PascalBool(forceClose))
}

// TryCloseBrowser 尝试关闭浏览器，同时尝试关闭窗口
func (m *ICefBrowser) TryCloseBrowser() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_TryCloseBrowser).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

// SetFocus 设置焦点
func (m *ICefBrowser) SetFocus(focus bool) {
	imports.Proc(internale_CEFBrowser_SetFocus).Call(uintptr(m.browseId), api.PascalBool(focus))
}

// GetZoomLevel 获取缩放级别
func (m *ICefBrowser) GetZoomLevel() (result float64) {
	imports.Proc(internale_CEFBrowser_GetZoomLevel).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetZoomLevel 设置缩放级别
func (m *ICefBrowser) SetZoomLevel(zoomLevel float64) {
	imports.Proc(internale_CEFBrowser_SetZoomLevel).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&zoomLevel)))
}

// RunFileDialog 运行文件选择窗口, 可配合在下载文件事件中使用
func (m *ICefBrowser) RunFileDialog(mode int32, title, defaultFilePath string, acceptFilters *lcl.TStrings) {
	imports.Proc(internale_CEFBrowser_RunFileDialog).Call(uintptr(m.browseId), uintptr(mode), api.PascalStr(title), api.PascalStr(defaultFilePath), acceptFilters.Instance())
}

// StartDownload 开始下载
func (m *ICefBrowser) StartDownload(url string) {
	imports.Proc(internale_CEFBrowser_StartDownload).Call(uintptr(m.browseId), api.PascalStr(url))
}

// DownloadImage 开始下载图片
func (m *ICefBrowser) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	imports.Proc(internale_CEFBrowser_DownloadImage).Call(uintptr(m.browseId), api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache))
}

// Print
func (m *ICefBrowser) Print() {
	imports.Proc(internale_CEFBrowser_Print).Call(uintptr(m.browseId))
}

// GetFocusedFrame 获取当前窗口有焦点的Frame
func (m *ICefBrowser) GetFocusedFrame() *ICefFrame {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetFocusedFrame).Call(uintptr(m.browseId))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = common.StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		return ret
	}
	return nil
}

// MainFrame 获取当前窗口的主Frame
func (m *ICefBrowser) MainFrame() *ICefFrame {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetMainFrame).Call(uintptr(m.browseId))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = common.StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		return ret
	}
	return nil
}

// GetFrameById 根据FrameId获取Frame对象
func (m *ICefBrowser) GetFrameById(frameId int64) *ICefFrame {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetFrameById).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&frameId)))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = common.StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		cf = nil
		return ret
	}
	return nil
}

// GetFrameByName 根据FrameName获取Frame对象
func (m *ICefBrowser) GetFrameByName(frameName string) *ICefFrame {
	r1, _, _ := imports.Proc(internale_CEFBrowser_GetFrameByName).Call(uintptr(m.browseId), api.PascalStr(frameName))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = common.StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		cf = nil
		return ret
	}
	return nil
}

//PrintToPdf
//func (m *ICefBrowser) PrintToPdf(path string) {
//	imports.Proc(internale_CEFBrowser_PrintToPdf).Call(uintptr(m.browseId), api.PascalStr(path))
//}

// ExecuteDevToolsMethod 执行开发者工具方法
func (m *ICefBrowser) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *DictionaryValue) {
	var data = []byte{}
	var dataPtr unsafe.Pointer
	var dataLen int = 0
	var argsLen int = 0
	if dictionaryValue != nil && dictionaryValue.dataLen > 0 {
		defer func() {
			dictionaryValue.Clear()
			dictionaryValue = nil
			data = nil
			dataPtr = nil
		}()
		data = dictionaryValue.Package()
		argsLen = dictionaryValue.dataLen
		dataPtr = unsafe.Pointer(&data[0])
		dataLen = len(data) - 1
	} else {
		dataPtr = unsafe.Pointer(&data)
	}
	imports.Proc(internale_CEFBrowser_ExecuteDevToolsMethod).Call(uintptr(m.browseId), uintptr(messageId), api.PascalStr(method), uintptr(argsLen), uintptr(dataPtr), uintptr(dataLen))
}

// SendKeyEvent 发送模拟键盘事件
func (m *ICefBrowser) SendKeyEvent(event *TCefKeyEvent) {
	imports.Proc(internale_CEFBrowser_SendKeyEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)))
}

// SetAudioMuted
func (m *ICefBrowser) SetAudioMuted(mute bool) {
	imports.Proc(internale_CEFBrowser_SetAudioMuted).Call(uintptr(m.browseId), api.PascalBool(mute))
}

// IsAudioMuted
func (m *ICefBrowser) IsAudioMuted() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_IsAudioMuted).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

// SetAutoResizeEnabled 设置启用自动调整大小
func (m *ICefBrowser) SetAutoResizeEnabled(enabled bool, minSize, maxSize *TCefSize) {
	imports.Proc(internale_CEFBrowser_SetAutoResizeEnabled).Call(uintptr(m.browseId), api.PascalBool(enabled), uintptr(unsafe.Pointer(minSize)), uintptr(unsafe.Pointer(maxSize)))
}

// SetAccessibilityState 设置可访问性状态
func (m *ICefBrowser) SetAccessibilityState(accessibilityState TCefState) {
	imports.Proc(internale_CEFBrowser_SetAccessibilityState).Call(uintptr(m.browseId), uintptr(accessibilityState))
}

// NotifyMoveOrResizeStarted 通用移动或大小开始
func (m *ICefBrowser) NotifyMoveOrResizeStarted() {
	imports.Proc(internale_CEFBrowser_NotifyMoveOrResizeStarted).Call(uintptr(m.browseId))
}

// NotifyScreenInfoChanged 通知屏幕改变
func (m *ICefBrowser) NotifyScreenInfoChanged() {
	imports.Proc(internale_CEFBrowser_NotifyScreenInfoChanged).Call(uintptr(m.browseId))
}

// SendCaptureLostEvent 发送失去标题事件
func (m *ICefBrowser) SendCaptureLostEvent() {
	imports.Proc(internale_CEFBrowser_SendCaptureLostEvent).Call(uintptr(m.browseId))
}

// SendTouchEvent 发送触摸事件
func (m *ICefBrowser) SendTouchEvent(event *TCefTouchEvent) {
	imports.Proc(internale_CEFBrowser_SendTouchEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)))
}

// SendMouseWheelEvent 发送鼠标滚轮事件
func (m *ICefBrowser) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	imports.Proc(internale_CEFBrowser_SendMouseWheelEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), uintptr(deltaX), uintptr(deltaY))
}

// SendMouseMoveEvent 发送鼠标移动事件
func (m *ICefBrowser) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	imports.Proc(internale_CEFBrowser_SendMouseMoveEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), api.PascalBool(mouseLeave))
}

// SendMouseClickEvent 发送鼠标点击事件
func (m *ICefBrowser) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	imports.Proc(internale_CEFBrowser_SendMouseClickEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), uintptr(type_), api.PascalBool(mouseUp), uintptr(clickCount))
}

// ViewSource 显示网页源码
func (m *ICefBrowser) ViewSource(frame *ICefFrame) {
	m.createBrowserViewSource(frame)
}

// ShowDevTools 显示开发者工具
func (m *ICefBrowser) ShowDevTools() {
	if browserWinInfo := BrowserWindow.GetWindowInfo(m.Identifier()); browserWinInfo != nil {
		m.createBrowserDevTools(browserWinInfo)
	}
}

// CloseDevTools 关闭开发者工具
func (m *ICefBrowser) CloseDevTools() {
	imports.Proc(internale_CEFBrowser_CloseDevTools).Call(uintptr(m.browseId))
}

// HasDevTools 判断启用开发者工具
func (m *ICefBrowser) HasDevTools() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_HasDevTools).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

// CanGoBack 是否允许后退
func (m *ICefBrowser) CanGoBack() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_CanGoBack).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

// GoBack 历史-后退
func (m *ICefBrowser) GoBack() {
	if m.CanGoBack() {
		imports.Proc(internale_CEFBrowser_GoBack).Call(uintptr(m.browseId))
	}
}

// CanGoForward 是否允许前进
func (m *ICefBrowser) CanGoForward() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_CanGoForward).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

// GoForward 历史-前进
func (m *ICefBrowser) GoForward() {
	if m.CanGoForward() {
		imports.Proc(internale_CEFBrowser_GoForward).Call(uintptr(m.browseId))
	}
}

// IsLoading 是否正在加载
func (m *ICefBrowser) IsLoading() bool {
	r1, _, _ := imports.Proc(internale_CEFBrowser_IsLoading).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

// Reload 重新加载
func (m *ICefBrowser) Reload() {
	imports.Proc(internale_CEFBrowser_Reload).Call(uintptr(m.browseId))
}

// ReloadIgnoreCache 重新加载忽略缓存
func (m *ICefBrowser) ReloadIgnoreCache() {
	imports.Proc(internale_CEFBrowser_ReloadIgnoreCache).Call(uintptr(m.browseId))
}

// StopLoad 停止加载
func (m *ICefBrowser) StopLoad() {
	imports.Proc(internale_CEFBrowser_StopLoad).Call(uintptr(m.browseId))
}

// FrameCount 获取当前窗口Frame数量
func (m *ICefBrowser) FrameCount() int {
	r1, _, _ := imports.Proc(internale_CEFBrowser_FrameCount).Call(uintptr(m.browseId))
	return int(r1)
}

// GetFrameNames 获取当前窗口所有Frame名称
func (m *ICefBrowser) GetFrameNames() []*FrameNames {
	var result uintptr
	var resultSize int32
	imports.Proc(internale_CEFBrowser_GetFrameNames).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&resultSize)))
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
	imports.Proc(internale_CEFBrowser_Find).Call(uintptr(m.browseId), api.PascalStr(searchText), api.PascalBool(forward), api.PascalBool(matchCase), api.PascalBool(findNext))
}

// StopFinding 停止加载
func (m *ICefBrowser) StopFinding(clearSelection bool) {
	imports.Proc(internale_CEFBrowser_StopFinding).Call(uintptr(m.browseId), api.PascalBool(clearSelection))
}

// _CEFBrowser_ShowDevTools 开发者工具，内部调用
func _CEFBrowser_ShowDevTools(chromium, browser, windowParent, name uintptr) {
	imports.Proc(internale_CEFBrowser_ShowDevTools).Call(chromium, browser, windowParent, name)
}
