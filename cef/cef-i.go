//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"bytes"
	"github.com/energye/golcl/lcl"
	"strings"
	"time"
	"unsafe"
)

type TCefCloseBrowsesAction CBS

type String struct {
	value string
}

type CommonInstance struct {
	lcl.IObject
	instance uintptr
	ptr      unsafe.Pointer
}

type cefV8Context struct {
	Browse uintptr
	Frame  uintptr
	Global uintptr
}

//GoEmit相关事件的接收目标
type GoEmitTarget struct {
	BrowseId int32
	FrameId  int64
}

//Type ICefCookie
type ICefCookie struct {
	Url, Name, Value, Domain, Path string
	Secure, Httponly, HasExpires   bool
	Creation, LastAccess, Expires  time.Time
	Count, Total, ID               int32
	SameSite                       TCefCookieSameSite
	Priority                       TCefCookiePriority
	SetImmediately                 bool
	DeleteCookie                   bool
	Result                         bool
}

type cefCookie struct {
	url, name, value, domain, path uintptr //string
	secure, httponly, hasExpires   uintptr //bool
	creation, lastAccess, expires  uintptr //float64
	count, total, aID              uintptr //int32
	sameSite                       uintptr //int32 TCefCookieSameSite
	priority                       uintptr //int32 TCefCookiePriority
	aSetImmediately                uintptr //bool
	aDeleteCookie                  uintptr //bool
	aResult                        uintptr //bool
}

type TCefKeyEvent struct {
	Kind                 TCefKeyEventType // called 'type' in the original CEF source code
	Modifiers            TCefEventFlags
	WindowsKeyCode       int32
	NativeKeyCode        int32
	IsSystemKey          int32
	Character            uint16
	UnmodifiedCharacter  uint16
	FocusOnEditableField int32
}

type TCefCommandLine struct {
	commandLines map[string]string
}

type tCefProxy struct {
	ProxyType              uintptr
	ProxyScheme            uintptr
	ProxyServer            uintptr
	ProxyPort              uintptr
	ProxyUsername          uintptr
	ProxyPassword          uintptr
	ProxyScriptURL         uintptr
	ProxyByPassList        uintptr
	MaxConnectionsPerProxy uintptr
	CustomHeaderName       uintptr
	CustomHeaderValue      uintptr
}

type TCefProxy struct {
	ProxyType              TCefProxyType
	ProxyScheme            TCefProxyScheme
	ProxyServer            string
	ProxyPort              int32
	ProxyUsername          string
	ProxyPassword          string
	ProxyScriptURL         string
	ProxyByPassList        string
	MaxConnectionsPerProxy int32
	CustomHeaderName       string
	CustomHeaderValue      string
}

type TCefSize struct {
	Width  int32
	Height int32
}

type TCefTouchEvent struct {
	Id            int32
	X             float32
	Y             float32
	RadiusX       float32
	RadiusY       float32
	RotationAngle float32
	Pressure      float32
	Type          TCefTouchEeventType
	Modifiers     TCefEventFlags
	PointerType   TCefPointerType
}

type TCefMouseEvent struct {
	X         int32
	Y         int32
	Modifiers TCefEventFlags
}

type BeforePopupInfo struct {
	TargetUrl         string
	TargetFrameName   string
	TargetDisposition TCefWindowOpenDisposition
	UserGesture       bool
}

type beforePopupInfo struct {
	TargetUrl         uintptr //string
	TargetFrameName   uintptr //string
	TargetDisposition uintptr //int32
	UserGesture       uintptr //bool
}

type TCefRect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type tCefRect struct {
	X      uintptr //int32
	Y      uintptr //int32
	Width  uintptr //int32
	Height uintptr //int32
}

type ICefClient struct {
	instance uintptr
	ptr      unsafe.Pointer
}

func (m *TCefCommandLine) AppendSwitch(name, value string) {
	m.commandLines[name] = value
}

func (m *TCefCommandLine) AppendArgument(argument string) {
	m.commandLines[argument] = ""
}

func (m *TCefCommandLine) toString() string {
	var str bytes.Buffer
	var i = 0
	var replace = func(s, old, new string) string {
		return strings.ReplaceAll(s, old, new)
	}
	for name, value := range m.commandLines {
		if i > 0 {
			str.WriteString(" ")
		}
		if value != "" {
			str.WriteString(replace(replace(name, " ", ""), "=", ""))
			str.WriteString("=")
			str.WriteString(replace(replace(value, " ", ""), "=", ""))
		} else {
			str.WriteString(replace(name, " ", ""))
		}
		i++
	}
	return str.String()
}

func (m *TCefKeyEvent) KeyDown() bool {
	return m.Kind == KEYEVENT_RAW_KEYDOWN || m.Kind == KEYEVENT_KEYDOWN
}

func (m *TCefKeyEvent) KeyUp() bool {
	return m.Kind == KEYEVENT_KEYUP
}

func (m *String) SetValue(v string) {
	m.value = v
}

func (m *String) GetValue() string {
	return m.value
}
