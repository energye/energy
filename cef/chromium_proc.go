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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/cef/ipc/argument"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	. "github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
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
	Options() *TChromiumOptions
	FontOptions() *TChromiumFontOptions
	Config() *TCefChromiumConfig
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
	CreateBrowserByWinControl(browserParent *lcl.TWinControl, windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool
	CreateBrowserByWindowHandle(parentHandle TCefWindowHandle, rect types.TRect, windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue, forceAsPopup bool) bool
	CreateBrowserByBrowserViewComponent(homePage string, browserViewComponent *TCEFBrowserViewComponent, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool
	Initialized() bool
	IsSameBrowser(browser *ICefBrowser) bool
	PrintToPDF(saveFilePath string)
	Print()
	BrowserZoom(zoom ZOOM)
	GoBack()
	GoForward()
	NotifyMoveOrResizeStarted()
	Invalidate(type_ TCefPaintElementType)
	CloseBrowser(forceClose bool)
	ExecuteJavaScript(code, scriptURL string, frame *ICefFrame, startLine int32)
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
	SetProxy(cefProxy TCefProxy)
	UpdatePreferences()
	SendDevToolsMessage(message string) bool
	ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) int32
	//SendProcessMessage(targetProcess CefProcessId, processMessage *ICefProcessMessage)
	Client() *ICefClient
	SendProcessMessageForJSONBytes(name string, targetProcess CefProcessId, data []byte)
	CreateClientHandler(alsOSR bool) (*ICefClient, bool)
	SetFocus(value bool)
	SendExternalBeginFrame()
	SendKeyEvent(event *TCefKeyEvent)
	SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32)
	SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool)
	SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32)
	SendTouchEvent(event *TCefTouchEvent)
	SendCaptureLostEvent()
	FrameIsFocused() bool
	TryCloseBrowser() bool
	BrowserHandle() types.HWND
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
	SetAcceptLanguageList(languageList string) // Remove CEF 118
	AcceptLanguageList() string                // Remove CEF 118
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
	IncZoomCommand()
	DecZoomCommand()
	ResetZoomCommand()
	DefaultZoomLevel() float64
	CanIncZoom()
	CanDecZoom()
	CanResetZoom()
	WasResized()
	WasHidden(hidden bool)
	NotifyScreenInfoChanged()
	IMESetComposition(text string, underlines []*TCefCompositionUnderline, replacementRange, selectionRange TCefRange)
	IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32)
	IMEFinishComposingText(keepSelection bool)
	IMECancelComposition()
	HasDevTools() bool
	InitializeDragAndDrop(dropTargetCtrl lcl.IWinControl)
	InitializeDragAndDropByHWND(aDropTargetWnd types.HWND)
	ShutdownDragAndDrop()
	SetNewBrowserParent(aNewParentHwnd types.HWND)
	Fullscreen() bool
	ExitFullscreen(willCauseResize bool)
	GetWebsiteSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes) *ICefValue
	SetWebsiteSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes, value *ICefValue)
	GetContentSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes) TCefContentSettingValues
	SetContentSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes, value TCefContentSettingValues)
	WindowHandle() TCefWindowHandle // Calls ICefBrowserHost.GetWindowHandle and returns the window handle for this browser.
	AsTargetWindow() target.IWindow
	IsClosing() bool
	setClosing(v bool)
	SimulateKeyEvent(keyEvent TSimulateKeyEvent, timestamp float32, text, unmodifiedtext, keyIdentifier, code, key string)
	SimulateMouseEvent(mouseEvent TSimulateMouseEvent, x, y, timestamp, force, tangentialPressure, tiltX, tiltY, deltaX, deltaY float32)
	SimulateTouchEvent(type_ TCefSimulatedTouchEventType, touchPoints []*TCefSimulatedTouchPoint, modifiers int32, timestamp float32)
	SimulateEditingCommand(command TCefEditingCommand)
	ClearCertificateExceptions(clearImmediately bool) bool
	ClearHttpAuthCredentials(clearImmediately bool) bool
	GetNavigationEntries(currentOnly bool)
	SavePreferences(fileName string)
	ExecuteTaskOnCefThread(cefThreadId TCefThreadId, taskID uint32, delayMs int64) bool
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

