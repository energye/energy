//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Chromium 功能函数接口定义

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/cef/ipc/target"
	. "github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/json"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"strings"
	"time"
	"unsafe"
)

// IChromiumProc
// Chromium 功能函数接口
type IChromiumProc interface {
	lcl.IObject
	Instance() uintptr
	ChromiumConfig() *tCefChromiumConfig
	Browser() *ICefBrowser
	BrowserById(id int32) *ICefBrowser
	BrowserIdByIndex(index int32) int32
	BrowserCount() int32
	BrowserId() int32
	SetDefaultURL(defaultURL string)
	SetEnableMultiBrowserMode(enableMultiBrowserMode bool)
	LoadUrl(url string)
	LoadHtml(html string)
	StartDownload(url string)
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool)
	Reload()
	ReloadIgnoreCache()
	StopLoad()
	ResetZoomLevel()
	CloseAllBrowsers()
	CreateBrowser(window ICEFWindowParent, windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool
	CreateBrowserByBrowserViewComponent(homePage string, browserViewComponent *TCEFBrowserViewComponent) bool
	Initialized() bool
	IsSameBrowser(browser *ICefBrowser) bool
	PrintToPDF(saveFilePath string)
	Print()
	BrowserZoom(zoom ZOOM)
	GoBack()
	GoForward()
	NotifyMoveOrResizeStarted()
	CloseBrowser(forceClose bool)
	ExecuteJavaScript(code, scriptURL string, startLine int32)
	ShowDevTools(window ICEFWindowParent)
	CloseDevTools(window ICEFWindowParent)
	VisitAllCookies(id int32)
	VisitURLCookies(url string, includeHttpOnly bool, id int32)
	DeleteCookies(url, cookieName string, deleteImmediately bool)
	SetCookie(url, name, value, domain, path string,
		secure, httponly, hasExpires bool,
		creation, lastAccess, expires time.Time,
		sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32)
	FlushCookieStore(flushImmediately bool) bool // flushImmediately = true
	SetProxy(cefProxy *TCefProxy)
	UpdatePreferences()
	ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue)
	//SendProcessMessage(targetProcess CefProcessId, processMessage *ICefProcessMessage)
	CefClient() *ICefClient
	SendProcessMessageForJSONBytes(name string, targetProcess CefProcessId, message argument.IList)
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
	SendProcessMessageForV8Value(messageName string, targetProcess CefProcessId, arguments *ICefV8Value)
	SimulateMouseWheel(deltaX, deltaY int32)
	CloseAllConnections(closeImmediately bool) bool
	RetrieveHTMLByFrameName(frameName string)
	RetrieveHTMLByFrame(frame *ICefFrame)
	RetrieveHTMLByFrameIdentifier(frameIdentifier int64)
	RetrieveTextByFrameName(frameName string)
	RetrieveTextByFrame(frame *ICefFrame)
	RetrieveTextByFrameIdentifier(frameIdentifier int64)
	ResolveHost(url string)
	SetUserAgentOverride(userAgent, acceptLanguage, platform string)
	ClearDataForOrigin(origin string, storageTypes ...TCefClearDataStorageTypes)
	ClearCache()
	ToggleAudioMuted()
	ClipboardCopy()
	ClipboardPaste()
	ClipboardCut()
	ClipboardUndo()
	ClipboardRedo()
	ClipboardDel()
	SelectAll()
	IncZoomStep()
	DecZoomStep()
	IncZoomPct()
	DecZoomPct()
	ResetZoomStep()
	ResetZoomPct()
	ReadZoom()
	WasResized()
	WasHidden(hidden bool)
	NotifyScreenInfoChanged()
	IMESetComposition(text string, underlines []*TCefCompositionUnderline, replacementRange, selectionRange *TCefRange)
	IMECommitText(text string, replacementRange *TCefRange, relativeCursorPos int32)
	IMEFinishComposingText(keepSelection bool)
	IMECancelComposition()
}

// IsValid 实例有效
func (m *TCEFChromium) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// UnsafeAddr 实例指针
func (m *TCEFChromium) UnsafeAddr() unsafe.Pointer {
	return m.instance
}

func (m *TCEFChromium) ClassName() string {
	r1, _, _ := imports.Proc(def.CEFChromium_ClassName).Call()
	return api.GoStr(r1)
}

func (m *TCEFChromium) Free() {
	if m.idBrowsers != nil {
		for _, browse := range m.idBrowsers {
			browse.Free()
		}
		m.idBrowsers = nil
	}
	imports.Proc(def.CEFChromium_Free).Call()
}

func (m *TCEFChromium) HashCode() int32 {
	r1, _, _ := imports.Proc(def.CEFChromium_GetHashCode).Call()
	return int32(r1)
}

func (m *TCEFChromium) Equals(object lcl.IObject) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_Equals).Call(lcl.CheckPtr(object))
	return api.GoBool(r1)
}

func (m *TCEFChromium) ClassType() types.TClass {
	r1, _, _ := imports.Proc(def.CEFChromium_ClassType).Call()
	return types.TClass(r1)
}

func (m *TCEFChromium) InstanceSize() int32 {
	r1, _, _ := imports.Proc(def.CEFChromium_InstanceSize).Call()
	return int32(r1)
}

