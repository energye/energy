// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build darwin && !prod
// +build darwin,!prod

package macapp

// MacApp
//
//  1. macos调式时临时创建一个符合macapp的程序包
//
//  2. 如果基于cef，需要指定cef frameworks 根目录【/homt/xxx/cef_binary_xxxxxxx_macosx64/Release】
var MacApp = &macApp{}

type macApp struct {
	execName             string
	execFile             string
	macContentsDir       string
	macOSDir             string
	macResources         string
	lclLibFileName       string
	plistFileName        string
	pkgInfoFileName      string
	macAppFrameworksDir  string
	isLinked             bool
	isMain               bool
	cefFrameworksDir     string
	browseSubprocessPath string
	lsUIElement          string
}
