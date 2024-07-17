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
var executionID uint32

// MessageType IPC Message type
type MessageType uint8

const (
	MT_GO_EMIT MessageType = iota + 1
	MT_JS_EMIT
	MT_GO_EMIT_CALLBACK
	MT_JS_EMIT_CALLBACK
)

type ProcessMessage struct {
	Type MessageType `json:"t"`
	Name string      `json:"n"`
	Data interface{} `json:"d"`
	Id   uint32      `json:"i"`
}

// ToJSON
func (m *ProcessMessage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// NextExecutionID
func NextExecutionID() uint32 {
	atomic.AddUint32(&executionID, 1)
	return executionID
}

// ResetExecutionID
func ResetExecutionID() {
	atomic.StoreUint32(&executionID, 0)
}

func CheckMessageType(t MessageType) bool {
	return t == MT_GO_EMIT || t == MT_JS_EMIT || t == MT_GO_EMIT_CALLBACK || t == MT_JS_EMIT_CALLBACK
}
