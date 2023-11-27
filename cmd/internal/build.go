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
	UsageLine: "build -p [path] -u [upx] --upxFlag -d [dll] --tempDllFlag",
	Short:     "build energy project",
	Long: `
	Building energy project
	-p Project path, default current path. Can be configured in energy.json
	-u Set this parameter and install upx. Use upx to compress the execution file.
	  --upxFlag: Upx command line parameters
	-d TempDll Enable built-in liblcl build, 
	  --tempDllFlag [dll] TempDll parameters, default latest.
		windows:
			386: -tags="tempdll latest"
			amd64: -tags="tempdll latest"
		windows(Windows 7, 8/8.1 and Windows Server 2012):
			386: -tags="tempdll 109"
			amd64: -tags="tempdll 109"
		linux(gtk3):
			amd64: -tags="tempdll latest"
			arm64: -tags="tempdll latest"
		linux(gtk2):
			amd64: -tags="tempdll 106"
			arm64: -tags="tempdll 106"
		macos:
			amd64: -tags="tempdll latest"
			arm64: -tags="tempdll latest"
	
	.  Execute command
`,
}

func init() {
	CmdBuild.Run = runBuild
}

func runBuild(c *command.Config) error {
	return build.Build(c)
}
