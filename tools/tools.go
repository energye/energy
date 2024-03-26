//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package tools

import (
	"os"
	"reflect"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func MkdirAll(path string) {
	var err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

type StructTags struct {
	tag map[string]*Tags
}

type Tags struct {
	tags map[string]string
}

// str struct
// tags tag name
// NewFindStructTags(&struct, "tagName1","tagName2")
func NewFindStructTags(str interface{}, tags ...string) *StructTags {
	m := &StructTags{
		tag: make(map[string]*Tags),
	}
	m.initTags(str, tags...)
	return m
}

func (m *StructTags) initTags(str interface{}, tags ...string) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		var fieldName = t.Field(i).Name
		m.tag[fieldName] = &Tags{
			tags: make(map[string]string),
		}
		for _, tagName := range tags {
			var tag = t.Field(i).Tag
			var tagVal = tag.Get(tagName)
			if tagVal != "" {
				m.tag[fieldName].tags[tagName] = tagVal
			}
		}
	}
}

// GetFieldTag(结构里定义字段名区分大小写)
func (m *StructTags) GetTags(fieldName string) *Tags {
	return m.tag[fieldName]
}

// Get 获得标签值
func (m *Tags) Get(tagName string) string {
	return m.tags[tagName]
}

// Set 设置标签值
func (m *Tags) Set(tagName, value string) {
	m.tags[tagName] = value
}
