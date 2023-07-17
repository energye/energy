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

/* Button Notification Codes */
const (
	BN_CLICKED       = 0
	BN_PAINT         = 1
	BN_HILITE        = 2
	BN_UNHILITE      = 3
	BN_DISABLE       = 4
	BN_DOUBLECLICKED = 5
	BN_PUSHED        = BN_HILITE
	BN_UNPUSHED      = BN_UNHILITE
	BN_DBLCLK        = BN_DOUBLECLICKED
	BN_SETFOCUS      = 6
	BN_KILLFOCUS     = 7
)

/* Button Control Messages */
const (
	BM_GETCHECK     = 0x00F0
	BM_SETCHECK     = 0x00F1
	BM_GETSTATE     = 0x00F2
	BM_SETSTATE     = 0x00F3
	BM_SETSTYLE     = 0x00F4
	BM_CLICK        = 0x00F5
	BM_GETIMAGE     = 0x00F6
	BM_SETIMAGE     = 0x00F7
	BM_SETDONTCLICK = 0x00F8
)

/* Listbox Notification Codes */
const (
	LBN_ERRSPACE  = -2
	LBN_SELCHANGE = 1
	LBN_DBLCLK    = 2
	LBN_SELCANCEL = 3
	LBN_SETFOCUS  = 4
	LBN_KILLFOCUS = 5
)

/* Listbox messages */
const (
	LB_ADDSTRING           = 0x0180
	LB_INSERTSTRING        = 0x0181
	LB_DELETESTRING        = 0x0182
	LB_SELITEMRANGEEX      = 0x0183
	LB_RESETCONTENT        = 0x0184
	LB_SETSEL              = 0x0185
	LB_SETCURSEL           = 0x0186
	LB_GETSEL              = 0x0187
	LB_GETCURSEL           = 0x0188
	LB_GETTEXT             = 0x0189
	LB_GETTEXTLEN          = 0x018A
	LB_GETCOUNT            = 0x018B
	LB_SELECTSTRING        = 0x018C
	LB_DIR                 = 0x018D
	LB_GETTOPINDEX         = 0x018E
	LB_FINDSTRING          = 0x018F
	LB_GETSELCOUNT         = 0x0190
	LB_GETSELITEMS         = 0x0191
	LB_SETTABSTOPS         = 0x0192
	LB_GETHORIZONTALEXTENT = 0x0193
	LB_SETHORIZONTALEXTENT = 0x0194
	LB_SETCOLUMNWIDTH      = 0x0195
	LB_ADDFILE             = 0x0196
	LB_SETTOPINDEX         = 0x0197
	LB_GETITEMRECT         = 0x0198
	LB_GETITEMDATA         = 0x0199
	LB_SETITEMDATA         = 0x019A
	LB_SELITEMRANGE        = 0x019B
	LB_SETANCHORINDEX      = 0x019C
	LB_GETANCHORINDEX      = 0x019D
	LB_SETCARETINDEX       = 0x019E
	LB_GETCARETINDEX       = 0x019F
	LB_SETITEMHEIGHT       = 0x01A0
	LB_GETITEMHEIGHT       = 0x01A1
	LB_FINDSTRINGEXACT     = 0x01A2
	LB_SETLOCALE           = 0x01A5
	LB_GETLOCALE           = 0x01A6
	LB_SETCOUNT            = 0x01A7
	LB_INITSTORAGE         = 0x01A8
	LB_ITEMFROMPOINT       = 0x01A9
	LB_MSGMAX              = 0x01B3 /* if _WIN32_WINNT >= 0x0501 */
	//LB_MSGMAX             = 0x01B1 /* else if _WIN32_WCE >= 0x0400 */
	//LB_MSGMAX             = 0x01B0 /* else if WINVER >= 0x0400 */
	//LB_MSGMAX             = 0x01A8] /* else */
)

