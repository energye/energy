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
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/types"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl/api"
	"os"
	"path"
	"unsafe"
)

// SetObjectRootName
//
// GO绑定JS对象类型所属对象名定义
func (m *TCEFApplication) SetObjectRootName(name string) {
	if name == "" {
		//默认值
		name = internalObjectRootName
	} else {
		//自定义
		internalObjectRootName = name
	}
	//imports.Proc(internale_CEFV8ValueRef_SetObjectRootName).Call(api.PascalStr(internalObjectRootName))
}

//initDefaultSettings 初始 energy 默认设置
func (m *TCEFApplication) initDefaultSettings() {
	lp := libPath()
	if lp != "" {
		m.SetFrameworkDirPath(lp)
		m.SetResourcesDirPath(lp)
		m.SetLocalesDirPath(path.Join(lp, "locales"))
	}
	m.SetAcceptLanguageList(string(LANGUAGE_zh_CN))
	m.SetLocale(string(LANGUAGE_zh_CN))
	m.SetLogSeverity(LOGSEVERITY_DISABLE)
	// 以下条件判断根据不同平台, 启动不同的窗口组件
	// ViewsFrameworkBrowserWindow 简称(VF)窗口组件, 同时支持 Windows/Linux/MacOSX
	// LCL 窗口组件,同时支持 Windows/MacOSX, CEF版本<=106.xx时支持GTK2, CEF版本 >= 107.xx时默认开启 GTK3 且不支持 GTK2 和 LCL提供的各种组件
	if common.IsLinux() { // (VF)View Framework 窗口
		// Linux CEF >= 107.xxx 版本以后，默认启用的GTK3，106及以前版本默认支持GTK2但无法正常输入中文
		// 强制使用GTK3方式，但又无法正常创建lcl组件到窗口中，该框架只对浏览器应用做封装
		// 所以初衷以浏览器应用建设为目标
		// Linux平台默认设置为false,将启用 ViewsFrameworkBrowserWindow 窗口
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(false)
		// 这是一个解决“GPU不可用错误”问题的方法 linux
		// https://bitbucket.org/chromiumembedded/cef/issues/2964/gpu-is-not-usable-error-during-cef
		m.SetDisableZygote(true)
	} else if common.IsDarwin() { // LCL窗口
		m.AddCrDelegate()
		// MacOSX 在使用LCL窗口组件必须将 ExternalMessagePump=true 和 MultiThreadedMessageLoop=false
		// 或
		// 同 Linux 一样使用 ViewsFrameworkBrowserWindow 窗口组件
		m.SetExternalMessagePump(true)
		m.SetMultiThreadedMessageLoop(false)
	} else { // LCL窗口
		//Windows
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(true)
	}
}

func libCef() string {
	if common.IsWindows() {
		return "libcef.dll"
	} else if common.IsLinux() {
		return "libcef.so"
	}
	return ""
}

func libPath() string {
	var lib = libCef()
	if lib != "" {
		//当前目录
		if tools.IsExist(ExePath + Separator + lib) {
			return ExePath
		}
		//环境变量
		var env = os.Getenv(ENERGY_HOME_KEY)
		if tools.IsExist(env + Separator + lib) {
			return env
		}
	}
	return ""
}

/*** 设置 TCefSettings (cef_settings_t) 属性 ***/

func (m *TCEFApplication) NoSandbox() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_NoSandbox).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetNoSandbox(value bool) {
	imports.Proc(internale_CEFAppConfig_SetNoSandbox).Call(api.PascalBool(value))
}

func (m *TCEFApplication) BrowserSubprocessPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_BrowserSubprocessPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetBrowserSubprocessPath(value string) {
	imports.Proc(internale_CEFAppConfig_SetBrowserSubprocessPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) FrameworkDirPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_FrameworkDirPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetFrameworkDirPath(value string) {
	imports.Proc(internale_CEFAppConfig_SetFrameworkDirPath).Call(api.PascalStr(value))
}

// MainBundlePath 仅用于macOS
func (m *TCEFApplication) MainBundlePath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MainBundlePath).Call()
	return api.GoStr(r1)
}

