//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于IPC的字段数据绑定 - 渲染(子)进程
package cef

import (
	"fmt"
	iterBind "github.com/energye/energy/v2/cef/internal/bind"
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/pkgs/json"
	"reflect"
	"sync"
)

var bindRender *bindRenderProcess

type bindRenderProcess struct {
	isInitBindIPC bool
	handler       *ICefV8Handler
	syncChan      *ipc.SyncChan
}

func (m *bindRenderProcess) registerContextEvent() {
	if m.isInitBindIPC {
		return
	}
	m.isInitBindIPC = true
	ipc.RenderChan().AddCallback(func(channelId int64, argument argument.IList) bool {
		if argument != nil {
			//messageJSON := data.JSONObject()
			//messageId := messageJSON.GetIntByKey(ipc_id)// messageId: 同步永远是1
			//name := messageJSON.GetStringByKey(ipc_name)
			//argumentList := messageJSON.GetArrayByKey(ipc_argumentList)
			//if name == internalIPCJSExecuteGoSyncEventReplay {
			//	return true
			//}
		}
		return false
	})
}

// registerFieldBindResultEvent
//  webkit 初始化获取并创建绑定字段和事件
func (m *bindRenderProcess) registerFieldBindResultEvent() {
	var (
		bindObject json.JSONObject
		waitResult = new(sync.WaitGroup)
	)
	waitResult.Add(1)
	// 注册绑定字段返回相关事件
	ipc.RenderChan().AddCallback(func(channelId int64, argument argument.IList) bool {
		if argument == nil {
			return false
		}
		name := argument.GetName()
		// 获取绑定字段数据结构消息
		if name == internalGetFieldBindResult { // webkit
			if argument.JSON() != nil {
				bindObject = argument.JSON().JSONObject()
			} else {
				bindObject = json.NewJSONObject(nil)
			}
			waitResult.Done() //返回结果
			return true
		} else if name == internalGETResult { // getter
			if argument.JSON() != nil {
				m.syncChan.ResultSyncChan <- argument.JSON().JSONObject()
			} else {
				m.syncChan.ResultSyncChan <- json.NewJSONObject(nil)
			}
			return true
		} else if name == internalSETResult { // setter

			return true
		}
		return false
	})
	// 获取绑定字段数据结构
	message := &argument.List{
		Id:   1,
		BId:  ipc.RenderChan().BrowserId(),
		Name: internalGetFieldBind,
	}

	//发送数据到主进程
	ipc.RenderChan().IPC().Send(message.Bytes())
	message.Reset()
	// TODO VF 组件有问题
	waitResult.Wait() //等待返回
	fmt.Println("webkit-bindObject-size::", bindObject.Size())
	fmt.Println("webkit-bindObject-ToJSONString::", bindObject.ToJSONString())
	m.webKitMakeBind(bindObject) // render webkit bind make
	bindObject.Free()
}

// webKitMakeBind
func (m *bindRenderProcess) webKitMakeBind(bindObject json.JSONObject) {
	m.handler = V8HandlerRef.New()
	m.handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		if ipc.RenderChan() != nil && arguments.Size() > 0 {
			//renderIPC.ipc.Send()
			fmt.Println("v8Handler.Execute", name, ipc.RenderChan(), arguments.Size(), arguments.Get(0).GetStringValue())
			//获取绑定字段值
			if name == internalGET {
				m.getterBindValue(arguments, retVal)
				return true
			} else if name == internalSET {
				m.setterBindValue(arguments, retVal)
				return true
			} else if name == internalCALL {
				m.call(arguments, retVal)
				return true
			}
		}
		return false
	})

	// 注册绑定JS扩展
	m.registerBindExtension(bindObject)
}

// call
//  调用绑定函数
func (m *bindRenderProcess) call(arguments *TCefV8ValueArray, retVal *ResultV8Value) {
	if arguments == nil || arguments.Size() == 0 {
		return
	}
	ptr := arguments.Get(0)
	defer ptr.Free()
	if !ptr.IsString() {
		return
	}
	fmt.Println("call-ptr:", ptr.GetStringValue(), "size:", arguments.Size())
	if arguments.Size() == 2 {
		args := arguments.Get(1)
		defer args.Free()
		if !args.IsArray() {
			return
		}
		fmt.Println("call-ptr:", ptr.GetStringValue(), "GetArrayLength:", args.GetArrayLength())

	}
}

// setterBindValue
//	设置绑定值
func (m *bindRenderProcess) setterBindValue(arguments *TCefV8ValueArray, retVal *ResultV8Value) {
	if arguments == nil || arguments.Size() != 2 {
		return
	}
	ptr := arguments.Get(0)
	defer ptr.Free()
	if !ptr.IsString() {
		return
	}
	newValue := arguments.Get(1)
	defer newValue.Free()
	bd := &iterBind.Data{}
	if newValue.IsString() {
		bd.T = reflect.String
		bd.V = newValue.GetStringValue()
	} else if newValue.IsInt() {
		bd.T = reflect.Int
		bd.V = newValue.GetIntValue()
	} else if newValue.IsUInt() {
		bd.T = reflect.Uint
		bd.V = newValue.GetUIntValue()
	} else if newValue.IsDouble() {
		bd.T = reflect.Float64
		bd.V = newValue.GetDoubleValue()
	} else if newValue.IsBool() {
		bd.T = reflect.Bool
		bd.V = newValue.GetBoolValue()
	} else {
		// 暂时不支持复合类型赋值
		return
	}
	bd.P = ptr.GetStringValue()
	message := &argument.List{
		Name: internalSET,
		BId:  ipc.RenderChan().BrowserId(),
		Data: bd,
	}
	fmt.Println("setterBindValue-ptr:", string(message.Bytes()))
	ipc.RenderChan().IPC().Send(message.Bytes())
	message.Reset()
}

// getterBindValue
//	获取绑定值
func (m *bindRenderProcess) getterBindValue(arguments *TCefV8ValueArray, retVal *ResultV8Value) {
	if arguments == nil || arguments.Size() != 1 {
		return
	}
	m.syncChan.DelayWaiting()
	ptr := arguments.Get(0)
	defer ptr.Free()
	if !ptr.IsString() {
		return
	}
	message := &argument.List{
		Name: internalGET,
		BId:  ipc.RenderChan().BrowserId(),
		Data: ptr.GetStringValue(),
	}
	ipc.RenderChan().IPC().Send(message.Bytes())
	message.Reset()
	data := <-m.syncChan.ResultSyncChan
	m.syncChan.Stop()
	var resultJSON json.JSONObject
	if data != nil {
		resultJSON = data.(json.JSONObject)
	}
	fmt.Println("获取返回结果:", resultJSON.ToJSONString(), resultJSON.GetStringByKey("V"))
	if resultJSON != nil && resultJSON.Size() > 0 {
		if v := ValueConvert.BindDataToV8Value(resultJSON); v != nil {
			retVal.SetResult(v)
			return
		}
	}
}
