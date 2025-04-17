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
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/install"
)

var CmdInstall = &command.Command{
	UsageLine: "install -p [path] -v [version] -n [name] -d [download] --os --arch --cef",
	Short:     "Automatic installation and configuration of the energy framework complete development environment",
	Long: `
	-p Installation directory Default current directory
	-v Release number, Default latest
	-n Name of the framework directory after installation, Default EnergyFramework
	-d Download Source. Details: https://energye.github.io/data/model-base-config.json
	--all Skip select. Install All Software
	--os Specify install OS: [windows, linux, darwin], default current os
	--arch Specify install ARCH: [386, amd64, arm, arm64], default current arch
	--cef Install system supports CEF version
		 latest : Current system architecture Latest supported version.
		    109 : CEF 109.1.18 is the last one to support Windows 7
		    101 : CEF 101.0.8 is the last one supports Linux32-bit
		    87  : CEF 87.1.14 is the last one to support Flash
		    49  : CEF 49.0.2623 is the last on to support Windows XP
	--ws Set this parameter when GTK2 is used on Linux

Auto installation and configuration of the energy framework complete development environment.
Installation package is downloaded over the network during the installation process.
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
