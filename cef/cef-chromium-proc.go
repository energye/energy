//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"strings"
	"time"
	"unsafe"
)

type IChromiumProc interface {
	lcl.IObject
	On(name string, eventCallback ipc.EventCallback)
	ExecuteJavaScript(code, scriptURL string, startLine int32) //单纯的执行js，没有返回值
	Emit(eventName string, args ipc.IArgumentList, target IEmitTarget) ProcessMessageError
	EmitAndCallback(eventName string, args ipc.IArgumentList, target IEmitTarget, callback ipc.IPCCallback) ProcessMessageError
	EmitAndReturn(eventName string, args ipc.IArgumentList, target IEmitTarget) (ipc.IIPCContext, ProcessMessageError)
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
	CreateBrowser(window ITCefWindowParent) bool
	CreateBrowserByBrowserViewComponent(homePage string, browserViewComponent *TCEFBrowserViewComponent) bool
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
	ShowDevTools(window ITCefWindowParent)
	CloseDevTools(window ITCefWindowParent)
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
	TryCloseBrowser() bool
	BrowserHandle() types.HWND
	WidgetHandle() types.HWND
	RenderHandle() types.HWND
	SetCustomHeader(customHeader *TCustomHeader)
	CustomHeader() *TCustomHeader
	SetJavascriptEnabled(value bool)
	JavascriptEnabled() bool
	SetWebRTCIPHandlingPolicy(value TCefWebRTCHandlingPolicy)
	WebRTCIPHandlingPolicy() TCefWebRTCHandlingPolicy
	SetWebRTCMultipleRoutes(value TCefState)
	WebRTCMultipleRoutes() TCefState
	SetWebRTCNonproxiedUDP(value TCefState)
	WebRTCNonproxiedUDP() TCefState
	SetBatterySaverModeState(value TCefBatterySaverModeState)
	BatterySaverModeState() TCefBatterySaverModeState
	SetHighEfficiencyMode(value TCefState)
	HighEfficiencyMode() TCefState
	SetLoadImagesAutomatically(value bool)
	LoadImagesAutomatically() bool
	SetQuicAllowed(value bool)
	QuicAllowed() bool
	SetOffline(value bool)
	Offline() bool
	SetDefaultWindowInfoExStyle(exStyle types.DWORD)
	DefaultWindowInfoExStyle() types.DWORD
	SetBlock3rdPartyCookies(value bool)
	Block3rdPartyCookies() bool
	SetAcceptCookies(cp TCefCookiePref)
	AcceptCookies() TCefCookiePref
	SetAcceptLanguageList(languageList string)
	AcceptLanguageList() string
	SetPrintingEnabled(value bool)
	PrintingEnabled() bool
	SetYouTubeRestrict(value bool)
	YouTubeRestrict() bool
	SetSafeSearch(value bool)
	SafeSearch() bool
	SetAudioMuted(value bool)
	AudioMuted() bool
	SetDragOperations(value TCefDragOperations)
	DragOperations() TCefDragOperations
	FrameCount() uint32
	SetSpellCheckerDicts(value string)
	SpellCheckerDicts() string
	SetSpellChecking(value bool)
	SpellChecking() bool
	SetAlwaysOpenPDFExternally(value bool)
	AlwaysOpenPDFExternally() bool
	SetAlwaysAuthorizePlugins(value bool)
	AlwaysAuthorizePlugins() bool
	SetAllowOutdatedPlugins(value bool)
	AllowOutdatedPlugins() bool
	SetSendReferrer(value bool)
	SendReferrer() bool
	SetDoNotTrack(value bool)
	DoNotTrack() bool
	SetZoomStep(value int8)
	ZoomStep() int8
	SetZoomPct(value float64)
	ZoomPct() float64
	SetZoomLevel(value float64)
	ZoomLevel() float64
	SetDefaultEncoding(value string)
	DefaultEncoding() string
}

func (m *TCEFChromium) IsValid() bool {
	return m.instance != nil
}

func (m *TCEFChromium) UnsafeAddr() unsafe.Pointer {
	return m.instance
}

