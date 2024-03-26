//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build amd64 || arm64
// +build amd64 arm64

package messages

import "github.com/energye/energy/v2/types"

type TDWordFiller struct {
	Filler [4]uint8
}

type TPaint struct {
	Msg         types.Cardinal
	_UnusedMsg  types.Cardinal
	DC          types.HDC
	PaintStruct types.Paint
	Result      types.LResult
}

type TMove struct {
	Msg        types.Cardinal
	_UnusedMsg types.Cardinal
	MoveType   types.PtrInt // 0 = update, 1 = force RequestAlign, 128 = Source is Interface (Widget has moved)
	Dummy      types.LPARAM // needed for64 bit alignment
	Result     types.LResult
}

type TSize struct {
	Msg          types.Cardinal
	MsgFiller    TDWordFiller
	SizeType     types.PtrInt // see LCLType.pp (e.g. Size_Restored)
	Width        types.Word
	Height       types.Word
	LParamfiller TDWordFiller
	Result       types.LResult
}

type TWindowPosChanged struct {
	Msg        types.Cardinal
	_UnusedMsg types.Cardinal
	Unused     types.WPARAM
	WindowPos  types.WindowPos
	Result     types.LPARAM
}