func (m *TCEFChromium) Options() *TChromiumOptions {
	if m.options == nil {
		m.options = NewChromiumOptions(m)
	}
	return m.options
}

func (m *TCEFChromium) FontOptions() *TChromiumFontOptions {
	if m.fontOptions == nil {
		m.fontOptions = NewChromiumFontOptions(m)
	}
	return m.fontOptions
}

func (m *TCEFChromium) Config() *TCefChromiumConfig {
	return m.config
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
	imports.Proc(def.CEFChromium_SetDefaultURL).Call(m.Instance(), api.PascalStr(defaultURL))
}

// SetEnableMultiBrowserMode 设置启用多浏览器模式
func (m *TCEFChromium) SetEnableMultiBrowserMode(enableMultiBrowserMode bool) {
	if !m.IsValid() || application.Is49() {
		return
	}
	imports.Proc(def.CEFChromium_SetMultiBrowserMode).Call(m.Instance(), api.PascalBool(enableMultiBrowserMode))
}

// LoadUrl 加载一个URL地址
func (m *TCEFChromium) LoadUrl(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_LoadURL).Call(m.Instance(), api.PascalStr(url))
}

// LoadHtml 加载HTML
func (m *TCEFChromium) LoadHtml(html string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_LoadString).Call(m.Instance(), api.PascalStr(html))
}

// StartDownload 开始下载
func (m *TCEFChromium) StartDownload(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_StartDownload).Call(m.Instance(), api.PascalStr(url))
}

// DownloadImage 开始下载图片
func (m *TCEFChromium) DownloadImage(imageUrl string, isFavicon bool, maxImageSize int32, bypassCache bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_DownloadImage).Call(m.Instance(), api.PascalStr(imageUrl), api.PascalBool(isFavicon), uintptr(maxImageSize), api.PascalBool(bypassCache))
}

func (m *TCEFChromium) Reload() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_Reload).Call(m.Instance())
}

func (m *TCEFChromium) ReloadIgnoreCache() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ReloadIgnoreCache).Call(m.Instance())
}

func (m *TCEFChromium) StopLoad() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_StopLoad).Call(m.Instance())
}

func (m *TCEFChromium) ResetZoomLevel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ResetZoomLevel).Call(m.Instance())
}

func (m *TCEFChromium) CloseAllBrowsers() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_CloseAllBrowsers).Call(m.Instance())
}

func (m *TCEFChromium) CreateBrowser(window ICEFWindowParent, windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	var (
		parentHandle TCefWindowHandle
		rect         types.TRect
	)
	if window != nil {
		parentHandle = TCefWindowHandle(window.Handle())
		rect = window.BoundsRect()
	}
	return m.CreateBrowserByWindowHandle(parentHandle, rect, windowName, context, extraInfo, false)
}

func (m *TCEFChromium) CreateBrowserByWinControl(browserParent *lcl.TWinControl, windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool {
	var control uintptr
	if browserParent != nil {
		control = browserParent.Instance()
	}
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowserByWinControl).Call(m.Instance(), control, api.PascalStr(windowName), context.Instance(), extraInfo.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) CreateBrowserByWindowHandle(parentHandle TCefWindowHandle, rect types.TRect, windowName string, context *ICefRequestContext,
	extraInfo *ICefDictionaryValue, forceAsPopup bool) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowserByWindowHandle).Call(m.Instance(), uintptr(parentHandle), uintptr(unsafe.Pointer(&rect)),
		api.PascalStr(windowName), context.Instance(), extraInfo.Instance(), api.PascalBool(forceAsPopup))
	return api.GoBool(r1)
}