func (m *TCEFChromium) InheritsFrom(class types.TClass) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_InheritsFrom).Call(uintptr(class))
	return api.GoBool(r1)
}

func (m *TCEFChromium) ToString() string {
	r1, _, _ := imports.Proc(def.CEFChromium_ToString).Call()
	return api.GoStr(r1)
}

func (m *TCEFChromium) ChromiumConfig() *tCefChromiumConfig {
	return m.cfg
}

func (m *TCEFChromium) Browser() *ICefBrowser {
	if !m.IsValid() {
		return nil
	}
	if m.browser != nil && m.browser.instance != nil && m.browser.IsValid() {
		return m.browser
	}
	var result uintptr
	imports.Proc(def.CEFChromium_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	m.browser = &ICefBrowser{instance: unsafe.Pointer(result)}
	return m.browser
}

func (m *TCEFChromium) BrowserById(id int32) *ICefBrowser {
	if !m.IsValid() {
		return nil
	}
	if m.idBrowsers == nil {
		m.idBrowsers = make(map[int32]*ICefBrowser)
	}
	if browse, ok := m.idBrowsers[id]; ok {
		if browse.instance != nil && browse.IsValid() {
			return browse
		}
		delete(m.idBrowsers, id)
	}
	var result uintptr
	imports.Proc(def.CEFChromium_BrowserById).Call(m.Instance(), uintptr(id), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		browse := &ICefBrowser{instance: unsafe.Pointer(result)}
		m.idBrowsers[id] = browse
		return browse
	}
	return nil
}

func (m *TCEFChromium) BrowserIdByIndex(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_BrowserIdByIndex).Call(m.Instance(), uintptr(index))
	return int32(r1)
}

func (m *TCEFChromium) BrowserCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_BrowserCount).Call(m.Instance())
	return int32(r1)
}

func (m *TCEFChromium) BrowserId() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_BrowserId).Call(m.Instance())
	return int32(r1)
}

// SetDefaultURL 设置默认地址
func (m *TCEFChromium) SetDefaultURL(defaultURL string) {
	if !m.IsValid() {
		return
	}
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

// SetEnableMultiBrowserMode 设置启用多浏览器模式
func (m *TCEFChromium) SetEnableMultiBrowserMode(enableMultiBrowserMode bool) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetMultiBrowserMode(m.Instance(), enableMultiBrowserMode)
}

// LoadUrl 加载一个URL地址
func (m *TCEFChromium) LoadUrl(url string) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_LoadURL(m.Instance(), url)
}

// LoadHtml 加载HTML
func (m *TCEFChromium) LoadHtml(html string) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_LoadString(m.Instance(), html)
}

// StartDownload 开始下载
func (m *TCEFChromium) StartDownload(url string) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_StartDownload(m.Instance(), url)
}

// DownloadImage 开始下载图片
func (m *TCEFChromium) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_DownloadImage(m.Instance(), imageUrl, isFavicon, maxImageSize, bypassCache)
}

func (m *TCEFChromium) Reload() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_Reload(m.Instance())
}

func (m *TCEFChromium) ReloadIgnoreCache() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_ReloadIgnoreCache(m.Instance())
}

func (m *TCEFChromium) StopLoad() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_StopLoad(m.Instance())
}

func (m *TCEFChromium) ResetZoomLevel() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_ResetZoomLevel(m.Instance())
}

func (m *TCEFChromium) CloseAllBrowsers() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_CloseAllBrowses(m.Instance())
}

func (m *TCEFChromium) CreateBrowser(window ICEFWindowParent, windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	var (
		windowParent  uintptr
		windowNamePtr uintptr
		contextPtr    uintptr
		extraInfoPtr  uintptr
	)
	if window != nil {
		windowParent = m.Instance()
	}
	if windowName != "" {
		windowNamePtr = api.PascalStr(windowName)
	}
	if context != nil && context.IsValid() {
		contextPtr = context.Instance()
	}
	if extraInfo != nil && extraInfo.IsValid() {
		extraInfoPtr = extraInfo.Instance()
	}
	return _CEFChromium_CreateBrowse(m.Instance(), windowParent, windowNamePtr, contextPtr, extraInfoPtr)
	//if window == nil {
	//	return _CEFChromium_CreateBrowse(m.Instance(), windowParent, windowNamePtr, contextPtr, extraInfoPtr)
	//} else {
	//	if window.Type() == Wht_WindowParent {
	//		return _CEFChromium_CreateBrowseByWindow(m.Instance(), window.Instance())
	//	} else if window.Type() == Wht_LinkedWindowParent {
	//		return _CEFChromium_CreateBrowseByLinkedWindow(m.Instance(), window.Instance())
	//	}
	//}
	//return false
}

func (m *TCEFChromium) CreateBrowserByBrowserViewComponent(homePage string, browserViewComponent *TCEFBrowserViewComponent) bool {
	if !m.IsValid() {
		return false
	}
	return _CEFChromium_CreateBrowserByBrowserViewComponent(m.Instance(), api.PascalStr(homePage), browserViewComponent.Instance())
}

