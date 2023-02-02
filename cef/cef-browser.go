//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

//type ICefBrowser
type ICefBrowser struct {
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

func (m *ICefBrowser) GetBrowserId() int32 {
	return m.browseId
}

func (m *ICefBrowser) GetFrameId() int64 {
	if mainFrame := m.MainFrame(); mainFrame != nil {
		return mainFrame.Id
	}
	return 0
}

//浏览器ID号
func (m *ICefBrowser) Identifier() int32 {
	return m.browseId
}

//HostWindowHandle
func (m *ICefBrowser) HostWindowHandle() types.HWND {
	r1, _, _ := Proc(internale_CEFBrowser_GetHostWindowHandle).Call(uintptr(m.browseId))
	return r1
}

//CloseBrowser
func (m *ICefBrowser) CloseBrowser(forceClose bool) {
	Proc(internale_CEFBrowser_CloseBrowser).Call(uintptr(m.browseId), api.PascalBool(forceClose))
}

//TryCloseBrowser
func (m *ICefBrowser) TryCloseBrowser() bool {
	r1, _, _ := Proc(internale_CEFBrowser_TryCloseBrowser).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

//SetFocus
func (m *ICefBrowser) SetFocus(focus bool) {
	Proc(internale_CEFBrowser_SetFocus).Call(uintptr(m.browseId), api.PascalBool(focus))
}

//GetZoomLevel
func (m *ICefBrowser) GetZoomLevel() (result float64) {
	Proc(internale_CEFBrowser_GetZoomLevel).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&result)))
	return result
}

//SetZoomLevel
func (m *ICefBrowser) SetZoomLevel(zoomLevel float64) {
	Proc(internale_CEFBrowser_SetZoomLevel).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&zoomLevel)))
}

//RunFileDialog
func (m *ICefBrowser) RunFileDialog(mode int32, title, defaultFilePath string, acceptFilters *lcl.TStrings) {
	Proc(internale_CEFBrowser_RunFileDialog).Call(uintptr(m.browseId), uintptr(mode), api.PascalStr(title), api.PascalStr(defaultFilePath), acceptFilters.Instance())
}

//StartDownload
func (m *ICefBrowser) StartDownload(url string) {
	Proc(internale_CEFBrowser_StartDownload).Call(uintptr(m.browseId), api.PascalStr(url))
}

//DownloadImage
func (m *ICefBrowser) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	Proc(internale_CEFBrowser_DownloadImage).Call(uintptr(m.browseId), api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache))
}

//Print
func (m *ICefBrowser) Print() {
	Proc(internale_CEFBrowser_Print).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) GetFocusedFrame() *ICefFrame {
	r1, _, _ := Proc(internale_CEFBrowser_GetFocusedFrame).Call(uintptr(m.browseId))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		return ret
	}
	return nil
}
func (m *ICefBrowser) MainFrame() *ICefFrame {
	r1, _, _ := Proc(internale_CEFBrowser_GetMainFrame).Call(uintptr(m.browseId))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		return ret
	}
	return nil
}

func (m *ICefBrowser) GetFrameById(frameId int64) *ICefFrame {
	r1, _, _ := Proc(internale_CEFBrowser_GetFrameById).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&frameId)))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		cf = nil
		return ret
	}
	return nil
}

func (m *ICefBrowser) GetFrameByName(frameName string) *ICefFrame {
	r1, _, _ := Proc(internale_CEFBrowser_GetFrameByName).Call(uintptr(m.browseId), api.PascalStr(frameName))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.GoStr(cf.Identifier))
		ret.Name = api.GoStr(cf.Name)
		ret.Url = api.GoStr(cf.Url)
		cf = nil
		return ret
	}
	return nil
}

//PrintToPdf
//func (m *ICefBrowser) PrintToPdf(path string) {
//	Proc(internale_CEFBrowser_PrintToPdf).Call(uintptr(m.browseId), api.PascalStr(path))
//}

//ExecuteDevToolsMethod
func (m *ICefBrowser) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
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
	Proc(internale_CEFBrowser_ExecuteDevToolsMethod).Call(uintptr(m.browseId), uintptr(messageId), api.PascalStr(method), uintptr(argsLen), uintptr(dataPtr), uintptr(dataLen))
}

//SendKeyEvent
func (m *ICefBrowser) SendKeyEvent(event *TCefKeyEvent) {
	Proc(internale_CEFBrowser_SendKeyEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)))
}

func (m *ICefBrowser) SetAudioMuted(mute bool) {
	Proc(internale_CEFBrowser_SetAudioMuted).Call(uintptr(m.browseId), api.PascalBool(mute))
}