// SetMainBundlePath 仅用于macOS
func (m *TCEFApplication) SetMainBundlePath(value string) {
	imports.Proc(internale_CEFAppConfig_SetMainBundlePath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ChromeRuntime() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeRuntime).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetChromeRuntime(value bool) {
	imports.Proc(internale_CEFAppConfig_SetChromeRuntime).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MultiThreadedMessageLoop() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MultiThreadedMessageLoop).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMultiThreadedMessageLoop(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMultiThreadedMessageLoop).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ExternalMessagePump() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ExternalMessagePump).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetExternalMessagePump(value bool) {
	imports.Proc(internale_CEFAppConfig_SetExternalMessagePump).Call(api.PascalBool(value))
}

func (m *TCEFApplication) WindowlessRenderingEnabled() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_WindowlessRenderingEnabled).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetWindowlessRenderingEnabled(value bool) {
	imports.Proc(internale_CEFAppConfig_SetWindowlessRenderingEnabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) CommandLineArgsDisabled() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_CommandLineArgsDisabled).Call()
	return api.GoBool(r1)
}

// SetCommandLineArgsDisabled 开启/禁用设置命令行参数
func (m *TCEFApplication) SetCommandLineArgsDisabled(value bool) {
	imports.Proc(internale_CEFAppConfig_SetCommandLineArgsDisabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) Cache() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_Cache).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetCache(value string) {
	imports.Proc(internale_CEFAppConfig_SetCache).Call(api.PascalStr(value))
}

func (m *TCEFApplication) RootCache() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_RootCache).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetRootCache(value string) {
	imports.Proc(internale_CEFAppConfig_SetRootCache).Call(api.PascalStr(value))
}

func (m *TCEFApplication) UserDataPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UserDataPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetUserDataPath(value string) {
	imports.Proc(internale_CEFAppConfig_SetUserDataPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) PersistSessionCookies() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_PersistSessionCookies).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetPersistSessionCookies(value bool) {
	imports.Proc(internale_CEFAppConfig_SetPersistSessionCookies).Call(api.PascalBool(value))
}

func (m *TCEFApplication) PersistUserPreferences() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_PersistUserPreferences).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetPersistUserPreferences(value bool) {
	imports.Proc(internale_CEFAppConfig_SetPersistUserPreferences).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UserAgent() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UserAgent).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetUserAgent(value string) {
	imports.Proc(internale_CEFAppConfig_SetUserAgent).Call(api.PascalStr(value))
}

func (m *TCEFApplication) UserAgentProduct() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UserAgentProduct).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetUserAgentProduct(value string) {
	imports.Proc(internale_CEFAppConfig_SetUserAgentProduct).Call(api.PascalStr(value))
}

func (m *TCEFApplication) Locale() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_Locale).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLocale(value string) {
	imports.Proc(internale_CEFAppConfig_SetLocale).Call(api.PascalStr(value))
}

func (m *TCEFApplication) LogFile() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LogFile).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLogFile(value string) {
	imports.Proc(internale_CEFAppConfig_SetLogFile).Call(api.PascalStr(value))
}

func (m *TCEFApplication) LogSeverity() LogSeverity {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LogSeverity).Call()
	return LogSeverity(r1)
}

func (m *TCEFApplication) SetLogSeverity(value LogSeverity) {
	imports.Proc(internale_CEFAppConfig_SetLogSeverity).Call(value.ToPtr())
}

func (m *TCEFApplication) JavaScriptFlags() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_JavaScriptFlags).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetJavaScriptFlags(value string) {
	imports.Proc(internale_CEFAppConfig_SetJavaScriptFlags).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ResourcesDirPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ResourcesDirPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetResourcesDirPath(value string) {
	imports.Proc(internale_CEFAppConfig_SetResourcesDirPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) LocalesDirPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LocalesDirPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLocalesDirPath(value string) {
	imports.Proc(internale_CEFAppConfig_SetLocalesDirPath).Call(api.PascalStr(value))
}

func (m *TCEFApplication) PackLoadingDisabled() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_PackLoadingDisabled).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetPackLoadingDisabled(value bool) {
	imports.Proc(internale_CEFAppConfig_SetPackLoadingDisabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) RemoteDebuggingPort() int32 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_RemoteDebuggingPort).Call()
	return int32(r1)
}