/* Combo Box Notification Codes */
const (
	CBN_ERRSPACE     = -1
	CBN_SELCHANGE    = 1
	CBN_DBLCLK       = 2
	CBN_SETFOCUS     = 3
	CBN_KILLFOCUS    = 4
	CBN_EDITCHANGE   = 5
	CBN_EDITUPDATE   = 6
	CBN_DROPDOWN     = 7
	CBN_CLOSEUP      = 8
	CBN_SELENDOK     = 9
	CBN_SELENDCANCEL = 10

	/* Combo Box messages */

	CB_GETEDITSEL            = 0x0140
	CB_LIMITTEXT             = 0x0141
	CB_SETEDITSEL            = 0x0142
	CB_ADDSTRING             = 0x0143
	CB_DELETESTRING          = 0x0144
	CB_DIR                   = 0x0145
	CB_GETCOUNT              = 0x0146
	CB_GETCURSEL             = 0x0147
	CB_GETLBTEXT             = 0x0148
	CB_GETLBTEXTLEN          = 0x0149
	CB_INSERTSTRING          = 0x014A
	CB_RESETCONTENT          = 0x014B
	CB_FINDSTRING            = 0x014C
	CB_SELECTSTRING          = 0x014D
	CB_SETCURSEL             = 0x014E
	CB_SHOWDROPDOWN          = 0x014F
	CB_GETITEMDATA           = 0x0150
	CB_SETITEMDATA           = 0x0151
	CB_GETDROPPEDCONTROLRECT = 0x0152
	CB_SETITEMHEIGHT         = 0x0153
	CB_GETITEMHEIGHT         = 0x0154
	CB_SETEXTENDEDUI         = 0x0155
	CB_GETEXTENDEDUI         = 0x0156
	CB_GETDROPPEDSTATE       = 0x0157
	CB_FINDSTRINGEXACT       = 0x0158
	CB_SETLOCALE             = 345
	CB_GETLOCALE             = 346
	CB_GETTOPINDEX           = 347
	CB_SETTOPINDEX           = 348
	CB_GETHORIZONTALEXTENT   = 349
	CB_SETHORIZONTALEXTENT   = 350
	CB_GETDROPPEDWIDTH       = 351
	CB_SETDROPPEDWIDTH       = 352
	CB_INITSTORAGE           = 353
	CB_MSGMAX                = 0x165 /* if _WIN32_WINNT >= 0x0501 */
	//CB_MSGMAX                = 0x163 /* else if _WIN32_WCE >= 0x0400 */
	//CB_MSGMAX                = 0x162 /* else if _WIN32_VER >= 0x0400 */
	//CB_MSGMAX                = 0x15B /* else */
)

/* Edit Control Notification Codes */
const (
	EN_SETFOCUS  = 0x0100
	EN_KILLFOCUS = 0x0200
	EN_CHANGE    = 0x0300
	EN_UPDATE    = 0x0400
	EN_ERRSPACE  = 0x0500
	EN_MAXTEXT   = 0x0501
	EN_HSCROLL   = 0x0601
	EN_VSCROLL   = 0x0602
)

/* Edit Control Messages */
const (
	EM_GETSEL              = 0x00B0
	EM_SETSEL              = 0x00B1
	EM_GETRECT             = 0x00B2
	EM_SETRECT             = 0x00B3
	EM_SETRECTNP           = 0x00B4
	EM_SCROLL              = 0x00B5
	EM_LINESCROLL          = 0x00B6
	EM_SCROLLCARET         = 0x00B7
	EM_GETMODIFY           = 0x00B8
	EM_SETMODIFY           = 0x00B9
	EM_GETLINECOUNT        = 0x00BA
	EM_LINEINDEX           = 0x00BB
	EM_SETHANDLE           = 0x00BC
	EM_GETHANDLE           = 0x00BD
	EM_GETTHUMB            = 0x00BE
	EM_LINELENGTH          = 0x00C1
	EM_REPLACESEL          = 0x00C2
	EM_GETLINE             = 0x00C4
	EM_LIMITTEXT           = 0x00C5
	EM_CANUNDO             = 0x00C6
	EM_UNDO                = 0x00C7
	EM_FMTLINES            = 0x00C8
	EM_LINEFROMCHAR        = 0x00C9
	EM_SETTABSTOPS         = 0x00CB
	EM_SETPASSWORDCHAR     = 0x00CC
	EM_EMPTYUNDOBUFFER     = 0x00CD
	EM_GETFIRSTVISIBLELINE = 0x00CE
	EM_SETREADONLY         = 0x00CF
	EM_SETWORDBREAKPROC    = 0x00D0
	EM_GETWORDBREAKPROC    = 0x00D1
	EM_GETPASSWORDCHAR     = 0x00D2
	EM_SETMARGINS          = 211
	EM_GETMARGINS          = 212
	EM_SETLIMITTEXT        = EM_LIMITTEXT //win40 Name change
	EM_GETLIMITTEXT        = 213
	EM_POSFROMCHAR         = 214
	EM_CHARFROMPOS         = 215
	EM_SETIMESTATUS        = 216
	EM_GETIMESTATUS        = 217
)

