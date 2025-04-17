//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC
//
// event listeners
// event triggered

package ipc

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
)

// On
//
//	IPC GO 监听事件
//
// 参数
//
//	支持 JavaScript 对应 Go 的基本类型和复合类型
//	name: 事件名称
//	fn : 事件回调函数 EmitContextCallback 或 func(...) [result...] {}
//	options: 监听选项, 配置监听规则, 异步或同步规则
//
// 入参
//
//	基本类型: int(int8 ~ uint64), bool, float(float32、float64), string
//	复合类型: slice, map, struct
//	slice: 根据js实际类型定义, []interface{} | []interface{} | [][data type]
//	map: key 只能 string 类型, value 基本类型+复合类型
//	struct: 首字母大写, 字段类型匹配
//	    type ArgsStructDemo struct {
//	       Key1 string
//			  Key2 string
//			  Key3 string
//			  Key4 int
//			  Key5 float64
//			  Key6 bool
//			  Sub1  SubStructXXX
//			  Sub2  *SubStructXXX
//	    }
//
// 出参
//
//	fn 回调函数的出参与入参使用方式相同
func On(name string, fn interface{}, options ...types.OnOptions) {
	ipc.On(name, fn, options...)
}

// RemoveOn
// IPC GO 移除监听事件
func RemoveOn(name string) {
	ipc.RemoveOn(name)
}

// Emit
// IPC GO 中触发 Go | JS 监听的事件, 主窗口
//
// 参数
//
//		name: 监听的事件名
//	 []argument: 入参
//					基本类型: int(int8 ~ uint64), bool, float(float32、float64), string
//					复合类型: slice, map, struct
func Emit(name string, argument ...interface{}) bool {
	return ipc.Emit(name, argument...)
}

// EmitAndCallback
// IPC GO 中触发 Go | JS 监听的事件, 主窗口
//
// 参数
//
//		name: 监听的事件名
//	 []argument: 入参
//					基本类型: int(int8 ~ uint64), bool, float(float32、float64), string
//					复合类型: slice, map, struct
//	 callback: 回调函数, 接收返回值. 函数类型 EmitContextCallback 或 func(...) [result...] {}
func EmitAndCallback(name string, argument []interface{}, callback interface{}) bool {
	return ipc.EmitAndCallback(name, argument, callback)
}

// EmitTarget
// IPC GO 中触发指定目标 Go | JS 监听的事件
//
// 参数
//
//		name: 监听的事件名
//		target: 接收事件的目标
//	 []argument: 入参
//					基本类型: int(int8 ~ uint64), bool, float(float32、float64), string
//					复合类型: slice, map, struct
func EmitTarget(name string, target target.ITarget, argument ...interface{}) bool {
	return ipc.EmitTarget(name, target, argument...)
}

// EmitTargetAndCallback
// IPC GO 中触发指定目标 Go | JS 监听的事件
//
// 参数
//
//		name: 监听的事件名
//		target: 接收事件的目标
//	 []argument: 入参
//					基本类型: int(int8 ~ uint64), bool, float(float32、float64), string
//					复合类型: slice, map, struct
//	 callback: 回调函数, 接收返回值. 函数类型 EmitContextCallback 或 func(...) [result...] {}
func EmitTargetAndCallback(name string, target target.ITarget, argument []interface{}, callback interface{}) bool {
	return ipc.EmitTargetAndCallback(name, target, argument, callback)
}
