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
)

var CmdInit = &command.Command{
	UsageLine: "init -n [name]",
	Short:     "init energy project",
	Long: `
	-n initialized project name
	Initialize energy golang project
	.  Execute default command
`,
}

func init() {
	CmdInit.Run = runInit
}

func runInit(c *command.Config) error {
	m := c.Init
	if m.Name == "" {
		println("initialize project name:")
		fmt.Scan(&m.Name)
	}
	fmt.Println("name:", m.Name)
	return nil
}
