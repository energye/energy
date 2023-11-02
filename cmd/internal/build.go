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
	UsageLine: "build -p [path] -u [upx] --UpxFlag --gtk -d [dll]",
	Short:     "build energy project",
	Long: `
	Building energy project
	-p Project path, default current path. Can be configured in energy.json
	-u Set this parameter and install upx. Use upx to compress the execution file.
	  --upxFlag: Upx command line parameters 
	--gtk Compile on Linux, enable TempDll. gtk2 or gtk3
	-d Enable built-in liblcl build, 
	  windows: 
	    if < windows10 use cef_109: -tags="tempdll && win7"
	    else >= windows10 use latest: -tags="tempdll"
	  linux: gtk2 or gtk3: 
	    gtk2 use cef_106: -tags="tempdll && gtk2" 
	    gtk3 use latest: -tags="tempdll && gtk3" 
	  macos:
	    use latest: -tags="tempdll"
	
	.  Execute command
`,
}

func init() {
	CmdBuild.Run = runBuild
}

func runBuild(c *command.Config) error {
	return build.Build(c)
}
