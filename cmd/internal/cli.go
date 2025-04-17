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
	"github.com/cyber-xxm/energy/v2/cmd/internal/cli"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
)

var CmdCli = &command.Command{
	UsageLine: "energy cli -v [version] -u [update]",
	Short:     "",
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
	var downloadURL string
	if c.Cli.Version || c.Cli.Update {
		downloadURL = cli.CheckVersion()
	} else {
		cli.PrintCLIVersion()
	}
	if c.Cli.Update && downloadURL != "" {
		err := cli.OnlineUpdate(downloadURL)
		if err != nil {
			return err
		}
		term.Section.Println("ENERGY CLI UPDATE SUCCESS")
	}
	return nil
}
