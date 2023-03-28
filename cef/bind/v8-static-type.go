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
//
// 字段值仅主进程有效, 非主进程字段值为默认值
package bind

// Bind 绑定Go类型
//
// 变量和函数绑定, 在Go中定义的字段绑定到JS中, 在Go中定义的函数导出到JS
//
// 支持类型 String = string , Integer = (int8 ~ uint64) , Double = (float32, float64), Boolean = bool, Function = func, Object = struct | map,  Array = Slice
func Bind(name string, bind any) error {
	return nil
}
