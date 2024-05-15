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

// ICefv8Value Parent: ICefBaseRefCounted
//
//	Interface representing a V8 value handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8value_t))</a>
type ICefv8Value interface {
	ICefBaseRefCounted
	// IsValid
	//  Returns true (1) if the underlying handle is valid and it can be accessed on the current thread. Do not call any other functions if this function returns false (0).
	IsValid() bool // function
	// IsUndefined
	//  True if the value type is undefined.
	IsUndefined() bool // function
	// IsNull
	//  True if the value type is null.
	IsNull() bool // function
	// IsBool
	//  True if the value type is bool.
	IsBool() bool // function
	// IsInt
	//  True if the value type is int.
	IsInt() bool // function
	// IsUInt
	//  True if the value type is unsigned int.
	IsUInt() bool // function
	// IsDouble
	//  True if the value type is double.
	IsDouble() bool // function
	// IsDate
	//  True if the value type is Date.
	IsDate() bool // function
	// IsString
	//  True if the value type is string.
	IsString() bool // function
	// IsObject
	//  True if the value type is object.
	IsObject() bool // function
	// IsArray
	//  True if the value type is array.
	IsArray() bool // function
	// IsArrayBuffer
	//  True if the value type is an ArrayBuffer.
	IsArrayBuffer() bool // function
	// IsFunction
	//  True if the value type is function.
	IsFunction() bool // function
	// IsPromise
	//  True if the value type is a Promise.
	IsPromise() bool // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefv8Value) bool // function
	// GetBoolValue
	//  Return a bool value.
	GetBoolValue() bool // function
	// GetIntValue
	//  Return an int value.
	GetIntValue() int32 // function
	// GetUIntValue
	//  Return an unsigned int value.
	GetUIntValue() uint32 // function
	// GetDoubleValue
	//  Return a double value.
	GetDoubleValue() (resultFloat64 float64) // function
	// GetDateValue
	//  Return a Date value.
	GetDateValue() (resultDateTime TDateTime) // function
	// GetStringValue
	//  Return a string value.
	GetStringValue() string // function
	// IsUserCreated
	//  Returns true (1) if this is a user created object.
	IsUserCreated() bool // function
	// HasException
	//  Returns true (1) if the last function call resulted in an exception. This attribute exists only in the scope of the current CEF value object.
	HasException() bool // function
	// GetException
	//  Returns the exception resulting from the last function call. This attribute exists only in the scope of the current CEF value object.
	GetException() ICefV8Exception // function
	// ClearException
	//  Clears the last exception and returns true (1) on success.
	ClearException() bool // function
	// WillRethrowExceptions
	//  Returns true (1) if this object will re-throw future exceptions. This attribute exists only in the scope of the current CEF value object.
	WillRethrowExceptions() bool // function
	// SetRethrowExceptions
	//  Set whether this object will re-throw future exceptions. By default exceptions are not re-thrown. If a exception is re-thrown the current context should not be accessed again until after the exception has been caught and not re-thrown. Returns true (1) on success. This attribute exists only in the scope of the current CEF value object.
	SetRethrowExceptions(rethrow bool) bool // function
	// HasValueByKey
	//  Returns true (1) if the object has a value with the specified identifier.
	HasValueByKey(key string) bool // function
	// HasValueByIndex
	//  Returns true (1) if the object has a value with the specified identifier.
	HasValueByIndex(index int32) bool // function
	// DeleteValueByKey
	//  Deletes the value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only and don't-delete values this function will return true (1) even though deletion failed.
	DeleteValueByKey(key string) bool // function
	// DeleteValueByIndex
	//  Deletes the value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly, deletion fails or an exception is thrown. For read-only and don't-delete values this function will return true (1) even though deletion failed.
	DeleteValueByIndex(index int32) bool // function
	// GetValueByKey
	//  Returns the value with the specified identifier on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	GetValueByKey(key string) ICefv8Value // function
	// GetValueByIndex
	//  Returns the value with the specified identifier on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	GetValueByIndex(index int32) ICefv8Value // function
	// SetValueByKey
	//  Associates a value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only values this function will return true (1) even though assignment failed.
	SetValueByKey(key string, value ICefv8Value, attribute TCefV8PropertyAttributes) bool // function
	// SetValueByIndex
	//  Associates a value with the specified identifier and returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only values this function will return true (1) even though assignment failed.
	SetValueByIndex(index int32, value ICefv8Value) bool // function
	// SetValueByAccessor
	//  Registers an identifier and returns true (1) on success. Access to the identifier will be forwarded to the ICefV8Accessor instance passed to cef_v8value_create_object(). Returns false (0) if this function is called incorrectly or an exception is thrown. For read-only values this function will return true (1) even though assignment failed.
	SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool // function
	// GetKeys
	//  Read the keys for the object's values into the specified vector. Integer- based keys will also be returned as strings.
	GetKeys(keys IStrings) int32 // function
	// SetUserData
	//  Sets the user data for this object and returns true (1) on success. Returns false (0) if this function is called incorrectly. This function can only be called on user created objects.
	SetUserData(data ICefv8Value) bool // function
	// GetUserData
	//  Returns the user data, if any, assigned to this object.
	GetUserData() ICefv8Value // function
	// GetExternallyAllocatedMemory
	//  Returns the amount of externally allocated memory registered for the object.
	GetExternallyAllocatedMemory() int32 // function
	// AdjustExternallyAllocatedMemory
	//  Adjusts the amount of registered external memory for the object. Used to give V8 an indication of the amount of externally allocated memory that is kept alive by JavaScript objects. V8 uses this information to decide when to perform global garbage collection. Each ICefv8Value tracks the amount of external memory associated with it and automatically decreases the global total by the appropriate amount on its destruction. |change_in_bytes| specifies the number of bytes to adjust by. This function returns the number of bytes associated with the object after the adjustment. This function can only be called on user created objects.
	AdjustExternallyAllocatedMemory(changeInBytes int32) int32 // function
	// GetArrayLength
	//  Returns the number of elements in the array.
	GetArrayLength() int32 // function
	// GetArrayBufferReleaseCallback
	//  Returns the ReleaseCallback object associated with the ArrayBuffer or NULL if the ArrayBuffer was not created with CreateArrayBuffer.
	GetArrayBufferReleaseCallback() ICefv8ArrayBufferReleaseCallback // function
	// NeuterArrayBuffer
	//  Prevent the ArrayBuffer from using it's memory block by setting the length to zero. This operation cannot be undone. If the ArrayBuffer was created with CreateArrayBuffer then ICefv8ArrayBufferReleaseCallback.ReleaseBuffer will be called to release the underlying buffer.
	NeuterArrayBuffer() bool // function
	// GetFunctionName
	//  Returns the function name.
	GetFunctionName() string // function
	// GetFunctionHandler
	//  Returns the function handler or NULL if not a CEF-created function.
	GetFunctionHandler() ICefv8Handler // function
	// ExecuteFunction
	//  Execute the function using the current V8 context. This function should only be called from within the scope of a ICefv8Handler or ICefV8Accessor callback, or in combination with calling enter() and exit() on a stored ICefv8Context reference. |object| is the receiver ('this' object) of the function. If |object| is NULL the current context's global object will be used. |arguments| is the list of arguments that will be passed to the function. Returns the function return value on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	ExecuteFunction(obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value // function
	// ExecuteFunctionWithContext
	//  Execute the function using the specified V8 context. |object| is the receiver ('this' object) of the function. If |object| is NULL the specified context's global object will be used. |arguments| is the list of arguments that will be passed to the function. Returns the function return value on success. Returns NULL if this function is called incorrectly or an exception is thrown.
	ExecuteFunctionWithContext(context ICefv8Context, obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value // function
	// ResolvePromise
	//  Resolve the Promise using the current V8 context. This function should only be called from within the scope of a ICefv8Handler or ICefV8Accessor callback, or in combination with calling enter() and exit() on a stored ICefv8Context reference. |arg| is the argument passed to the resolved promise. Returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown.
	ResolvePromise(arg ICefv8Value) bool // function
	// RejectPromise
	//  Reject the Promise using the current V8 context. This function should only be called from within the scope of a ICefv8Handler or ICefV8Accessor callback, or in combination with calling enter() and exit() on a stored ICefv8Context reference. Returns true (1) on success. Returns false (0) if this function is called incorrectly or an exception is thrown.
	RejectPromise(errorMsg string) bool // function
}

// TCefv8Value Parent: TCefBaseRefCounted
//
//	Interface representing a V8 value handle. V8 handles can only be accessed from the thread on which they are created. Valid threads for creating a V8 handle include the render process main thread (TID_RENDERER) and WebWorker threads. A task runner for posting tasks on the associated thread can be retrieved via the ICefv8context.GetTaskRunner() function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8value_t))</a>
type TCefv8Value struct {
	TCefBaseRefCounted
}

// V8ValueRef -> ICefv8Value
var V8ValueRef v8Value

// v8Value TCefv8Value Ref
type v8Value uintptr

func (m *v8Value) UnWrap(data uintptr) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1705, uintptr(data), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewUndefined() ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1697, uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewNull() ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1692, uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewBool(value bool) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1687, PascalBool(value), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewInt(value int32) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1691, uintptr(value), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewUInt(value uint32) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1696, uintptr(value), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewDouble(value float64) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1689, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewDate(value TDateTime) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1688, uintptr(unsafePointer(&value)), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewString(str string) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1695, PascalStr(str), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewObject(accessor ICefV8Accessor, interceptor ICefV8Interceptor) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1693, GetObjectUintptr(accessor), GetObjectUintptr(interceptor), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewArray(len int32) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1685, uintptr(len), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewArrayBuffer(buffer uintptr, length NativeUInt, callback ICefv8ArrayBufferReleaseCallback) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1686, uintptr(buffer), uintptr(length), GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewFunction(name string, handler ICefv8Handler) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1690, PascalStr(name), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *v8Value) NewPromise() ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1694, uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) IsValid() bool {
	r1 := CEF().SysCallN(1683, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsUndefined() bool {
	r1 := CEF().SysCallN(1681, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsNull() bool {
	r1 := CEF().SysCallN(1675, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsBool() bool {
	r1 := CEF().SysCallN(1670, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsInt() bool {
	r1 := CEF().SysCallN(1674, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsUInt() bool {
	r1 := CEF().SysCallN(1680, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsDouble() bool {
	r1 := CEF().SysCallN(1672, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsDate() bool {
	r1 := CEF().SysCallN(1671, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsString() bool {
	r1 := CEF().SysCallN(1679, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsObject() bool {
	r1 := CEF().SysCallN(1676, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsArray() bool {
	r1 := CEF().SysCallN(1668, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsArrayBuffer() bool {
	r1 := CEF().SysCallN(1669, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsFunction() bool {
	r1 := CEF().SysCallN(1673, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsPromise() bool {
	r1 := CEF().SysCallN(1677, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) IsSame(that ICefv8Value) bool {
	r1 := CEF().SysCallN(1678, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefv8Value) GetBoolValue() bool {
	r1 := CEF().SysCallN(1651, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) GetIntValue() int32 {
	r1 := CEF().SysCallN(1658, m.Instance())
	return int32(r1)
}

func (m *TCefv8Value) GetUIntValue() uint32 {
	r1 := CEF().SysCallN(1661, m.Instance())
	return uint32(r1)
}

func (m *TCefv8Value) GetDoubleValue() (resultFloat64 float64) {
	CEF().SysCallN(1653, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefv8Value) GetDateValue() (resultDateTime TDateTime) {
	CEF().SysCallN(1652, m.Instance(), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCefv8Value) GetStringValue() string {
	r1 := CEF().SysCallN(1660, m.Instance())
	return GoStr(r1)
}

func (m *TCefv8Value) IsUserCreated() bool {
	r1 := CEF().SysCallN(1682, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) HasException() bool {
	r1 := CEF().SysCallN(1665, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) GetException() ICefV8Exception {
	var resultCefV8Exception uintptr
	CEF().SysCallN(1654, m.Instance(), uintptr(unsafePointer(&resultCefV8Exception)))
	return AsCefV8Exception(resultCefV8Exception)
}

func (m *TCefv8Value) ClearException() bool {
	r1 := CEF().SysCallN(1644, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) WillRethrowExceptions() bool {
	r1 := CEF().SysCallN(1706, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) SetRethrowExceptions(rethrow bool) bool {
	r1 := CEF().SysCallN(1700, m.Instance(), PascalBool(rethrow))
	return GoBool(r1)
}

func (m *TCefv8Value) HasValueByKey(key string) bool {
	r1 := CEF().SysCallN(1667, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefv8Value) HasValueByIndex(index int32) bool {
	r1 := CEF().SysCallN(1666, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefv8Value) DeleteValueByKey(key string) bool {
	r1 := CEF().SysCallN(1646, m.Instance(), PascalStr(key))
	return GoBool(r1)
}

func (m *TCefv8Value) DeleteValueByIndex(index int32) bool {
	r1 := CEF().SysCallN(1645, m.Instance(), uintptr(index))
	return GoBool(r1)
}

func (m *TCefv8Value) GetValueByKey(key string) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1664, m.Instance(), PascalStr(key), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) GetValueByIndex(index int32) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1663, m.Instance(), uintptr(index), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) SetValueByKey(key string, value ICefv8Value, attribute TCefV8PropertyAttributes) bool {
	r1 := CEF().SysCallN(1704, m.Instance(), PascalStr(key), GetObjectUintptr(value), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefv8Value) SetValueByIndex(index int32, value ICefv8Value) bool {
	r1 := CEF().SysCallN(1703, m.Instance(), uintptr(index), GetObjectUintptr(value))
	return GoBool(r1)
}

func (m *TCefv8Value) SetValueByAccessor(key string, settings TCefV8AccessControls, attribute TCefV8PropertyAttributes) bool {
	r1 := CEF().SysCallN(1702, m.Instance(), PascalStr(key), uintptr(settings), uintptr(attribute))
	return GoBool(r1)
}

func (m *TCefv8Value) GetKeys(keys IStrings) int32 {
	r1 := CEF().SysCallN(1659, m.Instance(), GetObjectUintptr(keys))
	return int32(r1)
}

func (m *TCefv8Value) SetUserData(data ICefv8Value) bool {
	r1 := CEF().SysCallN(1701, m.Instance(), GetObjectUintptr(data))
	return GoBool(r1)
}

func (m *TCefv8Value) GetUserData() ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1662, m.Instance(), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) GetExternallyAllocatedMemory() int32 {
	r1 := CEF().SysCallN(1655, m.Instance())
	return int32(r1)
}

func (m *TCefv8Value) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	r1 := CEF().SysCallN(1643, m.Instance(), uintptr(changeInBytes))
	return int32(r1)
}

func (m *TCefv8Value) GetArrayLength() int32 {
	r1 := CEF().SysCallN(1650, m.Instance())
	return int32(r1)
}

func (m *TCefv8Value) GetArrayBufferReleaseCallback() ICefv8ArrayBufferReleaseCallback {
	var resultCefv8ArrayBufferReleaseCallback uintptr
	CEF().SysCallN(1649, m.Instance(), uintptr(unsafePointer(&resultCefv8ArrayBufferReleaseCallback)))
	return AsCefv8ArrayBufferReleaseCallback(resultCefv8ArrayBufferReleaseCallback)
}

func (m *TCefv8Value) NeuterArrayBuffer() bool {
	r1 := CEF().SysCallN(1684, m.Instance())
	return GoBool(r1)
}

func (m *TCefv8Value) GetFunctionName() string {
	r1 := CEF().SysCallN(1657, m.Instance())
	return GoStr(r1)
}

func (m *TCefv8Value) GetFunctionHandler() ICefv8Handler {
	var resultCefv8Handler uintptr
	CEF().SysCallN(1656, m.Instance(), uintptr(unsafePointer(&resultCefv8Handler)))
	return AsCefv8Handler(resultCefv8Handler)
}

func (m *TCefv8Value) ExecuteFunction(obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1647, m.Instance(), GetObjectUintptr(obj), GetObjectUintptr(arguments), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) ExecuteFunctionWithContext(context ICefv8Context, obj ICefv8Value, arguments ICefV8ValueArray) ICefv8Value {
	var resultCefv8Value uintptr
	CEF().SysCallN(1648, m.Instance(), GetObjectUintptr(context), GetObjectUintptr(obj), GetObjectUintptr(arguments), uintptr(unsafePointer(&resultCefv8Value)))
	return AsCefv8Value(resultCefv8Value)
}

func (m *TCefv8Value) ResolvePromise(arg ICefv8Value) bool {
	r1 := CEF().SysCallN(1699, m.Instance(), GetObjectUintptr(arg))
	return GoBool(r1)
}

func (m *TCefv8Value) RejectPromise(errorMsg string) bool {
	r1 := CEF().SysCallN(1698, m.Instance(), PascalStr(errorMsg))
	return GoBool(r1)
}
