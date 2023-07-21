//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

type WindowPos struct {
	Hwnd            THandle
	HwndInsertAfter THandle
	X               Integer
	Y               Integer
	Cx              Integer
	Cy              Integer
	Flags           Cardinal
}
