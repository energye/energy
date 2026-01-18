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

func IsWindowsVersionLeast(major, minor, buildNumber int) bool {
	return version.OSVersion.Major >= major &&
		version.OSVersion.Minor >= minor &&
		version.OSVersion.Build >= buildNumber
}

// Windows101809 Windows >= 10 1809
func Windows101809() bool {
	return IsWindowsVersionLeast(10, 0, 17763)
}

// Windows102004 Windows >= 10 2004
func Windows102004() bool {
	return IsWindowsVersionLeast(10, 0, 18985)
}

// Windows1019041 Windows >= 10 19041
func Windows1019041() bool {
	return IsWindowsVersionLeast(10, 0, 19041)
}

// Windows1122H2 Windows >= 11 22H2
func Windows1122H2() bool {
	return IsWindowsVersionLeast(10, 0, 22621)
}

// ChangeTheme windows 10 theme
func ChangeTheme(hWnd types.HWND, useDarkMode bool) {
	if Windows101809() {
		attr := win.DwmwaUseImmersiveDarkModeBefore20h1
		if Windows102004() {
			attr = win.DwmwaUseImmersiveDarkMode
		}
		var winDark int32
		if useDarkMode {
			winDark = 1
		}
		win.DwmSetWindowAttribute(hWnd, attr, unsafe.Pointer(&winDark), unsafe.Sizeof(winDark))
	}
}
