//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 进程命令行
package cef

import (
	"bytes"
	"strings"
)

func (m *TCefCommandLine) AppendSwitch(name, value string) {
	m.commandLines[name] = value
}

func (m *TCefCommandLine) AppendArgument(argument string) {
	m.commandLines[argument] = ""
}

func (m *TCefCommandLine) toString() string {
	var str bytes.Buffer
	var i = 0
	var replace = func(s, old, new string) string {
		return strings.ReplaceAll(s, old, new)
	}
	for name, value := range m.commandLines {
		if i > 0 {
			str.WriteString(" ")
		}
		if value != "" {
			str.WriteString(replace(replace(name, " ", ""), "=", ""))
			str.WriteString("=")
			str.WriteString(replace(replace(value, " ", ""), "=", ""))
		} else {
			str.WriteString(replace(name, " ", ""))
		}
		i++
	}
	return str.String()
}
