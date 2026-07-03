//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type InfoBar struct {
	Box
}

func AsInfoBar(ptr unsafe.Pointer) *InfoBar {
	if ptr == nil {
		return nil
	}
	m := new(InfoBar)
	m.instance = ptr
	return m
}

func NewInfoBar() *InfoBar {
	r := gtk3.SysCall("gtk_info_bar_new")
	if r == 0 {
		return nil
	}
	return AsInfoBar(unsafe.Pointer(r))
}

func (m *InfoBar) AddActionWidget(child IWidget, responseId int) {
	gtk3.SysCall("gtk_info_bar_add_action_widget", m.Instance(), child.Instance(), uintptr(responseId))
}

func (m *InfoBar) AddButton(buttonText string, responseId int) {
	gtk3.SysCall("gtk_info_bar_add_button", m.Instance(), CStr(buttonText), uintptr(responseId))
}

func (m *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	gtk3.SysCall("gtk_info_bar_set_response_sensitive", m.Instance(), uintptr(responseId), ToCBool(setting))
}

func (m *InfoBar) SetDefaultResponse(responseId int) {
	gtk3.SysCall("gtk_info_bar_set_default_response", m.Instance(), uintptr(responseId))
}

func (m *InfoBar) SetMessageType(messageType MessageType) {
	gtk3.SysCall("gtk_info_bar_set_message_type", m.Instance(), uintptr(messageType))
}

func (m *InfoBar) GetMessageType() MessageType {
	r := gtk3.SysCall("gtk_info_bar_get_message_type", m.Instance())
	return MessageType(r)
}

func (m *InfoBar) GetActionArea() IWidget {
	r := gtk3.SysCall("gtk_info_bar_get_action_area", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

func (m *InfoBar) GetContentArea() IBox {
	r := gtk3.SysCall("gtk_info_bar_get_content_area", m.Instance())
	if r == 0 {
		return nil
	}
	return AsBox(unsafe.Pointer(r))
}

func (m *InfoBar) SetShowCloseButton(setting bool) {
	gtk3.SysCall("gtk_info_bar_set_show_close_button", m.Instance(), ToCBool(setting))
}

func (m *InfoBar) GetShowCloseButton() bool {
	r := gtk3.SysCall("gtk_info_bar_get_show_close_button", m.Instance())
	return ToGoBool(r)
}
