//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package bind

var bind = &v8bind{fieldCollection: make(map[string]JSValue)}

type v8bind struct {
	fieldCollection map[string]JSValue
}

func (m *v8bind) Add(name string, value JSValue) {
	m.fieldCollection[name] = value
}
