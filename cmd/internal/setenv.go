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
	"errors"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/tools"
)

var CmdSetenv = &command.Command{
	UsageLine: "setenv -p [path]",
	Short:     "Set energy framework development environment",
	Long: `
	-p Set the Framework pointed to by the ENERGY_HOME development environment variable
`,
}

func init() {
	CmdSetenv.Run = runSetenv
}

func runSetenv(c *command.Config) error {
	if c.Setenv.Path == "" {
		return errors.New("ERROR: ENERGY environment variable, command line argument [-p] directory to empty")
	}
	if !tools.IsExist(c.Setenv.Path) {
		return errors.New("Directory [" + c.Setenv.Path + "] does not exist")
	}
	env.GlobalDevEnvConfig.Framework = c.Setenv.Path
	env.GlobalDevEnvConfig.Update()
	println("SUCCESS")
	return nil
}
