//go:build windows
// +build windows

//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package inits

import "github.com/energye/energy/v2/pkgs/win"

func winInit() {
	//win
	win.Init()
}
