//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 构建编译 energy 项目

package internal

var CmdBuild = &Command{
	UsageLine: "build",
	Short:     "build energy project",
	Long: `
	Building energy project
	.  Execute default command
`,
}

func init() {
	CmdBuild.Run = runBuild
}

func runBuild(c *CommandConfig) error {
	return nil
}
