//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package bind Energy has not been exported or added yet
package bind

import (
	"container/list"
	"github.com/energye/energy/v2/cef/process"
	"reflect"
	"sync"
	"unsafe"
)

var (
	isMainProcess bool
	isSubProcess  bool
	bind          = &V8bind{hasFieldCollection: make(map[string]uintptr), fieldCollection: list.New()}
)

func init() {
	isMainProcess = process.Args.IsMain()
	isSubProcess = process.Args.IsRender()
}

// V8bind
type V8bind struct {
	hasFieldCollection  map[string]uintptr
	fieldCollection     *list.List
	setLock, removeLock sync.Mutex
}

type Data struct {
	P string       `json:"P"`
	T reflect.Kind `json:"T"`
	V any          `json:"V"`
}

func (m *V8bind) HasSize() int {
	return len(m.hasFieldCollection)
}

func (m *V8bind) Size() int {
	return m.fieldCollection.Len()
}

func (m *V8bind) HasFieldCollection() map[string]uintptr {
	return m.hasFieldCollection
}

func (m *V8bind) FieldCollection() *list.List {
	return m.fieldCollection
}

// Set 添加或修改
//
//	参数
//		 name: 唯一字段名, 重复将被覆盖
//		value: 值
func (m *V8bind) Set(value JSValue) {
	m.setLock.Lock()
	defer m.setLock.Unlock()
	// create set new value and return new id
	if value.Id() == 0 {
		value.setId(m.Add(value))
	}

	//if id, ok := m.hasFieldCollection[value.Name()]; ok {
	//	if value.Id() != id {
	//		// remove old id
	//		old := m.Remove(id)
	//		// gen add value and return new id
	//		value.setId(m.Add(value))
	//		// update name id
	//		m.hasFieldCollection[value.Name()] = value.Id()
	//		switch old.(type) {
	//		case JSValue:
	//			// old value set new id
	//			old.(JSValue).setId(value.Id())
	//		}
	//	}
	//} else {
	//	// create set new value and return new id
	//	value.setId(m.Add(value))
	//	m.hasFieldCollection[value.Name()] = value.Id()
	//}
}

// GetJSValue 返回 JSValue
func (m *V8bind) GetJSValue(id uintptr) JSValue {
	if v := m.Get(id); v != nil {
		return v.Value.(JSValue)
	}
	return nil
}

func (m *V8bind) ElementToJSValue(item *list.Element) JSValue {
	if item != nil {
		r, ok := item.Value.(JSValue)
		if ok {
			return r
		}
	}
	return nil
}

// Add 添加 JSValue 并返回 id
func (m *V8bind) Add(value JSValue) uintptr {
	return uintptr(unsafe.Pointer(m.fieldCollection.PushBack(value)))
}

// Get list element
func (m *V8bind) Get(id uintptr) *list.Element {
	if id <= 0xFF {
		return nil
	}
	return (*list.Element)(unsafe.Pointer(id))
}

// Remove 删除
func (m *V8bind) Remove(id uintptr) any {
	m.removeLock.Lock()
	defer m.removeLock.Unlock()
	if v := m.Get(id); v != nil {
		r := m.fieldCollection.Remove(v)
		v.Value = nil
		return r
	}
	return nil
}

// GetBinds 获取绑定的字段
func GetBinds(fn func(bind *V8bind)) {
	fn(bind)
}
