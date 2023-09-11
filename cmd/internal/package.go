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
	"github.com/energye/energy/v2/cmd/internal/packager"
	"github.com/energye/energy/v2/cmd/internal/project"
)

var CmdPackage = &command.Command{
	UsageLine: "package -p [path] -c [clean]",
	Short:     "Making an Installation Package",
	Long: `
	-p Project path, default current path. Can be configured in energy.json
	-c Clear configuration and regenerate the default configuration
	.  Execute default command

Making an Installation Package
	Windows: 
		Creating an installation program using NSIS for Windows
	Linux: 
		Creating deb installation packages using dpkg
	MacOS:
		Generate app package for energy
`,
}

func init() {
	CmdPackage.Run = runPackage
}

func runPackage(c *command.Config) error {
	if proj, err := project.NewProject(c.Package.Path); err != nil {
		return err
	} else {
		proj.Clean = c.Package.Clean
		if err = packager.GeneraInstaller(proj); err != nil {
			return err
		}
	}
	return nil
}