func (m *TCEFChromium) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	if m.initialized {
		return true
	}
	m.initialized = _CEFChromium_Initialized(m.Instance())
	return m.initialized
}

func (m *TCEFChromium) IsSameBrowser(browser *ICefBrowser) bool {
	if !m.IsValid() {
		return false
	}
	return _CEFChromium_IsSameBrowser(m.Instance(), browser.Instance())
}

func (m *TCEFChromium) PrintToPDF(saveFilePath string) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_PrintToPDF(m.Instance(), saveFilePath)
}

func (m *TCEFChromium) Print() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_Print(m.Instance())
}

func (m *TCEFChromium) BrowserZoom(zoom ZOOM) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_BrowserZoom(m.Instance(), zoom)
}

func (m *TCEFChromium) GoBack() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_GoBackForward(m.Instance(), BF_GOBACK)
}

func (m *TCEFChromium) GoForward() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_GoBackForward(m.Instance(), BF_GOFORWARD)
}

func (m *TCEFChromium) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_NotifyMoveOrResizeStarted(m.Instance())
}

func (m *TCEFChromium) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_CloseBrowser(m.Instance(), api.PascalBool(forceClose))
}

func (m *TCEFChromium) ShowDevTools(window ICEFWindowParent) {
	if !m.IsValid() {
		return
	}
	if window == nil {
		_CEFChromium_ShowDevTools(m.Instance())
	} else {
		_CEFChromium_ShowDevToolsByWindowParent(m.Instance(), window.Instance())
	}
}

func (m *TCEFChromium) CloseDevTools(window ICEFWindowParent) {
	if !m.IsValid() {
		return
	}
	if window == nil {
		_CEFChromium_CloseDevTools(m.Instance())
	} else {
		_CEFChromium_CloseDevToolsByWindowParent(m.Instance(), window.Instance())
	}
}

// 查看所有cookie,该函数触发 OnCookiesVisited 事件返回结果
func (m *TCEFChromium) VisitAllCookies(id int32) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_VisitAllCookies(m.Instance(), id)
}

// 查看指针URL cookie,该函数触发 OnCookiesVisited 事件返回结果
// url https://www.demo.com
func (m *TCEFChromium) VisitURLCookies(url string, includeHttpOnly bool, id int32) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_VisitURLCookies(m.Instance(), url, includeHttpOnly, id)
}

// 删除所有cookie
func (m *TCEFChromium) DeleteCookies(url, cookieName string, deleteImmediately bool) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_DeleteCookies(m.Instance(), url, cookieName, deleteImmediately)
}

func (m *TCEFChromium) SetCookie(url, name, value, domain, path string,
	secure, httponly, hasExpires bool,
	creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetCookie(m.Instance(), url, name, value, domain, path, secure, httponly, hasExpires, creation, lastAccess, expires, sameSite, priority, aSetImmediately, aID)
}

func (m *TCEFChromium) FlushCookieStore(flushImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	return _CEFChromium_FlushCookieStore(m.Instance(), flushImmediately)
}

func (m *TCEFChromium) SetProxy(cefProxy *TCefProxy) {
	if !m.IsValid() {
		return
	}
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
	if !m.IsValid() {
		return
	}
	_CEFChromium_UpdatePreferences(m.Instance())
}

func (m *TCEFChromium) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	if !m.IsValid() {
		return
	}
	if dictionaryValue == nil {
		dictionaryValue = DictionaryValueRef.New()
	}
	_CEFChromium_ExecuteDevToolsMethod(m.Instance(), messageId, method, dictionaryValue)
}

func (m *TCEFChromium) CreateClientHandler(client *ICefClient, alsOSR bool) bool {
	if !m.IsValid() {
		return false
	}
	return api.GoBool(_CEFChromium_CreateClientHandler(m.Instance(), client.Instance(), api.PascalBool(alsOSR)))
}

func (m *TCEFChromium) SetFocus(value bool) {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SetFocus(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	_CEFChromium_SendCaptureLostEvent(m.Instance())
}

func (m *TCEFChromium) FrameIsFocused() bool {
	if !m.IsValid() {
		return false
	}
	return api.GoBool(_CEFChromium_FrameIsFocused(m.Instance()))
}

func (m *TCEFChromium) TryCloseBrowser() bool {
	if !m.IsValid() {
		return false
	}
	return api.GoBool(_CEFChromium_TryCloseBrowser(m.Instance()))
}

func (m *TCEFChromium) BrowserHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	if m.browserHandle == 0 {
		m.browserHandle = types.HWND(_CEFChromium_BrowserHandle(m.Instance()))
	}
	return m.browserHandle
}

func (m *TCEFChromium) WidgetHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	if m.widgetHandle == 0 {
		m.widgetHandle = types.HWND(_CEFChromium_WidgetHandle(m.Instance()))
	}
	return m.widgetHandle
}

func (m *TCEFChromium) RenderHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	if m.renderHandle == 0 {
		m.renderHandle = types.HWND(_CEFChromium_RenderHandle(m.Instance()))
	}
	return m.renderHandle
}