func (m *TCEFChromium) ClassName() string {
	r1, _, _ := imports.Proc(internale_CEFChromium_ClassName).Call()
	return api.GoStr(r1)
}

func (m *TCEFChromium) Free() {
	imports.Proc(internale_CEFChromium_Free).Call()
}

func (m *TCEFChromium) HashCode() int32 {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetHashCode).Call()
	return int32(r1)
}

func (m *TCEFChromium) Equals(object lcl.IObject) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_Equals).Call(lcl.CheckPtr(object))
	return api.GoBool(r1)
}

func (m *TCEFChromium) ClassType() types.TClass {
	r1, _, _ := imports.Proc(internale_CEFChromium_ClassType).Call()
	return types.TClass(r1)
}

func (m *TCEFChromium) InstanceSize() int32 {
	r1, _, _ := imports.Proc(internale_CEFChromium_InstanceSize).Call()
	return int32(r1)
}

func (m *TCEFChromium) InheritsFrom(class types.TClass) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_InheritsFrom).Call(uintptr(class))
	return api.GoBool(r1)
}

func (m *TCEFChromium) ToString() string {
	r1, _, _ := imports.Proc(internale_CEFChromium_ToString).Call()
	return api.GoStr(r1)
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
	_CEFChromium_SetDefaultURL(m.Instance(), defaultURL)
}

func (m *TCEFChromium) SetEnableMultiBrowserMode(enableMultiBrowserMode bool) {
	_CEFChromium_SetMultiBrowserMode(m.Instance(), enableMultiBrowserMode)
}

func (m *TCEFChromium) LoadUrl(url string) {
	_CEFChromium_LoadURL(m.Instance(), url)
}

func (m *TCEFChromium) LoadHtml(html string) {
	_CEFChromium_LoadString(m.Instance(), html)
}

func (m *TCEFChromium) StartDownload(url string) {
	_CEFChromium_StartDownload(m.Instance(), url)
}

func (m *TCEFChromium) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	_CEFChromium_DownloadImage(m.Instance(), imageUrl, isFavicon, maxImageSize, bypassCache)
}

func (m *TCEFChromium) Reload() {
	_CEFChromium_Reload(m.Instance())
}

func (m *TCEFChromium) StopLoad() {
	_CEFChromium_StopLoad(m.Instance())
}

func (m *TCEFChromium) ResetZoomLevel() {
	_CEFChromium_ResetZoomLevel(m.Instance())
}

func (m *TCEFChromium) CloseAllBrowsers() {
	_CEFChromium_CloseAllBrowses(m.Instance())
}

func (m *TCEFChromium) CreateBrowser(window ITCefWindowParent) bool {
	if window.Type() == Wht_WindowParent {
		return _CEFChromium_CreateBrowseByWindow(m.Instance(), window.Instance())
	} else if window.Type() == Wht_LinkedWindowParent {
		return _CEFChromium_CreateBrowseByLinkedWindow(m.Instance(), window.Instance())
	}
	return false
}

func (m *TCEFChromium) CreateBrowserByBrowserViewComponent(homePage string, browserViewComponent *TCEFBrowserViewComponent) bool {
	return _CEFChromium_CreateBrowserByBrowserViewComponent(m.Instance(), api.PascalStr(homePage), browserViewComponent.Instance())
}

func (m *TCEFChromium) Initialized() bool {
	return _CEFChromium_Initialized(m.Instance())
}

func (m *TCEFChromium) BrowserId() int32 {
	return _CEFChromium_GetBrowserId(m.Instance())
}

func (m *TCEFChromium) Browser() *ICefBrowser {
	return &ICefBrowser{
		browseId: m.BrowserId(),
	}
}

func (m *TCEFChromium) IsSameBrowser(browser *ICefBrowser) bool {
	return _CEFChromium_IsSameBrowser(m.Instance(), uintptr(browser.browseId))
}

func (m *TCEFChromium) PrintToPDF(saveFilePath string) {
	_CEFChromium_PrintToPDF(m.Instance(), saveFilePath)
}

func (m *TCEFChromium) Print() {
	_CEFChromium_Print(m.Instance())
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
	_CEFChromium_BrowserZoom(m.Instance(), zoom)
}

