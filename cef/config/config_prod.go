//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build prod
// +build prod

package config

import (
	"github.com/cyber-xxm/energy/v2/consts"
)

func init() {
	config = &Config{Framework: consts.ExeDir}
}

func (m *Config) FrameworkPath() string {
	return m.Framework
}
