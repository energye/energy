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

// EventCallback IPC context callback
type EventCallback func(context IContext)

type ICallback interface {
	Invoke(context IContext) interface{}
}

// Callback IPC Listening callback function
type Callback struct {
	callback EventCallback
}

// Invoke event function
func (m *Callback) Invoke(context IContext) interface{} {
	if m.callback != nil {
		m.callback(context)
		return context.getResult()
	}
	return nil
}

func New(callback EventCallback) ICallback {
	return &Callback{callback: callback}
}
