//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows && 386
// +build windows,386

package assets

import (
	"embed"
	"io/fs"
)

//go:embed upx/upx-windows-amd32.exe
var upx embed.FS

func UpxBytes() (fs.File, error) {
	return upx.Open("upx/upx-windows-amd32.exe")
}