func (m *TCEFChromium) GoBack() {
	_CEFChromium_GoBackForward(m.Instance(), BF_GOBACK)
}

func (m *TCEFChromium) GoForward() {
	_CEFChromium_GoBackForward(m.Instance(), BF_GOFORWARD)
}

func (m *TCEFChromium) NotifyMoveOrResizeStarted() {
	_CEFChromium_NotifyMoveOrResizeStarted(m.Instance())
}

func (m *TCEFChromium) CloseBrowser(forceClose bool) {
	_CEFChromium_CloseBrowser(m.Instance(), api.PascalBool(forceClose))
}

func (m *TCEFChromium) ShowDevTools(window ITCefWindowParent) {
	if window == nil {
		_CEFChromium_ShowDevTools(m.Instance())
	} else {
		_CEFChromium_ShowDevToolsByWindowParent(m.Instance(), window.Instance())
	}
}

func (m *TCEFChromium) CloseDevTools(window ITCefWindowParent) {
	if window == nil {
		_CEFChromium_CloseDevTools(m.Instance())
	} else {
		_CEFChromium_CloseDevToolsByWindowParent(m.Instance(), window.Instance())
	}
}

// 查看所有cookie,该函数触发 OnCookiesVisited 事件返回结果
func (m *TCEFChromium) VisitAllCookies(id int32) {
	_CEFChromium_VisitAllCookies(m.Instance(), id)
}

// 查看指针URL cookie,该函数触发 OnCookiesVisited 事件返回结果
// url https://www.demo.com
func (m *TCEFChromium) VisitURLCookies(url string, includeHttpOnly bool, id int32) {
	_CEFChromium_VisitURLCookies(m.Instance(), url, includeHttpOnly, id)
}

// 删除所有cookie
func (m *TCEFChromium) DeleteCookies(url, cookieName string, deleteImmediately bool) {
	_CEFChromium_DeleteCookies(m.Instance(), url, cookieName, deleteImmediately)
}

func (m *TCEFChromium) SetCookie(url, name, value, domain, path string,
	secure, httponly, hasExpires bool,
	creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32) {
	_CEFChromium_SetCookie(m.Instance(), url, name, value, domain, path, secure, httponly, hasExpires, creation, lastAccess, expires, sameSite, priority, aSetImmediately, aID)
}

func (m *TCEFChromium) SetProxy(cefProxy *TCefProxy) {
	proxy := &tCefProxyPtr{
		ProxyType:              uintptr(cefProxy.ProxyType),
		ProxyScheme:            uintptr(cefProxy.ProxyScheme),
		ProxyServer:            api.PascalStr(cefProxy.ProxyServer),
		ProxyPort:              uintptr(cefProxy.ProxyPort),
		ProxyUsername:          api.PascalStr(cefProxy.ProxyUsername),
		ProxyPassword:          api.PascalStr(cefProxy.ProxyPassword),
		ProxyScriptURL:         api.PascalStr(cefProxy.ProxyScriptURL),
		ProxyByPassList:        api.PascalStr(cefProxy.ProxyByPassList),
		MaxConnectionsPerProxy: uintptr(cefProxy.MaxConnectionsPerProxy),
	}
	_CEFChromium_SetProxy(m.Instance(), proxy)
}

func (m *TCEFChromium) UpdatePreferences() {
	_CEFChromium_UpdatePreferences(m.Instance())
}

func (m *TCEFChromium) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	_CEFChromium_ExecuteDevToolsMethod(m.Instance(), messageId, method, dictionaryValue)
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
	return api.GoBool(_CEFChromium_CreateClientHandler(m.Instance(), uintptr(client.instance), api.PascalBool(alsOSR)))
}

