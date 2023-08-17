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

var CmdCreate = &Command{
	UsageLine: "create",
	Short:     "create energy project",
	Long: `
	Initialize and create an energy golang project
	.  Execute default command
`,
}

func init() {
	CmdCreate.Run = runCreate
}

func runCreate(c *CommandConfig) error {
	return nil
}
