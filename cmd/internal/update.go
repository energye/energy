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
	"github.com/cyber-xxm/energy/v2/cmd/internal/update"
)

var CmdUpdate = &command.Command{
	UsageLine: "update",
	Short:     "Check and update the version",
	Long: `
	Check and update the version
	Check the LibLCL used by the current environment
	Check the energy used by the current project
`,
}

func init() {
	CmdUpdate.Run = runUpdate
}

func runUpdate(c *command.Config) error {
	return update.Update(c)
}
