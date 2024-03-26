//go:build !windows
// +build !windows

package thread

import (
	"syscall"
)

func GetCurrentThreadId() int {
	return syscall.Gettid()
}