func (m *TCEFChromium) SetCustomHeader(customHeader *TCustomHeader) {
	if !m.IsValid() {
		return
	}
	ptrCustomHeader := &tCustomHeader{
		CustomHeaderName:  api.PascalStr(customHeader.CustomHeaderName),
		CustomHeaderValue: api.PascalStr(customHeader.CustomHeaderValue),
	}
	imports.Proc(def.CEFChromium_SetCustomHeader).Call(m.Instance(), uintptr(unsafe.Pointer(ptrCustomHeader)))
}

func (m *TCEFChromium) CustomHeader() *TCustomHeader {
	if !m.IsValid() {
		return nil
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetCustomHeader).Call(m.Instance())
	ptrCustomHeader := (*tCustomHeader)(unsafe.Pointer(r1))
	return &TCustomHeader{
		CustomHeaderName:  api.GoStr(ptrCustomHeader.CustomHeaderName),
		CustomHeaderValue: api.GoStr(ptrCustomHeader.CustomHeaderValue),
	}
}

func (m *TCEFChromium) SetJavascriptEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetJavascriptEnabled).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) JavascriptEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetJavascriptEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetWebRTCIPHandlingPolicy(value TCefWebRTCHandlingPolicy) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetWebRTCIPHandlingPolicy).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) WebRTCIPHandlingPolicy() TCefWebRTCHandlingPolicy {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetWebRTCIPHandlingPolicy).Call(m.Instance())
	return TCefWebRTCHandlingPolicy(r1)
}

func (m *TCEFChromium) SetWebRTCMultipleRoutes(value TCefState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetWebRTCMultipleRoutes).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) WebRTCMultipleRoutes() TCefState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetWebRTCMultipleRoutes).Call(m.Instance())
	return TCefState(r1)
}

func (m *TCEFChromium) SetWebRTCNonproxiedUDP(value TCefState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetWebRTCNonproxiedUDP).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) WebRTCNonproxiedUDP() TCefState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetWebRTCNonproxiedUDP).Call(m.Instance())
	return TCefState(r1)
}

func (m *TCEFChromium) SetBatterySaverModeState(value TCefBatterySaverModeState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetBatterySaverModeState).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) BatterySaverModeState() TCefBatterySaverModeState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetBatterySaverModeState).Call(m.Instance())
	return TCefBatterySaverModeState(r1)
}

func (m *TCEFChromium) SetHighEfficiencyMode(value TCefState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetHighEfficiencyMode).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) HighEfficiencyMode() TCefState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetHighEfficiencyMode).Call(m.Instance())
	return TCefState(r1)
}

func (m *TCEFChromium) SetLoadImagesAutomatically(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetLoadImagesAutomatically).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) LoadImagesAutomatically() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetLoadImagesAutomatically).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetQuicAllowed(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetQuicAllowed).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) QuicAllowed() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetQuicAllowed).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetOffline(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetOffline).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) Offline() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetOffline).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetDefaultWindowInfoExStyle(exStyle types.DWORD) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetDefaultWindowInfoExStyle).Call(m.Instance(), uintptr(exStyle))
}

func (m *TCEFChromium) DefaultWindowInfoExStyle() types.DWORD {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetDefaultWindowInfoExStyle).Call(m.Instance())
	return types.DWORD(r1)
}

func (m *TCEFChromium) SetBlock3rdPartyCookies(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetBlock3rdPartyCookies).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) Block3rdPartyCookies() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetBlock3rdPartyCookies).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAcceptCookies(cp TCefCookiePref) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAcceptCookies).Call(m.Instance(), cp.ToPtr())
}

func (m *TCEFChromium) AcceptCookies() TCefCookiePref {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetAcceptCookies).Call(m.Instance())
	return TCefCookiePref(r1)
}

func (m *TCEFChromium) SetAcceptLanguageList(languageList string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAcceptLanguageList).Call(m.Instance(), api.PascalStr(languageList))
}

func (m *TCEFChromium) AcceptLanguageList() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetAcceptLanguageList).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFChromium) SetPrintingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetPrintingEnabled).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) PrintingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetPrintingEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetYouTubeRestrict(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetYouTubeRestrict).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) YouTubeRestrict() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetYouTubeRestrict).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetSafeSearch(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetSafeSearch).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SafeSearch() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetSafeSearch).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAudioMuted(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAudioMuted).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AudioMuted() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetAudioMuted).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetDragOperations(value TCefDragOperations) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetDragOperations).Call(m.Instance(), value.ToPtr())
}

func (m *TCEFChromium) DragOperations() TCefDragOperations {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetDragOperations).Call(m.Instance())
	return TCefDragOperations(r1)
}

func (m *TCEFChromium) FrameCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetFrameCount).Call(m.Instance())
	return uint32(r1)
}

func (m *TCEFChromium) SetSpellCheckerDicts(value string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetSpellCheckerDicts).Call(m.Instance(), api.PascalStr(value))
}

