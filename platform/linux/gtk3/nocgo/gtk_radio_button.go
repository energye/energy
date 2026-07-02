//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
// Licensed under Apache License Version 2.0, January 2004
// https://www.apache.org/licenses/LICENSE-2.0
//----------------------------------------

package nocgo

import "unsafe"

// RadioButton is a representation of GTK's GtkRadioButton.
type RadioButton struct {
	CheckButton
}

func AsRadioButton(ptr unsafe.Pointer) *RadioButton {
	if ptr == nil {
		return nil
	}
	m := new(RadioButton)
	m.instance = ptr
	return m
}

// NewRadioButtonWithLabelFromWidget is a wrapper around gtk_radio_button_new_with_label_from_widget().
func NewRadioButtonWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	var member uintptr
	if radioGroupMember != nil {
		member = radioGroupMember.Instance()
	}
	r := gtk3.SysCall("gtk_radio_button_new_with_label_from_widget", member, CStr(label))
	if r == 0 {
		return nil
	}
	return AsRadioButton(unsafe.Pointer(r))
}
