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
	"unsafe"
)

// IValueListEditor Parent: ICustomStringGrid
type IValueListEditor interface {
	ICustomStringGrid
	Modified() bool                                      // property
	SetModified(AValue bool)                             // property
	Keys(Index int32) string                             // property
	SetKeys(Index int32, AValue string)                  // property
	Values(Key string) string                            // property
	SetValues(Key string, AValue string)                 // property
	ItemProps(AKeyOrIndex uintptr) IItemProp             // property
	SetItemProps(AKeyOrIndex uintptr, AValue IItemProp)  // property
	AlternateColor() TColor                              // property
	SetAlternateColor(AValue TColor)                     // property
	AutoEdit() bool                                      // property
	SetAutoEdit(AValue bool)                             // property
	DragCursor() TCursor                                 // property
	SetDragCursor(AValue TCursor)                        // property
	DragKind() TDragKind                                 // property
	SetDragKind(AValue TDragKind)                        // property
	DragMode() TDragMode                                 // property
	SetDragMode(AValue TDragMode)                        // property
	HeaderHotZones() TGridZoneSet                        // property
	SetHeaderHotZones(AValue TGridZoneSet)               // property
	HeaderPushZones() TGridZoneSet                       // property
	SetHeaderPushZones(AValue TGridZoneSet)              // property
	MouseWheelOption() TMouseWheelOption                 // property
	SetMouseWheelOption(AValue TMouseWheelOption)        // property
	ParentColor() bool                                   // property
	SetParentColor(AValue bool)                          // property
	ParentFont() bool                                    // property
	SetParentFont(AValue bool)                           // property
	TitleFont() IFont                                    // property
	SetTitleFont(AValue IFont)                           // property
	TitleImageList() IImageList                          // property
	SetTitleImageList(AValue IImageList)                 // property
	TitleStyle() TTitleStyle                             // property
	SetTitleStyle(AValue TTitleStyle)                    // property
	DisplayOptions() TDisplayOptions                     // property
	SetDisplayOptions(AValue TDisplayOptions)            // property
	DropDownRows() int32                                 // property
	SetDropDownRows(AValue int32)                        // property
	KeyOptions() TKeyOptions                             // property
	SetKeyOptions(AValue TKeyOptions)                    // property
	Strings() IValueListStrings                          // property
	SetStrings(AValue IValueListStrings)                 // property
	TitleCaptions() IStrings                             // property
	SetTitleCaptions(AValue IStrings)                    // property
	FindRow(KeyName string, OutRow *int32) bool          // function
	InsertRow(KeyName, Value string, Append bool) int32  // function
	IsEmptyRow() bool                                    // function
	IsEmptyRow1(aRow int32) bool                         // function
	RestoreCurrentRow() bool                             // function
	Sort(Index, IndxFrom, IndxTo int32)                  // procedure
	Sort1(ACol TVleSortCol)                              // procedure
	SetOnCheckboxToggled(fn TToggledCheckboxEvent)       // property event
	SetOnEditingDone(fn TNotifyEvent)                    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) // property event
	SetOnGetPickList(fn TGetPickListEvent)               // property event
	SetOnStringsChange(fn TNotifyEvent)                  // property event
	SetOnStringsChanging(fn TNotifyEvent)                // property event
	SetOnValidate(fn TOnValidateEvent)                   // property event
}

// TValueListEditor Parent: TCustomStringGrid
type TValueListEditor struct {
	TCustomStringGrid
	checkboxToggledPtr    uintptr
	editingDonePtr        uintptr
	mouseWheelHorzPtr     uintptr
	mouseWheelLeftPtr     uintptr
	mouseWheelRightPtr    uintptr
	userCheckboxBitmapPtr uintptr
	getPickListPtr        uintptr
	stringsChangePtr      uintptr
	stringsChangingPtr    uintptr
	validatePtr           uintptr
}

func NewValueListEditor(AOwner IComponent) IValueListEditor {
	r1 := LCL().SysCallN(5156, GetObjectUintptr(AOwner))
	return AsValueListEditor(r1)
}

