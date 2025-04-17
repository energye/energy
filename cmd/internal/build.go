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
	"github.com/cyber-xxm/energy/v2/cmd/internal/build"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
)

var CmdBuild = &command.Command{
	UsageLine: "build -p [path] -u [upx] -o [out] --os --arch --upxFlag --libemfs --buildargs",
	Short:     "Build ENERGY PROJECT",
	Long: `
Building ENERGY PROJECT:
  -p Project path, default current path. Can be configured in energy_[os].json
  -u Set this parameter and install upx. Use upx to compress the execution file.
     --upxFlag: Upx command line parameters
  -o Build out file path
  --os Build OS for windows | darwin | linux
  --arch Build ARCH for 386 | amd64 | arm | arm64
  --libemfs Built in dynamic libraries to executable files, Copy liblcl to the built-in directory every compilation
  --args go build [args]. Parameter tags should be in the last position, and the parameters following them will be passed as build parameters.
     energy build -p=/app/root/path --buildargs -tags prod,other1,other2 -ldflags="-X 'main.config=value'"
`,
}

func init() {
	CmdBuild.Run = runBuild
}

func runBuild(c *command.Config) error {
	return build.Build(c)
}
