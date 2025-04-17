//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF v8 value V8所有类型
//
// ICefV8Value 创建和使用一搬在 v8context 上下文中使用

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"time"
	"unsafe"
)

// ICefV8Value
//
//	CEF V8 值类型, 对应到 JavaScrip 的类型, 使用该对象时需要合理的管理释放
type ICefV8Value struct {
	base         TCefBaseRefCounted
	instance     unsafe.Pointer
	valueType    consts.V8ValueType      // 值类型
	valueByIndex []*ICefV8Value          // 当前对象的所有数组集合
	valueByKeys  map[string]*ICefV8Value // 当前对象的所有key=value子集合
	cantNotFree  bool                    // 是否允许释放, false时允许释放
}

func (m *ICefV8Value) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8Value) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) IsUndefined() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtUndefined {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsUndefined).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtUndefined
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsNull() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtNull {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsNull).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtNull
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsBool() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtBool {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsBool).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtBool
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsInt() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtInt {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsInt).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtInt
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsUInt() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtUInt {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsUInt).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtUInt
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsDouble() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtDouble {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsDouble).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtDouble
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsDate() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtDate {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsDate).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtDate
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsString() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtString {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsString).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtString
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsObject() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtObject {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsObject).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtObject
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsArray() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtArray {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsArray).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtArray
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsArrayBuffer() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtArrayBuffer {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsArrayBuffer).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtArrayBuffer
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsFunction() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtFunction {
		//return true
	}
	//if m.valueType == consts.V8vtInvalid {
	r1, _, _ := imports.Proc(def.CefV8Value_IsFunction).Call(m.Instance())
	if api.GoBool(r1) {
		m.valueType = consts.V8vtFunction
		return true
	}
	//}
	return false
}

func (m *ICefV8Value) IsPromise() bool {
	if !m.IsValid() {
		return false
	}
	if m.valueType == consts.V8vtPromise {
		return true
	}
	if m.valueType == consts.V8vtInvalid {
		r1, _, _ := imports.Proc(def.CefV8Value_IsPromise).Call(m.Instance())
		if api.GoBool(r1) {
			m.valueType = consts.V8vtPromise
			return true
		}
	}
	return false
}

func (m *ICefV8Value) IsSame() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_IsSame).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) GetBoolValue() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_GetBoolValue).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) GetIntValue() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Value_GetIntValue).Call(m.Instance())
	return int32(r1)
}

func (m *ICefV8Value) GetUIntValue() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Value_GetUIntValue).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefV8Value) GetDoubleValue() (result float64) {
	if !m.IsValid() {
		return 0.0
	}
	imports.Proc(def.CefV8Value_GetDoubleValue).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *ICefV8Value) GetDateValue() time.Time {
	if !m.IsValid() {
		return time.Time{}
	}
	var result float64
	imports.Proc(def.CefV8Value_GetDateValue).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return common.DDateTimeToGoDateTime(result)
}

func (m *ICefV8Value) GetStringValue() (value string) {
	if !m.IsValid() {
		return ""
	}
	val := NewTString()
	imports.Proc(def.CefV8Value_GetStringValue).Call(m.Instance(), val.Instance())
	value = val.Value()
	val.Free()
	return
}

func (m *ICefV8Value) IsUserCreated() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_IsUserCreated).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) HasException() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_HasException).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) GetException() *ICefV8Exception {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefV8Value_GetException).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefV8Exception{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefV8Value) ClearException() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_ClearException).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) WillRethrowExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_WillRethrowExceptions).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) SetRethrowExceptions(reThrow bool) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_SetRethrowExceptions).Call(m.Instance(), api.PascalBool(reThrow))
	return api.GoBool(r1)
}

func (m *ICefV8Value) HasValueByKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_HasValueByKey).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

func (m *ICefV8Value) HasValueByIndex(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_HasValueByIndex).Call(m.Instance(), uintptr(index))
	return api.GoBool(r1)
}

// deleteValueByKey internal
func (m *ICefV8Value) deleteValueByKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	if m.valueByKeys != nil {
		if v, ok := m.valueByKeys[key]; ok {
			v.Free()
			delete(m.valueByKeys, key)
		}
	}
	r1, _, _ := imports.Proc(def.CefV8Value_DeleteValueByKey).Call(m.Instance(), api.PascalStr(key))
	return api.GoBool(r1)
}

