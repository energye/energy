//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package build

import "github.com/energye/energy/v2/cmd/internal/command"

const (
	assetsFSPath = "assets/build/"
)

func Build(c *command.Config) error {
	return build(c)
}
