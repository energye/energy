//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 应用程序的属性配置

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/cef/i18n"
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/energy/consts"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl/api"
	"path/filepath"
	"strings"
	"unsafe"
)

// initDefaultSettings 初始 energy 默认设置
func (m *TCEFApplication) initDefaultSettings() {
	if m.FrameworkDirPath() == "" {
		// 默认 CEF Framework 目录
		cfg := config.Get()
		if cfg != nil {
			libCef := func() string {
				if common.IsWindows() {
					return "libcef.dll"
				} else if common.IsLinux() {
					return "libcef.so"
				}
				return ""
			}()
			if libCef != "" {
				if tools.IsExist(filepath.Join(consts.ExeDir, libCef)) {
					m.SetFrameworkDirPath(consts.ExeDir)
				} else if frameworkDir := cfg.FrameworkPath(); tools.IsExist(filepath.Join(frameworkDir, libCef)) {
					m.SetFrameworkDirPath(frameworkDir)
				}
			}
		}
	}

	m.SetLocale(LANGUAGE_zh_CN)
	m.SetLogSeverity(LOGSEVERITY_DISABLE)
	m.SetEnablePrintPreview(true)
	// m.SetEnableGPU(true) 默认还是关闭GPU加速
	// DefaultMessageLoop() 根据不同平台, 启动不同的窗口组件
	// ViewsFrameworkBrowserWindow 简称(VF)窗口组件, 同时支持 Windows/Linux/MacOSX
	// LCL 窗口组件,同时支持 Windows/MacOSX, CEF版本<=106.xx时支持GTK2, CEF版本 >= 107.xx时默认开启 GTK3 且不支持 GTK2 和 LCL提供的各种组件
	m.DefaultMessageLoop()
}

// DefaultMessageLoop 默认消息轮询, 在创建 CEF Application 时确定使用什么方式
func (m *TCEFApplication) DefaultMessageLoop() {
	if common.IsLinux() { // Linux => (VF)View Framework 窗口
		ui := api.WidgetUI()
		if ui.IsGTK3() {
			// Linux CEF >= 107.xxx 版本以后，默认启用的GTK3，106及以前版本默认支持GTK2但无法正常输入中文
			// Linux 默认设置为false,将启用 ViewsFrameworkBrowserWindow 窗口
			m.SetExternalMessagePump(false)
			m.SetMultiThreadedMessageLoop(false)
		} else if ui.IsGTK2() {
			// GTK2 默认支持LCL,但还未解决无法输入中文问题
			m.SetExternalMessagePump(false)
			m.SetMultiThreadedMessageLoop(true)
		}
		// 这是一个解决“GPU不可用错误”问题的方法 linux
		// https://bitbucket.org/chromiumembedded/cef/issues/2964/gpu-is-not-usable-error-during-cef
		m.SetDisableZygote(true)
	} else if common.IsDarwin() { // Darwin => LCL窗口
		GlobalWorkSchedulerCreate(nil)
		m.SetOnScheduleMessagePumpWork(nil)
		// MacOSX 在使用LCL窗口组件必须将 ExternalMessagePump=true 和 MultiThreadedMessageLoop=false
		// 或
		// 同 Linux 一样使用 ViewsFrameworkBrowserWindow 窗口组件
		m.SetExternalMessagePump(true)
		m.SetMultiThreadedMessageLoop(false)
	} else { // Windows => LCL窗口
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(true)
	}
}

/*** 自定义属性 ***/

/*** 设置 TCefSettings (cef_settings_t) 属性 ***/

func (m *TCEFApplication) NoSandbox() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_NoSandbox).Call()
	return api.GoBool(r1)
}

// SetNoSandbox
// Set to true (1) to disable the sandbox for sub-processes. See
// cef_sandbox_win.h for requirements to enable the sandbox on Windows. Also
// configurable using the "no-sandbox" command-line switch.
func (m *TCEFApplication) SetNoSandbox(value bool) {
	imports.Proc(def.CEFAppConfig_SetNoSandbox).Call(api.PascalBool(value))
}

func (m *TCEFApplication) BrowserSubprocessPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_BrowserSubprocessPath).Call()
	return api.GoStr(r1)
}

