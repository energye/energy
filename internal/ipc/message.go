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
	MT_EVENT_GO_EMIT
	MT_EVENT_JS_EMIT
	MT_EVENT_GO_EMIT_CALLBACK
	MT_EVENT_JS_EMIT_CALLBACK
	MT_DRAG_MOVE
	MT_DRAG_DOWN
	MT_DRAG_UP
	MT_DRAG_DBLCLICK
	MT_DRAG_RESIZE
	MT_DRAG_BORDER_WMSZ
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
