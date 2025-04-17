//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package process Current process
package process

import "github.com/cyber-xxm/energy/v2/cef/internal/process"

// BrowserId renderer process create success
//
//	Returns the browser ID of the current process
func BrowserId() int32 {
	return process.Current.BrowserId()
}

// FrameId renderer process create success
//
//	Returns the main FrameId (channelId) of the current process
func FrameId() string {
	return process.Current.FrameId()
}