func (m *TCEFChromium) SetFocus(value bool) {
	_CEFChromium_SetFocus(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SendCaptureLostEvent() {
	_CEFChromium_SendCaptureLostEvent(m.Instance())
}

func (m *TCEFChromium) FrameIsFocused() bool {
	return api.GoBool(_CEFChromium_FrameIsFocused(m.Instance()))
}

func (m *TCEFChromium) TryCloseBrowser() bool {
	return api.GoBool(_CEFChromium_TryCloseBrowser(m.Instance()))
}

func (m *TCEFChromium) BrowserHandle() types.HWND {
	if m.browserHandle == 0 {
		m.browserHandle = types.HWND(_CEFChromium_BrowserHandle(m.Instance()))
	}
	return m.browserHandle
}

func (m *TCEFChromium) WidgetHandle() types.HWND {
	if m.widgetHandle == 0 {
		m.widgetHandle = types.HWND(_CEFChromium_WidgetHandle(m.Instance()))
	}
	return m.widgetHandle
}

func (m *TCEFChromium) RenderHandle() types.HWND {
	if m.renderHandle == 0 {
		m.renderHandle = types.HWND(_CEFChromium_RenderHandle(m.Instance()))
	}
	return m.renderHandle
}

func (m *TCEFChromium) SetCustomHeader(customHeader *TCustomHeader) {
	ptrCustomHeader := &tCustomHeader{
		CustomHeaderName:  api.PascalStr(customHeader.CustomHeaderName),
		CustomHeaderValue: api.PascalStr(customHeader.CustomHeaderValue),
	}
	imports.Proc(internale_CEFChromium_SetCustomHeader).Call(m.Instance(), uintptr(unsafe.Pointer(ptrCustomHeader)))
}

func (m *TCEFChromium) CustomHeader() *TCustomHeader {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetCustomHeader).Call(m.Instance())
	ptrCustomHeader := (*tCustomHeader)(unsafe.Pointer(r1))
	return &TCustomHeader{
		CustomHeaderName:  api.GoStr(ptrCustomHeader.CustomHeaderName),
		CustomHeaderValue: api.GoStr(ptrCustomHeader.CustomHeaderValue),
	}
}

func (m *TCEFChromium) SetJavascriptEnabled(value bool) {
	imports.Proc(internale_CEFChromium_SetJavascriptEnabled).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) JavascriptEnabled() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetJavascriptEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetWebRTCIPHandlingPolicy(value TCefWebRTCHandlingPolicy) {
	imports.Proc(internale_CEFChromium_SetWebRTCIPHandlingPolicy).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) WebRTCIPHandlingPolicy() TCefWebRTCHandlingPolicy {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetWebRTCIPHandlingPolicy).Call(m.Instance())
	return TCefWebRTCHandlingPolicy(r1)
}

func (m *TCEFChromium) SetWebRTCMultipleRoutes(value TCefState) {
	imports.Proc(internale_CEFChromium_SetWebRTCMultipleRoutes).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) WebRTCMultipleRoutes() TCefState {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetWebRTCMultipleRoutes).Call(m.Instance())
	return TCefState(r1)
}

func (m *TCEFChromium) SetWebRTCNonproxiedUDP(value TCefState) {
	imports.Proc(internale_CEFChromium_SetWebRTCNonproxiedUDP).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) WebRTCNonproxiedUDP() TCefState {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetWebRTCNonproxiedUDP).Call(m.Instance())
	return TCefState(r1)
}

func (m *TCEFChromium) SetBatterySaverModeState(value TCefBatterySaverModeState) {
	imports.Proc(internale_CEFChromium_SetBatterySaverModeState).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) BatterySaverModeState() TCefBatterySaverModeState {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetBatterySaverModeState).Call(m.Instance())
	return TCefBatterySaverModeState(r1)
}

func (m *TCEFChromium) SetHighEfficiencyMode(value TCefState) {
	imports.Proc(internale_CEFChromium_SetHighEfficiencyMode).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) HighEfficiencyMode() TCefState {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetHighEfficiencyMode).Call(m.Instance())
	return TCefState(r1)
}

func (m *TCEFChromium) SetLoadImagesAutomatically(value bool) {
	imports.Proc(internale_CEFChromium_SetLoadImagesAutomatically).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) LoadImagesAutomatically() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetLoadImagesAutomatically).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetQuicAllowed(value bool) {
	imports.Proc(internale_CEFChromium_SetQuicAllowed).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) QuicAllowed() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetQuicAllowed).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetOffline(value bool) {
	imports.Proc(internale_CEFChromium_SetOffline).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) Offline() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetOffline).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetDefaultWindowInfoExStyle(exStyle types.DWORD) {
	imports.Proc(internale_CEFChromium_SetDefaultWindowInfoExStyle).Call(m.Instance(), uintptr(exStyle))
}

