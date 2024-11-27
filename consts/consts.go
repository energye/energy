//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package consts CEF const
package consts

import (
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/energy/homedir"
	"math"
	"os"
	"path/filepath"
	"reflect"
)

var (
	CurrentExecuteDir string                       // 当前执行目录
	ExeDir            string                       // 执行文件所在目录
	ExePath           string                       // 执行文件所在完整目录
	ExeName           string                       // 执行文件名称
	HomeDir, _        = homedir.Dir()              // 当前系统用户目录
	Separator         = string(filepath.Separator) // 平台目录分隔符
)

const (
	ENERGY_HOME_KEY = "ENERGY_HOME"
	MemoryNetwork   = "unix"
)

func init() {
	CurrentExecuteDir, _ = os.Getwd()
	ExePath = os.Args[0]
	ExeDir, ExeName = filepath.Split(ExePath)
}

type TCefMenuAccelerator = int32

const (
	MA_Shift                          = "SHIFT"
	MA_Shift_Code TCefMenuAccelerator = 0x10 //  16
	MA_Ctrl                           = "CTRL"
	MA_Ctrl_Code  TCefMenuAccelerator = 0x11 //  17
	MA_Alt                            = "ALT"
	MA_Alt_Code   TCefMenuAccelerator = 0x12 //  18
)

// Proc Get value and Set value
const (
	GetValue uintptr = 0
	SetValue uintptr = 1
)

//0:net 1:unix
type IPC_TYPE = types.Int8

const (
	IPCT_NET = IPC_TYPE(iota)
	IPCT_UNIX
)

type ChannelType = types.Int8

const (
	Ct_Server = ChannelType(iota)
	Ct_Client
)

// SpecificVersion 特定版本: CEF49，CEF87，CEF106，CEF109
type SpecificVersion = types.Int32

const (
	SV_INVALID = SpecificVersion(-1)  // 无效
	SV_CEF     = SpecificVersion(0)   // 非特定版本，当前版本或当前最新版本
	SV_CEF49   = SpecificVersion(49)  // 特定 WindowsXP
	SV_CEF87   = SpecificVersion(87)  // 特定 Flash
	SV_CEF106  = SpecificVersion(106) // 特定 Linux GTK2
	SV_CEF109  = SpecificVersion(109) // 特定 Windows 7, 8/8.1 and Windows Server 2012
)

// 功能和消息常量
const (
	WM_APP                   = 0x008000
	MINIBROWSER_SHOWDEVTOOLS = WM_APP + 0x00101 //开发工具展示
	MINIBROWSER_HIDEDEVTOOLS = WM_APP + 0x00102 //开发工具隐藏
)

// 缩放、放大
type ZOOM = types.Int32

const (
	ZOOM_RESET = ZOOM(iota)
	ZOOM_INC
	ZOOM_DEC
)

// TCefTerminationStatus
//
//	进程结束的状态
//	/include/internal/cef_types.h (cef_termination_status_t)
type TCefTerminationStatus = types.Int32

const (
	TS_ABNORMAL_TERMINATION = TCefTerminationStatus(iota)
	TS_PROCESS_WAS_KILLED
	TS_PROCESS_CRASHED
	TS_PROCESS_OOM
)

// 前进 & 后退
type BF = types.Int32

const (
	BF_GOBACK = BF(iota)
	BF_GOFORWARD
)

// 日志等级
type LogSeverity = types.UInt32

const (
	LOGSEVERITY_DEFAULT LogSeverity = 0
	LOGSEVERITY_VERBOSE LogSeverity = 1
	LOGSEVERITY_DEBUG   LogSeverity = LOGSEVERITY_VERBOSE
	LOGSEVERITY_INFO    LogSeverity = 2
	LOGSEVERITY_WARNING LogSeverity = 3
	LOGSEVERITY_ERROR   LogSeverity = 4
	LOGSEVERITY_FATAL   LogSeverity = 5
	LOGSEVERITY_DISABLE LogSeverity = 99
)

type LANGUAGE = types.String

const (
	LANGUAGE_zh_CN  LANGUAGE = "zh-CN"
	LANGUAGE_zh_TW  LANGUAGE = "zh-TW"
	LANGUAGE_am     LANGUAGE = "am"
	LANGUAGE_ar     LANGUAGE = "ar"
	LANGUAGE_bg     LANGUAGE = "bg"
	LANGUAGE_bn     LANGUAGE = "bn"
	LANGUAGE_ca     LANGUAGE = "ca"
	LANGUAGE_cs     LANGUAGE = "cs"
	LANGUAGE_da     LANGUAGE = "da"
	LANGUAGE_de     LANGUAGE = "de"
	LANGUAGE_el     LANGUAGE = "el"
	LANGUAGE_en_GB  LANGUAGE = "en-GB"
	LANGUAGE_en_US  LANGUAGE = "en-US"
	LANGUAGE_es     LANGUAGE = "es"
	LANGUAGE_es_419 LANGUAGE = "es-419"
	LANGUAGE_et     LANGUAGE = "et"
	LANGUAGE_fa     LANGUAGE = "fa"
	LANGUAGE_fi     LANGUAGE = "fi"
	LANGUAGE_fil    LANGUAGE = "fil"
	LANGUAGE_fr     LANGUAGE = "fr"
	LANGUAGE_gu     LANGUAGE = "gu"
	LANGUAGE_he     LANGUAGE = "he"
	LANGUAGE_hi     LANGUAGE = "hi"
	LANGUAGE_hr     LANGUAGE = "hr"
	LANGUAGE_hu     LANGUAGE = "hu"
	LANGUAGE_id     LANGUAGE = "channelId"
	LANGUAGE_it     LANGUAGE = "it"
	LANGUAGE_ja     LANGUAGE = "ja"
	LANGUAGE_kn     LANGUAGE = "kn"
	LANGUAGE_ko     LANGUAGE = "ko"
	LANGUAGE_lt     LANGUAGE = "lt"
	LANGUAGE_lv     LANGUAGE = "lv"
	LANGUAGE_ml     LANGUAGE = "ml"
	LANGUAGE_mr     LANGUAGE = "mr"
	LANGUAGE_ms     LANGUAGE = "ms"
	LANGUAGE_nb     LANGUAGE = "nb"
	LANGUAGE_nl     LANGUAGE = "nl"
	LANGUAGE_pl     LANGUAGE = "pl"
	LANGUAGE_pt_BR  LANGUAGE = "pt-BR"
	LANGUAGE_pt_PT  LANGUAGE = "pt-PT"
	LANGUAGE_ro     LANGUAGE = "ro"
	LANGUAGE_ru     LANGUAGE = "ru"
	LANGUAGE_sk     LANGUAGE = "sk"
	LANGUAGE_sl     LANGUAGE = "sl"
	LANGUAGE_sr     LANGUAGE = "sr"
	LANGUAGE_sv     LANGUAGE = "sv"
	LANGUAGE_sw     LANGUAGE = "sw"
	LANGUAGE_ta     LANGUAGE = "ta"
	LANGUAGE_te     LANGUAGE = "te"
	LANGUAGE_th     LANGUAGE = "th"
	LANGUAGE_tr     LANGUAGE = "tr"
	LANGUAGE_uk     LANGUAGE = "uk"
	LANGUAGE_vi     LANGUAGE = "vi"
)

// TCefCloseBrowserAction
// Chromium关闭的操作类型
// 在 TChromium.Onclose 使用
// -------------------------
// cbaCancel : 停止关闭浏览器
// cbaClose  : 继续关闭浏览器
// cbaDelay  : 暂时停止关闭浏览器
//
//	: 当应用程序需要在关闭浏览器之前执行一些自定义进程时使用。在关闭浏览器之前，通常需要在主线程中销毁TCEFWindowParent。
type TCefCloseBrowserAction = types.Int32

const (
	CbaClose = TCefCloseBrowserAction(iota)
	CbaDelay
	CbaCancel
)

// CefProcessId CEF 进程 Id
type CefProcessId = types.Int32

const (
	PID_BROWSER = CefProcessId(iota)
	PID_RENDER
)

// Go Kind 扩展常量
const (
	SLICE_BYTE reflect.Kind = iota + 30 // []byte
	JD                                  // JsonData
	NIL                                 // nil
)

// FN_TYPE 函数类型
type FN_TYPE = types.Int8

const (
	FN_TYPE_COMMON = FN_TYPE(iota) //普通函数，直接定义的
	FN_TYPE_OBJECT                 //对象函数，所属对象
)

// IS_CO 通用类型或对象类型
type IS_CO = types.Int8

const (
	IS_COMMON = IS_CO(iota)
	IS_OBJECT
)

// 进程消息类型
type PROCESS_MESSAGE_TYPE = types.Int8

const (
	PMT_JS_CODE = PROCESS_MESSAGE_TYPE(iota) //执行JS代码消息
	PMT_TEXT                                 //文本传递消息
	PMT_BINARY                               //二进制消息
)

type TDateTime = types.Float64

// include/internal/cef_types.h (cef_cookie_same_site_t)
type TCefCookieSameSite = types.Int32

const (
	Ccss_CEF_COOKIE_SAME_SITE_UNSPECIFIED = TCefCookieSameSite(iota)
	Ccss_CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	Ccss_CEF_COOKIE_SAME_SITE_LAX_MODE
	Ccss_CEF_COOKIE_SAME_SITE_STRICT_MODE
)

// include/internal/cef_types.h (cef_cookie_priority_t)
type TCefCookiePriority = types.Int32

const (
	CEF_COOKIE_PRIORITY_LOW    TCefCookiePriority = -1
	CEF_COOKIE_PRIORITY_MEDIUM TCefCookiePriority = 0
	CEF_COOKIE_PRIORITY_HIGH   TCefCookiePriority = 1
)

type TCefProxyType = types.Int32

const (
	PtDirect       = TCefProxyType(iota) // mode dict => direct
	PtAutodetect                         // mode dict => auto_detect
	PtSystem                             // mode dict => system
	PtFixedServers                       // mode dict => fixed_servers
	PtPACScript                          // mode dict => pac_script
)

type TCefProxyScheme = types.Int32

const (
	PsHTTP = TCefProxyScheme(iota)
	PsSOCKS4
	PsSOCKS5
)

type TCefContextMenuType = types.Int32

const (
	CMT_NONE = TCefContextMenuType(iota)
	CMT_CHECK
	CMT_RADIO
)

// include/internal/cef_types.h (cef_context_menu_media_type_t)
type TCefContextMenuMediaType = types.Int32

const (
	CM_MEDIATYPE_NONE = TCefContextMenuMediaType(iota)
	CM_MEDIATYPE_IMAGE
	CM_MEDIATYPE_VIDEO
	CM_MEDIATYPE_AUDIO
	CM_MEDIATYPE_CANVAS
	CM_MEDIATYPE_FILE
	CM_MEDIATYPE_PLUGIN
)

type MenuId = int32

const (
	MENU_ID_BACK                       MenuId = 100
	MENU_ID_FORWARD                    MenuId = 101
	MENU_ID_RELOAD                     MenuId = 102
	MENU_ID_RELOAD_NOCACHE             MenuId = 103
	MENU_ID_STOPLOAD                   MenuId = 104
	MENU_ID_UNDO                       MenuId = 110
	MENU_ID_REDO                       MenuId = 111
	MENU_ID_CUT                        MenuId = 112
	MENU_ID_COPY                       MenuId = 113
	MENU_ID_PASTE                      MenuId = 114
	MENU_ID_DELETE                     MenuId = 115
	MENU_ID_SELECT_ALL                 MenuId = 116
	MENU_ID_FIND                       MenuId = 130
	MENU_ID_PRINT                      MenuId = 131
	MENU_ID_VIEW_SOURCE                MenuId = 132
	MENU_ID_SPELLCHECK_SUGGESTION_0    MenuId = 200
	MENU_ID_SPELLCHECK_SUGGESTION_1    MenuId = 201
	MENU_ID_SPELLCHECK_SUGGESTION_2    MenuId = 202
	MENU_ID_SPELLCHECK_SUGGESTION_3    MenuId = 203
	MENU_ID_SPELLCHECK_SUGGESTION_4    MenuId = 204
	MENU_ID_SPELLCHECK_SUGGESTION_LAST MenuId = 204
	MENU_ID_NO_SPELLING_SUGGESTIONS    MenuId = 205
	MENU_ID_ADD_TO_DICTIONARY          MenuId = 206
	MENU_ID_CUSTOM_FIRST               MenuId = 220
	MENU_ID_CUSTOM_LAST                MenuId = 250
	MENU_ID_USER_FIRST                 MenuId = 26500
	MENU_ID_USER_LAST                  MenuId = 28500
)

// include/internal/cef_types.h (cef_menu_color_type_t)
type TCefMenuColorType = types.Int32

