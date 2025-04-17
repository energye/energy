//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !prod
// +build !prod

package config

import (
	"encoding/json"
	"github.com/cyber-xxm/energy/v2/consts"
	"io/ioutil"
	"path/filepath"
)

func init() {
	homeDIR := consts.HomeDir
	if homeDIR == "" {
		println("Warning, failed to obtain the current user directory.")
	} else {
		filePath := filepath.Join(homeDIR, ".energy")
		energyConfig, err := ioutil.ReadFile(filePath)
		if err != nil {
			println("Read .energy Error:", err.Error())
			return
		}
		tempConfig := &Config{}
		err = json.Unmarshal(energyConfig, tempConfig)
		if err != nil {
			println("Parsing .energy file Error:", err.Error())
			return
		}
		config = tempConfig
	}
}

func (m *Config) FrameworkPath() string {
	return filepath.Join(m.Root, "energy", m.Framework)
}