func (m *TCEFChromium) CreateBrowserByBrowserViewComponent(homePage string, browserViewComponent *TCEFBrowserViewComponent, context *ICefRequestContext, extraInfo *ICefDictionaryValue) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_CreateBrowserByBrowserViewComponent).Call(m.Instance(), api.PascalStr(homePage), browserViewComponent.Instance(), context.Instance(), extraInfo.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	if m.initialized {
		return true
	}
	r1, _, _ := imports.Proc(def.CEFChromium_Initialized).Call(m.Instance())
	m.initialized = api.GoBool(r1)
	return m.initialized
}

func (m *TCEFChromium) IsSameBrowser(browser *ICefBrowser) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_IsSameBrowser).Call(m.Instance(), browser.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) PrintToPDF(saveFilePath string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_PrintToPDF).Call(m.Instance(), api.PascalStr(saveFilePath))
}

func (m *TCEFChromium) Print() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_Print).Call(m.Instance())
}

func (m *TCEFChromium) BrowserZoom(zoom ZOOM) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_BrowserZoom).Call(m.Instance(), zoom.ToPtr())
}

func (m *TCEFChromium) GoBack() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_GoBackForward).Call(m.Instance(), uintptr(BF_GOBACK))
}

func (m *TCEFChromium) GoForward() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_GoBackForward).Call(m.Instance(), uintptr(BF_GOFORWARD))
}

func (m *TCEFChromium) NotifyMoveOrResizeStarted() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_NotifyMoveOrResizeStarted).Call(m.Instance())
}

func (m *TCEFChromium) Invalidate(type_ TCefPaintElementType) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_Invalidate).Call(m.Instance(), uintptr(type_))
}

func (m *TCEFChromium) CloseBrowser(forceClose bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_CloseBrowser).Call(m.Instance(), api.PascalBool(forceClose))
}

// ExecuteJavaScript
// 执行JS代码
//
// code: js代码
//
// scriptURL: js脚本地址 默认about:blank
//
// startLine: js脚本启始执行行号
func (m *TCEFChromium) ExecuteJavaScript(code, scriptURL string, frame *ICefFrame, startLine int32) {
	imports.Proc(def.CEFChromium_ExecuteJavaScript).Call(m.Instance(), api.PascalStr(code), api.PascalStr(scriptURL), frame.Instance(), uintptr(startLine))
}

func (m *TCEFChromium) ShowDevTools(window ICEFWindowParent) {
	if !m.IsValid() {
		return
	}
	if window == nil {
		imports.Proc(def.CEFChromium_ShowDevTools).Call(m.Instance())
	} else {
		imports.Proc(def.CEFChromium_ShowDevToolsByWindowParent).Call(m.Instance(), window.Instance())
	}
}

func (m *TCEFChromium) CloseDevTools(window ICEFWindowParent) {
	if !m.IsValid() {
		return
	}
	if window == nil {
		imports.Proc(def.CEFChromium_CloseDevTools).Call(m.Instance())
	} else {
		imports.Proc(def.CEFChromium_CloseDevToolsByWindowParent).Call(m.Instance(), window.Instance())
	}
}

// VisitAllCookies 查看所有cookie,该函数触发 OnCookiesVisited 事件返回结果
func (m *TCEFChromium) VisitAllCookies(id int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_VisitAllCookies).Call(m.Instance(), uintptr(id))
}

// VisitURLCookies 查看URL cookie,该函数触发 OnCookiesVisited 事件返回结果
// url https://www.demo.com
func (m *TCEFChromium) VisitURLCookies(url string, includeHttpOnly bool, id int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_VisitURLCookies).Call(m.Instance(), api.PascalStr(url), api.PascalBool(includeHttpOnly), uintptr(id))
}

// 删除所有cookie
func (m *TCEFChromium) DeleteCookies(url, cookieName string, deleteImmediately bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_DeleteCookies).Call(m.Instance(), api.PascalStr(url), api.PascalStr(cookieName), api.PascalBool(deleteImmediately))
}

