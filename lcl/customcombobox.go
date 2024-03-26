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
	r1 := LCL().SysCallN(1264, GetObjectUintptr(TheOwner))
	return AsCustomComboBox(r1)
}

func (m *TCustomComboBox) CharCase() TEditCharCase {
	r1 := LCL().SysCallN(1260, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomComboBox) SetCharCase(AValue TEditCharCase) {
	LCL().SysCallN(1260, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) DroppedDown() bool {
	r1 := LCL().SysCallN(1266, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetDroppedDown(AValue bool) {
	LCL().SysCallN(1266, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoComplete() bool {
	r1 := LCL().SysCallN(1254, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoComplete(AValue bool) {
	LCL().SysCallN(1254, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoCompleteText() TComboBoxAutoCompleteText {
	r1 := LCL().SysCallN(1255, 0, m.Instance(), 0)
	return TComboBoxAutoCompleteText(r1)
}

func (m *TCustomComboBox) SetAutoCompleteText(AValue TComboBoxAutoCompleteText) {
	LCL().SysCallN(1255, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) AutoDropDown() bool {
	r1 := LCL().SysCallN(1256, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoDropDown(AValue bool) {
	LCL().SysCallN(1256, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoSelect() bool {
	r1 := LCL().SysCallN(1257, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoSelect(AValue bool) {
	LCL().SysCallN(1257, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoSelected() bool {
	r1 := LCL().SysCallN(1258, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoSelected(AValue bool) {
	LCL().SysCallN(1258, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) ArrowKeysTraverseList() bool {
	r1 := LCL().SysCallN(1253, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetArrowKeysTraverseList(AValue bool) {
	LCL().SysCallN(1253, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) Canvas() ICanvas {
	r1 := LCL().SysCallN(1259, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomComboBox) DropDownCount() int32 {
	r1 := LCL().SysCallN(1265, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetDropDownCount(AValue int32) {
	LCL().SysCallN(1265, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) EmulatedTextHintStatus() TEmulatedTextHintStatus {
	r1 := LCL().SysCallN(1267, m.Instance())
	return TEmulatedTextHintStatus(r1)
}

func (m *TCustomComboBox) Items() IStrings {
	r1 := LCL().SysCallN(1270, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomComboBox) SetItems(AValue IStrings) {
	LCL().SysCallN(1270, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBox) ItemIndex() int32 {
	r1 := LCL().SysCallN(1269, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetItemIndex(AValue int32) {
	LCL().SysCallN(1269, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) ReadOnly() bool {
	r1 := LCL().SysCallN(1272, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetReadOnly(AValue bool) {
	LCL().SysCallN(1272, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) SelLength() int32 {
	r1 := LCL().SysCallN(1273, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetSelLength(AValue int32) {
	LCL().SysCallN(1273, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) SelStart() int32 {
	r1 := LCL().SysCallN(1274, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetSelStart(AValue int32) {
	LCL().SysCallN(1274, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) SelText() string {
	r1 := LCL().SysCallN(1275, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetSelText(AValue string) {
	LCL().SysCallN(1275, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) Style() TComboBoxStyle {
	r1 := LCL().SysCallN(1277, 0, m.Instance(), 0)
	return TComboBoxStyle(r1)
}

func (m *TCustomComboBox) SetStyle(AValue TComboBoxStyle) {
	LCL().SysCallN(1277, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) Text() string {
	r1 := LCL().SysCallN(1278, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetText(AValue string) {
	LCL().SysCallN(1278, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) TextHint() string {
	r1 := LCL().SysCallN(1279, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetTextHint(AValue string) {
	LCL().SysCallN(1279, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) MatchListItem(AValue string) int32 {
	r1 := LCL().SysCallN(1271, m.Instance(), PascalStr(AValue))
	return int32(r1)
}

func CustomComboBoxClass() TClass {
	ret := LCL().SysCallN(1261)
	return TClass(ret)
}

func (m *TCustomComboBox) IntfGetItems() {
	LCL().SysCallN(1268, m.Instance())
}

func (m *TCustomComboBox) AddItem(Item string, AnObject IObject) {
	LCL().SysCallN(1252, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject))
}

func (m *TCustomComboBox) AddHistoryItem(Item string, MaxHistoryCount int32, SetAsText, CaseSensitive bool) {
	LCL().SysCallN(1250, m.Instance(), PascalStr(Item), uintptr(MaxHistoryCount), PascalBool(SetAsText), PascalBool(CaseSensitive))
}

func (m *TCustomComboBox) AddHistoryItem1(Item string, AnObject IObject, MaxHistoryCount int32, SetAsText, CaseSensitive bool) {
	LCL().SysCallN(1251, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject), uintptr(MaxHistoryCount), PascalBool(SetAsText), PascalBool(CaseSensitive))
}

func (m *TCustomComboBox) Clear() {
	LCL().SysCallN(1262, m.Instance())
}

func (m *TCustomComboBox) ClearSelection() {
	LCL().SysCallN(1263, m.Instance())
}

func (m *TCustomComboBox) SelectAll() {
	LCL().SysCallN(1276, m.Instance())
}
