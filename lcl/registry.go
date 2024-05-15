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

// IRegistry Parent: IObject
type IRegistry interface {
	IObject
	Access() uint32                                            // property
	SetAccess(AValue uint32)                                   // property
	CurrentKey() HKEY                                          // property
	CurrentPath() string                                       // property
	LazyWrite() bool                                           // property
	SetLazyWrite(AValue bool)                                  // property
	RootKey() HKEY                                             // property
	SetRootKey(AValue HKEY)                                    // property
	StringSizeIncludesNull() bool                              // property
	LastError() int32                                          // property
	LastErrorMsg() string                                      // property
	CreateKey(Key string) bool                                 // function
	DeleteKey(Key string) bool                                 // function
	DeleteValue(Name string) bool                              // function
	GetDataInfo(ValueName string, OutValue *TRegDataInfo) bool // function
	GetDataSize(ValueName string) int32                        // function
	GetDataType(ValueName string) TRegDataType                 // function
	GetKeyInfo(OutValue *TRegKeyInfo) bool                     // function
	HasSubKeys() bool                                          // function
	KeyExists(Key string) bool                                 // function
	LoadKey(Key, FileName string) bool                         // function
	OpenKey(Key string, CanCreate bool) bool                   // function
	OpenKeyReadOnly(Key string) bool                           // function
	ReadCurrency(Name string) Currency                         // function
	ReadBinaryData(Name string, count int32) []byte            // function
	ReadBool(Name string) bool                                 // function
	ReadDate(Name string) TDateTime                            // function
	ReadDateTime(Name string) TDateTime                        // function
	ReadFloat(Name string) (resultDouble float64)              // function
	ReadInteger(Name string) int32                             // function
	ReadInt64(Name string) (resultInt64 int64)                 // function
	ReadString(Name string) string                             // function
	ReadStringArray(Name string) TStringArray                  // function
	ReadTime(Name string) TDateTime                            // function
	RegistryConnect(UNCName string) bool                       // function
	ReplaceKey(Key, FileName, BackUpFileName string) bool      // function
	RestoreKey(Key, FileName string) bool                      // function
	SaveKey(Key, FileName string) bool                         // function
	UnLoadKey(Key string) bool                                 // function
	ValueExists(Name string) bool                              // function
	GetKeyNames() uintptr                                      // function
	GetValueNames() uintptr                                    // function
	ReadStringList(Name string, AList IStrings)                // procedure
	CloseKey()                                                 // procedure
	CloseKey1(key HKEY)                                        // procedure
	GetKeyNames1(Strings IStrings)                             // procedure
	GetValueNames1(Strings IStrings)                           // procedure
	MoveKey(OldName, NewName string, Delete bool)              // procedure
	RenameValue(OldName, NewName string)                       // procedure
	WriteCurrency(Name string, Value Currency)                 // procedure
	WriteBinaryData(Name string, Buffer []byte)                // procedure
	WriteBool(Name string, Value bool)                         // procedure
	WriteDate(Name string, Value TDateTime)                    // procedure
	WriteDateTime(Name string, Value TDateTime)                // procedure
	WriteFloat(Name string, Value float64)                     // procedure
	WriteInteger(Name string, Value int32)                     // procedure
	WriteInt64(Name string, Value int64)                       // procedure
	WriteString(Name, Value string)                            // procedure
	WriteExpandString(Name, Value string)                      // procedure
	WriteStringArray(Name string, Arr TStringArray)            // procedure
	WriteTime(Name string, Value TDateTime)                    // procedure
}

// TRegistry Parent: TObject
type TRegistry struct {
	TObject
}

func NewRegistry() IRegistry {
	r1 := LCL().SysCallN(4745)
	return AsRegistry(r1)
}

func NewRegistry1(aaccess uint32) IRegistry {
	r1 := LCL().SysCallN(4746, uintptr(aaccess))
	return AsRegistry(r1)
}

