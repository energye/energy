//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

// Instance 实例
func (m *ICefContextMenuParams) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefContextMenuParams) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefContextMenuParams) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// XCoord
func (m *ICefContextMenuParams) XCoord() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_XCoord).Call(m.Instance())
	return int32(r1)
}

// YCoord
func (m *ICefContextMenuParams) YCoord() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_YCoord).Call(m.Instance())
	return int32(r1)
}

// TypeFlags
func (m *ICefContextMenuParams) TypeFlags() consts.TCefContextMenuTypeFlags {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_TypeFlags).Call(m.Instance())
	return consts.TCefContextMenuTypeFlags(r1)
}

// LinkUrl
func (m *ICefContextMenuParams) LinkUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_LinkUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// UnfilteredLinkUrl
func (m *ICefContextMenuParams) UnfilteredLinkUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_UnfilteredLinkUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// SourceUrl
func (m *ICefContextMenuParams) SourceUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_SourceUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// HasImageContents
func (m *ICefContextMenuParams) HasImageContents() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_HasImageContents).Call(m.Instance())
	return api.GoBool(r1)
}

// TitleText
func (m *ICefContextMenuParams) TitleText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_TitleText).Call(m.Instance())
	return api.GoStr(r1)
}

// PageUrl
func (m *ICefContextMenuParams) PageUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_PageUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// FrameUrl
func (m *ICefContextMenuParams) FrameUrl() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_FrameUrl).Call(m.Instance())
	return api.GoStr(r1)
}

// FrameCharset
func (m *ICefContextMenuParams) FrameCharset() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_FrameCharset).Call(m.Instance())
	return api.GoStr(r1)
}

// MediaType
func (m *ICefContextMenuParams) MediaType() consts.TCefContextMenuMediaType {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_MediaType).Call(m.Instance())
	return consts.TCefContextMenuMediaType(r1)
}

// MediaStateFlags
func (m *ICefContextMenuParams) MediaStateFlags() consts.TCefContextMenuMediaStateFlags {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_MediaStateFlags).Call(m.Instance())
	return consts.TCefContextMenuMediaStateFlags(r1)
}

// SelectionText
func (m *ICefContextMenuParams) SelectionText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_SelectionText).Call(m.Instance())
	return api.GoStr(r1)
}

// MisspelledWord
func (m *ICefContextMenuParams) MisspelledWord() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_MisspelledWord).Call(m.Instance())
	return api.GoStr(r1)
}

// DictionarySuggestions
func (m *ICefContextMenuParams) DictionarySuggestions(suggestions []string) bool {
	if !m.IsValid() {
		return false
	}
	slist := lcl.NewStringList()
	for _, s := range suggestions {
		slist.Add(s)
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_DictionarySuggestions).Call(m.Instance(), slist.Instance())
	slist.Free()
	return api.GoBool(r1)
}

// IsEditable
func (m *ICefContextMenuParams) IsEditable() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_IsEditable).Call(m.Instance())
	return api.GoBool(r1)
}

// IsSpellCheckEnabled
func (m *ICefContextMenuParams) IsSpellCheckEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_IsSpellCheckEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

// EditStateFlags
func (m *ICefContextMenuParams) EditStateFlags() consts.TCefContextMenuEditStateFlags {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_EditStateFlags).Call(m.Instance())
	return consts.TCefContextMenuEditStateFlags(r1)
}

// IsCustomMenu
func (m *ICefContextMenuParams) IsCustomMenu() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_ContextMenuParams_IsCustomMenu).Call(m.Instance())
	return api.GoBool(r1)
}
