//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// TCefPlatformThreadId
//
//	Platform thread ID.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_thread_internal.h">CEF source file: /include/internal/cef_thread_internal.h (cef_platform_thread_id_t))</a>
type TCefPlatformThreadId = DWORD

// TCefPlatformThreadHandle
//
//	Platform thread handle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_thread_internal.h">CEF source file: /include/internal/cef_thread_internal.h (cef_platform_thread_handle_t))</a>
type TCefPlatformThreadHandle = DWORD

// TCefTransitionType
//
//	Transition type for a request. Made up of one source value and 0 or more qualifiers.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t))</a>
type TCefTransitionType = Cardinal

// TCefColor
//
//	32-bit ARGB color value, not premultiplied. The color components are always in a known order. Equivalent to the SkColor type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_color_t))</a>
type TCefColor = Cardinal

// TCefErrorCode
//
//	Supported error code values.
//	Ranges:
//	0- 99 System related errors
//	100-199 Connection related errors
//	200-299 Certificate errors
//	300-399 HTTP errors
//	400-499 Cache errors
//	500-599 ?
//	600-699 FTP errors
//	700-799 Certificate manager errors
//	800-899 DNS resolver errors
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t))</a>
//	<a href="https://chromium.googlesource.com/chromium/src/+/master/net/base/net_error_list.h">For the complete list of error values see include/base/internal/cef_net_error_list.h which includes this Chromium source file /net/base/net_error_list.h)</a>
type TCefErrorCode = Integer

// TCefCertStatus
//
//	Supported certificate status code values. See net\cert\cert_status_flags.h for more information. CERT_STATUS_NONE is new in CEF because we use an enum while cert_status_flags.h uses a typedef and static const variables.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cert_status_t))</a>
type TCefCertStatus = Integer

// TCefSSLVersion
//
//	Supported SSL version values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_version_t))</a>
//	<a href="https://source.chromium.org/chromium/chromium/src/+/main:net/ssl/ssl_connection_status_flags.h">See net/ssl/ssl_connection_status_flags.h for more information.)</a>
type TCefSSLVersion = integer

// TCefStringList
//
//	CEF string maps are a set of key/value string pairs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_list.h">CEF source file: /include/internal/cef_string_list.h (cef_string_list_t))</a>
type TCefStringList = Pointer

// TCefStringMapHandle
//
//	CEF string maps are a set of key/value string pairs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_map.h">CEF source file: /include/internal/cef_string_map.h (cef_string_map_t))</a>
type TCefStringMapHandle = Pointer

// TCefStringMultimapHandle
//
//	CEF string multimaps are a set of key/value string pairs. More than one value can be assigned to a single key.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_string_multimap.h">CEF source file: /include/internal/cef_string_multimap.h (cef_string_multimap_t))</a>
type TCefStringMultimapHandle = Pointer

// TCefUriUnescapeRule
//
//	URI unescape rules passed to CefURIDecode().
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t))</a>
type TCefUriUnescapeRule = Integer

// TCefDomEventCategory
//
//	DOM event category flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t))</a>
type TCefDomEventCategory = Integer

// TCefEventFlags
//
//	Supported event bit flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t))</a>
type TCefEventFlags = Cardinal

// TCefDragOperations
//
//	"Verb" of a drag-and-drop operation as negotiated between the source and destination. These constants match their equivalents in WebCore's DragActions.h and should not be renumbered.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t))</a>
type TCefDragOperations = Cardinal

// TCefDragOperation
type TCefDragOperation = Cardinal

// TCefV8AccessControls
//
//	V8 access control values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_accesscontrol_t))</a>
type TCefV8AccessControls = Cardinal

// TCefV8PropertyAttributes
//
//	V8 property attribute values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_propertyattribute_t))</a>
type TCefV8PropertyAttributes = Cardinal

// TCefUrlRequestFlags
//
//	Flags used to customize the behavior of CefURLRequest.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t))</a>
type TCefUrlRequestFlags = Cardinal

// TCefContextMenuTypeFlags
//
//	Supported context menu type flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t))</a>
type TCefContextMenuTypeFlags = Cardinal

// TCefContextMenuMediaStateFlags
//
//	Supported context menu media state bit flags. These constants match their equivalents in Chromium's ContextMenuData::MediaFlags and should not be renumbered.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_media_state_flags_t))</a>
type TCefContextMenuMediaStateFlags = Cardinal

// TCefContextMenuEditStateFlags
//
//	Supported context menu edit state bit flags. These constants match their equivalents in Chromium's ContextMenuDataEditFlags and should not be renumbered.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_edit_state_flags_t))</a>
type TCefContextMenuEditStateFlags = Cardinal

// TCefJsonWriterOptions
//
//	Options that can be passed to CefWriteJSON.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t))</a>
type TCefJsonWriterOptions = Cardinal

// TCefSSLContentStatus
//
//	Supported SSL content status flags. See content/public/common/ssl_status.h for more information.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_content_status_t))</a>
type TCefSSLContentStatus = Cardinal

// TCefLogSeverity
//
//	Log severity levels.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t))</a>
type TCefLogSeverity = Cardinal

// TCefFileDialogMode
//
//	Supported file dialog modes.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t))</a>
type TCefFileDialogMode = Cardinal

// TCefDuplexMode
//
//	Print job duplex mode values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_duplex_mode_t))</a>
type TCefDuplexMode = Integer

// TCefSchemeOptions
//
//	Configuration options for registering a custom scheme. These values are used when calling AddCustomScheme.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t))</a>
type TCefSchemeOptions = Integer

// TCefMediaRouterCreateResult
//
//	Result codes for ICefMediaRouter.CreateRoute. Should be kept in sync with Chromium's media_router::mojom::RouteRequestResultCode type.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_create_result_t))</a>
type TCefMediaRouterCreateResult = Integer

// TCefCookiePriority
//
//	Cookie priority values.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cookie_priority_t))</a>
type TCefCookiePriority = Integer

// TCefTextFieldCommands
//
//	Represents commands available to TextField.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_field_commands_t))</a>
type TCefTextFieldCommands = Integer

// TCefChromeToolbarType
//
//	Chrome toolbar types.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_toolbar_type_t))</a>
type TCefChromeToolbarType = Integer

// TCefDockingMode
//
//	Docking modes supported by ICefWindow.AddOverlay.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_docking_mode_t))</a>
type TCefDockingMode = Integer

// TCefShowState
//
//	Show states supported by ICefWindowDelegate.GetInitialShowState.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t))</a>
type TCefShowState = Integer

// TCefQuickMenuEditStateFlags
//
//	Supported quick menu state bit flags.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_quick_menu_edit_state_flags_t))</a>
type TCefQuickMenuEditStateFlags = Integer

