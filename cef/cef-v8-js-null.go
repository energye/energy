//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

type JSNull struct {
	ICEFv8Value
}

func (m *JSNull) ToString() string {
	return "null"
}
