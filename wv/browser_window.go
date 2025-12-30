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

import "sync/atomic"

// global browser id
var globalBrowserID uint32

// return next browser id
func getNextBrowserID() uint32 {
	atomic.AddUint32(&globalBrowserID, 1)
	return globalBrowserID
}

type IBrowserWindow interface {
}
