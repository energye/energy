//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import (
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/install"
)

var CmdInstall = &command.Command{
	UsageLine: "install -p [path] -v [version] -n [name] -d [download] -c [cef]",
	Short:     "Automatically configure the CEF and Energy framework",
	Long: `
	-p Installation directory Default current directory
	-v Specifying a version number,Default latest
	-n Name of the frame after installation
	-d Download Source, 0:gitee or 1:github, Default empty
	-c Install system supports CEF version, provide 4 options, default empty
		default : Automatically select support for the latest version based on the current system.
		109 : CEF 109.1.18 is the last one to support Windows 7.
		106 : CEF 106.1.1 is the last default support for GTK2 in Linux.
		87  : CEF 87.1.14 is the last one to support Flash.
	.  Execute default command

Automatically configure the CEF and Energy framework.

During this process, CEF and Energy are downloaded.

Default framework name is "EnergyFramework".
`,
}

type downloadInfo struct {
	fileName      string
	frameworkPath string
	downloadPath  string
	url           string
	success       bool
	isSupport     bool
	module        string
}

func init() {
	CmdInstall.Run = runInstall
}

// https://cef-builds.spotifycdn.com/cef_binary_107.1.11%2Bg26c0b5e%2Bchromium-107.0.5304.110_windows64.tar.bz2
// 运行安装
func runInstall(c *command.Config) error {
	install.Install(c)
	return nil
}
