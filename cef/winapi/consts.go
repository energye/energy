//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package winapi

// messages.WM_SIZE
const (
	SIZE_RESTORED  = 0
	SIZE_MINIMIZED = 1
	SIZE_MAXIMIZED = 2
	SIZE_MAXSHOW   = 3
	SIZE_MAXHIDE   = 4
)

const (
	// Mouse message key states
	MK_LBUTTON  = 1
	MK_RBUTTON  = 2
	MK_SHIFT    = 4
	MK_CONTROL  = 8
	MK_MBUTTON  = 0x10
	MK_XBUTTON1 = 0x20
	MK_XBUTTON2 = 0x40
	// following are "virtual" key states
	MK_DOUBLECLICK = 0x80
	MK_TRIPLECLICK = 0x100
	MK_QUADCLICK   = 0x200
	MK_ALT         = 0x20000000
)

// ------------
// KeyFlags (High word part !!!)
// ------------
const (
	KF_EXTENDED = 0x100
	KF_DLGMODE  = 0x800
	KF_MENUMODE = 0x1000
	KF_ALTDOWN  = 0x2000
	KF_REPEAT   = 0x4000
	KF_UP       = 0x8000
)

// -------------
// Virtual keys
// -------------
//
// Basic keys up to $FF have values and meaning compatible with the Windows API as described here:
// http://msdn.microsoft.com/library/default.asp?url=/library/en-us/winui/WinUI/WindowsUserInterface/UserInput/VirtualKeyCodes.asp
//
// Starting with $100 and upwards the key constants are LCL additions
const (
	VK_UNKNOWN  = 0 // defined by LCL
	VK_LBUTTON  = 1
	VK_RBUTTON  = 2
	VK_CANCEL   = 3
	VK_MBUTTON  = 4
	VK_XBUTTON1 = 5
	VK_XBUTTON2 = 6
	VK_BACK     = 8 // The "Backspace" key, dont confuse with the
	// Android BACK key which is mapped to VK_ESCAPE
	VK_TAB        = 9
	VK_CLEAR      = 12
	VK_RETURN     = 13 // The "Enter" key, also used for a keypad center press
	VK_SHIFT      = 16 // See also VK_LSHIFT, VK_RSHIFT
	VK_CONTROL    = 17 // See also VK_LCONTROL, VK_RCONTROL
	VK_MENU       = 18 // The ALT key. Also called "Option" in Mac OS X. See also VK_LMENU, VK_RMENU
	VK_PAUSE      = 19 // Pause/Break key
	VK_CAPITAL    = 20 // CapsLock key
	VK_KANA       = 21
	VK_HANGUL     = 21
	VK_JUNJA      = 23
	VK_FINAL      = 24
	VK_HANJA      = 25
	VK_KANJI      = 25
	VK_ESCAPE     = 27 // Also used for the hardware Back key in Android
	VK_CONVERT    = 28
	VK_NONCONVERT = 29
	VK_ACCEPT     = 30
	VK_MODECHANGE = 31
	VK_SPACE      = 32
	VK_PRIOR      = 33 // Page Up
	VK_NEXT       = 34 // Page Down
	VK_END        = 35
	VK_HOME       = 36
	VK_LEFT       = 37
	VK_UP         = 38
	VK_RIGHT      = 39
	VK_DOWN       = 40
	VK_SELECT     = 41
	VK_PRINT      = 42 // PrintScreen key
	VK_EXECUTE    = 43
	VK_SNAPSHOT   = 44
	VK_INSERT     = 45
	VK_DELETE     = 46
	VK_HELP       = 47
	VK_0          = 0x30
	VK_1          = 0x31
	VK_2          = 0x32
	VK_3          = 0x33
	VK_4          = 0x34
	VK_5          = 0x35
	VK_6          = 0x36
	VK_7          = 0x37
	VK_8          = 0x38
	VK_9          = 0x39
	//3A-40 Undefined
	VK_A    = 0x41
	VK_B    = 0x42
	VK_C    = 0x43
	VK_D    = 0x44
	VK_E    = 0x45
	VK_F    = 0x46
	VK_G    = 0x47
	VK_H    = 0x48
	VK_I    = 0x49
	VK_J    = 0x4A
	VK_K    = 0x4B
	VK_L    = 0x4C
	VK_M    = 0x4D
	VK_N    = 0x4E
	VK_O    = 0x4F
	VK_P    = 0x50
	VK_Q    = 0x51
	VK_R    = 0x52
	VK_S    = 0x53
	VK_T    = 0x54
	VK_U    = 0x55
	VK_V    = 0x56
	VK_W    = 0x57
	VK_X    = 0x58
	VK_Y    = 0x59
	VK_Z    = 0x5A
	VK_LWIN = 0x5B // In Mac OS X this is the Apple, or Command key. Windows Key in PC keyboards
	VK_RWIN = 0x5C // In Mac OS X this is the Apple, or Command key. Windows Key in PC keyboards
	VK_APPS = 0x5D // The PopUp key in PC keyboards
	// $5E reserved
	VK_SLEEP     = 0x5F
	VK_NUMPAD0   = 96 // $60
	VK_NUMPAD1   = 97
	VK_NUMPAD2   = 98
	VK_NUMPAD3   = 99
	VK_NUMPAD4   = 100
	VK_NUMPAD5   = 101
	VK_NUMPAD6   = 102
	VK_NUMPAD7   = 103
	VK_NUMPAD8   = 104
	VK_NUMPAD9   = 105
	VK_MULTIPLY  = 106 // VK_MULTIPLY up to VK_DIVIDE are usually in the numeric keypad in PC keyboards
	VK_ADD       = 107
	VK_SEPARATOR = 108
	VK_SUBTRACT  = 109
	VK_DECIMAL   = 110
	VK_DIVIDE    = 111
	VK_F1        = 112
	VK_F2        = 113
	VK_F3        = 114
	VK_F4        = 115
	VK_F5        = 116
	VK_F6        = 117
	VK_F7        = 118
	VK_F8        = 119
	VK_F9        = 120
	VK_F10       = 121
	VK_F11       = 122
	VK_F12       = 123
	VK_F13       = 124
	VK_F14       = 125
	VK_F15       = 126
	VK_F16       = 127
	VK_F17       = 128
	VK_F18       = 129
	VK_F19       = 130
	VK_F20       = 131
	VK_F21       = 132
	VK_F22       = 133
	VK_F23       = 134
	VK_F24       = 135 // $87

	// $88-$8F unassigned

	VK_NUMLOCK = 0x90
	VK_SCROLL  = 0x91

	// $92-$96  OEM specific
	// $97-$9F  Unassigned

	// not in VCL defined:
	//MWE: And should not be used.
	//     The keys they are on map to another VK

	//  VK_SEMICOLON  = 186;
	//  VK_EQUAL      = 187; // $BB
	//  VK_COMMA      = 188;
	//  VK_POINT      = 190;
	//  VK_SLASH      = 191;
	//  VK_AT         = 192;

	// VK_L & VK_R - left and right Alt, Ctrl and Shift virtual keys.
	// When Application.ExtendedKeysSupport is false, these keys are
	// used only as parameters to GetAsyncKeyState() and GetKeyState().
	// No other API or message will distinguish left and right keys in this way
	//
	// When Application.ExtendedKeysSupport is true, these keys will be sent
	// on KeyDown / KeyUp instead of the generic VK_SHIFT, VK_CONTROL, etc.
	VK_LSHIFT   = 0xA0
	VK_RSHIFT   = 0xA1
	VK_LCONTROL = 0xA2
	VK_RCONTROL = 0xA3
	VK_LMENU    = 0xA4 // Left ALT key (also named Option in Mac OS X)
	VK_RMENU    = 0xA5 // Right ALT key (also named Option in Mac OS X)

	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7

	// VK_OEM keys are utilized only when Application.ExtendedKeysSupport is false

	// $B8-$B9 Reserved
	VK_OEM_1 = 0xBA // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the ';:' key
	VK_OEM_PLUS   = 0xBB // For any country/region, the '+' key
	VK_OEM_COMMA  = 0xBC // For any country/region, the ',' key
	VK_OEM_MINUS  = 0xBD // For any country/region, the '-' key
	VK_OEM_PERIOD = 0xBE // For any country/region, the '.' key
	VK_OEM_2      = 0xBF // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the '/?' key
	VK_OEM_3 = 0xC0 // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the '`~' key
	// $C1-$D7 Reserved
	// $D8-$DA Unassigned
	VK_OEM_4 = 0xDB // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the '[{' key
	VK_OEM_5 = 0xDC // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the '\|' key
	VK_OEM_6 = 0xDD // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the ']}' key
	VK_OEM_7 = 0xDE // Used for miscellaneous characters; it can vary by keyboard.
	// For the US standard keyboard, the 'single-quote/double-quote' key
	VK_OEM_8 = 0xDF // Used for miscellaneous characters; it can vary by keyboard.

	// $E0 Reserved
	// $E1 OEM specific
	VK_OEM_102 = 0xE2 // Either the angle bracket key or the backslash key on the RT 102-key keyboard

	// $E3-$E4 OEM specific

	VK_PROCESSKEY = 0xE7 // IME Process key

	// $E8 Unassigned
	// $E9-$F5 OEM specific

	VK_ATTN         = 0xF6
	VK_CRSEL        = 0xF7
	VK_EXSEL        = 0xF8
	VK_EREOF        = 0xF9
	VK_PLAY         = 0xFA
	VK_ZOOM         = 0xFB
	VK_NONAME       = 0xFC
	VK_PA1          = 0xFD
	VK_OEM_CLEAR    = 0xFE
	VK_HIGHESTVALUE = 0xFFFF
	VK_UNDEFINED    = 0xFF // defined by LCL

	//==============================================
	// LCL aliases for more clear naming of keys
	//==============================================

	VK_LCL_EQUAL      = VK_OEM_PLUS   // The "=+" Key
	VK_LCL_COMMA      = VK_OEM_COMMA  // The ",<" Key
	VK_LCL_POINT      = VK_OEM_PERIOD // The ".>" Key
	VK_LCL_SLASH      = VK_OEM_2      // The "/?" Key
	VK_LCL_SEMI_COMMA = VK_OEM_1      // The ";:" Key
	VK_LCL_MINUS      = VK_OEM_MINUS  // The "-_" Key

	VK_LCL_OPEN_BRAKET   = VK_OEM_4 //deprecated 'Use VK_LCL_OPEN_BRACKET instead';
	VK_LCL_CLOSE_BRAKET  = VK_OEM_6 //deprecated 'Use VK_LCL_CLOSE_BRACKET instead';
	VK_LCL_OPEN_BRACKET  = VK_OEM_4 // The "[{" Key
	VK_LCL_CLOSE_BRACKET = VK_OEM_6 // The "]}" Key

	VK_LCL_BACKSLASH = VK_OEM_5 // The "\|" Key
	VK_LCL_TILDE     = VK_OEM_3 // The "`~" Key
	VK_LCL_QUOTE     = VK_OEM_7 // The "'"" Key

	VK_LCL_ALT      = VK_MENU
	VK_LCL_LALT     = VK_LMENU
	VK_LCL_RALT     = VK_RMENU
	VK_LCL_CAPSLOCK = VK_CAPITAL

	//==============================================
	// New LCL defined keys
	//==============================================

	VK_LCL_POWER   = 0x100
	VK_LCL_CALL    = 0x101
	VK_LCL_ENDCALL = 0x102
	VK_LCL_AT      = 0x103 // Not equivalent to anything < $FF, will only be sent by a primary "@" key
	// but not for a @ key as secondary action of a "2" key for example
)

