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
	"github.com/cyber-xxm/energy/v2/cmd/internal/remotecfg"
	"sort"
)

var CmdVersion = &command.Command{
	UsageLine: "version",
	Short:     "Get version list",
	Long:      ``,
}

func init() {
	CmdVersion.Run = runVersion
}

func runVersion(c *command.Config) error {
	vu, err := remotecfg.VersionUpgradeList()
	if err != nil {
		return err
	}
	latestVersion, err := remotecfg.LatestVersion()
	if err != nil {
		return err
	}
	var keys []string
	for k, _ := range vu {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	println("Latest:", latestVersion.Version)
	println("Version list")
	for i := len(keys) - 1; i >= 0; i-- {
		println("  ", keys[i])
	}
	return nil
}