// TCefTouchHandleStateFlags
//
//	Values indicating what state of the touch handle is set.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_handle_state_flags_t))</a>
type TCefTouchHandleStateFlags = Integer

// TCefMediaAccessPermissionTypes
//
//	Media access permissions used by OnRequestMediaAccessPermission.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t))</a>
type TCefMediaAccessPermissionTypes = Integer

// TCefPermissionRequestTypes
//
//	Permission types used with OnShowPermissionPrompt. Some types are platform-specific or only supported with the Chrome runtime. Should be kept in sync with Chromium's permissions::RequestType type.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_permission_request_types_t))</a>
type TCefPermissionRequestTypes = Integer

// TCefDownloadInterruptReason
//
//	Download interrupt reasons. Should be kept in sync with Chromium's download::DownloadInterruptReason type.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t))</a>
type TCefDownloadInterruptReason = Integer

// TCefMenuId
//
//	Supported menu IDs. Non-English translations can be provided for the IDS_MENU_* strings in ICefResourceBundleHandler.GetLocalizedString().
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t))</a>
type TCefMenuId = Integer

// TCefLogItems
//
//	Log items prepended to each log line.
//	See the uCEFConstants unit for all possible values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t))</a>
type TCefLogItems = Cardinal

// TCefCloseBrowserAction
//
//	Action taken after the TChromium.Onclose event. cbaCancel : stop closing the browser. cbaClose : continue closing the browser. cbaDelay : stop closing the browser momentarily. Used when the application needs to execute some custom processes before closing the browser. This is usually needed to destroy a TCEFWindowParent in the main thread before closing the browser.
type TCefCloseBrowserAction = int32

const (
	CbaClose TCefCloseBrowserAction = iota
	CbaDelay
	CbaCancel
)

// TCefProcessType
//
//	Sub-process types of Chromium.
type TCefProcessType = int32

const (
	PtBrowser TCefProcessType = iota
	PtRenderer
	PtZygote
	PtGPU
	PtUtility
	PtBroker
	PtCrashpad
	PtOther
)

// TCefCookiePref
//
//	Used in TChromium preferences to allow or block cookies.
type TCefCookiePref = int32

const (
	CpDefault TCefCookiePref = iota
	CpAllow
	CpBlock
)

// TCefBrowserNavigation
//
//	Used by TCefBrowserNavigationTask to navigate in the right CEF thread.
type TCefBrowserNavigation = int32

const (
	BnBack TCefBrowserNavigation = iota
	BnForward
	BnReload
	BnReloadIgnoreCache
	BnStopLoad
)

// TCefAplicationStatus
//
//	Status of TCefAplicationCore.
type TCefAplicationStatus = int32

const (
	AsLoading TCefAplicationStatus = iota
	AsLoaded
	AsInitialized
	AsShuttingDown
	AsUnloaded
	AsErrorMissingFiles
	AsErrorDLLVersion
	AsErrorWindowsVersion
	AsErrorLoadingLibrary
	AsErrorInitializingLibrary
	AsErrorExecutingProcess
)

// TCefUIColorMode
//
//	Color mode in UI for platforms that support it.
type TCefUIColorMode = int32

const (
	UicmSystemDefault TCefUIColorMode = iota
	UicmForceLight
	UicmForceDark
)

// TCefProxyScheme
//
//	Supported proxy schemes in Chromium.
type TCefProxyScheme = int32

const (
	PsHTTP TCefProxyScheme = iota
	PsSOCKS4
	PsSOCKS5
)

// TCefClearDataStorageTypes
//
//	Storage types used by the Storage.clearDataForOrigin DevTools method in TChromiumCore.ClearDataForOrigin.
type TCefClearDataStorageTypes = int32

const (
	CdstAppCache TCefClearDataStorageTypes = iota
	CdstCookies
	CdstFileSystems
	CdstIndexeddb
	CdstLocalStorage
	CdstShaderCache
	CdstWebsql
	CdstServiceWorkers
	CdstCacheStorage
	CdstAll
)

// TCefAutoplayPolicy
//
//	Autoplay policy types used by TCefApplicationCore.AutoplayPolicy. See the --autoplay-policy switch.
type TCefAutoplayPolicy = int32

const (
	AppDefault TCefAutoplayPolicy = iota
	AppDocumentUserActivationRequired
	AppNoUserGestureRequired
	AppUserGestureRequired
)

// TCefWebRTCHandlingPolicy
//
//	WebRTC handling policy types used by TChromiumCore.WebRTCIPHandlingPolicy.
type TCefWebRTCHandlingPolicy = int32

const (
	HpDefault TCefWebRTCHandlingPolicy = iota
	HpDefaultPublicAndPrivateInterfaces
	HpDefaultPublicInterfaceOnly
	HpDisableNonProxiedUDP
)

// TCefNetLogCaptureMode
//
//	Values used by the --net-log-capture-mode command line switch. Sets the granularity of events to capture in the network log.
//	<a href="https://source.chromium.org/chromium/chromium/src/+/main:content/browser/network_service_instance_impl.cc">network_service_instance_impl.cc)</a>
//	<a href="https://source.chromium.org/chromium/chromium/src/+/main:net/log/net_log_capture_mode.h">net_log_capture_mode.h)</a>
type TCefNetLogCaptureMode = int32

const (
	NlcmDefault TCefNetLogCaptureMode = iota
	NlcmIncludeSensitive
	NlcmEverything
)

// TCefBatterySaverModeState
//
//	Values used by the battery saver mode state preference.
//	<a href="https://source.chromium.org/chromium/chromium/src/+/main:components/performance_manager/public/user_tuning/prefs.h">components/performance_manager/public/user_tuning/prefs.h)</a>
type TCefBatterySaverModeState = int32

const (
	BsmsDisabled TCefBatterySaverModeState = iota
	BsmsEnabledBelowThreshold
	BsmsEnabledOnBattery
	BsmsEnabled
	BsmsDefault
)

// TCefHighEfficiencyModeState
//
//	Values used by the high efficiency mode state preference.
//	<a href="https://source.chromium.org/chromium/chromium/src/+/main:components/performance_manager/public/user_tuning/prefs.h">components/performance_manager/public/user_tuning/prefs.h)</a>
type TCefHighEfficiencyModeState = int32

const (
	KDisabled TCefHighEfficiencyModeState = iota
	KEnabled
	KEnabledOnTimer
	KDefault
)

// TCEFDialogType
//
//	Used by TCEFFileDialogInfo.
type TCEFDialogType = int32

const (
	DtOpen TCEFDialogType = iota
	DtOpenMultiple
	DtOpenFolder
	DtSave
)