type MONITOR_DPI_TYPE int32

const (
	MDT_EFFECTIVE_DPI MONITOR_DPI_TYPE = 0
	MDT_ANGULAR_DPI   MONITOR_DPI_TYPE = 1
	MDT_RAW_DPI       MONITOR_DPI_TYPE = 2
	MDT_DEFAULT       MONITOR_DPI_TYPE = 0
)

const (
	DISPID_UNKNOWN     = -1
	DISPID_VALUE       = 0
	DISPID_PROPERTYPUT = -3
	DISPID_NEWENUM     = -4
	DISPID_EVALUATE    = -5
	DISPID_CONSTRUCTOR = -6
	DISPID_DESTRUCTOR  = -7
	DISPID_COLLECT     = -8
)

const (
	MONITOR_DEFAULTTONULL    = 0x00000000
	MONITOR_DEFAULTTOPRIMARY = 0x00000001
	MONITOR_DEFAULTTONEAREST = 0x00000002
	MONITORINFOF_PRIMARY     = 0x00000001
)

// DeviceCapabilities capabilities
const (
	DC_FIELDS            = 1
	DC_PAPERS            = 2
	DC_PAPERSIZE         = 3
	DC_MINEXTENT         = 4
	DC_MAXEXTENT         = 5
	DC_BINS              = 6
	DC_DUPLEX            = 7
	DC_SIZE              = 8
	DC_EXTRA             = 9
	DC_VERSION           = 10
	DC_DRIVER            = 11
	DC_BINNAMES          = 12
	DC_ENUMRESOLUTIONS   = 13
	DC_FILEDEPENDENCIES  = 14
	DC_TRUETYPE          = 15
	DC_PAPERNAMES        = 16
	DC_ORIENTATION       = 17
	DC_COPIES            = 18
	DC_BINADJUST         = 19
	DC_EMF_COMPLIANT     = 20
	DC_DATATYPE_PRODUCED = 21
	DC_COLLATE           = 22
	DC_MANUFACTURER      = 23
	DC_MODEL             = 24
	DC_PERSONALITY       = 25
	DC_PRINTRATE         = 26
	DC_PRINTRATEUNIT     = 27
	DC_PRINTERMEM        = 28
	DC_MEDIAREADY        = 29
	DC_STAPLE            = 30
	DC_PRINTRATEPPM      = 31
	DC_COLORDEVICE       = 32
	DC_NUP               = 33
	DC_MEDIATYPENAMES    = 34
	DC_MEDIATYPES        = 35
)