// DeleteValueByKey export
func (m *ICefV8Value) DeleteValueByKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	if isIPCInternalKey(key) {
		return false
	}
	return m.deleteValueByKey(key)
}

func (m *ICefV8Value) DeleteValueByIndex(index int) bool {
	if !m.IsValid() {
		return false
	}
	if m.valueByIndex != nil {
		if v := m.valueByIndex[index]; v != nil {
			v.Free()
			m.valueByIndex[index] = nil
		}
	}
	r1, _, _ := imports.Proc(def.CefV8Value_DeleteValueByIndex).Call(m.Instance(), uintptr(int32(index)))
	return api.GoBool(r1)
}

// getValueByKey internal
func (m *ICefV8Value) getValueByKey(key string) *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	if !m.IsObject() {
		return nil
	}
	if m.valueByKeys == nil {
		m.valueByKeys = make(map[string]*ICefV8Value)
	}
	if v, ok := m.valueByKeys[key]; ok {
		v.Free()
	}
	var result uintptr
	imports.Proc(def.CefV8Value_GetValueByKey).Call(m.Instance(), api.PascalStr(key), uintptr(unsafe.Pointer(&result)))
	v := &ICefV8Value{
		instance: unsafe.Pointer(result),
	}
	m.valueByKeys[key] = v
	return v
}

// GetValueByKey export
func (m *ICefV8Value) GetValueByKey(key string) *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	if key == internalIPC {
		return nil
	}
	return m.getValueByKey(key)
}

// GetValueByIndex 当前是数组时，通过下标取值V8Value
func (m *ICefV8Value) GetValueByIndex(index int) *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	if !m.IsArray() {
		return nil
	}
	if m.valueByIndex == nil {
		m.valueByIndex = make([]*ICefV8Value, m.GetArrayLength())
	}
	if v := m.valueByIndex[index]; v != nil {
		v.Free()
	}
	var result uintptr
	imports.Proc(def.CefV8Value_GetValueByIndex).Call(m.Instance(), uintptr(int32(index)), uintptr(unsafe.Pointer(&result)))
	v := &ICefV8Value{
		instance: unsafe.Pointer(result),
	}
	m.valueByIndex[index] = v
	return v
}

// setValueByKey internal
func (m *ICefV8Value) setValueByKey(key string, value *ICefV8Value, attribute consts.TCefV8PropertyAttributes) bool {
	if !m.IsValid() {
		return false
	}
	if m.valueByKeys != nil {
		if v, ok := m.valueByKeys[key]; ok {
			if v != value {
				v.Free()
			}
		}
		m.valueByKeys[key] = value
	}
	r1, _, _ := imports.Proc(def.CefV8Value_SetValueByKey).Call(m.Instance(), api.PascalStr(key), value.Instance(), attribute.ToPtr())
	return api.GoBool(r1)
}

// SetValueByKey export
func (m *ICefV8Value) SetValueByKey(key string, value *ICefV8Value, attribute consts.TCefV8PropertyAttributes) bool {
	if !m.IsValid() {
		return false
	}
	if isIPCInternalKey(key) {
		return false
	}
	return m.setValueByKey(key, value, attribute)
}

func (m *ICefV8Value) SetValueByIndex(index int32, value *ICefV8Value) bool {
	if !m.IsValid() {
		return false
	}
	if m.valueByIndex != nil {
		if v := m.valueByIndex[index]; v != nil {
			if v != value {
				v.Free()
			}
		}
		m.valueByIndex[index] = value
	}
	r1, _, _ := imports.Proc(def.CefV8Value_SetValueByIndex).Call(m.Instance(), uintptr(index), value.Instance())
	return api.GoBool(r1)
}

// SetValueByAccessor internal
func (m *ICefV8Value) setValueByAccessor(key string, attribute consts.TCefV8PropertyAttributes) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_SetValueByAccessor).Call(m.Instance(), api.PascalStr(key), attribute.ToPtr())
	return api.GoBool(r1)
}

// SetValueByAccessor export
func (m *ICefV8Value) SetValueByAccessor(key string, attribute consts.TCefV8PropertyAttributes) bool {
	if !m.IsValid() {
		return false
	}
	if isIPCInternalKey(key) {
		return false
	}
	return m.setValueByAccessor(key, attribute)
}