// TCefMediaType
//
//	Used by TCefMediaSinkInfo and TCefMediaSourceInfo.
type TCefMediaType = int32

const (
	MtCast TCefMediaType = iota
	MtDial
	MtUnknown
)

// TCefState
//
//	Represents the state of a setting.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_state_t))</a>
type TCefState = int32

const (
	STATE_DEFAULT TCefState = iota
	STATE_ENABLED
	STATE_DISABLED
)

// TCefScaleFactor
//
//	Supported UI scale factors for the platform. SCALE_FACTOR_NONE is used for density independent resources such as string, html/js files or an image that can be used for any scale factors (such as wallpapers).
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scale_factor_t))</a>
type TCefScaleFactor = int32

const (
	SCALE_FACTOR_NONE TCefScaleFactor = iota
	SCALE_FACTOR_100P
	SCALE_FACTOR_125P
	SCALE_FACTOR_133P
	SCALE_FACTOR_140P
	SCALE_FACTOR_150P
	SCALE_FACTOR_180P
	SCALE_FACTOR_200P
	SCALE_FACTOR_250P
	SCALE_FACTOR_300P
)

// TCefValueType
//
//	Supported value types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_value_type_t))</a>
type TCefValueType = int32

const (
	VTYPE_INVALID TCefValueType = iota
	VTYPE_NULL
	VTYPE_BOOL
	VTYPE_INT
	VTYPE_DOUBLE
	VTYPE_STRING
	VTYPE_BINARY
	VTYPE_DICTIONARY
	VTYPE_LIST
)

// TCefMediaRouteConnectionState
//
//	Connection state for a MediaRoute object.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_connection_state_t))</a>
type TCefMediaRouteConnectionState = int32

const (
	CEF_MRCS_UNKNOWN TCefMediaRouteConnectionState = iota
	CEF_MRCS_CONNECTING
	CEF_MRCS_CONNECTED
	CEF_MRCS_CLOSED
	CEF_MRCS_TERMINATED
)

// TCefMediaSinkIconType
//
//	Icon types for a MediaSink object. Should be kept in sync with Chromium's media_router::SinkIconType type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_sink_icon_type_t))</a>
type TCefMediaSinkIconType = int32

const (
	CEF_MSIT_CAST TCefMediaSinkIconType = iota
	CEF_MSIT_CAST_AUDIO_GROUP
	CEF_MSIT_CAST_AUDIO
	CEF_MSIT_MEETING
	CEF_MSIT_HANGOUT
	CEF_MSIT_EDUCATION
	CEF_MSIT_WIRED_DISPLAY
	CEF_MSIT_GENERIC
	CEF_MSIT_TOTAL_COUNT
)

// TCefReferrerPolicy
//
//	Policy for how the Referrer HTTP header value will be sent during navigation. If the `--no-referrers` command-line flag is specified then the policy value will be ignored and the Referrer value will never be sent. Must be kept synchronized with net::URLRequest::ReferrerPolicy from Chromium.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_referrer_policy_t))</a>
type TCefReferrerPolicy = int32

const (
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_FROM_SECURE_TO_INSECURE TCefReferrerPolicy = iota
	REFERRER_POLICY_REDUCE_REFERRER_GRANULARITY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_ONLY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_NEVER_CLEAR_REFERRER
	REFERRER_POLICY_ORIGIN
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_CLEAR_ON_TRANSITION_FROM_SECURE_TO_INSECURE
	REFERRER_POLICY_NO_REFERRER
)

// TCefPostDataElementType
//
//	Post data elements may represent either bytes or files.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_postdataelement_type_t))</a>
type TCefPostDataElementType = int32

const (
	PDE_TYPE_EMPTY TCefPostDataElementType = iota
	PDE_TYPE_BYTES
	PDE_TYPE_FILE
)

// TCefResourceType
//
//	Resource type for a request. These constants match their equivalents in Chromium's ResourceType and should not be renumbered.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resource_type_t))</a>
type TCefResourceType = int32

const (
	RT_MAIN_FRAME TCefResourceType = iota
	RT_SUB_FRAME
	RT_STYLESHEET
	RT_SCRIPT
	RT_IMAGE
	RT_FONT_RESOURCE
	RT_SUB_RESOURCE
	RT_OBJECT
	RT_MEDIA
	RT_WORKER
	RT_SHARED_WORKER
	RT_PREFETCH
	RT_FAVICON
	RT_XHR
	RT_PING
	RT_SERVICE_WORKER
	RT_CSP_REPORT
	RT_PLUGIN_RESOURCE
	RT_EMPTY_FILLER_TYPE_DO_NOT_USE
	RT_NAVIGATION_PRELOAD_MAIN_FRAME
	RT_NAVIGATION_PRELOAD_SUB_FRAME
)

// TCefDomDocumentType
//
//	DOM document types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_document_type_t))</a>
type TCefDomDocumentType = int32

const (
	DOM_DOCUMENT_TYPE_UNKNOWN TCefDomDocumentType = iota
	DOM_DOCUMENT_TYPE_HTML
	DOM_DOCUMENT_TYPE_XHTML
	DOM_DOCUMENT_TYPE_PLUGIN
)

// TCefDomNodeType
//
//	DOM node types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_node_type_t))</a>
type TCefDomNodeType = int32

const (
	DOM_NODE_TYPE_UNSUPPORTED TCefDomNodeType = iota
	DOM_NODE_TYPE_ELEMENT
	DOM_NODE_TYPE_ATTRIBUTE
	DOM_NODE_TYPE_TEXT
	DOM_NODE_TYPE_CDATA_SECTION
	DOM_NODE_TYPE_PROCESSING_INSTRUCTIONS
	DOM_NODE_TYPE_COMMENT
	DOM_NODE_TYPE_DOCUMENT
	DOM_NODE_TYPE_DOCUMENT_TYPE
	DOM_NODE_TYPE_DOCUMENT_FRAGMENT
)

// TCefContextMenuMediaType
//
//	Supported context menu media types. These constants match their equivalents in Chromium's ContextMenuDataMediaType and should not be renumbered.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_media_type_t))</a>
type TCefContextMenuMediaType = int32

const (
	CM_MEDIATYPE_NONE TCefContextMenuMediaType = iota
	CM_MEDIATYPE_IMAGE
	CM_MEDIATYPE_VIDEO
	CM_MEDIATYPE_AUDIO
	CM_MEDIATYPE_CANVAS
	CM_MEDIATYPE_FILE
	CM_MEDIATYPE_PLUGIN
)

// TCefMenuItemType
//
//	Supported menu item types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_item_type_t))</a>
type TCefMenuItemType = int32