// GetDeviceCaps index constants
const (
	DRIVERVERSION   = 0
	TECHNOLOGY      = 2
	HORZSIZE        = 4
	VERTSIZE        = 6
	HORZRES         = 8
	VERTRES         = 10
	LOGPIXELSX      = 88
	LOGPIXELSY      = 90
	BITSPIXEL       = 12
	PLANES          = 14
	NUMBRUSHES      = 16
	NUMPENS         = 18
	NUMFONTS        = 22
	NUMCOLORS       = 24
	NUMMARKERS      = 20
	ASPECTX         = 40
	ASPECTY         = 42
	ASPECTXY        = 44
	PDEVICESIZE     = 26
	CLIPCAPS        = 36
	SIZEPALETTE     = 104
	NUMRESERVED     = 106
	COLORRES        = 108
	PHYSICALWIDTH   = 110
	PHYSICALHEIGHT  = 111
	PHYSICALOFFSETX = 112
	PHYSICALOFFSETY = 113
	SCALINGFACTORX  = 114
	SCALINGFACTORY  = 115
	VREFRESH        = 116
	DESKTOPHORZRES  = 118
	DESKTOPVERTRES  = 117
	BLTALIGNMENT    = 119
	SHADEBLENDCAPS  = 120
	COLORMGMTCAPS   = 121
	RASTERCAPS      = 38
	CURVECAPS       = 28
	LINECAPS        = 30
	POLYGONALCAPS   = 32
	TEXTCAPS        = 34
)

