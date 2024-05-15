//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin || (windows && arm64)
// +build darwin windows,arm64

package assets

import (
	"errors"
	"io/fs"
)

func UpxBytes() (fs.File, error) {
	return nil, errors.New("not support")
}