func (m *TCEFChromium) SpellCheckerDicts() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetSpellCheckerDicts).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFChromium) SetSpellChecking(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetSpellChecking).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SpellChecking() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetSpellChecking).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAlwaysOpenPDFExternally(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAlwaysOpenPDFExternally).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AlwaysOpenPDFExternally() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetAlwaysOpenPDFExternally).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAlwaysAuthorizePlugins(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAlwaysAuthorizePlugins).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AlwaysAuthorizePlugins() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetAlwaysAuthorizePlugins).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetAllowOutdatedPlugins(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAllowOutdatedPlugins).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) AllowOutdatedPlugins() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetAllowOutdatedPlugins).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetSendReferrer(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetSendReferrer).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SendReferrer() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetSendReferrer).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetDoNotTrack(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetDoNotTrack).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) DoNotTrack() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetDoNotTrack).Call(m.Instance())
	return api.GoBool(r1)
}

// SetZoomStep 设置缩放步 0~255
func (m *TCEFChromium) SetZoomStep(value int8) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetZoomStep).Call(m.Instance(), uintptr(value))
}

// ZoomStep 获取缩放步 0~255
func (m *TCEFChromium) ZoomStep() int8 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetZoomStep).Call(m.Instance())
	return int8(r1)
}

// SetZoomPct 设置缩放百分比
func (m *TCEFChromium) SetZoomPct(value float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetZoomPct).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}

// ZoomPct 获取缩放百分比
func (m *TCEFChromium) ZoomPct() (result float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_GetZoomPct).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// SetZoomLevel 设置缩放级别
func (m *TCEFChromium) SetZoomLevel(value float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}

// ZoomLevel 获取缩放级别
func (m *TCEFChromium) ZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_GetZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// SetDefaultEncoding 设置默认编码
func (m *TCEFChromium) SetDefaultEncoding(value string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetDefaultEncoding).Call(m.Instance(), api.PascalStr(value))
}

// DefaultEncoding 获取默认编码
func (m *TCEFChromium) DefaultEncoding() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetDefaultEncoding).Call(m.Instance())
	return api.GoStr(r1)
}

// SendProcessMessage 发送进程消息
func (m *TCEFChromium) SendProcessMessage(targetProcess CefProcessId, message *ICefProcessMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendProcessMessage).Call(m.Instance(), targetProcess.ToPtr(), message.Instance())
	message.Free()
}

// SendProcessMessageForJSONBytes 发送进程消息
func (m *TCEFChromium) SendProcessMessageForJSONBytes(name string, targetProcess CefProcessId, message argument.IList) {
	if !m.IsValid() {
		return
	}
	if !m.initialized {
		m.initialized = m.Initialized()
		if !m.initialized {
			return
		}
	}
	var data = message.Bytes()
	imports.Proc(def.CEFChromium_SendProcessMessageForJSONBytes).Call(m.Instance(), api.PascalStr(name), targetProcess.ToPtr(), uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))))
}

// CefClient 获取Client
func (m *TCEFChromium) CefClient() *ICefClient {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFChromium_CefClient).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefClient{instance: unsafe.Pointer(result)}
	}
	return nil
}

// SendProcessMessageForV8Value 发送进程消息
func (m *TCEFChromium) SendProcessMessageForV8Value(messageName string, targetProcess CefProcessId, arguments *ICefV8Value) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendProcessMessageForV8Value).Call(m.Instance(), api.PascalStr(messageName), targetProcess.ToPtr(), arguments.Instance())
}

func (m *TCEFChromium) SimulateMouseWheel(deltaX, deltaY int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SimulateMouseWheel).Call(m.Instance(), uintptr(deltaX), uintptr(deltaY))
}

func (m *TCEFChromium) CloseAllConnections(closeImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_CloseAllConnections).Call(m.Instance(), api.PascalBool(closeImmediately))
	return api.GoBool(r1)
}

func (m *TCEFChromium) RetrieveHTMLByFrameName(frameName string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_RetrieveHTMLByFrameName).Call(m.Instance(), api.PascalStr(frameName))
}

func (m *TCEFChromium) RetrieveHTMLByFrame(frame *ICefFrame) {
	if !m.IsValid() || !frame.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_RetrieveHTMLByFrame).Call(m.Instance(), frame.Instance())
}

func (m *TCEFChromium) RetrieveHTMLByFrameIdentifier(frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_RetrieveHTMLByFrameIdentifier).Call(m.Instance(), uintptr(unsafe.Pointer(&frameIdentifier)))
}

func (m *TCEFChromium) RetrieveTextByFrameName(frameName string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_RetrieveTextByFrameName).Call(m.Instance(), api.PascalStr(frameName))
}

func (m *TCEFChromium) RetrieveTextByFrame(frame *ICefFrame) {
	if !m.IsValid() || !frame.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_RetrieveTextByFrame).Call(m.Instance(), frame.Instance())
}

func (m *TCEFChromium) RetrieveTextByFrameIdentifier(frameIdentifier int64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_RetrieveTextByFrameIdentifier).Call(m.Instance(), uintptr(unsafe.Pointer(&frameIdentifier)))
}

func (m *TCEFChromium) ResolveHost(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ResolveHost).Call(m.Instance(), api.PascalStr(url))
}

func (m *TCEFChromium) SetUserAgentOverride(userAgent, acceptLanguage, platform string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetUserAgentOverride).Call(m.Instance(), api.PascalStr(userAgent), api.PascalStr(acceptLanguage), api.PascalStr(platform))
}

