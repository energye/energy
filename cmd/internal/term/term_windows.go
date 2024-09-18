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

package term

import (
	"github.com/pterm/pterm"
	"syscall"
	"unsafe"
)

func init() {
	// < windows 10 禁用颜色
	if osVersion() < 10 {
		pterm.DisableColor()
		IsWindows10 = false
	}
}

func osVersion() int {
	type TOSVersion struct {
		Name  string
		Build int
		Major int
		Minor int
	}
	type TOSVersionInfoEx struct {
		OSVersionInfoSize uint32
		MajorVersion      uint32
		MinorVersion      uint32
		BuildNumber       uint32
		PlatformId        uint32
		CSDVersion        [128]uint16 // Maintenance UnicodeString for PSS usage
		ServicePackMajor  uint16
		ServicePackMinor  uint16
		SuiteMask         uint16
		ProductType       uint8
		Reserved          uint8
	}
	var kernel32dll = syscall.NewLazyDLL("kernel32.dll")
	_GetVersionEx := kernel32dll.NewProc("GetVersionExW")
	var GetVersionEx = func(lpVersionInformation *TOSVersionInfoEx) bool {
		if lpVersionInformation != nil {
			lpVersionInformation.OSVersionInfoSize = uint32(unsafe.Sizeof(*lpVersionInformation))
		}
		r, _, _ := _GetVersionEx.Call(uintptr(unsafe.Pointer(lpVersionInformation)))
		return r != 0
	}

	var GetProductVersion = func(AFileName string, AMajor, AMinor, ABuild *uint32) bool {
		var (
			versiondll              = syscall.NewLazyDLL("version.dll")
			_GetFileVersionInfoSize = versiondll.NewProc("GetFileVersionInfoSizeW")
			_GetFileVersionInfo     = versiondll.NewProc("GetFileVersionInfoW")
			_VerQueryValue          = versiondll.NewProc("VerQueryValueW")
		)
		var CStr = func(str string) uintptr {
			return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
		}
		// GetFileVersionInfoSize
		var GetFileVersionInfoSize = func(lptstrFileName string, lpdwhandle *uint32) uint32 {
			r, _, _ := _GetFileVersionInfoSize.Call(CStr(lptstrFileName),
				uintptr(unsafe.Pointer(lpdwhandle)))
			return uint32(r)
		}

		// GetFileVersionInfo
		var GetFileVersionInfo = func(lptstrFilename string, dwHandle, dwLen uint32, lpData uintptr) bool {
			r, _, _ := _GetFileVersionInfo.Call(CStr(lptstrFilename),
				uintptr(dwHandle), uintptr(dwLen), lpData)
			return r != 0
		}

		// VerQueryValue
		var VerQueryValue = func(pBlock uintptr, lpSubBlock string, lplpBuffer *uintptr, puLen *uint32) bool {
			r, _, _ := _VerQueryValue.Call(uintptr(pBlock),
				uintptr(unsafe.Pointer(CStr(lpSubBlock))),
				uintptr(unsafe.Pointer(lplpBuffer)), uintptr(unsafe.Pointer(puLen)))
			return r != 0
		}

		type TVSFixedFileInfo struct {
			Signature        uint32 // e.g. $feef04bd
			StrucVersion     uint32 // e.g. $00000042 = "0.42"
			FileVersionMS    uint32 // e.g. $00030075 = "3.75"
			FileVersionLS    uint32 // e.g. $00000031 = "0.31"
			ProductVersionMS uint32 // e.g. $00030010 = "3.10"
			ProductVersionLS uint32 // e.g. $00000031 = "0.31"
			FileFlagsMask    uint32 // = $3F for version "0.42"
			FileFlags        uint32 // e.g. VFF_DEBUG | VFF_PRERELEASE
			FileOS           uint32 // e.g. VOS_DOS_WINDOWS16
			FileType         uint32 // e.g. VFT_DRIVER
			FileSubtype      uint32 // e.g. VFT2_DRV_KEYBOARD
			FileDateMS       uint32 // e.g. 0
			FileDateLS       uint32 // e.g. 0
		}

		var wnd uint32
		infoSize := GetFileVersionInfoSize(AFileName, &wnd)
		if infoSize != 0 {
			verBuf := make([]byte, infoSize)
			bufPtr := uintptr(unsafe.Pointer(&verBuf[0]))
			if GetFileVersionInfo(AFileName, wnd, infoSize, bufPtr) {
				var verSize uint32
				var fI *TVSFixedFileInfo
				if VerQueryValue(bufPtr, "\\", (*uintptr)(unsafe.Pointer(&fI)), &verSize) {
					*AMajor = fI.ProductVersionMS >> 16
					*AMinor = uint32(uint16(fI.ProductVersionMS))
					*ABuild = fI.ProductVersionLS >> 16
					return true
				}
			}
		}
		return false
	}
	var GetNetWkstaMajorMinor = func(MajorVersion, MinorVersion *uint32) bool {
		type WKSTA_INFO_100 struct {
			Wki100_platform_id  uint32
			Wki100_computername uintptr // LPWSTR
			Wki100_langroup     uintptr // LPWSTR
			Wki100_ver_major    uint32
			Wki100_ver_minor    uint32
		}
		const (
			netapi       = "netapi32.dll"
			NERR_Success = 0
		)
		var (
			netapidll           = syscall.NewLazyDLL(netapi)
			_NetWkstaGetInfo100 = netapidll.NewProc("NetWkstaGetInfo")
			_NetApiBufferFree   = netapidll.NewProc("NetApiBufferFree")
		)

		var NetWkstaGetInfo100 = func(ServerName string, Level uint32, BufPtr **WKSTA_INFO_100) uint32 {
			r, _, _ := _NetWkstaGetInfo100.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ServerName))), uintptr(Level),
				uintptr(unsafe.Pointer(BufPtr)))
			return uint32(r)
		}

		var NetApiBufferFree = func(BufPtr *WKSTA_INFO_100) uint32 {
			r, _, _ := _NetApiBufferFree.Call(uintptr(unsafe.Pointer(BufPtr)))
			return uint32(r)
		}

		var LBuf *WKSTA_INFO_100
		result := NetWkstaGetInfo100("", 100, &LBuf) == NERR_Success
		if result {
			*MajorVersion = LBuf.Wki100_ver_major
			*MinorVersion = LBuf.Wki100_ver_minor
		} else {
			*MajorVersion = 0
			*MinorVersion = 0
		}
		if LBuf != nil {
			NetApiBufferFree(LBuf)
		}
		return result
	}
	var (
		version TOSVersion
		verInfo TOSVersionInfoEx
	)
	GetVersionEx(&verInfo)
	version.Major = int(verInfo.MajorVersion)
	version.Minor = int(verInfo.MinorVersion)
	version.Build = int(verInfo.BuildNumber)

	var majorNum, minorNum, buildNum uint32
	if version.Major > 6 || (version.Major == 6 && version.Minor > 1) {
		if GetProductVersion("kernel32.dll", &majorNum, &minorNum, &buildNum) {
			version.Major = int(majorNum)
			version.Minor = int(minorNum)
			version.Build = int(buildNum)
		} else if GetNetWkstaMajorMinor(&majorNum, &minorNum) {
			version.Major = int(majorNum)
			version.Minor = int(minorNum)
		}
	}
	return int(version.Major)
}
