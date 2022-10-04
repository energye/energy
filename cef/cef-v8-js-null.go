//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

type JSNull struct {
	ICEFv8Value
}

func (m *JSNull) ToString() string {
	return "null"
}
