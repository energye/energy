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

import "encoding/json"

type ProcessMessage struct {
	Name string      `json:"n"`
	Data interface{} `json:"d"`
	Id   int         `json:"i"`
}

func (m *ProcessMessage) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}
