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
	"github.com/energye/energy/v2/cmd/internal/upgrade"
)

var CmdUpgrade = &command.Command{
	UsageLine: "upg",
	Short:     "check and upgrade to the latest version library",
	Long: `
	check and upgrade to the latest version library
`,
}

func init() {
	CmdUpgrade.Run = runUpgrade
}

func runUpgrade(c *command.Config) error {
	return upgrade.Upgrade(c)
}