const (
	MENUITEMTYPE_NONE TCefMenuItemType = iota
	MENUITEMTYPE_COMMAND
	MENUITEMTYPE_CHECK
	MENUITEMTYPE_RADIO
	MENUITEMTYPE_SEPARATOR
	MENUITEMTYPE_SUBMENU
)

// TCefFocusSource
//
//	Focus sources.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_focus_source_t))</a>
type TCefFocusSource = int32

const (
	FOCUS_SOURCE_NAVIGATION TCefFocusSource = iota
	FOCUS_SOURCE_SYSTEM
)

// TCefJsDialogType
//
//	Supported JavaScript dialog types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_jsdialog_type_t))</a>
type TCefJsDialogType = int32

const (
	JSDIALOGTYPE_ALERT TCefJsDialogType = iota
	JSDIALOGTYPE_CONFIRM
	JSDIALOGTYPE_PROMPT
)

// TCefKeyEventType
//
//	Notification that a character was typed. Use this for text input. Key down events may generate 0, 1, or more than one character event depending on the key, locale, and operating system.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_key_event_type_t))</a>
type TCefKeyEventType = int32

const (
	KEYEVENT_RAWKEYDOWN TCefKeyEventType = iota
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

// TCefWindowOpenDisposition
//
//	The manner in which a link click should be opened. These constants match their equivalents in Chromium's window_open_disposition.h and should not be renumbered.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_window_open_disposition_t))</a>
type TCefWindowOpenDisposition = int32

const (
	WOD_UNKNOWN TCefWindowOpenDisposition = iota
	WOD_CURRENT_TAB
	WOD_SINGLETON_TAB
	WOD_NEW_FOREGROUND_TAB
	WOD_NEW_BACKGROUND_TAB
	WOD_NEW_POPUP
	WOD_NEW_WINDOW
	WOD_SAVE_TO_DISK
	WOD_OFF_THE_RECORD
	WOD_IGNORE_ACTION
	WOD_SWITCH_TO_TAB
	WOD_NEW_PICTURE_IN_PICTURE
)

// TCefTextInpuMode
//
//	Input mode of a virtual keyboard. These constants match their equivalents in Chromium's text_input_mode.h and should not be renumbered. See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_input_mode_t))</a>
type TCefTextInpuMode = int32

const (
	CEF_TEXT_INPUT_MODE_DEFAULT TCefTextInpuMode = iota
	CEF_TEXT_INPUT_MODE_NONE
	CEF_TEXT_INPUT_MODE_TEXT
	CEF_TEXT_INPUT_MODE_TEL
	CEF_TEXT_INPUT_MODE_URL
	CEF_TEXT_INPUT_MODE_EMAIL
	CEF_TEXT_INPUT_MODE_NUMERIC
	CEF_TEXT_INPUT_MODE_DECIMAL
	CEF_TEXT_INPUT_MODE_SEARCH
)

// TCefTouchEeventType
//
//	Touch points states types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_event_type_t))</a>
type TCefTouchEeventType = int32

const (
	CEF_TET_RELEASED TCefTouchEeventType = iota
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

// TCefPointerType
//
//	The device type that caused the event.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_pointer_type_t))</a>
type TCefPointerType = int32

const (
	CEF_POINTER_TYPE_TOUCH TCefPointerType = iota
	CEF_POINTER_TYPE_MOUSE
	CEF_POINTER_TYPE_PEN
	CEF_POINTER_TYPE_ERASER
	CEF_POINTER_TYPE_UNKNOWN
)

// TCefChannelLayout
//
//	Enumerates the various representations of the ordering of audio channels. Must be kept synchronized with media::ChannelLayout from Chromium. See media\base\channel_layout.h
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_channel_layout_t))</a>
type TCefChannelLayout = int32

const (
	CEF_CHANNEL_LAYOUT_NONE TCefChannelLayout = iota
	CEF_CHANNEL_LAYOUT_UNSUPPORTED
	CEF_CHANNEL_LAYOUT_MONO
	CEF_CHANNEL_LAYOUT_STEREO
	CEF_CHANNEL_LAYOUT_2_1
	CEF_CHANNEL_LAYOUT_SURROUND
	CEF_CHANNEL_LAYOUT_4_0
	CEF_CHANNEL_LAYOUT_2_2
	CEF_CHANNEL_LAYOUT_QUAD
	CEF_CHANNEL_LAYOUT_5_0
	CEF_CHANNEL_LAYOUT_5_1
	CEF_CHANNEL_LAYOUT_5_0_BACK
	CEF_CHANNEL_LAYOUT_5_1_BACK
	CEF_CHANNEL_LAYOUT_7_0
	CEF_CHANNEL_LAYOUT_7_1
	CEF_CHANNEL_LAYOUT_7_1_WIDE
	CEF_CHANNEL_LAYOUT_STEREO_DOWNMIX
	CEF_CHANNEL_LAYOUT_2POINT1
	CEF_CHANNEL_LAYOUT_3_1
	CEF_CHANNEL_LAYOUT_4_1
	CEF_CHANNEL_LAYOUT_6_0
	CEF_CHANNEL_LAYOUT_6_0_FRONT
	CEF_CHANNEL_LAYOUT_HEXAGONAL
	CEF_CHANNEL_LAYOUT_6_1
	CEF_CHANNEL_LAYOUT_6_1_BACK
	CEF_CHANNEL_LAYOUT_6_1_FRONT
	CEF_CHANNEL_LAYOUT_7_0_FRONT
	CEF_CHANNEL_LAYOUT_7_1_WIDE_BACK
	CEF_CHANNEL_LAYOUT_OCTAGONAL
	CEF_CHANNEL_LAYOUT_DISCRETE
	CEF_CHANNEL_LAYOUT_STEREO_AND_KEYBOARD_MIC
	CEF_CHANNEL_LAYOUT_4_1_QUAD_SIDE
	CEF_CHANNEL_LAYOUT_BITSTREAM
	CEF_CHANNEL_LAYOUT_5_1_4_DOWNMIX
)

// TCefCookieSameSite
//
//	Cookie same site values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cookie_same_site_t))</a>
type TCefCookieSameSite = int32

const (
	CEF_COOKIE_SAME_SITE_UNSPECIFIED TCefCookieSameSite = iota
	CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	CEF_COOKIE_SAME_SITE_LAX_MODE
	CEF_COOKIE_SAME_SITE_STRICT_MODE
)

// TCefPaintElementType
//
//	Paint element types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_paint_element_type_t))</a>
type TCefPaintElementType = int32

const (
	PET_VIEW TCefPaintElementType = iota
	PET_POPUP
)

