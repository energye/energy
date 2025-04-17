//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 当前进程的命令行参数获取

package process

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/logger"
	"os"
	"strings"
)

// Args
// 当前进程命令行参数解析实例
var Args = parse()

type _args struct {
	isMain bool
	size   int
	value  map[string]PRCESS_TYPE
}

func parse() *_args {
	var args = &_args{value: make(map[string]PRCESS_TYPE, len(os.Args))}
	args.size = len(os.Args)
	for _, v := range os.Args {
		a := strings.Split(v, "=")
		if len(a) == 2 {
			args.value[strings.Replace(a[0], "--", "", 1)] = PRCESS_TYPE(a[1])
		}
	}
	if v, ok := args.value["type"]; ok {
		args.isMain = v == PT_MAIN
	} else {
		args.isMain = true
	}
	return args
}

func (m *_args) Size() int {
	return m.size
}

// ProcessType
//
//	返回进程类型
func (m *_args) ProcessType() PRCESS_TYPE {
	if v, ok := m.value["type"]; ok {
		return v
	}
	return PT_MAIN
}

func (m *_args) IsMain() bool {
	return m.isMain
}

func (m *_args) IsRender() bool {
	if v, ok := m.value["type"]; ok {
		return v == PT_RENDERER
	}
	return false
}

func (m *_args) IsGPU() bool {
	if v, ok := m.value["type"]; ok {
		return v == PT_GPU
	}
	return false
}

func (m *_args) IsUtility() bool {
	if v, ok := m.value["type"]; ok {
		return v == PT_UTILITY
	}
	return false
}

// Args
//
//	根据name返回指定的命令行参数值
func (m *_args) Args(name string) string {
	if v, ok := m.value[name]; ok {
		return string(v)
	}
	return ""
}

// Print
//
//	打印到控制台
func (m *_args) Print() {
	logger.Debug("command line:", m.size)
	for key, value := range m.value {
		fmt.Printf("args = [%v = %v]\n", key, value)
	}
}

type PRCESS_TYPE string

const (
	PT_MAIN     PRCESS_TYPE = ""            // main
	PT_GPU                  = "gpu-process" // gpu
	PT_UTILITY              = "utility"     // utility
	PT_RENDERER             = "renderer"    // renderer
	PT_DEVTOOLS             = "devtools"    //
)
