//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICEFJson Parent: IObject
type ICEFJson interface {
	IObject
}

// TCEFJson Parent: TObject
type TCEFJson struct {
	TObject
}

// JsonRef -> ICEFJson
var JsonRef json

// json TCEFJson Ref
type json uintptr

func (m *json) ReadValue(aDictionary ICefDictionaryValue, aKey string, aValue *ICefValue) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(136, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefValue(result2)
	return GoBool(r1)
}

func (m *json) ReadBoolean(aDictionary ICefDictionaryValue, aKey string, aValue *bool) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(130, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = GoBool(result2)
	return GoBool(r1)
}

func (m *json) ReadInteger(aDictionary ICefDictionaryValue, aKey string, aValue *int32) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(133, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = int32(result2)
	return GoBool(r1)
}

func (m *json) ReadDouble(aDictionary ICefDictionaryValue, aKey string, aValue *float64) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(132, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = float64(result2)
	return GoBool(r1)
}

func (m *json) ReadString(aDictionary ICefDictionaryValue, aKey string, aValue *string) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(135, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = GoStr(result2)
	return GoBool(r1)
}

func (m *json) ReadBinary(aDictionary ICefDictionaryValue, aKey string, aValue *ICefBinaryValue) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(129, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefBinaryValue(result2)
	return GoBool(r1)
}

func (m *json) ReadDictionary(aDictionary ICefDictionaryValue, aKey string, aValue *ICefDictionaryValue) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(131, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefDictionaryValue(result2)
	return GoBool(r1)
}

func (m *json) ReadList(aDictionary ICefDictionaryValue, aKey string, aValue *ICefListValue) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(134, GetObjectUintptr(aDictionary), PascalStr(aKey), uintptr(unsafePointer(&result2)))
	*aValue = AsCefListValue(result2)
	return GoBool(r1)
}

func (m *json) Parse(jsonString string, options TCefJsonParserOptions) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(126, PascalStr(jsonString), uintptr(options), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *json) Parse1(json uintptr, jsonsize NativeUInt, options TCefJsonParserOptions) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(127, uintptr(json), uintptr(jsonsize), uintptr(options), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *json) ParseAndReturnError(jsonString string, options TCefJsonParserOptions, outErrorMsgOut *string) ICefValue {
	var result2 uintptr
	var resultCefValue uintptr
	CEF().SysCallN(128, PascalStr(jsonString), uintptr(options), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&resultCefValue)))
	*outErrorMsgOut = GoStr(result2)
	return AsCefValue(resultCefValue)
}

func (m *json) Write(node ICefValue, options TCefJsonWriterOptions) string {
	r1 := CEF().SysCallN(139, GetObjectUintptr(node), uintptr(options))
	return GoStr(r1)
}

func (m *json) Write1(node ICefDictionaryValue, options TCefJsonWriterOptions) string {
	r1 := CEF().SysCallN(140, GetObjectUintptr(node), uintptr(options))
	return GoStr(r1)
}

func (m *json) Write2(node ICefValue, aRsltStrings *IStringList) bool {
	var result1 uintptr
	r1 := CEF().SysCallN(141, GetObjectUintptr(node), uintptr(unsafePointer(&result1)))
	*aRsltStrings = AsStringList(result1)
	return GoBool(r1)
}

func (m *json) Write3(node ICefDictionaryValue, aRsltStrings *IStringList) bool {
	var result1 uintptr
	r1 := CEF().SysCallN(142, GetObjectUintptr(node), uintptr(unsafePointer(&result1)))
	*aRsltStrings = AsStringList(result1)
	return GoBool(r1)
}

func (m *json) SaveToFile(node ICefValue, aFileName string) bool {
	r1 := CEF().SysCallN(137, GetObjectUintptr(node), PascalStr(aFileName))
	return GoBool(r1)
}

func (m *json) SaveToFile1(node ICefDictionaryValue, aFileName string) bool {
	r1 := CEF().SysCallN(138, GetObjectUintptr(node), PascalStr(aFileName))
	return GoBool(r1)
}

func (m *json) LoadFromFile(aFileName string, aRsltNode *ICefValue, options TCefJsonParserOptions) bool {
	var result1 uintptr
	r1 := CEF().SysCallN(125, PascalStr(aFileName), uintptr(unsafePointer(&result1)), uintptr(options))
	*aRsltNode = AsCefValue(result1)
	return GoBool(r1)
}