// TCefCursorType
//
//	Cursor type values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cursor_type_t))</a>
type TCefCursorType = int32

const (
	CT_POINTER TCefCursorType = iota
	CT_CROSS
	CT_HAND
	CT_IBEAM
	CT_WAIT
	CT_HELP
	CT_EASTRESIZE
	CT_NORTHRESIZE
	CT_NORTHEASTRESIZE
	CT_NORTHWESTRESIZE
	CT_SOUTHRESIZE
	CT_SOUTHEASTRESIZE
	CT_SOUTHWESTRESIZE
	CT_WESTRESIZE
	CT_NORTHSOUTHRESIZE
	CT_EASTWESTRESIZE
	CT_NORTHEASTSOUTHWESTRESIZE
	CT_NORTHWESTSOUTHEASTRESIZE
	CT_COLUMNRESIZE
	CT_ROWRESIZE
	CT_MIDDLEPANNING
	CT_EASTPANNING
	CT_NORTHPANNING
	CT_NORTHEASTPANNING
	CT_NORTHWESTPANNING
	CT_SOUTHPANNING
	CT_SOUTHEASTPANNING
	CT_SOUTHWESTPANNING
	CT_WESTPANNING
	CT_MOVE
	CT_VERTICALTEXT
	CT_CELL
	CT_CONTEXTMENU
	CT_ALIAS
	CT_PROGRESS
	CT_NODROP
	CT_COPY
	CT_NONE
	CT_NOTALLOWED
	CT_ZOOMIN
	CT_ZOOMOUT
	CT_GRAB
	CT_GRABBING
	CT_MIDDLE_PANNING_VERTICAL
	CT_MIDDLE_PANNING_HORIZONTAL
	CT_CUSTOM
	CT_DND_NONE
	CT_DND_MOVE
	CT_DND_COPY
	CT_DND_LIN
)

// TCefNavigationType
//
//	Navigation types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_navigation_type_t))</a>
type TCefNavigationType = int32

const (
	NAVIGATION_LINK_CLICKED TCefNavigationType = iota
	NAVIGATION_FORM_SUBMITTED
	NAVIGATION_BACK_FORWARD
	NAVIGATION_RELOAD
	NAVIGATION_FORM_RESUBMITTED
	NAVIGATION_OTHER
)

// TCefProcessId
//
//	Existing process IDs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_process_id_t))</a>
type TCefProcessId = int32

const (
	PID_BROWSER TCefProcessId = iota
	PID_RENDERER
)

// TCefThreadId
//
//	Existing thread IDs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_thread_id_t))</a>
type TCefThreadId = int32

const (
	TID_UI TCefThreadId = iota
	TID_FILE_BACKGROUND
	TID_FILE_USER_VISIBLE
	TID_FILE_USER_BLOCKING
	TID_PROCESS_LAUNCHER
	TID_IO
	TID_RENDERER
)

// TCefThreadPriority
//
//	Thread priority values listed in increasing order of importance.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_thread_priority_t))</a>
type TCefThreadPriority = int32

const (
	TP_BACKGROUND TCefThreadPriority = iota
	TP_NORMAL
	TP_DISPLAY
	TP_REALTIME_AUDIO
)

// TCefMessageLoopType
//
//	Flags used to customize the behavior of CefURLRequest.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_message_loop_type_t))</a>
type TCefMessageLoopType = int32

const (
	ML_TYPE_DEFAULT TCefMessageLoopType = iota
	ML_TYPE_UI
	ML_TYPE_IO
)

// TCefCOMInitMode
//
//	Flags used to customize the behavior of CefURLRequest.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_com_init_mode_t))</a>
type TCefCOMInitMode = int32

const (
	COM_INIT_MODE_NONE TCefCOMInitMode = iota
	COM_INIT_MODE_STA
	COM_INIT_MODE_MTA
)

// TCefMouseButtonType
//
//	Mouse button types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_mouse_button_type_t))</a>
type TCefMouseButtonType = int32

const (
	MBT_LEFT TCefMouseButtonType = iota
	MBT_MIDDLE
	MBT_RIGHT
)

// TCefReturnValue
//
//	Return value types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_return_value_t))</a>
type TCefReturnValue = int32

const (
	RV_CANCEL TCefReturnValue = iota
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

// TCefUrlRequestStatus
//
//	Flags that represent CefURLRequest status.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_status_t))</a>
type TCefUrlRequestStatus = int32

const (
	UR_UNKNOWN TCefUrlRequestStatus = iota
	UR_SUCCESS
	UR_IO_PENDING
	UR_CANCELED
	UR_FAILED
)

// TCefTerminationStatus
//
//	Process termination status values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_termination_status_t))</a>
type TCefTerminationStatus = int32

const (
	TS_ABNORMAL_TERMINATION TCefTerminationStatus = iota
	TS_PROCESS_WAS_KILLED
	TS_PROCESS_CRASHED
	TS_PROCESS_OOM
)

// TCefPathKey
//
//	Process termination status values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_path_key_t))</a>
type TCefPathKey = int32

const (
	PK_DIR_CURRENT TCefPathKey = iota
	PK_DIR_EXE
	PK_DIR_MODULE
	PK_DIR_TEMP
	PK_FILE_EXE
	PK_FILE_MODULE
	PK_LOCAL_APP_DATA
	PK_USER_DATA
	PK_DIR_RESOURCES
)

// TCefStorageType
//
//	Storage types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_storage_type_t))</a>
type TCefStorageType = int32

const (
	ST_LOCALSTORAGE TCefStorageType = iota
	ST_SESSIONSTORAGE
)

// TCefResponseFilterStatus
//
//	Return values for ICefResponseFilter.Filter().
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_response_filter_status_t))</a>
type TCefResponseFilterStatus = int32

const (
	RESPONSE_FILTER_NEED_MORE_DATA TCefResponseFilterStatus = iota
	RESPONSE_FILTER_DONE
	RESPONSE_FILTER_ERROR
)

// TCefColorType
//
//	Describes how to interpret the components of a pixel.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_color_type_t))</a>
type TCefColorType = int32

const (
	CEF_COLOR_TYPE_RGBA_8888 TCefColorType = iota
	CEF_COLOR_TYPE_BGRA_8888
)

// TCefAlphaType
//
//	Describes how to interpret the alpha component of a pixel.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_alpha_type_t))</a>
type TCefAlphaType = int32

const (
	CEF_ALPHA_TYPE_OPAQUE TCefAlphaType = iota
	CEF_ALPHA_TYPE_PREMULTIPLIED
	CEF_ALPHA_TYPE_POSTMULTIPLIED
)

// TCefTextStyle
//
//	Text style types. Should be kepy in sync with gfx::TextStyle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_style_t))</a>
type TCefTextStyle = int32

