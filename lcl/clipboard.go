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

// IClipboard Parent: IPersistent
type IClipboard interface {
	IPersistent
	AsText() string                                           // property
	SetAsText(AValue string)                                  // property
	ClipboardType() TClipboardType                            // property
	FormatCount() int32                                       // property
	Formats(Index int32) TClipboardFormat                     // property
	AddFormat(FormatID TClipboardFormat, Stream IStream) bool // function
	AddFormat1(FormatID TClipboardFormat, Buffer []byte) bool // function
	FindPictureFormatID() TClipboardFormat                    // function
	FindFormatID(FormatName string) TClipboardFormat          // function
	GetAsHtml(ExtractFragmentOnly bool) string                // function
	GetComponent(Owner, Parent IComponent) IComponent         // function
	GetFormat(FormatID TClipboardFormat, Stream IStream) bool // function
	GetTextBuf(Buffer *string, BufSize int32) int32           // function
	HasFormat(FormatID TClipboardFormat) bool                 // function
	HasFormatName(FormatName string) bool                     // function
	HasPictureFormat() bool                                   // function
	SetComponent(Component IComponent) bool                   // function
	SetComponentAsText(Component IComponent) bool             // function
	SetFormat(FormatID TClipboardFormat, Stream IStream) bool // function
	AssignTo(Dest IPersistent)                                // procedure
	Clear()                                                   // procedure
	Close()                                                   // procedure
	SupportedFormats(List IStrings)                           // procedure
	Open()                                                    // procedure
	SetAsHtml(Html string)                                    // procedure
	SetAsHtml1(Html string, PlainText string)                 // procedure
	SetTextBuf(Buffer string)                                 // procedure
	SetOnRequest(fn TClipboardRequestEvent)                   // property event
}

// TClipboard Parent: TPersistent
type TClipboard struct {
	TPersistent
	requestPtr uintptr
}

func NewClipboard() IClipboard {
	r1 := LCL().SysCallN(665)
	return AsClipboard(r1)
}

func NewClipboard1(AClipboardType TClipboardType) IClipboard {
	r1 := LCL().SysCallN(666, uintptr(AClipboardType))
	return AsClipboard(r1)
}

func (m *TClipboard) AsText() string {
	r1 := LCL().SysCallN(659, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TClipboard) SetAsText(AValue string) {
	LCL().SysCallN(659, 1, m.Instance(), PascalStr(AValue))
}

func (m *TClipboard) ClipboardType() TClipboardType {
	r1 := LCL().SysCallN(663, m.Instance())
	return TClipboardType(r1)
}

func (m *TClipboard) FormatCount() int32 {
	r1 := LCL().SysCallN(669, m.Instance())
	return int32(r1)
}

func (m *TClipboard) Formats(Index int32) TClipboardFormat {
	r1 := LCL().SysCallN(670, m.Instance(), uintptr(Index))
	return TClipboardFormat(r1)
}

func (m *TClipboard) AddFormat(FormatID TClipboardFormat, Stream IStream) bool {
	r1 := LCL().SysCallN(657, m.Instance(), uintptr(FormatID), GetObjectUintptr(Stream))
	return GoBool(r1)
}

func (m *TClipboard) AddFormat1(FormatID TClipboardFormat, Buffer []byte) bool {
	r1 := LCL().SysCallN(658, m.Instance(), uintptr(FormatID), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)))
	return GoBool(r1)
}

func (m *TClipboard) FindPictureFormatID() TClipboardFormat {
	r1 := LCL().SysCallN(668, m.Instance())
	return TClipboardFormat(r1)
}

func (m *TClipboard) FindFormatID(FormatName string) TClipboardFormat {
	r1 := LCL().SysCallN(667, m.Instance(), PascalStr(FormatName))
	return TClipboardFormat(r1)
}

func (m *TClipboard) GetAsHtml(ExtractFragmentOnly bool) string {
	r1 := LCL().SysCallN(671, m.Instance(), PascalBool(ExtractFragmentOnly))
	return GoStr(r1)
}

func (m *TClipboard) GetComponent(Owner, Parent IComponent) IComponent {
	r1 := LCL().SysCallN(672, m.Instance(), GetObjectUintptr(Owner), GetObjectUintptr(Parent))
	return AsComponent(r1)
}

func (m *TClipboard) GetFormat(FormatID TClipboardFormat, Stream IStream) bool {
	r1 := LCL().SysCallN(673, m.Instance(), uintptr(FormatID), GetObjectUintptr(Stream))
	return GoBool(r1)
}

func (m *TClipboard) GetTextBuf(Buffer *string, BufSize int32) int32 {
	r1 := sysCallGetTextBuf(674, m.Instance(), Buffer, BufSize)
	return int32(r1)
}

func (m *TClipboard) HasFormat(FormatID TClipboardFormat) bool {
	r1 := LCL().SysCallN(675, m.Instance(), uintptr(FormatID))
	return GoBool(r1)
}

func (m *TClipboard) HasFormatName(FormatName string) bool {
	r1 := LCL().SysCallN(676, m.Instance(), PascalStr(FormatName))
	return GoBool(r1)
}

func (m *TClipboard) HasPictureFormat() bool {
	r1 := LCL().SysCallN(677, m.Instance())
	return GoBool(r1)
}

func (m *TClipboard) SetComponent(Component IComponent) bool {
	r1 := LCL().SysCallN(681, m.Instance(), GetObjectUintptr(Component))
	return GoBool(r1)
}

func (m *TClipboard) SetComponentAsText(Component IComponent) bool {
	r1 := LCL().SysCallN(682, m.Instance(), GetObjectUintptr(Component))
	return GoBool(r1)
}

func (m *TClipboard) SetFormat(FormatID TClipboardFormat, Stream IStream) bool {
	r1 := LCL().SysCallN(683, m.Instance(), uintptr(FormatID), GetObjectUintptr(Stream))
	return GoBool(r1)
}

func ClipboardClass() TClass {
	ret := LCL().SysCallN(661)
	return TClass(ret)
}

func (m *TClipboard) AssignTo(Dest IPersistent) {
	LCL().SysCallN(660, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TClipboard) Clear() {
	LCL().SysCallN(662, m.Instance())
}

func (m *TClipboard) Close() {
	LCL().SysCallN(664, m.Instance())
}

func (m *TClipboard) SupportedFormats(List IStrings) {
	LCL().SysCallN(686, m.Instance(), GetObjectUintptr(List))
}

func (m *TClipboard) Open() {
	LCL().SysCallN(678, m.Instance())
}

func (m *TClipboard) SetAsHtml(Html string) {
	LCL().SysCallN(679, m.Instance(), PascalStr(Html))
}

func (m *TClipboard) SetAsHtml1(Html string, PlainText string) {
	LCL().SysCallN(680, m.Instance(), PascalStr(Html), PascalStr(PlainText))
}

func (m *TClipboard) SetTextBuf(Buffer string) {
	LCL().SysCallN(685, m.Instance(), PascalStr(Buffer))
}

func (m *TClipboard) SetOnRequest(fn TClipboardRequestEvent) {
	if m.requestPtr != 0 {
		RemoveEventElement(m.requestPtr)
	}
	m.requestPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(684, m.Instance(), m.requestPtr)
}
