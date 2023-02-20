//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy Id生成器
package cef

var (
	_bind_id  = 10000
	_event_id = 10000
)

// __bind_id 绑定变量ID
func __bind_id() int {
	_bind_id++
	return _bind_id
}

// __idReset 重置ID
func __idReset() {
	_bind_id = 10000
}

// __event_id 事件ID
func __event_id() int {
	_event_id++
	return _event_id
}
