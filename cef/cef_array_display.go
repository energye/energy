//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefDisplayArray = array of ICefDisplay
type ICefDisplayArray interface {
	Instance() uintptr
	Get(index int) ICefDisplay
	Size() int
	Free()
	Add(value ICefDisplay)
	Set(value []ICefDisplay)
}

// TCefDisplayArray = array of ICefDisplay
type TCefDisplayArray struct {
	instance unsafePointer
	count    int
	values   []ICefDisplay
}

// DisplayArrayRef -> TCefDisplayArray
var DisplayArrayRef displayArray

// displayArray
type displayArray uintptr

func (*displayArray) New(count int, instance uintptr) ICefDisplayArray {
	return &TCefDisplayArray{
		count:    count,
		instance: unsafePointer(instance),
		values:   make([]ICefDisplay, count),
	}
}

func (m *TCefDisplayArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefDisplay
func (m *TCefDisplayArray) Get(index int) ICefDisplay {
	if m == nil {
		return nil
	}
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefDisplay(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

// Size 返回 ICefDisplay 数组长度
func (m *TCefDisplayArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefDisplayArray) Free() {
	if m == nil {
		return
	}
	if m.values != nil {
		for _, v := range m.values {
			if v != nil {
				v.Free()
			}
		}
		m.values = nil
	}
	m.instance = nil
	m.count = 0
}

func (m *TCefDisplayArray) Add(value ICefDisplay) {
	m.values = append(m.values, value)
	m.count++
	m.instance = unsafePointer(m.values[0].Instance())
}

func (m *TCefDisplayArray) Set(value []ICefDisplay) {
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
