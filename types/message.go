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

type TRect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type WindowPos struct {
	Hwnd            THandle
	HwndInsertAfter THandle
	X               Integer
	Y               Integer
	Cx              Integer
	Cy              Integer
	Flags           Cardinal
}

type Paint struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     TRect
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]uint8
}
