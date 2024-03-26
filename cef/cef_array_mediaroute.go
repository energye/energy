//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefMediaRouteArray = array of ICefMediaRoute
type ICefMediaRouteArray interface {
	Instance() uintptr
	Get(index int) ICefMediaRoute
	Size() int
	Free()
	Add(value ICefMediaRoute)
	Set(value []ICefMediaRoute)
}

// TCefMediaRouteArray = array of ICefMediaRoute
type TCefMediaRouteArray struct {
	instance unsafePointer
	count    int
	values   []ICefMediaRoute
}

// MediaRouteArrayRef -> TCefMediaRouteArray
var MediaRouteArrayRef mediaRouteArray

// mediaRouteArray
type mediaRouteArray uintptr

func (*mediaRouteArray) New(count int, instance uintptr) ICefMediaRouteArray {
	return &TCefMediaRouteArray{
		instance: unsafePointer(instance),
		count:    count,
		values:   make([]ICefMediaRoute, count),
	}
}

func (m *TCefMediaRouteArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefMediaRoute
func (m *TCefMediaRouteArray) Get(index int) ICefMediaRoute {
	if m == nil {
		return nil
	}
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefMediaRoute(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

// Size 返回 ICefMediaRoute 数组长度
func (m *TCefMediaRouteArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefMediaRouteArray) Free() {
	if m == nil {
		return
	}
	if m.values != nil {
		for i, v := range m.values {
			if v != nil {
				v.Free()
				m.values[i] = nil
			}
		}
		m.values = nil
	}
	m.instance = nil
	m.count = 0
}

func (m *TCefMediaRouteArray) Add(value ICefMediaRoute) {
	m.values = append(m.values, value)
	m.count++
	m.instance = unsafePointer(m.values[0].Instance())
}

func (m *TCefMediaRouteArray) Set(value []ICefMediaRoute) {
	if m.values != nil {
		for i, v := range m.values {
			if v != nil && v.Instance() != 0 {
				v.Free()
				m.values[i] = nil
			}
		}
	}
	m.values = value
	m.count = len(value)
	m.instance = unsafePointer(m.values[0].Instance())
}
