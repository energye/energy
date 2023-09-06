//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package tools

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"strings"
)

type Registry struct {
	read registry.Key
	set  registry.Key
}

func NewRegistryCurrentUser() *Registry {
	read, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.READ)
	if err != nil {
		fmt.Println("open read", err)
		return nil
	}
	set, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.SET_VALUE)
	if err != nil {
		fmt.Println("open set", err)
		return nil
	}
	return &Registry{read: read, set: set}
}

// Read 读取
func (m *Registry) Read(name string) (string, error) {
	v, _, err := m.read.GetStringValue(name)
	return v, err
}

// Set 覆盖
func (m *Registry) Set(name, value string) error {
	return m.set.SetExpandStringValue(name, value)
	//return m.set.SetStringValue(name, value)
}

// Append 追加
func (m *Registry) Append(name, value string) error {
	if v, err := m.Read(name); err == nil {
		vals := strings.Split(v, ";")
		for _, val := range vals {
			val = strings.TrimSpace(val)
			if val == value { // 如果已经存在,就不添加了
				return nil
			}
		}
		if v[len(v)-1] == ';' {
			v += value
		} else {
			v += ";" + value
		}
		return m.Set(name, v)
	} else {
		return m.Set(name, value)
	}
}

// DeleteValue 删除指定key name
func (m *Registry) DeleteValue(name string) error {
	return m.set.DeleteValue(name)
}

// Close 关闭
func (m *Registry) Close() {
	if m.read != 0 {
		_ = m.read.Close()
	}
	if m.set != 0 {
		_ = m.set.Close()
	}
}
