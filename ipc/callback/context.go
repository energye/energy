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
	Data() interface{}          // ArgumentList
	Result(data ...interface{}) // callback function return Result
	BrowserId() uint32          // browser id
}

type IOnCallbackContext interface {
	IContext
	GetResult() interface{} // callback function result
}

// Context IPC Context
type Context struct {
	browserId uint32
	data      interface{}
	result    []interface{}
}

// NewContext create IPC event message
func NewContext(browserId uint32, data interface{}) IOnCallbackContext {
	ctx := &Context{
		browserId: browserId,
		data:      data,
		result:    nil,
	}
	return ctx
}

func (m *Context) BrowserId() uint32 {
	return m.browserId
}

func (m *Context) Data() interface{} {
	return m.data
}

func (m *Context) Result(results ...interface{}) {
	m.result = results
}

func (m *Context) GetResult() interface{} {
	return m.result
}
