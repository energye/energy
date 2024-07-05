//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package assets energy internal assets
package assets

import _ "embed"

//go:embed assets/icon.png
var pngByte []byte

//go:embed assets/icon.ico
var icoByte []byte

type icon uintptr

var ICON icon

func (icon) PNG() []byte {
	return pngByte
}

func (icon) ICO() []byte {
	return icoByte
}