func (m *ICefV8Value) GetKeys() *ICefV8ValueKeys {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	r1, _, _ := imports.Proc(def.CefV8Value_GetKeys).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8ValueKeys{keys: lcl.AsStrings(result), count: int(int32(r1))}
}

func (m *ICefV8Value) GetIKeys() types.IV8ValueKeys {
	if !m.IsValid() {
		return nil
	}
	return m.GetKeys()
}

func (m *ICefV8Value) SetUserData(data *ICefV8Value) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_SetUserData).Call(m.Instance(), data.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) GetUserData() *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefV8Value_GetUserData).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefV8Value) GetExternallyAllocatedMemory() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Value_GetExternallyAllocatedMemory).Call(m.Instance())
	return int32(r1)
}

func (m *ICefV8Value) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Value_AdjustExternallyAllocatedMemory).Call(m.Instance(), uintptr(changeInBytes))
	return int32(r1)
}

func (m *ICefV8Value) GetArrayLength() int {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CefV8Value_GetArrayLength).Call(m.Instance())
	return int(int32(r1))
}

func (m *ICefV8Value) GetArrayBufferReleaseCallback() *ICefV8ArrayBufferReleaseCallback {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefV8Value_GetArrayBufferReleaseCallback).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8ArrayBufferReleaseCallback{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefV8Value) NeuterArrayBuffer() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_NeuterArrayBuffer).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) GetFunctionName() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CefV8Value_GetFunctionName).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefV8Value) GetFunctionHandler() *ICefV8Handler {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefV8Value_GetFunctionHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Handler{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefV8Value) ExecuteFunction(obj *ICefV8Value, arguments *TCefV8ValueArray) *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	var argumentsPtr = arguments.Instance()
	if arguments.Size() > 0 && arguments.argumentsCollect != nil {
		var args = make([]uintptr, arguments.Size(), arguments.Size())
		for i, a := range arguments.argumentsCollect {
			args[i] = a.Instance()
		}
		argumentsPtr = uintptr(unsafe.Pointer(&args[0]))
	}
	imports.Proc(def.CefV8Value_ExecuteFunction).Call(m.Instance(), obj.Instance(), argumentsPtr, uintptr(int32(arguments.Size())), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefV8Value) ExecuteFunctionWithContext(v8Context *ICefV8Context, obj *ICefV8Value, arguments *TCefV8ValueArray) *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	var argumentsPtr = arguments.Instance()
	if arguments.Size() > 0 && arguments.argumentsCollect != nil {
		var args = make([]uintptr, arguments.Size(), arguments.Size())
		for i, v := range arguments.argumentsCollect {
			args[i] = v.Instance()
		}
		argumentsPtr = uintptr(unsafe.Pointer(&args[0]))
	}
	//argumentsPtr = uintptr(unsafe.Pointer(&argumentsPtr))
	imports.Proc(def.CefV8Value_ExecuteFunctionWithContext).Call(m.Instance(), v8Context.Instance(), obj.Instance(), argumentsPtr, uintptr(int32(arguments.Size())), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefV8Value) ExecuteFunctionWithContextForArgsBytes(v8Context *ICefV8Context, obj *ICefV8Value, arguments []byte) *ICefV8Value {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	var argumentsPtr = uintptr(unsafe.Pointer(&arguments[0]))
	var argumentsLength = uintptr(uint32(len(arguments)))
	imports.Proc(def.CefV8Value_ExecuteFunctionWithContextForArgsBytes).Call(m.Instance(), v8Context.Instance(), obj.Instance(), argumentsPtr, argumentsLength, uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance: unsafe.Pointer(result),
	}
}

func (m *ICefV8Value) ResolvePromise(arg *ICefV8Value) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_ResolvePromise).Call(m.Instance(), arg.Instance())
	return api.GoBool(r1)
}

func (m *ICefV8Value) RejectPromise(errorMsg string) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CefV8Value_RejectPromise).Call(m.Instance(), api.PascalStr(errorMsg))
	return api.GoBool(r1)
}