func (m *TCEFApplication) SetRemoteDebuggingPort(value int32) {
	if value > 1024 && value < 65535 {
		imports.Proc(internale_CEFAppConfig_SetRemoteDebuggingPort).Call(uintptr(value))
	}
}

func (m *TCEFApplication) UncaughtExceptionStackSize() int32 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UncaughtExceptionStackSize).Call()
	return int32(r1)
}

func (m *TCEFApplication) SetUncaughtExceptionStackSize(value int32) {
	imports.Proc(internale_CEFAppConfig_SetUncaughtExceptionStackSize).Call(uintptr(value))
}

func (m *TCEFApplication) IgnoreCertificateErrors() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_IgnoreCertificateErrors).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetIgnoreCertificateErrors(value bool) {
	imports.Proc(internale_CEFAppConfig_SetIgnoreCertificateErrors).Call(api.PascalBool(value))
}

func (m *TCEFApplication) BackgroundColor() types.TCefColor {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_BackgroundColor).Call()
	return types.TCefColor(r1)
}

func (m *TCEFApplication) SetBackgroundColor(value types.TCefColor) {
	imports.Proc(internale_CEFAppConfig_SetBackgroundColor).Call(value.ToPtr())
}

func (m *TCEFApplication) AcceptLanguageList() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AcceptLanguageList).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetAcceptLanguageList(value string) {
	imports.Proc(internale_CEFAppConfig_SetAcceptLanguageList).Call(api.PascalStr(value))
}

func (m *TCEFApplication) CookieableSchemesList() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_CookieableSchemesList).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetCookieableSchemesList(value string) {
	imports.Proc(internale_CEFAppConfig_SetCookieableSchemesList).Call(api.PascalStr(value))
}

func (m *TCEFApplication) CookieableSchemesExcludeDefaults() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_CookieableSchemesExcludeDefaults).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetCookieableSchemesExcludeDefaults(value bool) {
	imports.Proc(internale_CEFAppConfig_SetCookieableSchemesExcludeDefaults).Call(api.PascalBool(value))
}

/*** 设置常用的命令行参数属性 ***/

func (m *TCEFApplication) SingleProcess() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_SingleProcess).Call()
	SingleProcess = api.GoBool(r1)
	return SingleProcess
}

