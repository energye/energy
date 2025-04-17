//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package argument

import (
	goJSON "encoding/json"
	"github.com/cyber-xxm/energy/v2/pkgs/json"
	"reflect"
)

// IList
//
//	IPC Argument List Interface
type IList interface {
	MessageId() int32     // messageId
	BrowserId() int32     // browserId
	GetName() string      // messageName
	GetEventName() string // eventName
	GetData() interface{} // messageData
	JSON() json.JSON      // messageData convert JSON
	Bytes() []byte        // messageData convert []byte
	Reset()               // free
}

// List
//
//	IPC Argument List
type List struct {
	Id        int32       `json:"id"`        // messageId
	BId       int32       `json:"bid"`       // browserId
	Name      string      `json:"name"`      // messageName
	EventName string      `json:"eventName"` // eventName
	Data      interface{} `json:"data"`      // messageData
	jsonData  json.JSON   `json:"-"`
	bytesData []byte      `json:"-"`
}

func UnList(data []byte) IList {
	if data == nil {
		return nil
	}
	var v = &List{}
	if err := goJSON.Unmarshal(data, v); err == nil {
		return v
	}
	return nil
}

func (m *List) MessageId() int32 {
	return m.Id
}

func (m *List) BrowserId() int32 {
	return m.BId
}

func (m *List) GetName() string {
	return m.Name
}

func (m *List) GetEventName() string {
	return m.EventName
}

func (m *List) GetData() interface{} {
	return m.Data
}

func (m *List) JSON() json.JSON {
	if m.jsonData != nil {
		return m.jsonData
	}
	if m.Data == nil {
		return nil
	}
	switch m.Data.(type) {
	case []byte:
		m.jsonData = json.NewJSON(m.Data.([]byte))
	case string:
		m.jsonData = json.NewJSON([]byte(m.Data.(string)))
	case json.JsonData:
		v := m.Data.(json.JsonData)
		m.jsonData = &v
	case *json.JsonData:
		m.jsonData = m.Data.(*json.JsonData)
	case json.JSON:
		m.jsonData = m.Data.(json.JSON)
	case json.JSONObject:
		m.jsonData = m.Data.(json.JSONObject)
	case json.JSONArray:
		m.jsonData = m.Data.(json.JSONArray)
	}
	if m.jsonData != nil {
		return m.jsonData
	}
	rv := reflect.ValueOf(m.Data)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		kind = rv.Elem().Kind()
	}
	if kind == reflect.Slice || kind == reflect.Array {
		if v := json.NewJSONArray(m.Data); v != nil {
			m.jsonData = v.JSONArray()
		}
	} else if kind == reflect.Map || kind == reflect.Struct {
		if v := json.NewJSONObject(m.Data); v != nil {
			m.jsonData = v.JSONObject()
		}
	}
	return m.jsonData
}

func (m *List) Bytes() []byte {
	if m.bytesData != nil {
		return m.bytesData
	}
	if byt, err := goJSON.Marshal(m); err == nil {
		m.bytesData = byt
		return byt
	}
	return nil
}

func (m *List) Reset() {
	m.Id = 0
	m.Name = ""
	m.Data = nil
	if m.jsonData != nil {
		m.jsonData.Free()
		m.jsonData = nil
	}
	if m.bytesData != nil {
		m.bytesData = nil
	}
}