/* Scroll bar messages */
const (
	SBM_SETPOS           = 224 /* not in win3.1  */
	SBM_GETPOS           = 225 /* not in win3.1  */
	SBM_SETRANGE         = 226 /* not in win3.1  */
	SBM_SETRANGEREDRAW   = 230 /* not in win3.1  */
	SBM_GETRANGE         = 227 /* not in win3.1  */
	SBM_ENABLE_ARROWS    = 228 /* not in win3.1  */
	SBM_SETSCROLLINFO    = 233
	SBM_GETSCROLLINFO    = 234
	SBM_GETSCROLLBARINFO = 235 /* Win XP or later */

	/* Dialog messages */

	DM_GETDEFID   = WM_USER + 0
	DM_SETDEFID   = WM_USER + 1
	DM_REPOSITION = WM_USER + 2

	PSM_PAGEINFO  = WM_USER + 100
	PSM_SHEETINFO = WM_USER + 101
)

//--------------------------- lcl 消息------------------------------
/* lcl control Value IDs */
const (
	CM_BASE = 0xB000
	//CM_CLROFFSET                   = 0x100  CRL
	CM_CLROFFSET                   = 0x0 // Only applicable in CLR
	CM_ACTIVATE                    = CM_BASE + 0
	CM_DEACTIVATE                  = CM_BASE + 1
	CM_GOTFOCUS                    = CM_BASE + 2
	CM_LOSTFOCUS                   = CM_BASE + 3
	CM_CANCELMODE                  = CM_BASE + CM_CLROFFSET + 4
	CM_DIALOGKEY                   = CM_BASE + 5
	CM_DIALOGCHAR                  = CM_BASE + 6
	CM_FOCUSCHANGED                = CM_BASE + 7
	CM_PARENTFONTCHANGED           = CM_BASE + CM_CLROFFSET + 8
	CM_PARENTCOLORCHANGED          = CM_BASE + 9
	CM_HITTEST                     = CM_BASE + 10
	CM_VISIBLECHANGED              = CM_BASE + 11
	CM_ENABLEDCHANGED              = CM_BASE + 12
	CM_COLORCHANGED                = CM_BASE + 13
	CM_FONTCHANGED                 = CM_BASE + 14
	CM_CURSORCHANGED               = CM_BASE + 15
	CM_CTL3DCHANGED                = CM_BASE + 16
	CM_PARENTCTL3DCHANGED          = CM_BASE + 17
	CM_TEXTCHANGED                 = CM_BASE + 18
	CM_MOUSEENTER                  = CM_BASE + 19
	CM_MOUSELEAVE                  = CM_BASE + 20
	CM_MENUCHANGED                 = CM_BASE + 21
	CM_APPKEYDOWN                  = CM_BASE + 22
	CM_APPSYSCOMMAND               = CM_BASE + 23
	CM_BUTTONPRESSED               = CM_BASE + 24
	CM_SHOWINGCHANGED              = CM_BASE + 25
	CM_ENTER                       = CM_BASE + 26
	CM_EXIT                        = CM_BASE + 27
	CM_DESIGNHITTEST               = CM_BASE + 28
	CM_ICONCHANGED                 = CM_BASE + 29
	CM_WANTSPECIALKEY              = CM_BASE + 30
	CM_INVOKEHELP                  = CM_BASE + 31
	CM_WINDOWHOOK                  = CM_BASE + 32
	CM_RELEASE                     = CM_BASE + 33
	CM_SHOWHINTCHANGED             = CM_BASE + 34
	CM_PARENTSHOWHINTCHANGED       = CM_BASE + 35
	CM_SYSCOLORCHANGE              = CM_BASE + 36
	CM_WININICHANGE                = CM_BASE + 37
	CM_FONTCHANGE                  = CM_BASE + 38
	CM_TIMECHANGE                  = CM_BASE + 39
	CM_TABSTOPCHANGED              = CM_BASE + 40
	CM_UIACTIVATE                  = CM_BASE + 41
	CM_UIDEACTIVATE                = CM_BASE + 42
	CM_DOCWINDOWACTIVATE           = CM_BASE + 43
	CM_CONTROLLISTCHANGE           = CM_BASE + 44
	CM_GETDATALINK                 = CM_BASE + 45
	CM_CHILDKEY                    = CM_BASE + 46
	CM_DRAG                        = CM_BASE + CM_CLROFFSET + 47
	CM_HINTSHOW                    = CM_BASE + CM_CLROFFSET + 48
	CM_DIALOGHANDLE                = CM_BASE + 49
	CM_ISTOOLCONTROL               = CM_BASE + 50
	CM_RECREATEWND                 = CM_BASE + 51
	CM_INVALIDATE                  = CM_BASE + 52
	CM_SYSFONTCHANGED              = CM_BASE + 53
	CM_CONTROLCHANGE               = CM_BASE + 54
	CM_CHANGED                     = CM_BASE + 55
	CM_DOCKCLIENT                  = CM_BASE + 56
	CM_UNDOCKCLIENT                = CM_BASE + 57
	CM_FLOAT                       = CM_BASE + 58
	CM_BORDERCHANGED               = CM_BASE + 59
	CM_BIDIMODECHANGED             = CM_BASE + 60
	CM_PARENTBIDIMODECHANGED       = CM_BASE + 61
	CM_ALLCHILDRENFLIPPED          = CM_BASE + 62
	CM_ACTIONUPDATE                = CM_BASE + 63
	CM_ACTIONEXECUTE               = CM_BASE + 64
	CM_HINTSHOWPAUSE               = CM_BASE + 65
	CM_DOCKNOTIFICATION            = CM_BASE + CM_CLROFFSET + 66
	CM_MOUSEWHEEL                  = CM_BASE + 67
	CM_ISSHORTCUT                  = CM_BASE + 68
	CM_UPDATEACTIONS               = CM_BASE + 69
	CM_RAWX11EVENT                 = CM_BASE + 69
	CM_INVALIDATEDOCKHOST          = CM_BASE + CM_CLROFFSET + 70
	CM_SETACTIVECONTROL            = CM_BASE + 71
	CM_POPUPHWNDDESTROY            = CM_BASE + 72
	CM_CREATEPOPUP                 = CM_BASE + 73
	CM_DESTROYHANDLE               = CM_BASE + 74
	CM_MOUSEACTIVATE               = CM_BASE + 75
	CM_CONTROLLISTCHANGING         = CM_BASE + 76
	CM_BUFFEREDPRINTCLIENT         = CM_BASE + 77
	CM_UNTHEMECONTROL              = CM_BASE + 78
	CM_DOUBLEBUFFEREDCHANGED       = CM_BASE + 79
	CM_PARENTDOUBLEBUFFEREDCHANGED = CM_BASE + 80
	CM_STYLECHANGED                = CM_BASE + 81
	CM_THEMECHANGED                = CM_STYLECHANGED //deprecated 'Use CM_STYLECHANGED'
	CM_GESTURE                     = CM_BASE + 82
	CM_CUSTOMGESTURESCHANGED       = CM_BASE + 83
	CM_GESTUREMANAGERCHANGED       = CM_BASE + 84
	CM_STANDARDGESTURESCHANGED     = CM_BASE + 85
	CM_INPUTLANGCHANGE             = CM_BASE + 86
	CM_TABLETOPTIONSCHANGED        = CM_BASE + 87
	CM_PARENTTABLETOPTIONSCHANGED  = CM_BASE + 88
	CM_CUSTOMSTYLECHANGED          = CM_BASE + 89
	CM_SYSFONTSALLCHANGED          = CM_BASE + 90
)

