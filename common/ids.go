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
