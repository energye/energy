//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !tempdll
// +build !tempdll

// 编译命令: go build

package tempdll

func CheckAndReleaseDLL() (string, bool) {
	return "", false
}
