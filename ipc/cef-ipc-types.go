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

import (
	"github.com/energye/energy/consts"
	"github.com/energye/energy/types"
)

type IArrayValue interface {
	Size() uint32
	GetType(index types.NativeUInt) consts.TCefValueType
	GetBool(index types.NativeUInt) bool
	GetInt(index types.NativeUInt) int32
	GetDouble(index types.NativeUInt) (result float64)
	GetString(index types.NativeUInt) string
	GetIValue(index types.NativeUInt) IValue
	GetIBinary(index types.NativeUInt) IBinaryValue
	GetIObject(index types.NativeUInt) IObjectValue
	GetIArray(index types.NativeUInt) IArrayValue
}

type IBinaryValue interface {
	GetSize() uint32
	GetData(buffer []byte, dataOffset types.NativeUInt) uint32
}

type IObjectValue interface {
	GetSize() uint32
	GetType(key string) consts.TCefValueType
	GetBool(key string) bool
	GetInt(key string) int32
	GetDouble(key string) (result float64)
	GetString(key string) string
	GetIValue(key string) IValue
	GetIBinary(key string) IBinaryValue
	GetIObject(key string) IObjectValue
	GetIArray(key string) IArrayValue
}

type IValue interface {
	GetType() consts.TCefValueType
	GetBool() bool
	GetInt() int32
	GetDouble() (result float64)
	GetString() string
	GetIBinary() IBinaryValue
	GetIObject() IObjectValue
	GetIArray() IArrayValue
}
