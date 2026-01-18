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
	"github.com/energye/lcl/rtl/version"
	"github.com/energye/lcl/types"
	"unsafe"
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

// SetWindowAcrylic 设置窗口亚克力效果
// hWnd：LCL窗口句柄
// enable：true=开启亚克力，false=关闭亚克力（恢复系统默认背景）
func SetWindowAcrylic(hWnd types.HWND, enable bool) {
	var backdropType uint32
	if enable {
		backdropType = win.DWMSBT_ACRYLIC // 启用亚克力效果
	} else {
		backdropType = 0 // 恢复系统默认背景类型
	}

	// 调用DWM API设置亚克力效果
	win.DwmSetWindowAttribute(
		hWnd,
		win.DWMWA_SYSTEMBACKDROP_TYPE, // 通用系统背景属性
		unsafe.Pointer(&backdropType),
		unsafe.Sizeof(backdropType),
	)
}

// SetWindowMica ：设置窗口云母效果（独立效果，仅Windows 11+）
// hWnd：LCL窗口句柄
// enable：true=开启云母，false=关闭云母
// isMainWindow：true=主窗口风格（DWMSBT_MAINWINDOW），false=临时窗口风格（DWMSBT_TRANSIENTWINDOW）
func SetWindowMica(hWnd types.HWND, enable bool, isMainWindow bool) {
	var backdropType uint32
	if enable {
		if isMainWindow {
			backdropType = win.DWMSBT_MAINWINDOW // 云母（主窗口，系统默认）
		} else {
			backdropType = win.DWMSBT_TRANSIENTWINDOW // 云母（临时窗口，更浅的哑光）
		}
	} else {
		backdropType = 0 // 恢复系统默认背景类型
	}

	// 调用DWM API设置云母效果（两种调用方式兼容Win11不同版本）
	// 方式1：通用系统背景属性（推荐）
	win.DwmSetWindowAttribute(hWnd, win.DWMWA_SYSTEMBACKDROP_TYPE, unsafe.Pointer(&backdropType), unsafe.Sizeof(backdropType))

	// 方式2：Win11专属云母属性（兜底，兼容早期Win11版本）
	if enable {
		micaEffect := uint32(1) // 1=启用云母，0=禁用
		win.DwmSetWindowAttribute(hWnd, win.DWMWA_MICA_EFFECT, unsafe.Pointer(&micaEffect), unsafe.Sizeof(micaEffect))
	}
}

func EnableTranslucency(hWnd types.HWND, backdrop int32) {
	if SupportsBackdropTypes() {
		win.DwmSetWindowAttribute(hWnd, win.DwmwaSystemBackdropType, unsafe.Pointer(&backdrop), unsafe.Sizeof(backdrop))
	} else {
		println("Warning: Translucency type unavailable on Windows < 22621")
	}
}

func SetTitleBarColour(hWnd types.HWND, titleBarColour int32) {
	win.DwmSetWindowAttribute(hWnd, win.DwmwaCaptionColor, unsafe.Pointer(&titleBarColour), unsafe.Sizeof(titleBarColour))
}

func SetTitleTextColour(hWnd uintptr, titleTextColour int32) {
	win.DwmSetWindowAttribute(hWnd, win.DwmwaTextColor, unsafe.Pointer(&titleTextColour), unsafe.Sizeof(titleTextColour))
}

func SetBorderColour(hWnd uintptr, titleBorderColour int32) {
	win.DwmSetWindowAttribute(hWnd, win.DwmwaBorderColor, unsafe.Pointer(&titleBorderColour), unsafe.Sizeof(titleBorderColour))
}

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

// SetWindowExtendedStyle ：单独配置窗口扩展样式（添加/移除）
// hWnd：窗口句柄
func SetWindowExtendedStyle(hWnd types.HWND, style uintptr, enable bool) {
	exStyle := win.GetWindowLongPtr(hWnd, win.GWL_EXSTYLE)
	var newExStyle uintptr
	if enable {
		newExStyle = exStyle | style
	} else {
		newExStyle = exStyle &^ style
	}
	win.SetWindowLongPtr(hWnd, win.GWL_EXSTYLE, newExStyle)
}

// ConfigureWindowDefaultExStyles 设置默认扩展样式
// hWnd：窗口句柄
func ConfigureWindowDefaultExStyles(hWnd types.HWND) {
	SetWindowExtendedStyle(hWnd, uintptr(win.WS_EX_CONTROLPARENT), true)
	SetWindowExtendedStyle(hWnd, uintptr(win.WS_EX_APPWINDOW), true)
	SetWindowExtendedStyle(hWnd, uintptr(win.WS_EX_NOREDIRECTIONBITMAP), true)
}

func isWindowsVersionAtLeast(major, minor, build int) bool {
	windowsVersion := version.OSVersion
	if windowsVersion.Major > major {
		return true
	}
	if windowsVersion.Major < major {
		return false
	}
	if windowsVersion.Minor > minor {
		return true
	}
	if windowsVersion.Minor < minor {
		return false
	}
	return windowsVersion.Build >= build
}

func SetWindowDisplayAffinity(hWnd types.HWND, affinity uint32) bool {
	if affinity == win.WDA_EXCLUDEFROMCAPTURE && !isWindowsVersionAtLeast(10, 0, 19041) {
		// for older windows versions, use WDA_MONITOR
		affinity = win.WDA_MONITOR
	}
	return win.SetWindowDisplayAffinity(hWnd, affinity)
}

func SetBackgroundColour(hWnd types.HWND, r, g, b uint8) bool {
	sbrush := win.CreateSolidBrush(r, g, b)
	ret := win.SetClassLongPtr(hWnd, win.GCLP_HBRBACKGROUND, sbrush)
	return ret
}