const (
	CEF_MENU_COLOR_TEXT = TCefMenuColorType(iota)
	CEF_MENU_COLOR_TEXT_HOVERED
	CEF_MENU_COLOR_TEXT_ACCELERATOR
	CEF_MENU_COLOR_TEXT_ACCELERATOR_HOVERED
	CEF_MENU_COLOR_BACKGROUND
	CEF_MENU_COLOR_BACKGROUND_HOVERED
	CEF_MENU_COLOR_COUNT
)

type ARGB = types.UInt32

// include/internal/cef_types.h (cef_key_event_type_t)
type TCefKeyEventType = types.Int32

const (
	KEYEVENT_RAW_KEYDOWN = TCefKeyEventType(iota)
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

// include/internal/cef_types.h (cef_event_flags_t)
type TCefEventFlags = types.UInt32

const (
	EVENTFLAG_NONE                TCefEventFlags = 0
	EVENTFLAG_CAPS_LOCK_ON                       = 1 << 0
	EVENTFLAG_SHIFT_DOWN                         = 1 << 1
	EVENTFLAG_CONTROL_DOWN                       = 1 << 2
	EVENTFLAG_ALT_DOWN                           = 1 << 3
	EVENTFLAG_LEFT_MOUSE_BUTTON                  = 1 << 4
	EVENTFLAG_MIDDLE_MOUSE_BUTTON                = 1 << 5
	EVENTFLAG_RIGHT_MOUSE_BUTTON                 = 1 << 6
)

// Mac OS-X command key
const (
	EVENTFLAG_COMMAND_DOWN TCefEventFlags = 1 << 7
	EVENTFLAG_NUM_LOCK_ON                 = 1 << 8
	EVENTFLAG_IS_KEY_PAD                  = 1 << 9
	EVENTFLAG_IS_LEFT                     = 1 << 10
	EVENTFLAG_IS_RIGHT                    = 1 << 11
	EVENTFLAG_ALTGR_DOWN                  = 1 << 12
	EVENTFLAG_IS_REPEAT                   = 1 << 13
)

type TCefWindowHandleType = types.Int8

// include/internal/cef_types_win.h (cef_window_handle_t)
// include/internal/cef_types_mac.h (cef_window_handle_t)
// include/internal/cef_types_linux.h (cef_window_handle_t)
type TCefWindowHandle = types.UIntptr

const (
	Wht_WindowParent = TCefWindowHandleType(iota)
	Wht_LinkedWindowParent
)

// include/internal/cef_types.h (cef_return_value_t)
type TCefReturnValue int32

const (
	RV_CANCEL = TCefReturnValue(iota)
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

// include/internal/cef_types.h (cef_referrer_policy_t)
type TCefReferrerPolicy = types.Int32

const (
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_FROM_SECURE_TO_INSECURE = TCefReferrerPolicy(iota) // same value as REFERRER_POLICY_DEFAULT
	REFERRER_POLICY_REDUCE_REFERRER_GRANULARITY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_ONLY_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_NEVER_CLEAR_REFERRER
	REFERRER_POLICY_ORIGIN
	REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_CROSS_ORIGIN
	REFERRER_POLICY_ORIGIN_CLEAR_ON_TRANSITION_FROM_SECURE_TO_INSECURE
	REFERRER_POLICY_NO_REFERRER // REFERRER_POLICY_LAST_VALUE = REFERRER_POLICY_NO_REFERRER
)

// include/internal/cef_types.h (cef_urlrequest_flags_t)
type TCefUrlRequestFlags = types.Int

// include/internal/cef_types.h (cef_errorcode_t)
type TCefErrorCode = types.Int32

// include/internal/cef_types.h (cef_resource_type_t)
type TCefResourceType = types.Int32

const (
	RT_MAIN_FRAME                    = TCefResourceType(iota) // Top level page.
	RT_SUB_FRAME                                              // Frame or iframe.
	RT_STYLESHEET                                             // CSS stylesheet.
	RT_SCRIPT                                                 // External script.
	RT_IMAGE                                                  // Image (jpg/gif/png/etc).
	RT_FONT_RESOURCE                                          // Font.
	RT_SUB_RESOURCE                                           // Some other subresource. This is the default type if the actual type is unknown
	RT_OBJECT                                                 // Object (or embed) tag for a plugin, or a resource that a plugin requested.
	RT_MEDIA                                                  // Media resource.
	RT_WORKER                                                 // Main resource of a dedicated worker.
	RT_SHARED_WORKER                                          // Main resource of a shared worker.
	RT_PREFETCH                                               // Explicitly requested prefetch.
	RT_FAVICON                                                // Favicon
	RT_XHR                                                    // XMLHttpRequest
	RT_PING                                                   // A request for a "<ping>".
	RT_SERVICE_WORKER                                         // Main resource of a service worker.
	RT_CSP_REPORT                                             // A report of Content Security Policy violations.
	RT_PLUGIN_RESOURCE                                        // A resource that a plugin requested.
	RT_EMPTY_FILLER_TYPE_DO_NOT_USE                           // This type doesn't exist in CEF and it's here just to fill this position.
	RT_NAVIGATION_PRELOAD_MAIN_FRAME                          // A main-frame service worker navigation preload request.This type must have a value of 19
	RT_NAVIGATION_PRELOAD_SUB_FRAME                           // A sub-frame service worker navigation preload request.
)

// include/internal/cef_types.h (cef_transition_type_t)
type TCefTransitionType = types.Int

// include/internal/cef_types.h (cef_urlrequest_status_t)
type TCefUrlRequestStatus = types.Int32

const (
	UR_UNKNOWN = TCefUrlRequestStatus(iota)
	UR_SUCCESS
	UR_IO_PENDING
	UR_CANCELED
	UR_FAILED
)

// Represents the state of a setting.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_state_t)</see></para>
type TCefState = types.Int32

const (
	// Use the default state for the setting.
	STATE_DEFAULT = TCefState(iota)
	// Enable or allow the setting.
	STATE_ENABLED
	// Disable or disallow the setting.
	STATE_DISABLED
)

// Default values for the Windowsless framerate setting in TChromiumOptions
// The values are frames per second.

const CEF_OSR_FRAMERATE_DEFAULT = 30                 // Used when the shared textures are disabled.
const CEF_OSR_SHARED_TEXTURES_FRAMERATE_DEFAULT = 60 // Used when the shared textures are enabled.

const CEF_TIMER_MINIMUM = 0x0000000A

const CEF_TIMER_MAXIMUM = 0x7FFFFFFF

const CEF_TIMER_MAXDELAY = 1000 / CEF_OSR_FRAMERATE_DEFAULT

const CEF_TIMER_DEPLETEWORK_CYCLES = 10

const CEF_TIMER_DEPLETEWORK_DELAY = 50

// include/internal/cef_types.h (cef_touch_event_type_t)
type TCefTouchEeventType = types.Int32

const (
	CEF_TET_RELEASED = TCefTouchEeventType(iota)
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

// include/internal/cef_types.h (cef_pointer_type_t)
type TCefPointerType = types.Int32

const (
	CEF_POINTER_TYPE_TOUCH = TCefPointerType(iota)
	CEF_POINTER_TYPE_MOUSE
	CEF_POINTER_TYPE_PEN
	CEF_POINTER_TYPE_ERASER
	CEF_POINTER_TYPE_UNKNOWN
)

// include/internal/cef_types.h (cef_mouse_button_type_t)
type TCefMouseButtonType = types.Int32

const (
	MBT_LEFT = TCefMouseButtonType(iota)
	MBT_MIDDLE
	MBT_RIGHT
)

// 进程消息错误码
type ProcessMessageError = types.Int32

const (
	PME_OK                        ProcessMessageError = iota + 1 //发送成功
	PMErr_NOT_FOUND_FRAME         ProcessMessageError = -1       //没找到Frame
	PMErr_TARGET_PROCESS          ProcessMessageError = -2       //目标进程标识错误
	PMErr_NAME_IS_NULL            ProcessMessageError = -3       //消息名称为空
	PMErr_NO_INVALID_FRAME        ProcessMessageError = -4       //无效的Frame
	PMErr_REQUIRED_PARAMS_IS_NULL ProcessMessageError = -5       //必要参数为空
	PMErr_NAME_CANNOT_USED        ProcessMessageError = -6       //不能使用的消息名称
)

// The manner in which a link click should be opened. These constants match
// their equivalents in Chromium's window_open_disposition.h and should not be renumbered.
type TCefWindowOpenDisposition = types.Int32

const (
	// Unknown disposition.
	CEF_WOD_UNKNOWN = TCefWindowOpenDisposition(iota)
	// Current tab. This is the default in most cases.
	CEF_WOD_CURRENT_TAB
	// Indicates that only one tab with the url should exist in the same window.
	CEF_WOD_SINGLETON_TAB
	// Shift key + Middle mouse button or meta/ctrl key while clicking.
	CEF_WOD_NEW_FOREGROUND_TAB
	// Middle mouse button or meta/ctrl key while clicking.
	CEF_WOD_NEW_BACKGROUND_TAB
	// New popup window.
	CEF_WOD_NEW_POPUP
	// Shift key while clicking.
	CEF_WOD_NEW_WINDOW
	// Alt key while clicking.
	CEF_WOD_SAVE_TO_DISK
	// New off-the-record (incognito) window.
	CEF_WOD_OFF_THE_RECORD
	// Special case error condition from the renderer.
	CEF_WOD_IGNORE_ACTION
	// Activates an existing tab containing the url, rather than navigating.
	// This is similar to SINGLETON_TAB, but searches across all windows from
	// the current profile and anonymity (instead of just the current one);
	// closes the current tab on switching if the current tab was the NTP with
	// no session history; and behaves like CURRENT_TAB instead of
	// NEW_FOREGROUND_TAB when no existing tab is found.
	CEF_WOD_SWITCH_TO_TAB
	// Creates a new document picture-in-picture window showing a child WebView.
	CEF_WOD_NEW_PICTURE_IN_PICTURE
)

// WINDOW_TYPE 窗口类型
type WINDOW_TYPE = types.Int8

const (
	WT_MAIN_BROWSER      = WINDOW_TYPE(iota) // 主窗口 只允许有一个, 如果自己创建窗口需要设置为 WT_POPUP_SUB_BROWSER 子窗口选项
	WT_POPUP_SUB_BROWSER                     // 子窗口 允许有多个
	WT_DEV_TOOLS                             // 开发者工具窗口
	WT_VIEW_SOURCE                           // 显示源代码窗口
)

// include/internal/cef_types.h (cef_context_menu_type_flags_t)
type TCefContextMenuTypeFlags = types.UInt32

const (
	// No node is selected.
	CM_TYPEFLAG_NONE TCefContextMenuTypeFlags = 0
	// The top page is selected.
	CM_TYPEFLAG_PAGE = 1 << 0
	// A subframe page is selected.
	CM_TYPEFLAG_FRAME = 1 << 1
	// A link is selected.
	CM_TYPEFLAG_LINK = 1 << 2
	// A media node is selected.
	CM_TYPEFLAG_MEDIA = 1 << 3
	// There is a textual or mixed selection that is selected.
	CM_TYPEFLAG_SELECTION = 1 << 4
	// An editable element is selected.
	CM_TYPEFLAG_EDITABLE = 1 << 5
)

// include/internal/cef_types.h (cef_context_menu_media_state_flags_t)
type TCefContextMenuMediaStateFlags = types.UInt32

const (
	CM_MEDIAFLAG_NONE                   TCefContextMenuMediaStateFlags = 0
	CM_MEDIAFLAG_IN_ERROR                                              = 1 << 0
	CM_MEDIAFLAG_PAUSED                                                = 1 << 1
	CM_MEDIAFLAG_MUTED                                                 = 1 << 2
	CM_MEDIAFLAG_LOOP                                                  = 1 << 3
	CM_MEDIAFLAG_CAN_SAVE                                              = 1 << 4
	CM_MEDIAFLAG_HAS_AUDIO                                             = 1 << 5
	CM_MEDIAFLAG_CAN_TOGGLE_CONTROLS                                   = 1 << 6
	CM_MEDIAFLAG_CONTROLS                                              = 1 << 7
	CM_MEDIAFLAG_CAN_PRINT                                             = 1 << 8
	CM_MEDIAFLAG_CAN_ROTATE                                            = 1 << 9
	CM_MEDIAFLAG_CAN_PICTURE_IN_PICTURE                                = 1 << 10
	CM_MEDIAFLAG_PICTURE_IN_PICTURE                                    = 1 << 11
	CM_MEDIAFLAG_CAN_LOOP                                              = 1 << 12
)

// include/internal/cef_types.h (cef_context_menu_edit_state_flags_t)
type TCefContextMenuEditStateFlags = types.UInt32

const (
	CM_EDITFLAG_NONE            TCefContextMenuEditStateFlags = 0
	CM_EDITFLAG_CAN_UNDO                                      = 1 << 0
	CM_EDITFLAG_CAN_REDO                                      = 1 << 1
	CM_EDITFLAG_CAN_CUT                                       = 1 << 2
	CM_EDITFLAG_CAN_COPY                                      = 1 << 3
	CM_EDITFLAG_CAN_PASTE                                     = 1 << 4
	CM_EDITFLAG_CAN_DELETE                                    = 1 << 5
	CM_EDITFLAG_CAN_SELECT_ALL                                = 1 << 6
	CM_EDITFLAG_CAN_TRANSLATE                                 = 1 << 7
	CM_EDITFLAG_CAN_EDIT_RICHLY                               = 1 << 8
)

// include/internal/cef_types.h (cef_menu_anchor_position_t)
type TCefMenuAnchorPosition = types.Int32

const (
	CEF_MENU_ANCHOR_TOPLEFT = TCefMenuAnchorPosition(iota)
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

// include/internal/cef_types.h (cef_docking_mode_t)
type TCefDockingMode = types.Int32

const (
	CEF_DOCKING_MODE_TOP_LEFT TCefDockingMode = iota + 1
	CEF_DOCKING_MODE_TOP_RIGHT
	CEF_DOCKING_MODE_BOTTOM_LEFT
	CEF_DOCKING_MODE_BOTTOM_RIGHT
	CEF_DOCKING_MODE_CUSTOM
)

// include/internal/cef_types.h (cef_show_state_t)4
type TCefShowState = types.Int32

const (
	CEF_SHOW_STATE_NORMAL     TCefShowState = 1
	CEF_SHOW_STATE_MINIMIZED  TCefShowState = 2
	CEF_SHOW_STATE_MAXIMIZED  TCefShowState = 3
	CEF_SHOW_STATE_FULLSCREEN TCefShowState = 4
	CEF_SHOW_STATE_HIDDEN     TCefShowState = 5
)

// / Chrome toolbar types.
// / <para>TCefChromeToolbarType values.</para>
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_toolbar_type_t)</see></para>
type TCefChromeToolbarType = int32

const (
	CEF_CTT_NONE     TCefChromeToolbarType = 1
	CEF_CTT_NORMAL   TCefChromeToolbarType = 2
	CEF_CTT_LOCATION TCefChromeToolbarType = 3
)

// include/internal/cef_types.h (cef_drag_operations_mask_t)
type TCefDragOperations = types.Cardinal

const (
	DRAG_OPERATION_NONE    TCefDragOperations = 0
	DRAG_OPERATION_COPY    TCefDragOperations = 1 << 0
	DRAG_OPERATION_LINK    TCefDragOperations = 1 << 1
	DRAG_OPERATION_GENERIC TCefDragOperations = 1 << 2
	DRAG_OPERATION_PRIVATE TCefDragOperations = 1 << 3
	DRAG_OPERATION_MOVE    TCefDragOperations = 1 << 4
	DRAG_OPERATION_DELETE  TCefDragOperations = 1 << 5
	DRAG_OPERATION_EVERY   TCefDragOperations = math.MaxUint32
)

// include/internal/cef_types.h (cef_drag_operations_mask_t)
type TCefDragOperation = TCefDragOperations

type TrayType int8

const (
	TRAY_SYS TrayType = iota
	TRAY_LCL
)

// CombineRgn Mode flags
type RNGFnCombineMode = int32

const (
	RGN_AND RNGFnCombineMode = iota + 1
	RGN_OR
	RGN_XOR
	RGN_DIFF
	RGN_COPY
)

type TCefWebRTCHandlingPolicy = types.Int32

const (
	HpDefault TCefWebRTCHandlingPolicy = iota
	HpDefaultPublicAndPrivateInterfaces
	HpDefaultPublicInterfaceOnly
	HpDisableNonProxiedUDP
)

// Values used by the battery saver mode state preference
// https://source.chromium.org/chromium/chromium/src/+/main:components/performance_manager/public/user_tuning/prefs.h
type TCefBatterySaverModeState = types.Int32

const (
	BsmsDisabled TCefBatterySaverModeState = iota
	BsmsEnabledBelowThreshold
	BsmsEnabledOnBattery
	BsmsEnabled
	BsmsDefault // Custom value used to update the preferences only when there's a non-default value
)

// Used in TChromium preferences to allow or block cookies.
type TCefCookiePref = types.Int32

const (
	CpDefault TCefCookiePref = iota
	CpAllow
	CpBlock
)

// https://chromium.googlesource.com/chromium/src/+/refs/tags/77.0.3865.90/chrome/common/net/safe_search_util.h (YouTubeRestrictMode)
// https://www.chromium.org/administrators/policy-list-3#ForceYouTubeRestrict
type YouTubeRestrict = types.Int32

const (
	YOUTUBE_RESTRICT_OFF YouTubeRestrict = iota
	YOUTUBE_RESTRICT_MODERATE
	YOUTUBE_RESTRICT_STRICT
)

type ZoomStep = byte

const (
	ZOOM_STEP_25  ZoomStep = 0
	ZOOM_STEP_33  ZoomStep = 1
	ZOOM_STEP_50  ZoomStep = 2
	ZOOM_STEP_67  ZoomStep = 3
	ZOOM_STEP_75  ZoomStep = 4
	ZOOM_STEP_90  ZoomStep = 5
	ZOOM_STEP_100 ZoomStep = 6
	ZOOM_STEP_110 ZoomStep = 7
	ZOOM_STEP_125 ZoomStep = 8
	ZOOM_STEP_150 ZoomStep = 9
	ZOOM_STEP_175 ZoomStep = 10
	ZOOM_STEP_200 ZoomStep = 11
	ZOOM_STEP_250 ZoomStep = 12
	ZOOM_STEP_300 ZoomStep = 13
	ZOOM_STEP_400 ZoomStep = 14
	ZOOM_STEP_500 ZoomStep = 15
	ZOOM_STEP_UNK ZoomStep = 16
	ZOOM_STEP_MIN ZoomStep = ZOOM_STEP_25
	ZOOM_STEP_MAX ZoomStep = ZOOM_STEP_500
	ZOOM_STEP_DEF ZoomStep = ZOOM_STEP_100
)

// include/internal/cef_types.h (cef_v8_propertyattribute_t)
type TCefV8PropertyAttributes = types.Cardinal

const (
	V8_PROPERTY_ATTRIBUTE_NONE       TCefV8PropertyAttributes = 0
	V8_PROPERTY_ATTRIBUTE_READONLY   TCefV8PropertyAttributes = 1 << 0
	V8_PROPERTY_ATTRIBUTE_DONTENUM   TCefV8PropertyAttributes = 1 << 1
	V8_PROPERTY_ATTRIBUTE_DONTDELETE TCefV8PropertyAttributes = 1 << 2
)

// include/internal/cef_types.h (cef_value_type_t)
type TCefValueType = types.Int32

const (
	VTYPE_INVALID TCefValueType = iota
	VTYPE_NULL
	VTYPE_BOOL
	VTYPE_INT
	VTYPE_DOUBLE
	VTYPE_STRING
	VTYPE_BINARY
	VTYPE_DICTIONARY // Object
	VTYPE_LIST       // JSONArray
)

// include/internal/cef_types.h (cef_postdataelement_type_t)
type TCefPostDataElementType = types.Int32

const (
	PDE_TYPE_EMPTY TCefPostDataElementType = iota
	PDE_TYPE_BYTES
	PDE_TYPE_FILE
)

type TCefAutoplayPolicy = types.Int32

const (
	AppDefault TCefAutoplayPolicy = iota
	AppDocumentUserActivationRequired
	AppNoUserGestureRequired
	AppUserGestureRequired
)

// Values used by the --net-log-capture-mode command line switch.
// Sets the granularity of events to capture in the network log.
// https://source.chromium.org/chromium/chromium/src/+/main:content/browser/network_service_instance_impl.cc
// https://source.chromium.org/chromium/chromium/src/+/main:net/log/net_log_capture_mode.h
type TCefNetLogCaptureMode = types.Int32

const (
	NlcmDefault TCefNetLogCaptureMode = iota
	NlcmIncludeSensitive
	NlcmEverything
)

type TCefProcessType types.Int32

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

type TCefProcessTypeValue types.String

const (
	PtvBrowser  TCefProcessTypeValue = "browser"
	PtvRenderer TCefProcessTypeValue = "renderer"
	PtvZygote   TCefProcessTypeValue = "zygote"
	PtvGPU      TCefProcessTypeValue = "GPU"
	PtvUtility  TCefProcessTypeValue = "utility"
	PtvBroker   TCefProcessTypeValue = "broker"
	PtvCrashpad TCefProcessTypeValue = "crashpad"
	PtvOther    TCefProcessTypeValue = "other"
)

type TCefApplicationStatus = types.Int32

const (
	AsLoading TCefApplicationStatus = iota
	AsLoaded
	AsInitialized
	AsShuttingDown
	AsUnloaded
	AsErrorMissingFiles
	AsErrorDLLVersion
	AsErrorLoadingLibrary
	AsErrorInitializingLibrary
	AsErrorExecutingProcess
)

// net error
type CEF_NET_ERROR = types.Int32

const (
	ERR_NONE                                          CEF_NET_ERROR = 0
	ERR_IO_PENDING                                    CEF_NET_ERROR = -1
	ERR_FAILED                                        CEF_NET_ERROR = -2
	ERR_ABORTED                                       CEF_NET_ERROR = -3
	ERR_INVALID_ARGUMENT                              CEF_NET_ERROR = -4
	ERR_INVALID_HANDLE                                CEF_NET_ERROR = -5
	ERR_FILE_NOT_FOUND                                CEF_NET_ERROR = -6
	ERR_TIMED_OUT                                     CEF_NET_ERROR = -7
	ERR_FILE_TOO_BIG                                  CEF_NET_ERROR = -8
	ERR_UNEXPECTED                                    CEF_NET_ERROR = -9
	ERR_ACCESS_DENIED                                 CEF_NET_ERROR = -10
	ERR_NOT_IMPLEMENTED                               CEF_NET_ERROR = -11
	ERR_INSUFFICIENT_RESOURCES                        CEF_NET_ERROR = -12
	ERR_OUT_OF_MEMORY                                 CEF_NET_ERROR = -13
	ERR_UPLOAD_FILE_CHANGED                           CEF_NET_ERROR = -14
	ERR_SOCKET_NOT_CONNECTED                          CEF_NET_ERROR = -15
	ERR_FILE_EXISTS                                   CEF_NET_ERROR = -16
	ERR_FILE_PATH_TOO_LONG                            CEF_NET_ERROR = -17
	ERR_FILE_NO_SPACE                                 CEF_NET_ERROR = -18
	ERR_FILE_VIRUS_INFECTED                           CEF_NET_ERROR = -19
	ERR_BLOCKED_BY_CLIENT                             CEF_NET_ERROR = -20
	ERR_NETWORK_CHANGED                               CEF_NET_ERROR = -21
	ERR_BLOCKED_BY_ADMINISTRATOR                      CEF_NET_ERROR = -22
	ERR_SOCKET_IS_CONNECTED                           CEF_NET_ERROR = -23
	ERR_BLOCKED_ENROLLMENT_CHECK_PENDING              CEF_NET_ERROR = -24
	ERR_UPLOAD_STREAM_REWIND_NOT_SUPPORTED            CEF_NET_ERROR = -25
	ERR_CONTEXT_SHUT_DOWN                             CEF_NET_ERROR = -26
	ERR_BLOCKED_BY_RESPONSE                           CEF_NET_ERROR = -27
	ERR_BLOCKED_BY_XSS_AUDITOR                        CEF_NET_ERROR = -28
	ERR_CLEARTEXT_NOT_PERMITTED                       CEF_NET_ERROR = -29
	ERR_CONNECTION_CLOSED                             CEF_NET_ERROR = -100
	ERR_CONNECTION_RESET                              CEF_NET_ERROR = -101
	ERR_CONNECTION_REFUSED                            CEF_NET_ERROR = -102
	ERR_CONNECTION_ABORTED                            CEF_NET_ERROR = -103
	ERR_CONNECTION_FAILED                             CEF_NET_ERROR = -104
	ERR_NAME_NOT_RESOLVED                             CEF_NET_ERROR = -105
	ERR_INTERNET_DISCONNECTED                         CEF_NET_ERROR = -106
	ERR_SSL_PROTOCOL_ERROR                            CEF_NET_ERROR = -107
	ERR_ADDRESS_INVALID                               CEF_NET_ERROR = -108
	ERR_ADDRESS_UNREACHABLE                           CEF_NET_ERROR = -109
	ERR_SSL_CLIENT_AUTH_CERT_NEEDED                   CEF_NET_ERROR = -110
	ERR_TUNNEL_CONNECTION_FAILED                      CEF_NET_ERROR = -111
	ERR_NO_SSL_VERSIONS_ENABLED                       CEF_NET_ERROR = -112
	ERR_SSL_VERSION_OR_CIPHER_MISMATCH                CEF_NET_ERROR = -113
	ERR_SSL_RENEGOTIATION_REQUESTED                   CEF_NET_ERROR = -114
	ERR_PROXY_AUTH_UNSUPPORTED                        CEF_NET_ERROR = -115
	ERR_CERT_ERROR_IN_SSL_RENEGOTIATION               CEF_NET_ERROR = -116
	ERR_BAD_SSL_CLIENT_AUTH_CERT                      CEF_NET_ERROR = -117
	ERR_CONNECTION_TIMED_OUT                          CEF_NET_ERROR = -118
	ERR_HOST_RESOLVER_QUEUE_TOO_LARGE                 CEF_NET_ERROR = -119
	ERR_SOCKS_CONNECTION_FAILED                       CEF_NET_ERROR = -120
	ERR_SOCKS_CONNECTION_HOST_UNREACHABLE             CEF_NET_ERROR = -121
	ERR_ALPN_NEGOTIATION_FAILED                       CEF_NET_ERROR = -122
	ERR_SSL_NO_RENEGOTIATION                          CEF_NET_ERROR = -123
	ERR_WINSOCK_UNEXPECTED_WRITTEN_BYTES              CEF_NET_ERROR = -124
	ERR_SSL_DECOMPRESSION_FAILURE_ALERT               CEF_NET_ERROR = -125
	ERR_SSL_BAD_RECORD_MAC_ALERT                      CEF_NET_ERROR = -126
	ERR_PROXY_AUTH_REQUESTED                          CEF_NET_ERROR = -127
	ERR_SSL_WEAK_SERVER_EPHEMERAL_DH_KEY              CEF_NET_ERROR = -129
	ERR_PROXY_CONNECTION_FAILED                       CEF_NET_ERROR = -130
	ERR_MANDATORY_PROXY_CONFIGURATION_FAILED          CEF_NET_ERROR = -131
	ERR_PRECONNECT_MAX_SOCKET_LIMIT                   CEF_NET_ERROR = -133
	ERR_SSL_CLIENT_AUTH_PRIVATE_KEY_ACCESS_DENIED     CEF_NET_ERROR = -134
	ERR_SSL_CLIENT_AUTH_CERT_NO_PRIVATE_KEY           CEF_NET_ERROR = -135
	ERR_PROXY_CERTIFICATE_INVALID                     CEF_NET_ERROR = -136
	ERR_NAME_RESOLUTION_FAILED                        CEF_NET_ERROR = -137
	ERR_NETWORK_ACCESS_DENIED                         CEF_NET_ERROR = -138
	ERR_TEMPORARILY_THROTTLED                         CEF_NET_ERROR = -139
	ERR_HTTPS_PROXY_TUNNEL_RESPONSE_REDIRECT          CEF_NET_ERROR = -140
	ERR_SSL_CLIENT_AUTH_SIGNATURE_FAILED              CEF_NET_ERROR = -141
	ERR_MSG_TOO_BIG                                   CEF_NET_ERROR = -142
	ERR_SPDY_SESSION_ALREADY_EXISTS                   CEF_NET_ERROR = -143
	ERR_WS_PROTOCOL_ERROR                             CEF_NET_ERROR = -145
	ERR_ADDRESS_IN_USE                                CEF_NET_ERROR = -147
	ERR_SSL_HANDSHAKE_NOT_COMPLETED                   CEF_NET_ERROR = -148
	ERR_SSL_BAD_PEER_PUBLIC_KEY                       CEF_NET_ERROR = -149
	ERR_SSL_PINNED_KEY_NOT_IN_CERT_CHAIN              CEF_NET_ERROR = -150
	ERR_CLIENT_AUTH_CERT_TYPE_UNSUPPORTED             CEF_NET_ERROR = -151
	ERR_ORIGIN_BOUND_CERT_GENERATION_TYPE_MISMATCH    CEF_NET_ERROR = -152
	ERR_SSL_DECRYPT_ERROR_ALERT                       CEF_NET_ERROR = -153
	ERR_WS_THROTTLE_QUEUE_TOO_LARGE                   CEF_NET_ERROR = -154
	ERR_SSL_SERVER_CERT_CHANGED                       CEF_NET_ERROR = -156
	ERR_SSL_UNRECOGNIZED_NAME_ALERT                   CEF_NET_ERROR = -159
	ERR_SOCKET_SET_RECEIVE_BUFFER_SIZE_ERROR          CEF_NET_ERROR = -160
	ERR_SOCKET_SET_SEND_BUFFER_SIZE_ERROR             CEF_NET_ERROR = -161
	ERR_SOCKET_RECEIVE_BUFFER_SIZE_UNCHANGEABLE       CEF_NET_ERROR = -162
	ERR_SOCKET_SEND_BUFFER_SIZE_UNCHANGEABLE          CEF_NET_ERROR = -163
	ERR_SSL_CLIENT_AUTH_CERT_BAD_FORMAT               CEF_NET_ERROR = -164
	ERR_ICANN_NAME_COLLISION                          CEF_NET_ERROR = -166
	ERR_SSL_SERVER_CERT_BAD_FORMAT                    CEF_NET_ERROR = -167
	ERR_CT_STH_PARSING_FAILED                         CEF_NET_ERROR = -168
	ERR_CT_STH_INCOMPLETE                             CEF_NET_ERROR = -169
	ERR_UNABLE_TO_REUSE_CONNECTION_FOR_PROXY_AUTH     CEF_NET_ERROR = -170
	ERR_CT_CONSISTENCY_PROOF_PARSING_FAILED           CEF_NET_ERROR = -171
	ERR_SSL_OBSOLETE_CIPHER                           CEF_NET_ERROR = -172
	ERR_WS_UPGRADE                                    CEF_NET_ERROR = -173
	ERR_READ_IF_READY_NOT_IMPLEMENTED                 CEF_NET_ERROR = -174
	ERR_SSL_VERSION_INTERFERENCE                      CEF_NET_ERROR = -175
	ERR_NO_BUFFER_SPACE                               CEF_NET_ERROR = -176
	ERR_SSL_CLIENT_AUTH_NO_COMMON_ALGORITHMS          CEF_NET_ERROR = -177
	ERR_EARLY_DATA_REJECTED                           CEF_NET_ERROR = -178
	ERR_WRONG_VERSION_ON_EARLY_DATA                   CEF_NET_ERROR = -179
	ERR_TLS13_DOWNGRADE_DETECTED                      CEF_NET_ERROR = -180
	ERR_SSL_KEY_USAGE_INCOMPATIBLE                    CEF_NET_ERROR = -181
	ERR_CERT_COMMON_NAME_INVALID                      CEF_NET_ERROR = -200
	ERR_CERT_DATE_INVALID                             CEF_NET_ERROR = -201
	ERR_CERT_AUTHORITY_INVALID                        CEF_NET_ERROR = -202
	ERR_CERT_CONTAINS_ERRORS                          CEF_NET_ERROR = -203
	ERR_CERT_NO_REVOCATION_MECHANISM                  CEF_NET_ERROR = -204
	ERR_CERT_UNABLE_TO_CHECK_REVOCATION               CEF_NET_ERROR = -205
	ERR_CERT_REVOKED                                  CEF_NET_ERROR = -206
	ERR_CERT_INVALID                                  CEF_NET_ERROR = -207
	ERR_CERT_WEAK_SIGNATURE_ALGORITHM                 CEF_NET_ERROR = -208
	ERR_CERT_NON_UNIQUE_NAME                          CEF_NET_ERROR = -210
	ERR_CERT_WEAK_KEY                                 CEF_NET_ERROR = -211
	ERR_CERT_NAME_CONSTRAINT_VIOLATION                CEF_NET_ERROR = -212
	ERR_CERT_VALIDITY_TOO_LONG                        CEF_NET_ERROR = -213
	ERR_CERTIFICATE_TRANSPARENCY_REQUIRED             CEF_NET_ERROR = -214
	ERR_CERT_SYMANTEC_LEGACY                          CEF_NET_ERROR = -215
	ERR_CERT_END                                      CEF_NET_ERROR = -216
	ERR_INVALID_URL                                   CEF_NET_ERROR = -300
	ERR_DISALLOWED_URL_SCHEME                         CEF_NET_ERROR = -301
	ERR_UNKNOWN_URL_SCHEME                            CEF_NET_ERROR = -302
	ERR_INVALID_REDIRECT                              CEF_NET_ERROR = -303
	ERR_TOO_MANY_REDIRECTS                            CEF_NET_ERROR = -310
	ERR_UNSAFE_REDIRECT                               CEF_NET_ERROR = -311
	ERR_UNSAFE_PORT                                   CEF_NET_ERROR = -312
	ERR_INVALID_RESPONSE                              CEF_NET_ERROR = -320
	ERR_INVALID_CHUNKED_ENCODING                      CEF_NET_ERROR = -321
	ERR_METHOD_NOT_SUPPORTED                          CEF_NET_ERROR = -322
	ERR_UNEXPECTED_PROXY_AUTH                         CEF_NET_ERROR = -323
	ERR_EMPTY_RESPONSE                                CEF_NET_ERROR = -324
	ERR_RESPONSE_HEADERS_TOO_BIG                      CEF_NET_ERROR = -325
	ERR_PAC_STATUS_NOT_OK                             CEF_NET_ERROR = -326
	ERR_PAC_SCRIPT_FAILED                             CEF_NET_ERROR = -327
	ERR_REQUEST_RANGE_NOT_SATISFIABLE                 CEF_NET_ERROR = -328
	ERR_MALFORMED_IDENTITY                            CEF_NET_ERROR = -329
	ERR_CONTENT_DECODING_FAILED                       CEF_NET_ERROR = -330
	ERR_NETWORK_IO_SUSPENDED                          CEF_NET_ERROR = -331
	ERR_SYN_REPLY_NOT_RECEIVED                        CEF_NET_ERROR = -332
	ERR_ENCODING_CONVERSION_FAILED                    CEF_NET_ERROR = -333
	ERR_UNRECOGNIZED_FTP_DIRECTORY_LISTING_FORMAT     CEF_NET_ERROR = -334
	ERR_NO_SUPPORTED_PROXIES                          CEF_NET_ERROR = -336
	ERR_SPDY_PROTOCOL_ERROR                           CEF_NET_ERROR = -337
	ERR_INVALID_AUTH_CREDENTIALS                      CEF_NET_ERROR = -338
	ERR_UNSUPPORTED_AUTH_SCHEME                       CEF_NET_ERROR = -339
	ERR_ENCODING_DETECTION_FAILED                     CEF_NET_ERROR = -340
	ERR_MISSING_AUTH_CREDENTIALS                      CEF_NET_ERROR = -341
	ERR_UNEXPECTED_SECURITY_LIBRARY_STATUS            CEF_NET_ERROR = -342
	ERR_MISCONFIGURED_AUTH_ENVIRONMENT                CEF_NET_ERROR = -343
	ERR_UNDOCUMENTED_SECURITY_LIBRARY_STATUS          CEF_NET_ERROR = -344
	ERR_RESPONSE_BODY_TOO_BIG_TO_DRAIN                CEF_NET_ERROR = -345
	ERR_RESPONSE_HEADERS_MULTIPLE_CONTENT_LENGTH      CEF_NET_ERROR = -346
	ERR_INCOMPLETE_SPDY_HEADERS                       CEF_NET_ERROR = -347
	ERR_PAC_NOT_IN_DHCP                               CEF_NET_ERROR = -348
	ERR_RESPONSE_HEADERS_MULTIPLE_CONTENT_DISPOSITION CEF_NET_ERROR = -349
	ERR_RESPONSE_HEADERS_MULTIPLE_LOCATION            CEF_NET_ERROR = -350
	ERR_SPDY_SERVER_REFUSED_STREAM                    CEF_NET_ERROR = -351
	ERR_SPDY_PING_FAILED                              CEF_NET_ERROR = -352
	ERR_CONTENT_LENGTH_MISMATCH                       CEF_NET_ERROR = -354
	ERR_INCOMPLETE_CHUNKED_ENCODING                   CEF_NET_ERROR = -355
	ERR_QUIC_PROTOCOL_ERROR                           CEF_NET_ERROR = -356
	ERR_RESPONSE_HEADERS_TRUNCATED                    CEF_NET_ERROR = -357
	ERR_QUIC_HANDSHAKE_FAILED                         CEF_NET_ERROR = -358
	ERR_SPDY_INADEQUATE_TRANSPORT_SECURITY            CEF_NET_ERROR = -360
	ERR_SPDY_FLOW_CONTROL_ERROR                       CEF_NET_ERROR = -361
	ERR_SPDY_FRAME_SIZE_ERROR                         CEF_NET_ERROR = -362
	ERR_SPDY_COMPRESSION_ERROR                        CEF_NET_ERROR = -363
	ERR_PROXY_AUTH_REQUESTED_WITH_NO_CONNECTION       CEF_NET_ERROR = -364
	ERR_HTTP_1_1_REQUIRED                             CEF_NET_ERROR = -365
	ERR_PROXY_HTTP_1_1_REQUIRED                       CEF_NET_ERROR = -366
	ERR_PAC_SCRIPT_TERMINATED                         CEF_NET_ERROR = -367
	ERR_INVALID_HTTP_RESPONSE                         CEF_NET_ERROR = -370
	ERR_CONTENT_DECODING_INIT_FAILED                  CEF_NET_ERROR = -371
	ERR_SPDY_RST_STREAM_NO_ERROR_RECEIVED             CEF_NET_ERROR = -372
	ERR_SPDY_PUSHED_STREAM_NOT_AVAILABLE              CEF_NET_ERROR = -373
	ERR_SPDY_CLAIMED_PUSHED_STREAM_RESET_BY_SERVER    CEF_NET_ERROR = -374
	ERR_TOO_MANY_RETRIES                              CEF_NET_ERROR = -375
	ERR_SPDY_STREAM_CLOSED                            CEF_NET_ERROR = -376
	ERR_SPDY_CLIENT_REFUSED_STREAM                    CEF_NET_ERROR = -377
	ERR_SPDY_PUSHED_RESPONSE_DOES_NOT_MATCH           CEF_NET_ERROR = -378
	ERR_CACHE_MISS                                    CEF_NET_ERROR = -400
	ERR_CACHE_READ_FAILURE                            CEF_NET_ERROR = -401
	ERR_CACHE_WRITE_FAILURE                           CEF_NET_ERROR = -402
	ERR_CACHE_OPERATION_NOT_SUPPORTED                 CEF_NET_ERROR = -403
	ERR_CACHE_OPEN_FAILURE                            CEF_NET_ERROR = -404
	ERR_CACHE_CREATE_FAILURE                          CEF_NET_ERROR = -405
	ERR_CACHE_RACE                                    CEF_NET_ERROR = -406
	ERR_CACHE_CHECKSUM_READ_FAILURE                   CEF_NET_ERROR = -407
	ERR_CACHE_CHECKSUM_MISMATCH                       CEF_NET_ERROR = -408
	ERR_CACHE_LOCK_TIMEOUT                            CEF_NET_ERROR = -409
	ERR_CACHE_AUTH_FAILURE_AFTER_READ                 CEF_NET_ERROR = -410
	ERR_CACHE_ENTRY_NOT_SUITABLE                      CEF_NET_ERROR = -411
	ERR_CACHE_DOOM_FAILURE                            CEF_NET_ERROR = -412
	ERR_CACHE_OPEN_OR_CREATE_FAILURE                  CEF_NET_ERROR = -413
	ERR_INSECURE_RESPONSE                             CEF_NET_ERROR = -501
	ERR_NO_PRIVATE_KEY_FOR_CERT                       CEF_NET_ERROR = -502
	ERR_ADD_USER_CERT_FAILED                          CEF_NET_ERROR = -503
	ERR_INVALID_SIGNED_EXCHANGE                       CEF_NET_ERROR = -504
	ERR_FTP_FAILED                                    CEF_NET_ERROR = -601
	ERR_FTP_SERVICE_UNAVAILABLE                       CEF_NET_ERROR = -602
	ERR_FTP_TRANSFER_ABORTED                          CEF_NET_ERROR = -603
	ERR_FTP_FILE_BUSY                                 CEF_NET_ERROR = -604
	ERR_FTP_SYNTAX_ERROR                              CEF_NET_ERROR = -605
	ERR_FTP_COMMAND_NOT_SUPPORTED                     CEF_NET_ERROR = -606
	ERR_FTP_BAD_COMMAND_SEQUENCE                      CEF_NET_ERROR = -607
	ERR_PKCS12_IMPORT_BAD_PASSWORD                    CEF_NET_ERROR = -701
	ERR_PKCS12_IMPORT_FAILED                          CEF_NET_ERROR = -702
	ERR_IMPORT_CA_CERT_NOT_CA                         CEF_NET_ERROR = -703
	ERR_IMPORT_CERT_ALREADY_EXISTS                    CEF_NET_ERROR = -704
	ERR_IMPORT_CA_CERT_FAILED                         CEF_NET_ERROR = -705
	ERR_IMPORT_SERVER_CERT_FAILED                     CEF_NET_ERROR = -706
	ERR_PKCS12_IMPORT_INVALID_MAC                     CEF_NET_ERROR = -707
	ERR_PKCS12_IMPORT_INVALID_FILE                    CEF_NET_ERROR = -708
	ERR_PKCS12_IMPORT_UNSUPPORTED                     CEF_NET_ERROR = -709
	ERR_KEY_GENERATION_FAILED                         CEF_NET_ERROR = -710
	ERR_PRIVATE_KEY_EXPORT_FAILED                     CEF_NET_ERROR = -712
	ERR_SELF_SIGNED_CERT_GENERATION_FAILED            CEF_NET_ERROR = -713
	ERR_CERT_DATABASE_CHANGED                         CEF_NET_ERROR = -714
	ERR_DNS_MALFORMED_RESPONSE                        CEF_NET_ERROR = -800
	ERR_DNS_SERVER_REQUIRES_TCP                       CEF_NET_ERROR = -801
	ERR_DNS_SERVER_FAILED                             CEF_NET_ERROR = -802
	ERR_DNS_TIMED_OUT                                 CEF_NET_ERROR = -803
	ERR_NS_CACHE_MISS                                 CEF_NET_ERROR = -804
	ERR_DNS_SEARCH_EMPTY                              CEF_NET_ERROR = -805
	ERR_DNS_SORT_ERROR                                CEF_NET_ERROR = -806
	ERR_DNS_HTTP_FAILED                               CEF_NET_ERROR = -807
)

// include/internal/cef_types.h (cef_color_type_t)
type TCefColorType = types.Int32

const (
	CEF_COLOR_TYPE_RGBA_8888 TCefColorType = iota
	CEF_COLOR_TYPE_BGRA_8888
)

// include/internal/cef_types.h (cef_alpha_type_t)
type TCefAlphaType = types.Int32

const (
	CEF_ALPHA_TYPE_OPAQUE TCefAlphaType = iota
	CEF_ALPHA_TYPE_PREMULTIPLIED
	CEF_ALPHA_TYPE_POSTMULTIPLIED
)

// Margin type for PDF printing.
// include/internal/cef_types.h (cef_pdf_print_margin_type_t)
type TCefPdfPrintMarginType = types.Int32

const (
	// Default margins of 1cm (~0.4 inches).
	PDF_PRINT_MARGIN_DEFAULT TCefPdfPrintMarginType = iota
	// No margins.
	PDF_PRINT_MARGIN_NONE
	// Custom margins using the |margin_*| values from TCefPdfPrintSettings.
	PDF_PRINT_MARGIN_CUSTOM
)

// V8ValueType ICefV8Value ValueType
type V8ValueType = types.Int32

const (
	V8vtInvalid V8ValueType = iota
	V8vtUndefined
	V8vtNull
	V8vtBool
	V8vtInt
	V8vtUInt
	V8vtDouble
	V8vtDate
	V8vtString
	V8vtObject
	V8vtArray
	V8vtArrayBuffer
	V8vtFunction
	V8vtPromise
)

// include/internal/cef_types.h (cef_preferences_type_t)
type TCefPreferencesType = types.Int32

const (
	CEF_PREFERENCES_TYPE_GLOBAL TCefPreferencesType = iota
	CEF_PREFERENCES_TYPE_REQUEST_CONTEXT
)

type TCefScaleFactor = types.Int32

// include/internal/cef_types.h (cef_scale_factor_t)
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

// include/internal/cef_types.h (cef_channel_layout_t)
type TCefChannelLayout = types.Int32

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
	CEF_CHANNEL_LAYOUT_5_1_4_DOWNMIX // CEF_CHANNEL_LAYOUT_MAX = CEF_CHANNEL_LAYOUT_5_1_4_DOWNMIX
)

// CefCreateType
//
//	CtSelfOwn 自己处理事件
//	CtOther   使用其他组件处理事件
type CefCreateType = types.Int8

const (
	CtSelfOwn CefCreateType = iota // CtSelfOwn 自己处理事件
	CtOther                        // CtOther   使用其他组件处理事件
)

// include/internal/cef_types.h (cef_quick_menu_edit_state_flags_t)
type TCefQuickMenuEditStateFlags = types.Int32

// include/internal/cef_types.h (cef_log_severity_t)
type TCefLogSeverity = types.Cardinal

// TCefCursorHandle
//
//	/include/internal/cef_types_win.h (cef_cursor_handle_t)
//	/include/internal/cef_types_mac.h (cef_cursor_handle_t)
//	/include/internal/cef_types_linux.h (cef_cursor_handle_t)
type TCefCursorHandle uintptr

// include/internal/cef_types.h (cef_cursor_type_t)
type TCefCursorType = types.Int32

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

// include/internal/cef_types.h (cef_focus_source_t)
type TCefFocusSource = types.Int32

const (
	FOCUS_SOURCE_NAVIGATION TCefFocusSource = iota
	FOCUS_SOURCE_SYSTEM
)

// include/internal/cef_types.h (cef_permission_request_result_t)
type TCefPermissionRequestResult = types.Int32

const (
	CEF_PERMISSION_RESULT_ACCEPT TCefPermissionRequestResult = iota
	CEF_PERMISSION_RESULT_DENY
	CEF_PERMISSION_RESULT_DISMISS
	CEF_PERMISSION_RESULT_IGNORE
)

// include/internal/cef_types.h (cef_media_access_permission_types_t)
type TCefMediaAccessPermissionTypes = types.Int32

// include/internal/cef_types.h (cef_jsdialog_type_t)
type TCefJsDialogType = types.Int32

const (
	JSDIALOGTYPE_ALERT TCefJsDialogType = iota
	JSDIALOGTYPE_CONFIRM
	JSDIALOGTYPE_PROMPT
)

// TCefDuplexMode
//
//	/include/internal/cef_types.h (cef_duplex_mode_t)
type TCefDuplexMode = types.Int32

// include/internal/cef_types.h (cef_color_model_t)
type TCefColorModel = types.Int32

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

// CefSchemeOption
//
//	Configuration options for registering a custom scheme.
//	These values are used when calling AddCustomScheme.
//
// include/internal/cef_types.h (cef_scheme_options_t)
type CefSchemeOption = types.Int32

const (
	CEF_SCHEME_OPTION_NONE CefSchemeOption = 0

	// CEF_SCHEME_OPTION_STANDARD
	//  If CEF_SCHEME_OPTION_STANDARD is set the scheme will be treated as a
	//  standard scheme. Standard schemes are subject to URL canonicalization and
	//  parsing rules as defined in the Common Internet Scheme Syntax RFC 1738
	//  Section 3.1 available at http://www.ietf.org/rfc/rfc1738.txt
	//
	//  In particular, the syntax for standard scheme URLs must be of the form:
	//  <pre>
	//   [scheme]://[username]:[password]@[host]:[port]/[url-path]
	//  </pre> Standard scheme URLs must have a host component that is a fully
	//  qualified domain name as defined in Section 3.5 of RFC 1034 [13] and
	//  Section 2.1 of RFC 1123. These URLs will be canonicalized to
	//  "scheme://host/path" in the simplest case and
	//  "scheme://username:password@host:port/path" in the most explicit case. For
	//  example, "scheme:host/path" and "scheme://host/path" will both be
	//  canonicalized to "scheme://host/path". The origin of a standard scheme URL
	//  is the combination of scheme, host and port (i.e., "scheme://host:port" in
	//  the most explicit case).
	//
	//  For non-standard scheme URLs only the "scheme:" component is parsed and
	//  canonicalized. The remainder of the URL will be passed to the handler as-
	//  is. For example, "scheme://some%20text" will remain the same.
	//  Non-standard scheme URLs cannot be used as a target for form submission.
	CEF_SCHEME_OPTION_STANDARD = 1 << 0

	// CEF_SCHEME_OPTION_LOCAL
	//  If CEF_SCHEME_OPTION_LOCAL is set the scheme will be treated with the same
	//  security rules as those applied to "file" URLs. Normal pages cannot link
	//  to or access local URLs. Also, by default, local URLs can only perform
	//  XMLHttpRequest calls to the same URL (origin + path) that originated the
	//  request. To allow XMLHttpRequest calls from a local URL to other URLs with
	//  the same origin set the CefSettings.file_access_from_file_urls_allowed
	//  value to true (1). To allow XMLHttpRequest calls from a local URL to all
	//  origins set the CefSettings.universal_access_from_file_urls_allowed value
	//  to true (1).
	CEF_SCHEME_OPTION_LOCAL = 1 << 1

	// CEF_SCHEME_OPTION_DISPLAY_ISOLATED
	//  If CEF_SCHEME_OPTION_DISPLAY_ISOLATED is set the scheme can only be
	//  displayed from other content hosted with the same scheme. For example,
	//  pages in other origins cannot create iframes or hyperlinks to URLs with
	//  the scheme. For schemes that must be accessible from other schemes don't
	//  set this, set CEF_SCHEME_OPTION_CORS_ENABLED, and use CORS
	//  "Access-Control-Allow-Origin" headers to further restrict access.
	CEF_SCHEME_OPTION_DISPLAY_ISOLATED = 1 << 2

	// CEF_SCHEME_OPTION_SECURE
	//  If CEF_SCHEME_OPTION_SECURE is set the scheme will be treated with the
	//  same security rules as those applied to "https" URLs. For example, loading
	//  this scheme from other secure schemes will not trigger mixed content
	//  warnings.
	CEF_SCHEME_OPTION_SECURE = 1 << 3

	// CEF_SCHEME_OPTION_CORS_ENABLED
	//  If CEF_SCHEME_OPTION_CORS_ENABLED is set the scheme can be sent CORS
	//  requests. This value should be set in most cases where
	//  CEF_SCHEME_OPTION_STANDARD is set.
	CEF_SCHEME_OPTION_CORS_ENABLED = 1 << 4

	// CEF_SCHEME_OPTION_CSP_BYPASSING
	//  If CEF_SCHEME_OPTION_CSP_BYPASSING is set the scheme can bypass Content-
	//  Security-Policy (CSP) checks. This value should not be set in most cases
	//  where CEF_SCHEME_OPTION_STANDARD is set.
	CEF_SCHEME_OPTION_CSP_BYPASSING = 1 << 5

	// CEF_SCHEME_OPTION_FETCH_ENABLED
	//  If CEF_SCHEME_OPTION_FETCH_ENABLED is set the scheme can perform Fetch API requests.
	CEF_SCHEME_OPTION_FETCH_ENABLED = 1 << 6
)

// TCefResponseFilterStatus
//
//	/include/internal/cef_types.h (cef_response_filter_status_t)
type TCefResponseFilterStatus = types.Int32

const (
	RESPONSE_FILTER_NEED_MORE_DATA TCefResponseFilterStatus = iota
	RESPONSE_FILTER_DONE
	RESPONSE_FILTER_ERROR
)

// include/internal/cef_types.h (cef_paint_element_type_t)
type TCefPaintElementType = types.Int32

const (
	PET_VIEW TCefPaintElementType = iota
	PET_POPUP
)

// include/internal/cef_types.h (cef_horizontal_alignment_t)
type TCefHorizontalAlignment = types.Int32

const (
	CEF_HORIZONTAL_ALIGNMENT_LEFT TCefHorizontalAlignment = iota
	CEF_HORIZONTAL_ALIGNMENT_CENTER
	CEF_HORIZONTAL_ALIGNMENT_RIGHT
)

// include/internal/cef_types.h (cef_text_input_mode_t)
type TCefTextInputMode = types.Int32

const (
	CEF_TEXT_INPUT_MODE_DEFAULT TCefTextInputMode = iota
	CEF_TEXT_INPUT_MODE_NONE
	CEF_TEXT_INPUT_MODE_TEXT
	CEF_TEXT_INPUT_MODE_TEL
	CEF_TEXT_INPUT_MODE_URL
	CEF_TEXT_INPUT_MODE_EMAIL
	CEF_TEXT_INPUT_MODE_NUMERIC
	CEF_TEXT_INPUT_MODE_DECIMAL
	CEF_TEXT_INPUT_MODE_SEARCH // CEF_TEXT_INPUT_MODE_MAX = CEF_TEXT_INPUT_MODE_SEARCH
)

// include/internal/cef_types.h (cef_cert_status_t)
type TCefCertStatus = types.Int32

// / Supported SSL version values.
// / <para>See the uCEFConstants unit for all possible values.</para>
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_version_t)</see></para>
// / <para><see href="https://source.chromium.org/chromium/chromium/src/+/main:net/ssl/ssl_connection_status_flags.h">See net/ssl/ssl_connection_status_flags.h for more information.</see></para>
type TCefSSLVersion = types.Int32

// / Supported SSL content status flags. See content/public/common/ssl_status.h for more information.
// / <para>See the uCEFConstants unit for all possible values.</para>
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_content_status_t)</see></para>
type TCefSSLContentStatus = types.UInt32

