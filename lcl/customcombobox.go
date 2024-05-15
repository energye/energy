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

// ICustomComboBox Parent: IWinControl
type ICustomComboBox interface {
	IWinControl
	CharCase() TEditCharCase                                                                             // property
	SetCharCase(AValue TEditCharCase)                                                                    // property
	DroppedDown() bool                                                                                   // property
	SetDroppedDown(AValue bool)                                                                          // property
	AutoComplete() bool                                                                                  // property
	SetAutoComplete(AValue bool)                                                                         // property
	AutoCompleteText() TComboBoxAutoCompleteText                                                         // property
	SetAutoCompleteText(AValue TComboBoxAutoCompleteText)                                                // property
	AutoDropDown() bool                                                                                  // property
	SetAutoDropDown(AValue bool)                                                                         // property
	AutoSelect() bool                                                                                    // property
	SetAutoSelect(AValue bool)                                                                           // property
	AutoSelected() bool                                                                                  // property
	SetAutoSelected(AValue bool)                                                                         // property
	ArrowKeysTraverseList() bool                                                                         // property
	SetArrowKeysTraverseList(AValue bool)                                                                // property
	Canvas() ICanvas                                                                                     // property
	DropDownCount() int32                                                                                // property
	SetDropDownCount(AValue int32)                                                                       // property
	EmulatedTextHintStatus() TEmulatedTextHintStatus                                                     // property
	Items() IStrings                                                                                     // property
	SetItems(AValue IStrings)                                                                            // property
	ItemIndex() int32                                                                                    // property
	SetItemIndex(AValue int32)                                                                           // property
	ReadOnly() bool                                                                                      // property
	SetReadOnly(AValue bool)                                                                             // property
	SelLength() int32                                                                                    // property
	SetSelLength(AValue int32)                                                                           // property
	SelStart() int32                                                                                     // property
	SetSelStart(AValue int32)                                                                            // property
	SelText() string                                                                                     // property
	SetSelText(AValue string)                                                                            // property
	Style() TComboBoxStyle                                                                               // property
	SetStyle(AValue TComboBoxStyle)                                                                      // property
	Text() string                                                                                        // property
	SetText(AValue string)                                                                               // property
	TextHint() string                                                                                    // property
	SetTextHint(AValue string)                                                                           // property
	MatchListItem(AValue string) int32                                                                   // function
	IntfGetItems()                                                                                       // procedure
	AddItem(Item string, AnObject IObject)                                                               // procedure
	AddHistoryItem(Item string, MaxHistoryCount int32, SetAsText, CaseSensitive bool)                    // procedure
	AddHistoryItem1(Item string, AnObject IObject, MaxHistoryCount int32, SetAsText, CaseSensitive bool) // procedure
	Clear()                                                                                              // procedure
	ClearSelection()                                                                                     // procedure
	SelectAll()                                                                                          // procedure
}

// TCustomComboBox Parent: TWinControl
type TCustomComboBox struct {
	TWinControl
}

func NewCustomComboBox(TheOwner IComponent) ICustomComboBox {
	r1 := LCL().SysCallN(1454, GetObjectUintptr(TheOwner))
	return AsCustomComboBox(r1)
}

