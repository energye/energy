//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

// Options App config option
type Options struct {
	Caption            string
	DefaultURL         string
	ICON               []byte
	X                  int
	Y                  int
	Width              int
	Height             int
	DisableDevTools    bool
	DisableContextMenu bool
	Frameless          bool
	LocalLoad          *LocalLoad
}
