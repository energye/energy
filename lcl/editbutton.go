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

// IEditButton Parent: ICustomEditButton
type IEditButton interface {
	ICustomEditButton
	AutoSelected() bool                   // property
	SetAutoSelected(AValue bool)          // property
	Button() ISpeedButton                 // property
	Edit() IEbEdit                        // property
	AutoSelect() bool                     // property
	SetAutoSelect(AValue bool)            // property
	ButtonCaption() string                // property
	SetButtonCaption(AValue string)       // property
	ButtonCursor() TCursor                // property
	SetButtonCursor(AValue TCursor)       // property
	ButtonHint() string                   // property
	SetButtonHint(AValue string)          // property
	ButtonOnlyWhenFocused() bool          // property
	SetButtonOnlyWhenFocused(AValue bool) // property
	ButtonWidth() int32                   // property
	SetButtonWidth(AValue int32)          // property
	DirectInput() bool                    // property
	SetDirectInput(AValue bool)           // property
	Flat() bool                           // property
	SetFlat(AValue bool)                  // property
	FocusOnButtonClick() bool             // property
	SetFocusOnButtonClick(AValue bool)    // property
	Glyph() IBitmap                       // property
	SetGlyph(AValue IBitmap)              // property
	Images() ICustomImageList             // property
	SetImages(AValue ICustomImageList)    // property
	ImageIndex() TImageIndex              // property
	SetImageIndex(AValue TImageIndex)     // property
	ImageWidth() int32                    // property
	SetImageWidth(AValue int32)           // property
	Layout() TLeftRight                   // property
	SetLayout(AValue TLeftRight)          // property
	NumGlyphs() int32                     // property
	SetNumGlyphs(AValue int32)            // property
	ParentFont() bool                     // property
	SetParentFont(AValue bool)            // property
	ParentShowHint() bool                 // property
	SetParentShowHint(AValue bool)        // property
	Spacing() int32                       // property
	SetSpacing(AValue int32)              // property
	SetOnButtonClick(fn TNotifyEvent)     // property event
}

// TEditButton Parent: TCustomEditButton
type TEditButton struct {
	TCustomEditButton
	buttonClickPtr uintptr
}

func NewEditButton(AOwner IComponent) IEditButton {
	r1 := LCL().SysCallN(2532, GetObjectUintptr(AOwner))
	return AsEditButton(r1)
}

func (m *TEditButton) AutoSelected() bool {
	r1 := LCL().SysCallN(2524, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetAutoSelected(AValue bool) {
	LCL().SysCallN(2524, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Button() ISpeedButton {
	r1 := LCL().SysCallN(2525, m.Instance())
	return AsSpeedButton(r1)
}

func (m *TEditButton) Edit() IEbEdit {
	r1 := LCL().SysCallN(2534, m.Instance())
	return AsEbEdit(r1)
}

func (m *TEditButton) AutoSelect() bool {
	r1 := LCL().SysCallN(2523, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetAutoSelect(AValue bool) {
	LCL().SysCallN(2523, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) ButtonCaption() string {
	r1 := LCL().SysCallN(2526, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TEditButton) SetButtonCaption(AValue string) {
	LCL().SysCallN(2526, 1, m.Instance(), PascalStr(AValue))
}

func (m *TEditButton) ButtonCursor() TCursor {
	r1 := LCL().SysCallN(2527, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TEditButton) SetButtonCursor(AValue TCursor) {
	LCL().SysCallN(2527, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) ButtonHint() string {
	r1 := LCL().SysCallN(2528, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TEditButton) SetButtonHint(AValue string) {
	LCL().SysCallN(2528, 1, m.Instance(), PascalStr(AValue))
}

func (m *TEditButton) ButtonOnlyWhenFocused() bool {
	r1 := LCL().SysCallN(2529, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetButtonOnlyWhenFocused(AValue bool) {
	LCL().SysCallN(2529, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) ButtonWidth() int32 {
	r1 := LCL().SysCallN(2530, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetButtonWidth(AValue int32) {
	LCL().SysCallN(2530, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) DirectInput() bool {
	r1 := LCL().SysCallN(2533, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetDirectInput(AValue bool) {
	LCL().SysCallN(2533, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Flat() bool {
	r1 := LCL().SysCallN(2535, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetFlat(AValue bool) {
	LCL().SysCallN(2535, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) FocusOnButtonClick() bool {
	r1 := LCL().SysCallN(2536, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetFocusOnButtonClick(AValue bool) {
	LCL().SysCallN(2536, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Glyph() IBitmap {
	r1 := LCL().SysCallN(2537, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TEditButton) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(2537, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TEditButton) Images() ICustomImageList {
	r1 := LCL().SysCallN(2540, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TEditButton) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2540, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TEditButton) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2538, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TEditButton) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2538, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) ImageWidth() int32 {
	r1 := LCL().SysCallN(2539, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetImageWidth(AValue int32) {
	LCL().SysCallN(2539, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) Layout() TLeftRight {
	r1 := LCL().SysCallN(2541, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TEditButton) SetLayout(AValue TLeftRight) {
	LCL().SysCallN(2541, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) NumGlyphs() int32 {
	r1 := LCL().SysCallN(2542, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(2542, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) ParentFont() bool {
	r1 := LCL().SysCallN(2543, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetParentFont(AValue bool) {
	LCL().SysCallN(2543, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(2544, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2544, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Spacing() int32 {
	r1 := LCL().SysCallN(2546, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetSpacing(AValue int32) {
	LCL().SysCallN(2546, 1, m.Instance(), uintptr(AValue))
}

func EditButtonClass() TClass {
	ret := LCL().SysCallN(2531)
	return TClass(ret)
}

func (m *TEditButton) SetOnButtonClick(fn TNotifyEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2545, m.Instance(), m.buttonClickPtr)
}