// include/internal/cef_types.h (cef_media_route_create_result_t)
type TCefMediaRouterCreateResult = types.Int32

// include/internal/cef_types.h (cef_media_route_connection_state_t)
type TCefMediaRouteConnectionState = types.Int32

const (
	CEF_MRCS_UNKNOWN TCefMediaRouteConnectionState = iota
	CEF_MRCS_CONNECTING
	CEF_MRCS_CONNECTED
	CEF_MRCS_CLOSED
	CEF_MRCS_TERMINATED
)

// include/internal/cef_types.h (cef_dom_document_type_t)
type TCefDomDocumentType = types.Int32

const (
	DOM_DOCUMENT_TYPE_UNKNOWN TCefDomDocumentType = iota
	DOM_DOCUMENT_TYPE_HTML
	DOM_DOCUMENT_TYPE_XHTML
	DOM_DOCUMENT_TYPE_PLUGIN
)

// include/internal/cef_types.h (cef_dom_node_type_t)
type TCefDomNodeType = types.Int32

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

type TCefClearDataStorageTypes = types.Int32

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
	CdstAll // default
)

// include/internal/cef_types.h (cef_composition_underline_style_t)
type TCefCompositionUnderlineStyle int32

const (
	CEF_CUS_SOLID TCefCompositionUnderlineStyle = iota
	CEF_CUS_DOT
	CEF_CUS_DASH
	CEF_CUS_NONE
)