// SetBrowserSubprocessPath The path to a separate executable that will be launched for sub-processes.
// If this value is empty on Windows or Linux then the main process
// executable will be used. If this value is empty on macOS then a helper
// executable must exist at "Contents/Frameworks/<app>
// Helper.app/Contents/MacOS/<app> Helper" in the top-level app bundle. See
// the comments on CefExecuteProcess() for details. If this value is
// non-empty then it must be an absolute path. Also configurable using the
// "browser-subprocess-path" command-line switch.
func (m *TCEFApplication) SetBrowserSubprocessPath(value string) {
	imports.Proc(def.CEFAppConfig_SetBrowserSubprocessPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) FrameworkDirPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_FrameworkDirPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetFrameworkDirPath(value string) {
	imports.Proc(def.CEFAppConfig_SetFrameworkDirPath).Call(api.PascalStr(value))
	// resources 和 locals 在同一目录
	m.SetResourcesDirPath(value)
	m.SetLocalesDirPath(filepath.Join(value, "locales"))
}

// MainBundlePath 仅用于macOS
func (m *TCEFApplication) MainBundlePath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MainBundlePath).Call()
	return api.GoStr(r1)
}

// SetMainBundlePath 仅用于macOS
func (m *TCEFApplication) SetMainBundlePath(value string) {
	imports.Proc(def.CEFAppConfig_SetMainBundlePath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ChromeRuntime() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeRuntime).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetChromeRuntime(value bool) {
	imports.Proc(def.CEFAppConfig_SetChromeRuntime).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MultiThreadedMessageLoop() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MultiThreadedMessageLoop).Call()
	return api.GoBool(r1)
}

// IsMessageLoop
//
//	不同的窗口组件使用不同的消息轮询
//	return
//		true : VF(views framework)窗口组件
//		false: LCL窗口组件
func (m *TCEFApplication) IsMessageLoop() bool {
	return !m.externalMessagePump && !m.multiThreadedMessageLoop
}

func (m *TCEFApplication) SetMultiThreadedMessageLoop(value bool) {
	m.multiThreadedMessageLoop = value
	imports.Proc(def.CEFAppConfig_SetMultiThreadedMessageLoop).Call(api.PascalBool(value))
}

// EnableVFWindow 启用VF(ViewsFramework)窗口, Linux默认该模式，非Linux需要强制开启才可使用
func (m *TCEFApplication) EnableVFWindow(e bool) {
	if e {
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(false)
	} else {
		m.DefaultMessageLoop()
	}
}

func (m *TCEFApplication) ExternalMessagePump() bool {
	if !m.Is49() {
		r1, _, _ := imports.Proc(def.CEFAppConfig_ExternalMessagePump).Call()
		return api.GoBool(r1)
	}
	return false
}

func (m *TCEFApplication) SetExternalMessagePump(value bool) {
	if !m.Is49() {
		m.externalMessagePump = value
		imports.Proc(def.CEFAppConfig_SetExternalMessagePump).Call(api.PascalBool(value))
	}
}

func (m *TCEFApplication) WindowlessRenderingEnabled() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_WindowlessRenderingEnabled).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetWindowlessRenderingEnabled(value bool) {
	imports.Proc(def.CEFAppConfig_SetWindowlessRenderingEnabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) CommandLineArgsDisabled() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_CommandLineArgsDisabled).Call()
	return api.GoBool(r1)
}

// SetCommandLineArgsDisabled 开启/禁用设置命令行参数
func (m *TCEFApplication) SetCommandLineArgsDisabled(value bool) {
	imports.Proc(def.CEFAppConfig_SetCommandLineArgsDisabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) Cache() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_Cache).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetCache(value string) {
	imports.Proc(def.CEFAppConfig_SetCache).Call(api.PascalStr(value))
}

func (m *TCEFApplication) RootCache() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_RootCache).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetRootCache(value string) {
	imports.Proc(def.CEFAppConfig_SetRootCache).Call(api.PascalStr(value))
}

// UserDataPath
//
//	CEF 115 Remove
func (m *TCEFApplication) UserDataPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UserDataPath).Call()
	return api.GoStr(r1)
}

// SetUserDataPath
//
//	CEF 115 Remove
func (m *TCEFApplication) SetUserDataPath(value string) {
	imports.Proc(def.CEFAppConfig_SetUserDataPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) PersistSessionCookies() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_PersistSessionCookies).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetPersistSessionCookies(value bool) {
	imports.Proc(def.CEFAppConfig_SetPersistSessionCookies).Call(api.PascalBool(value))
}

