//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IDirectoryEdit Parent: ICustomEditButton
type IDirectoryEdit interface {
	ICustomEditButton
	AutoSelected() bool                           // property
	SetAutoSelected(AValue bool)                  // property
	Directory() string                            // property
	SetDirectory(AValue string)                   // property
	RootDir() string                              // property
	SetRootDir(AValue string)                     // property
	DialogTitle() string                          // property
	SetDialogTitle(AValue string)                 // property
	DialogOptions() TOpenOptions                  // property
	SetDialogOptions(AValue TOpenOptions)         // property
	ShowHidden() bool                             // property
	SetShowHidden(AValue bool)                    // property
	ButtonCaption() string                        // property
	SetButtonCaption(AValue string)               // property
	ButtonCursor() TCursor                        // property
	SetButtonCursor(AValue TCursor)               // property
	ButtonHint() string                           // property
	SetButtonHint(AValue string)                  // property
	ButtonOnlyWhenFocused() bool                  // property
	SetButtonOnlyWhenFocused(AValue bool)         // property
	ButtonWidth() int32                           // property
	SetButtonWidth(AValue int32)                  // property
	DirectInput() bool                            // property
	SetDirectInput(AValue bool)                   // property
	Glyph() IBitmap                               // property
	SetGlyph(AValue IBitmap)                      // property
	NumGlyphs() int32                             // property
	SetNumGlyphs(AValue int32)                    // property
	Images() ICustomImageList                     // property
	SetImages(AValue ICustomImageList)            // property
	ImageIndex() TImageIndex                      // property
	SetImageIndex(AValue TImageIndex)             // property
	ImageWidth() int32                            // property
	SetImageWidth(AValue int32)                   // property
	Flat() bool                                   // property
	SetFlat(AValue bool)                          // property
	FocusOnButtonClick() bool                     // property
	SetFocusOnButtonClick(AValue bool)            // property
	AutoSelect() bool                             // property
	SetAutoSelect(AValue bool)                    // property
	DragCursor() TCursor                          // property
	SetDragCursor(AValue TCursor)                 // property
	DragMode() TDragMode                          // property
	SetDragMode(AValue TDragMode)                 // property
	Layout() TLeftRight                           // property
	SetLayout(AValue TLeftRight)                  // property
	ParentFont() bool                             // property
	SetParentFont(AValue bool)                    // property
	ParentShowHint() bool                         // property
	SetParentShowHint(AValue bool)                // property
	Spacing() int32                               // property
	SetSpacing(AValue int32)                      // property
	RunDialog()                                   // procedure
	SetOnAcceptDirectory(fn TAcceptFileNameEvent) // property event
	SetOnButtonClick(fn TNotifyEvent)             // property event
}

// TDirectoryEdit Parent: TCustomEditButton
type TDirectoryEdit struct {
	TCustomEditButton
	acceptDirectoryPtr uintptr
	buttonClickPtr     uintptr
}

func NewDirectoryEdit(AOwner IComponent) IDirectoryEdit {
	r1 := LCL().SysCallN(2358, GetObjectUintptr(AOwner))
	return AsDirectoryEdit(r1)
}

func (m *TDirectoryEdit) AutoSelected() bool {
	r1 := LCL().SysCallN(2351, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelected(AValue bool) {
	LCL().SysCallN(2351, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Directory() string {
	r1 := LCL().SysCallN(2362, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDirectory(AValue string) {
	LCL().SysCallN(2362, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) RootDir() string {
	r1 := LCL().SysCallN(2375, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetRootDir(AValue string) {
	LCL().SysCallN(2375, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogTitle() string {
	r1 := LCL().SysCallN(2360, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDialogTitle(AValue string) {
	LCL().SysCallN(2360, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogOptions() TOpenOptions {
	r1 := LCL().SysCallN(2359, 0, m.Instance(), 0)
	return TOpenOptions(r1)
}

func (m *TDirectoryEdit) SetDialogOptions(AValue TOpenOptions) {
	LCL().SysCallN(2359, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ShowHidden() bool {
	r1 := LCL().SysCallN(2379, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetShowHidden(AValue bool) {
	LCL().SysCallN(2379, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonCaption() string {
	r1 := LCL().SysCallN(2352, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonCaption(AValue string) {
	LCL().SysCallN(2352, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonCursor() TCursor {
	r1 := LCL().SysCallN(2353, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetButtonCursor(AValue TCursor) {
	LCL().SysCallN(2353, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ButtonHint() string {
	r1 := LCL().SysCallN(2354, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonHint(AValue string) {
	LCL().SysCallN(2354, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonOnlyWhenFocused() bool {
	r1 := LCL().SysCallN(2355, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetButtonOnlyWhenFocused(AValue bool) {
	LCL().SysCallN(2355, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonWidth() int32 {
	r1 := LCL().SysCallN(2356, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetButtonWidth(AValue int32) {
	LCL().SysCallN(2356, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DirectInput() bool {
	r1 := LCL().SysCallN(2361, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetDirectInput(AValue bool) {
	LCL().SysCallN(2361, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Glyph() IBitmap {
	r1 := LCL().SysCallN(2367, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TDirectoryEdit) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(2367, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) NumGlyphs() int32 {
	r1 := LCL().SysCallN(2372, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(2372, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Images() ICustomImageList {
	r1 := LCL().SysCallN(2370, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TDirectoryEdit) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2370, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2368, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDirectoryEdit) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2368, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ImageWidth() int32 {
	r1 := LCL().SysCallN(2369, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetImageWidth(AValue int32) {
	LCL().SysCallN(2369, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Flat() bool {
	r1 := LCL().SysCallN(2365, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFlat(AValue bool) {
	LCL().SysCallN(2365, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) FocusOnButtonClick() bool {
	r1 := LCL().SysCallN(2366, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFocusOnButtonClick(AValue bool) {
	LCL().SysCallN(2366, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(2350, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(2350, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) DragCursor() TCursor {
	r1 := LCL().SysCallN(2363, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2363, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DragMode() TDragMode {
	r1 := LCL().SysCallN(2364, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDirectoryEdit) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2364, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Layout() TLeftRight {
	r1 := LCL().SysCallN(2371, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TDirectoryEdit) SetLayout(AValue TLeftRight) {
	LCL().SysCallN(2371, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ParentFont() bool {
	r1 := LCL().SysCallN(2373, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(2373, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(2374, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2374, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Spacing() int32 {
	r1 := LCL().SysCallN(2380, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetSpacing(AValue int32) {
	LCL().SysCallN(2380, 1, m.Instance(), uintptr(AValue))
}

func DirectoryEditClass() TClass {
	ret := LCL().SysCallN(2357)
	return TClass(ret)
}

func (m *TDirectoryEdit) RunDialog() {
	LCL().SysCallN(2376, m.Instance())
}

func (m *TDirectoryEdit) SetOnAcceptDirectory(fn TAcceptFileNameEvent) {
	if m.acceptDirectoryPtr != 0 {
		RemoveEventElement(m.acceptDirectoryPtr)
	}
	m.acceptDirectoryPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2377, m.Instance(), m.acceptDirectoryPtr)
}

func (m *TDirectoryEdit) SetOnButtonClick(fn TNotifyEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2378, m.Instance(), m.buttonClickPtr)
}
