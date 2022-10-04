//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

type JSUndefined struct {
	ICEFv8Value
}

func (m *JSUndefined) ToString() string {
	return "undefined"
}
