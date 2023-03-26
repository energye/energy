//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package bind

import (
	"container/list"
	"fmt"
	"sync"
	"unsafe"
)

var bind = &v8bind{hasFieldCollection: make(map[string]uintptr), fieldCollection: list.New()}

type v8bind struct {
	hasFieldCollection map[string]uintptr
	fieldCollection    *list.List
	lock               sync.Mutex
}

// Set 添加或修改
//
//	参数
//		 name: 唯一字段名, 重复将被覆盖
//		value: 值
func (m *v8bind) Set(name string, value JSValue) {
	if id, ok := m.hasFieldCollection[name]; ok {
		m.Remove(id)
		id = m.Add(value)
		value.setId(id)
		m.hasFieldCollection[name] = id
	} else {
		id = m.Add(value)
		value.setId(id)
		m.hasFieldCollection[name] = id
	}
	//m.fieldCollection[name] = value
}

func (m *v8bind) GetJSValue(id uintptr) JSValue {
	if v := m.Get(id); v != nil {
		return v.Value.(JSValue)
	}
	return nil
}

func (m *v8bind) Add(value JSValue) uintptr {
	return uintptr(unsafe.Pointer(m.fieldCollection.PushBack(value)))
}

func (m *v8bind) Get(id uintptr) *list.Element {
	return (*list.Element)(unsafe.Pointer(id))
}

func (m *v8bind) Remove(id uintptr) {
	if v := m.Get(id); v != nil {
		m.fieldCollection.Remove(v)
	}
}

func (m *v8bind) Binds() map[string]JSValue {
	//return m.fieldCollection
	return nil
}

func GetBinds(fn func(binds map[string]JSValue)) {
	//fn(bind.fieldCollection)
}

func Test() {
	stringKey := NewString("stringKey", "字符串值")
	fmt.Println("stringKey", stringKey, stringKey.IsString())
	stringKey = NewString("stringKey", "字符串值")
}
