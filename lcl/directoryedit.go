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
	r1 := LCL().SysCallN(2575, GetObjectUintptr(AOwner))
	return AsDirectoryEdit(r1)
}

func (m *TDirectoryEdit) AutoSelected() bool {
	r1 := LCL().SysCallN(2568, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelected(AValue bool) {
	LCL().SysCallN(2568, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Directory() string {
	r1 := LCL().SysCallN(2579, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDirectory(AValue string) {
	LCL().SysCallN(2579, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) RootDir() string {
	r1 := LCL().SysCallN(2592, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetRootDir(AValue string) {
	LCL().SysCallN(2592, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogTitle() string {
	r1 := LCL().SysCallN(2577, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDialogTitle(AValue string) {
	LCL().SysCallN(2577, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogOptions() TOpenOptions {
	r1 := LCL().SysCallN(2576, 0, m.Instance(), 0)
	return TOpenOptions(r1)
}

func (m *TDirectoryEdit) SetDialogOptions(AValue TOpenOptions) {
	LCL().SysCallN(2576, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ShowHidden() bool {
	r1 := LCL().SysCallN(2596, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetShowHidden(AValue bool) {
	LCL().SysCallN(2596, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonCaption() string {
	r1 := LCL().SysCallN(2569, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonCaption(AValue string) {
	LCL().SysCallN(2569, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonCursor() TCursor {
	r1 := LCL().SysCallN(2570, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetButtonCursor(AValue TCursor) {
	LCL().SysCallN(2570, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ButtonHint() string {
	r1 := LCL().SysCallN(2571, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonHint(AValue string) {
	LCL().SysCallN(2571, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonOnlyWhenFocused() bool {
	r1 := LCL().SysCallN(2572, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetButtonOnlyWhenFocused(AValue bool) {
	LCL().SysCallN(2572, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonWidth() int32 {
	r1 := LCL().SysCallN(2573, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetButtonWidth(AValue int32) {
	LCL().SysCallN(2573, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DirectInput() bool {
	r1 := LCL().SysCallN(2578, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetDirectInput(AValue bool) {
	LCL().SysCallN(2578, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Glyph() IBitmap {
	r1 := LCL().SysCallN(2584, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TDirectoryEdit) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(2584, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) NumGlyphs() int32 {
	r1 := LCL().SysCallN(2589, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(2589, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Images() ICustomImageList {
	r1 := LCL().SysCallN(2587, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TDirectoryEdit) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2587, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2585, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDirectoryEdit) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2585, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ImageWidth() int32 {
	r1 := LCL().SysCallN(2586, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetImageWidth(AValue int32) {
	LCL().SysCallN(2586, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Flat() bool {
	r1 := LCL().SysCallN(2582, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFlat(AValue bool) {
	LCL().SysCallN(2582, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) FocusOnButtonClick() bool {
	r1 := LCL().SysCallN(2583, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFocusOnButtonClick(AValue bool) {
	LCL().SysCallN(2583, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(2567, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(2567, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) DragCursor() TCursor {
	r1 := LCL().SysCallN(2580, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2580, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DragMode() TDragMode {
	r1 := LCL().SysCallN(2581, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDirectoryEdit) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2581, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Layout() TLeftRight {
	r1 := LCL().SysCallN(2588, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TDirectoryEdit) SetLayout(AValue TLeftRight) {
	LCL().SysCallN(2588, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ParentFont() bool {
	r1 := LCL().SysCallN(2590, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(2590, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(2591, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2591, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Spacing() int32 {
	r1 := LCL().SysCallN(2597, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetSpacing(AValue int32) {
	LCL().SysCallN(2597, 1, m.Instance(), uintptr(AValue))
}

func DirectoryEditClass() TClass {
	ret := LCL().SysCallN(2574)
	return TClass(ret)
}

func (m *TDirectoryEdit) RunDialog() {
	LCL().SysCallN(2593, m.Instance())
}

func (m *TDirectoryEdit) SetOnAcceptDirectory(fn TAcceptFileNameEvent) {
	if m.acceptDirectoryPtr != 0 {
		RemoveEventElement(m.acceptDirectoryPtr)
	}
	m.acceptDirectoryPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2594, m.Instance(), m.acceptDirectoryPtr)
}

func (m *TDirectoryEdit) SetOnButtonClick(fn TNotifyEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2595, m.Instance(), m.buttonClickPtr)
}