// MessageBox() Flags
const (
	MB_OK               = 0x00000000
	MB_OKCANCEL         = 0x00000001
	MB_ABORTRETRYIGNORE = 0x00000002
	MB_YESNOCANCEL      = 0x00000003
	MB_YESNO            = 0x00000004
	MB_RETRYCANCEL      = 0x00000005

	MB_ICONHAND        = 0x00000010
	MB_ICONQUESTION    = 0x00000020
	MB_ICONEXCLAMATION = 0x00000030
	MB_ICONASTERISK    = 0x00000040
	MB_USERICON        = 0x00000080
	MB_ICONWARNING     = MB_ICONEXCLAMATION
	MB_ICONERROR       = MB_ICONHAND
	MB_ICONINFORMATION = MB_ICONASTERISK
	MB_ICONSTOP        = MB_ICONHAND
	MB_DEFBUTTON1      = 0x00000000
	MB_DEFBUTTON2      = 0x00000100
	MB_DEFBUTTON3      = 0x00000200
	MB_DEFBUTTON4      = 0x00000300
	MB_APPLMODAL       = 0x00000000
	MB_SYSTEMMODAL     = 0x00001000
	MB_TASKMODAL       = 0x00002000
	MB_HELP            = 0x00004000 // Help Button

	MB_NOFOCUS              = 0x00008000
	MB_SETFOREGROUND        = 0x00010000
	MB_DEFAULT_DESKTOP_ONLY = 0x00020000

	MB_TOPMOST    = 0x00040000
	MB_RIGHT      = 0x00080000
	MB_RTLREADING = 0x00100000

	MB_SERVICE_NOTIFICATION      = 0x00200000
	MB_SERVICE_NOTIFICATION_NT3X = 0x00040000

	MB_TYPEMASK = 0x0000000F
	MB_ICONMASK = 0x000000F0
	MB_DEFMASK  = 0x00000F00
	MB_MODEMASK = 0x00003000
	MB_MISCMASK = 0x0000C000
)