const (
	CEF_TEXT_STYLE_BOLD TCefTextStyle = iota
	CEF_TEXT_STYLE_ITALIC
	CEF_TEXT_STYLE_STRIKE
	CEF_TEXT_STYLE_DIAGONAL_STRIKE
	CEF_TEXT_STYLE_UNDERLINE
)

// TCefMainAxisAlignment
//
//	Specifies where along the main axis the CefBoxLayout child views should be laid out.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_main_axis_alignment_t))</a>
type TCefMainAxisAlignment = int32

const (
	CEF_MAIN_AXIS_ALIGNMENT_START TCefMainAxisAlignment = iota
	CEF_MAIN_AXIS_ALIGNMENT_CENTER
	CEF_MAIN_AXIS_ALIGNMENT_END
)

// TCefCrossAxisAlignment
//
//	Specifies where along the main axis the CefBoxLayout child views should be laid out.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cross_axis_alignment_t))</a>
type TCefCrossAxisAlignment = int32

const (
	CEF_CROSS_AXIS_ALIGNMENT_STRETCH TCefCrossAxisAlignment = iota
	CEF_CROSS_AXIS_ALIGNMENT_START
	CEF_CROSS_AXIS_ALIGNMENT_CENTER
	CEF_CROSS_AXIS_ALIGNMENT_END
)

// TCefPdfPrintMarginType
//
//	Margin type for PDF printing.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_pdf_print_margin_type_t))</a>
type TCefPdfPrintMarginType = int32

const (
	PDF_PRINT_MARGIN_DEFAULT TCefPdfPrintMarginType = iota
	PDF_PRINT_MARGIN_NONE
	PDF_PRINT_MARGIN_CUSTOM
)

// TCefColorModel
//
//	Print job color mode values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_color_model_t))</a>
type TCefColorModel = int32

const (
	COLOR_MODEL_UNKNOWN TCefColorModel = iota
	COLOR_MODEL_GRAY
	COLOR_MODEL_COLOR
	COLOR_MODEL_CMYK
	COLOR_MODEL_CMY
	COLOR_MODEL_KCMY
	COLOR_MODEL_CMY_K
	COLOR_MODEL_BLACK
	COLOR_MODEL_GRAYSCALE
	COLOR_MODEL_RGB
	COLOR_MODEL_RGB16
	COLOR_MODEL_RGBA
	COLOR_MODEL_COLORMODE_COLOR
	COLOR_MODEL_COLORMODE_MONOCHROME
	COLOR_MODEL_HP_COLOR_COLOR
	COLOR_MODEL_HP_COLOR_BLACK
	COLOR_MODEL_PRINTOUTMODE_NORMAL
	COLOR_MODEL_PRINTOUTMODE_NORMAL_GRAY
	COLOR_MODEL_PROCESSCOLORMODEL_CMYK
	COLOR_MODEL_PROCESSCOLORMODEL_GREYSCALE
	COLOR_MODEL_PROCESSCOLORMODEL_RGB
)

// TCefJsonParserOptions
//
//	Options that can be passed to CefParseJSON.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_parser_options_t))</a>
type TCefJsonParserOptions = int32

const (
	JSON_PARSER_RFC TCefJsonParserOptions = iota
	JSON_PARSER_ALLOW_TRAILING_COMMAS
)

// TCefXmlEncodingType
//
//	Supported XML encoding types. The parser supports ASCII, ISO-8859-1, and UTF16 (LE and BE) by default. All other types must be translated to UTF8 before being passed to the parser. If a BOM is detected and the correct decoder is available then that decoder will be used automatically.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_xml_encoding_type_t))</a>
type TCefXmlEncodingType = int32

const (
	XML_ENCODING_NONE TCefXmlEncodingType = iota
	XML_ENCODING_UTF8
	XML_ENCODING_UTF16LE
	XML_ENCODING_UTF16BE
	XML_ENCODING_ASCII
)

// TCefXmlNodeType
//
//	XML node types.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_xml_node_type_t))</a>
type TCefXmlNodeType = int32

const (
	XML_NODE_UNSUPPORTED TCefXmlNodeType = iota
	XML_NODE_PROCESSING_INSTRUCTION
	XML_NODE_DOCUMENT_TYPE
	XML_NODE_ELEMENT_START
	XML_NODE_ELEMENT_END
	XML_NODE_ATTRIBUTE
	XML_NODE_TEXT
	XML_NODE_CDATA
	XML_NODE_ENTITY_REFERENCE
	XML_NODE_WHITESPACE
	XML_NODE_COMMENT
)

// TCefDomEventPhase
//
//	DOM event processing phases.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_phase_t))</a>
type TCefDomEventPhase = int32

const (
	DOM_EVENT_PHASE_UNKNOWN TCefDomEventPhase = iota
	DOM_EVENT_PHASE_CAPTURING
	DOM_EVENT_PHASE_AT_TARGET
	DOM_EVENT_PHASE_BUBBLING
)

// TCefButtonState
//
//	Specifies the button display state.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_button_state_t))</a>
type TCefButtonState = int32

const (
	CEF_BUTTON_STATE_NORMAL TCefButtonState = iota
	CEF_BUTTON_STATE_HOVERED
	CEF_BUTTON_STATE_PRESSED
	CEF_BUTTON_STATE_DISABLED
)

// TCefHorizontalAlignment
//
//	Specifies the horizontal text alignment mode.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_horizontal_alignment_t))</a>
type TCefHorizontalAlignment = int32

const (
	CEF_HORIZONTAL_ALIGNMENT_LEFT TCefHorizontalAlignment = iota
	CEF_HORIZONTAL_ALIGNMENT_CENTER
	CEF_HORIZONTAL_ALIGNMENT_RIGHT
)

// TCefMenuAnchorPosition
//
//	Specifies how a menu will be anchored for non-RTL languages. The opposite position will be used for RTL languages.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_anchor_position_t))</a>
type TCefMenuAnchorPosition = int32

