//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import (
	"github.com/cyber-xxm/energy/v2/consts"
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
	Free()
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
	Free()
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
	Free()
}

type IV8ValueKeys interface {
	Count() int
	Get(index int) string
	Free()
}

// OnType listening type
type OnType int8

const (
	OtMain OnType = iota // Only the main process
	OtSub                // Only the sub process
	OtAll                // All process
)

// Mode IPC mode of the browser process
type Mode = int8

const (
	// MSync Synchronization, the default way CEF is used
	//  In JS, ipc.emit triggers the Go event and processes long-term tasks. The window will remain frozen until the task processing is completed.
	MSync Mode = iota
	// MAsync
	//  Asynchronous, using coroutines, coroutines (within the event) cannot be debugged, there are no other unforeseen problems found so far.
	//  异步 (Asynchronous): Refers to an approach where operations can continue without waiting for the previous operations' completion.
	//  使用协程 (using coroutines): Indicates the implementation or employment of coroutines, which are a way to manage the execution flow in a non-preemptive manner.
	//  协程(事件内)无法Debug (coroutines (within the event) cannot be debugged): Points out the inability to debug coroutines when they are inside an event.
	//  暂未发现其它无法预料的问题 (there are no other unforeseen problems found so far): Indicates that, at the time of the statement, no other unforeseen issues have been encountered or identified.
	//
	// 使用场景 (Usage scenarios):
	//
	//  Only applicable when using JS ipc.emit to trigger events.
	//  Recommended for use in the Go UI main thread when performing long-duration tasks, otherwise it will freeze the UI window.
	MAsync
)

// OnOptions Listening options
type OnOptions struct {
	OnType OnType // Listening type, default main process
	Mode   Mode   // IPC emit mode of the browser process
}
