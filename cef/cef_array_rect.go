//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"unsafe"
)

// TCefRectDynArray  = array of TCefRect
type TCefRectDynArray []TCefRect

type ICefRectArray interface {
	Instance() uintptr      // 返回数组指针
	Count() int             // 返回总数
	Rects() []TCefRect      // 返回所有Rect
	Get(index int) TCefRect // 返回指定的Rect
	Free()                  // 释放掉引用
}

// TCefRectArray = array[0..(High(Integer) div SizeOf(TCefRect))-1] of TCefRect
type TCefRectArray struct {
	instance unsafePointer
	sizeOf   uintptr
	count    int
	values   []TCefRect
}

// NewCefRectArray
//
//	TCefRect 动态数组结构, 通过指针引用取值
func NewCefRectArray(count int, instance uintptr) ICefRectArray {
	return &TCefRectArray{
		instance: unsafePointer(instance),
		sizeOf:   unsafe.Sizeof(TCefRect{}),
		count:    count,
		values:   make([]TCefRect, count),
	}
}

func (m *TCefRectArray) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *TCefRectArray) Count() int {
	return m.count
}

func (m *TCefRectArray) Rects() []TCefRect {
	if len(m.values) == 0 && m.count > 0 {
		m.values = make([]TCefRect, m.count, m.count)
		for i := 0; i < m.count; i++ {
			m.values[i] = *(*TCefRect)(getPointerOffset(m.Instance(), uintptr(i)*m.sizeOf))
		}
	}
	return m.values
}

func (m *TCefRectArray) Get(index int) TCefRect {
	values := m.Rects()
	if len(values) > 0 && index < m.count {
		return values[index]
	}
	return TCefRect{}
}

func (m *TCefRectArray) Free() {
	if m == nil {
		return
	}
	m.values = nil
	m.instance = nil
	m.count = 0
}