type TThreadPriority = types.Int32

const (
	TpIdle TThreadPriority = iota
	TpLowest
	TpLower
	TpNormal
	TpHigher
	TpHighest
	TpTimeCritical
)

// include/internal/cef_types.h (cef_file_dialog_mode_t)
type FileDialogMode = types.Int32

const (
	FILE_DIALOG_OPEN          FileDialogMode = 0x00000000
	FILE_DIALOG_OPEN_MULTIPLE                = 0x00000001
	FILE_DIALOG_OPEN_FOLDER                  = 0x00000002
	FILE_DIALOG_SAVE                         = 0x00000003
)

// include/internal/cef_types.h (cef_menu_item_type_t)
type TCefMenuItemType = types.Int32

const (
	MENUITEMTYPE_NONE TCefMenuItemType = iota
	MENUITEMTYPE_COMMAND
	MENUITEMTYPE_CHECK
	MENUITEMTYPE_RADIO
	MENUITEMTYPE_SEPARATOR
	MENUITEMTYPE_SUBMENU
)

// include/internal/cef_types.h (cef_button_state_t)
type TCefButtonState = types.Int32

const (
	CEF_BUTTON_STATE_NORMAL TCefButtonState = iota
	CEF_BUTTON_STATE_HOVERED
	CEF_BUTTON_STATE_PRESSED
	CEF_BUTTON_STATE_DISABLED
)

