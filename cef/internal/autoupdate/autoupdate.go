//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package autoupdate Energy check auto update
package autoupdate

var (
	isCheckUpdate = true
)

// CheckUpdate
//	检查更新, isCheckUpdate 为true时
func CheckUpdate() {
	if isCheckUpdate {

	}
}

// IsCheckUpdate
//	设置是否检查更新
func IsCheckUpdate(v bool) {
	isCheckUpdate = v
}

func UpdateLog() []string {
	return nil
}

func updatePrompt() {

}

func updateDownload() {

}
