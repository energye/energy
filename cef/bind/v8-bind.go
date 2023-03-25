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

import "sync"

var bind = &v8bind{fieldCollection: make(map[string]JSValue)}

type v8bind struct {
	fieldCollection map[string]JSValue
	lock            sync.Mutex
}

// set 添加或修改
func (m *v8bind) set(name string, value JSValue) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.fieldCollection[name] = value
}