// include/internal/cef_types.h (cef_main_axis_alignment_t)
type TCefMainAxisAlignment = types.Int32

const (
	// Child views will be left/top-aligned.
	CEF_MAIN_AXIS_ALIGNMENT_START TCefMainAxisAlignment = iota
	// Child views will be center-aligned.
	CEF_AXIS_ALIGNMENT_CENTER
	// Child views will be right/bottom-aligned.
	CEF_AXIS_ALIGNMENT_END
	// Child views will be stretched to fit.
	CEF_AXIS_ALIGNMENT_STRETCH
)

// Text style types. Should be kepy in sync with gfx::TextStyle.
//
// include/internal/cef_types.h (cef_text_style_t)
type TCefTextStyle = types.Int32

const (
	CEF_TEXT_STYLE_BOLD TCefTextStyle = iota
	CEF_TEXT_STYLE_ITALIC
	CEF_TEXT_STYLE_STRIKE
	CEF_TEXT_STYLE_DIAGONAL_STRIKE
	CEF_TEXT_STYLE_UNDERLINE
)

// include/internal/cef_types.h (cef_text_field_commands_t)
type TCefTextFieldCommands = types.Int32

const (
	CEF_TFC_CUT TCefTextFieldCommands = iota + 1
	CEF_TFC_COPY
	CEF_TFC_PASTE
	CEF_TFC_UNDO
	CEF_TFC_DELETE
	CEF_TFC_SELECT_ALL
)

// UI tool classes used
type UITool int8

const (
	UitInvalid UITool = iota - 1 // invalid
	UitWin32                     // windows
	UitGtk2                      // linux
	UitGtk3                      // linux
	UitCocoa                     // macos
)

// LocalCustomerScheme 本地资源加载自定义固定协议
//
//	file, fs
type LocalCustomerScheme string

const (
	LcsLocal LocalCustomerScheme = "local" // 本地目录 local://energy/index.html
	LcsFS    LocalCustomerScheme = "fs"    // 内置 fs://energy/index.html
)

// LocalProxyScheme
//
//	本地加载资源，在浏览器发起xhr请求时的代理协议
//	http, https
type LocalProxyScheme int

const (
	LpsHttp  LocalProxyScheme = iota // http
	LpsHttps                         // https
	//LpsTcp                           // tcp
)

type TCefPermissionRequestTypes int32

