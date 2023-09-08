//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux && amd64
// +build linux,amd64

package assets

import (
	"embed"
	"io/fs"
)

//go:embed upx/upx-linux-amd64
var upx embed.FS

func UpxBytes() (fs.File, error) {
	return upx.Open("upx/upx-linux-amd64")
}
