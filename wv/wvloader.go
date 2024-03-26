//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IWVLoader Parent: IComponent
//
//	Class used to simplify the WebView2 initialization and destruction.
type IWVLoader interface {
	IComponent
	SetOnEnvironmentCreated(fn TOnLoaderNotifyEvent)
	SetOnInitializationError(fn TOnLoaderNotifyEvent)
	SetOnGetCustomSchemes(fn TOnLoaderGetCustomSchemesEvent)
	SetOnNewBrowserVersionAvailable(fn TOnLoaderNewBrowserVersionAvailableEvent)
	SetOnBrowserProcessExited(fn TOnLoaderBrowserProcessExitedEvent)
	SetOnProcessInfosChanged(fn TOnLoaderProcessInfosChangedEvent)
	// Environment
	//  Represents the global WebView2 Environment.
	Environment() ICoreWebView2Environment // property
	// Status
	//  Returns the TWVLoader initialization status.
	Status() TWV2LoaderStatus // property
	// AvailableBrowserVersion
	//  Get the browser version info including channel name if it is not the
	//  WebView2 Runtime. Channel names are Beta, Dev, and Canary.
	AvailableBrowserVersion() string // property
	// ErrorMessage
	//  Returns all the text appended to the error log with AppendErrorLog.
	ErrorMessage() string // property
	// ErrorCode
	//  Returns the last initialization error code.
	ErrorCode() (resultInt64 int64) // property
	// SetCurrentDir
	//  Used to set the current directory when the WebView2 library is loaded. This is required if the application is launched from a different application.
	SetCurrentDir() bool // property
	// SetSetCurrentDir Set SetCurrentDir
	SetSetCurrentDir(AValue bool) // property
	// Initialized
	//  Returns true if the Status is wvlsInitialized.
	Initialized() bool // property
	// InitializationError
	//  Returns true if the Status is wvlsError.
	InitializationError() bool // property
	// CheckFiles
	//  Checks if the WebView2 library is present and the DLL version.
	CheckFiles() bool // property
	// SetCheckFiles Set CheckFiles
	SetCheckFiles(AValue bool) // property
	// ShowMessageDlg
	//  Set to true when you need to use a showmessage dialog to show the error messages.
	ShowMessageDlg() bool // property
	// SetShowMessageDlg Set ShowMessageDlg
	SetShowMessageDlg(AValue bool) // property
	// InitCOMLibrary
	//  Set to true to call CoInitializeEx and CoUnInitialize in TWVLoader.Create and TWVLoader.Destroy.
	InitCOMLibrary() bool // property
	// SetInitCOMLibrary Set InitCOMLibrary
	SetInitCOMLibrary(AValue bool) // property
	// CustomCommandLineSwitches
	//  Custom command line switches used by TCoreWebView2EnvironmentOptions.Create to initialize WebView2.
	CustomCommandLineSwitches() string // property
	// DeviceScaleFactor
	//  Returns the device scale factor.
	DeviceScaleFactor() float32 // property
	// ReRaiseExceptions
	//  Set to true to raise all exceptions.
	ReRaiseExceptions() bool // property
	// SetReRaiseExceptions Set ReRaiseExceptions
	SetReRaiseExceptions(AValue bool) // property
	// InstalledRuntimeVersion
	//  Returns the installed WebView2 runtime version.
	InstalledRuntimeVersion() string // property
	// LoaderDllPath
	//  Full path to WebView2Loader.dll. Leave empty to load WebView2Loader.dll from the current directory.
	LoaderDllPath() string // property
	// SetLoaderDllPath Set LoaderDllPath
	SetLoaderDllPath(AValue string) // property
	// UseInternalLoader
	//  Use a WebView2Loader.dll replacement based on the OpenWebView2Loader project.
	//  <a href="https://github.com/jchv/OpenWebView2Loader">See the OpenWebView2Loader project repository at GitHub.</a>
	UseInternalLoader() bool // property
	// SetUseInternalLoader Set UseInternalLoader
	SetUseInternalLoader(AValue bool) // property
	// BrowserExecPath
	//  Use BrowserExecPath to specify whether WebView2 controls use a fixed or
	//  installed version of the WebView2 Runtime that exists on a user machine.
	//  To use a fixed version of the WebView2 Runtime, pass the folder path that
	//  contains the fixed version of the WebView2 Runtime to BrowserExecPath.
	//  BrowserExecPath supports both relative(to the application's executable)
	//  and absolute files paths. To create WebView2 controls that use the installed
	//  version of the WebView2 Runtime that exists on user machines,
	//  pass an empty string to BrowserExecPath.
	//  Property used to create the environment. Used as the browserExecutableFolder parameter of CreateCoreWebView2EnvironmentWithOptions.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</a>
	BrowserExecPath() string // property
	// SetBrowserExecPath Set BrowserExecPath
	SetBrowserExecPath(AValue string) // property
	// UserDataFolder
	//  You may specify the userDataFolder to change the default user data folder
	//  location for WebView2. The path is either an absolute file path or a relative
	//  file path that is interpreted as relative to the compiled code for the
	//  current process.
	//  Property used to create the environment. Used as the userDataFolder parameter of CreateCoreWebView2EnvironmentWithOptions.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</a>
	UserDataFolder() string // property
	// SetUserDataFolder Set UserDataFolder
	SetUserDataFolder(AValue string) // property
	// AdditionalBrowserArguments
	//  Additional command line switches.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AdditionalBrowserArguments.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	AdditionalBrowserArguments() string // property
	// SetAdditionalBrowserArguments Set AdditionalBrowserArguments
	SetAdditionalBrowserArguments(AValue string) // property
	// Language
	//  The default display language for WebView. It applies to browser UI such as
	//  context menu and dialogs. It also applies to the `accept-languages` HTTP
	//  header that WebView sends to websites.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_Language.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	Language() string // property
	// SetLanguage Set Language
	SetLanguage(AValue string) // property
	// TargetCompatibleBrowserVersion
	//  Specifies the version of the WebView2 Runtime binaries required to be
	//  compatible with your app.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_TargetCompatibleBrowserVersion.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	TargetCompatibleBrowserVersion() string // property
	// SetTargetCompatibleBrowserVersion Set TargetCompatibleBrowserVersion
	SetTargetCompatibleBrowserVersion(AValue string) // property
	// AllowSingleSignOnUsingOSPrimaryAccount
	//  Used to enable single sign on with Azure Active Directory(AAD) and personal Microsoft
	//  Account(MSA) resources inside WebView.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AllowSingleSignOnUsingOSPrimaryAccount.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	AllowSingleSignOnUsingOSPrimaryAccount() bool // property
	// SetAllowSingleSignOnUsingOSPrimaryAccount Set AllowSingleSignOnUsingOSPrimaryAccount
	SetAllowSingleSignOnUsingOSPrimaryAccount(AValue bool) // property
	// ExclusiveUserDataFolderAccess
	//  Whether other processes can create WebView2 from WebView2Environment created with the
	//  same user data folder and therefore sharing the same WebView browser process instance.
	//  Default is FALSE.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions2.Get_ExclusiveUserDataFolderAccess.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2">See the ICoreWebView2EnvironmentOptions2 article.</a>
	ExclusiveUserDataFolderAccess() bool // property
	// SetExclusiveUserDataFolderAccess Set ExclusiveUserDataFolderAccess
	SetExclusiveUserDataFolderAccess(AValue bool) // property
	// CustomCrashReportingEnabled
	//  When `CustomCrashReportingEnabled` is set to `TRUE`, Windows won't send crash data to Microsoft endpoint.
	//  `CustomCrashReportingEnabled` is default to be `FALSE`, in this case, WebView will respect OS consent.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions3.Get_IsCustomCrashReportingEnabled.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3">See the ICoreWebView2EnvironmentOptions3 article.</a>
	CustomCrashReportingEnabled() bool // property
	// SetCustomCrashReportingEnabled Set CustomCrashReportingEnabled
	SetCustomCrashReportingEnabled(AValue bool) // property
	// EnableTrackingPrevention
	//  The `EnableTrackingPrevention` property is used to enable/disable tracking prevention
	//  feature in WebView2. This property enable/disable tracking prevention for all the
	//  WebView2's created in the same environment.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions5.Get_EnableTrackingPrevention.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5">See the ICoreWebView2EnvironmentOptions5 article.</a>
	EnableTrackingPrevention() bool // property
	// SetEnableTrackingPrevention Set EnableTrackingPrevention
	SetEnableTrackingPrevention(AValue bool) // property
	// AreBrowserExtensionsEnabled
	//  When `AreBrowserExtensionsEnabled` is set to `TRUE`, new extensions can be added to user
	//  profile and used. `AreBrowserExtensionsEnabled` is default to be `FALSE`, in this case,
	//  new extensions can't be installed, and already installed extension won't be
	//  available to use in user profile.
	//  If connecting to an already running environment with a different value for `AreBrowserExtensionsEnabled`
	//  property, it will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions6.Get_AreBrowserExtensionsEnabled.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions6">See the ICoreWebView2EnvironmentOptions6 article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension">See the ICoreWebView2BrowserExtension article for Extensions API details.</a>
	AreBrowserExtensionsEnabled() bool // property
	// SetAreBrowserExtensionsEnabled Set AreBrowserExtensionsEnabled
	SetAreBrowserExtensionsEnabled(AValue bool) // property
	// EnableGPU
	//  Enable GPU hardware acceleration.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu-compositing</a>
	EnableGPU() bool // property
	// SetEnableGPU Set EnableGPU
	SetEnableGPU(AValue bool) // property
	// EnableFeatures
	//  List of feature names to enable.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-features</a>
	//  The list of features you can enable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	EnableFeatures() string // property
	// SetEnableFeatures Set EnableFeatures
	SetEnableFeatures(AValue string) // property
	// DisableFeatures
	//  List of feature names to disable.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-features</a>
	//  The list of features you can disable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	DisableFeatures() string // property
	// SetDisableFeatures Set DisableFeatures
	SetDisableFeatures(AValue string) // property
	// EnableBlinkFeatures
	//  Enable one or more Blink runtime-enabled features.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-blink-features</a>
	//  The list of Blink features you can enable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	EnableBlinkFeatures() string // property
	// SetEnableBlinkFeatures Set EnableBlinkFeatures
	SetEnableBlinkFeatures(AValue string) // property
	// DisableBlinkFeatures
	//  Disable one or more Blink runtime-enabled features.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-blink-features</a>
	//  The list of Blink features you can disable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	DisableBlinkFeatures() string // property
	// SetDisableBlinkFeatures Set DisableBlinkFeatures
	SetDisableBlinkFeatures(AValue string) // property
	// BlinkSettings
	//  Set blink settings. Format is <name>[=<value],<name>[=<value>],...
	//  The names are declared in Settings.json5. For boolean type, use "true", "false",
	//  or omit '=<value>' part to set to true. For enum type, use the int value of the
	//  enum value. Applied after other command line flags and prefs.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --blink-settings</a>
	//  The list of Blink settings you can disable is here:
	//  https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/frame/settings.json5
	BlinkSettings() string // property
	// SetBlinkSettings Set BlinkSettings
	SetBlinkSettings(AValue string) // property
	// ForceFieldTrials
	//  This option can be used to force field trials when testing changes locally.
	//  The argument is a list of name and value pairs, separated by slashes.
	//  If a trial name is prefixed with an asterisk, that trial will start activated.
	//  For example, the following argument defines two trials, with the second one
	//  activated: "GoogleNow/Enable/*MaterialDesignNTP/Default/" This option can also
	//  be used by the browser process to send the list of trials to a non-browser
	//  process, using the same format. See FieldTrialList::CreateTrialsFromString()
	//  in field_trial.h for details.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrials</a>
	//  https://source.chromium.org/chromium/chromium/src/+/master:base/base_switches.cc
	ForceFieldTrials() string // property
	// SetForceFieldTrials Set ForceFieldTrials
	SetForceFieldTrials(AValue string) // property
	// ForceFieldTrialParams
	//  This option can be used to force parameters of field trials when testing
	//  changes locally. The argument is a param list of(key, value) pairs prefixed
	//  by an associated(trial, group) pair. You specify the param list for multiple
	// (trial, group) pairs with a comma separator.
	//  Example: "Trial1.Group1:k1/v1/k2/v2,Trial2.Group2:k3/v3/k4/v4"
	//  Trial names, groups names, parameter names, and value should all be URL
	//  escaped for all non-alphanumeric characters.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrial-params</a>
	//  https://source.chromium.org/chromium/chromium/src/+/master:components/variations/variations_switches.cc
	ForceFieldTrialParams() string // property
	// SetForceFieldTrialParams Set ForceFieldTrialParams
	SetForceFieldTrialParams(AValue string) // property
	// SmartScreenProtectionEnabled
	//  Workaround given my Microsoft to disable the SmartScreen protection.
	SmartScreenProtectionEnabled() bool // property
	// SetSmartScreenProtectionEnabled Set SmartScreenProtectionEnabled
	SetSmartScreenProtectionEnabled(AValue bool) // property
	// AllowInsecureLocalhost
	//  Enables TLS/SSL errors on localhost to be ignored(no interstitial, no blocking of requests).
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-insecure-localhost</a>
	AllowInsecureLocalhost() bool // property
	// SetAllowInsecureLocalhost Set AllowInsecureLocalhost
	SetAllowInsecureLocalhost(AValue bool) // property
	// DisableWebSecurity
	//  Don't enforce the same-origin policy.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-web-security</a>
	DisableWebSecurity() bool // property
	// SetDisableWebSecurity Set DisableWebSecurity
	SetDisableWebSecurity(AValue bool) // property
	// TouchEvents
	//  Enable support for touch event feature detection.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --touch-events</a>
	TouchEvents() TWVState // property
	// SetTouchEvents Set TouchEvents
	SetTouchEvents(AValue TWVState) // property
	// HyperlinkAuditing
	//  Don't send hyperlink auditing pings.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-pings</a>
	HyperlinkAuditing() bool // property
	// SetHyperlinkAuditing Set HyperlinkAuditing
	SetHyperlinkAuditing(AValue bool) // property
	// AutoplayPolicy
	//  Autoplay policy.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --autoplay-policy</a>
	AutoplayPolicy() TWVAutoplayPolicy // property
	// SetAutoplayPolicy Set AutoplayPolicy
	SetAutoplayPolicy(AValue TWVAutoplayPolicy) // property
	// MuteAudio
	//  Mutes audio sent to the audio device so it is not audible during automated testing.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --mute-audio</a>
	MuteAudio() bool // property
	// SetMuteAudio Set MuteAudio
	SetMuteAudio(AValue bool) // property
	// KioskPrinting
	//  Default encoding.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --default-encoding</a>
	//  Enable automatically pressing the print button in print preview.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --kiosk-printing</a>
	KioskPrinting() bool // property
	// SetKioskPrinting Set KioskPrinting
	SetKioskPrinting(AValue bool) // property
	// ProxySettings
	//  Configure the browser to use a proxy server.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</a>
	ProxySettings() IWVProxySettings // property
	// AllowFileAccessFromFiles
	//  By default, file:// URIs cannot read other file:// URIs. This is an override for developers who need the old behavior for testing.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-file-access-from-files</a>
	AllowFileAccessFromFiles() bool // property
	// SetAllowFileAccessFromFiles Set AllowFileAccessFromFiles
	SetAllowFileAccessFromFiles(AValue bool) // property
	// AllowRunningInsecureContent
	//  By default, an https page cannot run JavaScript, CSS or plugins from http URLs. This provides an override to get the old insecure behavior.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-running-insecure-content</a>
	AllowRunningInsecureContent() bool // property
	// SetAllowRunningInsecureContent Set AllowRunningInsecureContent
	SetAllowRunningInsecureContent(AValue bool) // property
	// DisableBackgroundNetworking
	//  Disable several subsystems which run network requests in the background.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-background-networking</a>
	DisableBackgroundNetworking() bool // property
	// SetDisableBackgroundNetworking Set DisableBackgroundNetworking
	SetDisableBackgroundNetworking(AValue bool) // property
	// ForcedDeviceScaleFactor
	//  Overrides the device scale factor for the browser UI and the contents.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-device-scale-factor</a>
	ForcedDeviceScaleFactor() float32 // property
	// SetForcedDeviceScaleFactor Set ForcedDeviceScaleFactor
	SetForcedDeviceScaleFactor(AValue float32) // property
	// RemoteDebuggingPort
	//  Set to a value between 1024 and 65535 to enable remote debugging on the
	//  specified port. Also configurable using the "remote-debugging-port"
	//  command-line switch. Remote debugging can be accessed by loading the
	//  chrome://inspect page in Google Chrome. Port numbers 9222 and 9229 are
	//  discoverable by default. Other port numbers may need to be configured via
	//  "Discover network targets" on the Devices tab.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-debugging-port</a>
	RemoteDebuggingPort() int32 // property
	// SetRemoteDebuggingPort Set RemoteDebuggingPort
	SetRemoteDebuggingPort(AValue int32) // property
	// RemoteAllowOrigins
	//  Enables web socket connections from the specified origins only. '*' allows any origin.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-allow-origins</a>
	RemoteAllowOrigins() string // property
	// SetRemoteAllowOrigins Set RemoteAllowOrigins
	SetRemoteAllowOrigins(AValue string) // property
	// DebugLog
	//  Force logging to be enabled. Logging is disabled by default in release builds.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-logging</a>
	DebugLog() TWV2DebugLog // property
	// SetDebugLog Set DebugLog
	SetDebugLog(AValue TWV2DebugLog) // property
	// DebugLogLevel
	//  Sets the minimum log level.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-level</a>
	DebugLogLevel() TWV2DebugLogLevel // property
	// SetDebugLogLevel Set DebugLogLevel
	SetDebugLogLevel(AValue TWV2DebugLogLevel) // property
	// JavaScriptFlags
	//  Specifies the flags passed to JS engine.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --js-flags</a>
	JavaScriptFlags() string // property
	// SetJavaScriptFlags Set JavaScriptFlags
	SetJavaScriptFlags(AValue string) // property
	// DisableEdgePitchNotification
	//  Workaround given my Microsoft to disable the "Download Edge" notifications.
	DisableEdgePitchNotification() bool // property
	// SetDisableEdgePitchNotification Set DisableEdgePitchNotification
	SetDisableEdgePitchNotification(AValue bool) // property
	// TreatInsecureOriginAsSecure
	//  Treat given(insecure) origins as secure origins.
	//  Multiple origins can be supplied as a comma-separated list.
	//  For the definition of secure contexts, see https://w3c.github.io/webappsec-secure-contexts/
	//  and https://www.w3.org/TR/powerful-features/#is-origin-trustworthy
	//  Example: --unsafely-treat-insecure-origin-as-secure=http://a.test,http://b.test
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --unsafely-treat-insecure-origin-as-secure</a>
	TreatInsecureOriginAsSecure() string // property
	// SetTreatInsecureOriginAsSecure Set TreatInsecureOriginAsSecure
	SetTreatInsecureOriginAsSecure(AValue string) // property
	// AutoAcceptCamAndMicCapture
	//  Bypasses the dialog prompting the user for permission to capture cameras and microphones.
	//  Useful in automatic tests of video-conferencing Web applications. This is nearly
	//  identical to kUseFakeUIForMediaStream, with the exception being that this flag does NOT
	//  affect screen-capture.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --auto-accept-camera-and-microphone-capture</a>
	AutoAcceptCamAndMicCapture() bool // property
	// SetAutoAcceptCamAndMicCapture Set AutoAcceptCamAndMicCapture
	SetAutoAcceptCamAndMicCapture(AValue bool) // property
	// SupportsCompositionController
	//  Returns true if the current WebView2 runtime version supports Composition Controllers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3">See the ICoreWebView2Environment3 article.</a>
	SupportsCompositionController() bool // property
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`
	//  Provide a list of all process using same user data folder except for crashpad process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8">See the ICoreWebView2Environment8 article.</a>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property
	// SupportsControllerOptions
	//  Returns true if the current WebView2 runtime version supports Controller Options.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10">See the ICoreWebView2Environment10 article.</a>
	SupportsControllerOptions() bool // property
	// FailureReportFolderPath
	//  `FailureReportFolderPath` returns the path of the folder where minidump files are written.
	//  Whenever a WebView2 process crashes, a crash dump file will be created in the crash dump folder.
	//  The crash dump format is minidump files.
	//  Please see [Minidump Files documentation](/windows/win32/debug/minidump-files) for detailed information.
	//  Normally when a single child process fails, a minidump will be generated and written to disk,
	//  then the `ProcessFailed` event is raised. But for unexpected crashes, a minidump file might not be generated
	//  at all, despite whether `ProcessFailed` event is raised. If there are multiple
	//  process failures at once, multiple minidump files could be generated. Thus `FailureReportFolderPath`
	//  could contain old minidump files that are not associated with a specific `ProcessFailed` event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment11#get_failurereportfolderpath">See the ICoreWebView2Environment11 article.</a>
	FailureReportFolderPath() string // property
	// StartWebView2
	//  This function is used to initialize WebView2.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</a>
	StartWebView2() bool // function
	// CompareVersions
	//  This method is for anyone want to compare version correctly to determine
	//  which version is newer, older or same.Use it to determine whether
	//  to use webview2 or certain feature based upon version. Sets the value of
	//  aCompRslt to -1, 0 or 1 if aVersion1 is less than, equal or greater
	//  than aVersion2 respectively. Returns false if it fails to parse
	//  any of the version strings.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#comparebrowserversions">See the Globals article.</a>
	CompareVersions(aVersion1, aVersion2 string, aCompRslt *int32) bool // function
	// UpdateDeviceScaleFactor
	//  Update the DeviceScaleFactor property value with the current scale or the ForcedDeviceScaleFactor value.
	UpdateDeviceScaleFactor() // procedure
	// AppendErrorLog
	//  Append aText to the ErrorMessage property.
	AppendErrorLog(aText string) // procedure
}

// TWVLoader Parent: TComponent
//
//	Class used to simplify the WebView2 initialization and destruction.
type TWVLoader struct {
	TComponent
}

func NewWVLoader(AOwner IComponent) IWVLoader {
	r1 := WV().SysCallN(1067, GetObjectUintptr(AOwner))
	return AsWVLoader(r1)
}

func (m *TWVLoader) Environment() ICoreWebView2Environment {
	var resultCoreWebView2Environment uintptr
	WV().SysCallN(1082, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Environment)))
	return AsCoreWebView2Environment(resultCoreWebView2Environment)
}

func (m *TWVLoader) Status() TWV2LoaderStatus {
	r1 := WV().SysCallN(1109, m.Instance())
	return TWV2LoaderStatus(r1)
}

func (m *TWVLoader) AvailableBrowserVersion() string {
	r1 := WV().SysCallN(1061, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) ErrorMessage() string {
	r1 := WV().SysCallN(1084, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) ErrorCode() (resultInt64 int64) {
	WV().SysCallN(1083, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TWVLoader) SetCurrentDir() bool {
	r1 := WV().SysCallN(1105, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetSetCurrentDir(AValue bool) {
	WV().SysCallN(1105, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) Initialized() bool {
	r1 := WV().SysCallN(1093, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) InitializationError() bool {
	r1 := WV().SysCallN(1092, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) CheckFiles() bool {
	r1 := WV().SysCallN(1064, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetCheckFiles(AValue bool) {
	WV().SysCallN(1064, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ShowMessageDlg() bool {
	r1 := WV().SysCallN(1106, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetShowMessageDlg(AValue bool) {
	WV().SysCallN(1106, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) InitCOMLibrary() bool {
	r1 := WV().SysCallN(1091, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetInitCOMLibrary(AValue bool) {
	WV().SysCallN(1091, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) CustomCommandLineSwitches() string {
	r1 := WV().SysCallN(1068, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) DeviceScaleFactor() float32 {
	r1 := WV().SysCallN(1072, m.Instance())
	return float32(r1)
}

func (m *TWVLoader) ReRaiseExceptions() bool {
	r1 := WV().SysCallN(1102, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetReRaiseExceptions(AValue bool) {
	WV().SysCallN(1102, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) InstalledRuntimeVersion() string {
	r1 := WV().SysCallN(1094, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) LoaderDllPath() string {
	r1 := WV().SysCallN(1098, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetLoaderDllPath(AValue string) {
	WV().SysCallN(1098, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) UseInternalLoader() bool {
	r1 := WV().SysCallN(1116, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetUseInternalLoader(AValue bool) {
	WV().SysCallN(1116, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) BrowserExecPath() string {
	r1 := WV().SysCallN(1063, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetBrowserExecPath(AValue string) {
	WV().SysCallN(1063, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) UserDataFolder() string {
	r1 := WV().SysCallN(1117, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetUserDataFolder(AValue string) {
	WV().SysCallN(1117, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) AdditionalBrowserArguments() string {
	r1 := WV().SysCallN(1052, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetAdditionalBrowserArguments(AValue string) {
	WV().SysCallN(1052, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) Language() string {
	r1 := WV().SysCallN(1097, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetLanguage(AValue string) {
	WV().SysCallN(1097, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) TargetCompatibleBrowserVersion() string {
	r1 := WV().SysCallN(1112, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetTargetCompatibleBrowserVersion(AValue string) {
	WV().SysCallN(1112, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) AllowSingleSignOnUsingOSPrimaryAccount() bool {
	r1 := WV().SysCallN(1056, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowSingleSignOnUsingOSPrimaryAccount(AValue bool) {
	WV().SysCallN(1056, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ExclusiveUserDataFolderAccess() bool {
	r1 := WV().SysCallN(1085, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetExclusiveUserDataFolderAccess(AValue bool) {
	WV().SysCallN(1085, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) CustomCrashReportingEnabled() bool {
	r1 := WV().SysCallN(1069, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetCustomCrashReportingEnabled(AValue bool) {
	WV().SysCallN(1069, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) EnableTrackingPrevention() bool {
	r1 := WV().SysCallN(1081, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetEnableTrackingPrevention(AValue bool) {
	WV().SysCallN(1081, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AreBrowserExtensionsEnabled() bool {
	r1 := WV().SysCallN(1058, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAreBrowserExtensionsEnabled(AValue bool) {
	WV().SysCallN(1058, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) EnableGPU() bool {
	r1 := WV().SysCallN(1080, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetEnableGPU(AValue bool) {
	WV().SysCallN(1080, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) EnableFeatures() string {
	r1 := WV().SysCallN(1079, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetEnableFeatures(AValue string) {
	WV().SysCallN(1079, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DisableFeatures() string {
	r1 := WV().SysCallN(1076, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetDisableFeatures(AValue string) {
	WV().SysCallN(1076, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) EnableBlinkFeatures() string {
	r1 := WV().SysCallN(1078, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetEnableBlinkFeatures(AValue string) {
	WV().SysCallN(1078, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DisableBlinkFeatures() string {
	r1 := WV().SysCallN(1074, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetDisableBlinkFeatures(AValue string) {
	WV().SysCallN(1074, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) BlinkSettings() string {
	r1 := WV().SysCallN(1062, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetBlinkSettings(AValue string) {
	WV().SysCallN(1062, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) ForceFieldTrials() string {
	r1 := WV().SysCallN(1088, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetForceFieldTrials(AValue string) {
	WV().SysCallN(1088, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) ForceFieldTrialParams() string {
	r1 := WV().SysCallN(1087, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetForceFieldTrialParams(AValue string) {
	WV().SysCallN(1087, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) SmartScreenProtectionEnabled() bool {
	r1 := WV().SysCallN(1107, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetSmartScreenProtectionEnabled(AValue bool) {
	WV().SysCallN(1107, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AllowInsecureLocalhost() bool {
	r1 := WV().SysCallN(1054, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowInsecureLocalhost(AValue bool) {
	WV().SysCallN(1054, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) DisableWebSecurity() bool {
	r1 := WV().SysCallN(1077, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetDisableWebSecurity(AValue bool) {
	WV().SysCallN(1077, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) TouchEvents() TWVState {
	r1 := WV().SysCallN(1113, 0, m.Instance(), 0)
	return TWVState(r1)
}

func (m *TWVLoader) SetTouchEvents(AValue TWVState) {
	WV().SysCallN(1113, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) HyperlinkAuditing() bool {
	r1 := WV().SysCallN(1090, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetHyperlinkAuditing(AValue bool) {
	WV().SysCallN(1090, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AutoplayPolicy() TWVAutoplayPolicy {
	r1 := WV().SysCallN(1060, 0, m.Instance(), 0)
	return TWVAutoplayPolicy(r1)
}

func (m *TWVLoader) SetAutoplayPolicy(AValue TWVAutoplayPolicy) {
	WV().SysCallN(1060, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) MuteAudio() bool {
	r1 := WV().SysCallN(1099, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetMuteAudio(AValue bool) {
	WV().SysCallN(1099, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) KioskPrinting() bool {
	r1 := WV().SysCallN(1096, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetKioskPrinting(AValue bool) {
	WV().SysCallN(1096, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ProxySettings() IWVProxySettings {
	r1 := WV().SysCallN(1101, m.Instance())
	return AsWVProxySettings(r1)
}

func (m *TWVLoader) AllowFileAccessFromFiles() bool {
	r1 := WV().SysCallN(1053, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowFileAccessFromFiles(AValue bool) {
	WV().SysCallN(1053, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AllowRunningInsecureContent() bool {
	r1 := WV().SysCallN(1055, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowRunningInsecureContent(AValue bool) {
	WV().SysCallN(1055, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) DisableBackgroundNetworking() bool {
	r1 := WV().SysCallN(1073, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetDisableBackgroundNetworking(AValue bool) {
	WV().SysCallN(1073, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ForcedDeviceScaleFactor() float32 {
	r1 := WV().SysCallN(1089, 0, m.Instance(), 0)
	return float32(r1)
}

func (m *TWVLoader) SetForcedDeviceScaleFactor(AValue float32) {
	WV().SysCallN(1089, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) RemoteDebuggingPort() int32 {
	r1 := WV().SysCallN(1104, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TWVLoader) SetRemoteDebuggingPort(AValue int32) {
	WV().SysCallN(1104, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) RemoteAllowOrigins() string {
	r1 := WV().SysCallN(1103, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetRemoteAllowOrigins(AValue string) {
	WV().SysCallN(1103, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DebugLog() TWV2DebugLog {
	r1 := WV().SysCallN(1070, 0, m.Instance(), 0)
	return TWV2DebugLog(r1)
}

func (m *TWVLoader) SetDebugLog(AValue TWV2DebugLog) {
	WV().SysCallN(1070, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) DebugLogLevel() TWV2DebugLogLevel {
	r1 := WV().SysCallN(1071, 0, m.Instance(), 0)
	return TWV2DebugLogLevel(r1)
}

func (m *TWVLoader) SetDebugLogLevel(AValue TWV2DebugLogLevel) {
	WV().SysCallN(1071, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) JavaScriptFlags() string {
	r1 := WV().SysCallN(1095, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetJavaScriptFlags(AValue string) {
	WV().SysCallN(1095, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DisableEdgePitchNotification() bool {
	r1 := WV().SysCallN(1075, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetDisableEdgePitchNotification(AValue bool) {
	WV().SysCallN(1075, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) TreatInsecureOriginAsSecure() string {
	r1 := WV().SysCallN(1114, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetTreatInsecureOriginAsSecure(AValue string) {
	WV().SysCallN(1114, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) AutoAcceptCamAndMicCapture() bool {
	r1 := WV().SysCallN(1059, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAutoAcceptCamAndMicCapture(AValue bool) {
	WV().SysCallN(1059, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) SupportsCompositionController() bool {
	r1 := WV().SysCallN(1110, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) ProcessInfos() ICoreWebView2ProcessInfoCollection {
	var resultCoreWebView2ProcessInfoCollection uintptr
	WV().SysCallN(1100, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessInfoCollection)))
	return AsCoreWebView2ProcessInfoCollection(resultCoreWebView2ProcessInfoCollection)
}

func (m *TWVLoader) SupportsControllerOptions() bool {
	r1 := WV().SysCallN(1111, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) FailureReportFolderPath() string {
	r1 := WV().SysCallN(1086, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) StartWebView2() bool {
	r1 := WV().SysCallN(1108, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) CompareVersions(aVersion1, aVersion2 string, aCompRslt *int32) bool {
	var result1 uintptr
	r1 := WV().SysCallN(1066, m.Instance(), PascalStr(aVersion1), PascalStr(aVersion2), uintptr(unsafePointer(&result1)))
	*aCompRslt = int32(result1)
	return GoBool(r1)
}

func WVLoaderClass() TClass {
	ret := WV().SysCallN(1065)
	return TClass(ret)
}

func (m *TWVLoader) UpdateDeviceScaleFactor() {
	WV().SysCallN(1115, m.Instance())
}

func (m *TWVLoader) AppendErrorLog(aText string) {
	WV().SysCallN(1057, m.Instance(), PascalStr(aText))
}
