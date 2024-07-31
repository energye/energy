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
	"encoding/json"
	"sync/atomic"
)

// callback function id
var (
	goExecutionID uint32
)

// MessageType IPC Message type
type MessageType uint8

const (
	MT_READY MessageType = iota + 1
	MT_GO_EMIT
	MT_JS_EMIT
	MT_GO_EMIT_CALLBACK
	MT_JS_EMIT_CALLBACK
	MT_DRAG_MOVE
	MT_DRAG_DOWN
	MT_DRAG_UP
	MT_DRAG_DBLCLICK
)

type ProcessMessage struct {
	Type MessageType `json:"t"` // MessageType, Used to distinguish message types, this constant is defined the same in GO and JS
	Name string      `json:"n"` // Message name
	Data interface{} `json:"d"` // Message payload
	Id   uint32      `json:"i"` // Not 0 (js or go) ipc.emit has callback function
}

func (m *ProcessMessage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// NextExecutionID
func NextExecutionID() uint32 {
	atomic.AddUint32(&goExecutionID, 1)
	return goExecutionID
}

// ResetExecutionID
func ResetExecutionID() {
	atomic.StoreUint32(&goExecutionID, 0)
}

func CheckIPCMessage(t MessageType) bool {
	return t == MT_GO_EMIT || t == MT_JS_EMIT ||
		t == MT_GO_EMIT_CALLBACK || t == MT_JS_EMIT_CALLBACK
}

func CheckDragMessage(t MessageType) bool {
	return t == MT_DRAG_MOVE || t == MT_DRAG_DOWN || t == MT_DRAG_UP || t == MT_DRAG_DBLCLICK
}
