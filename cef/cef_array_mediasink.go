//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefMediaSinkArray = array of ICefMediaSink
type ICefMediaSinkArray interface {
	Instance() uintptr
	Get(index int) ICefMediaSink
	Size() int
	Free()
	Add(value ICefMediaSink)
	Set(value []ICefMediaSink)
}

// TCefMediaSinkArray = array of ICefMediaSink
type TCefMediaSinkArray struct {
	instance unsafePointer
	count    int
	values   []ICefMediaSink
}

// MediaSinkArrayRef -> TCefMediaSinkArray
var MediaSinkArrayRef mediaSinkArray

// mediaSinkArray
type mediaSinkArray uintptr

func (*mediaSinkArray) New(count int, instance uintptr) ICefMediaSinkArray {
	return &TCefMediaSinkArray{
		instance: unsafePointer(instance),
		count:    count,
		values:   make([]ICefMediaSink, count),
	}
}

func (m *TCefMediaSinkArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefMediaSink
func (m *TCefMediaSinkArray) Get(index int) ICefMediaSink {
	if m == nil {
		return nil
	}
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefMediaSink(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

// Size 返回 ICefMediaSink 数组长度
func (m *TCefMediaSinkArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefMediaSinkArray) Free() {
	if m == nil {
		return
	}
	if m.values != nil {
		for i, v := range m.values {
			if v != nil && v.Instance() != 0 {
				v.Free()
				m.values[i] = nil
			}
		}
		m.values = nil
	}
	m.instance = nil
	m.count = 0
}

func (m *TCefMediaSinkArray) Add(value ICefMediaSink) {
	m.values = append(m.values, value)
	m.count++
	m.instance = unsafePointer(m.values[0].Instance())
}

func (m *TCefMediaSinkArray) Set(value []ICefMediaSink) {
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
