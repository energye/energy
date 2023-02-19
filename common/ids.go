//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package common

import "sync"

// IPC ID生成
type IPCIDGen struct {
	_id   int32
	mutex sync.Mutex
}
type CliID struct {
	IPCIDGen
}

// 消息ID生成
type MsgID struct {
	IPCIDGen
}

func (m *IPCIDGen) New() int32 {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m._id++
	return m._id
}
