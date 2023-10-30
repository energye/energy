//go:build windows

package install

import (
	"syscall"
	"unsafe"
)

func versionNumber() (majorVersion, minorVersion, buildNumber uint32) {
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procRtlGetNtVersionNumbers := ntdll.NewProc("RtlGetNtVersionNumbers")
	procRtlGetNtVersionNumbers.Call(uintptr(unsafe.Pointer(&majorVersion)), uintptr(unsafe.Pointer(&minorVersion)), uintptr(unsafe.Pointer(&buildNumber)))
	buildNumber &= 0xffff
	return
}
