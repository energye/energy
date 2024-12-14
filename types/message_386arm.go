//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build 386 || arm
// +build 386 arm

package types

type TPaint struct {
	Msg         Cardinal
	DC          HDC
	PaintStruct Paint
	Result      LResult
}

type TMove struct {
	Msg      Cardinal
	MoveType PtrInt // 0 = update, 1 = force RequestAlign, 128 = Source is Interface (Widget has moved)
	Dummy    LPARAM // needed for64 bit alignment
	Result   LResult
}

type TSize struct {
	Msg      Cardinal
	SizeType PtrInt // see LCLType.pp (e.g. Size_Restored)
	Width    Word
	Height   Word
	Result   LResult
}

type TWindowPosChanged struct {
	Msg       Cardinal
	Unused    WPARAM
	WindowPos WindowPos
	Result    LPARAM
}