func (m *TCEFApplication) SetSingleProcess(value bool) {
	SingleProcess = value
	imports.Proc(internale_CEFAppConfig_SetSingleProcess).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableMediaStream() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableMediaStream).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableMediaStream(value bool) {
	imports.Proc(internale_CEFAppConfig_SetEnableMediaStream).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableSpeechInput() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableSpeechInput).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableSpeechInput(value bool) {
	imports.Proc(internale_CEFAppConfig_SetEnableSpeechInput).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UseFakeUIForMediaStream() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UseFakeUIForMediaStream).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetUseFakeUIForMediaStream(value bool) {
	imports.Proc(internale_CEFAppConfig_SetUseFakeUIForMediaStream).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableUsermediaScreenCapturing() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableUsermediaScreenCapturing).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableUsermediaScreenCapturing(value bool) {
	imports.Proc(internale_CEFAppConfig_SetEnableUsermediaScreenCapturing).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableGPU() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableGPU).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableGPU(value bool) {
	imports.Proc(internale_CEFAppConfig_SetEnableGPU).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableFeatures() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetEnableFeatures(value string) {
	imports.Proc(internale_CEFAppConfig_SetEnableFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) DisableFeatures() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDisableFeatures(value string) {
	imports.Proc(internale_CEFAppConfig_SetDisableFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) EnableBlinkFeatures() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableBlinkFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetEnableBlinkFeatures(value string) {
	imports.Proc(internale_CEFAppConfig_SetEnableBlinkFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) DisableBlinkFeatures() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableBlinkFeatures).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDisableBlinkFeatures(value string) {
	imports.Proc(internale_CEFAppConfig_SetDisableBlinkFeatures).Call(api.PascalStr(value))
}

func (m *TCEFApplication) BlinkSettings() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_BlinkSettings).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetBlinkSettings(value string) {
	imports.Proc(internale_CEFAppConfig_SetBlinkSettings).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ForceFieldTrials() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ForceFieldTrials).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetForceFieldTrials(value string) {
	imports.Proc(internale_CEFAppConfig_SetForceFieldTrials).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ForceFieldTrialParams() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ForceFieldTrialParams).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetForceFieldTrialParams(value string) {
	imports.Proc(internale_CEFAppConfig_SetForceFieldTrialParams).Call(api.PascalStr(value))
}

func (m *TCEFApplication) SmoothScrolling() TCefState {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_SmoothScrolling).Call()
	return TCefState(r1)
}

func (m *TCEFApplication) SetSmoothScrolling(value TCefState) {
	imports.Proc(internale_CEFAppConfig_SetSmoothScrolling).Call(value.ToPtr())
}

func (m *TCEFApplication) FastUnload() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_FastUnload).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetFastUnload(value bool) {
	imports.Proc(internale_CEFAppConfig_SetFastUnload).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableSafeBrowsing() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableSafeBrowsing).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableSafeBrowsing(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableSafeBrowsing).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MuteAudio() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MuteAudio).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMuteAudio(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMuteAudio).Call(api.PascalBool(value))
}

func (m *TCEFApplication) SitePerProcess() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_SitePerProcess).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetSitePerProcess(value bool) {
	imports.Proc(internale_CEFAppConfig_SetSitePerProcess).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableWebSecurity() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableWebSecurity).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableWebSecurity(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableWebSecurity).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisablePDFExtension() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisablePDFExtension).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisablePDFExtension(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisablePDFExtension).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableSiteIsolationTrials() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableSiteIsolationTrials).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableSiteIsolationTrials(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableSiteIsolationTrials).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableChromeLoginPrompt() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableChromeLoginPrompt).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableChromeLoginPrompt(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableChromeLoginPrompt).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableExtensions() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableExtensions).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableExtensions(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableExtensions).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AutoplayPolicy() TCefAutoplayPolicy {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AutoplayPolicy).Call()
	return TCefAutoplayPolicy(r1)
}

func (m *TCEFApplication) SetAutoplayPolicy(value TCefAutoplayPolicy) {
	imports.Proc(internale_CEFAppConfig_SetAutoplayPolicy).Call(value.ToPtr())
}

func (m *TCEFApplication) DisableBackgroundNetworking() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableBackgroundNetworking).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableBackgroundNetworking(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableBackgroundNetworking).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MetricsRecordingOnly() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MetricsRecordingOnly).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMetricsRecordingOnly(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMetricsRecordingOnly).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowFileAccessFromFiles() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AllowFileAccessFromFiles).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowFileAccessFromFiles(value bool) {
	imports.Proc(internale_CEFAppConfig_SetAllowFileAccessFromFiles).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowRunningInsecureContent() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AllowRunningInsecureContent).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowRunningInsecureContent(value bool) {
	imports.Proc(internale_CEFAppConfig_SetAllowRunningInsecureContent).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnablePrintPreview() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnablePrintPreview).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnablePrintPreview(value bool) {
	imports.Proc(internale_CEFAppConfig_SetEnablePrintPreview).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DefaultEncoding() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DefaultEncoding).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDefaultEncoding(value string) {
	imports.Proc(internale_CEFAppConfig_SetDefaultEncoding).Call(api.PascalStr(value))
}

