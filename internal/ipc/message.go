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

var executionID uint32

type ProcessMessage struct {
	Name string      `json:"n"`
	Data interface{} `json:"d"`
	Id   uint32      `json:"i"`
}

func (m *ProcessMessage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func NextExecutionID() uint32 {
	atomic.AddUint32(&executionID, 1)
	return executionID
}

func ResetExecutionID() {
	atomic.StoreUint32(&executionID, 0)
}