func (m *ICefBrowser) IsAudioMuted() bool {
	r1, _, _ := Proc(internale_CEFBrowser_IsAudioMuted).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

func (m *ICefBrowser) SetAutoResizeEnabled(enabled bool, minSize, maxSize *TCefSize) {
	Proc(internale_CEFBrowser_SetAutoResizeEnabled).Call(uintptr(m.browseId), api.PascalBool(enabled), uintptr(unsafe.Pointer(minSize)), uintptr(unsafe.Pointer(maxSize)))
}

func (m *ICefBrowser) SetAccessibilityState(accessibilityState TCefState) {
	Proc(internale_CEFBrowser_SetAccessibilityState).Call(uintptr(m.browseId), uintptr(accessibilityState))
}

func (m *ICefBrowser) NotifyMoveOrResizeStarted() {
	Proc(internale_CEFBrowser_NotifyMoveOrResizeStarted).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) NotifyScreenInfoChanged() {
	Proc(internale_CEFBrowser_NotifyScreenInfoChanged).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) SendCaptureLostEvent() {
	Proc(internale_CEFBrowser_SendCaptureLostEvent).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) SendTouchEvent(event *TCefTouchEvent) {
	Proc(internale_CEFBrowser_SendTouchEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)))
}

func (m *ICefBrowser) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	Proc(internale_CEFBrowser_SendMouseWheelEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *ICefBrowser) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	Proc(internale_CEFBrowser_SendMouseMoveEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), api.PascalBool(mouseLeave))
}

func (m *ICefBrowser) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	Proc(internale_CEFBrowser_SendMouseClickEvent).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), uintptr(type_), api.PascalBool(mouseUp), uintptr(clickCount))
}

func (m *ICefBrowser) ViewSource(frame *ICefFrame) {
	m.createBrowserViewSource(frame)
}

//显示开发者工具
func (m *ICefBrowser) ShowDevTools() {
	if browserWinInfo := BrowserWindow.GetWindowInfo(m.Identifier()); browserWinInfo != nil {
		m.createBrowserDevTools(browserWinInfo)
	}
}

//关闭开发者工具
func (m *ICefBrowser) CloseDevTools() {
	Proc(internale_CEFBrowser_CloseDevTools).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) HasDevTools() bool {
	r1, _, _ := Proc(internale_CEFBrowser_HasDevTools).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

func (m *ICefBrowser) CanGoBack() bool {
	r1, _, _ := Proc(internale_CEFBrowser_CanGoBack).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

func (m *ICefBrowser) GoBack() {
	if m.CanGoBack() {
		Proc(internale_CEFBrowser_GoBack).Call(uintptr(m.browseId))
	}
}

func (m *ICefBrowser) CanGoForward() bool {
	r1, _, _ := Proc(internale_CEFBrowser_CanGoForward).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

func (m *ICefBrowser) GoForward() {
	if m.CanGoForward() {
		Proc(internale_CEFBrowser_GoForward).Call(uintptr(m.browseId))
	}
}

func (m *ICefBrowser) IsLoading() bool {
	r1, _, _ := Proc(internale_CEFBrowser_IsLoading).Call(uintptr(m.browseId))
	return api.GoBool(r1)
}

func (m *ICefBrowser) Reload() {
	Proc(internale_CEFBrowser_Reload).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) ReloadIgnoreCache() {
	Proc(internale_CEFBrowser_ReloadIgnoreCache).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) StopLoad() {
	Proc(internale_CEFBrowser_StopLoad).Call(uintptr(m.browseId))
}

func (m *ICefBrowser) FrameCount() int {
	r1, _, _ := Proc(internale_CEFBrowser_FrameCount).Call(uintptr(m.browseId))
	return int(r1)
}

func (m *ICefBrowser) GetFrameNames() []*FrameNames {
	var result uintptr
	var resultSize int32
	Proc(internale_CEFBrowser_GetFrameNames).Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&resultSize)))
	frameNames := make([]*FrameNames, resultSize, resultSize)
	for i := 0; i < int(resultSize); i++ {
		fnsPtr := (*frameNamesPtr)(unsafe.Pointer(GetParamOf(i, result)))
		frameNames[i] = &FrameNames{
			Name:  api.GoStr(fnsPtr.Name),
			Value: api.GoStr(fnsPtr.Value),
		}
	}
	result = 0
	resultSize = 0
	return frameNames
}

func (m *ICefBrowser) Find(searchText string, forward, matchCase, findNext bool) {
	Proc(internale_CEFBrowser_Find).Call(uintptr(m.browseId), api.PascalStr(searchText), api.PascalBool(forward), api.PascalBool(matchCase), api.PascalBool(findNext))
}

func (m *ICefBrowser) StopFinding(clearSelection bool) {
	Proc(internale_CEFBrowser_StopFinding).Call(uintptr(m.browseId), api.PascalBool(clearSelection))
}

// ICefBrowser _CEFBrowser_ShowDevTools
func _CEFBrowser_ShowDevTools(chromium, browser, windowParent, name uintptr) {
	Proc(internale_CEFBrowser_ShowDevTools).Call(chromium, browser, windowParent, name)
}