/* lcl control notification IDs */
const (
	CN_BASE              = 0xBC00
	CN_CHARTOITEM        = CN_BASE + WM_CHARTOITEM
	CN_COMMAND           = CN_BASE + WM_COMMAND
	CN_COMPAREITEM       = CN_BASE + WM_COMPAREITEM
	CN_CTLCOLORBTN       = CN_BASE + WM_CTLCOLORBTN
	CN_CTLCOLORDLG       = CN_BASE + WM_CTLCOLORDLG
	CN_CTLCOLOREDIT      = CN_BASE + WM_CTLCOLOREDIT
	CN_CTLCOLORLISTBOX   = CN_BASE + WM_CTLCOLORLISTBOX
	CN_CTLCOLORMSGBOX    = CN_BASE + WM_CTLCOLORMSGBOX
	CN_CTLCOLORSCROLLBAR = CN_BASE + WM_CTLCOLORSCROLLBAR
	CN_CTLCOLORSTATIC    = CN_BASE + WM_CTLCOLORSTATIC
	CN_DELETEITEM        = CN_BASE + WM_DELETEITEM
	CN_DRAWITEM          = CN_BASE + WM_DRAWITEM
	CN_HSCROLL           = CN_BASE + WM_HSCROLL
	CN_MEASUREITEM       = CN_BASE + WM_MEASUREITEM
	CN_PARENTNOTIFY      = CN_BASE + WM_PARENTNOTIFY
	CN_VKEYTOITEM        = CN_BASE + WM_VKEYTOITEM
	CN_VSCROLL           = CN_BASE + WM_VSCROLL
	CN_KEYDOWN           = CN_BASE + WM_KEYDOWN
	CN_KEYUP             = CN_BASE + WM_KEYUP
	CN_CHAR              = CN_BASE + WM_CHAR
	CN_SYSKEYDOWN        = CN_BASE + WM_SYSKEYDOWN
	CN_SYSCHAR           = CN_BASE + WM_SYSCHAR
	CN_NOTIFY            = CN_BASE + WM_NOTIFY
)
