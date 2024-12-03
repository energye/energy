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

// TCefDraggableRegion 拖拽区域集
type TCefDraggableRegion struct {
	Bounds    TCefRect
	Draggable bool
}

// TCefDraggableRegions 拖拽区域集合
type TCefDraggableRegions struct {
	regions      []TCefDraggableRegion
	regionsCount int
}

// NewCefDraggableRegion 创建一个拖拽区域
func NewCefDraggableRegion(rect TCefRect, draggable bool) TCefDraggableRegion {
	return TCefDraggableRegion{
		Bounds:    rect,
		Draggable: draggable,
	}
}

// NewCefDraggableRegions 创建拖拽区域
func NewCefDraggableRegions() *TCefDraggableRegions {
	return &TCefDraggableRegions{
		regions: make([]TCefDraggableRegion, 0),
	}
}

// Regions 获取拖拽区域
func (m *TCefDraggableRegions) Regions() []TCefDraggableRegion {
	if m.RegionsCount() == 0 || m.regions == nil || len(m.regions) == 0 {
		m.Append(NewCefDraggableRegion(NewCefRect(0, 0, 0, 0), false))
	}
	return m.regions
}

// Region 获取指定的拖拽区域
func (m *TCefDraggableRegions) Region(i int) *TCefDraggableRegion {
	if m.regions != nil && i < m.regionsCount {
		return &m.regions[i]
	}
	return nil
}

// Append 添加拖拽区域
func (m *TCefDraggableRegions) Append(region TCefDraggableRegion) {
	m.regions = append(m.regions, region)
	m.regionsCount = len(m.regions)
}

// RegionsCount 拖拽区域数量
func (m *TCefDraggableRegions) RegionsCount() int {
	return m.regionsCount
}