// Permission types used with OnShowPermissionPrompt. Some types are
// platform-specific or only supported with the Chrome runtime. Should be kept
// in sync with Chromium's permissions::RequestType type.
// <para>TCefPermissionRequestTypes values.</para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_permission_request_types_t)</see></para>
const (
	CEF_PERMISSION_TYPE_NONE                       TCefPermissionRequestTypes = 0
	CEF_PERMISSION_TYPE_ACCESSIBILITY_EVENTS       TCefPermissionRequestTypes = 1 << 0
	CEF_PERMISSION_TYPE_AR_SESSION                 TCefPermissionRequestTypes = 1 << 1
	CEF_PERMISSION_TYPE_CAMERA_PAN_TILT_ZOOM       TCefPermissionRequestTypes = 1 << 2
	CEF_PERMISSION_TYPE_CAMERA_STREAM              TCefPermissionRequestTypes = 1 << 3
	CEF_PERMISSION_TYPE_CLIPBOARD                  TCefPermissionRequestTypes = 1 << 4
	CEF_PERMISSION_TYPE_TOP_LEVEL_STORAGE_ACCESS   TCefPermissionRequestTypes = 1 << 5
	CEF_PERMISSION_TYPE_DISK_QUOTA                 TCefPermissionRequestTypes = 1 << 6
	CEF_PERMISSION_TYPE_LOCAL_FONTS                TCefPermissionRequestTypes = 1 << 7
	CEF_PERMISSION_TYPE_GEOLOCATION                TCefPermissionRequestTypes = 1 << 8
	CEF_PERMISSION_TYPE_IDLE_DETECTION             TCefPermissionRequestTypes = 1 << 9
	CEF_PERMISSION_TYPE_MIC_STREAM                 TCefPermissionRequestTypes = 1 << 10
	CEF_PERMISSION_TYPE_MIDI                       TCefPermissionRequestTypes = 1 << 11
	CEF_PERMISSION_TYPE_MIDI_SYSEX                 TCefPermissionRequestTypes = 1 << 12
	CEF_PERMISSION_TYPE_MULTIPLE_DOWNLOADS         TCefPermissionRequestTypes = 1 << 13
	CEF_PERMISSION_TYPE_NOTIFICATIONS              TCefPermissionRequestTypes = 1 << 14
	CEF_PERMISSION_TYPE_PROTECTED_MEDIA_IDENTIFIER TCefPermissionRequestTypes = 1 << 15
	CEF_PERMISSION_TYPE_REGISTER_PROTOCOL_HANDLER  TCefPermissionRequestTypes = 1 << 16
	CEF_PERMISSION_TYPE_STORAGE_ACCESS             TCefPermissionRequestTypes = 1 << 17
	CEF_PERMISSION_TYPE_VR_SESSION                 TCefPermissionRequestTypes = 1 << 18
	CEF_PERMISSION_TYPE_WINDOW_MANAGEMENT          TCefPermissionRequestTypes = 1 << 19
)

type TCefChromePageActionIconType int32

// Chrome page action icon types. Should be kept in sync with Chromium's
// PageActionIconType type.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_page_action_icon_type_t)</see></para>
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
	//{* CEF_CPAIT_MAX_VALUE = CEF_CPAIT_PRICE_INSIGHTS *}
)

type TCefChromeToolbarButtonType int32

// Chrome toolbar button types. Should be kept in sync with CEF's internal
// ToolbarButtonType type.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_chrome_toolbar_button_type_t)</see></para>
const (
	CEF_CTBT_CAST TCefChromeToolbarButtonType = iota
	CEF_CTBT_DOWNLOAD
	CEF_CTBT_SEND_TAB_TO_SELF
	CEF_CTBT_SIDE_PANEL
	//{* CEF_CTBT_MAX_VALUE = CEF_CTBT_SIDE_PANEL *}
)

type TCefGestureCommand int32

// Specifies the gesture commands.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_gesture_command_t)</see></para>
const (
	CEF_GESTURE_COMMAND_BACK TCefGestureCommand = iota
	CEF_GESTURE_COMMAND_FORWARD
)

// Download interrupt reasons. Should be kept in sync with
// Chromium's download::DownloadInterruptReason type.
// <para>See the uCEFConstants unit for all possible values.</para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see></para>
type TCefDownloadInterruptReason int32

type TCefUIColorMode int32

// Color mode in UI for platforms that support it.
const (

	// System default.

	UICMSystemDefault TCefUIColorMode = iota

	// Forces light color mode in UI for platforms that support it.

	UICMForceLight

	// Forces dark color mode in UI for platforms that support it.

	UICMForceDark
)

type TCefContentSettingTypes int32

// Supported content setting types. Some types are platform-specific or only
// supported with the Chrome runtime. Should be kept in sync with Chromium's
// ContentSettingsType type.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_content_setting_types_t)</see></para>
const (
	CEF_CONTENT_SETTING_TYPE_COOKIES TCefContentSettingTypes = iota
	CEF_CONTENT_SETTING_TYPE_IMAGES
	CEF_CONTENT_SETTING_TYPE_JAVASCRIPT

	// This setting governs both popups and unwanted redirects like tab-unders and framebusting.

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
	// Website setting which stores metadata for the subresource filter to aid in
	// decisions for whether or not to show the UI.
	CEF_CONTENT_SETTING_TYPE_ADS_DATA
	// This is special-cased in the permissions layer to always allow, and as
	// such doesn't have associated prefs data.
	CEF_CONTENT_SETTING_TYPE_MIDI
	// This content setting type is for caching password protection service's
	// verdicts of each origin.
	CEF_CONTENT_SETTING_TYPE_PASSWORD_PROTECTION
	// Website setting which stores engagement data for media related to a
	// specific origin.
	CEF_CONTENT_SETTING_TYPE_MEDIA_ENGAGEMENT
	// Content setting which stores whether or not the site can play audible
	// sound. This will not block playback but instead the user will not hear it.
	CEF_CONTENT_SETTING_TYPE_SOUND
	// Website setting which stores the list of client hints that the origin
	// requested the browser to remember. The browser is expected to send all
	// client hints in the HTTP request headers for every resource requested
	// from that origin.
	CEF_CONTENT_SETTING_TYPE_CLIENT_HINTS
	// Generic Sensor API covering ambient-light-sensor, accelerometer, gyroscope
	// and magnetometer are all mapped to a single content_settings_type.
	// Setting for the Generic Sensor API covering ambient-light-sensor,
	// accelerometer, gyroscope and magnetometer. These are all mapped to a
	// single ContentSettingsType.
	CEF_CONTENT_SETTING_TYPE_SENSORS

	// Content setting which stores whether or not the user has granted the site
	// permission to respond to accessibility events, which can be used to
	// provide a custom accessibility experience. Requires explicit user consent
	// because some users may not want sites to know they're using assistive
	// technology.

	CEF_CONTENT_SETTING_TYPE_ACCESSIBILITY_EVENTS

	// Used to store whether to allow a website to install a payment handler.

	CEF_CONTENT_SETTING_TYPE_PAYMENT_HANDLER

	// Content setting which stores whether to allow sites to ask for permission
	// to access USB devices. If this is allowed specific device permissions are
	// stored under USB_CHOOSER_DATA.

	CEF_CONTENT_SETTING_TYPE_USB_GUARD

	// Nothing is stored in this setting at present. Please refer to
	// BackgroundFetchPermissionContext for details on how this permission
	// is ascertained.

	CEF_CONTENT_SETTING_TYPE_BACKGROUND_FETCH

	// Website setting which stores the amount of times the user has dismissed
	// intent picker UI without explicitly choosing an option.

	CEF_CONTENT_SETTING_TYPE_INTENT_PICKER_DISPLAY

	// Used to store whether to allow a website to detect user active/idle state.

	CEF_CONTENT_SETTING_TYPE_IDLE_DETECTION

	// Setting for enabling auto-select of all screens for getDisplayMediaSet.

	CEF_CONTENT_SETTING_TYPE_GET_DISPLAY_MEDIA_SET_SELECT_ALL_SCREENS

	// Content settings for access to serial ports. The "guard" content setting
	// stores whether to allow sites to ask for permission to access a port. The
	// permissions granted to access particular ports are stored in the "chooser
	// data" website setting.

	CEF_CONTENT_SETTING_TYPE_SERIAL_GUARD
	CEF_CONTENT_SETTING_TYPE_SERIAL_CHOOSER_DATA

	// Nothing is stored in this setting at present. Please refer to
	// PeriodicBackgroundSyncPermissionContext for details on how this permission
	// is ascertained.
	// This content setting is not registered because it does not require access
	// to any existing providers.

	CEF_CONTENT_SETTING_TYPE_PERIODIC_BACKGROUND_SYNC

	// Content setting which stores whether to allow sites to ask for permission
	// to do Bluetooth scanning.

	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_SCANNING

	// Content settings for access to HID devices. The "guard" content setting
	// stores whether to allow sites to ask for permission to access a device.
	// The permissions granted to access particular devices are stored in the
	// "chooser data" website setting.

	CEF_CONTENT_SETTING_TYPE_HID_GUARD
	CEF_CONTENT_SETTING_TYPE_HID_CHOOSER_DATA

	// Wake Lock API, which has two lock types: screen and system locks.
	// Currently, screen locks do not need any additional permission, and system
	// locks are always denied while the right UI is worked out.

	CEF_CONTENT_SETTING_TYPE_WAKE_LOCK_SCREEN
	CEF_CONTENT_SETTING_TYPE_WAKE_LOCK_SYSTEM

	// Legacy SameSite cookie behavior. This disables SameSite=Lax-by-default,
	// SameSite=None requires Secure, and Schemeful Same-Site, forcing the
	// legacy behavior wherein 1) cookies that don't specify SameSite are treated
	// as SameSite=None, 2) SameSite=None cookies are not required to be Secure,
	// and 3) schemeful same-site is not active.
	//
	// This will also be used to revert to legacy behavior when future changes
	// in cookie handling are introduced.

	CEF_CONTENT_SETTING_TYPE_LEGACY_COOKIE_ACCESS

	// Content settings which stores whether to allow sites to ask for permission
	// to save changes to an original file selected by the user through the
	// File System Access API.

	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_WRITE_GUARD

	// Used to store whether to allow a website to exchange data with NFC
	// devices.

	CEF_CONTENT_SETTING_TYPE_NFC

	// Website setting to store permissions granted to access particular
	// Bluetooth devices.

	CEF_CONTENT_SETTING_TYPE_BLUETOOTH_CHOOSER_DATA

	// Full access to the system clipboard (sanitized read without user gesture,
	// and unsanitized read and write with user gesture).

	CEF_CONTENT_SETTING_TYPE_CLIPBOARD_READ_WRITE

	// This is special-cased in the permissions layer to always allow, and as
	// such doesn't have associated prefs data.

	CEF_CONTENT_SETTING_TYPE_CLIPBOARD_SANITIZED_WRITE

	// This content setting type is for caching safe browsing real time url
	// check's verdicts of each origin.

	CEF_CONTENT_SETTING_TYPE_SAFE_BROWSING_URL_CHECK_DATA

	// Used to store whether a site is allowed to request AR or VR sessions with
	// the WebXr Device API.

	CEF_CONTENT_SETTING_TYPE_VR
	CEF_CONTENT_SETTING_TYPE_AR

	// Content setting which stores whether to allow site to open and read files
	// and directories selected through the File System Access API.

	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_READ_GUARD

	// Access to first party storage in a third-party context. Exceptions are
	// scoped to the combination of requesting/top-level origin, and are managed
	// through the Storage Access API. For the time being, this content setting
	// exists in parallel to third-party cookie rules stored in COOKIES.

	CEF_CONTENT_SETTING_TYPE_STORAGE_ACCESS

	// Content setting which stores whether to allow a site to control camera
	// movements. It does not give access to camera.

	CEF_CONTENT_SETTING_TYPE_CAMERA_PAN_TILT_ZOOM

	// Content setting for Screen Enumeration and Screen Detail functionality.
	// Permits access to detailed multi-screen information, like size and
	// position. Permits placing fullscreen and windowed content on specific
	// screens. See also: https://w3c.github.io/window-placement

	CEF_CONTENT_SETTING_TYPE_WINDOW_MANAGEMENT

	// Stores whether to allow insecure websites to make local network requests.
	// See also: https://wicg.github.io/local-network-access
	// Set through enterprise policies only.

	CEF_CONTENT_SETTING_TYPE_INSECURE_LOCAL_NETWORK

	// Content setting which stores whether or not a site can access low-level
	// locally installed font data using the Local Fonts Access API.

	CEF_CONTENT_SETTING_TYPE_LOCAL_FONTS

	// Stores per-origin state for permission auto-revocation (for all permission
	// types).

	CEF_CONTENT_SETTING_TYPE_PERMISSION_AUTOREVOCATION_DATA

	// Stores per-origin state of the most recently selected directory for the
	// use by the File System Access API.

	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_LAST_PICKED_DIRECTORY

	// Controls access to the getDisplayMedia API when {preferCurrentTab: true}
	// is specified.

	CEF_CONTENT_SETTING_TYPE_DISPLAY_CAPTURE

	// Website setting to store permissions metadata granted to paths on the
	// local file system via the File System Access API.
	// |FILE_SYSTEM_WRITE_GUARD| is the corresponding "guard" setting.

	CEF_CONTENT_SETTING_TYPE_FILE_SYSTEM_ACCESS_CHOOSER_DATA

	// Stores a grant that allows a relying party to send a request for identity
	// information to specified identity providers, potentially through any
	// anti-tracking measures that would otherwise prevent it. This setting is
	// associated with the relying party's origin.

	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_SHARING

	// Whether to use the v8 optimized JIT for running JavaScript on the page.

	CEF_CONTENT_SETTING_TYPE_JAVASCRIPT_JIT

	// Content setting which stores user decisions to allow loading a site over
	// HTTP. Entries are added by hostname when a user bypasses the HTTPS-First
	// Mode interstitial warning when a site does not support HTTPS. Allowed
	// hosts are exact hostname matches -- subdomains of a host on the allowlist
	// must be separately allowlisted.

	CEF_CONTENT_SETTING_TYPE_HTTP_ALLOWED

	// Stores metadata related to form fill, such as e.g. whether user data was
	// autofilled on a specific website.

	CEF_CONTENT_SETTING_TYPE_FORMFILL_METADATA

	// Setting to indicate that there is an active federated sign-in session
	// between a specified relying party and a specified identity provider for
	// a specified account. When this is present it allows access to session
	// management capabilities between the sites. This setting is associated
	// with the relying party's origin.

	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_ACTIVE_SESSION

	// Setting to indicate whether Chrome should automatically apply darkening to
	// web content.

	CEF_CONTENT_SETTING_TYPE_AUTO_DARK_WEB_CONTENT

	// Setting to indicate whether Chrome should request the desktop view of a
	// site instead of the mobile one.

	CEF_CONTENT_SETTING_TYPE_REQUEST_DESKTOP_SITE

	// Setting to indicate whether browser should allow signing into a website
	// via the browser FedCM API.

	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_API

	// Stores notification interactions per origin for the past 90 days.
	// Interactions per origin are pre-aggregated over seven-day windows: A
	// notification interaction or display is assigned to the last Monday
	// midnight in local time.

	CEF_CONTENT_SETTING_TYPE_NOTIFICATION_INTERACTIONS

	// Website setting which stores the last reduced accept language negotiated
	// for a given origin, to be used on future visits to the origin.

	CEF_CONTENT_SETTING_TYPE_REDUCED_ACCEPT_LANGUAGE

	// Website setting which is used for NotificationPermissionReviewService to
	// store origin blocklist from review notification permissions feature.

	CEF_CONTENT_SETTING_TYPE_NOTIFICATION_PERMISSION_REVIEW

	// Website setting to store permissions granted to access particular devices
	// in private network.

	CEF_CONTENT_SETTING_TYPE_PRIVATE_NETWORK_GUARD
	CEF_CONTENT_SETTING_TYPE_PRIVATE_NETWORK_CHOOSER_DATA

	// Website setting which stores whether the browser has observed the user
	// signing into an identity-provider based on observing the IdP-SignIn-Status
	// HTTP header.

	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_IDENTITY_PROVIDER_SIGNIN_STATUS

	// Website setting which is used for UnusedSitePermissionsService to
	// store revoked permissions of unused sites from unused site permissions
	// feature.

	CEF_CONTENT_SETTING_TYPE_REVOKED_UNUSED_SITE_PERMISSIONS

	// Similar to STORAGE_ACCESS, but applicable at the page-level rather than
	// being specific to a frame.

	CEF_CONTENT_SETTING_TYPE_TOP_LEVEL_STORAGE_ACCESS

	// Setting to indicate whether user has opted in to allowing auto re-authn
	// via the FedCM API.

	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_AUTO_REAUTHN_PERMISSION

	// Website setting which stores whether the user has explicitly registered
	// a website as an identity-provider.

	CEF_CONTENT_SETTING_TYPE_FEDERATED_IDENTITY_IDENTITY_PROVIDER_REGISTRATION

	// Content setting which is used to indicate whether anti-abuse functionality
	// should be enabled.

	CEF_CONTENT_SETTING_TYPE_ANTI_ABUSE

	// Content setting used to indicate whether third-party storage partitioning
	// should be enabled.

	CEF_CONTENT_SETTING_TYPE_THIRD_PARTY_STORAGE_PARTITIONING

	// Used to indicate whether HTTPS-First Mode is enabled on the hostname.

	CEF_CONTENT_SETTING_TYPE_HTTPS_ENFORCED

	// Stores per origin metadata for cookie controls.

	CEF_CONTENT_SETTING_TYPE_COOKIE_CONTROLS_METADATA

	// Setting for supporting 3PCD.

	CEF_CONTENT_SETTING_TYPE_TPCD_SUPPORT
	CEF_CONTENT_SETTING_TYPE_NUM_TYPES
)

