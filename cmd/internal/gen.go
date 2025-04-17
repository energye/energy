//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 将资源以二进制形式生成go资源文件

package internal

import (
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/gen"
)

var CmdGen = &command.Command{
	UsageLine: "gen --icon --syso",
	Short:     "generate icons or syso commands",
	Long: `
	generate icons or syso commands
	--icon Used to generate application icons, can convert .png to .ico.
			Generate pixel size of: [256, 128, 64, 48, 32, 16]
		-p --iconFilePath: Icon source file directory
		-o --outPath: Save directory after icon generation. if empty, the current directory
	--syso Generate the application program xxx.syso, and when compiling the execution file
			the execution file information can be written into it
		-p --iconFilePath: Icon source file directory
		-o --outPath: Save directory after icon generation. if empty, the current directory
		-n --name: Generate the syso file name and move it to the application name
		-m --manifestFilePath: Manifest file directory. if empty, will use the default template
		-a --arch: amd64 or i386 or arm64. if empty, the current system architecture
		-i --infoFilePath: Generate directory for syso information data files in JSON format
`,
}

func init() {
	CmdGen.Run = runGen
}

func runGen(c *command.Config) error {
	return gen.Gen(c)
}
