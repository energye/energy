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

import "embed"

//go:embed assets
var assets embed.FS

// DefaultPNGICON energy app default icon.png
func DefaultPNGICON() []byte {
	if d, err := assets.ReadFile("assets/icon.png"); err == nil {
		return d
	}
	return nil
}

// DefaultICOICON energy app default icon.ico
func DefaultICOICON() []byte {
	if d, err := assets.ReadFile("assets/icon.ico"); err == nil {
		return d
	}
	return nil
}
