//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package commons

import (
	"bytes"
	"fmt"
	"github.com/energye/energy/logger"
	"os"
	"strings"
)

var Args = argsParse()

type _args struct {
	isMain      bool
	size        int
	value       map[string]PRCESS_TYPE
	commandLine string
}

func argsParse() *_args {
	var commandBuf = bytes.Buffer{}
	var args = &_args{value: make(map[string]PRCESS_TYPE, len(os.Args))}
	args.size = len(os.Args)
	for i, v := range os.Args {
		a := strings.Split(v, "=")
		if len(a) == 2 {
			args.value[strings.Replace(a[0], "--", "", 1)] = PRCESS_TYPE(a[1])
		}
		if i > 0 {
			commandBuf.WriteString("|,|")
		}
		commandBuf.WriteString(v)
	}
	args.commandLine = commandBuf.String()
	if v, ok := args.value["type"]; ok {
		args.isMain = v == PT_MAIN
	} else {
		args.isMain = true
	}
	return args
}

func (m *_args) CommandLine() string {
	return m.commandLine
}

func (m *_args) Size() int {
	return m.size

}

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
		return v == PT_RENDER
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

func (m *_args) IsSubprocess() bool {
	_, ok := m.value["type"]
	return ok
}

func (m *_args) SetArgs(name, value string) {
	m.value[name] = PRCESS_TYPE(value)
}
func (m *_args) Args(name string) string {
	if v, ok := m.value[name]; ok {
		return string(v)
	}
	return ""
}

func (m *_args) Print() {
	logger.Debug("command line:", m.size)
	for key, value := range m.value {
		fmt.Printf("_args[%v = %v]\n", key, value)
	}
}

type PRCESS_TYPE string

const (
	PT_MAIN     PRCESS_TYPE = ""
	PT_GPU                  = "gpu-process"
	PT_UTILITY              = "utility"
	PT_RENDER               = "render"
	PT_DEVTOOLS             = "devtools"
)