func (m *TCEFChromium) DefaultWindowInfoExStyle() types.DWORD {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetDefaultWindowInfoExStyle).Call(m.Instance())
	return types.DWORD(r1)
}

func (m *TCEFChromium) SetBlock3rdPartyCookies(value bool) {
	imports.Proc(internale_CEFChromium_SetBlock3rdPartyCookies).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) Block3rdPartyCookies() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetBlock3rdPartyCookies).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAcceptCookies(cp TCefCookiePref) {
	imports.Proc(internale_CEFChromium_SetAcceptCookies).Call(m.Instance(), cp.ToPtr())
}

func (m *TCEFChromium) AcceptCookies() TCefCookiePref {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetAcceptCookies).Call(m.Instance())
	return TCefCookiePref(r1)
}

func (m *TCEFChromium) SetAcceptLanguageList(languageList string) {
	imports.Proc(internale_CEFChromium_SetAcceptLanguageList).Call(m.Instance(), api.PascalStr(languageList))
}

func (m *TCEFChromium) AcceptLanguageList() string {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetAcceptLanguageList).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFChromium) SetPrintingEnabled(value bool) {
	imports.Proc(internale_CEFChromium_SetPrintingEnabled).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) PrintingEnabled() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetPrintingEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetYouTubeRestrict(value bool) {
	imports.Proc(internale_CEFChromium_SetYouTubeRestrict).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) YouTubeRestrict() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetYouTubeRestrict).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetSafeSearch(value bool) {
	imports.Proc(internale_CEFChromium_SetSafeSearch).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SafeSearch() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetSafeSearch).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAudioMuted(value bool) {
	imports.Proc(internale_CEFChromium_SetAudioMuted).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AudioMuted() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetAudioMuted).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetDragOperations(value TCefDragOperations) {
	imports.Proc(internale_CEFChromium_SetDragOperations).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) DragOperations() TCefDragOperations {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetDragOperations).Call(m.Instance())
	return TCefDragOperations(r1)
}

func (m *TCEFChromium) FrameCount() uint32 {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetFrameCount).Call(m.Instance())
	return uint32(r1)
}

func (m *TCEFChromium) SetSpellCheckerDicts(value string) {
	imports.Proc(internale_CEFChromium_SetSpellCheckerDicts).Call(m.Instance(), api.PascalStr(value))
}

func (m *TCEFChromium) SpellCheckerDicts() string {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetSpellCheckerDicts).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFChromium) SetSpellChecking(value bool) {
	imports.Proc(internale_CEFChromium_SetSpellChecking).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SpellChecking() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetSpellChecking).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAlwaysOpenPDFExternally(value bool) {
	imports.Proc(internale_CEFChromium_SetAlwaysOpenPDFExternally).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AlwaysOpenPDFExternally() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetAlwaysOpenPDFExternally).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAlwaysAuthorizePlugins(value bool) {
	imports.Proc(internale_CEFChromium_SetAlwaysAuthorizePlugins).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AlwaysAuthorizePlugins() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetAlwaysAuthorizePlugins).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAllowOutdatedPlugins(value bool) {
	imports.Proc(internale_CEFChromium_SetAllowOutdatedPlugins).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AllowOutdatedPlugins() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetAllowOutdatedPlugins).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetSendReferrer(value bool) {
	imports.Proc(internale_CEFChromium_SetSendReferrer).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SendReferrer() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetSendReferrer).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetDoNotTrack(value bool) {
	imports.Proc(internale_CEFChromium_SetDoNotTrack).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) DoNotTrack() bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetDoNotTrack).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetZoomStep(value int8) {
	imports.Proc(internale_CEFChromium_SetZoomStep).Call(m.Instance(), uintptr(value))
}

func (m *TCEFChromium) ZoomStep() int8 {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetZoomStep).Call(m.Instance())
	return int8(r1)
}