func (m *TCEFChromium) ClearDataForOrigin(origin string, storageTypes ...TCefClearDataStorageTypes) {
	if !m.IsValid() {
		return
	}
	var st = CdstAll // default
	if len(storageTypes) > 0 {
		st = storageTypes[0]
	}
	imports.Proc(def.CEFChromium_ClearDataForOrigin).Call(m.Instance(), api.PascalStr(origin), st.ToPtr())
}

func (m *TCEFChromium) ClearCache() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClearCache).Call(m.Instance())
}

func (m *TCEFChromium) ToggleAudioMuted() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ToggleAudioMuted).Call(m.Instance())
}

func (m *TCEFChromium) ClipboardCopy() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClipboardCopy).Call(m.Instance())
}

func (m *TCEFChromium) ClipboardPaste() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClipboardPaste).Call(m.Instance())
}

func (m *TCEFChromium) ClipboardCut() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClipboardCut).Call(m.Instance())
}

func (m *TCEFChromium) ClipboardUndo() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClipboardUndo).Call(m.Instance())
}

func (m *TCEFChromium) ClipboardRedo() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClipboardRedo).Call(m.Instance())
}

func (m *TCEFChromium) ClipboardDel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ClipboardDel).Call(m.Instance())
}

func (m *TCEFChromium) SelectAll() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SelectAll).Call(m.Instance())
}

func (m *TCEFChromium) IncZoomStep() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_IncZoomStep).Call(m.Instance())
}

func (m *TCEFChromium) DecZoomStep() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_DecZoomStep).Call(m.Instance())
}

func (m *TCEFChromium) IncZoomPct() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_IncZoomPct).Call(m.Instance())
}

func (m *TCEFChromium) DecZoomPct() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_DecZoomPct).Call(m.Instance())
}

func (m *TCEFChromium) ResetZoomStep() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ResetZoomStep).Call(m.Instance())
}

func (m *TCEFChromium) ResetZoomPct() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ResetZoomPct).Call(m.Instance())
}

func (m *TCEFChromium) ReadZoom() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ReadZoom).Call(m.Instance())
}

func (m *TCEFChromium) WasResized() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_WasResized).Call(m.Instance())
}

func (m *TCEFChromium) WasHidden(hidden bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_WasHidden).Call(m.Instance(), api.PascalBool(hidden))
}

func (m *TCEFChromium) NotifyScreenInfoChanged() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_NotifyScreenInfoChanged).Call(m.Instance())
}

func (m *TCEFChromium) IMESetComposition(text string, underlines []*TCefCompositionUnderline, replacementRange, selectionRange *TCefRange) {
	if !m.IsValid() {
		return
	}
	var size = len(underlines)
	// 在这里使用结构指针的方式将复合类型传递过去
	underlinesPtr := make([]tCefCompositionUnderlinePtr, size, size)
	for i := 0; i < size; i++ {
		underlinesPtr[i] = tCefCompositionUnderlinePtr{
			Range:           uintptr(unsafe.Pointer(&underlines[i].Range)),
			Color:           uintptr(underlines[i].Color),
			BackgroundColor: uintptr(underlines[i].BackgroundColor),
			Thick:           uintptr(underlines[i].Thick),
			Style:           uintptr(underlines[i].Style),
		}
	}
	imports.Proc(def.CEFChromium_IMESetComposition).Call(m.Instance(), api.PascalStr(text), uintptr(unsafe.Pointer(replacementRange)), uintptr(unsafe.Pointer(selectionRange)),
		uintptr(unsafe.Pointer(&underlinesPtr[0])), uintptr(int32(size)))
}

func (m *TCEFChromium) IMECommitText(text string, replacementRange *TCefRange, relativeCursorPos int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_IMECommitText).Call(m.Instance(), api.PascalStr(text), uintptr(unsafe.Pointer(replacementRange)), uintptr(relativeCursorPos))
}

func (m *TCEFChromium) IMEFinishComposingText(keepSelection bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_IMEFinishComposingText).Call(m.Instance(), api.PascalBool(keepSelection))
}

func (m *TCEFChromium) IMECancelComposition() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_IMECancelComposition).Call(m.Instance())
}

// EmitRender IPC 发送进程 消息
//
// messageId != 0 是带有回调函数消息
func (m *TCEFChromium) EmitRender(messageId int32, eventName string, target target.ITarget, data ...any) bool {
	if !m.initialized {
		m.initialized = m.Initialized()
		if !m.initialized {
			return false
		}
	}
	if target == nil || target.BrowserId() <= 0 || target.ChannelId() <= 0 {
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
		m.SendProcessMessageForJSONBytes(internalIPCGoExecuteJSEvent, PID_RENDER, message)
		message.Reset()
		return true
	} else {
		browse := BrowserWindow.GetBrowser(target.BrowserId())
		//browse := m.BrowserById(target.BrowserId())
		if browse != nil && browse.IsValid() {
			frame := browse.GetFrameById(target.ChannelId())
			if frame != nil && frame.IsValid() {
				return frame.EmitRender(messageId, eventName, target, data...)
			}
		}
	}
	return false
}

//--------TCEFChromium proc begin--------

// TCEFChromium _CEFChromium_Create
func _CEFChromium_Create(owner, config uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_Create).Call(owner, config)
	return r1
}

