//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 无标题窗口在html中通过css属性配置拖拽区域

package cef

import (
	"unsafe"
)

// TCefDraggableRegionArray = array[0..(High(Integer) div SizeOf(TCefDraggableRegion))-1]  of TCefDraggableRegion;
type TCefDraggableRegionArray []TCefDraggableRegion

// ICefDraggableRegion 拖拽区域集合
type ICefDraggableRegion interface {
	Instance() uintptr
	Count() int
	Regions() []TCefDraggableRegion
	Get(index int) TCefDraggableRegion
	Free()
}

// TCefDraggableRegions 拖拽区域集合
type TCefDraggableRegions struct {
	instance unsafePointer
	sizeOf   uintptr
	count    int
	values   []TCefDraggableRegion
}

// NewCefDraggableRegion 创建一个拖拽区域
func NewCefDraggableRegion(rect *TCefRect, isDraggable bool) TCefDraggableRegion {
	var draggable int32
	if isDraggable {
		draggable = 1 // true(1)
	}
	return TCefDraggableRegion{
		Bounds:    *rect,
		Draggable: draggable,
	}
}

// NewCefDraggableRegions 创建拖拽区域
func NewCefDraggableRegions(count int, instance uintptr) ICefDraggableRegion {
	return &TCefDraggableRegions{
		instance: unsafePointer(instance),
		sizeOf:   unsafe.Sizeof(TCefDraggableRegion{}),
		count:    count,
	}
}

func (m *TCefDraggableRegions) Instance() uintptr {
	return uintptr(m.instance)
}

// Regions 获取拖拽区域
func (m *TCefDraggableRegions) Regions() []TCefDraggableRegion {
	if len(m.values) == 0 && m.count > 0 {
		m.values = make([]TCefDraggableRegion, m.count, m.count)
		for i := 0; i < m.count; i++ {
			m.values[i] = *(*TCefDraggableRegion)(getPointerOffset(m.Instance(), uintptr(i)*m.sizeOf))
		}
	}
	return m.values
}

func (m *TCefDraggableRegions) Get(index int) TCefDraggableRegion {
	values := m.Regions()
	if len(values) > 0 && index < m.count {
		return values[index]
	}
	return TCefDraggableRegion{}
}

// Count 拖拽区域数量
func (m *TCefDraggableRegions) Count() int {
	return m.count
}

func (m *TCefDraggableRegions) Free() {
	if m == nil {
		return
	}
	m.values = nil
	m.instance = nil
	m.count = 0
}