func (m *TCEFApplication) PersistUserPreferences() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_PersistUserPreferences).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetPersistUserPreferences(value bool) {
	imports.Proc(def.CEFAppConfig_SetPersistUserPreferences).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UserAgent() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UserAgent).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetUserAgent(value string) {
	imports.Proc(def.CEFAppConfig_SetUserAgent).Call(api.PascalStr(value))
}

func (m *TCEFApplication) UserAgentProduct() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UserAgentProduct).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetUserAgentProduct(value string) {
	imports.Proc(def.CEFAppConfig_SetUserAgentProduct).Call(api.PascalStr(value))
}

func (m *TCEFApplication) Locale() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_Locale).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLocale(value LANGUAGE) {
	i18n.Switch(value)
	imports.Proc(def.CEFAppConfig_SetLocale).Call(value.ToPtr())
}

func (m *TCEFApplication) LogFile() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LogFile).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLogFile(value string) {
	imports.Proc(def.CEFAppConfig_SetLogFile).Call(api.PascalStr(value))
}

func (m *TCEFApplication) LogSeverity() LogSeverity {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LogSeverity).Call()
	return LogSeverity(r1)
}

func (m *TCEFApplication) SetLogSeverity(value LogSeverity) {
	imports.Proc(def.CEFAppConfig_SetLogSeverity).Call(value.ToPtr())
}

func (m *TCEFApplication) LogItems() TCefLogItems {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LogItems).Call(GetValue, 0)
	return TCefLogItems(r1)
}

func (m *TCEFApplication) SetLogItems(value TCefLogItems) {
	imports.Proc(def.CEFAppConfig_LogItems).Call(SetValue, uintptr(value))
}

func (m *TCEFApplication) JavaScriptFlags() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_JavaScriptFlags).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetJavaScriptFlags(value string) {
	imports.Proc(def.CEFAppConfig_SetJavaScriptFlags).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ResourcesDirPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ResourcesDirPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetResourcesDirPath(value string) {
	imports.Proc(def.CEFAppConfig_SetResourcesDirPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) LocalesDirPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LocalesDirPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLocalesDirPath(value string) {
	imports.Proc(def.CEFAppConfig_SetLocalesDirPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) PackLoadingDisabled() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_PackLoadingDisabled).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetPackLoadingDisabled(value bool) {
	imports.Proc(def.CEFAppConfig_SetPackLoadingDisabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) RemoteDebuggingPort() int32 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_RemoteDebuggingPort).Call()
	return int32(r1)
}

func (m *TCEFApplication) SetRemoteDebuggingPort(value int32) {
	if value > 1024 && value < 65535 {
		imports.Proc(def.CEFAppConfig_SetRemoteDebuggingPort).Call(uintptr(value))
	}
}

func (m *TCEFApplication) UncaughtExceptionStackSize() int32 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UncaughtExceptionStackSize).Call()
	return int32(r1)
}

func (m *TCEFApplication) SetUncaughtExceptionStackSize(value int32) {
	imports.Proc(def.CEFAppConfig_SetUncaughtExceptionStackSize).Call(uintptr(value))
}

func (m *TCEFApplication) IgnoreCertificateErrors() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_IgnoreCertificateErrors).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetIgnoreCertificateErrors(value bool) {
	imports.Proc(def.CEFAppConfig_SetIgnoreCertificateErrors).Call(api.PascalBool(value))
}

func (m *TCEFApplication) BackgroundColor() types.TCefColor {
	r1, _, _ := imports.Proc(def.CEFAppConfig_BackgroundColor).Call()
	return types.TCefColor(r1)
}

func (m *TCEFApplication) SetBackgroundColor(value types.TCefColor) {
	imports.Proc(def.CEFAppConfig_SetBackgroundColor).Call(value.ToPtr())
}

// AcceptLanguageList Remove CEF 118
func (m *TCEFApplication) AcceptLanguageList() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AcceptLanguageList).Call()
	return api.GoStr(r1)
}

// SetAcceptLanguageList Remove CEF 118
func (m *TCEFApplication) SetAcceptLanguageList(value string) {
	imports.Proc(def.CEFAppConfig_SetAcceptLanguageList).Call(api.PascalStr(value))
}

