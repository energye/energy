//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICEFFileDialogInfo Parent: IObject
type ICEFFileDialogInfo interface {
	IObject
	Mode() uint32                              // property
	SetMode(AValue uint32)                     // property
	Title() string                             // property
	SetTitle(AValue string)                    // property
	DefaultFilePath() string                   // property
	SetDefaultFilePath(AValue string)          // property
	SetAcceptFilters(AValue IStrings)          // property
	Callback() ICefFileDialogCallback          // property
	SetCallback(AValue ICefFileDialogCallback) // property
	DialogFilter() string                      // property
	DialogType() TCEFDialogType                // property
	DefaultAudioFileDesc() string              // property
	SetDefaultAudioFileDesc(AValue string)     // property
	DefaultVideoFileDesc() string              // property
	SetDefaultVideoFileDesc(AValue string)     // property
	DefaultTextFileDesc() string               // property
	SetDefaultTextFileDesc(AValue string)      // property
	DefaultImageFileDesc() string              // property
	SetDefaultImageFileDesc(AValue string)     // property
	DefaultAllFileDesc() string                // property
	SetDefaultAllFileDesc(AValue string)       // property
	DefaultUnknownFileDesc() string            // property
	SetDefaultUnknownFileDesc(AValue string)   // property
	Clear()                                    // procedure
}

// TCEFFileDialogInfo Parent: TObject
type TCEFFileDialogInfo struct {
	TObject
}

func NewCEFFileDialogInfo() ICEFFileDialogInfo {
	r1 := CEF().SysCallN(113)
	return AsCEFFileDialogInfo(r1)
}

func (m *TCEFFileDialogInfo) Mode() uint32 {
	r1 := CEF().SysCallN(123, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFFileDialogInfo) SetMode(AValue uint32) {
	CEF().SysCallN(123, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFFileDialogInfo) Title() string {
	r1 := CEF().SysCallN(124, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetTitle(AValue string) {
	CEF().SysCallN(124, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultFilePath() string {
	r1 := CEF().SysCallN(116, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultFilePath(AValue string) {
	CEF().SysCallN(116, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) SetAcceptFilters(AValue IStrings) {
	CEF().SysCallN(109, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCEFFileDialogInfo) Callback() ICefFileDialogCallback {
	var resultCefFileDialogCallback uintptr
	CEF().SysCallN(110, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCefFileDialogCallback)))
	return AsCefFileDialogCallback(resultCefFileDialogCallback)
}

func (m *TCEFFileDialogInfo) SetCallback(AValue ICefFileDialogCallback) {
	CEF().SysCallN(110, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCEFFileDialogInfo) DialogFilter() string {
	r1 := CEF().SysCallN(121, m.Instance())
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) DialogType() TCEFDialogType {
	r1 := CEF().SysCallN(122, m.Instance())
	return TCEFDialogType(r1)
}

func (m *TCEFFileDialogInfo) DefaultAudioFileDesc() string {
	r1 := CEF().SysCallN(115, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultAudioFileDesc(AValue string) {
	CEF().SysCallN(115, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultVideoFileDesc() string {
	r1 := CEF().SysCallN(120, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultVideoFileDesc(AValue string) {
	CEF().SysCallN(120, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultTextFileDesc() string {
	r1 := CEF().SysCallN(118, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultTextFileDesc(AValue string) {
	CEF().SysCallN(118, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultImageFileDesc() string {
	r1 := CEF().SysCallN(117, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultImageFileDesc(AValue string) {
	CEF().SysCallN(117, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultAllFileDesc() string {
	r1 := CEF().SysCallN(114, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultAllFileDesc(AValue string) {
	CEF().SysCallN(114, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCEFFileDialogInfo) DefaultUnknownFileDesc() string {
	r1 := CEF().SysCallN(119, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCEFFileDialogInfo) SetDefaultUnknownFileDesc(AValue string) {
	CEF().SysCallN(119, 1, m.Instance(), PascalStr(AValue))
}

func CEFFileDialogInfoClass() TClass {
	ret := CEF().SysCallN(111)
	return TClass(ret)
}

func (m *TCEFFileDialogInfo) Clear() {
	CEF().SysCallN(112, m.Instance())
}