const (
	CEF_MENU_ANCHOR_TOPLEFT TCefMenuAnchorPosition = iota
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

// TCefMenuColorType
//
//	Supported color types for menu items.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_color_type_t))</a>
type TCefMenuColorType = int32

const (
	CEF_MENU_COLOR_TEXT TCefMenuColorType = iota
	CEF_MENU_COLOR_TEXT_HOVERED
	CEF_MENU_COLOR_TEXT_ACCELERATOR
	CEF_MENU_COLOR_TEXT_ACCELERATOR_HOVERED
	CEF_MENU_COLOR_BACKGROUND
	CEF_MENU_COLOR_BACKGROUND_HOVERED
	CEF_MENU_COLOR_COUNT
)

// TCefCompositionUnderlineStyle
//
//	Composition underline style.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_composition_underline_style_t))</a>
type TCefCompositionUnderlineStyle = int32

const (
	CEF_CUS_SOLID TCefCompositionUnderlineStyle = iota
	CEF_CUS_DOT
	CEF_CUS_DASH
	CEF_CUS_NONE
)

// TCefPermissionRequestResult
//
//	Permission request results.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_permission_request_result_t))</a>
type TCefPermissionRequestResult = int32

const (
	CEF_PERMISSION_RESULT_ACCEPT TCefPermissionRequestResult = iota
	CEF_PERMISSION_RESULT_DENY
	CEF_PERMISSION_RESULT_DISMISS
	CEF_PERMISSION_RESULT_IGNORE
)

// TCefPreferencesType
//
//	Preferences type passed to ICefBrowserProcessHandler.OnRegisterCustomPreferences.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_preferences_type_t))</a>
type TCefPreferencesType = int32

const (
	CEF_PREFERENCES_TYPE_GLOBAL TCefPreferencesType = iota
	CEF_PREFERENCES_TYPE_REQUEST_CONTEXT
)

// TCefGestureCommand
//
//	Specifies the gesture commands.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_gesture_command_t))</a>
type TCefGestureCommand = int32

const (
	CEF_GESTURE_COMMAND_BACK TCefGestureCommand = iota
	CEF_GESTURE_COMMAND_FORWARD
)

// TCefZoomCommand
//
//	Specifies the zoom commands supported by ICefBrowserHost.Zoom.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_zoom_command_t))</a>
type TCefZoomCommand = int32

const (
	CEF_ZOOM_COMMAND_OUT TCefZoomCommand = iota
	CEF_ZOOM_COMMAND_RESET
	CEF_ZOOM_COMMAND_IN
)

// TCefTestCertType
//
//	Specifies the gesture commands.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_test_cert_type_t))</a>
type TCefTestCertType = int32

const (
	CEF_TEST_CERT_OK_IP TCefTestCertType = iota
	CEF_TEST_CERT_OK_DOMAIN
	CEF_TEST_CERT_EXPIRED
)

// TCefChromePageActionIconType
//
//	Chrome page action icon types. Should be kept in sync with Chromium's PageActionIconType type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_page_action_icon_type_t))</a>
type TCefChromePageActionIconType = int32

const (
	CEF_CPAIT_BOOKMARK_STAR TCefChromePageActionIconType = iota
	CEF_CPAIT_CLICK_TO_CALL
	CEF_CPAIT_COOKIE_CONTROLS
	CEF_CPAIT_FILE_SYSTEM_ACCESS
	CEF_CPAIT_FIND
	CEF_CPAIT_HIGH_EFFICIENCY
	CEF_CPAIT_INTENT_PICKER
	CEF_CPAIT_LOCAL_CARD_MIGRATION
	CEF_CPAIT_MANAGE_PASSWORDS
	CEF_CPAIT_PAYMENTS_OFFER_NOTIFICATION
	CEF_CPAIT_PRICE_TRACKING
	CEF_CPAIT_PWA_INSTALL
	CEF_CPAIT_QR_CODE_GENERATOR
	CEF_CPAIT_READER_MODE
	CEF_CPAIT_SAVE_AUTOFILL_ADDRESS
	CEF_CPAIT_SAVE_CARD
	CEF_CPAIT_SEND_TAB_TO_SELF
	CEF_CPAIT_SHARING_HUB
	CEF_CPAIT_SIDE_SEARCH
	CEF_CPAIT_SMS_REMOTE_FETCHER
	CEF_CPAIT_TRANSLATE
	CEF_CPAIT_VIRTUAL_CARD_ENROLL
	CEF_CPAIT_VIRTUAL_CARD_MANUAL_FALLBACK
	CEF_CPAIT_ZOOM
	CEF_CPAIT_SAVE_IBAN
	CEF_CPAIT_MANDATORY_REAUTH
	CEF_CPAIT_PRICE_INSIGHTS
)

// TCefChromeToolbarButtonType
//
//	Chrome toolbar button types. Should be kept in sync with CEF's internal ToolbarButtonType type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_chrome_toolbar_button_type_t))</a>
type TCefChromeToolbarButtonType = int32

const (
	CEF_CTBT_CAST TCefChromeToolbarButtonType = iota
	CEF_CTBT_DOWNLOAD
	CEF_CTBT_SEND_TAB_TO_SELF
	CEF_CTBT_SIDE_PANEL
)

// TCefBaseTime
//
//	Represents a wall clock time in UTC. Values are not guaranteed to be monotonically non-decreasing and are subject to large amounts of skew. Time is stored internally as microseconds since the Windows epoch (1601). This is equivalent of Chromium `base::Time` (see base/time/time.h).
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_basetime_t))</a>
type TCefBaseTime = int64

// TCefContentSettingTypes
//
//	Supported content setting types. Some types are platform-specific or only supported with the Chrome runtime. Should be kept in sync with Chromium's ContentSettingsType type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_content_settings.h">CEF source file: /include/internal/cef_types_content_settings.h (cef_content_setting_types_t))</a>
type TCefContentSettingTypes = int32

