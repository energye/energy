//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package ipc

const (
	Ln_IPC_GoEmitJS         = "IPCGoEmitJS"                   //Go执行Js on监听
	Ln_GET_BIND_FIELD_VALUE = "internal_GET_BIND_FIELD_VALUE" //browse进程监听获取字段值
	Ln_SET_BIND_FIELD_VALUE = "internal_SET_BIND_FIELD_VALUE" //browse进程监听设置字段值
	Ln_EXECUTE_BIND_FUNC    = "internal_EXECUTE_BIND_FUNC"    //browse进程监听执行绑定函数
	Ln_onConnectEvent       = "connect"
)

func InternalIPCNameCheck(name string) bool {
	if name == Ln_IPC_GoEmitJS || name == Ln_GET_BIND_FIELD_VALUE || name == Ln_SET_BIND_FIELD_VALUE || name == Ln_EXECUTE_BIND_FUNC || name == Ln_onConnectEvent {
		return true
	}
	return false
}