func (m *TCEFChromium) SetCookie(url, name, value, domain, path string,
	secure, httponly, hasExpires bool,
	creation, lastAccess, expires time.Time,
	sameSite TCefCookieSameSite, priority TCefCookiePriority, aSetImmediately bool, aID int32) {
	if !m.IsValid() {
		return
	}
	cookie := &TCefCookie{
		Url:            url,
		Name:           name,
		Value:          value,
		Domain:         domain,
		Path:           path,
		Secure:         secure,
		Httponly:       httponly,
		HasExpires:     hasExpires,
		Creation:       creation,
		LastAccess:     lastAccess,
		Expires:        expires,
		SameSite:       sameSite,
		Priority:       priority,
		SetImmediately: aSetImmediately,
		ID:             aID,
	}
	cookiePtr := cookie.ToPtr()
	imports.Proc(def.CEFChromium_SetCookie).Call(m.Instance(), uintptr(unsafe.Pointer(cookiePtr)))
}

func (m *TCEFChromium) FlushCookieStore(flushImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_FlushCookieStore).Call(m.Instance(), api.PascalBool(flushImmediately))
	return api.GoBool(r1)
}

func (m *TCEFChromium) SetProxy(proxy TCefProxy) {
	if !m.IsValid() {
		return
	}
	proxyPtr := proxy.ToPtr()
	imports.Proc(def.CEFChromium_SetProxy).Call(m.Instance(), uintptr(unsafe.Pointer(proxyPtr)))
}

func (m *TCEFChromium) UpdatePreferences() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_UpdatePreferences).Call(m.Instance())
}

func (m *TCEFChromium) SendDevToolsMessage(message string) bool {
	r1, _, _ := imports.Proc(def.CEFChromium_SendDevToolsMessage).Call(m.Instance(), api.PascalStr(message))
	return api.GoBool(r1)
}

func (m *TCEFChromium) ExecuteDevToolsMethod(messageId int32, method string, dictionaryValue *ICefDictionaryValue) int32 {
	if !m.IsValid() {
		return 0
	}
	if dictionaryValue == nil {
		dictionaryValue = DictionaryValueRef.New()
	}
	r1, _, _ := imports.Proc(def.CEFChromium_ExecuteDevToolsMethod).Call(m.Instance(), uintptr(messageId), api.PascalStr(method), dictionaryValue.Instance())
	return int32(r1)
}

func (m *TCEFChromium) CreateClientHandler(alsOSR bool) (*ICefClient, bool) {
	if !m.IsValid() {
		return nil, false
	}
	var result uintptr
	r1, _, _ := imports.Proc(def.CEFChromium_CreateClientHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)), api.PascalBool(alsOSR))
	if result != 0 && api.GoBool(r1) {
		return &ICefClient{instance: unsafe.Pointer(result), ct: CtOther}, api.GoBool(r1)
	}
	return nil, false
}

func (m *TCEFChromium) SetFocus(value bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetFocus).Call(m.Instance(), api.PascalBool(value))
}

func (m *TCEFChromium) SendExternalBeginFrame() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendExternalBeginFrame).Call(m.Instance())
}

func (m *TCEFChromium) SendKeyEvent(event *TCefKeyEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendKeyEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)))
}

func (m *TCEFChromium) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendMouseClickEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), uintptr(type_), api.PascalBool(mouseUp), uintptr(clickCount))
}

func (m *TCEFChromium) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendMouseMoveEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), api.PascalBool(mouseLeave))
}

func (m *TCEFChromium) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendMouseWheelEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TCEFChromium) SendTouchEvent(event *TCefTouchEvent) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendTouchEvent).Call(m.Instance(), uintptr(unsafe.Pointer(event)))
}

func (m *TCEFChromium) SendCaptureLostEvent() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendCaptureLostEvent).Call(m.Instance())
}

