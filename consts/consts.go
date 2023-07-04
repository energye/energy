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
	ExePath       string                       //执行文件目录
	HomeDir, _    = homedir.Dir()              //系统用户目录
	Separator     = string(filepath.Separator) //系统目录分隔符
	IsMessageLoop = false                      //CEF应用的窗口, true: 使用VF(views framework)窗口组件, false: 使用LCL窗口组件, 其实是窗口消息轮询使用方式.
)

const (
	Empty           = ""
	ENERGY_HOME_KEY = "ENERGY_HOME"
	MemoryNetwork   = "unix"
)

func init() {
	ExePath, _ = os.Getwd()
}

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

type TriggerMode = types.Int8

const (
	Tm_Async    = TriggerMode(iota) //异步
	Tm_Callback                     //异步，带回调函数返回结果
	Tm_Sync                         //同步，阻塞等待结果返回值
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
//  进程结束的状态
//  /include/internal/cef_types.h (cef_termination_status_t)
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

// CEF 进程 ChannelId
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

// 函数类型
type FN_TYPE = types.Int8

const (
	FN_TYPE_COMMON = FN_TYPE(iota) //普通函数，直接定义的
	FN_TYPE_OBJECT                 //对象函数，所属对象
)

// 通用类型或对象类型
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

// /include/internal/cef_types.h (cef_cookie_same_site_t)
type TCefCookieSameSite = types.Int32

const (
	Ccss_CEF_COOKIE_SAME_SITE_UNSPECIFIED = TCefCookieSameSite(iota)
	Ccss_CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	Ccss_CEF_COOKIE_SAME_SITE_LAX_MODE
	Ccss_CEF_COOKIE_SAME_SITE_STRICT_MODE
)

// /include/internal/cef_types.h (cef_cookie_priority_t)
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

// /include/internal/cef_types.h (cef_context_menu_media_type_t)
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

type MenuId = types.Int32

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

// /include/internal/cef_types.h (cef_menu_color_type_t)
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

// /include/internal/cef_types.h (cef_key_event_type_t)
type TCefKeyEventType = types.Int32

const (
	KEYEVENT_RAW_KEYDOWN = TCefKeyEventType(iota)
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

// /include/internal/cef_types.h (cef_event_flags_t)
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

// /include/internal/cef_types_win.h (cef_window_handle_t)
// /include/internal/cef_types_mac.h (cef_window_handle_t)
// /include/internal/cef_types_linux.h (cef_window_handle_t)
type TCefWindowHandle = types.UIntptr

const (
	Wht_WindowParent = TCefWindowHandleType(iota)
	Wht_LinkedWindowParent
)

// /include/internal/cef_types.h (cef_return_value_t)
type TCefReturnValue int32

const (
	RV_CANCEL = TCefReturnValue(iota)
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

// /include/internal/cef_types.h (cef_referrer_policy_t)
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

// /include/internal/cef_types.h (cef_urlrequest_flags_t)
type TCefUrlRequestFlags = types.Int

// /include/internal/cef_types.h (cef_errorcode_t)
type TCefErrorCode = types.Int32

// /include/internal/cef_types.h (cef_resource_type_t)
type TCefResourceType = types.Int32

const (
	RT_MAIN_FRAME = TCefResourceType(iota)
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
	RT_EMPTY_FILLER_TYPE_DO_NOT_USE  // This type doesn't exist in CEF and it's here just to fill this position.
	RT_NAVIGATION_PRELOAD_MAIN_FRAME // This type must have a value of 19
	RT_NAVIGATION_PRELOAD_SUB_FRAME
)

// /include/internal/cef_types.h (cef_transition_type_t)
type TCefTransitionType = types.Int

// /include/internal/cef_types.h (cef_urlrequest_status_t)
type TCefUrlRequestStatus = types.Int32

const (
	UR_UNKNOWN = TCefUrlRequestStatus(iota)
	UR_SUCCESS
	UR_IO_PENDING
	UR_CANCELED
	UR_FAILED
)

// /include/internal/cef_types.h (cef_state_t)
type TCefState = types.Int32

const (
	STATE_DEFAULT = TCefState(iota)
	STATE_ENABLED
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

// /include/internal/cef_types.h (cef_touch_event_type_t)
type TCefTouchEeventType = types.Int32

const (
	CEF_TET_RELEASED = TCefTouchEeventType(iota)
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

// /include/internal/cef_types.h (cef_pointer_type_t)
type TCefPointerType = types.Int32

const (
	CEF_POINTER_TYPE_TOUCH = TCefPointerType(iota)
	CEF_POINTER_TYPE_MOUSE
	CEF_POINTER_TYPE_PEN
	CEF_POINTER_TYPE_ERASER
	CEF_POINTER_TYPE_UNKNOWN
)

// /include/internal/cef_types.h (cef_mouse_button_type_t)
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

// /include/internal/cef_types.h (cef_window_open_disposition_t)
type TCefWindowOpenDisposition = types.Int32

const (
	WOD_UNKNOWN = TCefWindowOpenDisposition(iota)
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

// WINDOW_TYPE 窗口类型
type WINDOW_TYPE = types.Int8

const (
	WT_MAIN_BROWSER = WINDOW_TYPE(iota)
	WT_POPUP_SUB_BROWSER
	WT_DEV_TOOLS
	WT_VIEW_SOURCE
)

// /include/internal/cef_types.h (cef_context_menu_type_flags_t)
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

// /include/internal/cef_types.h (cef_context_menu_media_state_flags_t)
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

// /include/internal/cef_types.h (cef_context_menu_edit_state_flags_t)
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

// /include/internal/cef_types.h (cef_menu_anchor_position_t)
type TCefMenuAnchorPosition = types.Int32

const (
	CEF_MENU_ANCHOR_TOPLEFT = TCefMenuAnchorPosition(iota)
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

// /include/internal/cef_types.h (cef_docking_mode_t)
type TCefDockingMode = types.Int32

const (
	CEF_DOCKING_MODE_TOP_LEFT TCefDockingMode = iota + 1
	CEF_DOCKING_MODE_TOP_RIGHT
	CEF_DOCKING_MODE_BOTTOM_LEFT
	CEF_DOCKING_MODE_BOTTOM_RIGHT
	CEF_DOCKING_MODE_CUSTOM
)

// /include/internal/cef_types.h (cef_show_state_t)4
type TCefShowState = types.Int32

const (
	CEF_SHOW_STATE_NORMAL     = TCefShowState(1)
	CEF_SHOW_STATE_MINIMIZED  = TCefShowState(2)
	CEF_SHOW_STATE_MAXIMIZED  = TCefShowState(3)
	CEF_SHOW_STATE_FULLSCREEN = TCefShowState(4)
)

// /include/internal/cef_types.h (cef_chrome_toolbar_type_t)
type TCefChromeToolbarType = types.Int32

const (
	CEF_CTT_NONE     = TCefChromeToolbarType(1)
	CEF_CTT_NORMAL   = TCefChromeToolbarType(2)
	CEF_CTT_LOCATION = TCefChromeToolbarType(3)
)

// /include/internal/cef_types.h (cef_drag_operations_mask_t)
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

// /include/internal/cef_types.h (cef_drag_operations_mask_t)
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

// /include/internal/cef_types.h (cef_v8_accesscontrol_t)
type TCefV8AccessControls = types.Cardinal

const (
	V8_ACCESS_CONTROL_DEFAULT               TCefV8AccessControls = 0
	V8_ACCESS_CONTROL_ALL_CAN_READ          TCefV8AccessControls = 1 << 0
	V8_ACCESS_CONTROL_ALL_CAN_WRITE         TCefV8AccessControls = 1 << 1
	V8_ACCESS_CONTROL_PROHIBITS_OVERWRITING TCefV8AccessControls = 1 << 2
)

// /include/internal/cef_types.h (cef_v8_propertyattribute_t)
type TCefV8PropertyAttributes = types.Cardinal

const (
	V8_PROPERTY_ATTRIBUTE_NONE       TCefV8PropertyAttributes = 0
	V8_PROPERTY_ATTRIBUTE_READONLY   TCefV8PropertyAttributes = 1 << 0
	V8_PROPERTY_ATTRIBUTE_DONTENUM   TCefV8PropertyAttributes = 1 << 1
	V8_PROPERTY_ATTRIBUTE_DONTDELETE TCefV8PropertyAttributes = 1 << 2
)

// /include/internal/cef_types.h (cef_value_type_t)
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

// /include/internal/cef_types.h (cef_postdataelement_type_t)
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

// /include/internal/cef_types.h (cef_color_type_t)
type TCefColorType = types.Int32

const (
	CEF_COLOR_TYPE_RGBA_8888 TCefColorType = iota
	CEF_COLOR_TYPE_BGRA_8888
)

// /include/internal/cef_types.h (cef_alpha_type_t)
type TCefAlphaType = types.Int32

const (
	CEF_ALPHA_TYPE_OPAQUE TCefAlphaType = iota
	CEF_ALPHA_TYPE_PREMULTIPLIED
	CEF_ALPHA_TYPE_POSTMULTIPLIED
)

// /include/internal/cef_types.h (cef_pdf_print_margin_type_t)
type TCefPdfPrintMarginType = types.Int32

const (
	PDF_PRINT_MARGIN_DEFAULT TCefPdfPrintMarginType = iota
	PDF_PRINT_MARGIN_NONE
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

// /include/internal/cef_types.h (cef_preferences_type_t)
type TCefPreferencesType = types.Int32

const (
	CEF_PREFERENCES_TYPE_GLOBAL TCefPreferencesType = iota
	CEF_PREFERENCES_TYPE_REQUEST_CONTEXT
)

type TCefScaleFactor = types.Int32

// /include/internal/cef_types.h (cef_scale_factor_t)
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

// /include/internal/cef_types.h (cef_channel_layout_t)
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
//	CtSelfOwn 自己处理事件
//	CtOther   使用其他组件处理事件
type CefCreateType = types.Int8

const (
	CtSelfOwn CefCreateType = iota // CtSelfOwn 自己处理事件
	CtOther                        // CtOther   使用其他组件处理事件
)

// /include/internal/cef_types.h (cef_quick_menu_edit_state_flags_t)
type TCefQuickMenuEditStateFlags = types.Int32

// /include/internal/cef_types.h (cef_file_dialog_mode_t)
type TCefFileDialogMode = types.Cardinal

// /include/internal/cef_types.h (cef_log_severity_t)
type TCefLogSeverity = types.Cardinal

// TCefCursorHandle
//  /include/internal/cef_types_win.h (cef_cursor_handle_t)
//  /include/internal/cef_types_mac.h (cef_cursor_handle_t)
//  /include/internal/cef_types_linux.h (cef_cursor_handle_t)
type TCefCursorHandle uintptr

// /include/internal/cef_types.h (cef_cursor_type_t)
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

// /include/internal/cef_types.h (cef_focus_source_t)
type TCefFocusSource = types.Int32

const (
	FOCUS_SOURCE_NAVIGATION TCefFocusSource = iota
	FOCUS_SOURCE_SYSTEM
)

// /include/internal/cef_types.h (cef_permission_request_result_t)
type TCefPermissionRequestResult = types.Int32

const (
	CEF_PERMISSION_RESULT_ACCEPT TCefPermissionRequestResult = iota
	CEF_PERMISSION_RESULT_DENY
	CEF_PERMISSION_RESULT_DISMISS
	CEF_PERMISSION_RESULT_IGNORE
)

// /include/internal/cef_types.h (cef_media_access_permission_types_t)
type TCefMediaAccessPermissionTypes = types.Int32

// /include/internal/cef_types.h (cef_jsdialog_type_t)
type TCefJsDialogType = types.Int32

const (
	JSDIALOGTYPE_ALERT TCefJsDialogType = iota
	JSDIALOGTYPE_CONFIRM
	JSDIALOGTYPE_PROMPT
)

// TCefDuplexMode
//  /include/internal/cef_types.h (cef_duplex_mode_t)
type TCefDuplexMode = types.Int32

// /include/internal/cef_types.h (cef_color_model_t)
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

// /include/internal/cef_types.h (cef_scheme_options_t)
type CefSchemeOption = types.Int32

const (
	CEF_SCHEME_OPTION_NONE             CefSchemeOption = 0
	CEF_SCHEME_OPTION_STANDARD                         = 1 << 0
	CEF_SCHEME_OPTION_LOCAL                            = 1 << 1
	CEF_SCHEME_OPTION_DISPLAY_ISOLATED                 = 1 << 2
	CEF_SCHEME_OPTION_SECURE                           = 1 << 3
	CEF_SCHEME_OPTION_CORS_ENABLED                     = 1 << 4
	CEF_SCHEME_OPTION_CSP_BYPASSING                    = 1 << 5
	CEF_SCHEME_OPTION_FETCH_ENABLED                    = 1 << 6
)

// TCefResponseFilterStatus
//  /include/internal/cef_types.h (cef_response_filter_status_t)
type TCefResponseFilterStatus = types.Int32

const (
	RESPONSE_FILTER_NEED_MORE_DATA TCefResponseFilterStatus = iota
	RESPONSE_FILTER_DONE
	RESPONSE_FILTER_ERROR
)

// /include/internal/cef_types.h (cef_paint_element_type_t)
type TCefPaintElementType = types.Int32

const (
	PET_VIEW TCefPaintElementType = iota
	PET_POPUP
)

// /include/internal/cef_types.h (cef_horizontal_alignment_t)
type TCefHorizontalAlignment = types.Int32

const (
	CEF_HORIZONTAL_ALIGNMENT_LEFT TCefHorizontalAlignment = iota
	CEF_HORIZONTAL_ALIGNMENT_CENTER
	CEF_HORIZONTAL_ALIGNMENT_RIGHT
)

// /include/internal/cef_types.h (cef_text_input_mode_t)
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

// /include/internal/cef_types.h (cef_cert_status_t)
type TCefCertStatus = types.Int32

// /include/internal/cef_types.h (cef_media_route_create_result_t)
type TCefMediaRouterCreateResult = types.Int32

// /include/internal/cef_types.h (cef_media_route_connection_state_t)
type TCefMediaRouteConnectionState = types.Int32

const (
	CEF_MRCS_UNKNOWN TCefMediaRouteConnectionState = iota
	CEF_MRCS_CONNECTING
	CEF_MRCS_CONNECTED
	CEF_MRCS_CLOSED
	CEF_MRCS_TERMINATED
)

// /include/internal/cef_types.h (cef_dom_document_type_t)
type TCefDomDocumentType = types.Int32

const (
	DOM_DOCUMENT_TYPE_UNKNOWN TCefDomDocumentType = iota
	DOM_DOCUMENT_TYPE_HTML
	DOM_DOCUMENT_TYPE_XHTML
	DOM_DOCUMENT_TYPE_PLUGIN
)

// /include/internal/cef_types.h (cef_dom_node_type_t)
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

// /include/internal/cef_types.h (cef_composition_underline_style_t)
type TCefCompositionUnderlineStyle int32

const (
	CEF_CUS_SOLID TCefCompositionUnderlineStyle = iota
	CEF_CUS_DOT
	CEF_CUS_DASH
	CEF_CUS_NONE
)

//  MessageBox() Flags
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

// /include/internal/cef_types.h (cef_file_dialog_mode_t)
type FileDialogMode = types.Int32

const (
	FILE_DIALOG_OPEN          FileDialogMode = 0x00000000
	FILE_DIALOG_OPEN_MULTIPLE                = 0x00000001
	FILE_DIALOG_OPEN_FOLDER                  = 0x00000002
	FILE_DIALOG_SAVE                         = 0x00000003
)

// /include/internal/cef_types.h (cef_menu_item_type_t)
type TCefMenuItemType = types.Int32

const (
	MENUITEMTYPE_NONE TCefMenuItemType = iota
	MENUITEMTYPE_COMMAND
	MENUITEMTYPE_CHECK
	MENUITEMTYPE_RADIO
	MENUITEMTYPE_SEPARATOR
	MENUITEMTYPE_SUBMENU
)

// /include/internal/cef_types.h (cef_button_state_t)
type TCefButtonState = types.Int32

const (
	CEF_BUTTON_STATE_NORMAL TCefButtonState = iota
	CEF_BUTTON_STATE_HOVERED
	CEF_BUTTON_STATE_PRESSED
	CEF_BUTTON_STATE_DISABLED
)

// /include/internal/cef_types.h (cef_main_axis_alignment_t)
type TCefMainAxisAlignment = types.Int32

const (
	CEF_MAIN_AXIS_ALIGNMENT_START TCefMainAxisAlignment = iota
	CEF_MAIN_AXIS_ALIGNMENT_CENTER
	CEF_MAIN_AXIS_ALIGNMENT_END
)

// /include/internal/cef_types.h (cef_cross_axis_alignment_t)
type TCefCrossAxisAlignment = types.Int32

const (
	CEF_CROSS_AXIS_ALIGNMENT_STRETCH TCefCrossAxisAlignment = iota
	CEF_CROSS_AXIS_ALIGNMENT_START
	CEF_CROSS_AXIS_ALIGNMENT_CENTER
	CEF_CROSS_AXIS_ALIGNMENT_END
)

// /include/internal/cef_types.h (cef_text_style_t)
type TCefTextStyle = types.Int32

const (
	CEF_TEXT_STYLE_BOLD TCefTextStyle = iota
	CEF_TEXT_STYLE_ITALIC
	CEF_TEXT_STYLE_STRIKE
	CEF_TEXT_STYLE_DIAGONAL_STRIKE
	CEF_TEXT_STYLE_UNDERLINE
)

// /include/internal/cef_types.h (cef_text_field_commands_t)
type TCefTextFieldCommands = types.Int32

const (
	CEF_TFC_CUT TCefTextFieldCommands = iota + 1
	CEF_TFC_COPY
	CEF_TFC_PASTE
	CEF_TFC_UNDO
	CEF_TFC_DELETE
	CEF_TFC_SELECT_ALL
)
