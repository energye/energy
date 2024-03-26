//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefFrameIdentifierArray = array of int64
type ICefFrameIdentifierArray interface {
	Instance() uintptr
	Get(index int) int64
	Size() int
	Free()
}

// TCefFrameIdentifierArray = array of int64
type TCefFrameIdentifierArray struct {
	instance unsafePointer
	count    int
}

// FrameIdentifierArrayRef -> TCefFrameIdentifierArray
var FrameIdentifierArrayRef frameIdentifierArray

// frameIdentifierArray
type frameIdentifierArray uintptr

func (*frameIdentifierArray) New(count int, instance uintptr) ICefFrameIdentifierArray {
	return &TCefFrameIdentifierArray{
		count:    count,
		instance: unsafePointer(instance),
	}
}

func (m *TCefFrameIdentifierArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefPostDataElement
func (m *TCefFrameIdentifierArray) Get(index int) int64 {
	if m == nil || m.instance == nil {
		return 0
	}
	if index < m.count {
		return *(*int64)(unsafePointer(m.Instance() + uintptr(index)*8))
	}
	return 0
}

// Size 返回 ICefPostDataElement 数组长度
func (m *TCefFrameIdentifierArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefFrameIdentifierArray) Free() {
	if m == nil {
		return
	}
	m.instance = nil
	m.count = 0
}
