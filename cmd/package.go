//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cmd

import "fmt"

var CmdPackage = &Command{
	UsageLine: "package -p [path] -m [model] -o [out]",
	Short:     "Making an Installation Package",
	Long: `

	-p Package directory
	-m Use mode to set online or offline, offline by default
	-o Output directory
	.  Execute default command

Making an Installation Package.

	Set the package directory by specifying Path. The current directory is the default
		Use mode to set online or offline, offline by default.
		The framework is automatically downloaded when installed in online mode, installation package will be much smaller.
		Offline mode makes the framework into the package,installation package will be large
`,
}

func init() {
	CmdPackage.Run = runPackage
}

func runPackage(c *CommandConfig) error {
	fmt.Println("runPackage", "mode:", c.Package.Mode, "path:", c.Package.Path, "out:", c.Package.Out)
	return nil
}
