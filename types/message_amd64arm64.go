//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build amd64 || arm64 || loong64
// +build amd64 arm64 loong64

package types

type TDWordFiller struct {
	Filler [4]uint8
}

type TPaint struct {
	Msg         Cardinal
	_UnusedMsg  Cardinal
	DC          HDC
	PaintStruct Paint
	Result      LResult
}

type TMove struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	MoveType   PtrInt // 0 = update, 1 = force RequestAlign, 128 = Source is Interface (Widget has moved)
	Dummy      LPARAM // needed for64 bit alignment
	Result     LResult
}

type TSize struct {
	Msg          Cardinal
	MsgFiller    TDWordFiller
	SizeType     PtrInt // see LCLType.pp (e.g. Size_Restored)
	Width        Word
	Height       Word
	LParamfiller TDWordFiller
	Result       LResult
}

type TWindowPosChanged struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	Unused     WPARAM
	WindowPos  WindowPos
	Result     LPARAM
}