// SetCanNotFree
//
//	 设置是否允许释放, 可以管理对象的独立释放
//		v=false 时允许释放
func (m *ICefV8Value) SetCanNotFree(v bool) {
	if m != nil && m.instance != nil {
		m.cantNotFree = v
	}
}

func (m *ICefV8Value) Free() {
	if m != nil && m.instance != nil {
		if m.valueByIndex != nil {
			for _, v := range m.valueByIndex {
				if v != nil {
					v.Free()
				}
			}
			m.valueByIndex = nil
		}
		if m.valueByKeys != nil {
			for _, v := range m.valueByKeys {
				if v != nil {
					v.Free()
				}
			}
			m.valueByKeys = nil
		}
		if !m.cantNotFree {
			//var ptr = m.Instance()
			//imports.Proc(CefV8Value_Free).Call(uintptr(unsafe.Pointer(&ptr)))
			m.base.Free(m.Instance())
			m.instance = nil
		}
	}
}

// ResultV8Value 返回 ICefV8Value 的替代结构
type ResultV8Value struct {
	v8value *ICefV8Value
}

// SetResult 设置 ICefV8Value 返回值
func (m *ResultV8Value) SetResult(v8value *ICefV8Value) {
	m.v8value = v8value
}

// V8ValueArrayRef -> TCefV8ValueArray
var V8ValueArrayRef v8ValueArray

// v8ValueArray
type v8ValueArray uintptr

func (*v8ValueArray) New() *TCefV8ValueArray {
	return &TCefV8ValueArray{
		argumentsCollect: []*ICefV8Value{},
	}
}

func (m *TCefV8ValueArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefV8Value
func (m *TCefV8ValueArray) Get(index int) *ICefV8Value {
	if m == nil {
		return nil
	}
	if index < m.argumentsLength {
		value := m.argumentsCollect[index]
		if value == nil {
			value = &ICefV8Value{instance: unsafe.Pointer(common.GetParamOf(index, m.arguments))}
			m.argumentsCollect[index] = value
		}
		return value
	}
	return nil
}

// TCefV8ValueArray ICefV8Value 数组的替代结构
type TCefV8ValueArray struct {
	instance         unsafe.Pointer
	arguments        uintptr
	argumentsLength  int
	argumentsCollect []*ICefV8Value
}

// Size 返回 ICefV8Value 数组长度
func (m *TCefV8ValueArray) Size() int {
	if m == nil {
		return 0
	}
	return m.argumentsLength
}

func (m *TCefV8ValueArray) Free() {
	if m == nil {
		return
	}
	if m.argumentsCollect != nil {
		for _, v := range m.argumentsCollect {
			if v != nil {
				v.Free()
			}
		}
		m.argumentsCollect = nil
	}
	m.instance = nil
	m.arguments = 0
	m.argumentsLength = 0
}

func (m *TCefV8ValueArray) Add(value *ICefV8Value) {
	m.argumentsCollect = append(m.argumentsCollect, value)
	m.argumentsLength++
	m.instance = unsafe.Pointer(m.argumentsCollect[0].Instance())
	m.arguments = uintptr(m.instance)
}

func (m *TCefV8ValueArray) Set(value []*ICefV8Value) {
	if m.argumentsCollect != nil {
		for _, v := range m.argumentsCollect {
			if v != nil && v.instance != nil {
				v.Free()
			}
		}
	}
	m.argumentsCollect = value
	m.argumentsLength = len(value)
	m.instance = unsafe.Pointer(m.argumentsCollect[0].Instance())
	m.arguments = uintptr(m.instance)
}

// V8ValueRef -> ICefV8Value
var V8ValueRef cefV8Value

// cefV8Value
type cefV8Value uintptr

func (*cefV8Value) NewUndefined() *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewUndefined).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtUndefined,
	}
}

func (*cefV8Value) NewNull() *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewNull).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtNull,
	}
}

func (*cefV8Value) NewBool(value bool) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewBool).Call(api.PascalBool(value), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtBool,
	}
}
func (*cefV8Value) NewInt(value int32) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewInt).Call(uintptr(value), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtInt,
	}
}

func (*cefV8Value) NewUInt(value uint32) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewUInt).Call(uintptr(value), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtUInt,
	}
}

func (*cefV8Value) NewDouble(value float64) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewDouble).Call(uintptr(unsafe.Pointer(&value)), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtDouble,
	}
}

