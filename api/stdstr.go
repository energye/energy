//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"
)

// StringToUTF8Ptr
//
// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return &utf8StrArr[0]
}

func PascalStr(str string) uintptr {
	if str == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(StringToUTF8Ptr(str)))
}

// 这种跟copyStr3基本一样，只是用go来处理了
func copyStr(src uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	str := make([]uint8, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}
	return string(str)
}

type GoStringHeader struct {
	Data uintptr
	Len  int
}

// 小点的字符适合此种方式，大了就不行了
func copyStr2(str uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	var ret string
	head := (*GoStringHeader)(unsafe.Pointer(&ret))
	head.Data = str
	head.Len = strLen
	return ret
}

// 最新的lz macOS下出问题了
func copyStr3(str uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	buffer := make([]uint8, strLen)
	DMove(str, uintptr(unsafe.Pointer(&buffer[0])), strLen)
	return string(buffer)
}

// GoStr pascal string to go string
func GoStr(str uintptr) string {
	l := DStrLen(str)
	if l == 0 {
		return ""
	}
	return copyStr(str, int(l))
}