func (m *TCEFApplication) CookieableSchemesList() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_CookieableSchemesList).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetCookieableSchemesList(value string) {
	imports.Proc(def.CEFAppConfig_SetCookieableSchemesList).Call(api.PascalStr(value))
}

func (m *TCEFApplication) CookieableSchemesExcludeDefaults() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_CookieableSchemesExcludeDefaults).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetCookieableSchemesExcludeDefaults(value bool) {
	imports.Proc(def.CEFAppConfig_SetCookieableSchemesExcludeDefaults).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ChromePolicyId() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromePolicyId).Call(GetValue)
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetChromePolicyId(value string) {
	imports.Proc(def.CEFAppConfig_ChromePolicyId).Call(SetValue, api.PascalStr(value))
}

/*** 设置常用的命令行参数属性 ***/

func (m *TCEFApplication) SingleProcess() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_SingleProcess).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetSingleProcess(value bool) {
	imports.Proc(def.CEFAppConfig_SetSingleProcess).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableMediaStream() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableMediaStream).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableMediaStream(value bool) {
	imports.Proc(def.CEFAppConfig_SetEnableMediaStream).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableSpeechInput() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableSpeechInput).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableSpeechInput(value bool) {
	imports.Proc(def.CEFAppConfig_SetEnableSpeechInput).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UseFakeUIForMediaStream() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UseFakeUIForMediaStream).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetUseFakeUIForMediaStream(value bool) {
	imports.Proc(def.CEFAppConfig_SetUseFakeUIForMediaStream).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableUsermediaScreenCapturing() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableUsermediaScreenCapturing).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableUsermediaScreenCapturing(value bool) {
	imports.Proc(def.CEFAppConfig_SetEnableUsermediaScreenCapturing).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableGPU() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableGPU).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableGPU(value bool) {
	imports.Proc(def.CEFAppConfig_SetEnableGPU).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableFeatures() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetEnableFeatures(value string) {
	imports.Proc(def.CEFAppConfig_SetEnableFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) DisableFeatures() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDisableFeatures(value string) {
	imports.Proc(def.CEFAppConfig_SetDisableFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) EnableBlinkFeatures() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableBlinkFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetEnableBlinkFeatures(value string) {
	imports.Proc(def.CEFAppConfig_SetEnableBlinkFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) DisableBlinkFeatures() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableBlinkFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDisableBlinkFeatures(value string) {
	imports.Proc(def.CEFAppConfig_SetDisableBlinkFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) BlinkSettings() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_BlinkSettings).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetBlinkSettings(value string) {
	imports.Proc(def.CEFAppConfig_SetBlinkSettings).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ForceFieldTrials() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ForceFieldTrials).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetForceFieldTrials(value string) {
	imports.Proc(def.CEFAppConfig_SetForceFieldTrials).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ForceFieldTrialParams() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ForceFieldTrialParams).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetForceFieldTrialParams(value string) {
	imports.Proc(def.CEFAppConfig_SetForceFieldTrialParams).Call(api.PascalStr(value))
}

func (m *TCEFApplication) SmoothScrolling() TCefState {
	r1, _, _ := imports.Proc(def.CEFAppConfig_SmoothScrolling).Call()
	return TCefState(r1)
}

func (m *TCEFApplication) SetSmoothScrolling(value TCefState) {
	imports.Proc(def.CEFAppConfig_SetSmoothScrolling).Call(value.ToPtr())
}

func (m *TCEFApplication) FastUnload() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_FastUnload).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetFastUnload(value bool) {
	imports.Proc(def.CEFAppConfig_SetFastUnload).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableSafeBrowsing() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableSafeBrowsing).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableSafeBrowsing(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableSafeBrowsing).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MuteAudio() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MuteAudio).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMuteAudio(value bool) {
	imports.Proc(def.CEFAppConfig_SetMuteAudio).Call(api.PascalBool(value))
}

func (m *TCEFApplication) SitePerProcess() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_SitePerProcess).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetSitePerProcess(value bool) {
	imports.Proc(def.CEFAppConfig_SetSitePerProcess).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableWebSecurity() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableWebSecurity).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableWebSecurity(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableWebSecurity).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisablePDFExtension() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisablePDFExtension).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisablePDFExtension(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisablePDFExtension).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableSiteIsolationTrials() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableSiteIsolationTrials).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableSiteIsolationTrials(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableSiteIsolationTrials).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableChromeLoginPrompt() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableChromeLoginPrompt).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableChromeLoginPrompt(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableChromeLoginPrompt).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableExtensions() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableExtensions).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableExtensions(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableExtensions).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AutoplayPolicy() TCefAutoplayPolicy {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AutoplayPolicy).Call()
	return TCefAutoplayPolicy(r1)
}

