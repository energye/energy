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

// IStream Parent: IObject
type IStream interface {
	IObject
	Position() (resultInt64 int64)                                    // property
	SetPosition(AValue int64)                                         // property
	Size() (resultInt64 int64)                                        // property
	SetSize(AValue int64)                                             // property
	Read(count int32) []byte                                          // function
	Write(Buffer []byte) int32                                        // function
	Seek(Offset int32, Origin Word) int32                             // function
	Seek1(Offset int64, Origin TSeekOrigin) (resultInt64 int64)       // function
	CopyFrom(Source IStream, Count int64) (resultInt64 int64)         // function
	ReadComponent(Instance IComponent) IComponent                     // function
	ReadComponentRes(Instance IComponent) IComponent                  // function
	ReadByte() Byte                                                   // function
	ReadWord() Word                                                   // function
	ReadDWord() uint32                                                // function
	ReadQWord() QWord                                                 // function
	ReadAnsiString() string                                           // function
	ReadBuffer(count int32) []byte                                    // procedure
	WriteBuffer(Buffer []byte)                                        // procedure
	WriteComponent(Instance IComponent)                               // procedure
	WriteComponentRes(ResName string, Instance IComponent)            // procedure
	WriteDescendent(Instance, Ancestor IComponent)                    // procedure
	WriteDescendentRes(ResName string, Instance, Ancestor IComponent) // procedure
	WriteResourceHeader(ResName string, FixupInfo *int32)             // procedure
	FixupResourceHeader(FixupInfo int32)                              // procedure
	ReadResHeader()                                                   // procedure
	WriteByte(b Byte)                                                 // procedure
	WriteWord(w Word)                                                 // procedure
	WriteDWord(d uint32)                                              // procedure
	WriteQWord(q QWord)                                               // procedure
	WriteAnsiString(S string)                                         // procedure
}

// TStream Parent: TObject
type TStream struct {
	TObject
}

func NewStream() IStream {
	r1 := LCL().SysCallN(4510)
	return AsStream(r1)
}

func (m *TStream) Position() (resultInt64 int64) {
	LCL().SysCallN(4512, 0, m.Instance(), uintptr(unsafe.Pointer(&resultInt64)), uintptr(unsafe.Pointer(&resultInt64)))
	return
}

func (m *TStream) SetPosition(AValue int64) {
	LCL().SysCallN(4512, 1, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&AValue)))
}

func (m *TStream) Size() (resultInt64 int64) {
	LCL().SysCallN(4525, 0, m.Instance(), uintptr(unsafe.Pointer(&resultInt64)), uintptr(unsafe.Pointer(&resultInt64)))
	return
}

func (m *TStream) SetSize(AValue int64) {
	LCL().SysCallN(4525, 1, m.Instance(), uintptr(unsafe.Pointer(&AValue)), uintptr(unsafe.Pointer(&AValue)))
}

func (m *TStream) Read(count int32) []byte {
	_, d := sysCallBufferRead(4513, m.Instance(), count)
	return d
}

func (m *TStream) Write(Buffer []byte) int32 {
	r1 := sysCallBufferWrite(4526, m.Instance(), Buffer)
	return int32(r1)
}

func (m *TStream) Seek(Offset int32, Origin Word) int32 {
	r1 := LCL().SysCallN(4523, m.Instance(), uintptr(Offset), uintptr(Origin))
	return int32(r1)
}

func (m *TStream) Seek1(Offset int64, Origin TSeekOrigin) (resultInt64 int64) {
	LCL().SysCallN(4524, m.Instance(), uintptr(unsafe.Pointer(&Offset)), uintptr(Origin), uintptr(unsafe.Pointer(&resultInt64)))
	return
}

func (m *TStream) CopyFrom(Source IStream, Count int64) (resultInt64 int64) {
	LCL().SysCallN(4509, m.Instance(), GetObjectUintptr(Source), uintptr(unsafe.Pointer(&Count)), uintptr(unsafe.Pointer(&resultInt64)))
	return
}

func (m *TStream) ReadComponent(Instance IComponent) IComponent {
	r1 := LCL().SysCallN(4517, m.Instance(), GetObjectUintptr(Instance))
	return AsComponent(r1)
}

func (m *TStream) ReadComponentRes(Instance IComponent) IComponent {
	r1 := LCL().SysCallN(4518, m.Instance(), GetObjectUintptr(Instance))
	return AsComponent(r1)
}

func (m *TStream) ReadByte() Byte {
	r1 := LCL().SysCallN(4516, m.Instance())
	return Byte(r1)
}

func (m *TStream) ReadWord() Word {
	r1 := LCL().SysCallN(4522, m.Instance())
	return Word(r1)
}

func (m *TStream) ReadDWord() uint32 {
	r1 := LCL().SysCallN(4519, m.Instance())
	return uint32(r1)
}

func (m *TStream) ReadQWord() QWord {
	r1 := LCL().SysCallN(4520, m.Instance())
	return QWord(r1)
}

func (m *TStream) ReadAnsiString() string {
	r1 := LCL().SysCallN(4514, m.Instance())
	return GoStr(r1)
}

func StreamClass() TClass {
	ret := LCL().SysCallN(4508)
	return TClass(ret)
}

func (m *TStream) ReadBuffer(count int32) []byte {
	_, d := sysCallBufferRead(4515, m.Instance(), count)
	return d
}

func (m *TStream) WriteBuffer(Buffer []byte) {
	sysCallBufferWrite(4528, m.Instance(), Buffer)
}

func (m *TStream) WriteComponent(Instance IComponent) {
	LCL().SysCallN(4530, m.Instance(), GetObjectUintptr(Instance))
}

func (m *TStream) WriteComponentRes(ResName string, Instance IComponent) {
	LCL().SysCallN(4531, m.Instance(), PascalStr(ResName), GetObjectUintptr(Instance))
}

func (m *TStream) WriteDescendent(Instance, Ancestor IComponent) {
	LCL().SysCallN(4533, m.Instance(), GetObjectUintptr(Instance), GetObjectUintptr(Ancestor))
}

func (m *TStream) WriteDescendentRes(ResName string, Instance, Ancestor IComponent) {
	LCL().SysCallN(4534, m.Instance(), PascalStr(ResName), GetObjectUintptr(Instance), GetObjectUintptr(Ancestor))
}

func (m *TStream) WriteResourceHeader(ResName string, FixupInfo *int32) {
	var result1 uintptr
	LCL().SysCallN(4536, m.Instance(), PascalStr(ResName), uintptr(unsafe.Pointer(&result1)))
	*FixupInfo = int32(result1)
}

func (m *TStream) FixupResourceHeader(FixupInfo int32) {
	LCL().SysCallN(4511, m.Instance(), uintptr(FixupInfo))
}

func (m *TStream) ReadResHeader() {
	LCL().SysCallN(4521, m.Instance())
}

func (m *TStream) WriteByte(b Byte) {
	LCL().SysCallN(4529, m.Instance(), uintptr(b))
}

func (m *TStream) WriteWord(w Word) {
	LCL().SysCallN(4537, m.Instance(), uintptr(w))
}

func (m *TStream) WriteDWord(d uint32) {
	LCL().SysCallN(4532, m.Instance(), uintptr(d))
}

func (m *TStream) WriteQWord(q QWord) {
	LCL().SysCallN(4535, m.Instance(), uintptr(q))
}

func (m *TStream) WriteAnsiString(S string) {
	LCL().SysCallN(4527, m.Instance(), PascalStr(S))
}