func (m *TCEFChromium) SetZoomPct(value float64) {
	imports.Proc(internale_CEFChromium_SetZoomPct).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}

func (m *TCEFChromium) ZoomPct() (result float64) {
	imports.Proc(internale_CEFChromium_GetZoomPct).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFChromium) SetZoomLevel(value float64) {
	imports.Proc(internale_CEFChromium_SetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}

func (m *TCEFChromium) ZoomLevel() (result float64) {
	imports.Proc(internale_CEFChromium_GetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFChromium) SetDefaultEncoding(value string) {
	imports.Proc(internale_CEFChromium_SetDefaultEncoding).Call(m.Instance(), api.PascalStr(value))
}

func (m *TCEFChromium) DefaultEncoding() string {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetDefaultEncoding).Call(m.Instance())
	return api.GoStr(r1)
}

//--------TCEFChromium proc begin--------

// TCEFChromium _CEFChromium_Create
func _CEFChromium_Create(owner, config uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_Create).Call(owner, config)
	return r1
}

// TCEFChromium _CEFChromium_SetDefaultURL
func _CEFChromium_SetDefaultURL(instance uintptr, url string) {
	imports.Proc(internale_CEFChromium_SetDefaultURL).Call(instance, api.PascalStr(url))
}

// TCEFChromium _CEFChromium_SetDefaultURL
func _CEFChromium_SetMultiBrowserMode(instance uintptr, url bool) {
	imports.Proc(internale_CEFChromium_SetMultiBrowserMode).Call(instance, api.PascalBool(url))
}

// TCEFChromium _CEFChromium_LoadURL
func _CEFChromium_LoadURL(instance uintptr, url string) {
	imports.Proc(internale_CEFChromium_LoadURL).Call(instance, api.PascalStr(url))
}

// TCEFChromium _CEFChromium_LoadString
func _CEFChromium_LoadString(instance uintptr, html string) {
	imports.Proc(internale_CEFChromium_LoadString).Call(instance, api.PascalStr(html))
}

// TCEFChromium _CEFChromium_StartDownload
func _CEFChromium_StartDownload(instance uintptr, url string) {
	imports.Proc(internale_CEFChromium_StartDownload).Call(instance, api.PascalStr(url))
}

// TCEFChromium _CEFChromium_DownloadImage
func _CEFChromium_DownloadImage(instance uintptr, imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	imports.Proc(internale_CEFChromium_DownloadImage).Call(instance, api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache))
}

// TCEFChromium _CEFChromium_Reload
func _CEFChromium_Reload(instance uintptr) {
	imports.Proc(internale_CEFChromium_Reload).Call(instance)
}

// TCEFChromium _CEFChromium_StopLoad
func _CEFChromium_StopLoad(instance uintptr) {
	imports.Proc(internale_CEFChromium_StopLoad).Call(instance)
}

// TCEFChromium _CEFChromium_ResetZoomLevel
func _CEFChromium_ResetZoomLevel(instance uintptr) {
	imports.Proc(internale_CEFChromium_ResetZoomLevel).Call(instance)
}

// TCEFChromium _CEFChromium_CloseAllBrowses
func _CEFChromium_CloseAllBrowses(instance uintptr) {
	imports.Proc(internale_CEFChromium_CloseAllBrowsers).Call(instance)
}

// TCEFChromium _CEFChromium_CreateBrowseByWindow
func _CEFChromium_CreateBrowseByWindow(instance, window uintptr) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_CreateBrowserByWindow).Call(instance, window)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_CreateBrowseByLinkedWindow
func _CEFChromium_CreateBrowseByLinkedWindow(instance, window uintptr) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_CreateBrowserByLinkedWindow).Call(instance, window)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_CreateBrowserByBrowserViewComponent
func _CEFChromium_CreateBrowserByBrowserViewComponent(instance, homePage, browserViewComponent uintptr) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_CreateBrowserByBrowserViewComponent).Call(instance, homePage, browserViewComponent)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_Initialized
func _CEFChromium_Initialized(instance uintptr) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_Initialized).Call(instance)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_GetBrowserId
func _CEFChromium_GetBrowserId(instance uintptr) int32 {
	r1, _, _ := imports.Proc(internale_CEFChromium_GetBrowserId).Call(instance)
	return int32(r1)
}

