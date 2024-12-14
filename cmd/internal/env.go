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
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/term"
)

var CmdEnv = &command.Command{
	UsageLine: "env",
	Short:     "Get ENERGY Framework Development Environment",
	Long:      ``,
}

func init() {
	CmdEnv.Run = runGetEnv
}

func runGetEnv(c *command.Config) error {
	//term.Section.Println(consts.GolanHomeKey, os.Getenv(consts.GolanHomeKey))
	//term.Section.Println(consts.EnergyHomeKey, os.Getenv(consts.EnergyHomeKey))
	//if consts.IsWindows {
	//	term.Section.Println(consts.NSISHomeKey, os.Getenv(consts.NSISHomeKey))
	//	term.Section.Println(consts.Z7ZHomeKey, os.Getenv(consts.Z7ZHomeKey))
	//}
	//if !consts.IsDarwin {
	//	term.Section.Println(consts.UPXHomeKey, os.Getenv(consts.UPXHomeKey))
	//}

	term.Section.Println("Golang", env.GlobalDevEnvConfig.GoRoot)
	term.Section.Println("ENERGY_HOME", env.GlobalDevEnvConfig.Framework)
	term.Section.Println("NSIS", env.GlobalDevEnvConfig.NSIS)
	term.Section.Println("7z", env.GlobalDevEnvConfig.Z7Z)
	term.Section.Println("UPX", env.GlobalDevEnvConfig.UPX)
	return nil
}