func (m *TCEFApplication) SetAutoplayPolicy(value TCefAutoplayPolicy) {
	imports.Proc(def.CEFAppConfig_SetAutoplayPolicy).Call(value.ToPtr())
}

func (m *TCEFApplication) DisableBackgroundNetworking() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableBackgroundNetworking).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableBackgroundNetworking(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableBackgroundNetworking).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MetricsRecordingOnly() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MetricsRecordingOnly).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMetricsRecordingOnly(value bool) {
	imports.Proc(def.CEFAppConfig_SetMetricsRecordingOnly).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowFileAccessFromFiles() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AllowFileAccessFromFiles).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowFileAccessFromFiles(value bool) {
	imports.Proc(def.CEFAppConfig_SetAllowFileAccessFromFiles).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowRunningInsecureContent() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AllowRunningInsecureContent).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowRunningInsecureContent(value bool) {
	imports.Proc(def.CEFAppConfig_SetAllowRunningInsecureContent).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnablePrintPreview() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnablePrintPreview).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnablePrintPreview(value bool) {
	if !m.Is49() {
		imports.Proc(def.CEFAppConfig_SetEnablePrintPreview).Call(api.PascalBool(value))
	}
}

func (m *TCEFApplication) DefaultEncoding() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DefaultEncoding).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDefaultEncoding(value string) {
	imports.Proc(def.CEFAppConfig_SetDefaultEncoding).Call(api.PascalStr(value))
}

func (m *TCEFApplication) DisableJavascript() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableJavascript).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascript(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableJavascript).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableJavascriptCloseWindows() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableJavascriptCloseWindows).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascriptCloseWindows(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableJavascriptCloseWindows).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableJavascriptAccessClipboard() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableJavascriptAccessClipboard).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascriptAccessClipboard(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableJavascriptAccessClipboard).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableJavascriptDomPaste() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableJavascriptDomPaste).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascriptDomPaste(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableJavascriptDomPaste).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowUniversalAccessFromFileUrls() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AllowUniversalAccessFromFileUrls).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowUniversalAccessFromFileUrls(value bool) {
	imports.Proc(def.CEFAppConfig_SetAllowUniversalAccessFromFileUrls).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableImageLoading() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableImageLoading).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableImageLoading(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableImageLoading).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ImageShrinkStandaloneToFit() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ImageShrinkStandaloneToFit).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetImageShrinkStandaloneToFit(value bool) {
	imports.Proc(def.CEFAppConfig_SetImageShrinkStandaloneToFit).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableTextAreaResize() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableTextAreaResize).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableTextAreaResize(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableTextAreaResize).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableTabToLinks() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableTabToLinks).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableTabToLinks(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableTabToLinks).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableProfanityFilter() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_EnableProfanityFilter).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableProfanityFilter(value bool) {
	imports.Proc(def.CEFAppConfig_SetEnableProfanityFilter).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableSpellChecking() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableSpellChecking).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableSpellChecking(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableSpellChecking).Call(api.PascalBool(value))
}

func (m *TCEFApplication) OverrideSpellCheckLang() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_OverrideSpellCheckLang).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetOverrideSpellCheckLang(value string) {
	imports.Proc(def.CEFAppConfig_SetOverrideSpellCheckLang).Call(api.PascalStr(value))
}

func (m *TCEFApplication) TouchEvents() TCefState {
	r1, _, _ := imports.Proc(def.CEFAppConfig_TouchEvents).Call()
	return TCefState(r1)
}

func (m *TCEFApplication) SetTouchEvents(value TCefState) {
	imports.Proc(def.CEFAppConfig_SetTouchEvents).Call(value.ToPtr())
}

func (m *TCEFApplication) DisableReadingFromCanvas() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableReadingFromCanvas).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableReadingFromCanvas(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableReadingFromCanvas).Call(api.PascalBool(value))
}

