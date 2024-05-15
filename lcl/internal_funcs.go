//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import "unsafe"

// GetInstance As操作的简化。
//
// Simplification of As operation.
//
//go:noinline
func GetInstance(value interface{}) unsafe.Pointer {
	switch value.(type) {
	case uintptr:
		// 一个对象来自已经存在的对象实例指针
		// an object from a pointer to an existing object instance
		return unsafe.Pointer(value.(uintptr))
	case unsafe.Pointer:
		// 一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
		// An object from an unsafe address. Note: Using this function may cause some unknown situations. Use it with caution.
		return value.(unsafe.Pointer)
	case IObject:
		// 一个对象来自已经存在的对象实例。
		// An object from an existing object instance.
		return unsafe.Pointer(CheckPtr(value))
	default:
		// 尝试
		return unsafe.Pointer(getUIntPtr(value))
	}
}

// GetObjectUintptr 获取类指针
func GetObjectUintptr(object IObject) uintptr {
	if object == nil {
		return 0
	}
	switch object.(type) {
	case IObject:
		return object.Instance()
	case IUnknown:
		return GetUnknownUintptr(object.(IUnknown))
	}
	return 0
}

// GetUnknownUintptr 获取接口指针
func GetUnknownUintptr(unknown IUnknown) uintptr {
	if unknown == nil {
		return 0
	}
	return unknown.Instance()
}

func getUIntPtr(v interface{}) uintptr {
	switch v.(type) {
	case int:
		return uintptr(v.(int))
	case uint:
		return uintptr(v.(uint))
	case int8:
		return uintptr(v.(int8))
	case uint8:
		return uintptr(v.(uint8))
	case int16:
		return uintptr(v.(int16))
	case uint16:
		return uintptr(v.(uint16))
	case int32:
		return uintptr(v.(int32))
	case uint32:
		return uintptr(v.(uint32))
	case int64:
		return uintptr(v.(int64))
	case uint64:
		return uintptr(v.(uint64))
	}
	return 0
}

// 获取控件buffer文本，内部使用的
//
//go:noinline
func getControlBufferText(aGetTextLen func() int32, aGetTextBuf func(Buffer *string, BufSize int32) int32) string {
	strLen := aGetTextLen()
	if strLen != 0 {
		var buffStr string
		aGetTextBuf(&buffStr, strLen+1)
		return buffStr
	}
	return ""
}

// 内部的从流获取
func getStreamText(stream IStream) string {
	if stream != nil && stream.IsValid() {
		size := stream.Size()
		stream.SetPosition(0)
		if size > 0 {
			buff := stream.Read(int32(size))
			if buff[len(buff)-1] == 0 {
				return string(buff[:len(buff)-1])
			} else {
				return string(buff)
			}
		}
	}
	return ""
}