func (m *TCEFApplication) DisableJavascript() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableJavascript).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascript(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableJavascript).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableJavascriptCloseWindows() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableJavascriptCloseWindows).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascriptCloseWindows(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableJavascriptCloseWindows).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableJavascriptAccessClipboard() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableJavascriptAccessClipboard).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascriptAccessClipboard(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableJavascriptAccessClipboard).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableJavascriptDomPaste() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableJavascriptDomPaste).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableJavascriptDomPaste(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableJavascriptDomPaste).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowUniversalAccessFromFileUrls() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AllowUniversalAccessFromFileUrls).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowUniversalAccessFromFileUrls(value bool) {
	imports.Proc(internale_CEFAppConfig_SetAllowUniversalAccessFromFileUrls).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableImageLoading() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableImageLoading).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableImageLoading(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableImageLoading).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ImageShrinkStandaloneToFit() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ImageShrinkStandaloneToFit).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetImageShrinkStandaloneToFit(value bool) {
	imports.Proc(internale_CEFAppConfig_SetImageShrinkStandaloneToFit).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableTextAreaResize() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableTextAreaResize).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableTextAreaResize(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableTextAreaResize).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableTabToLinks() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableTabToLinks).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableTabToLinks(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableTabToLinks).Call(api.PascalBool(value))
}

func (m *TCEFApplication) EnableProfanityFilter() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableProfanityFilter).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetEnableProfanityFilter(value bool) {
	imports.Proc(internale_CEFAppConfig_SetEnableProfanityFilter).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableSpellChecking() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableSpellChecking).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableSpellChecking(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableSpellChecking).Call(api.PascalBool(value))
}

func (m *TCEFApplication) OverrideSpellCheckLang() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_OverrideSpellCheckLang).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetOverrideSpellCheckLang(value string) {
	imports.Proc(internale_CEFAppConfig_SetOverrideSpellCheckLang).Call(api.PascalStr(value))
}

func (m *TCEFApplication) TouchEvents() TCefState {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_TouchEvents).Call()
	return TCefState(r1)
}

func (m *TCEFApplication) SetTouchEvents(value TCefState) {
	imports.Proc(internale_CEFAppConfig_SetTouchEvents).Call(value.ToPtr())
}

func (m *TCEFApplication) DisableReadingFromCanvas() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableReadingFromCanvas).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableReadingFromCanvas(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableReadingFromCanvas).Call(api.PascalBool(value))
}

func (m *TCEFApplication) HyperlinkAuditing() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_HyperlinkAuditing).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetHyperlinkAuditing(value bool) {
	imports.Proc(internale_CEFAppConfig_SetHyperlinkAuditing).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableNewBrowserInfoTimeout() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableNewBrowserInfoTimeout).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableNewBrowserInfoTimeout(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableNewBrowserInfoTimeout).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DevToolsProtocolLogFile() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DevToolsProtocolLogFile).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetDevToolsProtocolLogFile(value string) {
	imports.Proc(internale_CEFAppConfig_SetDevToolsProtocolLogFile).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ForcedDeviceScaleFactor() float32 { //single
	var result uintptr
	imports.Proc(internale_CEFAppConfig_ForcedDeviceScaleFactor).Call(uintptr(unsafe.Pointer(&result)))
	return *(*float32)(unsafe.Pointer(result))
}

func (m *TCEFApplication) SetForcedDeviceScaleFactor(value float32) { //single
	imports.Proc(internale_CEFAppConfig_SetForcedDeviceScaleFactor).Call(uintptr(unsafe.Pointer(&value)))
}

func (m *TCEFApplication) DisableZygote() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableZygote).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableZygote(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableZygote).Call(api.PascalBool(value))
}

