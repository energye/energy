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

import (
	"github.com/energye/energy/v2/cmd/internal/build"
	"github.com/energye/energy/v2/cmd/internal/command"
)

var CmdBuild = &command.Command{
	UsageLine: "build -p [path] -u [upx] --upxFlag --args --libemfs",
	Short:     "build energy project",
	Long: `
Building energy project:
  -p Project path, default current path. Can be configured in energy.json
  -u Set this parameter and install upx. Use upx to compress the execution file.
     --upxFlag: Upx command line parameters
  --args Set go build [args]
  --libemfs Built in dynamic libraries to executable files, Copy liblcl to the built-in directory every compilation
`,
}

func init() {
	CmdBuild.Run = runBuild
}

func runBuild(c *command.Config) error {
	return build.Build(c)
}