// TCEFChromium _CEFChromium_SetDefaultURL
func _CEFChromium_SetDefaultURL(instance uintptr, url string) {
	imports.Proc(def.CEFChromium_SetDefaultURL).Call(instance, api.PascalStr(url))
}

// TCEFChromium _CEFChromium_SetMultiBrowserMode
func _CEFChromium_SetMultiBrowserMode(instance uintptr, url bool) {
	imports.Proc(def.CEFChromium_SetMultiBrowserMode).Call(instance, api.PascalBool(url))
}

// TCEFChromium _CEFChromium_LoadURL
func _CEFChromium_LoadURL(instance uintptr, url string) {
	imports.Proc(def.CEFChromium_LoadURL).Call(instance, api.PascalStr(url))
}

// TCEFChromium _CEFChromium_LoadString
func _CEFChromium_LoadString(instance uintptr, html string) {
	imports.Proc(def.CEFChromium_LoadString).Call(instance, api.PascalStr(html))
}

// TCEFChromium _CEFChromium_StartDownload
func _CEFChromium_StartDownload(instance uintptr, url string) {
	imports.Proc(def.CEFChromium_StartDownload).Call(instance, api.PascalStr(url))
}

// TCEFChromium _CEFChromium_DownloadImage
func _CEFChromium_DownloadImage(instance uintptr, imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	imports.Proc(def.CEFChromium_DownloadImage).Call(instance, api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache))
}

// TCEFChromium _CEFChromium_Reload
func _CEFChromium_Reload(instance uintptr) {
	imports.Proc(def.CEFChromium_Reload).Call(instance)
}

// TCEFChromium _CEFChromium_ReloadIgnoreCache
func _CEFChromium_ReloadIgnoreCache(instance uintptr) {
	imports.Proc(def.CEFChromium_ReloadIgnoreCache).Call(instance)
}

// TCEFChromium _CEFChromium_StopLoad
func _CEFChromium_StopLoad(instance uintptr) {
	imports.Proc(def.CEFChromium_StopLoad).Call(instance)
}

// TCEFChromium _CEFChromium_ResetZoomLevel
func _CEFChromium_ResetZoomLevel(instance uintptr) {
	imports.Proc(def.CEFChromium_ResetZoomLevel).Call(instance)
}

// TCEFChromium _CEFChromium_CloseAllBrowses
func _CEFChromium_CloseAllBrowses(instance uintptr) {
	imports.Proc(def.CEFChromium_CloseAllBrowsers).Call(instance)
}

// TCEFChromium _CEFChromium_CreateBrowse
func _CEFChromium_CreateBrowse(instance, window, windowName, context, extraInfo uintptr) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowser).Call(instance, window, windowName, context, extraInfo)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_CreateBrowseByWindow
func _CEFChromium_CreateBrowseByWindow(instance, window uintptr) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowserByWindow).Call(instance, window)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_CreateBrowseByLinkedWindow
func _CEFChromium_CreateBrowseByLinkedWindow(instance, window uintptr) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowserByLinkedWindow).Call(instance, window)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_CreateBrowserByBrowserViewComponent
func _CEFChromium_CreateBrowserByBrowserViewComponent(instance, homePage, browserViewComponent uintptr) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowserByBrowserViewComponent).Call(instance, homePage, browserViewComponent)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_Initialized
func _CEFChromium_Initialized(instance uintptr) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_Initialized).Call(instance)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_IsSameBrowser
func _CEFChromium_IsSameBrowser(instance, browser uintptr) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_IsSameBrowser).Call(instance, browser)
	return api.GoBool(r1)
}

// TCEFChromium _CEFChromium_PrintToPDF
func _CEFChromium_PrintToPDF(instance uintptr, saveFilePath string) {
	imports.Proc(def.CEFChromium_PrintToPDF).Call(instance, api.PascalStr(saveFilePath))
}

// TCEFChromium _CEFChromium_Print
func _CEFChromium_Print(instance uintptr) {
	imports.Proc(def.CEFChromium_Print).Call(instance)
}

// TCEFChromium _CEFChromium_BrowserZoom
func _CEFChromium_BrowserZoom(instance uintptr, zoom ZOOM) {
	imports.Proc(def.CEFChromium_BrowserZoom).Call(instance, uintptr(zoom))
}

// TCEFChromium _CEFChromium_GoBackForward
func _CEFChromium_GoBackForward(instance uintptr, bf BF) {
	imports.Proc(def.CEFChromium_GoBackForward).Call(instance, uintptr(bf))
}

// TCEFChromium _CEFChromium_NotifyMoveOrResizeStarted
func _CEFChromium_NotifyMoveOrResizeStarted(instance uintptr) {
	imports.Proc(def.CEFChromium_NotifyMoveOrResizeStarted).Call(instance)
}

// TCEFChromium _CEFChromium_CloseBrowser
func _CEFChromium_CloseBrowser(instance, forceClose uintptr) {
	imports.Proc(def.CEFChromium_CloseBrowser).Call(instance, forceClose)
}