func (m *TValueListEditor) Modified() bool {
	r1 := LCL().SysCallN(5171, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetModified(AValue bool) {
	LCL().SysCallN(5171, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) Keys(Index int32) string {
	r1 := LCL().SysCallN(5170, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TValueListEditor) SetKeys(Index int32, AValue string) {
	LCL().SysCallN(5170, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TValueListEditor) Values(Key string) string {
	r1 := LCL().SysCallN(5193, 0, m.Instance(), PascalStr(Key))
	return GoStr(r1)
}

func (m *TValueListEditor) SetValues(Key string, AValue string) {
	LCL().SysCallN(5193, 1, m.Instance(), PascalStr(Key), PascalStr(AValue))
}

func (m *TValueListEditor) ItemProps(AKeyOrIndex uintptr) IItemProp {
	r1 := LCL().SysCallN(5168, 0, m.Instance(), uintptr(AKeyOrIndex))
	return AsItemProp(r1)
}

func (m *TValueListEditor) SetItemProps(AKeyOrIndex uintptr, AValue IItemProp) {
	LCL().SysCallN(5168, 1, m.Instance(), uintptr(AKeyOrIndex), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) AlternateColor() TColor {
	r1 := LCL().SysCallN(5153, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TValueListEditor) SetAlternateColor(AValue TColor) {
	LCL().SysCallN(5153, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) AutoEdit() bool {
	r1 := LCL().SysCallN(5154, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetAutoEdit(AValue bool) {
	LCL().SysCallN(5154, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) DragCursor() TCursor {
	r1 := LCL().SysCallN(5158, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TValueListEditor) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5158, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DragKind() TDragKind {
	r1 := LCL().SysCallN(5159, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TValueListEditor) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5159, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DragMode() TDragMode {
	r1 := LCL().SysCallN(5160, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TValueListEditor) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5160, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) HeaderHotZones() TGridZoneSet {
	r1 := LCL().SysCallN(5163, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TValueListEditor) SetHeaderHotZones(AValue TGridZoneSet) {
	LCL().SysCallN(5163, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) HeaderPushZones() TGridZoneSet {
	r1 := LCL().SysCallN(5164, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TValueListEditor) SetHeaderPushZones(AValue TGridZoneSet) {
	LCL().SysCallN(5164, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) MouseWheelOption() TMouseWheelOption {
	r1 := LCL().SysCallN(5172, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TValueListEditor) SetMouseWheelOption(AValue TMouseWheelOption) {
	LCL().SysCallN(5172, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) ParentColor() bool {
	r1 := LCL().SysCallN(5173, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetParentColor(AValue bool) {
	LCL().SysCallN(5173, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) ParentFont() bool {
	r1 := LCL().SysCallN(5174, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetParentFont(AValue bool) {
	LCL().SysCallN(5174, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) TitleFont() IFont {
	r1 := LCL().SysCallN(5190, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TValueListEditor) SetTitleFont(AValue IFont) {
	LCL().SysCallN(5190, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleImageList() IImageList {
	r1 := LCL().SysCallN(5191, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TValueListEditor) SetTitleImageList(AValue IImageList) {
	LCL().SysCallN(5191, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleStyle() TTitleStyle {
	r1 := LCL().SysCallN(5192, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TValueListEditor) SetTitleStyle(AValue TTitleStyle) {
	LCL().SysCallN(5192, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DisplayOptions() TDisplayOptions {
	r1 := LCL().SysCallN(5157, 0, m.Instance(), 0)
	return TDisplayOptions(r1)
}

func (m *TValueListEditor) SetDisplayOptions(AValue TDisplayOptions) {
	LCL().SysCallN(5157, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DropDownRows() int32 {
	r1 := LCL().SysCallN(5161, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TValueListEditor) SetDropDownRows(AValue int32) {
	LCL().SysCallN(5161, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) KeyOptions() TKeyOptions {
	r1 := LCL().SysCallN(5169, 0, m.Instance(), 0)
	return TKeyOptions(r1)
}

func (m *TValueListEditor) SetKeyOptions(AValue TKeyOptions) {
	LCL().SysCallN(5169, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) Strings() IValueListStrings {
	r1 := LCL().SysCallN(5188, 0, m.Instance(), 0)
	return AsValueListStrings(r1)
}

func (m *TValueListEditor) SetStrings(AValue IValueListStrings) {
	LCL().SysCallN(5188, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleCaptions() IStrings {
	r1 := LCL().SysCallN(5189, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TValueListEditor) SetTitleCaptions(AValue IStrings) {
	LCL().SysCallN(5189, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) FindRow(KeyName string, OutRow *int32) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(5162, m.Instance(), PascalStr(KeyName), uintptr(unsafe.Pointer(&result1)))
	*OutRow = int32(result1)
	return GoBool(r1)
}

func (m *TValueListEditor) InsertRow(KeyName, Value string, Append bool) int32 {
	r1 := LCL().SysCallN(5165, m.Instance(), PascalStr(KeyName), PascalStr(Value), PascalBool(Append))
	return int32(r1)
}

func (m *TValueListEditor) IsEmptyRow() bool {
	r1 := LCL().SysCallN(5166, m.Instance())
	return GoBool(r1)
}

func (m *TValueListEditor) IsEmptyRow1(aRow int32) bool {
	r1 := LCL().SysCallN(5167, m.Instance(), uintptr(aRow))
	return GoBool(r1)
}

func (m *TValueListEditor) RestoreCurrentRow() bool {
	r1 := LCL().SysCallN(5175, m.Instance())
	return GoBool(r1)
}

func ValueListEditorClass() TClass {
	ret := LCL().SysCallN(5155)
	return TClass(ret)
}

func (m *TValueListEditor) Sort(Index, IndxFrom, IndxTo int32) {
	LCL().SysCallN(5186, m.Instance(), uintptr(Index), uintptr(IndxFrom), uintptr(IndxTo))
}

func (m *TValueListEditor) Sort1(ACol TVleSortCol) {
	LCL().SysCallN(5187, m.Instance(), uintptr(ACol))
}

func (m *TValueListEditor) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5176, m.Instance(), m.checkboxToggledPtr)
}

func (m *TValueListEditor) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5177, m.Instance(), m.editingDonePtr)
}

func (m *TValueListEditor) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5179, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TValueListEditor) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5180, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TValueListEditor) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5181, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TValueListEditor) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5184, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TValueListEditor) SetOnGetPickList(fn TGetPickListEvent) {
	if m.getPickListPtr != 0 {
		RemoveEventElement(m.getPickListPtr)
	}
	m.getPickListPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5178, m.Instance(), m.getPickListPtr)
}

func (m *TValueListEditor) SetOnStringsChange(fn TNotifyEvent) {
	if m.stringsChangePtr != 0 {
		RemoveEventElement(m.stringsChangePtr)
	}
	m.stringsChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5182, m.Instance(), m.stringsChangePtr)
}

func (m *TValueListEditor) SetOnStringsChanging(fn TNotifyEvent) {
	if m.stringsChangingPtr != 0 {
		RemoveEventElement(m.stringsChangingPtr)
	}
	m.stringsChangingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5183, m.Instance(), m.stringsChangingPtr)
}

func (m *TValueListEditor) SetOnValidate(fn TOnValidateEvent) {
	if m.validatePtr != 0 {
		RemoveEventElement(m.validatePtr)
	}
	m.validatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5185, m.Instance(), m.validatePtr)
}