func (m *TCEFApplication) HyperlinkAuditing() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_HyperlinkAuditing).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetHyperlinkAuditing(value bool) {
	imports.Proc(def.CEFAppConfig_SetHyperlinkAuditing).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableNewBrowserInfoTimeout() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableNewBrowserInfoTimeout).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableNewBrowserInfoTimeout(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableNewBrowserInfoTimeout).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DevToolsProtocolLogFile() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DevToolsProtocolLogFile).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDevToolsProtocolLogFile(value string) {
	imports.Proc(def.CEFAppConfig_SetDevToolsProtocolLogFile).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ForcedDeviceScaleFactor() (result float32) { //single
	imports.Proc(def.CEFAppConfig_ForcedDeviceScaleFactor).Call(uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFApplication) SetForcedDeviceScaleFactor(value float32) { //single
	imports.Proc(def.CEFAppConfig_SetForcedDeviceScaleFactor).Call(uintptr(unsafe.Pointer(&value)))
}

func (m *TCEFApplication) DisableZygote() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableZygote).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableZygote(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableZygote).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UseMockKeyChain() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UseMockKeyChain).Call()
	return api.GoBool(r1)
}

// SetUseMockKeyChain Uses mock keychain for testing purposes, which prevents blocking dialogs from causing timeouts.
// <para><see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --use-mock-keychain</see></para>
func (m *TCEFApplication) SetUseMockKeyChain(value bool) {
	imports.Proc(def.CEFAppConfig_SetUseMockKeyChain).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableRequestHandlingForTesting() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableRequestHandlingForTesting).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableRequestHandlingForTesting(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableRequestHandlingForTesting).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisablePopupBlocking() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisablePopupBlocking).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisablePopupBlocking(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisablePopupBlocking).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableBackForwardCache() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableBackForwardCache).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableBackForwardCache(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableBackForwardCache).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableComponentUpdate() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DisableComponentUpdate).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableComponentUpdate(value bool) {
	imports.Proc(def.CEFAppConfig_SetDisableComponentUpdate).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowInsecureLocalhost() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AllowInsecureLocalhost).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowInsecureLocalhost(value bool) {
	imports.Proc(def.CEFAppConfig_SetAllowInsecureLocalhost).Call(api.PascalBool(value))
}

func (m *TCEFApplication) KioskPrinting() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_KioskPrinting).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetKioskPrinting(value bool) {
	imports.Proc(def.CEFAppConfig_SetKioskPrinting).Call(api.PascalBool(value))
}

func (m *TCEFApplication) TreatInsecureOriginAsSecure() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_TreatInsecureOriginAsSecure).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetTreatInsecureOriginAsSecure(value string) {
	imports.Proc(def.CEFAppConfig_SetTreatInsecureOriginAsSecure).Call(api.PascalStr(value))
}

func (m *TCEFApplication) NetLogEnabled() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_NetLogEnabled).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetNetLogEnabled(value bool) {
	imports.Proc(def.CEFAppConfig_SetNetLogEnabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) NetLogFile() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_NetLogFile).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetNetLogFile(value string) {
	imports.Proc(def.CEFAppConfig_SetNetLogFile).Call(api.PascalStr(value))
}

func (m *TCEFApplication) NetLogCaptureMode() TCefNetLogCaptureMode {
	r1, _, _ := imports.Proc(def.CEFAppConfig_NetLogCaptureMode).Call()
	return TCefNetLogCaptureMode(r1)
}

func (m *TCEFApplication) SetNetLogCaptureMode(value TCefNetLogCaptureMode) {
	imports.Proc(def.CEFAppConfig_SetNetLogCaptureMode).Call(value.ToPtr())
}

func (m *TCEFApplication) RemoteAllowOrigins() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_RemoteAllowOrigins).Call(GetValue)
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetRemoteAllowOrigins(value string) {
	imports.Proc(def.CEFAppConfig_RemoteAllowOrigins).Call(SetValue, api.PascalStr(value))
}

func (m *TCEFApplication) AutoAcceptCamAndMicCapture() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AutoAcceptCamAndMicCapture).Call(GetValue)
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAutoAcceptCamAndMicCapture(value bool) {
	imports.Proc(def.CEFAppConfig_AutoAcceptCamAndMicCapture).Call(SetValue, api.PascalBool(value))
}

