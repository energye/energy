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

// ICustomComboBoxEx Parent: ICustomComboBox
type ICustomComboBoxEx interface {
	ICustomComboBox
	AutoCompleteOptions() TAutoCompleteOptions                                                                                         // property
	SetAutoCompleteOptions(AValue TAutoCompleteOptions)                                                                                // property
	Images() ICustomImageList                                                                                                          // property
	SetImages(AValue ICustomImageList)                                                                                                 // property
	ImagesWidth() int32                                                                                                                // property
	SetImagesWidth(AValue int32)                                                                                                       // property
	ItemsEx() IComboExItems                                                                                                            // property
	SetItemsEx(AValue IComboExItems)                                                                                                   // property
	StyleForComboBoxExStyle() TComboBoxExStyle                                                                                         // property
	SetStyleForComboBoxExStyle(AValue TComboBoxExStyle)                                                                                // property
	StyleEx() TComboBoxExStyles                                                                                                        // property
	SetStyleEx(AValue TComboBoxExStyles)                                                                                               // property
	Add() int32                                                                                                                        // function
	Add1(ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex)                 // procedure
	AssignItemsEx(AItems IStrings)                                                                                                     // procedure
	AssignItemsEx1(AItemsEx IComboExItems)                                                                                             // procedure
	Delete(AIndex int32)                                                                                                               // procedure
	DeleteSelected()                                                                                                                   // procedure
	Insert(AIndex int32, ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex) // procedure
}

// TCustomComboBoxEx Parent: TCustomComboBox
type TCustomComboBoxEx struct {
	TCustomComboBox
}

func NewCustomComboBoxEx(TheOwner IComponent) ICustomComboBoxEx {
	r1 := LCL().SysCallN(1431, GetObjectUintptr(TheOwner))
	return AsCustomComboBoxEx(r1)
}

func (m *TCustomComboBoxEx) AutoCompleteOptions() TAutoCompleteOptions {
	r1 := LCL().SysCallN(1429, 0, m.Instance(), 0)
	return TAutoCompleteOptions(r1)
}

func (m *TCustomComboBoxEx) SetAutoCompleteOptions(AValue TAutoCompleteOptions) {
	LCL().SysCallN(1429, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) Images() ICustomImageList {
	r1 := LCL().SysCallN(1434, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomComboBoxEx) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1434, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBoxEx) ImagesWidth() int32 {
	r1 := LCL().SysCallN(1435, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBoxEx) SetImagesWidth(AValue int32) {
	LCL().SysCallN(1435, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) ItemsEx() IComboExItems {
	r1 := LCL().SysCallN(1437, 0, m.Instance(), 0)
	return AsComboExItems(r1)
}

func (m *TCustomComboBoxEx) SetItemsEx(AValue IComboExItems) {
	LCL().SysCallN(1437, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBoxEx) StyleForComboBoxExStyle() TComboBoxExStyle {
	r1 := LCL().SysCallN(1439, 0, m.Instance(), 0)
	return TComboBoxExStyle(r1)
}

func (m *TCustomComboBoxEx) SetStyleForComboBoxExStyle(AValue TComboBoxExStyle) {
	LCL().SysCallN(1439, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) StyleEx() TComboBoxExStyles {
	r1 := LCL().SysCallN(1438, 0, m.Instance(), 0)
	return TComboBoxExStyles(r1)
}

func (m *TCustomComboBoxEx) SetStyleEx(AValue TComboBoxExStyles) {
	LCL().SysCallN(1438, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) Add() int32 {
	r1 := LCL().SysCallN(1425, m.Instance())
	return int32(r1)
}

func CustomComboBoxExClass() TClass {
	ret := LCL().SysCallN(1430)
	return TClass(ret)
}

func (m *TCustomComboBoxEx) Add1(ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex) {
	LCL().SysCallN(1426, m.Instance(), PascalStr(ACaption), uintptr(AIndent), uintptr(AImgIdx), uintptr(AOverlayImgIdx), uintptr(ASelectedImgIdx))
}

func (m *TCustomComboBoxEx) AssignItemsEx(AItems IStrings) {
	LCL().SysCallN(1427, m.Instance(), GetObjectUintptr(AItems))
}

func (m *TCustomComboBoxEx) AssignItemsEx1(AItemsEx IComboExItems) {
	LCL().SysCallN(1428, m.Instance(), GetObjectUintptr(AItemsEx))
}

func (m *TCustomComboBoxEx) Delete(AIndex int32) {
	LCL().SysCallN(1432, m.Instance(), uintptr(AIndex))
}

func (m *TCustomComboBoxEx) DeleteSelected() {
	LCL().SysCallN(1433, m.Instance())
}

func (m *TCustomComboBoxEx) Insert(AIndex int32, ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex) {
	LCL().SysCallN(1436, m.Instance(), uintptr(AIndex), PascalStr(ACaption), uintptr(AIndent), uintptr(AImgIdx), uintptr(AOverlayImgIdx), uintptr(ASelectedImgIdx))
}
