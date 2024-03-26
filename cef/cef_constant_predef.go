//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/types"
	"math"
)

// SpecificVersion 特定版本: CEF49，CEF87，CEF106，CEF109
type SpecificVersion = types.Int32

const (
	SvINVALID = SpecificVersion(-1)  // 无效
	SvCEF     = SpecificVersion(0)   // CEF 非特定版本，当前版本或当前最新版本
	SvCEF49   = SpecificVersion(49)  // 特定 WindowsXP
	SvCEF87   = SpecificVersion(87)  // 特定 Flash
	SvCEF106  = SpecificVersion(106) // 特定 Linux GTK2
	SvCEF109  = SpecificVersion(109) // 特定 Windows 7, 8/8.1 and Windows Server 2012
)

type TCefString = string

// TCefWindowHandle
//
//	Native Window handle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_win.h">CEF source file: /include/internal/cef_types_win.h (cef_window_handle_t)</a>
type TCefWindowHandle = types.HWND

// TCefCursorHandle
//
//	Native Cursor handle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_win.h">CEF source file: /include/internal/cef_types_win.h (cef_cursor_handle_t)</a>
type TCefCursorHandle uintptr

// TCefResourceHandlerClass
type TCefResourceHandlerClass uintptr

// TCefProxyType
type TCefProxyType = int32

const (
	PtDirect       TCefProxyType = iota // mode dict => direct
	PtAutodetect                        // mode dict => auto_detect
	PtSystem                            // mode dict => system
	PtFixedServers                      // mode dict => fixed_servers
	PtPACScript                         // mode dict => pac_script
)

// TCefEventFlags
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

// TCefEventFlags
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

// TCefDragOperations
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

// TCefV8AccessControls
const (
	V8_ACCESS_CONTROL_DEFAULT               TCefV8AccessControls = 0
	V8_ACCESS_CONTROL_ALL_CAN_READ          TCefV8AccessControls = 1 << 0
	V8_ACCESS_CONTROL_ALL_CAN_WRITE         TCefV8AccessControls = 1 << 1
	V8_ACCESS_CONTROL_PROHIBITS_OVERWRITING TCefV8AccessControls = 1 << 2
)

// TCefV8PropertyAttributes
const (
	V8_PROPERTY_ATTRIBUTE_NONE       TCefV8PropertyAttributes = 0
	V8_PROPERTY_ATTRIBUTE_READONLY   TCefV8PropertyAttributes = 1 << 0
	V8_PROPERTY_ATTRIBUTE_DONTENUM   TCefV8PropertyAttributes = 1 << 1
	V8_PROPERTY_ATTRIBUTE_DONTDELETE TCefV8PropertyAttributes = 1 << 2
)

// TCefContextMenuTypeFlags
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

// TCefContextMenuMediaStateFlags
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

// TCefContextMenuEditStateFlags
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

// TCefCookiePriority
const (
	CEF_COOKIE_PRIORITY_LOW    TCefCookiePriority = -1
	CEF_COOKIE_PRIORITY_MEDIUM TCefCookiePriority = 0
	CEF_COOKIE_PRIORITY_HIGH   TCefCookiePriority = 1
)

// TCefTextFieldCommands
const (
	CEF_TFC_CUT TCefTextFieldCommands = iota + 1
	CEF_TFC_COPY
	CEF_TFC_PASTE
	CEF_TFC_UNDO
	CEF_TFC_DELETE
	CEF_TFC_SELECT_ALL
)

// TCefChromeToolbarType
const (
	CEF_CTT_NONE TCefChromeToolbarType = iota + 1
	CEF_CTT_NORMAL
	CEF_CTT_LOCATION
)

// TCefDockingMode
const (
	CEF_DOCKING_MODE_TOP_LEFT TCefDockingMode = iota + 1
	CEF_DOCKING_MODE_TOP_RIGHT
	CEF_DOCKING_MODE_BOTTOM_LEFT
	CEF_DOCKING_MODE_BOTTOM_RIGHT
	CEF_DOCKING_MODE_CUSTOM
)

// TCefShowState
const (
	CEF_SHOW_STATE_NORMAL TCefShowState = iota + 1
	CEF_SHOW_STATE_MINIMIZED
	CEF_SHOW_STATE_MAXIMIZED
	CEF_SHOW_STATE_FULLSCREEN
)

// TCefPermissionRequestTypes
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

// TCefLogItems
const (
	/// Prepend the default list of items.
	LOG_ITEMS_DEFAULT TCefLogItems = 0
	/// Prepend no items.
	LOG_ITEMS_NONE TCefLogItems = 1
	/// Prepend the process ID.
	LOG_ITEMS_FLAG_PROCESS_ID TCefLogItems = 1 << 1
	/// Prepend the thread ID.
	LOG_ITEMS_FLAG_THREAD_ID TCefLogItems = 1 << 2
	/// Prepend the timestamp.
	LOG_ITEMS_FLAG_TIME_STAMP TCefLogItems = 1 << 3
	/// Prepend the tickcount.
	LOG_ITEMS_FLAG_TICK_COUNT TCefLogItems = 1 << 4
)

type TSentinelStatus = int32

const (
	SsIdle TSentinelStatus = iota
	SsInitialDelay
	SsCheckingChildren
	SsClosing
)

// MenuId ContextMenuId
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

const (
	LOGSEVERITY_DEFAULT TCefLogSeverity = 0
	LOGSEVERITY_VERBOSE TCefLogSeverity = 1
	LOGSEVERITY_DEBUG   TCefLogSeverity = LOGSEVERITY_VERBOSE
	LOGSEVERITY_INFO    TCefLogSeverity = 2
	LOGSEVERITY_WARNING TCefLogSeverity = 3
	LOGSEVERITY_ERROR   TCefLogSeverity = 4
	LOGSEVERITY_FATAL   TCefLogSeverity = 5
	LOGSEVERITY_DISABLE TCefLogSeverity = 99
)

// LANGUAGE 本地语言
type LANGUAGE = string

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

// UITool UI tool classes used
type UITool int8

const (
	UitInvalid UITool = iota - 1 // invalid
	UitWin32                     // windows
	UitGtk2                      // linux
	UitGtk3                      // linux
	UitCocoa                     // macos
)
