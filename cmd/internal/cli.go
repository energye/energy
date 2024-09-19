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
	"github.com/energye/energy/v2/cmd/internal/cli"
	"github.com/energye/energy/v2/cmd/internal/command"
)

var CmdCli = &command.Command{
	UsageLine: "energy cli -v [version] -u [update]",
	Short:     "energy cli cmd",
	Long: `
energy cli version and update:
  -v Current CLI version check
  -u Attempt to update the current CLI
`,
}

func init() {
	CmdCli.Run = runCli
}

func runCli(c *command.Config) error {
	if c.Cli.Version {
		cli.Version()
	}
	if c.Cli.Update {
		return cli.Update(c.Cli)
	}
	return nil
}
