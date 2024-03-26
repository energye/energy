//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package lcl

import (
	"github.com/energye/energy/v2/pkgs/win"
)

// SetIconResId
//
// 从资源中设置图标的id
//
// Sets the id of the icon from the resource.
func (m *TApplication) SetIconResId(id int) {
	hIcon := win.LoadIcon(win.GetSelfModuleHandle(), id)
	if hIcon != 0 {
		m.Icon().SetHandle(hIcon)
	}
}
