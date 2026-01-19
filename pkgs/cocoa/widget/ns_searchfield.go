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

package widget

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "config.h"
*/
import "C"

type NSSearchField struct {
	TextField
}

func NewNSSearchField(owner *NSToolBar, config ControlTextField, property ControlProperty) *NSSearchField {
	if config.Identifier == "" {
		config.Identifier = nextSerialNumber("SearchField")
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
	cTextField := C.NewSearchField(owner.delegate, cPlaceholder, cTooltip, cProperty)
	m := &NSSearchField{}
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
