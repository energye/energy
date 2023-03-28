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
	"github.com/energye/energy/pkgs/json"
	"sync"
	"unsafe"
)

var bind = &v8bind{hasFieldCollection: make(map[string]uintptr), fieldCollection: list.New()}

type v8bind struct {
	hasFieldCollection  map[string]uintptr
	fieldCollection     *list.List
	setLock, removeLock sync.Mutex
}

// Set 添加或修改
//
//	参数
//		 name: 唯一字段名, 重复将被覆盖
//		value: 值
func (m *v8bind) Set(name string, value JSValue) {
	m.setLock.Lock()
	defer m.setLock.Unlock()
	if id, ok := m.hasFieldCollection[name]; ok {
		if value.Id() != id {
			// remove old id
			old := m.Remove(id)
			// gen add value and return new id
			id = m.Add(value)
			value.setId(id)
			//update name id
			m.hasFieldCollection[name] = id
			switch old.(type) {
			case JSValue:
				//old value set new id
				old.(JSValue).setId(id)
			}
		}
	} else {
		// create set new value and return new id
		id = m.Add(value)
		value.setId(id)
		m.hasFieldCollection[name] = id
	}
}

// GetJSValue 返回 JSValue
func (m *v8bind) GetJSValue(id uintptr) JSValue {
	if v := m.Get(id); v != nil {
		return v.Value.(JSValue)
	}
	return nil
}

// Add 添加 JSValue 并返回 id
func (m *v8bind) Add(value JSValue) uintptr {
	return uintptr(unsafe.Pointer(m.fieldCollection.PushBack(value)))
}

// Get list element
func (m *v8bind) Get(id uintptr) *list.Element {
	return (*list.Element)(unsafe.Pointer(id))
}

// Remove 删除
func (m *v8bind) Remove(id uintptr) any {
	m.removeLock.Lock()
	defer m.removeLock.Unlock()
	if v := m.Get(id); v != nil {
		r := m.fieldCollection.Remove(v)
		v.Value = nil
		return r
	}
	return nil
}

func GetBinds(fn func(binds map[string]JSValue)) {
	//fn(bind.fieldCollection)
}