func (*cefV8Value) NewDate(value time.Time) *ICefV8Value {
	var result uintptr
	val := common.GoDateTimeToDDateTime(value)
	imports.Proc(def.CefV8ValueRef_NewDate).Call(uintptr(unsafe.Pointer(&val)), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtDate,
	}
}

func (*cefV8Value) NewString(value string) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewString).Call(api.PascalStr(value), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtString,
	}
}

// NewObject
func (*cefV8Value) NewObject(accessor *ICefV8Accessor, interceptor *ICefV8Interceptor) *ICefV8Value {
	var result uintptr
	if application.Is49() {
		// CEF49
		if accessor == nil || accessor.instance == nil {
			imports.Proc(def.CefV8ValueRef_NewObject).Call(uintptr(0), uintptr(unsafe.Pointer(&result)))
		} else {
			imports.Proc(def.CefV8ValueRef_NewObject).Call(accessor.Instance(), uintptr(unsafe.Pointer(&result)))
		}
	} else {
		var accessorPtr, interceptorPtr uintptr = 0, 0
		if accessor != nil {
			accessorPtr = accessor.Instance()
		}
		if interceptor != nil {
			interceptorPtr = interceptor.Instance()
		}
		imports.Proc(def.CefV8ValueRef_NewObject).Call(accessorPtr, interceptorPtr, uintptr(unsafe.Pointer(&result)))
	}
	return &ICefV8Value{
		instance:    unsafe.Pointer(result),
		valueType:   consts.V8vtObject,
		valueByKeys: make(map[string]*ICefV8Value),
	}
}

func (*cefV8Value) NewArray(len int32) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewArray).Call(uintptr(len), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:     unsafe.Pointer(result),
		valueType:    consts.V8vtArray,
		valueByIndex: make([]*ICefV8Value, len),
	}
}

func (*cefV8Value) NewArrayBuffer(buffer []byte, callback *ICefV8ArrayBufferReleaseCallback) *ICefV8Value {
	var bufLen = len(buffer)
	if bufLen == 0 {
		return nil
	}
	if callback == nil {
		callback = V8ArrayBufferReleaseCallbackRef.New()
	}
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewArrayBuffer).Call(uintptr(unsafe.Pointer(&buffer[0])), uintptr(int32(bufLen)), callback.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtArrayBuffer,
	}
}

// newFunction internal
func (*cefV8Value) newFunction(name string, handler *ICefV8Handler) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewFunction).Call(api.PascalStr(name), handler.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtFunction,
	}
}

// NewFunction export
func (m *cefV8Value) NewFunction(name string, handler *ICefV8Handler) *ICefV8Value {
	if isIPCInternalKey(name) {
		return nil
	}
	return m.newFunction(name, handler)
}

func (*cefV8Value) NewPromise() *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_NewPromise).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Value{
		instance:  unsafe.Pointer(result),
		valueType: consts.V8vtPromise,
	}
}

// UnWrap 指针包裹引用
func (*cefV8Value) UnWrap(data *ICefV8Value) *ICefV8Value {
	var result uintptr
	imports.Proc(def.CefV8ValueRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return data
}

// ICefV8ValueKeys
type ICefV8ValueKeys struct {
	keys     *lcl.TStrings
	count    int
	keyArray []string
}

func (m *ICefV8ValueKeys) Count() int {
	if m == nil || m.keys == nil {
		return 0
	}
	return m.count
}

func (m *ICefV8ValueKeys) Get(index int) string {
	if m == nil || m.keys == nil {
		return ""
	}
	count := m.Count()
	if index < count {
		if m.keyArray == nil {
			m.keyArray = make([]string, count)
		}
		if len(m.keyArray) < count {
			keyArray := m.keyArray
			m.keyArray = make([]string, count)
			copy(m.keyArray, keyArray)
		}
		value := m.keyArray[index]
		if value == "" {
			value = m.keys.Strings(int32(index))
			m.keyArray[index] = value
		}
		return value
	}
	return ""
}

func (m *ICefV8ValueKeys) Free() {
	if m == nil || m.keys == nil {
		return
	}
	m.keyArray = nil
	m.keys.Free()
	m.keys = nil
}
