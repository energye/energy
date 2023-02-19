//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

var (
	_bind_id  = 10000
	_event_id = 10000
)

func __bind_id() int {
	_bind_id++
	return _bind_id
}

func __idReset() {
	_bind_id = 10000
}

func __event_id() int {
	_event_id++
	return _event_id
}
