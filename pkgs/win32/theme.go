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
// +build windows

package win32

import (
	"github.com/energye/lcl/pkgs/win"
	"github.com/energye/lcl/rtl/version"
	"golang.org/x/sys/windows/registry"
	"unsafe"
)

func IsCurrentlyDarkMode() bool {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer key.Close()

	AppsUseLightTheme, _, err := key.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return false
	}
	return AppsUseLightTheme == 0
}

func IsWindowsVersionAtLeast(major, minor, buildNumber int) bool {
	return version.OSVersion.Major >= major && version.OSVersion.Minor >= minor && version.OSVersion.Build >= buildNumber
}

func SupportsThemes() bool {
	return IsWindowsVersionAtLeast(10, 0, 17763)
}

func SupportsImmersiveDarkMode() bool {
	return IsWindowsVersionAtLeast(10, 0, 18985)
}

func ChangeTheme(hwnd uintptr, useDarkMode bool) {
	if SupportsThemes() {
		attr := win.DwmwaUseImmersiveDarkModeBefore20h1
		if SupportsImmersiveDarkMode() {
			attr = win.DwmwaUseImmersiveDarkMode
		}
		var winDark int32
		if useDarkMode {
			winDark = 1
		}
		win.DwmSetWindowAttribute(hwnd, attr, unsafe.Pointer(&winDark), unsafe.Sizeof(winDark))
	}
}
