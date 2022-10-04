//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

//type ICefBrowser
type ICefBrowser struct {
	browseId int32
	chromium uintptr
}

type frameNamesPtr struct {
	Name  uintptr
	Value uintptr
}

type FrameNames struct {
	Name  string
	Value string
}

func (m *ICefBrowser) Free() {
	m.browseId = 0
}

func (m *ICefBrowser) Instance() uintptr {
	return uintptr(m.browseId)
}

func (m *ICefBrowser) IsValid() bool {
	return uintptr(m.browseId) != 0
}

//得到浏览器ID号
func (m *ICefBrowser) Identifier() int32 {
	return m.browseId
}

//得到HostWindowHandle
func (m *ICefBrowser) HostWindowHandle() types.HWND {
	r1, _, _ := Proc("CEFBrowser_GetHostWindowHandle").Call(uintptr(m.browseId))
	return r1
}

//CloseBrowser
func (m *ICefBrowser) CloseBrowser(forceClose bool) {
	if m.IsValid() {
		Proc("CEFBrowser_CloseBrowser").Call(uintptr(m.browseId), api.GoBoolToDBool(forceClose))
	}
}

//TryCloseBrowser
func (m *ICefBrowser) TryCloseBrowser() bool {
	r1, _, _ := Proc("CEFBrowser_TryCloseBrowser").Call(uintptr(m.browseId))
	return api.DBoolToGoBool(r1)
}

//SetFocus
func (m *ICefBrowser) SetFocus(focus bool) {
	Proc("CEFBrowser_SetFocus").Call(uintptr(m.browseId), api.GoBoolToDBool(focus))
}

//GetZoomLevel
func (m *ICefBrowser) GetZoomLevel() float64 {
	var ret float64
	Proc("CEFBrowser_GetZoomLevel").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&ret)))
	return ret
}

//SetZoomLevel
func (m *ICefBrowser) SetZoomLevel(zoomLevel float64) {
	Proc("CEFBrowser_SetZoomLevel").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&zoomLevel)))
}

//RunFileDialog
func (m *ICefBrowser) RunFileDialog(mode int32, title, defaultFilePath string, acceptFilters *lcl.TStrings) {
	Proc("CEFBrowser_RunFileDialog").Call(uintptr(m.browseId), uintptr(mode), api.GoStrToDStr(title), api.GoStrToDStr(defaultFilePath), acceptFilters.Instance())
}

//StartDownload
func (m *ICefBrowser) StartDownload(url string) {
	Proc("CEFBrowser_StartDownload").Call(uintptr(m.browseId), api.GoStrToDStr(url))
}

//DownloadImage
func (m *ICefBrowser) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	Proc("CEFBrowser_DownloadImage").Call(uintptr(m.browseId), api.GoStrToDStr(imageUrl), api.GoBoolToDBool(isFavicon), uintptr(maxImageSize), api.GoBoolToDBool(bypassCache))
}

//Print
func (m *ICefBrowser) Print() {
	Proc("CEFBrowser_Print").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) GetFocusedFrame() *ICefFrame {
	r1, _, _ := Proc("CEFBrowser_GetFocusedFrame").Call(uintptr(m.browseId))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.DStrToGoStr(cf.Identifier))
		ret.Name = api.DStrToGoStr(cf.Name)
		ret.Url = api.DStrToGoStr(cf.Url)
		return ret
	}
	return nil
}
func (m *ICefBrowser) MainFrame() *ICefFrame {
	r1, _, _ := Proc("CEFBrowser_GetMainFrame").Call(uintptr(m.browseId))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.DStrToGoStr(cf.Identifier))
		ret.Name = api.DStrToGoStr(cf.Name)
		ret.Url = api.DStrToGoStr(cf.Url)
		return ret
	}
	return nil
}

func (m *ICefBrowser) GetFrameById(frameId int64) *ICefFrame {
	r1, _, _ := Proc("CEFBrowser_GetFrameById").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&frameId)))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.DStrToGoStr(cf.Identifier))
		ret.Name = api.DStrToGoStr(cf.Name)
		ret.Url = api.DStrToGoStr(cf.Url)
		cf = nil
		return ret
	}
	return nil
}

func (m *ICefBrowser) GetFrameByName(frameName string) *ICefFrame {
	r1, _, _ := Proc("CEFBrowser_GetFrameByName").Call(uintptr(m.browseId), api.GoStrToDStr(frameName))
	cf := (*cefFrame)(unsafe.Pointer(r1))
	if cf != nil {
		ret := &ICefFrame{}
		ret.Browser = m
		ret.Id = StrToInt64(api.DStrToGoStr(cf.Identifier))
		ret.Name = api.DStrToGoStr(cf.Name)
		ret.Url = api.DStrToGoStr(cf.Url)
		cf = nil
		return ret
	}
	return nil
}

//PrintToPdf
//func (m *ICefBrowser) PrintToPdf(path string) {
//	Proc("CEFBrowser_PrintToPdf").Call(uintptr(m.browseId), api.GoStrToDStr(path))
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
	Proc("CEFBrowser_ExecuteDevToolsMethod").Call(uintptr(m.browseId), uintptr(messageId), api.GoStrToDStr(method), uintptr(argsLen), uintptr(dataPtr), uintptr(dataLen))
}

//SendKeyEvent
func (m *ICefBrowser) SendKeyEvent(event *TCefKeyEvent) {
	Proc("CEFBrowser_SendKeyEvent").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)))
}

