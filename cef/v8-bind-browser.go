//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于IPC的字段数据绑定 - 主进程
package cef

import (
	"fmt"
	iterBind "github.com/energye/energy/v2/cef/internal/bind"
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/pkgs/json"
	"reflect"
	"strings"
	"sync"
)

var bindBrowser *bindBrowserProcess

type bindBrowserProcess struct {
	bindObject   json.JSONObject
	bind         *iterBind.V8bind
	lock         sync.Mutex
	bindResultWG *sync.WaitGroup
}

func (m *bindBrowserProcess) objectFieldBind(object json.JSONObject, jsv iterBind.JSValue) {
	names := strings.Split(jsv.Name(), ".")
	if len(names) == 1 {
		if jsv.IsObject() {
			object.Set(jsv.Name(), json.NewJSONObject(nil))
		} else if jsv.IsArray() {
			object.Set(jsv.Name(), json.NewJSONArray(nil))
		} else if jsv.IsFunction() {
			fn := json.NewJSONObject(nil)
			fn.Set("t", int(jsv.Type()))
			fn.Set("p", jsv.Id())
			object.Set(jsv.Name(), fn)
		} else {
			object.Set(jsv.Name(), jsv.Id())
		}
	} else {
		name := names[len(names)-1]
		var pObject = object
		for i := 0; i < len(names)-1; i++ {
			if pObject.IsObject() {
				pObject = pObject.GetByKey(names[i])
			} else {
				pObject = pObject.JSONArray().GetByIndex(int(common.StrToInt32(names[i])))
			}
		}
		if jsv.IsObject() {
			if pObject.IsArray() {
				pObject.JSONArray().Add(json.NewJSONObject(nil))
			} else if pObject.IsObject() {
				pObject.Set(name, json.NewJSONObject(nil))
			}
		} else if jsv.IsArray() {
			if pObject.IsArray() {
				pObject.JSONArray().Add(json.NewJSONArray(nil))
			} else if pObject.IsObject() {
				pObject.Set(name, json.NewJSONArray(nil))
			}
		} else if jsv.IsFunction() {
			fn := json.NewJSONObject(nil)
			fn.Set("t", int(jsv.Type()))
			fn.Set("p", jsv.Id())
			if pObject.IsArray() {
				pObject.JSONArray().Add(fn)
			} else if pObject.IsObject() {
				pObject.Set(name, fn)
			}
		} else {
			if pObject.IsObject() {
				pObject.Set(name, jsv.Id())
			} else if pObject.IsArray() {
				pObject.JSONArray().Add(jsv.Id())
			}
		}
	}
}

// registerFieldBindEvent
//  绑定字段IPC事件初始化
//	生成绑定数据结构
//	注册事件
func (m *bindBrowserProcess) registerFieldBindEvent() {
	go iterBind.GetBinds(func(bind *iterBind.V8bind) {
		m.bind = bind
		var bindObject = json.NewJSONObject(nil)
		fields := bind.FieldCollection()
		for item := fields.Front(); item != nil; item = item.Next() {
			jsv := bind.ElementToJSValue(item)
			m.objectFieldBind(bindObject, jsv)
		}
		m.bindObject = bindObject
		if m.bindResultWG != nil {
			m.bindResultWG.Done()
		}
		fmt.Println("browser-bindObject-size:", m.bindObject.Size())
		fmt.Println("browser-bindObject-json:", m.bindObject.ToJSONString())
	})
	//	注册事件
	ipc.BrowserChan().AddCallback(func(channelId int64, argument argument.IList) bool {
		if argument != nil {
			name := argument.GetName()
			fmt.Println("name", name)
			if name == internalGetFieldBind {
				// 发送绑定字段数据结构结果
				go m.sendFieldBindResult(channelId)
				return true
			} else if name == internalGET {
				// 获取绑定值
				go m.getterBindValue(channelId, argument.GetData().(string))
				return true
			} else if name == internalSET {
				go m.setterBindValue(channelId, argument)
				return true
			} else if name == internalCALL {
				m.call(channelId, argument)
				return true
			}
		}
		return false
	})
}

// call
//  调用绑定函数
func (m *bindBrowserProcess) call(channelId int64, argument argument.IList) {
	if argument.JSON() != nil {
		fmt.Println("call-argument", argument.JSON().ToJSONString())
	}
}

// setterBindValue
//  设置绑定值
func (m *bindBrowserProcess) setterBindValue(channelId int64, argument argument.IList) {
	if argument.JSON() != nil {
		data := argument.JSON().JSONObject()
		ptr := data.GetStringByKey("P")
		p := uintptr(common.StrToInt64(ptr))
		t := reflect.Kind(data.GetIntByKey("T"))
		if jsv := m.bind.GetJSValue(p); jsv != nil {
			if t == jsv.Type() && jsv.BindType() == iterBind.BtStatic {
				jsv.SetValue(data.GetByKey("V").Data())
			} else if jsv.BindType() == iterBind.BtDynamic {
				jsv.SetValue(data.GetByKey("V").Data())
			}
		}
	}
}

// getterBindValue
//  获取绑定值
func (m *bindBrowserProcess) getterBindValue(channelId int64, ptr string) {
	fmt.Println("ptr:", ptr)
	p := uintptr(common.StrToInt64(ptr))
	jsv := m.bind.GetJSValue(p)
	//发送返回结果
	message := &argument.List{
		Name: internalGETResult,
	}
	if jsv != nil {
		bd := &iterBind.Data{
			P: ptr,
			T: jsv.Type(),
		}
		if data, ok := jsv.JSON().Data().(*reflect.Value); ok {
			bd.V = data.Interface()
		} else {
			bd.V = jsv.JSON().Data()
		}
		message.Data = bd
	}
	ipc.BrowserChan().IPC().Send(channelId, message.Bytes())
	message.Reset()
}

// sendFieldBindResult
//  发送绑定字段结果
func (m *bindBrowserProcess) sendFieldBindResult(channelId int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.bindObject == nil {
		m.bindResultWG = new(sync.WaitGroup)
		m.bindResultWG.Add(1)
		m.bindResultWG.Wait()
		m.bindResultWG = nil
	}
	message := &argument.List{
		Name: internalGetFieldBindResult,
		Data: m.bindObject.JsonData().ConvertToData(),
	}
	ipc.BrowserChan().IPC().Send(channelId, message.Bytes())
	message.Reset()
}
