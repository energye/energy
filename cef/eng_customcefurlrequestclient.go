//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICustomCefUrlrequestClient Parent: ICefUrlRequestClient
type ICustomCefUrlrequestClient interface {
	ICefUrlRequestClient
}

// TCustomCefUrlrequestClient Parent: TCefUrlRequestClient
type TCustomCefUrlrequestClient struct {
	TCefUrlRequestClient
}

func NewCustomCefUrlrequestClient(events ICEFUrlRequestClientEvents) ICustomCefUrlrequestClient {
	r1 := CEF().SysCallN(2132, GetObjectUintptr(events))
	return AsCustomCefUrlrequestClient(r1)
}
