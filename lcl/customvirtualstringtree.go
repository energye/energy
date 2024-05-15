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

// ICustomVirtualStringTree Parent: IBaseVirtualTree
type ICustomVirtualStringTree interface {
	IBaseVirtualTree
	ImageText(Node IVirtualNode, Kind TVTImageKind, Column TColumnIndex) string                // property
	StaticText(Node IVirtualNode, Column TColumnIndex) string                                  // property
	Text(Node IVirtualNode, Column TColumnIndex) string                                        // property
	SetText(Node IVirtualNode, Column TColumnIndex, AValue string)                             // property
	ComputeNodeHeight(Canvas ICanvas, Node IVirtualNode, Column TColumnIndex, S string) int32  // function
	ContentToClipboard(Format TClipboardFormat, Source TVSTTextSourceType) HGLOBAL             // function
	ContentToHTML(Source TVSTTextSourceType, Caption string) string                            // function
	ContentToRTF(Source TVSTTextSourceType) string                                             // function
	ContentToAnsi(Source TVSTTextSourceType, Separator string) string                          // function
	ContentToText(Source TVSTTextSourceType, Separator string) string                          // function
	ContentToUnicode(Source TVSTTextSourceType, Separator string) string                       // function
	ContentToUTF16(Source TVSTTextSourceType, Separator string) string                         // function
	ContentToUTF8(Source TVSTTextSourceType, Separator string) string                          // function
	Path(Node IVirtualNode, Column TColumnIndex, TextType TVSTTextType, Delimiter Char) string // function
	SaveToCSVFile(FileNameWithPath string, IncludeHeading bool) bool                           // function
	ContentToCustom(Source TVSTTextSourceType)                                                 // procedure
	AddToSelection(Node IVirtualNode)                                                          // procedure
	RemoveFromSelection(Node IVirtualNode)                                                     // procedure
}

// TCustomVirtualStringTree Parent: TBaseVirtualTree
type TCustomVirtualStringTree struct {
	TBaseVirtualTree
}

func NewCustomVirtualStringTree(AOwner IComponent) ICustomVirtualStringTree {
	r1 := LCL().SysCallN(2492, GetObjectUintptr(AOwner))
	return AsCustomVirtualStringTree(r1)
}

func (m *TCustomVirtualStringTree) ImageText(Node IVirtualNode, Kind TVTImageKind, Column TColumnIndex) string {
	r1 := LCL().SysCallN(2493, m.Instance(), GetObjectUintptr(Node), uintptr(Kind), uintptr(Column))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) StaticText(Node IVirtualNode, Column TColumnIndex) string {
	r1 := LCL().SysCallN(2497, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) Text(Node IVirtualNode, Column TColumnIndex) string {
	r1 := LCL().SysCallN(2498, 0, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) SetText(Node IVirtualNode, Column TColumnIndex, AValue string) {
	LCL().SysCallN(2498, 1, m.Instance(), GetObjectUintptr(Node), uintptr(Column), PascalStr(AValue))
}

func (m *TCustomVirtualStringTree) ComputeNodeHeight(Canvas ICanvas, Node IVirtualNode, Column TColumnIndex, S string) int32 {
	r1 := LCL().SysCallN(2482, m.Instance(), GetObjectUintptr(Canvas), GetObjectUintptr(Node), uintptr(Column), PascalStr(S))
	return int32(r1)
}

func (m *TCustomVirtualStringTree) ContentToClipboard(Format TClipboardFormat, Source TVSTTextSourceType) HGLOBAL {
	r1 := LCL().SysCallN(2484, m.Instance(), uintptr(Format), uintptr(Source))
	return HGLOBAL(r1)
}

func (m *TCustomVirtualStringTree) ContentToHTML(Source TVSTTextSourceType, Caption string) string {
	r1 := LCL().SysCallN(2486, m.Instance(), uintptr(Source), PascalStr(Caption))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToRTF(Source TVSTTextSourceType) string {
	r1 := LCL().SysCallN(2487, m.Instance(), uintptr(Source))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToAnsi(Source TVSTTextSourceType, Separator string) string {
	r1 := LCL().SysCallN(2483, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToText(Source TVSTTextSourceType, Separator string) string {
	r1 := LCL().SysCallN(2488, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToUnicode(Source TVSTTextSourceType, Separator string) string {
	r1 := LCL().SysCallN(2491, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToUTF16(Source TVSTTextSourceType, Separator string) string {
	r1 := LCL().SysCallN(2489, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToUTF8(Source TVSTTextSourceType, Separator string) string {
	r1 := LCL().SysCallN(2490, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) Path(Node IVirtualNode, Column TColumnIndex, TextType TVSTTextType, Delimiter Char) string {
	r1 := LCL().SysCallN(2494, m.Instance(), GetObjectUintptr(Node), uintptr(Column), uintptr(TextType), uintptr(Delimiter))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) SaveToCSVFile(FileNameWithPath string, IncludeHeading bool) bool {
	r1 := LCL().SysCallN(2496, m.Instance(), PascalStr(FileNameWithPath), PascalBool(IncludeHeading))
	return GoBool(r1)
}

func CustomVirtualStringTreeClass() TClass {
	ret := LCL().SysCallN(2481)
	return TClass(ret)
}

func (m *TCustomVirtualStringTree) ContentToCustom(Source TVSTTextSourceType) {
	LCL().SysCallN(2485, m.Instance(), uintptr(Source))
}

func (m *TCustomVirtualStringTree) AddToSelection(Node IVirtualNode) {
	LCL().SysCallN(2480, m.Instance(), GetObjectUintptr(Node))
}

func (m *TCustomVirtualStringTree) RemoveFromSelection(Node IVirtualNode) {
	LCL().SysCallN(2495, m.Instance(), GetObjectUintptr(Node))
}
