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
)

var CmdHelp = &command.Command{
	UsageLine: "energy [cmd] help",
	Short:     "energy [cmd] help",
	Long:      `energy [cmd] help`,
}

func init() {
	CmdHelp.Run = runHelp
}

func runHelp(c *command.Config) error {
	return nil
}
