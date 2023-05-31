//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package process

import "github.com/energye/energy/v2/cef/internal/process"

// BrowserId
//  renderer process create success
//	Returns the browser ID of the current process
func BrowserId() int32 {
	return process.Current.BrowserId()
}

// FrameId
//  renderer process create success
//	Returns the main FrameId (channelId) of the current process
func FrameId() int64 {
	return process.Current.FrameId()
}