func (m *ICefBrowser) SetAudioMuted(mute bool) {
	Proc("CEFBrowser_SetAudioMuted").Call(uintptr(m.browseId), api.GoBoolToDBool(mute))
}

func (m *ICefBrowser) IsAudioMuted() bool {
	r1, _, _ := Proc("CEFBrowser_IsAudioMuted").Call(uintptr(m.browseId))
	return api.DBoolToGoBool(r1)
}

func (m *ICefBrowser) SetAutoResizeEnabled(enabled bool, minSize, maxSize *TCefSize) {
	Proc("CEFBrowser_SetAutoResizeEnabled").Call(uintptr(m.browseId), api.GoBoolToDBool(enabled), uintptr(unsafe.Pointer(minSize)), uintptr(unsafe.Pointer(maxSize)))
}

func (m *ICefBrowser) SetAccessibilityState(accessibilityState TCefState) {
	Proc("CEFBrowser_SetAccessibilityState").Call(uintptr(m.browseId), uintptr(accessibilityState))
}

func (m *ICefBrowser) NotifyMoveOrResizeStarted() {
	Proc("CEFBrowser_NotifyMoveOrResizeStarted").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) NotifyScreenInfoChanged() {
	Proc("CEFBrowser_NotifyScreenInfoChanged").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) SendCaptureLostEvent() {
	Proc("CEFBrowser_SendCaptureLostEvent").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) SendTouchEvent(event *TCefTouchEvent) {
	Proc("CEFBrowser_SendTouchEvent").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)))
}

func (m *ICefBrowser) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	Proc("CEFBrowser_SendMouseWheelEvent").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *ICefBrowser) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	Proc("CEFBrowser_SendMouseMoveEvent").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), api.GoBoolToDBool(mouseLeave))
}

func (m *ICefBrowser) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	Proc("CEFBrowser_SendMouseClickEvent").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(event)), uintptr(type_), api.GoBoolToDBool(mouseUp), uintptr(clickCount))
}

func (m *ICefBrowser) ViewSource(frame *ICefFrame) {
	createBrowserViewSource(m, frame)
}

//显示开发者工具
func (m *ICefBrowser) ShowDevTools() {
	if browserWinInfo := BrowserWindow.GetWindowInfo(m.Identifier()); browserWinInfo != nil {
		BrowserWindow.uiLock.Lock()
		defer BrowserWindow.uiLock.Unlock()
		createBrowserDevTools(m, browserWinInfo)
	}
}

//关闭开发者工具
func (m *ICefBrowser) CloseDevTools() {
	Proc("CEFBrowser_CloseDevTools").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) HasDevTools() bool {
	r1, _, _ := Proc("CEFBrowser_HasDevTools").Call(uintptr(m.browseId))
	return api.DBoolToGoBool(r1)
}

func (m *ICefBrowser) CanGoBack() bool {
	r1, _, _ := Proc("CEFBrowser_CanGoBack").Call(uintptr(m.browseId))
	return api.DBoolToGoBool(r1)
}

func (m *ICefBrowser) GoBack() {
	if m.CanGoBack() {
		Proc("CEFBrowser_GoBack").Call(uintptr(m.browseId))
	}
}

func (m *ICefBrowser) CanGoForward() bool {
	r1, _, _ := Proc("CEFBrowser_CanGoForward").Call(uintptr(m.browseId))
	return api.DBoolToGoBool(r1)
}

func (m *ICefBrowser) GoForward() {
	if m.CanGoForward() {
		Proc("CEFBrowser_GoForward").Call(uintptr(m.browseId))
	}
}

func (m *ICefBrowser) IsLoading() bool {
	r1, _, _ := Proc("CEFBrowser_IsLoading").Call(uintptr(m.browseId))
	return api.DBoolToGoBool(r1)
}

func (m *ICefBrowser) Reload() {
	Proc("CEFBrowser_Reload").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) ReloadIgnoreCache() {
	Proc("CEFBrowser_ReloadIgnoreCache").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) StopLoad() {
	Proc("CEFBrowser_StopLoad").Call(uintptr(m.browseId))
}

func (m *ICefBrowser) FrameCount() int {
	r1, _, _ := Proc("CEFBrowser_FrameCount").Call(uintptr(m.browseId))
	return int(r1)
}

func (m *ICefBrowser) GetFrameNames() []*FrameNames {
	var result uintptr
	var resultSize int32
	Proc("CEFBrowser_GetFrameNames").Call(uintptr(m.browseId), uintptr(unsafe.Pointer(&result)), uintptr(unsafe.Pointer(&resultSize)))
	frameNames := make([]*FrameNames, resultSize, resultSize)
	for i := 0; i < int(resultSize); i++ {
		fnsPtr := (*frameNamesPtr)(unsafe.Pointer(getParamOf(i, result)))
		frameNames[i] = &FrameNames{
			Name:  api.DStrToGoStr(fnsPtr.Name),
			Value: api.DStrToGoStr(fnsPtr.Value),
		}
	}
	result = 0
	resultSize = 0
	return frameNames
}

func (m *ICefBrowser) Find(searchText string, forward, matchCase, findNext bool) {
	Proc("CEFBrowser_Find").Call(uintptr(m.browseId), api.GoStrToDStr(searchText), api.GoBoolToDBool(forward), api.GoBoolToDBool(matchCase), api.GoBoolToDBool(findNext))
}

func (m *ICefBrowser) StopFinding(clearSelection bool) {
	Proc("CEFBrowser_StopFinding").Call(uintptr(m.browseId), api.GoBoolToDBool(clearSelection))
}
