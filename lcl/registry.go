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
	r1 := LCL().SysCallN(4103)
	return AsRegistry(r1)
}

func NewRegistry1(aaccess uint32) IRegistry {
	r1 := LCL().SysCallN(4104, uintptr(aaccess))
	return AsRegistry(r1)
}

func (m *TRegistry) Access() uint32 {
	r1 := LCL().SysCallN(4099, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TRegistry) SetAccess(AValue uint32) {
	LCL().SysCallN(4099, 1, m.Instance(), uintptr(AValue))
}

func (m *TRegistry) CurrentKey() HKEY {
	r1 := LCL().SysCallN(4106, m.Instance())
	return HKEY(r1)
}

func (m *TRegistry) CurrentPath() string {
	r1 := LCL().SysCallN(4107, m.Instance())
	return GoStr(r1)
}

func (m *TRegistry) LazyWrite() bool {
	r1 := LCL().SysCallN(4122, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRegistry) SetLazyWrite(AValue bool) {
	LCL().SysCallN(4122, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRegistry) RootKey() HKEY {
	r1 := LCL().SysCallN(4143, 0, m.Instance(), 0)
	return HKEY(r1)
}

func (m *TRegistry) SetRootKey(AValue HKEY) {
	LCL().SysCallN(4143, 1, m.Instance(), uintptr(AValue))
}

func (m *TRegistry) StringSizeIncludesNull() bool {
	r1 := LCL().SysCallN(4145, m.Instance())
	return GoBool(r1)
}

func (m *TRegistry) LastError() int32 {
	r1 := LCL().SysCallN(4120, m.Instance())
	return int32(r1)
}

func (m *TRegistry) LastErrorMsg() string {
	r1 := LCL().SysCallN(4121, m.Instance())
	return GoStr(r1)
}

func (m *TRegistry) CreateKey(Key string) bool {
	r1 := LCL().SysCallN(4105, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) DeleteKey(Key string) bool {
	r1 := LCL().SysCallN(4108, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) DeleteValue(Name string) bool {
	r1 := LCL().SysCallN(4109, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) GetDataInfo(ValueName string, OutValue *TRegDataInfo) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(4110, m.Instance(), PascalStr(ValueName), uintptr(unsafe.Pointer(&result1)))
	*OutValue = *(*TRegDataInfo)(getPointer(result1))
	return GoBool(r1)
}

func (m *TRegistry) GetDataSize(ValueName string) int32 {
	r1 := LCL().SysCallN(4111, m.Instance(), PascalStr(ValueName))
	return int32(r1)
}

func (m *TRegistry) GetDataType(ValueName string) TRegDataType {
	r1 := LCL().SysCallN(4112, m.Instance(), PascalStr(ValueName))
	return TRegDataType(r1)
}

func (m *TRegistry) GetKeyInfo(OutValue *TRegKeyInfo) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(4113, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*OutValue = *(*TRegKeyInfo)(getPointer(result0))
	return GoBool(r1)
}

func (m *TRegistry) HasSubKeys() bool {
	r1 := LCL().SysCallN(4118, m.Instance())
	return GoBool(r1)
}

func (m *TRegistry) KeyExists(Key string) bool {
	r1 := LCL().SysCallN(4119, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) LoadKey(Key, FileName string) bool {
	r1 := LCL().SysCallN(4123, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) OpenKey(Key string, CanCreate bool) bool {
	r1 := LCL().SysCallN(4125, m.Instance(), PascalStr(Key), PascalBool(CanCreate))
	return GoBool(r1)
}

func (m *TRegistry) OpenKeyReadOnly(Key string) bool {
	r1 := LCL().SysCallN(4126, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) ReadCurrency(Name string) Currency {
	r1 := LCL().SysCallN(4129, m.Instance(), PascalStr(Name))
	return Currency(r1)
}

func (m *TRegistry) ReadBinaryData(Name string, count int32) []byte {
	if count <= 0 {
		return nil
	}
	buffer := make([]byte, count)
	LCL().SysCallN(4127, m.Instance(), PascalStr(Name), uintptr(unsafe.Pointer(&buffer[0])), uintptr(count))
	return buffer
}

func (m *TRegistry) ReadBool(Name string) bool {
	r1 := LCL().SysCallN(4128, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) ReadDate(Name string) TDateTime {
	r1 := LCL().SysCallN(4130, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) ReadDateTime(Name string) TDateTime {
	r1 := LCL().SysCallN(4131, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) ReadFloat(Name string) (resultDouble float64) {
	LCL().SysCallN(4132, m.Instance(), PascalStr(Name), uintptr(unsafe.Pointer(&resultDouble)))
	return
}

func (m *TRegistry) ReadInteger(Name string) int32 {
	r1 := LCL().SysCallN(4134, m.Instance(), PascalStr(Name))
	return int32(r1)
}

func (m *TRegistry) ReadInt64(Name string) (resultInt64 int64) {
	LCL().SysCallN(4133, m.Instance(), PascalStr(Name), uintptr(unsafe.Pointer(&resultInt64)))
	return
}

func (m *TRegistry) ReadString(Name string) string {
	r1 := LCL().SysCallN(4135, m.Instance(), PascalStr(Name))
	return GoStr(r1)
}

func (m *TRegistry) ReadStringArray(Name string) TStringArray {
	r1 := LCL().SysCallN(4136, m.Instance(), PascalStr(Name))
	return TStringArray(r1)
}

func (m *TRegistry) ReadTime(Name string) TDateTime {
	r1 := LCL().SysCallN(4138, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) RegistryConnect(UNCName string) bool {
	r1 := LCL().SysCallN(4139, m.Instance(), PascalStr(UNCName))
	return GoBool(r1)
}

func (m *TRegistry) ReplaceKey(Key, FileName, BackUpFileName string) bool {
	r1 := LCL().SysCallN(4141, m.Instance(), PascalStr(Key), PascalStr(FileName), PascalStr(BackUpFileName))
	return GoBool(r1)
}

func (m *TRegistry) RestoreKey(Key, FileName string) bool {
	r1 := LCL().SysCallN(4142, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) SaveKey(Key, FileName string) bool {
	r1 := LCL().SysCallN(4144, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) UnLoadKey(Key string) bool {
	r1 := LCL().SysCallN(4146, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) ValueExists(Name string) bool {
	r1 := LCL().SysCallN(4147, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) GetKeyNames() uintptr {
	r1 := LCL().SysCallN(4114, m.Instance())
	return uintptr(r1)
}

func (m *TRegistry) GetValueNames() uintptr {
	r1 := LCL().SysCallN(4116, m.Instance())
	return uintptr(r1)
}

func RegistryClass() TClass {
	ret := LCL().SysCallN(4100)
	return TClass(ret)
}

func (m *TRegistry) ReadStringList(Name string, AList IStrings) {
	LCL().SysCallN(4137, m.Instance(), PascalStr(Name), GetObjectUintptr(AList))
}

func (m *TRegistry) CloseKey() {
	LCL().SysCallN(4101, m.Instance())
}

func (m *TRegistry) CloseKey1(key HKEY) {
	LCL().SysCallN(4102, m.Instance(), uintptr(key))
}

func (m *TRegistry) GetKeyNames1(Strings IStrings) {
	LCL().SysCallN(4115, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TRegistry) GetValueNames1(Strings IStrings) {
	LCL().SysCallN(4117, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TRegistry) MoveKey(OldName, NewName string, Delete bool) {
	LCL().SysCallN(4124, m.Instance(), PascalStr(OldName), PascalStr(NewName), PascalBool(Delete))
}

func (m *TRegistry) RenameValue(OldName, NewName string) {
	LCL().SysCallN(4140, m.Instance(), PascalStr(OldName), PascalStr(NewName))
}

func (m *TRegistry) WriteCurrency(Name string, Value Currency) {
	LCL().SysCallN(4150, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteBinaryData(Name string, Buffer []byte) {
	LCL().SysCallN(4148, m.Instance(), PascalStr(Name), uintptr(unsafe.Pointer(&Buffer[0])), uintptr(len(Buffer)))
}

func (m *TRegistry) WriteBool(Name string, Value bool) {
	LCL().SysCallN(4149, m.Instance(), PascalStr(Name), PascalBool(Value))
}

func (m *TRegistry) WriteDate(Name string, Value TDateTime) {
	LCL().SysCallN(4151, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteDateTime(Name string, Value TDateTime) {
	LCL().SysCallN(4152, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteFloat(Name string, Value float64) {
	LCL().SysCallN(4154, m.Instance(), PascalStr(Name), uintptr(unsafe.Pointer(&Value)))
}

func (m *TRegistry) WriteInteger(Name string, Value int32) {
	LCL().SysCallN(4156, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteInt64(Name string, Value int64) {
	LCL().SysCallN(4155, m.Instance(), PascalStr(Name), uintptr(unsafe.Pointer(&Value)))
}

func (m *TRegistry) WriteString(Name, Value string) {
	LCL().SysCallN(4157, m.Instance(), PascalStr(Name), PascalStr(Value))
}

func (m *TRegistry) WriteExpandString(Name, Value string) {
	LCL().SysCallN(4153, m.Instance(), PascalStr(Name), PascalStr(Value))
}

func (m *TRegistry) WriteStringArray(Name string, Arr TStringArray) {
	LCL().SysCallN(4158, m.Instance(), PascalStr(Name), uintptr(Arr))
}

func (m *TRegistry) WriteTime(Name string, Value TDateTime) {
	LCL().SysCallN(4159, m.Instance(), PascalStr(Name), uintptr(Value))
}
