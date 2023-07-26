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
	"os"
)

var CmdEnv = &Command{
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

func runGetEnv(c *CommandConfig) error {
	envPath := os.Getenv(ENERGY_HOME_KEY)
	println("ENERGY_HOME_KEY", envPath)
	return nil
}
