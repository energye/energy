//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package callback

// IContext
//
//	Inter process IPC communication callback context
type IContext interface {
	Data() interface{}          //ArgumentList
	Result(data ...interface{}) //callback function return Result
}

// Context IPC Context
type Context struct {
	windowId uint32
	data     interface{}
	result   []interface{}
}

// NewContext create IPC event message
func NewContext(windowId uint32, data interface{}) IContext {
	ctx := &Context{
		windowId: windowId,
		data:     data,
		result:   nil,
	}
	return ctx
}

func (m *Context) Data() interface{} {
	return m.data
}

func (m *Context) Result(results ...interface{}) {
	m.result = results
}
