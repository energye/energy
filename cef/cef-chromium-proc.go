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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"strings"
	"time"
	"unsafe"
)

type IChromiumProc interface {
	lcl.IObject
	On(name string, eventCallback ipc.EventCallback)
	ExecuteJavaScript(code, scriptURL string, startLine int32) //单纯的执行js，没有返回值
	Emit(eventName string, args ipc.IArgumentList, target *GoEmitTarget) ProcessMessageError
	EmitAndCallback(eventName string, args ipc.IArgumentList, target *GoEmitTarget, callback ipc.IPCCallback) ProcessMessageError
	EmitAndReturn(eventName string, args ipc.IArgumentList, target *GoEmitTarget) (ipc.IIPCContext, ProcessMessageError)
	SetDefaultURL(defaultURL string)
	SetEnableMultiBrowserMode(enableMultiBrowserMode bool)
	LoadUrl(url string)
	LoadHtml(html string)
	StartDownload(url string)
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool)
	Reload()
	StopLoad()
	ResetZoomLevel()
	CloseAllBrowsers()
	CreateBrowser(window ITCefWindow) bool
	Initialized() bool
	BrowserId() int32
	IsSameBrowser(browser *ICefBrowser) bool
	PrintToPDF(saveFilePath string)
	Print()
	BrowserDownloadCancel(browseId, downloadId int32)
	BrowserDownloadPause(browseId, downloadId int32)
	BrowserZoom(zoom ZOOM)
	GoBack()
	GoForward()
	NotifyMoveOrResizeStarted()
	CloseBrowser(forceClose bool)
	ShowDevTools(window ITCefWindow)
	CloseDevTools(window ITCefWindow)
	VisitAllCookies(id int32)
	VisitURLCookies(url string, includeHttpOnly bool, id int32)
	DeleteCookies(url, cookieName string, deleteImmediately bool)
	SetCookie(url, name, value, domain, path string,
		secure, httponly, hasExpires bool,
		creation, lastAccess, expires time.Time,
		sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32)
	SetProxy(cefProxy *TCefProxy)
	UpdatePreferences()
	ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue)
	SendProcessMessage(targetProcess CefProcessId, processMessage *ipc.ICefProcessMessage) int
	CreateClientHandler(client *ICefClient, alsOSR bool) bool
	SetFocus(value bool)
	SendCaptureLostEvent()
	FrameIsFocused() bool
}

func (m *TCEFChromium) SetDefaultURL(defaultURL string) {
	if IsLinux() || IsDarwin() {
		httpIdx := strings.Index(defaultURL, "http")
		if httpIdx != 0 {
			if strings.Index(defaultURL, "file://") != 0 {
				defaultURL = "file://" + defaultURL
			}
		}
	}
	_CEFChromium_SetDefaultURL(m.instance, defaultURL)
}

func (m *TCEFChromium) SetEnableMultiBrowserMode(enableMultiBrowserMode bool) {
	_CEFChromium_SetMultiBrowserMode(m.instance, enableMultiBrowserMode)
}

func (m *TCEFChromium) LoadUrl(url string) {
	_CEFChromium_LoadURL(m.instance, url)
}

func (m *TCEFChromium) LoadHtml(html string) {
	_CEFChromium_LoadString(m.instance, html)
}

func (m *TCEFChromium) StartDownload(url string) {
	_CEFChromium_StartDownload(m.instance, url)
}

func (m *TCEFChromium) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	_CEFChromium_DownloadImage(m.instance, imageUrl, isFavicon, maxImageSize, bypassCache)
}

func (m *TCEFChromium) Reload() {
	_CEFChromium_Reload(m.instance)
}

func (m *TCEFChromium) StopLoad() {
	_CEFChromium_StopLoad(m.instance)
}

func (m *TCEFChromium) ResetZoomLevel() {
	_CEFChromium_ResetZoomLevel(m.instance)
}

func (m *TCEFChromium) CloseAllBrowsers() {
	_CEFChromium_CloseAllBrowses(m.instance)
}

func (m *TCEFChromium) CreateBrowser(window ITCefWindow) bool {
	if window.Type() == Wht_WindowParent {
		return _CEFChromium_CreateBrowseByWindow(m.instance, window.Instance())
	} else if window.Type() == Wht_LinkedWindowParent {
		return _CEFChromium_CreateBrowseByLinkedWindow(m.instance, window.Instance())
	}
	return false
}

func (m *TCEFChromium) Initialized() bool {
	return _CEFChromium_Initialized(m.instance)
}

func (m *TCEFChromium) BrowserId() int32 {
	return _CEFChromium_GetBrowserId(m.instance)
}