// GetSystemMetrics constants
const (
	SM_CXSCREEN             = 0
	SM_CYSCREEN             = 1
	SM_CXVSCROLL            = 2
	SM_CYHSCROLL            = 3
	SM_CYCAPTION            = 4
	SM_CXBORDER             = 5
	SM_CYBORDER             = 6
	SM_CXDLGFRAME           = 7
	SM_CYDLGFRAME           = 8
	SM_CYVTHUMB             = 9
	SM_CXHTHUMB             = 10
	SM_CXICON               = 11
	SM_CYICON               = 12
	SM_CXCURSOR             = 13
	SM_CYCURSOR             = 14
	SM_CYMENU               = 15
	SM_CXFULLSCREEN         = 16
	SM_CYFULLSCREEN         = 17
	SM_CYKANJIWINDOW        = 18
	SM_MOUSEPRESENT         = 19
	SM_CYVSCROLL            = 20
	SM_CXHSCROLL            = 21
	SM_DEBUG                = 22
	SM_SWAPBUTTON           = 23
	SM_RESERVED1            = 24
	SM_RESERVED2            = 25
	SM_RESERVED3            = 26
	SM_RESERVED4            = 27
	SM_CXMIN                = 28
	SM_CYMIN                = 29
	SM_CXSIZE               = 30
	SM_CYSIZE               = 31
	SM_CXFRAME              = 32
	SM_CYFRAME              = 33
	SM_CXMINTRACK           = 34
	SM_CYMINTRACK           = 35
	SM_CXDOUBLECLK          = 36
	SM_CYDOUBLECLK          = 37
	SM_CXICONSPACING        = 38
	SM_CYICONSPACING        = 39
	SM_MENUDROPALIGNMENT    = 40
	SM_PENWINDOWS           = 41
	SM_DBCSENABLED          = 42
	SM_CMOUSEBUTTONS        = 43
	SM_CXFIXEDFRAME         = SM_CXDLGFRAME
	SM_CYFIXEDFRAME         = SM_CYDLGFRAME
	SM_CXSIZEFRAME          = SM_CXFRAME
	SM_CYSIZEFRAME          = SM_CYFRAME
	SM_SECURE               = 44
	SM_CXEDGE               = 45
	SM_CYEDGE               = 46
	SM_CXMINSPACING         = 47
	SM_CYMINSPACING         = 48
	SM_CXSMICON             = 49
	SM_CYSMICON             = 50
	SM_CYSMCAPTION          = 51
	SM_CXSMSIZE             = 52
	SM_CYSMSIZE             = 53
	SM_CXMENUSIZE           = 54
	SM_CYMENUSIZE           = 55
	SM_ARRANGE              = 56
	SM_CXMINIMIZED          = 57
	SM_CYMINIMIZED          = 58
	SM_CXMAXTRACK           = 59
	SM_CYMAXTRACK           = 60
	SM_CXMAXIMIZED          = 61
	SM_CYMAXIMIZED          = 62
	SM_NETWORK              = 63
	SM_CLEANBOOT            = 67
	SM_CXDRAG               = 68
	SM_CYDRAG               = 69
	SM_SHOWSOUNDS           = 70
	SM_CXMENUCHECK          = 71
	SM_CYMENUCHECK          = 72
	SM_SLOWMACHINE          = 73
	SM_MIDEASTENABLED       = 74
	SM_MOUSEWHEELPRESENT    = 75
	SM_XVIRTUALSCREEN       = 76
	SM_YVIRTUALSCREEN       = 77
	SM_CXVIRTUALSCREEN      = 78
	SM_CYVIRTUALSCREEN      = 79
	SM_CMONITORS            = 80
	SM_SAMEDISPLAYFORMAT    = 81
	SM_IMMENABLED           = 82
	SM_CXFOCUSBORDER        = 83
	SM_CYFOCUSBORDER        = 84
	SM_TABLETPC             = 86
	SM_MEDIACENTER          = 87
	SM_STARTER              = 88
	SM_SERVERR2             = 89
	SM_CMETRICS             = 91
	SM_REMOTESESSION        = 0x1000
	SM_SHUTTINGDOWN         = 0x2000
	SM_REMOTECONTROL        = 0x2001
	SM_CARETBLINKINGENABLED = 0x2002
)

