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

// IKey int or string
type IKey interface {
	int | string
}

// IArguments IPC参数 数组形式
type IArguments[key IKey] interface {
	Size() int
	Type(key key) int
	String(key key) string
	Int(key key) int
	UInt(key key) uint
	Bytes(key key) []byte
	Float64(key key) float64
	Bool(key key) bool
	Get(key key) *Argument[key]
}

//["stringValue",345,true,2344.66,"字符串？",30344.66,{"stringField":"stringFieldValue","intField":1000,"arrayField":[100,200,"数组里的字符串",66996.99],"doubleField":999991.102,"booleanField":true},[100,200,"数组里的字符串",66996.99,{"stringField":"stringFieldValue","intField":1000,"arrayField":[100,200,"数组里的字符串",66996.99],"doubleField":999991.102,"booleanField":true},true,false],8888888889233,"null","undefined"]

type Argument[Key IKey] struct {
	T int
	V []any
}

func (m *Argument[IKey]) Size() int {
	return len(m.V)
}

// Type IKey index or name
func (m *Argument[IKey]) Type(key IKey) int {
	return 0
}

func (m *Argument[IKey]) String(key IKey) string {
	return ""
}
