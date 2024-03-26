//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build amd64 || arm64
// +build amd64 arm64

package types

type TDWordFiller struct {
	Filler [4]uint8
}

// TWMKey
type TWMKey struct {
	Msg       uint32
	MsgFiller TDWordFiller
	CharCode  [2]uint16
	// CharCode: Word;
	// Unused: Word;
	CharCodeUnusedFiller TDWordFiller
	KeyData              uint32
	KeyDataFiller        TDWordFiller
	Result               uintptr
}

// TLMKey
type TLMKey struct {
	Msg       uint32
	UnusedMsg Cardinal
	CharCode  [2]uint16
	Unused    uint16
	Unused2   int32
	KeyData   uint32
	Result    uintptr
}
