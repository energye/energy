//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue JSNull 实现
package cef

type JSNull struct {
	ICEFv8Value
}

func (m *JSNull) ToString() string {
	return "null"
}
