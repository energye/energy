//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package toolbar

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "config.h"

*/
import "C"

type TextField struct {
	Control
	config ControlTextField
}

type NSTextField struct {
	TextField
}

func NewNSTextField(owner *NSToolBar, config ControlTextField, property ControlProperty) *NSTextField {
	if config.Identifier == "" {
		config.Identifier = nextSerialNumber("TextField")
	}
	var cPlaceholder *C.char
	cPlaceholder = C.CString(config.Placeholder)
	defer C.free(Pointer(cPlaceholder))
	var cTooltip *C.char
	if config.Tips != "" {
		cTooltip = C.CString(config.Tips)
		defer C.free(Pointer(cTooltip))
	}
	cProperty := property.ToOC()
	cTextField := C.NewTextField(owner.delegate, cPlaceholder, cTooltip, cProperty)
	m := &NSTextField{}
	m.config = config
	m.Control = Control{
		instance: Pointer(cTextField),
		owner:    owner,
		property: &property,
		item:     config.ItemBase,
	}
	m.SetBindControlObjectIdentifier()
	return m
}

func (m *TextField) SetOnChange(fn TextEvent) {
	RegisterEvent(m.config.Identifier, MakeTextChangeEvent(fn))
}

func (m *TextField) SetOnCommit(fn TextEvent) {
	RegisterEvent(m.config.Identifier, MakeTextCommitEvent(fn))
}

func (m *TextField) Text() string {
	cText := C.GetTextFieldText(m.instance)
	return C.GoString(cText)
}

func (m *TextField) SetTextFieldCursorPosition(index int) {
	C.SetTextFieldCursorPosition(m.instance, C.int(index))
}

// SetText 设置搜索框文本
func (m *TextField) SetText(text string) {
	cText := C.CString(text)
	defer C.free(Pointer(cText))
	C.SetTextFieldText(m.instance, cText)
}

func (m *TextField) UpdateTextFieldWidth(width int) {
	cWidth := C.CGFloat(width)
	C.UpdateTextFieldWidth(m.instance, cWidth)
}