// TCEFChromium _CEFChromium_ExecuteJavaScript
func _CEFChromium_ExecuteJavaScript(instance uintptr, code, scriptURL string, startLine int32) {
	imports.Proc(def.CEFChromium_ExecuteJavaScript).Call(instance, api.PascalStr(code), api.PascalStr(scriptURL), uintptr(startLine))
}

// TCEFChromium _CEFChromium_ShowDevTools
func _CEFChromium_ShowDevTools(instance uintptr) {
	imports.Proc(def.CEFChromium_ShowDevTools).Call(instance)
}
func _CEFChromium_ShowDevToolsByWindowParent(instance, windowParent uintptr) {
	imports.Proc(def.CEFChromium_ShowDevToolsByWindowParent).Call(instance, windowParent)
}

// TCEFChromium _CEFChromium_CloseDevTools
func _CEFChromium_CloseDevTools(instance uintptr) {
	imports.Proc(def.CEFChromium_CloseDevTools).Call(instance)
}
func _CEFChromium_CloseDevToolsByWindowParent(instance, windowParent uintptr) {
	imports.Proc(def.CEFChromium_CloseDevToolsByWindowParent).Call(instance, windowParent)
}

// TCEFChromium _CEFChromium_VisitAllCookies
func _CEFChromium_VisitAllCookies(instance uintptr, id int32) {
	imports.Proc(def.CEFChromium_VisitAllCookies).Call(instance, uintptr(id))
}

// TCEFChromium _CEFChromium_VisitURLCookies
func _CEFChromium_VisitURLCookies(instance uintptr, url string, includeHttpOnly bool, id int32) {
	imports.Proc(def.CEFChromium_VisitURLCookies).Call(instance, api.PascalStr(url), api.PascalBool(includeHttpOnly), uintptr(id))
}

// TCEFChromium _CEFChromium_DeleteCookies
func _CEFChromium_DeleteCookies(instance uintptr, url, cookieName string, deleteImmediately bool) {
	imports.Proc(def.CEFChromium_DeleteCookies).Call(instance, api.PascalStr(url), api.PascalStr(cookieName), api.PascalBool(deleteImmediately))
}

// TCEFChromium _CEFChromium_SetCookie
func _CEFChromium_SetCookie(instance uintptr, url, name, value, domain, path string,
	secure, httponly, hasExpires bool, creation, lastAccess, expires time.Time,
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
	imports.Proc(def.CEFChromium_SetCookie).Call(instance, uintptr(unsafe.Pointer(cCookie)))
	cCookie = nil
}

// TCEFChromium  _CEFChromium_FlushCookieStore
func _CEFChromium_FlushCookieStore(instance uintptr, flushImmediately bool) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_FlushCookieStore).Call(instance, api.PascalBool(flushImmediately))
	return api.GoBool(r1)
}

// TCEFChromium  _CEFChromium_SetProxy
func _CEFChromium_SetProxy(instance uintptr, proxy *tCefProxyPtr) {
	imports.Proc(def.CEFChromium_SetProxy).Call(instance, uintptr(unsafe.Pointer(proxy)))
}

// TCEFChromium  _CEFChromium_UpdatePreferences
func _CEFChromium_UpdatePreferences(instance uintptr) {
	imports.Proc(def.CEFChromium_UpdatePreferences).Call(instance)
}

// TCEFChromium  _CEFChromium_ExecuteDevToolsMethod
func _CEFChromium_ExecuteDevToolsMethod(instance uintptr, messageId int32, method string, dictionaryValue *ICefDictionaryValue) {
	imports.Proc(def.CEFChromium_ExecuteDevToolsMethod).Call(instance, uintptr(messageId), api.PascalStr(method), dictionaryValue.Instance())
}

// TCEFChromium _CEFChromium_CreateClientHandler
func _CEFChromium_CreateClientHandler(instance, client, alsOSR uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_CreateClientHandler).Call(instance, client, alsOSR)
	return r1
}

// TCEFChromium _CEFChromium_SetFocus
func _CEFChromium_SetFocus(instance, value uintptr) {
	imports.Proc(def.CEFChromium_SetFocus).Call(instance, value)
}

// TCEFChromium _CEFChromium_SendCaptureLostEvent
func _CEFChromium_SendCaptureLostEvent(instance uintptr) {
	imports.Proc(def.CEFChromium_SendCaptureLostEvent).Call(instance)
}

// TCEFChromium _CEFChromium_FrameIsFocused
func _CEFChromium_FrameIsFocused(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_FrameIsFocused).Call(instance)
	return r1
}

// TCEFChromium _CEFChromium_TryCloseBrowser
func _CEFChromium_TryCloseBrowser(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_TryCloseBrowser).Call(instance)
	return r1
}

// TCEFChromium _CEFChromium_BrowserHandle
func _CEFChromium_BrowserHandle(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_BrowserHandle).Call(instance)
	return r1
}

// TCEFChromium _CEFChromium_WidgetHandle
func _CEFChromium_WidgetHandle(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_WidgetHandle).Call(instance)
	return r1
}

// TCEFChromium _CEFChromium_RenderHandle
func _CEFChromium_RenderHandle(instance uintptr) uintptr {
	r1, _, _ := imports.Proc(def.CEFChromium_RenderHandle).Call(instance)
	return r1
}

//--------TCEFChromium proc end--------
