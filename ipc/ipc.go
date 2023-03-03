//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ipc

import "github.com/energye/energy/types"

func On(name string, context IContext) {

}

func Emit(name string) {

}

type ArgumentList interface {
	Size() uint32
	GetBool(index types.NativeUInt) bool
	GetInt(index types.NativeUInt) int32
	GetDouble(index types.NativeUInt) (result float64)
	GetString(index types.NativeUInt) string
}

type ICefBinaryValue interface {
}

type IResult interface {
}

// IContext 进程间IPC通信回调上下文
type IContext interface {
	ArgumentList() ArgumentList
	Result() IResult
}