func (m *TCEFChromium) Browser() *ICefBrowser {
	return &ICefBrowser{
		browseId: m.BrowserId(),
	}
}

func (m *TCEFChromium) IsSameBrowser(browser *ICefBrowser) bool {
	return _CEFChromium_IsSameBrowser(m.instance, uintptr(browser.browseId))
}

func (m *TCEFChromium) PrintToPDF(saveFilePath string) {
	_CEFChromium_PrintToPDF(m.instance, saveFilePath)
}

func (m *TCEFChromium) Print() {
	_CEFChromium_Print(m.instance)
}

// 下载取消
func (m *TCEFChromium) BrowserDownloadCancel(browseId, downloadId int32) {
	_CEFChromium_BrowserDownloadCancel(uintptr(browseId), uintptr(downloadId))
}

// 下载暂停
func (m *TCEFChromium) BrowserDownloadPause(browseId, downloadId int32) {
	_CEFChromium_BrowserDownloadPause(uintptr(browseId), uintptr(downloadId))
}

func (m *TCEFChromium) BrowserZoom(zoom ZOOM) {
	_CEFChromium_BrowserZoom(m.instance, zoom)
}

func (m *TCEFChromium) GoBack() {
	_CEFChromium_GoBackForward(m.instance, BF_GOBACK)
}

func (m *TCEFChromium) GoForward() {
	_CEFChromium_GoBackForward(m.instance, BF_GOFORWARD)
}

func (m *TCEFChromium) NotifyMoveOrResizeStarted() {
	_CEFChromium_NotifyMoveOrResizeStarted(m.instance)
}

func (m *TCEFChromium) CloseBrowser(forceClose bool) {
	if m.IsValid() {
		_CEFChromium_CloseBrowser(m.instance, api.GoBoolToDBool(forceClose))
	}
}

func (m *TCEFChromium) ShowDevTools(window ITCefWindow) {
	if window == nil {
		_CEFChromium_ShowDevTools(m.instance)
	} else {
		_CEFChromium_ShowDevToolsByWindowParent(m.instance, window.Instance())
	}
}

func (m *TCEFChromium) CloseDevTools(window ITCefWindow) {
	if window == nil {
		_CEFChromium_CloseDevTools(m.instance)
	} else {
		_CEFChromium_CloseDevToolsByWindowParent(m.instance, window.Instance())
	}
}

// 查看所有cookie,该函数触发 OnCookiesVisited 事件返回结果
func (m *TCEFChromium) VisitAllCookies(id int32) {
	_CEFChromium_VisitAllCookies(m.instance, id)
}

// 查看指针URL cookie,该函数触发 OnCookiesVisited 事件返回结果
// url https://www.demo.com
func (m *TCEFChromium) VisitURLCookies(url string, includeHttpOnly bool, id int32) {
	_CEFChromium_VisitURLCookies(m.instance, url, includeHttpOnly, id)
}

// 删除所有cookie
func (m *TCEFChromium) DeleteCookies(url, cookieName string, deleteImmediately bool) {
	_CEFChromium_DeleteCookies(m.instance, url, cookieName, deleteImmediately)
}

func (m *TCEFChromium) SetCookie(url, name, value, domain, path string,
	secure, httponly, hasExpires bool,
	creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32) {
	_CEFChromium_SetCookie(m.instance, url, name, value, domain, path, secure, httponly, hasExpires, creation, lastAccess, expires, sameSite, priority, aSetImmediately, aID)
}

func (m *TCEFChromium) SetProxy(cefProxy *TCefProxy) {
	proxy := &tCefProxy{
		ProxyType:              uintptr(cefProxy.ProxyType),
		ProxyScheme:            uintptr(cefProxy.ProxyScheme),
		ProxyServer:            api.GoStrToDStr(cefProxy.ProxyServer),
		ProxyPort:              uintptr(cefProxy.ProxyPort),
		ProxyUsername:          api.GoStrToDStr(cefProxy.ProxyUsername),
		ProxyPassword:          api.GoStrToDStr(cefProxy.ProxyPassword),
		ProxyScriptURL:         api.GoStrToDStr(cefProxy.ProxyScriptURL),
		ProxyByPassList:        api.GoStrToDStr(cefProxy.ProxyByPassList),
		MaxConnectionsPerProxy: uintptr(cefProxy.MaxConnectionsPerProxy),
		CustomHeaderName:       api.GoStrToDStr(cefProxy.CustomHeaderName),
		CustomHeaderValue:      api.GoStrToDStr(cefProxy.CustomHeaderValue),
	}
	_CEFChromium_SetProxy(m.instance, proxy)
}