// TCEFChromium _CEFChromium_IsSameBrowser
func _CEFChromium_IsSameBrowser(instance, browser uintptr) bool {
	r1, _, _ := imports.Proc(internale_CEFChromium_IsSameBrowser).Call(instance, browser)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_PrintToPDF
func _CEFChromium_PrintToPDF(instance uintptr, saveFilePath string) {
	imports.Proc(internale_CEFChromium_PrintToPDF).Call(instance, api.PascalStr(saveFilePath))
}

// TCEFChromium _CEFChromium_Print
func _CEFChromium_Print(instance uintptr) {
	imports.Proc(internale_CEFChromium_Print).Call(instance)
}

// TCEFChromium _CEFChromium_BrowserDownloadCancel
func _CEFChromium_BrowserDownloadCancel(browseId, downloadId uintptr) {
	imports.Proc(internale_CEFChromium_BrowserDownloadCancel).Call(browseId, downloadId)
}

// TCEFChromium _CEFChromium_BrowserDownloadPause
func _CEFChromium_BrowserDownloadPause(browseId, downloadId uintptr) {
	imports.Proc(internale_CEFChromium_BrowserDownloadPause).Call(browseId, downloadId)
}

// TCEFChromium _CEFChromium_DownloadResume
func _CEFChromium_DownloadResume(browseId, downloadId uintptr) {
	imports.Proc(internale_CEFChromium_DownloadResume).Call(browseId, downloadId)
}

// TCEFChromium _CEFChromium_BrowserZoom
func _CEFChromium_BrowserZoom(instance uintptr, zoom ZOOM) {
	imports.Proc(internale_CEFChromium_BrowserZoom).Call(instance, uintptr(zoom))
}

// TCEFChromium _CEFChromium_GoBackForward
func _CEFChromium_GoBackForward(instance uintptr, bf BF) {
	imports.Proc(internale_CEFChromium_GoBackForward).Call(instance, uintptr(bf))
}

// TCEFChromium _CEFChromium_NotifyMoveOrResizeStarted
func _CEFChromium_NotifyMoveOrResizeStarted(instance uintptr) {
	imports.Proc(internale_CEFChromium_NotifyMoveOrResizeStarted).Call(instance)
}

// TCEFChromium _CEFChromium_CloseBrowser
func _CEFChromium_CloseBrowser(instance, forceClose uintptr) {
	imports.Proc(internale_CEFChromium_CloseBrowser).Call(instance, forceClose)
}

// TCEFChromium _CEFChromium_ExecuteJavaScript
func _CEFChromium_ExecuteJavaScript(instance uintptr, code, scriptURL string, startLine int32) {
	imports.Proc(internale_CEFChromium_ExecuteJavaScript).Call(instance, api.PascalStr(code), api.PascalStr(scriptURL), uintptr(startLine))
}

// TCEFChromium _CEFChromium_ShowDevTools
func _CEFChromium_ShowDevTools(instance uintptr) {
	imports.Proc(internale_CEFChromium_ShowDevTools).Call(instance)
}
func _CEFChromium_ShowDevToolsByWindowParent(instance, windowParent uintptr) {
	imports.Proc(internale_CEFChromium_ShowDevToolsByWindowParent).Call(instance, windowParent)
}

// TCEFChromium _CEFChromium_CloseDevTools
func _CEFChromium_CloseDevTools(instance uintptr) {
	imports.Proc(internale_CEFChromium_CloseDevTools).Call(instance)
}
func _CEFChromium_CloseDevToolsByWindowParent(instance, windowParent uintptr) {
	imports.Proc(internale_CEFChromium_CloseDevToolsByWindowParent).Call(instance, windowParent)
}

// TCEFChromium _CEFChromium_VisitAllCookies
func _CEFChromium_VisitAllCookies(instance uintptr, id int32) {
	imports.Proc(internale_CEFChromium_VisitAllCookies).Call(instance, uintptr(id))
}

// TCEFChromium _CEFChromium_VisitURLCookies
func _CEFChromium_VisitURLCookies(instance uintptr, url string, includeHttpOnly bool, id int32) {
	imports.Proc(internale_CEFChromium_VisitURLCookies).Call(instance, api.PascalStr(url), api.PascalBool(includeHttpOnly), uintptr(id))
}

// TCEFChromium _CEFChromium_DeleteCookies
func _CEFChromium_DeleteCookies(instance uintptr, url, cookieName string, deleteImmediately bool) {
	imports.Proc(internale_CEFChromium_DeleteCookies).Call(instance, api.PascalStr(url), api.PascalStr(cookieName), api.PascalBool(deleteImmediately))
}

// TCEFChromium _CEFChromium_SetCookie
func _CEFChromium_SetCookie(instance uintptr, url, name, value, domain, path string,
	secure, httponly, hasExpires bool,
	creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32) {
	creationPtr := GoDateTimeToDDateTime(creation)
	lastAccessPtr := GoDateTimeToDDateTime(lastAccess)
	expiresPtr := GoDateTimeToDDateTime(expires)
	cCookie := &iCefCookiePtr{
		url:             api.PascalStr(url),
		name:            api.PascalStr(name),
		value:           api.PascalStr(value),
		domain:          api.PascalStr(domain),
		path:            api.PascalStr(path),
		secure:          api.PascalBool(secure),
		httponly:        api.PascalBool(httponly),
		hasExpires:      api.PascalBool(hasExpires),
		creation:        uintptr(unsafe.Pointer(&creationPtr)),
		lastAccess:      uintptr(unsafe.Pointer(&lastAccessPtr)),
		expires:         uintptr(unsafe.Pointer(&expiresPtr)),
		sameSite:        uintptr(sameSite),
		priority:        uintptr(priority),
		aSetImmediately: api.PascalBool(aSetImmediately),
		aID:             uintptr(aID),
		aDeleteCookie:   uintptr(0),
		aResult:         uintptr(0),
		count:           uintptr(0),
		total:           uintptr(0),
	}
	imports.Proc(internale_CEFChromium_SetCookie).Call(instance, uintptr(unsafe.Pointer(cCookie)))
	cCookie = nil
}

// TCEFChromium  _CEFChromium_SetProxy
func _CEFChromium_SetProxy(instance uintptr, proxy *tCefProxyPtr) {
	imports.Proc(internale_CEFChromium_SetProxy).Call(instance, uintptr(unsafe.Pointer(proxy)))
}

// TCEFChromium  _CEFChromium_UpdatePreferences
func _CEFChromium_UpdatePreferences(instance uintptr) {
	imports.Proc(internale_CEFChromium_UpdatePreferences).Call(instance)
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
	imports.Proc(internale_CEFChromium_ExecuteDevToolsMethod).Call(instance, uintptr(messageId), api.PascalStr(method), uintptr(argsLen), uintptr(dataPtr), uintptr(dataLen))
}

// TCEFChromium  _CEFChromium_CreateClientHandler
func _CEFChromium_CreateClientHandler(instance, client, alsOSR uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_CreateClientHandler).Call(instance, client, alsOSR)
	return r1
}

func _CEFChromium_SetFocus(instance, value uintptr) {
	imports.Proc(internale_CEFChromium_SetFocus).Call(instance, value)
}

func _CEFChromium_SendCaptureLostEvent(instance uintptr) {
	imports.Proc(internale_CEFChromium_SendCaptureLostEvent).Call(instance)
}

func _CEFChromium_FrameIsFocused(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_FrameIsFocused).Call(instance)
	return r1
}

func _CEFChromium_TryCloseBrowser(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_TryCloseBrowser).Call(instance)
	return r1
}

func _CEFChromium_BrowserHandle(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_BrowserHandle).Call(instance)
	return r1
}

func _CEFChromium_WidgetHandle(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_WidgetHandle).Call(instance)
	return r1
}

func _CEFChromium_RenderHandle(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(internale_CEFChromium_RenderHandle).Call(instance)
	return r1
}

//--------TCEFChromium proc end--------