func (m *TCEFChromium) FrameIsFocused() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_FrameIsFocused).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) TryCloseBrowser() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_TryCloseBrowser).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) BrowserHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	if m.browserHandle == 0 {
		r1, _, _ := imports.Proc(def.CEFChromium_BrowserHandle).Call(m.Instance())
		m.browserHandle = types.HWND(r1)
	}
	return m.browserHandle
}

func (m *TCEFChromium) RenderHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	if m.renderHandle == 0 {
		r1, _, _ := imports.Proc(def.CEFChromium_RenderHandle).Call(m.Instance())
		m.renderHandle = types.HWND(r1)
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

// SetAcceptLanguageList Remove CEF 118
func (m *TCEFChromium) SetAcceptLanguageList(languageList string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetAcceptLanguageList).Call(m.Instance(), api.PascalStr(languageList))
}

// AcceptLanguageList Remove CEF 118
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
func (m *TCEFChromium) SendProcessMessageForJSONBytes(name string, targetProcess CefProcessId, data []byte) {
	if !m.IsValid() {
		return
	}
	if !m.initialized {
		m.initialized = m.Initialized()
		if !m.initialized {
			return
		}
	}
	imports.Proc(def.CEFChromium_SendProcessMessageForJSONBytes).Call(m.Instance(), api.PascalStr(name), targetProcess.ToPtr(), uintptr(unsafe.Pointer(&data[0])), uintptr(uint32(len(data))))
}

// SendProcessMessageForV8Value 发送进程消息
func (m *TCEFChromium) SendProcessMessageForV8Value(messageName string, targetProcess CefProcessId, arguments *ICefV8Value) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SendProcessMessageForV8Value).Call(m.Instance(), api.PascalStr(messageName), targetProcess.ToPtr(), arguments.Instance())
}

// Client 获取Client
func (m *TCEFChromium) Client() *ICefClient {
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
	imports.Proc(def.CEFChromium_ClearDataForOrigin).Call(m.Instance(), api.PascalStr(origin), uintptr(st))
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

func (m *TCEFChromium) IncZoomCommand() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ZoomCommand).Call(m.Instance(), ZcInc)
}

func (m *TCEFChromium) DecZoomCommand() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ZoomCommand).Call(m.Instance(), ZcDec)
}

func (m *TCEFChromium) ResetZoomCommand() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ZoomCommand).Call(m.Instance(), ZcReset)
}

func (m *TCEFChromium) DefaultZoomLevel() (result float64) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_DefaultZoomLevel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFChromium) CanIncZoom() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_CanZoom).Call(m.Instance(), CzInc)
}

func (m *TCEFChromium) CanDecZoom() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_CanZoom).Call(m.Instance(), CzDec)
}

func (m *TCEFChromium) CanResetZoom() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_CanZoom).Call(m.Instance(), CzReset)
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

func (m *TCEFChromium) IMESetComposition(text string, underlines []*TCefCompositionUnderline, replacementRange, selectionRange TCefRange) {
	if !m.IsValid() {
		return
	}
	var size = len(underlines)
	underlinesPtr := make([]tCefCompositionUnderlinePtr, size, size)
	for i := 0; i < size; i++ {
		line := underlines[i]
		underlinesPtr[i] = *line.ToPtr()
	}
	imports.Proc(def.CEFChromium_IMESetComposition).Call(m.Instance(), api.PascalStr(text), uintptr(unsafePointer(&replacementRange)), uintptr(unsafePointer(&selectionRange)),
		uintptr(unsafePointer(&underlinesPtr[0])), uintptr(int32(size)))
}

func (m *TCEFChromium) IMECommitText(text string, replacementRange TCefRange, relativeCursorPos int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_IMECommitText).Call(m.Instance(), api.PascalStr(text), uintptr(unsafe.Pointer(&replacementRange)), uintptr(relativeCursorPos))
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

func (m *TCEFChromium) HasDevTools() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_HasDevTools).Call(m.Instance())
	return api.GoBool(r1)
}

