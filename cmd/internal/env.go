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
	"os"
)

var CmdEnv = &command.Command{
	UsageLine: "env",
	Short:     "Get energy framework development environment",
	Long: `
	Get the Framework pointed to by the ENERGY_HOME development environment variable
	.  Execute default command
`,
}

func init() {
	CmdEnv.Run = runGetEnv
}

func runGetEnv(c *command.Config) error {
	envPath := os.Getenv(command.EnergyHomeKey)
	println("ENERGY_HOME_KEY", envPath)
	return nil
}
