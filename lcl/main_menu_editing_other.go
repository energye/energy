//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !linux

package lcl

import "github.com/energye/lcl/lcl"

func isWebview(control lcl.IWinControl) IMenuEditing {
	if control == nil {
		return nil
	}
}