// GetDeviceCaps TECHNOLOGY constants
const (
	DT_PLOTTER    = 0
	DT_RASDISPLAY = 1
	DT_RASPRINTER = 2
	DT_RASCAMERA  = 3
	DT_CHARSTREAM = 4
	DT_METAFILE   = 5
	DT_DISPFILE   = 6
)

// GetDeviceCaps SHADEBLENDCAPS constants
const (
	SB_NONE          = 0x00
	SB_CONST_ALPHA   = 0x01
	SB_PIXEL_ALPHA   = 0x02
	SB_PREMULT_ALPHA = 0x04
	SB_GRAD_RECT     = 0x10
	SB_GRAD_TRI      = 0x20
)

// GetDeviceCaps COLORMGMTCAPS constants
const (
	CM_NONE       = 0x00
	CM_DEVICE_ICM = 0x01
	CM_GAMMA_RAMP = 0x02
	CM_CMYK_COLOR = 0x04
)

// GetDeviceCaps RASTERCAPS constants
const (
	RC_BANDING      = 2
	RC_BITBLT       = 1
	RC_BITMAP64     = 8
	RC_DI_BITMAP    = 128
	RC_DIBTODEV     = 512
	RC_FLOODFILL    = 4096
	RC_GDI20_OUTPUT = 16
	RC_PALETTE      = 256
	RC_SCALING      = 4
	RC_STRETCHBLT   = 2048
	RC_STRETCHDIB   = 8192
	RC_DEVBITS      = 0x8000
	RC_OP_DX_OUTPUT = 0x4000
)

