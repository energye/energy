//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 导入动态链接库函数表

package imports

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type Table struct {
	name string
	addr ProcAddr
}

func NewTable(name string, addr ProcAddr) *Table {
	r := &Table{}
	r.name = name
	r.addr = addr
	return r
}

func (m *Table) Name() string {
	return m.name
}

func (m *Table) Addr() ProcAddr {
	return m.addr
}

func internalGetImportFunc(uiLib DLL, table []*Table, index int) ProcAddr {
	item := table[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.addr))) == nil {
		var err error
		item.addr, err = uiLib.GetProcAddr(item.name)
		if err != nil {
			fmt.Println(err, item.name)
			return 0
		}
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&table[index].addr)), unsafe.Pointer(item.addr))
	}
	return item.addr
}