func (m *TCustomComboBox) CharCase() TEditCharCase {
	r1 := LCL().SysCallN(1450, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomComboBox) SetCharCase(AValue TEditCharCase) {
	LCL().SysCallN(1450, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) DroppedDown() bool {
	r1 := LCL().SysCallN(1456, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetDroppedDown(AValue bool) {
	LCL().SysCallN(1456, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoComplete() bool {
	r1 := LCL().SysCallN(1444, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoComplete(AValue bool) {
	LCL().SysCallN(1444, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoCompleteText() TComboBoxAutoCompleteText {
	r1 := LCL().SysCallN(1445, 0, m.Instance(), 0)
	return TComboBoxAutoCompleteText(r1)
}

func (m *TCustomComboBox) SetAutoCompleteText(AValue TComboBoxAutoCompleteText) {
	LCL().SysCallN(1445, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) AutoDropDown() bool {
	r1 := LCL().SysCallN(1446, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoDropDown(AValue bool) {
	LCL().SysCallN(1446, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoSelect() bool {
	r1 := LCL().SysCallN(1447, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoSelect(AValue bool) {
	LCL().SysCallN(1447, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoSelected() bool {
	r1 := LCL().SysCallN(1448, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoSelected(AValue bool) {
	LCL().SysCallN(1448, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) ArrowKeysTraverseList() bool {
	r1 := LCL().SysCallN(1443, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetArrowKeysTraverseList(AValue bool) {
	LCL().SysCallN(1443, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) Canvas() ICanvas {
	r1 := LCL().SysCallN(1449, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomComboBox) DropDownCount() int32 {
	r1 := LCL().SysCallN(1455, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetDropDownCount(AValue int32) {
	LCL().SysCallN(1455, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) EmulatedTextHintStatus() TEmulatedTextHintStatus {
	r1 := LCL().SysCallN(1457, m.Instance())
	return TEmulatedTextHintStatus(r1)
}

func (m *TCustomComboBox) Items() IStrings {
	r1 := LCL().SysCallN(1460, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomComboBox) SetItems(AValue IStrings) {
	LCL().SysCallN(1460, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBox) ItemIndex() int32 {
	r1 := LCL().SysCallN(1459, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetItemIndex(AValue int32) {
	LCL().SysCallN(1459, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) ReadOnly() bool {
	r1 := LCL().SysCallN(1462, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetReadOnly(AValue bool) {
	LCL().SysCallN(1462, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) SelLength() int32 {
	r1 := LCL().SysCallN(1463, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetSelLength(AValue int32) {
	LCL().SysCallN(1463, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) SelStart() int32 {
	r1 := LCL().SysCallN(1464, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetSelStart(AValue int32) {
	LCL().SysCallN(1464, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) SelText() string {
	r1 := LCL().SysCallN(1465, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetSelText(AValue string) {
	LCL().SysCallN(1465, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) Style() TComboBoxStyle {
	r1 := LCL().SysCallN(1467, 0, m.Instance(), 0)
	return TComboBoxStyle(r1)
}

func (m *TCustomComboBox) SetStyle(AValue TComboBoxStyle) {
	LCL().SysCallN(1467, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) Text() string {
	r1 := LCL().SysCallN(1468, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetText(AValue string) {
	LCL().SysCallN(1468, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) TextHint() string {
	r1 := LCL().SysCallN(1469, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetTextHint(AValue string) {
	LCL().SysCallN(1469, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) MatchListItem(AValue string) int32 {
	r1 := LCL().SysCallN(1461, m.Instance(), PascalStr(AValue))
	return int32(r1)
}

func CustomComboBoxClass() TClass {
	ret := LCL().SysCallN(1451)
	return TClass(ret)
}

func (m *TCustomComboBox) IntfGetItems() {
	LCL().SysCallN(1458, m.Instance())
}

func (m *TCustomComboBox) AddItem(Item string, AnObject IObject) {
	LCL().SysCallN(1442, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject))
}

func (m *TCustomComboBox) AddHistoryItem(Item string, MaxHistoryCount int32, SetAsText, CaseSensitive bool) {
	LCL().SysCallN(1440, m.Instance(), PascalStr(Item), uintptr(MaxHistoryCount), PascalBool(SetAsText), PascalBool(CaseSensitive))
}

func (m *TCustomComboBox) AddHistoryItem1(Item string, AnObject IObject, MaxHistoryCount int32, SetAsText, CaseSensitive bool) {
	LCL().SysCallN(1441, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject), uintptr(MaxHistoryCount), PascalBool(SetAsText), PascalBool(CaseSensitive))
}

func (m *TCustomComboBox) Clear() {
	LCL().SysCallN(1452, m.Instance())
}

func (m *TCustomComboBox) ClearSelection() {
	LCL().SysCallN(1453, m.Instance())
}

func (m *TCustomComboBox) SelectAll() {
	LCL().SysCallN(1466, m.Instance())
}
