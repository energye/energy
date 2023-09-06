//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 创建 energy 项目

package internal

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/initialize"
	"strings"
)

var CmdInit = &command.Command{
	UsageLine: "init -n [name]",
	Short:     "Initialized energy project",
	Long: `
	-n Initialized project name
	Initialize energy golang project
	.  Execute default command
`,
}

func init() {
	CmdInit.Run = runInit
}

func runInit(c *command.Config) error {
	m := &c.Init
	if strings.TrimSpace(m.Name) == "" {
		for strings.TrimSpace(m.Name) == "" {
			print("Project Name: ")
			fmt.Scan(&m.Name)
			println()
		}
	}
	if strings.TrimSpace(m.ResLoad) == "" {
		println("Resource loading method, default 1 HTTP")
		print("1: HTTP\n2: Local Load\n  Number: ")
		fmt.Scan(&m.ResLoad)
		println()
	}
	m.ResLoad = strings.TrimSpace(m.ResLoad)
	if m.ResLoad == "" || (m.ResLoad != "1" && m.ResLoad != "2") {
		m.ResLoad = "1"
	}
	return initialize.InitEnergyProject(c)
}
