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
	UsageLine: "install -p [path] -v [version] -n [name] -d [download] -a [all] -os -arch -cef",
	Short:     "Automatic installation and configuration of the energy framework complete development environment",
	Long: `
	-p Installation directory Default current directory
	-v Specifying a version number,Default latest.\
	-n Name of the framework directory after installation, Default EnergyFramework.\
	-d Download Source, 0:gitee or 1:github, Default empty
	-a Install all, skip installation prompts (Y/n), default empty:false
	-os Specify install OS: [windows, linux, darwin], default current system: os
	-arch Specify install ARCH: [386, amd64, arm64], Default current system: architecture
	-cef Install system supports CEF version, provide 4 options, default empty
		default : Automatically select support for the latest version based on the current system.
		109 : CEF 109.1.18 is the last one to support Windows 7.
		106 : CEF 106.1.1 is the last default support for GTK2 in Linux.
		87  : CEF 87.1.14 is the last one to support Flash.
	.  Execute default command

Automatic installation and configuration of the energy framework complete development environment.
Installation package is downloaded over the network during the installation process.
According to the prompt (Y), install Golang、CEF(Chromium Embedded Framework)、NSIS, If installed, skip.
`,
}

func init() {
	CmdInstall.Run = runInstall
}

// https://cef-builds.spotifycdn.com/cef_binary_107.1.11%2Bg26c0b5e%2Bchromium-107.0.5304.110_windows64.tar.bz2
// 运行安装
func runInstall(c *command.Config) error {
	return install.Install(c)
}