func (m *TRegistry) Access() uint32 {
	r1 := LCL().SysCallN(4741, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TRegistry) SetAccess(AValue uint32) {
	LCL().SysCallN(4741, 1, m.Instance(), uintptr(AValue))
}

func (m *TRegistry) CurrentKey() HKEY {
	r1 := LCL().SysCallN(4748, m.Instance())
	return HKEY(r1)
}

func (m *TRegistry) CurrentPath() string {
	r1 := LCL().SysCallN(4749, m.Instance())
	return GoStr(r1)
}

func (m *TRegistry) LazyWrite() bool {
	r1 := LCL().SysCallN(4764, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRegistry) SetLazyWrite(AValue bool) {
	LCL().SysCallN(4764, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRegistry) RootKey() HKEY {
	r1 := LCL().SysCallN(4785, 0, m.Instance(), 0)
	return HKEY(r1)
}

func (m *TRegistry) SetRootKey(AValue HKEY) {
	LCL().SysCallN(4785, 1, m.Instance(), uintptr(AValue))
}

func (m *TRegistry) StringSizeIncludesNull() bool {
	r1 := LCL().SysCallN(4787, m.Instance())
	return GoBool(r1)
}

func (m *TRegistry) LastError() int32 {
	r1 := LCL().SysCallN(4762, m.Instance())
	return int32(r1)
}

func (m *TRegistry) LastErrorMsg() string {
	r1 := LCL().SysCallN(4763, m.Instance())
	return GoStr(r1)
}

func (m *TRegistry) CreateKey(Key string) bool {
	r1 := LCL().SysCallN(4747, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) DeleteKey(Key string) bool {
	r1 := LCL().SysCallN(4750, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) DeleteValue(Name string) bool {
	r1 := LCL().SysCallN(4751, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) GetDataInfo(ValueName string, OutValue *TRegDataInfo) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(4752, m.Instance(), PascalStr(ValueName), uintptr(unsafePointer(&result1)))
	*OutValue = *(*TRegDataInfo)(getPointer(result1))
	return GoBool(r1)
}

func (m *TRegistry) GetDataSize(ValueName string) int32 {
	r1 := LCL().SysCallN(4753, m.Instance(), PascalStr(ValueName))
	return int32(r1)
}

func (m *TRegistry) GetDataType(ValueName string) TRegDataType {
	r1 := LCL().SysCallN(4754, m.Instance(), PascalStr(ValueName))
	return TRegDataType(r1)
}

func (m *TRegistry) GetKeyInfo(OutValue *TRegKeyInfo) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(4755, m.Instance(), uintptr(unsafePointer(&result0)))
	*OutValue = *(*TRegKeyInfo)(getPointer(result0))
	return GoBool(r1)
}

func (m *TRegistry) HasSubKeys() bool {
	r1 := LCL().SysCallN(4760, m.Instance())
	return GoBool(r1)
}

func (m *TRegistry) KeyExists(Key string) bool {
	r1 := LCL().SysCallN(4761, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) LoadKey(Key, FileName string) bool {
	r1 := LCL().SysCallN(4765, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) OpenKey(Key string, CanCreate bool) bool {
	r1 := LCL().SysCallN(4767, m.Instance(), PascalStr(Key), PascalBool(CanCreate))
	return GoBool(r1)
}

func (m *TRegistry) OpenKeyReadOnly(Key string) bool {
	r1 := LCL().SysCallN(4768, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) ReadCurrency(Name string) Currency {
	r1 := LCL().SysCallN(4771, m.Instance(), PascalStr(Name))
	return Currency(r1)
}

func (m *TRegistry) ReadBinaryData(Name string, count int32) []byte {
	if count <= 0 {
		return nil
	}
	buffer := make([]byte, count)
	LCL().SysCallN(4769, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&buffer[0])), uintptr(count))
	return buffer
}

func (m *TRegistry) ReadBool(Name string) bool {
	r1 := LCL().SysCallN(4770, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) ReadDate(Name string) TDateTime {
	r1 := LCL().SysCallN(4772, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) ReadDateTime(Name string) TDateTime {
	r1 := LCL().SysCallN(4773, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) ReadFloat(Name string) (resultDouble float64) {
	LCL().SysCallN(4774, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TRegistry) ReadInteger(Name string) int32 {
	r1 := LCL().SysCallN(4776, m.Instance(), PascalStr(Name))
	return int32(r1)
}

func (m *TRegistry) ReadInt64(Name string) (resultInt64 int64) {
	LCL().SysCallN(4775, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TRegistry) ReadString(Name string) string {
	r1 := LCL().SysCallN(4777, m.Instance(), PascalStr(Name))
	return GoStr(r1)
}

func (m *TRegistry) ReadStringArray(Name string) TStringArray {
	r1 := LCL().SysCallN(4778, m.Instance(), PascalStr(Name))
	return TStringArray(r1)
}

func (m *TRegistry) ReadTime(Name string) TDateTime {
	r1 := LCL().SysCallN(4780, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) RegistryConnect(UNCName string) bool {
	r1 := LCL().SysCallN(4781, m.Instance(), PascalStr(UNCName))
	return GoBool(r1)
}

func (m *TRegistry) ReplaceKey(Key, FileName, BackUpFileName string) bool {
	r1 := LCL().SysCallN(4783, m.Instance(), PascalStr(Key), PascalStr(FileName), PascalStr(BackUpFileName))
	return GoBool(r1)
}

func (m *TRegistry) RestoreKey(Key, FileName string) bool {
	r1 := LCL().SysCallN(4784, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) SaveKey(Key, FileName string) bool {
	r1 := LCL().SysCallN(4786, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) UnLoadKey(Key string) bool {
	r1 := LCL().SysCallN(4788, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) ValueExists(Name string) bool {
	r1 := LCL().SysCallN(4789, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) GetKeyNames() uintptr {
	r1 := LCL().SysCallN(4756, m.Instance())
	return uintptr(r1)
}

func (m *TRegistry) GetValueNames() uintptr {
	r1 := LCL().SysCallN(4758, m.Instance())
	return uintptr(r1)
}

func RegistryClass() TClass {
	ret := LCL().SysCallN(4742)
	return TClass(ret)
}

func (m *TRegistry) ReadStringList(Name string, AList IStrings) {
	LCL().SysCallN(4779, m.Instance(), PascalStr(Name), GetObjectUintptr(AList))
}

func (m *TRegistry) CloseKey() {
	LCL().SysCallN(4743, m.Instance())
}

func (m *TRegistry) CloseKey1(key HKEY) {
	LCL().SysCallN(4744, m.Instance(), uintptr(key))
}

func (m *TRegistry) GetKeyNames1(Strings IStrings) {
	LCL().SysCallN(4757, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TRegistry) GetValueNames1(Strings IStrings) {
	LCL().SysCallN(4759, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TRegistry) MoveKey(OldName, NewName string, Delete bool) {
	LCL().SysCallN(4766, m.Instance(), PascalStr(OldName), PascalStr(NewName), PascalBool(Delete))
}

func (m *TRegistry) RenameValue(OldName, NewName string) {
	LCL().SysCallN(4782, m.Instance(), PascalStr(OldName), PascalStr(NewName))
}

func (m *TRegistry) WriteCurrency(Name string, Value Currency) {
	LCL().SysCallN(4792, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteBinaryData(Name string, Buffer []byte) {
	LCL().SysCallN(4790, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)))
}

func (m *TRegistry) WriteBool(Name string, Value bool) {
	LCL().SysCallN(4791, m.Instance(), PascalStr(Name), PascalBool(Value))
}

func (m *TRegistry) WriteDate(Name string, Value TDateTime) {
	LCL().SysCallN(4793, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteDateTime(Name string, Value TDateTime) {
	LCL().SysCallN(4794, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteFloat(Name string, Value float64) {
	LCL().SysCallN(4796, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&Value)))
}

func (m *TRegistry) WriteInteger(Name string, Value int32) {
	LCL().SysCallN(4798, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteInt64(Name string, Value int64) {
	LCL().SysCallN(4797, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&Value)))
}

func (m *TRegistry) WriteString(Name, Value string) {
	LCL().SysCallN(4799, m.Instance(), PascalStr(Name), PascalStr(Value))
}

func (m *TRegistry) WriteExpandString(Name, Value string) {
	LCL().SysCallN(4795, m.Instance(), PascalStr(Name), PascalStr(Value))
}

func (m *TRegistry) WriteStringArray(Name string, Arr TStringArray) {
	LCL().SysCallN(4800, m.Instance(), PascalStr(Name), uintptr(Arr))
}

func (m *TRegistry) WriteTime(Name string, Value TDateTime) {
	LCL().SysCallN(4801, m.Instance(), PascalStr(Name), uintptr(Value))
}
