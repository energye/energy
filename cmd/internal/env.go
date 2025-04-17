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
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
)

var CmdEnv = &command.Command{
	UsageLine: "env",
	Short:     "",
	Long:      ``,
}

func init() {
	CmdEnv.Run = runGetEnv
}

func runGetEnv(c *command.Config) error {

	return env.Env(c)
}