func (m *TCEFApplication) UseMockKeyChain() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UseMockKeyChain).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetUseMockKeyChain(value bool) {
	imports.Proc(internale_CEFAppConfig_SetUseMockKeyChain).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableRequestHandlingForTesting() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableRequestHandlingForTesting).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableRequestHandlingForTesting(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableRequestHandlingForTesting).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisablePopupBlocking() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisablePopupBlocking).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisablePopupBlocking(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisablePopupBlocking).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableBackForwardCache() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableBackForwardCache).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableBackForwardCache(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableBackForwardCache).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DisableComponentUpdate() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DisableComponentUpdate).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDisableComponentUpdate(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDisableComponentUpdate).Call(api.PascalBool(value))
}

func (m *TCEFApplication) AllowInsecureLocalhost() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AllowInsecureLocalhost).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetAllowInsecureLocalhost(value bool) {
	imports.Proc(internale_CEFAppConfig_SetAllowInsecureLocalhost).Call(api.PascalBool(value))
}

func (m *TCEFApplication) KioskPrinting() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_KioskPrinting).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetKioskPrinting(value bool) {
	imports.Proc(internale_CEFAppConfig_SetKioskPrinting).Call(api.PascalBool(value))
}

func (m *TCEFApplication) TreatInsecureOriginAsSecure() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_TreatInsecureOriginAsSecure).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetTreatInsecureOriginAsSecure(value string) {
	imports.Proc(internale_CEFAppConfig_SetTreatInsecureOriginAsSecure).Call(api.PascalStr(value))
}

func (m *TCEFApplication) NetLogEnabled() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_NetLogEnabled).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetNetLogEnabled(value bool) {
	imports.Proc(internale_CEFAppConfig_SetNetLogEnabled).Call(api.PascalBool(value))
}

func (m *TCEFApplication) NetLogFile() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_NetLogFile).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetNetLogFile(value string) {
	imports.Proc(internale_CEFAppConfig_SetNetLogFile).Call(api.PascalStr(value))
}

func (m *TCEFApplication) NetLogCaptureMode() TCefNetLogCaptureMode {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_NetLogCaptureMode).Call()
	return TCefNetLogCaptureMode(r1)
}

func (m *TCEFApplication) SetNetLogCaptureMode(value TCefNetLogCaptureMode) {
	imports.Proc(internale_CEFAppConfig_SetNetLogCaptureMode).Call(value.ToPtr())
}

func (m *TCEFApplication) EnableHighDPISupport() bool {
	if common.IsWindows() {
		r1, _, _ := imports.Proc(internale_CEFAppConfig_EnableHighDPISupport).Call()
		return api.GoBool(r1)
	}
	return false
}

func (m *TCEFApplication) SetEnableHighDPISupport(value bool) {
	if common.IsWindows() {
		imports.Proc(internale_CEFAppConfig_SetEnableHighDPISupport).Call(api.PascalBool(value))
	}
}

/*** 自定义属性 ***/

func (m *TCEFApplication) DeleteCache() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DeleteCache).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDeleteCache(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDeleteCache).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DeleteCookies() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_DeleteCookies).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetDeleteCookies(value bool) {
	imports.Proc(internale_CEFAppConfig_SetDeleteCookies).Call(api.PascalBool(value))
}

func (m *TCEFApplication) CheckCEFFiles() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_CheckCEFFiles).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetCheckCEFFiles(value bool) {
	imports.Proc(internale_CEFAppConfig_SetCheckCEFFiles).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ShowMessageDlg() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ShowMessageDlg).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetShowMessageDlg(value bool) {
	imports.Proc(internale_CEFAppConfig_SetShowMessageDlg).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MissingBinariesException() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MissingBinariesException).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMissingBinariesException(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMissingBinariesException).Call(api.PascalBool(value))
}

func (m *TCEFApplication) SetCurrentDir() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_SetCurrentDir).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetSetCurrentDir(value bool) {
	imports.Proc(internale_CEFAppConfig_SetSetCurrentDir).Call(api.PascalBool(value))
}