func Test() {
	//字段
	stringKey0 := NewString("stringKey", "字符串值0")
	fmt.Println("stringKey", stringKey0, stringKey0.Value())
	stringKey1 := NewString("stringKey", "字符串值1")
	integerKey := NewInteger("integerKey", 1000)
	fmt.Println("stringKey", stringKey0)
	fmt.Println("stringKey", stringKey1, stringKey1.Value())
	fmt.Println("integerKey", integerKey.Value())
	integerKey.SetValue("变成字符串")
	fmt.Println("integerKey", integerKey.AsString().Value())
	integerKey.SetValue(true)
	fmt.Println("integerKey", integerKey.AsBoolean().Value())
	boolField := integerKey.AsBoolean()
	fmt.Println("boolField", boolField.Value())
	fmt.Println("boolField", bind.GetJSValue(boolField.Id()).AsBoolean().Value())
	boolField.SetValue(false)
	fmt.Println("boolField", bind.GetJSValue(boolField.Id()).AsBoolean().Value())
	fmt.Println(bind.fieldCollection.Len())

	//函数
	funcKey := NewFunction("funcKey", func(in1 string) {
		fmt.Println("funcKey:", in1)
	})
	inArgument := json.NewJSONArray(nil)
	inArgument.Add("字符串参数")
	funcKey.Invoke(inArgument)

	funcKey.SetValue("函数变字符串")
	funcToString := funcKey.AsString()
	fmt.Println("funcToString:", funcToString.Value())

	// 对象
	fmt.Println("--------------")
	type objectDemo1 struct {
		Key1 string
		Key2 string
	}
	type objectDemo2 struct {
		Key1 string
		Key2 string
		Key3 int
	}
	type object struct {
		Key1 string
		Key2 string
		Key3 int
		Key4 float64
		Key5 bool
		Key6 *objectDemo1
		Key7 objectDemo2
		Key8 []string
	}
	var testObj = &object{
		Key1: "value1",
		Key2: "value2",
		Key3: 333,
		Key4: 555.3,
		Key5: true,
		//Key6: &objectDemo1{},
	}

	objectKey := NewObject(testObj)
	fmt.Println("objectKey:", objectKey.JSONString())
	objectKey.Set("Key1", "值1")
	objectKey.Set("Key2", "值2")
	objectKey.Set("Key3", 4444)
	objectKey.Set("Key4", 9999.99)
	objectKey.Set("Key5", false)
	objectKey.Set("Key6", &objectDemo1{Key1: "Key6值"})
	objectKey.Set("Key7", objectDemo2{Key1: "值值"})
	fmt.Println("objectKey:", objectKey.JSONString())
	objectKey1 := objectKey.Get("Key1")
	fmt.Println("objectKey1Name:", objectKey1.Name())
	objectKey1.SetValue("objectKey1设置新值 ")
	fmt.Println("objectKey1:", objectKey1.JSONString())
	objectKey6 := objectKey.Get("Key6")
	fmt.Println("objectKey:", objectKey.JSONString())
	objectKey6.SetValue(&objectDemo1{Key1: "objectKey6Key1"})
	fmt.Println("objectKey6:", objectKey6.JSONString())
	fmt.Println("objectKey:", objectKey.JSONString())
	objectKey.Set("Key8", []string{"v1", "v2"})
	objectKey8 := objectKey.Get("Key8")
	fmt.Println("objectKey8:", objectKey8.JSONString())
	objectKey8.AsArray()
	//object end
	fmt.Println("objectKey:", objectKey.JSONString())
	//object to string
	objectKey.SetValue("对象变成字符串")
	objectToString := objectKey.AsString()
	fmt.Println("objectToString:", objectToString.Value())
	fmt.Println("objectKey:", objectKey.JSONString())
	objectKey.SetValue(testObj)
	fmt.Println("objectKey:", objectKey.JSONString())
	bindObjectKey := bind.GetJSValue(objectKey.Id())
	fmt.Println("bindObjectKey", bindObjectKey.JSONString())

	objectKey.SetValue([]any{"字符串", 100001, 22222.333, true, testObj})
	fmt.Println("objectKey-to-array:", objectKey.JSONString())
	fmt.Println("objectKey-to-array:", objectKey.IsArray())

	// 数组
	var arrayFunc = func() string {
		fmt.Println("arrayFunc")
		return "调用函数 arrayFunc返回"
	}
	fmt.Println("arrayFunc", arrayFunc)
	arrayKey := NewArray("arrayKey", "字符串", 100001, 22222.333, true, testObj, arrayFunc)
	fmt.Println("arrayKey JSONString:", arrayKey.JSONString())
	fmt.Println("arrayKey index 1:", arrayKey.Get(1).AsInteger().Value())
	fmt.Println("arrayKey index 2:", arrayKey.Get(2).AsDouble().Value())
	fmt.Println("arrayKey index 4:", arrayKey.Get(4).AsObject().JSONString())
	fmt.Println("arrayKey index 5:", arrayKey.Get(5).AsFunction().Invoke(nil).GetStringByIndex(0))
	arrayKey.Set(1, "数字变字符串")
	fmt.Println("arrayKey index 1:", arrayKey.Get(1).AsString().Value())
	arrayIndex1 := arrayKey.Get(1)
	fmt.Println("arrayKey index 1:", arrayIndex1.AsString().Value())
	arrayIndex1 = bind.GetJSValue(arrayIndex1.Id())
	fmt.Println("arrayKey index 1:", arrayIndex1.AsString().Value())
	arrayKey.Add("添加一个字符串")

	//end
	fmt.Println("fieldCollection.Len():", bind.fieldCollection.Len(), len(bind.hasFieldCollection))
	for k, v := range bind.hasFieldCollection {
		jsv := bind.GetJSValue(v)
		fmt.Println("k:", k, "v:", v, "jsv:", jsv.Type(), v == jsv.Id())
	}

}

func TestBind() {

}
