//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue 静态类型实现
//
// 静态类型是不可变类型
package bind

// Bind V8Value
//
// 变量和函数绑定, 在Go中定义的字段绑定到JS字段中, 在Go中定义的函数导出到JS
//
// 支持类型 String = string , Integer = int32 , Double = float64, Boolean = bool, Function = func, Objects = struct | map,  JSONArray = Slice
//
// 主进程和子进程
func Bind(name string, bind any) error {
	return nil
}