type TCefContentSettingValues int32

// Supported content setting values. Should be kept in sync with Chromium's
// ContentSetting type.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_content_setting_values_t)</see></para>
const (
	CEF_CONTENT_SETTING_VALUE_DEFAULT TCefContentSettingValues = iota
	CEF_CONTENT_SETTING_VALUE_ALLOW
	CEF_CONTENT_SETTING_VALUE_BLOCK
	CEF_CONTENT_SETTING_VALUE_ASK
	CEF_CONTENT_SETTING_VALUE_SESSION_ONLY
	CEF_CONTENT_SETTING_VALUE_DETECT_IMPORTANT_CONTENT
	CEF_CONTENT_SETTING_VALUE_NUM_VALUES
)

// Log items prepended to each log line.
// <para>See the uCEFConstants unit for all possible values.</para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see></para>
type TCefLogItems = types.Cardinal

const (
	// Prepend the default list of items.
	LOG_ITEMS_DEFAULT TCefLogItems = 0
	// Prepend no items.
	LOG_ITEMS_NONE TCefLogItems = 1
	// Prepend the process ID.
	LOG_ITEMS_FLAG_PROCESS_ID TCefLogItems = 1 << 1
	// Prepend the thread ID.
	LOG_ITEMS_FLAG_THREAD_ID TCefLogItems = 1 << 2
	// Prepend the timestamp.
	LOG_ITEMS_FLAG_TIME_STAMP TCefLogItems = 1 << 3
	// Prepend the tickcount.
	LOG_ITEMS_FLAG_TICK_COUNT TCefLogItems = 1 << 4
)

type ZoomCommand = uintptr

const (
	ZcInc   ZoomCommand = iota + 1 // IncZoomCommand
	ZcDec                          // DecZoomCommand
	ZcReset                        // ResetZoomCommand
)

type CanZoom = uintptr

const (
	CzInc   CanZoom = iota + 1 // CanIncZoom
	CzDec                      // CanDecZoom
	CzReset                    // CanResetZoom
)

// TCefThreadId
//
//	Existing thread IDs.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_thread_id_t))</a>
type TCefThreadId = int32

const (
	// The main thread in the browser. This will be the same as the main
	// application thread if CefInitialize() is called with a
	// TCefSettings.multi_threaded_message_loop value of false. Do not perform
	// blocking tasks on this thread. All tasks posted after
	// ICefBrowserProcessHandler.OnContextInitialized() and before CefShutdown()
	// are guaranteed to run. This thread will outlive all other CEF threads.
	TID_UI TCefThreadId = iota
	// Used for blocking tasks like file system access where the user won't
	// notice if the task takes an arbitrarily long time to complete. All tasks
	// posted after ICefBrowserProcessHandler.OnContextInitialized() and before
	// CefShutdown() are guaranteed to run.
	TID_FILE_BACKGROUND
	// Used for blocking tasks like file system access that affect UI or
	// responsiveness of future user interactions. Do not use if an immediate
	// response to a user interaction is expected. All tasks posted after
	// ICefBrowserProcessHandler.OnContextInitialized() and before CefShutdown()
	// are guaranteed to run.
	// Examples:
	// - Updating the UI to reflect progress on a long task.
	// - Loading data that might be shown in the UI after a future user
	//   interaction.
	TID_FILE_USER_VISIBLE
	// Used for blocking tasks like file system access that affect UI
	// immediately after a user interaction. All tasks posted after
	// ICefBrowserProcessHandler.OnContextInitialized() and before CefShutdown()
	// are guaranteed to run.
	// Example: Generating data shown in the UI immediately after a click.
	TID_FILE_USER_BLOCKING
	// Used to launch and terminate browser processes.
	TID_PROCESS_LAUNCHER
	// Used to process IPC and network messages. Do not perform blocking tasks on
	// this thread. All tasks posted after
	// ICefBrowserProcessHandler.OnContextInitialized() and before CefShutdown()
	// are guaranteed to run.
	TID_IO
	// The main thread in the renderer. Used for all WebKit and V8 interaction.
	// Tasks may be posted to this thread after
	// ICefRenderProcessHandler.OnWebKitInitialized but are not guaranteed to
	// run before sub-process termination (sub-processes may be killed at any
	// time without warning).
	TID_RENDERER
)

// DOM form control types. Should be kept in sync with Chromium's
// blink::mojom::FormControlType type.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_form_control_type_t)</see></para>
type TCefDomFormControlType = int32

const (
	DOM_FORM_CONTROL_TYPE_UNSUPPORTED TCefDomFormControlType = iota
	DOM_FORM_CONTROL_TYPE_BUTTON_BUTTON
	DOM_FORM_CONTROL_TYPE_BUTTON_SUBMIT
	DOM_FORM_CONTROL_TYPE_BUTTON_RESET
	DOM_FORM_CONTROL_TYPE_BUTTON_SELECT_LIST
	DOM_FORM_CONTROL_TYPE_BUTTON_POPOVER
	DOM_FORM_CONTROL_TYPE_FIELDSET
	DOM_FORM_CONTROL_TYPE_INPUT_BUTTON
	DOM_FORM_CONTROL_TYPE_INPUT_CHECKBOX
	DOM_FORM_CONTROL_TYPE_INPUT_COLOR
	DOM_FORM_CONTROL_TYPE_INPUT_DATE
	DOM_FORM_CONTROL_TYPE_INPUT_DATETIME_LOCAL
	DOM_FORM_CONTROL_TYPE_INPUT_EMAIL
	DOM_FORM_CONTROL_TYPE_INPUT_FILE
	DOM_FORM_CONTROL_TYPE_INPUT_HIDDEN
	DOM_FORM_CONTROL_TYPE_INPUT_IMAGE
	DOM_FORM_CONTROL_TYPE_INPUT_MONTH
	DOM_FORM_CONTROL_TYPE_INPUT_NUMBER
	DOM_FORM_CONTROL_TYPE_INPUT_PASSWORD
	DOM_FORM_CONTROL_TYPE_INPUT_RADIO
	DOM_FORM_CONTROL_TYPE_INPUT_RANGE
	DOM_FORM_CONTROL_TYPE_INPUT_RESET
	DOM_FORM_CONTROL_TYPE_INPUT_SEARCH
	DOM_FORM_CONTROL_TYPE_INPUT_SUBMIT
	DOM_FORM_CONTROL_TYPE_INPUT_TELEPHONE
	DOM_FORM_CONTROL_TYPE_INPUT_TEXT
	DOM_FORM_CONTROL_TYPE_INPUT_TIME
	DOM_FORM_CONTROL_TYPE_INPUT_URL
	DOM_FORM_CONTROL_TYPE_INPUT_WEEK
	DOM_FORM_CONTROL_TYPE_OUTPUT
	DOM_FORM_CONTROL_TYPE_SELECT_ONE
	DOM_FORM_CONTROL_TYPE_SELECT_MULTIPLE
	DOM_FORM_CONTROL_TYPE_SELECT_LIST
	DOM_FORM_CONTROL_TYPE_TEXT_AREA
)

// CEF supports both a Chrome runtime style (based on the Chrome UI layer) and
// an Alloy runtime style (based on the Chromium content layer). Chrome style
// provides the full Chrome UI and browser functionality whereas Alloy style
// provides less default browser functionality but adds additional client
// callbacks and support for windowless (off-screen) rendering. The style type
// is individually configured for each window/browser at creation time and
// different styles can be mixed during runtime. For additional comparative
// details on runtime styles see
// https://bitbucket.org/chromiumembedded/cef/wiki/Architecture.md#markdown-header-cef3</para>
//
// <para>Windowless rendering will always use Alloy style. Windowed rendering with a
// default window or client-provided parent window can configure the style via
// TCefWindowInfo.runtime_style. Windowed rendering with the Views framework can
// configure the style via ICefWindowDelegate.GetWindowRuntimeStyle and
// ICefBrowserViewDelegate.GetBrowserRuntimeStyle. Alloy style Windows with the
// Views framework can host only Alloy style BrowserViews but Chrome style
// Windows can host both style BrowserViews. Additionally, a Chrome style
// Window can host at most one Chrome style BrowserView but potentially
// multiple Alloy style BrowserViews. See TCefWindowInfo.runtime_style
// documentation for any additional platform-specific limitations.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_runtime.h">CEF source file: /include/internal/cef_types_runtime.h (cef_runtime_style_t)</see></para>
type TCefRuntimeStyle = int32

const (
	// Use the default style. See above documentation for exceptions.
	CEF_RUNTIME_STYLE_DEFAULT TCefRuntimeStyle = iota
	// Use Chrome style.
	CEF_RUNTIME_STYLE_CHROME
	// Use Alloy style.
	CEF_RUNTIME_STYLE_ALLOY
)

// / Icon types for a MediaSink object. Should be kept in sync with Chromium's
// / media_router::SinkIconType type.
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_sink_icon_type_t)</see></para>
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
	/// The total number of values.
	CEF_MSIT_TOTAL_COUNT
)
