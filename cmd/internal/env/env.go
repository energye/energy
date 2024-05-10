//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package env

import (
	"encoding/json"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/energy/homedir"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DevEnvReadUpdate 开发环境信息读取和修改
func DevEnvReadUpdate(goroot, cefroot, nsisroot, upxroot, z7zroot string) *command.EnergyConfig {
	home, err := homedir.Dir()
	if err != nil {
		term.Section.Println(err.Error())
		return nil
	}
	energyDir := filepath.Join(home, ".energy")
	if !tools.IsExist(energyDir) {
		err = os.MkdirAll(energyDir, os.ModePerm)
		if err != nil {
			term.Section.Println(err.Error())
			return nil
		}
	}
	config := filepath.Join(energyDir, "energy.json")
	if goroot == "" {
		goroot = os.Getenv("GOROOT")
	}
	if cefroot == "" {
		cefroot = os.Getenv(consts.EnergyHomeKey)
	}
	if nsisroot == "" {
		nsisroot = os.Getenv(consts.NSISHomeKey)
	}
	if upxroot == "" {
		upxroot = os.Getenv(consts.UPXHomeKey)
	}
	if z7zroot == "" {
		z7zroot = os.Getenv(consts.Z7ZHomeKey)
	}
	if !tools.IsExist(config) {
		// 创建
		cfg := command.EnergyConfig{
			GoRoot:   goroot,
			CEFRoot:  cefroot,
			NSISRoot: nsisroot,
			UPXRoot:  upxroot,
			Z7zRoot:  z7zroot,
			Source: command.DownloadSource{
				Golang: consts.GolangDownloadSource,
				CEF:    consts.CEFDownloadSource,
			},
		}
		cfgJSON, err := json.MarshalIndent(&cfg, "", "\t")
		if err != nil {
			term.Section.Println(err.Error())
			return nil
		}
		if err := ioutil.WriteFile(config, cfgJSON, 0644); err != nil {
			term.Section.Println(err.Error())
			return nil
		}
		return &cfg
	} else {
		// 读取&更新
		cfgJSON, err := ioutil.ReadFile(config)
		if err != nil {
			term.Section.Println(err.Error())
			return nil
		}
		cfg := command.EnergyConfig{}
		err = json.Unmarshal(cfgJSON, &cfg)
		if err != nil {
			term.Section.Println(err.Error())
			return nil
		}
		isUpdate := false
		if goroot != "" {
			cfg.GoRoot = goroot
			isUpdate = true
		}
		if cefroot != "" {
			cfg.CEFRoot = cefroot
			isUpdate = true
		}
		if nsisroot != "" {
			cfg.NSISRoot = nsisroot
			isUpdate = true
		}
		if upxroot != "" {
			cfg.UPXRoot = upxroot
			isUpdate = true
		}
		if z7zroot != "" {
			cfg.Z7zRoot = z7zroot
			isUpdate = true
		}
		if isUpdate {
			cfgJSON, err = json.MarshalIndent(&cfg, "", "\t")
			if err != nil {
				term.Section.Println(err.Error())
				return nil
			}
			configFile, err := os.OpenFile(config, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
			if err != nil {
				term.Section.Println(err.Error())
				return nil
			}
			defer configFile.Close()
			_, err = configFile.Write(cfgJSON)
			if err != nil {
				term.Section.Println(err.Error())
				return nil
			}
		}
		return &cfg
	}
}
