//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

type JSUndefined struct {
	ICEFv8Value
}

func (m *JSUndefined) ToString() string {
	return "undefined"
}
