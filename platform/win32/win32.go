//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows

package win32

import (
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/types"
	"syscall"
	"unsafe"
)

var (
	user32dll = syscall.NewLazyDLL("user32.dll")
	_SetFocus = user32dll.NewProc("SetFocus")
	_GetFocus = user32dll.NewProc("GetFocus")
)

// SetWindowAlpha 设置窗口整体透明度
// hWnd：窗口句柄
// alpha：透明度值（0~255，0全透明，255不透明）
func SetWindowAlpha(hWnd types.HWND, alpha uint8) {
	exStyle := win.GetWindowLongPtr(hWnd, win.GWL_EXSTYLE)
	newExStyle := exStyle | uintptr(win.WS_EX_LAYERED)
	win.SetWindowLongPtr(hWnd, win.GWL_EXSTYLE, newExStyle)
	win.SetLayeredWindowAttributes(hWnd, 0, alpha, win.LWA_ALPHA)
}

// SetWindowColorKey 设置窗口颜色键透明
// hWnd：窗口句柄
// color：RGB颜色值（例：0x00FF00 表示绿色，需传入DWORD格式）
func SetWindowColorKey(hWnd types.HWND, color uint32) {
	exStyle := win.GetWindowLongPtr(hWnd, win.GWL_EXSTYLE)
	newExStyle := exStyle | uintptr(win.WS_EX_LAYERED)
	win.SetWindowLongPtr(hWnd, win.GWL_EXSTYLE, newExStyle)
	win.SetLayeredWindowAttributes(hWnd, color, 0, win.LWA_COLORKEY)
}

// SetWindowBlurBehind 设置窗口基础背景模糊
// hWnd：窗口句柄
// enable：true=开启模糊，false=关闭模糊
func SetWindowBlurBehind(hWnd types.HWND, enable bool) {
	blurBehind := win.DWMBlurBehind{
		DwFlags:                win.DWM_BB_ENABLE, // 启用fEnable成员
		FEnable:                enable,            // 开启/关闭模糊
		HRgnBlur:               0,                 // NULL=模糊整个客户区，可传入HRGN指定模糊区域
		FTransitionOnMaximized: false,             // 最大化时不保留过渡效果
	}
	_ = win.DwmEnableBlurBehindWindow(hWnd, blurBehind)
}

// EnableTranslucency 启用窗口透明效果
// 该函数根据Windows版本设置窗口的背景类型属性，实现透明效果
//
//	hWnd     - 窗口句柄
//	backdrop - 背景类型值，指定窗口背景的显示模式
func EnableTranslucency(hWnd types.HWND, backdrop int32) {
	if Windows1122H2() {
		win.DwmSetWindowAttribute(hWnd, win.DwmwaSystemBackdropType, unsafe.Pointer(&backdrop), unsafe.Sizeof(backdrop))
	} else {
		println("WARN: Windows < 11 22H2")
	}
}

// SetTitleBarColor 设置窗口标题栏颜色
//
//	hWnd   - 窗口句柄
//	color - 颜色值，32位整数表示的颜色值
func SetTitleBarColor(hWnd types.HWND, color int32) {
	win.DwmSetWindowAttribute(hWnd, win.DwmwaCaptionColor, unsafe.Pointer(&color), unsafe.Sizeof(color))
}

// SetTitleTextColor 设置窗口标题栏文本颜色
//
//	hWnd   - 窗口句柄
//	color - 颜色值，32位整数表示的颜色值
func SetTitleTextColor(hWnd uintptr, color int32) {
	win.DwmSetWindowAttribute(hWnd, win.DwmwaTextColor, unsafe.Pointer(&color), unsafe.Sizeof(color))
}

// SetBorderColor 设置窗口边框颜色
//
//	hWnd   - 窗口句柄
//	color - 颜色值，32位整数表示的颜色值
func SetBorderColor(hWnd uintptr, color int32) {
	win.DwmSetWindowAttribute(hWnd, win.DwmwaBorderColor, unsafe.Pointer(&color), unsafe.Sizeof(color))
}

// SetTranslucentBackground 设置窗口的半透明背景效果
// 该函数通过Windows API为指定窗口启用模糊背景效果
//
//	hWnd - 窗口句柄，类型为types.HWND，表示要设置背景效果的目标窗口
func SetTranslucentBackground(hWnd types.HWND) {
	accent := win.ACCENT_POLICY{
		AccentState: win.ACCENT_ENABLE_BLURBEHIND,
	}
	data := win.WINDOWCOMPOSITIONATTRIBDATA{}
	data.Attrib = win.WCA_ACCENT_POLICY
	data.PvData = unsafe.Pointer(&accent)
	data.CbData = unsafe.Sizeof(accent)
	win.SetWindowCompositionAttribute(hWnd, &data)
}

// SetWindowDisplayAffinity 设置窗口显示亲和性，用于控制窗口是否被屏幕捕获
// 在Windows 10 19041版本之前，如果设置为WDA_EXCLUDEFROMCAPTURE，则自动降级为WDA_MONITOR
//
//	hWnd: 窗口句柄
//	affinity: 显示亲和性标志，可选值包括WDA_MONITOR、WDA_EXCLUDEFROMCAPTURE等
func SetWindowDisplayAffinity(hWnd types.HWND, affinity uint32) bool {
	if affinity == win.WDA_EXCLUDEFROMCAPTURE && !Windows1019041() {
		affinity = win.WDA_MONITOR
	}
	return win.SetWindowDisplayAffinity(hWnd, affinity)
}

// SetBackgroundColor 设置窗口的背景颜色
//
//	hWnd: 窗口句柄，指定要设置背景颜色的窗口
//	r: 红色分量值，范围0-255
//	g: 绿色分量值，范围0-255
//	b: 蓝色分量值，范围0-255
func SetBackgroundColor(hWnd types.HWND, r, g, b uint8) bool {
	brush := win.CreateSolidBrush(r, g, b)
	ret := win.SetClassLongPtr(hWnd, win.GCLP_HBRBACKGROUND, brush)
	return ret
}

func SetFocus(hWnd types.HWND) types.HWND {
	ret, _, _ := _SetFocus.Call(
		uintptr(hWnd))
	return types.HWND(ret)
}

func GetFocus() types.HWND {
	ret, _, _ := _GetFocus.Call()
	return types.HWND(ret)
}
