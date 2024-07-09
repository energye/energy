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

type Options struct {
	Caption            string
	DefaultURL         string
	ICON               []byte
	Width              int
	Height             int
	DisableDevTools    bool
	DisableContextMenu bool
	LocalLoad          *LocalLoad
}
