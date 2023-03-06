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
)

type IArrayValue interface {
	Size() uint32
	GetType(index uint32) consts.TCefValueType
	GetBool(index uint32) bool
	GetInt(index uint32) int32
	GetDouble(index uint32) (result float64)
	GetString(index uint32) string
	GetIValue(index uint32) IValue
	GetIBinary(index uint32) IBinaryValue
	GetIObject(index uint32) IObjectValue
	GetIArray(index uint32) IArrayValue
}

type IBinaryValue interface {
	GetSize() uint32
	GetData(buffer []byte, dataOffset uint32) uint32
}

type IObjectValue interface {
	Size() uint32
	GetType(key string) consts.TCefValueType
	GetBool(key string) bool
	GetInt(key string) int32
	GetDouble(key string) (result float64)
	GetString(key string) string
	GetIKeys() IV8ValueKeys
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

type IV8ValueKeys interface {
	Count() int
	Get(index int) string
}