// GetDeviceCaps CURVECAPS constants
const (
	CC_NONE       = 0
	CC_CIRCLES    = 1
	CC_PIE        = 2
	CC_CHORD      = 4
	CC_ELLIPSES   = 8
	CC_WIDE       = 16
	CC_STYLED     = 32
	CC_WIDESTYLED = 64
	CC_INTERIORS  = 128
	CC_ROUNDRECT  = 256
)

// GetDeviceCaps LINECAPS constants
const (
	LC_NONE       = 0
	LC_POLYLINE   = 2
	LC_MARKER     = 4
	LC_POLYMARKER = 8
	LC_WIDE       = 16
	LC_STYLED     = 32
	LC_WIDESTYLED = 64
	LC_INTERIORS  = 128
)

// GetDeviceCaps POLYGONALCAPS constants
const (
	PC_NONE        = 0
	PC_POLYGON     = 1
	PC_POLYPOLYGON = 256
	PC_PATHS       = 512
	PC_RECTANGLE   = 2
	PC_WINDPOLYGON = 4
	PC_SCANLINE    = 8
	PC_TRAPEZOID   = 4
	PC_WIDE        = 16
	PC_STYLED      = 32
	PC_WIDESTYLED  = 64
	PC_INTERIORS   = 128
)

// GetDeviceCaps TEXTCAPS constants
const (
	TC_OP_CHARACTER = 1
	TC_OP_STROKE    = 2
	TC_CP_STROKE    = 4
	TC_CR_90        = 8
	TC_CR_ANY       = 16
	TC_SF_X_YINDEP  = 32
	TC_SA_DOUBLE    = 64
	TC_SA_INTEGER   = 128
	TC_SA_CONTIN    = 256
	TC_EA_DOUBLE    = 512
	TC_IA_ABLE      = 1024
	TC_UA_ABLE      = 2048
	TC_SO_ABLE      = 4096
	TC_RA_ABLE      = 8192
	TC_VA_ABLE      = 16384
	TC_RESERVED     = 32768
	TC_SCROLLBLT    = 65536
)