func (m *TCEFApplication) UIColorMode() TCefUIColorMode {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UIColorMode).Call(GetValue)
	return TCefUIColorMode(r1)
}

func (m *TCEFApplication) SetUIColorMode(value TCefUIColorMode) {
	imports.Proc(def.CEFAppConfig_UIColorMode).Call(SetValue, uintptr(value))
}

// EnableHighDPISupport
//
//	CEF 112 Remove
func (m *TCEFApplication) EnableHighDPISupport() bool {
	if common.IsWindows() {
		r1, _, _ := imports.Proc(def.CEFAppConfig_EnableHighDPISupport).Call()
		return api.GoBool(r1)
	}
	return false
}

// SetEnableHighDPISupport
//
//	CEF 112 Remove
func (m *TCEFApplication) SetEnableHighDPISupport(value bool) {
	if common.IsWindows() {
		imports.Proc(def.CEFAppConfig_SetEnableHighDPISupport).Call(api.PascalBool(value))
	}
}

/*** 自定义属性 ***/

func (m *TCEFApplication) DeleteCache() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DeleteCache).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDeleteCache(value bool) {
	imports.Proc(def.CEFAppConfig_SetDeleteCache).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DeleteCookies() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_DeleteCookies).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDeleteCookies(value bool) {
	imports.Proc(def.CEFAppConfig_SetDeleteCookies).Call(api.PascalBool(value))
}

func (m *TCEFApplication) CheckCEFFiles() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_CheckCEFFiles).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetCheckCEFFiles(value bool) {
	imports.Proc(def.CEFAppConfig_SetCheckCEFFiles).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ShowMessageDlg() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ShowMessageDlg).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetShowMessageDlg(value bool) {
	imports.Proc(def.CEFAppConfig_SetShowMessageDlg).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MissingBinariesException() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MissingBinariesException).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMissingBinariesException(value bool) {
	imports.Proc(def.CEFAppConfig_SetMissingBinariesException).Call(api.PascalBool(value))
}

func (m *TCEFApplication) SetCurrentDir() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_SetCurrentDir).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetSetCurrentDir(value bool) {
	imports.Proc(def.CEFAppConfig_SetSetCurrentDir).Call(api.PascalBool(value))
}

