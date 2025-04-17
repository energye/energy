//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy go string => pascal string
// 字符串引用, 用于读取string值

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"unsafe"
)

// TString
// Pointer reference
type TString struct {
	instance unsafe.Pointer
}

// IsValid
//
//	return true if created
func (m *TString) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// Value
//
//	 get string pointer, string length
//		bytes copy
func (m *TString) Value() string {
	if !m.IsValid() {
		return ""
	}
	var v uintptr
	s, _, _ := imports.Proc(def.TString_GetValue).Call(m.Instance(), uintptr(unsafe.Pointer(&v)))
	if v != 0 && s > 0 {
		str := make([]byte, s)
		for i := 0; i < int(s); i++ {
			str[i] = *(*byte)(unsafe.Pointer(v + uintptr(i)))
		}
		return string(str)
	}
	return ""
}

// Free
//
//	Destroy this reference
func (m *TString) Free() {
	if !m.IsValid() {
		return
	}
	var instance = m.Instance()
	imports.Proc(def.TString_Free).Call(uintptr(unsafe.Pointer(&instance)))
	m.instance = nil
}

// Instance
//
//	return string value pointer
func (m *TString) Instance() uintptr {
	return uintptr(m.instance)
}

// NewTString
//
//	Create TString pointer reference
func NewTString() *TString {
	var result uintptr
	imports.Proc(def.TString_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &TString{instance: unsafe.Pointer(result)}
}

// AsTString
//
//	Convert TString
//	instance must string(TString) pointer
func AsTString(instance uintptr) *TString {
	if instance == 0 {
		return nil
	}
	return &TString{instance: unsafe.Pointer(instance)}
}