// InitializeDragAndDrop
//
//	By Windows: Used with browsers in OSR mode to initialize drag and drop in Windows.
func (m *TCEFChromium) InitializeDragAndDrop(dropTargetCtrl lcl.IWinControl) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_InitializeDragAndDrop).Call(m.Instance(), lcl.CheckPtr(dropTargetCtrl))
}

// InitializeDragAndDropByHWND
//
//	By Windows: Used with browsers in OSR mode to initialize drag and drop in Windows.
func (m *TCEFChromium) InitializeDragAndDropByHWND(aDropTargetWnd types.HWND) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_InitializeDragAndDropByHWND).Call(m.Instance(), aDropTargetWnd)
}

// ShutdownDragAndDrop
//
//	By Windows: Used with browsers in OSR mode to shutdown drag and drop in Windows.
func (m *TCEFChromium) ShutdownDragAndDrop() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ShutdownDragAndDrop).Call(m.Instance())
}

// SetNewBrowserParent
//
//	By Windows: Used to reparent the browser to a different TCEFWindowParent.
func (m *TCEFChromium) SetNewBrowserParent(aNewParentHwnd types.HWND) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetNewBrowserParent).Call(m.Instance(), aNewParentHwnd)
}

// EmitRender IPC 发送进程 消息
//
// messageId != 0 是带有回调函数消息
func (m *TCEFChromium) EmitRender(messageId int32, eventName string, target target.ITarget, data ...interface{}) bool {
	if !m.initialized {
		m.initialized = m.Initialized()
		if !m.initialized {
			return false
		}
	}
	if target == nil || target.BrowserId() <= 0 || target.ChannelId() == "" {
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
	} else {
		browse := BrowserWindow.GetBrowser(target.BrowserId())
		//browse := m.BrowserById(target.BrowserId())
		if browse != nil && browse.IsValid() {
			if application.Is49() {
				// CEF49
				browse.EmitRender(messageId, eventName, target, data...)
			} else {
				frame := browse.GetFrameById(target.ChannelId())
				if frame != nil && frame.IsValid() {
					return frame.EmitRender(messageId, eventName, target, data...)
				}
			}
		}
	}
	return false
}

func (m *TCEFChromium) Fullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_Fullscreen).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFChromium) ExitFullscreen(willCauseResize bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_ExitFullscreen).Call(m.Instance(), api.PascalBool(willCauseResize))
}

func (m *TCEFChromium) GetWebsiteSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes) *ICefValue {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFChromium_GetWebsiteSetting).Call(m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType))
	if result > 0 {
		return &ICefValue{instance: getInstance(result)}
	}
	return nil
}

func (m *TCEFChromium) SetWebsiteSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes, value *ICefValue) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetWebsiteSetting).Call(m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType), value.Instance())
}

func (m *TCEFChromium) GetContentSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes) TCefContentSettingValues {
	if !m.IsValid() {
		return CEF_CONTENT_SETTING_VALUE_DEFAULT
	}
	r1, _, _ := imports.Proc(def.CEFChromium_GetContentSetting).Call(m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType))
	return TCefContentSettingValues(r1)
}

func (m *TCEFChromium) SetContentSetting(requestingUrl, topLevelUrl string, contentType TCefContentSettingTypes, value TCefContentSettingValues) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SetContentSetting).Call(m.Instance(), api.PascalStr(requestingUrl), api.PascalStr(topLevelUrl), uintptr(contentType), uintptr(value))
}

func (m *TCEFChromium) WindowHandle() TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	var result uintptr
	imports.Proc(def.CEFChromium_WindowHandle).Call(m.Instance(), uintptr(unsafePointer(&result)))
	return TCefWindowHandle(result)
}

// Target
//
//	IPC消息接收目标, 当前窗口chromium发送
//	参数: targetType 可选, 接收类型
func (m *TCEFChromium) Target() target.ITarget {
	if !m.IsValid() {
		return nil
	}
	browse := m.Browser()
	if !browse.IsValid() {
		return nil
	}
	return target.NewTarget(m, browse.Identifier(), browse.MainFrame().Identifier())
}