const (
	CEF_CONTENT_SETTING_TYPE_COOKIES TCefContentSettingTypes = iota
	CEF_CONTENT_SETTING_TYPE_IMAGES
	CEF_CONTENT_SETTING_TYPE_JAVASCRIPT
	CEF_CONTENT_SETTING_TYPE_POPUPS
	CEF_CONTENT_SETTING_TYPE_GEOLOCATION
	CEF_CONTENT_SETTING_TYPE_NOTIFICATIONS
	CEF_CONTENT_SETTING_TYPE_AUTO_SELECT_CERTIFICATE
	CEF_CONTENT_SETTING_TYPE_MIXEDSCRIPT
	CEF_CONTENT_SETTING_TYPE_MEDIASTREAM_MIC
	CEF_CONTENT_SETTING_TYPE_MEDIASTREAM_CAMERA
	CEF_CONTENT_SETTING_TYPE_PROTOCOL_HANDLERS
	CEF_CONTENT_SETTING_TYPE_DEPRECATED_PPAPI_BROKER
	CEF_CONTENT_SETTING_TYPE_AUTOMATIC_DOWNLOADS
	CEF_CONTENT_SETTING_TYPE_MIDI_SYSEX
	CEF_CONTENT_SETTING_TYPE_SSL_CERT_DECISIONS
	CEF_CONTENT_SETTING_TYPE_PROTECTED_MEDIA_IDENTIFIER
	CEF_CONTENT_SETTING_TYPE_APP_BANNER
	CEF_CONTENT_SETTING_TYPE_SITE_ENGAGEMENT
	CEF_CONTENT_SETTING_TYPE_DURABLE_STORAGE
	CEF_CONTENT_SETTING_TYPE_USB_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_GUARD
	CEF_CONTENT_SETTING_TYPE_BACKGROUND_SYNC
	CEF_CONTENT_SETTING_TYPE_AUTOPLAY
	CEF_CONTENT_SETTING_TYPE_IMPORTANT_SITE_INFO
	CEF_CONTENT_SETTING_TYPE_PERMISSION_AUTOBLOCKER_DATA
	CEF_CONTENT_SETTING_TYPE_ADS
	CEF_CONTENT_SETTING_TYPE_ADS_DATA
	CEF_CONTENT_SETTING_TYPE_MIDI
	CEF_CONTENT_SETTING_TYPE_PASSWORD_PROTECTION
	CEF_CONTENT_SETTING_TYPE_MEDIA_ENGAGEMENT
	CEF_CONTENT_SETTING_TYPE_SOUND
	CEF_CONTENT_SETTING_TYPE_CLIENT_HINTS
	CEF_CONTENT_SETTING_TYPE_SENSORS
	CEF_CONTENT_SETTING_TYPE_ACCESSIBILITY_EVENTS
	CEF_CONTENT_SETTING_TYPE_PAYMENT_HANDLER
	CEF_CONTENT_SETTING_TYPE_USB_GUARD
	CEF_CONTENT_SETTING_TYPE_BACKGROUND_FETCH
	CEF_CONTENT_SETTING_TYPE_INTENT_PICKER_DISPLAY
	CEF_CONTENT_SETTING_TYPE_IDLE_DETECTION
	CEF_CONTENT_SETTING_TYPE_SERIAL_GUARD
	CEF_CONTENT_SETTING_TYPE_SERIAL_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_PERIODIC_BACKGROUND_SYNC
	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_SCANNING
	CEF_CONTENT_SETTING_TYPE_HID_GUARD
	CEF_CONTENT_SETTING_TYPE_HID_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_WAKE_LOCK_SCREEN
	CEF_CONTENT_SETTING_TYPE_WAKE_LOCK_SYSTEM
	CEF_CONTENT_SETTING_TYPE_LEGACY_COOKIE_ACCESS
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_WRITE_GUARD
	CEF_CONTENT_SETTING_TYPE_NFC
	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_CLIPBOARD_READ_WRITE
	CEF_CONTENT_SETTING_TYPE_CLIPBOARD_SANITIZED_WRITE
	CEF_CONTENT_SETTING_TYPE_SAFE_BROWSING_URL_CHECK_DATA
	CEF_CONTENT_SETTING_TYPE_VR
	CEF_CONTENT_SETTING_TYPE_AR
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_READ_GUARD
	CEF_CONTENT_SETTING_TYPE_STORAGE_ACCESS
	CEF_CONTENT_SETTING_TYPE_CAMERA_PAN_TILT_ZOOM
	CEF_CONTENT_SETTING_TYPE_WINDOW_MANAGEMENT
	CEF_CONTENT_SETTING_TYPE_INSECURE_PRIVATE_NETWORK
	CEF_CONTENT_SETTING_TYPE_LOCAL_FONTS
	CEF_CONTENT_SETTING_TYPE_PERMISSION_AUTOREVOCATION_DATA
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_LAST_PICKED_DIRECTORY
	CEF_CONTENT_SETTING_TYPE_DISPLAY_CAPTURE
	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_ACCESS_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_SHARING
	CEF_CONTENT_SETTING_TYPE_JAVASCRIPT_JIT
	CEF_CONTENT_SETTING_TYPE_HTTP_ALLOWED
	CEF_CONTENT_SETTING_TYPE_FORMFILL_METADATA
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_ACTIVE_SESSION
	CEF_CONTENT_SETTING_TYPE_AUTO_DARK_WEB_CONTENT
	CEF_CONTENT_SETTING_TYPE_REQUEST_DESKTOP_SITE
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_API
	CEF_CONTENT_SETTING_TYPE_NOTIFICATION_INTERACTIONS
	CEF_CONTENT_SETTING_TYPE_REDUCED_ACCEPT_LANGUAGE
	CEF_CONTENT_SETTING_TYPE_NOTIFICATION_PERMISSION_REVIEW
	CEF_CONTENT_SETTING_TYPE_PRIVATE_NETWORK_GUARD
	CEF_CONTENT_SETTING_TYPE_PRIVATE_NETWORK_CHOOSER_DATA
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_IDENTITY_PROVIDER_SIGNIN_STATUS
	CEF_CONTENT_SETTING_TYPE_REVOKED_UNUSED_SITE_PERMISSIONS
	CEF_CONTENT_SETTING_TYPE_TOP_LEVEL_STORAGE_ACCESS
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_AUTO_REAUTHN_PERMISSION
	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_IDENTITY_PROVIDER_REGISTRATION
	CEF_CONTENT_SETTING_TYPE_ANTI_ABUSE
	CEF_CONTENT_SETTING_TYPE_THIRD_PARTY_STORAGE_PARTITIONING
	CEF_CONTENT_SETTING_TYPE_HTTPS_ENFORCED
	CEF_CONTENT_SETTING_TYPE_ALL_SCREEN_CAPTURE
	CEF_CONTENT_SETTING_TYPE_COOKIE_CONTROLS_METADATA
	CEF_CONTENT_SETTING_TYPE_TPCD_SUPPORT
	CEF_CONTENT_SETTING_TYPE_AUTO_PICTURE_IN_PICTURE
	CEF_CONTENT_SETTING_TYPE_TPCD_METADATA_GRANTS
	CEF_CONTENT_SETTING_TYPE_NUM_TYPES
)

// TCefContentSettingValues
//
//	Supported content setting values. Should be kept in sync with Chromium's ContentSetting type.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_content_settings.h">CEF source file: /include/internal/cef_types_content_settings.h (cef_content_setting_values_t))</a>
type TCefContentSettingValues = int32

const (
	CEF_CONTENT_SETTING_VALUE_DEFAULT TCefContentSettingValues = iota
	CEF_CONTENT_SETTING_VALUE_ALLOW
	CEF_CONTENT_SETTING_VALUE_BLOCK
	CEF_CONTENT_SETTING_VALUE_ASK
	CEF_CONTENT_SETTING_VALUE_SESSION_ONLY
	CEF_CONTENT_SETTING_VALUE_DETECT_IMPORTANT_CONTENT
	CEF_CONTENT_SETTING_VALUE_NUM_VALUES
)
