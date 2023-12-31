//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// notice is a cross-platform system notification library
// notice 是一个跨平台的系统通知库
package notice

import (
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"
)

var uniqueID string

// Notification 表示可发送到操作系统的用户通知
type Notification struct {
	Title, Content string
	Icon           []byte
	iconExt        string
	Timeout        int32
}

func (m *Notification) SetTimeout(v int32) *Notification {
	if v >= 0 {
		m.Timeout = v
	} else {
		m.Timeout = 0
	}
	return m
}

func (m *Notification) SetIcon(iconResourcePath string) *Notification {
	if tools.IsExist(iconResourcePath) {
		data, err := ioutil.ReadFile(iconResourcePath)
		if err == nil {
			m.Icon = data
			m.iconExt = filepath.Ext(iconResourcePath)
		}
	}
	return m
}

func (m *Notification) SetIconFS(iconResourcePath string) *Notification {
	data, err := emfs.GetResources(iconResourcePath)
	if err == nil {
		m.Icon = data
		m.iconExt = filepath.Ext(iconResourcePath)
	}
	return m
}

// NewNotification 创建一个通知，可以传递给SendNotification
func NewNotification(title, content string) *Notification {
	return &Notification{Title: title, Content: content}
}

func UniqueID() string {
	if uniqueID != "" {
		return uniqueID
	}
	uniqueID = "energy-id-" + strconv.FormatInt(time.Now().Unix(), 10) // This is a fake unique - it just has to not be reused...
	return uniqueID
}

func SetUniqueID(v string) {
	uniqueID = v
}
