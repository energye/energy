//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "unsafe"

// TCefCompositionUnderlineDynArray = array of TCefCompositionUnderline
type TCefCompositionUnderlineDynArray = []TCefCompositionUnderline

// ICefCompositionUnderlineArray = array[0..(High(integer) div SizeOf(TCefCompositionUnderline)) - 1] of TCefCompositionUnderline;
type ICefCompositionUnderlineArray interface {
	Instance() uintptr
	Count() int
	CompositionUnderlines() []TCefCompositionUnderline
	Get(index int) TCefCompositionUnderline
	Free()
}

// TCefCompositionUnderlineArray = array[0..(High(integer) div SizeOf(TCefCompositionUnderline)) - 1] of TCefCompositionUnderline;
type TCefCompositionUnderlineArray struct {
	instance unsafePointer
	sizeOf   uintptr
	count    int
	values   []TCefCompositionUnderline
}

// NewCefCompositionUnderlineArray
//
//	TCefRect 动态数组结构, 通过指针引用取值
func NewCefCompositionUnderlineArray(count int, instance uintptr) ICefCompositionUnderlineArray {
	return &TCefCompositionUnderlineArray{
		instance: unsafePointer(instance),
		sizeOf:   unsafe.Sizeof(TCefCompositionUnderline{}),
		count:    count,
		values:   make([]TCefCompositionUnderline, count),
	}
}

func (m *TCefCompositionUnderlineArray) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *TCefCompositionUnderlineArray) Count() int {
	return m.count
}

func (m *TCefCompositionUnderlineArray) CompositionUnderlines() []TCefCompositionUnderline {
	if len(m.values) == 0 && m.count > 0 {
		m.values = make([]TCefCompositionUnderline, m.count, m.count)
		for i := 0; i < m.count; i++ {
			m.values[i] = *(*TCefCompositionUnderline)(getPointerOffset(m.Instance(), uintptr(i)*m.sizeOf))
		}
	}
	return m.values
}

func (m *TCefCompositionUnderlineArray) Get(index int) TCefCompositionUnderline {
	values := m.CompositionUnderlines()
	if len(values) > 0 && index < m.count {
		return values[index]
	}
	return TCefCompositionUnderline{}
}

func (m *TCefCompositionUnderlineArray) Free() {
	if m == nil {
		return
	}
	m.values = nil
	m.instance = nil
	m.count = 0
}
