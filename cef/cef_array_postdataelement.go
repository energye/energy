//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefPostDataElementArray = array of ICefPostDataElement
type ICefPostDataElementArray interface {
	Instance() uintptr
	Get(index int) ICefPostDataElement
	Size() int
	Free()
	Add(value ICefPostDataElement)
	Set(value []ICefPostDataElement)
}

// TCefPostDataElementArray = array of ICefPostDataElement
type TCefPostDataElementArray struct {
	instance unsafePointer
	count    int
	values   []ICefPostDataElement
}

// PostDataElementArrayRef -> TCefPostDataElementArray
var PostDataElementArrayRef postDataElementArray

// postDataElementArray
type postDataElementArray uintptr

func (*postDataElementArray) New(count int, instance uintptr) ICefPostDataElementArray {
	return &TCefPostDataElementArray{
		count:    count,
		instance: unsafePointer(instance),
		values:   make([]ICefPostDataElement, count),
	}
}

func (m *TCefPostDataElementArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefPostDataElement
func (m *TCefPostDataElementArray) Get(index int) ICefPostDataElement {
	if m == nil {
		return nil
	}
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefPostDataElement(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

// Size 返回 ICefPostDataElement 数组长度
func (m *TCefPostDataElementArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefPostDataElementArray) Free() {
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

func (m *TCefPostDataElementArray) Add(value ICefPostDataElement) {
	m.values = append(m.values, value)
	m.count++
	m.instance = unsafePointer(m.values[0].Instance())
}

func (m *TCefPostDataElementArray) Set(value []ICefPostDataElement) {
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