// ProcessMessage
//
//	IPC消息触发当前Chromium
func (m *TCEFChromium) ProcessMessage() target.IProcessMessage {
	if m == nil {
		return nil
	}
	return m
}

// AsTargetWindow 转换为 IPC 目标接收窗口
func (m *TCEFChromium) AsTargetWindow() target.IWindow {
	return m
}

// IsClosing 返回窗口是否正在关闭/或已关闭 true正在或已关闭
func (m *TCEFChromium) IsClosing() bool {
	return m.isClosing
}

// 当窗口关闭时设置为true
func (m *TCEFChromium) setClosing(v bool) {
	m.isClosing = v
}

func (m *TCEFChromium) SimulateKeyEvent(keyEvent TSimulateKeyEvent, timestamp float32, text, unmodifiedtext, keyIdentifier, code, key string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SimulateKeyEvent).Call(m.Instance(), uintptr(unsafePointer(&keyEvent)), uintptr(unsafePointer(&timestamp)), api.PascalStr(text),
		api.PascalStr(unmodifiedtext), api.PascalStr(keyIdentifier), api.PascalStr(code), api.PascalStr(key))
}

func (m *TCEFChromium) SimulateMouseEvent(mouseEvent TSimulateMouseEvent, x, y, timestamp, force, tangentialPressure, tiltX, tiltY, deltaX, deltaY float32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SimulateMouseEvent).Call(m.Instance(), uintptr(unsafePointer(&mouseEvent)), uintptr(unsafePointer(&x)),
		uintptr(unsafePointer(&y)), uintptr(unsafePointer(&timestamp)), uintptr(unsafePointer(&force)), uintptr(unsafePointer(&tangentialPressure)),
		uintptr(unsafePointer(&tiltX)), uintptr(unsafePointer(&tiltY)), uintptr(unsafePointer(&deltaX)), uintptr(unsafePointer(&deltaY)))
}

func (m *TCEFChromium) SimulateTouchEvent(type_ TCefSimulatedTouchEventType, touchPoints []*TCefSimulatedTouchPoint, modifiers int32, timestamp float32) {
	if !m.IsValid() {
		return
	}
	var size = len(touchPoints)
	touchPointsPtr := make([]tCefSimulatedTouchPointPtr, size, size)
	for i := 0; i < size; i++ {
		tp := touchPoints[i]
		touchPointsPtr[i] = *tp.ToPtr()
	}
	imports.Proc(def.CEFChromium_SimulateTouchEvent).Call(m.Instance(), uintptr(type_), uintptr(unsafePointer(&touchPointsPtr[0])), uintptr(int32(size)), uintptr(modifiers), uintptr(unsafePointer(&timestamp)))
}

func (m *TCEFChromium) SimulateEditingCommand(command TCefEditingCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SimulateEditingCommand).Call(m.Instance(), uintptr(command))
}

func (m *TCEFChromium) ClearCertificateExceptions(clearImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_ClearCertificateExceptions).Call(m.Instance(), api.PascalBool(clearImmediately))
	return api.GoBool(r1)
}

func (m *TCEFChromium) ClearHttpAuthCredentials(clearImmediately bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_ClearHttpAuthCredentials).Call(m.Instance(), api.PascalBool(clearImmediately))
	return api.GoBool(r1)
}

func (m *TCEFChromium) GetNavigationEntries(currentOnly bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_GetNavigationEntries).Call(m.Instance(), api.PascalBool(currentOnly))
}

func (m *TCEFChromium) SavePreferences(fileName string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFChromium_SavePreferences).Call(m.Instance(), api.PascalStr(fileName))
}

func (m *TCEFChromium) ExecuteTaskOnCefThread(cefThreadId TCefThreadId, taskID uint32, delayMs int64) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFChromium_ExecuteTaskOnCefThread).Call(m.Instance(), uintptr(cefThreadId), uintptr(taskID), uintptr(unsafePointer(&delayMs)))
	return api.GoBool(r1)
}
