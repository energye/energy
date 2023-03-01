//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package consts

import (
	"github.com/energye/energy/types"
	"github.com/energye/golcl/energy/homedir"
	"math"
	"os"
	"path/filepath"
)

var (
	ExePath       string                       //执行文件目录
	HomeDir, _    = homedir.Dir()              //系统用户目录
	Separator     = string(filepath.Separator) //系统目录分隔符
	SingleProcess = false                      //进程启动方式, true 单进程 false 多进程
	IsMessageLoop = false                      //CEF应用的窗口, true: 使用VF(views framework)窗口组件, false: 使用LCL窗口组件, 其实是窗口消息轮询使用方式.
)

const (
	Empty               = ""
	MAINARGS_NETIPCPORT = "net-ipc-port"
	ENERGY_HOME_KEY     = "ENERGY_HOME"
	MemoryNetwork       = "unix"
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

// 进程结束的状态
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
type CBS = types.Int32

const (
	CbaClose = CBS(iota)
	CbaDelay
	CbaCancel
)

// CEF 进程 ChannelId
type CefProcessId = types.Int32

const (
	PID_BROWSER = CefProcessId(iota)
	PID_RENDER
)

// 支持的JS类型
type V8_JS_VALUE_TYPE = types.Int32

const (
	V8_VALUE_STRING = V8_JS_VALUE_TYPE(iota)
	V8_VALUE_INT
	V8_VALUE_DOUBLE
	V8_VALUE_BOOLEAN
	V8_VALUE_NULL
	V8_VALUE_UNDEFINED
	V8_VALUE_OBJECT
	V8_VALUE_ARRAY
	V8_VALUE_FUNCTION
	V8_VALUE_PTR
	V8_VALUE_EXCEPTION
	V8_VALUE_ROOT_OBJECT
	V8_NO_OUT_VALUE
)

// 支持的GO类型
type GO_VALUE_TYPE = types.Int32

const (
	GO_VALUE_STRING = GO_VALUE_TYPE(iota)
	GO_VALUE_INT
	GO_VALUE_INT8
	GO_VALUE_INT16
	GO_VALUE_INT32
	GO_VALUE_INT64
	GO_VALUE_UINT
	GO_VALUE_UINT8
	GO_VALUE_UINT16
	GO_VALUE_UINT32
	GO_VALUE_UINT64
	GO_VALUE_UINTPTR
	GO_VALUE_FLOAT32
	GO_VALUE_FLOAT64
	GO_VALUE_BOOL
	GO_VALUE_NIL
	GO_VALUE_STRUCT
	GO_VALUE_SLICE
	GO_VALUE_FUNC
	GO_VALUE_PTR
	GO_VALUE_MAP
	GO_VALUE_EXCEPTION
	GO_VALUE_INVALID_TYPE //无效类型
	GO_VALUE_ARGUMENT     //argument
	GO_VALUE_DICTVALUE    //dictValue
)

// JS交互绑定的事件类型
type BIND_EVENT = types.Int32

const (
	BE_SET = BIND_EVENT(iota)
	BE_GET
	BE_FUNC
	BE_CTX_CRT_BIND
)

// 异常信息
type CEF_V8_EXCEPTION = types.Int32

const (
	CVE_ERROR_OK                             = CEF_V8_EXCEPTION(iota) //操作成功
	CVE_ERROR_NOT_FOUND_FIELD                                         //未找到字段 或字段未定义
	CVE_ERROR_NOT_FOUND_FUNC                                          //未找到函数 或函数未定义
	CVE_ERROR_TYPE_NOT_SUPPORTED                                      //不支持的变量类型 变量类型只支持[string int double bool null undefined]
	CVE_ERROR_TYPE_CANNOT_CHANGE                                      //字段为普通类型不能变更为 array、object、function
	CVE_ERROR_TYPE_INVALID                                            //类型无效
	CVE_ERROR_GET_STRING_FAIL                                         //获取string类型失败
	CVE_ERROR_GET_INT_FAIL                                            //获取int类型失败
	CVE_ERROR_GET_DOUBLE_FAIL                                         //获取double类型失败
	CVE_ERROR_GET_BOOL_FAIL                                           //获取bool类型失败
	CVE_ERROR_GET_NULL_FAIL                                           //获取null类型失败
	CVE_ERROR_GET_UNDEFINED_FAIL                                      //获取undefined类型失败
	CVE_ERROR_FUNC_INVALID_P_L_9                                      //该函数非法 类型不正确 或参数个数大于9个
	CVE_ERROR_FUNC_IN_PAM                                             //入参类型不正确 只能为string int double boolean
	CVE_ERROR_FUNC_OUT_PAM                                            //出参类型不正确 只能为EefError 或 可选的[string int double boolean]
	CVE_ERROR_FUNC_GET_IN_PAM_STRING_FAIL                             //入参获取string类型值失败
	CVE_ERROR_FUNC_GET_IN_PAM_INT_FAIL                                //入参获取int类型值失败
	CVE_ERROR_FUNC_GET_IN_PAM_DOUBLE_FAIL                             //入参获取double类型值失败
	CVE_ERROR_FUNC_GET_IN_PAM_BOOLEAN_FAIL                            //入参获取boolean类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_STRING_FAIL                            //出参获取string类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_INT_FAIL                               //出参获取int类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_DOUBLE_FAIL                            //出参获取double类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_BOOLEAN_FAIL                           //出参获取boolean类型值失败
	CVE_ERROR_FUNC_GET_OUT_PAM_CEFERROR_FAIL                          //出参获取CefError值失败
	CVE_ERROR_IPC_GET_BIND_FIELD_VALUE_FAIL                           //IPC获取绑定值失败
	CVE_ERROR_UNKNOWN_ERROR                                           //未知错误
)

const (
	BIND_FUNC_IN_MAX_SUM  = 9 //函数最大入参数
	BIND_FUNC_OUT_MAX_SUM = 1 //函数最大出参数
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

type TCefCookieSameSite = types.Int32

const (
	Ccss_CEF_COOKIE_SAME_SITE_UNSPECIFIED = TCefCookieSameSite(iota)
	Ccss_CEF_COOKIE_SAME_SITE_NO_RESTRICTION
	Ccss_CEF_COOKIE_SAME_SITE_LAX_MODE
	Ccss_CEF_COOKIE_SAME_SITE_STRICT_MODE
)

type TCefCookiePriority = types.Int32

const (
	CEF_COOKIE_PRIORITY_LOW    TCefCookiePriority = -1
	CEF_COOKIE_PRIORITY_MEDIUM TCefCookiePriority = 0
	CEF_COOKIE_PRIORITY_HIGH   TCefCookiePriority = 1
)

type TCefProxyType = types.Int32

const (
	PtDirect = TCefProxyType(iota)
	PtAutodetect
	PtSystem
	PtFixedServers
	PtPACScript
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

type TCefKeyEventType = types.Int32

const (
	KEYEVENT_RAW_KEYDOWN = TCefKeyEventType(iota)
	KEYEVENT_KEYDOWN
	KEYEVENT_KEYUP
	KEYEVENT_CHAR
)

type TCefEventFlags = types.UInt32

type TCefWindowHandleType = types.Int8

type TCefWindowHandle = types.UIntptr

const (
	Wht_WindowParent = TCefWindowHandleType(iota)
	Wht_LinkedWindowParent
)

type TCefReturnValue = types.Int32

const (
	RV_CANCEL = TCefReturnValue(iota)
	RV_CONTINUE
	RV_CONTINUE_ASYNC
)

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

type TCefUrlRequestFlags = types.Int

type TCefErrorCode = types.Int32

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

type TCefTransitionType = types.Int

type TCefUrlRequestStatus = types.Int32

const (
	UR_UNKNOWN = TCefUrlRequestStatus(iota)
	UR_SUCCESS
	UR_IO_PENDING
	UR_CANCELED
	UR_FAILED
)

type TCefState = types.Int32

const (
	STATE_DEFAULT = TCefState(iota)
	STATE_ENABLED
	STATE_DISABLED
)

type TCefTouchEeventType = types.Int32

const (
	CEF_TET_RELEASED = TCefTouchEeventType(iota)
	CEF_TET_PRESSED
	CEF_TET_MOVED
	CEF_TET_CANCELLED
)

type TCefPointerType = types.Int32

const (
	CEF_POINTER_TYPE_TOUCH = TCefPointerType(iota)
	CEF_POINTER_TYPE_MOUSE
	CEF_POINTER_TYPE_PEN
	CEF_POINTER_TYPE_ERASER
	CEF_POINTER_TYPE_UNKNOWN
)

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

type TCefContextMenuTypeFlags = types.UInt32
type TCefContextMenuMediaStateFlags = types.UInt32
type TCefContextMenuEditStateFlags = types.UInt32

type TCefMenuAnchorPosition = types.Int32

const (
	CEF_MENU_ANCHOR_TOPLEFT = TCefMenuAnchorPosition(iota)
	CEF_MENU_ANCHOR_TOPRIGHT
	CEF_MENU_ANCHOR_BOTTOMCENTER
)

type TCefShowState = types.Int32

const (
	CEF_SHOW_STATE_NORMAL     = TCefShowState(1)
	CEF_SHOW_STATE_MINIMIZED  = TCefShowState(2)
	CEF_SHOW_STATE_MAXIMIZED  = TCefShowState(3)
	CEF_SHOW_STATE_FULLSCREEN = TCefShowState(4)
)

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

type TCefValueType = types.Int32

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

type TCefPdfPrintMarginType = types.Int32

const (
	PDF_PRINT_MARGIN_DEFAULT TCefPdfPrintMarginType = iota
	PDF_PRINT_MARGIN_NONE
	PDF_PRINT_MARGIN_CUSTOM
)