func (m *TCEFApplication) GlobalContextInitialized() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_GlobalContextInitialized).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) ChromeMajorVer() uint16 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeMajorVer).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeMinorVer() uint16 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeMinorVer).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeRelease() uint16 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeRelease).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeBuild() uint16 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeBuild).Call()
	return uint16(r1)
}

func (m *TCEFApplication) ChromeVersion() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeVersion).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LibCefVersion() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LibCefVersion).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LibCefPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LibCefPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) ChromeElfPath() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChromeElfPath).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LibLoaded() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LibLoaded).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) LogProcessInfo() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LogProcessInfo).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetLogProcessInfo(value bool) {
	imports.Proc(internale_CEFAppConfig_SetLogProcessInfo).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ReRaiseExceptions() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ReRaiseExceptions).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetReRaiseExceptions(value bool) {
	imports.Proc(internale_CEFAppConfig_SetReRaiseExceptions).Call(api.PascalBool(value))
}

func (m *TCEFApplication) DeviceScaleFactor() float32 {
	var result uintptr
	imports.Proc(internale_CEFAppConfig_DeviceScaleFactor).Call(uintptr(unsafe.Pointer(&result)))
	return *(*float32)(unsafe.Pointer(result))
}

func (m *TCEFApplication) LocalesRequired() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LocalesRequired).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) SetLocalesRequired(value string) {
	imports.Proc(internale_CEFAppConfig_SetLocalesRequired).Call(api.PascalStr(value))
}

func (m *TCEFApplication) ProcessType() TCefProcessType {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ProcessType).Call()
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
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MustCreateResourceBundleHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateResourceBundleHandler(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMustCreateResourceBundleHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MustCreateBrowserProcessHandler() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MustCreateBrowserProcessHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateBrowserProcessHandler(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMustCreateBrowserProcessHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MustCreateRenderProcessHandler() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MustCreateRenderProcessHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateRenderProcessHandler(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMustCreateRenderProcessHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) MustCreateLoadHandler() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MustCreateLoadHandler).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustCreateLoadHandler(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMustCreateLoadHandler).Call(api.PascalBool(value))
}

func (m *TCEFApplication) Status() TCefApplicationStatus {
	// type = TCefAplicationStatus = asLoading asLoaded asInitialized asShuttingDown asUnloaded asErrorMissingFiles asErrorDLLVersion asErrorLoadingLibrary asErrorInitializingLibrary asErrorExecutingProcess
	r1, _, _ := imports.Proc(internale_CEFAppConfig_Status).Call()
	return TCefApplicationStatus(r1)
}

func (m *TCEFApplication) MissingLibFiles() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MissingLibFiles).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) MustFreeLibrary() bool {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_MustFreeLibrary).Call()
	return api.GoBool(r1)
}

func (m *TCEFApplication) SetMustFreeLibrary(value bool) {
	imports.Proc(internale_CEFAppConfig_SetMustFreeLibrary).Call(api.PascalBool(value))
}

func (m *TCEFApplication) ChildProcessesCount() int32 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ChildProcessesCount).Call()
	return int32(r1)
}

func (m *TCEFApplication) UsedMemory() uint64 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_UsedMemory).Call()
	return uint64(r1)
}

func (m *TCEFApplication) TotalSystemMemory() uint64 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_TotalSystemMemory).Call()
	return uint64(r1)
}

func (m *TCEFApplication) AvailableSystemMemory() uint64 {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_AvailableSystemMemory).Call()
	return uint64(r1)
}

func (m *TCEFApplication) SystemMemoryLoad() types.Cardinal {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_SystemMemoryLoad).Call()
	return types.Cardinal(r1)
}

func (m *TCEFApplication) ApiHashUniversal() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ApiHashUniversal).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) ApiHashPlatform() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ApiHashPlatform).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) ApiHashCommit() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_ApiHashCommit).Call()
	return api.GoStr(r1)
}

func (m *TCEFApplication) LastErrorMessage() string {
	r1, _, _ := imports.Proc(internale_CEFAppConfig_LastErrorMessage).Call()
	return api.GoStr(r1)
}