func (m *TCEFApplication) GlobalContextInitialized() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_GlobalContextInitialized).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) ChromeMajorVer() uint16 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeMajorVer).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeMinorVer() uint16 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeMinorVer).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeRelease() uint16 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeRelease).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeBuild() uint16 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeBuild).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeVersion() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeVersion).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LibCefVersion() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LibCefVersion).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LibCefPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LibCefPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) ChromeElfPath() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChromeElfPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LibLoaded() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LibLoaded).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) LogProcessInfo() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LogProcessInfo).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetLogProcessInfo(value bool) {
	imports.Proc(def.CEFAppConfig_SetLogProcessInfo).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ReRaiseExceptions() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ReRaiseExceptions).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetReRaiseExceptions(value bool) {
	imports.Proc(def.CEFAppConfig_SetReRaiseExceptions).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DeviceScaleFactor() (result float32) {
	imports.Proc(def.CEFAppConfig_DeviceScaleFactor).Call(uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFApplication) LocalesRequired() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LocalesRequired).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLocalesRequired(value string) {
	imports.Proc(def.CEFAppConfig_SetLocalesRequired).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ProcessType() TCefProcessType {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ProcessType).Call()
	return TCefProcessType(r1)
}

func (m *TCEFApplication) ProcessTypeValue() (processTypeValue TCefProcessTypeValue) {
	switch m.ProcessType() {
	case PtBrowser:
		processTypeValue = PtvBrowser
	case PtRenderer:
		processTypeValue = PtvRenderer
	case PtZygote:
		processTypeValue = PtvZygote
	case PtGPU:
		processTypeValue = PtvGPU
	case PtUtility:
		processTypeValue = PtvUtility
	case PtBroker:
		processTypeValue = PtvBroker
	case PtCrashpad:
		processTypeValue = PtvCrashpad
	case PtOther:
		processTypeValue = PtvOther
	default:
		processTypeValue = ""
	}
	return
}

func (m *TCEFApplication) MustCreateResourceBundleHandler() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MustCreateResourceBundleHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateResourceBundleHandler(value bool) {
	imports.Proc(def.CEFAppConfig_SetMustCreateResourceBundleHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MustCreateBrowserProcessHandler() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MustCreateBrowserProcessHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateBrowserProcessHandler(value bool) {
	imports.Proc(def.CEFAppConfig_SetMustCreateBrowserProcessHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MustCreateRenderProcessHandler() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MustCreateRenderProcessHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateRenderProcessHandler(value bool) {
	imports.Proc(def.CEFAppConfig_SetMustCreateRenderProcessHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MustCreateLoadHandler() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MustCreateLoadHandler).Call()
	return api.GoBool(r1)
}

// By Windows:
// Set to true (1) before calling Windows APIs like TrackPopupMenu that enter a
// modal message loop. Set to false (0) after exiting the modal message loop.
func (m *TCEFApplication) OsmodalLoop(value bool) {
	imports.Proc(def.CEFAppConfig_OsmodalLoop).Call(api.PascalBool(value))
}

func (m *TCEFApplication) SetMustCreateLoadHandler(value bool) {
	imports.Proc(def.CEFAppConfig_SetMustCreateLoadHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) Status() TCefApplicationStatus {
	// type = TCefAplicationStatus = asLoading asLoaded asInitialized asShuttingDown asUnloaded asErrorMissingFiles asErrorDLLVersion asErrorLoadingLibrary asErrorInitializingLibrary asErrorExecutingProcess
	r1, _, _ := imports.Proc(def.CEFAppConfig_Status).Call()
	return TCefApplicationStatus(r1)
}

func (m *TCEFApplication) MissingLibFiles() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MissingLibFiles).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) MustFreeLibrary() bool {
	r1, _, _ := imports.Proc(def.CEFAppConfig_MustFreeLibrary).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustFreeLibrary(value bool) {
	imports.Proc(def.CEFAppConfig_SetMustFreeLibrary).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UpdateDeviceScaleFactor() {
	imports.Proc(def.CEFAppConfig_UpdateDeviceScaleFactor).Call()
}

func (m *TCEFApplication) ChildProcessesCount() int32 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ChildProcessesCount).Call()
	return int32(r1)
}

func (m *TCEFApplication) UsedMemory() uint64 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_UsedMemory).Call()
	return uint64(r1)
}

func (m *TCEFApplication) TotalSystemMemory() uint64 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_TotalSystemMemory).Call()
	return uint64(r1)
}

func (m *TCEFApplication) AvailableSystemMemory() uint64 {
	r1, _, _ := imports.Proc(def.CEFAppConfig_AvailableSystemMemory).Call()
	return uint64(r1)
}

func (m *TCEFApplication) SystemMemoryLoad() types.Cardinal {
	r1, _, _ := imports.Proc(def.CEFAppConfig_SystemMemoryLoad).Call()
	return types.Cardinal(r1)
}

func (m *TCEFApplication) ApiHashUniversal() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ApiHashUniversal).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) ApiHashPlatform() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ApiHashPlatform).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) ApiHashCommit() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_ApiHashCommit).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LastErrorMessage() string {
	r1, _, _ := imports.Proc(def.CEFAppConfig_LastErrorMessage).Call()
	return api.GoStr(r1)
}

// SpecificVersion CEF特定版本
func (m *TCEFApplication) SpecificVersion() SpecificVersion {
	if m.specificVersion == 0 {
		cefVersion := strings.Split(m.LibCefVersion(), ".")
		major := common.StrToInt32(cefVersion[0])
		switch SpecificVersion(major) {
		case Sv49, Sv87, Sv101, Sv109:
			m.specificVersion = SpecificVersion(major)
		default:
			m.specificVersion = SvLatest
		}
	}
	return m.specificVersion
}

// IsLatest The current version or the latest version
func (m *TCEFApplication) IsLatest() bool {
	return m.SpecificVersion() == SvLatest
}

// Is49 WindowsXP
func (m *TCEFApplication) Is49() bool {
	return m.SpecificVersion() == Sv49
}

// Is87 Flash
func (m *TCEFApplication) Is87() bool {
	return m.SpecificVersion() == Sv87
}

// Is101 Linux 32
func (m *TCEFApplication) Is101() bool {
	return m.SpecificVersion() == Sv101
}

// Is109  7, 8/8.1 and Windows Server 2012
func (m *TCEFApplication) Is109() bool {
	return m.SpecificVersion() == Sv109
}