func (m *TCEFChromium) UpdatePreferences() {
	_CEFChromium_UpdatePreferences(m.instance)
}

func (m *TCEFChromium) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	_CEFChromium_ExecuteDevToolsMethod(m.instance, messageId, method, dictionaryValue)
}

// 发送进程消息 默认主browser 和 主frame
func (m *TCEFChromium) SendProcessMessage(targetProcess CefProcessId, processMessage *ipc.ICefProcessMessage) int {
	if processMessage == nil || processMessage.Name == "" || processMessage.ArgumentList == nil || ipc.InternalIPCNameCheck(processMessage.Name) {
		return -3
	}
	//var browser = Browsers.MainBrowser()
	var browser = BrowserWindow.GetBrowser(1)
	data := processMessage.ArgumentList.Package()

	r1 := _CEFFrame_SendProcessMessage(1, browser.MainFrame().Id, processMessage.Name, targetProcess, int32(processMessage.ArgumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return int(r1)
}

func (m *TCEFChromium) CreateClientHandler(client *ICefClient, alsOSR bool) bool {
	return api.DBoolToGoBool(_CEFChromium_CreateClientHandler(m.instance, client.instance, api.GoBoolToDBool(alsOSR)))
}
func (m *TCEFChromium) SetFocus(value bool) {
	_CEFChromium_SetFocus(m.instance, api.GoBoolToDBool(value))
}
func (m *TCEFChromium) SendCaptureLostEvent() {
	_CEFChromium_SendCaptureLostEvent(m.instance)
}
func (m *TCEFChromium) FrameIsFocused() bool {
	return api.DBoolToGoBool(_CEFChromium_FrameIsFocused(m.instance))
}

//--------TCEFChromium proc begin--------

// TCEFChromium _CEFChromium_Create
func _CEFChromium_Create(owner, config uintptr) uintptr {
	r1, _, _ := Proc("CEFChromium_Create").Call(owner, config)
	return r1
}

// TCEFChromium _CEFChromium_SetDefaultURL
func _CEFChromium_SetDefaultURL(instance uintptr, url string) {
	Proc("CEFChromium_SetDefaultURL").Call(instance, api.GoStrToDStr(url))
}

// TCEFChromium _CEFChromium_SetDefaultURL
func _CEFChromium_SetMultiBrowserMode(instance uintptr, url bool) {
	Proc("CEFChromium_SetMultiBrowserMode").Call(instance, api.GoBoolToDBool(url))
}

// TCEFChromium _CEFChromium_LoadURL
func _CEFChromium_LoadURL(instance uintptr, url string) {
	Proc("CEFChromium_LoadURL").Call(instance, api.GoStrToDStr(url))
}

// TCEFChromium _CEFChromium_LoadString
func _CEFChromium_LoadString(instance uintptr, html string) {
	Proc("CEFChromium_LoadString").Call(instance, api.GoStrToDStr(html))
}

// TCEFChromium _CEFChromium_StartDownload
func _CEFChromium_StartDownload(instance uintptr, url string) {
	Proc("CEFChromium_StartDownload").Call(instance, api.GoStrToDStr(url))
}

// TCEFChromium _CEFChromium_DownloadImage
func _CEFChromium_DownloadImage(instance uintptr, imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	Proc("CEFChromium_DownloadImage").Call(instance, api.GoStrToDStr(imageUrl), api.GoBoolToDBool(isFavicon), uintptr(maxImageSize), api.GoBoolToDBool(bypassCache))
}

// TCEFChromium _CEFChromium_Reload
func _CEFChromium_Reload(instance uintptr) {
	Proc("CEFChromium_Reload").Call(instance)
}

// TCEFChromium _CEFChromium_StopLoad
func _CEFChromium_StopLoad(instance uintptr) {
	Proc("CEFChromium_StopLoad").Call(instance)
}

// TCEFChromium _CEFChromium_ResetZoomLevel
func _CEFChromium_ResetZoomLevel(instance uintptr) {
	Proc("CEFChromium_ResetZoomLevel").Call(instance)
}

// TCEFChromium _CEFChromium_CloseAllBrowses
func _CEFChromium_CloseAllBrowses(instance uintptr) {
	Proc("CEFChromium_CloseAllBrowsers").Call(instance)
}

// TCEFChromium _CEFChromium_CreateBrowseByWindow
func _CEFChromium_CreateBrowseByWindow(instance, window uintptr) bool {
	r1, _, _ := Proc("CEFChromium_CreateBrowserByWindow").Call(instance, window)
	return api.DBoolToGoBool(r1)
}

// TCEFChromium _CEFChromium_CreateBrowseByLinkedWindow
func _CEFChromium_CreateBrowseByLinkedWindow(instance, window uintptr) bool {
	r1, _, _ := Proc("CEFChromium_CreateBrowserByLinkedWindow").Call(instance, window)
	return api.DBoolToGoBool(r1)
}

// TCEFChromium _CEFChromium_Initialized
func _CEFChromium_Initialized(instance uintptr) bool {
	r1, _, _ := Proc("CEFChromium_Initialized").Call(instance)
	return api.DBoolToGoBool(r1)
}

// TCEFChromium _CEFChromium_GetBrowserId
func _CEFChromium_GetBrowserId(instance uintptr) int32 {
	r1, _, _ := Proc("CEFChromium_GetBrowserId").Call(instance)
	return int32(r1)
}

// TCEFChromium _CEFChromium_IsSameBrowser
func _CEFChromium_IsSameBrowser(instance, browser uintptr) bool {
	r1, _, _ := Proc("CEFChromium_IsSameBrowser").Call(instance, browser)
	return api.DBoolToGoBool(r1)
}

// TCEFChromium _CEFChromium_PrintToPDF
func _CEFChromium_PrintToPDF(instance uintptr, saveFilePath string) {
	Proc("CEFChromium_PrintToPDF").Call(instance, api.GoStrToDStr(saveFilePath))
}

// TCEFChromium _CEFChromium_Print
func _CEFChromium_Print(instance uintptr) {
	Proc("CEFChromium_Print").Call(instance)
}

// TCEFChromium _CEFChromium_BrowserDownloadCancel
func _CEFChromium_BrowserDownloadCancel(browseId, downloadId uintptr) {
	Proc("CEFChromium_BrowserDownloadCancel").Call(browseId, downloadId)
}

// TCEFChromium _CEFChromium_BrowserDownloadPause
func _CEFChromium_BrowserDownloadPause(browseId, downloadId uintptr) {
	Proc("CEFChromium_BrowserDownloadPause").Call(browseId, downloadId)
}

// TCEFChromium _CEFChromium_DownloadResume
func _CEFChromium_DownloadResume(browseId, downloadId uintptr) {
	Proc("CEFChromium_DownloadResume").Call(browseId, downloadId)
}

// TCEFChromium _CEFChromium_BrowserZoom
func _CEFChromium_BrowserZoom(instance uintptr, zoom ZOOM) {
	Proc("CEFChromium_BrowserZoom").Call(instance, uintptr(zoom))
}

// TCEFChromium _CEFChromium_GoBackForward
func _CEFChromium_GoBackForward(instance uintptr, bf BF) {
	Proc("CEFChromium_GoBackForward").Call(instance, uintptr(bf))
}

// TCEFChromium _CEFChromium_NotifyMoveOrResizeStarted
func _CEFChromium_NotifyMoveOrResizeStarted(instance uintptr) {
	Proc("CEFChromium_NotifyMoveOrResizeStarted").Call(instance)
}

// TCEFChromium _CEFChromium_CloseBrowser
func _CEFChromium_CloseBrowser(instance, forceClose uintptr) {
	Proc("CEFChromium_CloseBrowser").Call(instance, forceClose)
}

// TCEFChromium _CEFChromium_ExecuteJavaScript
func _CEFChromium_ExecuteJavaScript(instance uintptr, code, scriptURL string, startLine int32) {
	Proc("CEFChromium_ExecuteJavaScript").Call(instance, api.GoStrToDStr(code), api.GoStrToDStr(scriptURL), uintptr(startLine))
}

// TCEFChromium _CEFChromium_ShowDevTools
func _CEFChromium_ShowDevTools(instance uintptr) {
	Proc("CEFChromium_ShowDevTools").Call(instance)
}
func _CEFChromium_ShowDevToolsByWindowParent(instance, windowParent uintptr) {
	Proc("CEFChromium_ShowDevToolsByWindowParent").Call(instance, windowParent)
}

// TCEFChromium _CEFChromium_CloseDevTools
func _CEFChromium_CloseDevTools(instance uintptr) {
	Proc("CEFChromium_CloseDevTools").Call(instance)
}
func _CEFChromium_CloseDevToolsByWindowParent(instance, windowParent uintptr) {
	Proc("CEFChromium_CloseDevToolsByWindowParent").Call(instance, windowParent)
}

// TCEFChromium _CEFChromium_VisitAllCookies
func _CEFChromium_VisitAllCookies(instance uintptr, id int32) {
	Proc("CEFChromium_VisitAllCookies").Call(instance, uintptr(id))
}

// TCEFChromium _CEFChromium_VisitURLCookies
func _CEFChromium_VisitURLCookies(instance uintptr, url string, includeHttpOnly bool, id int32) {
	Proc("CEFChromium_VisitURLCookies").Call(instance, api.GoStrToDStr(url), api.GoBoolToDBool(includeHttpOnly), uintptr(id))
}

// TCEFChromium _CEFChromium_DeleteCookies
func _CEFChromium_DeleteCookies(instance uintptr, url, cookieName string, deleteImmediately bool) {
	Proc("CEFChromium_DeleteCookies").Call(instance, api.GoStrToDStr(url), api.GoStrToDStr(cookieName), api.GoBoolToDBool(deleteImmediately))
}

// TCEFChromium _CEFChromium_SetCookie
func _CEFChromium_SetCookie(instance uintptr, url, name, value, domain, path string,
	secure, httponly, hasExpires bool,
	creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32) {
	creationPtr := GoDateTimeToDDateTime(creation)
	lastAccessPtr := GoDateTimeToDDateTime(lastAccess)
	expiresPtr := GoDateTimeToDDateTime(expires)
	cCookie := &cefCookie{
		url:             api.GoStrToDStr(url),
		name:            api.GoStrToDStr(name),
		value:           api.GoStrToDStr(value),
		domain:          api.GoStrToDStr(domain),
		path:            api.GoStrToDStr(path),
		secure:          api.GoBoolToDBool(secure),
		httponly:        api.GoBoolToDBool(httponly),
		hasExpires:      api.GoBoolToDBool(hasExpires),
		creation:        uintptr(unsafe.Pointer(&creationPtr)),
		lastAccess:      uintptr(unsafe.Pointer(&lastAccessPtr)),
		expires:         uintptr(unsafe.Pointer(&expiresPtr)),
		sameSite:        uintptr(sameSite),
		priority:        uintptr(priority),
		aSetImmediately: api.GoBoolToDBool(aSetImmediately),
		aID:             uintptr(aID),
		aDeleteCookie:   uintptr(0),
		aResult:         uintptr(0),
		count:           uintptr(0),
		total:           uintptr(0),
	}
	Proc("CEFChromium_SetCookie").Call(instance, uintptr(unsafe.Pointer(cCookie)))
	cCookie = nil
}

// TCEFChromium  _CEFChromium_SetProxy
func _CEFChromium_SetProxy(instance uintptr, proxy *tCefProxy) {
	Proc("CEFChromium_SetProxy").Call(instance, uintptr(unsafe.Pointer(proxy)))
}

// TCEFChromium  _CEFChromium_UpdatePreferences
func _CEFChromium_UpdatePreferences(instance uintptr) {
	Proc("CEFChromium_UpdatePreferences").Call(instance)
}

// TCEFChromium  _CEFChromium_ExecuteDevToolsMethod
func _CEFChromium_ExecuteDevToolsMethod(instance uintptr, messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	var data = []byte{}
	var dataPtr unsafe.Pointer
	var dataLen int32 = 0
	var argsLen int32 = 0
	if dictionaryValue != nil && dictionaryValue.dataLen > 0 {
		defer func() {
			dictionaryValue.Clear()
			dictionaryValue = nil
			data = nil
			dataPtr = nil
		}()
		data = dictionaryValue.Package()
		argsLen = int32(dictionaryValue.dataLen)
		dataPtr = unsafe.Pointer(&data[0])
		dataLen = int32(len(data) - 1)
	} else {
		dataPtr = unsafe.Pointer(&data)
	}
	Proc("CEFChromium_ExecuteDevToolsMethod").Call(instance, uintptr(messageId), api.GoStrToDStr(method), uintptr(argsLen), uintptr(dataPtr), uintptr(dataLen))
}

// TCEFChromium  _CEFChromium_CreateClientHandler
func _CEFChromium_CreateClientHandler(instance, client, alsOSR uintptr) uintptr {
	r1, _, _ := Proc("CEFChromium_CreateClientHandler").Call(instance, client, alsOSR)
	return r1
}

func _CEFChromium_SetFocus(instance, value uintptr) {
	Proc("CEFChromium_SetFocus").Call(instance, value)
}

func _CEFChromium_SendCaptureLostEvent(instance uintptr) {
	Proc("CEFChromium_SendCaptureLostEvent").Call(instance)
}

func _CEFChromium_FrameIsFocused(instance uintptr) uintptr {
	r1, _, _ := Proc("CEFChromium_FrameIsFocused").Call(instance)
	return r1
}

//--------TCEFChromium proc end--------
