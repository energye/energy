//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package lcl

import (
	"fmt"
)

func showError(err interface{}) {
	fmt.Println(err)
}

func tryLoadAppIcon() {
	// no code
}
