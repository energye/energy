//go:build windows
// +build windows

package initialize

import "github.com/energye/golcl/lcl/win"

func APIInit() {
	win.Init()
}
