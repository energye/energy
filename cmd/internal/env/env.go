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
	"github.com/energye/energy/v2/cmd/internal/tools/homedir"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// DevEnvReadUpdate 开发环境信息读取和修改
func DevEnvReadUpdate() *command.EnergyConfig {
	home, err := homedir.Dir()
	if err != nil {
		term.Section.Println(err.Error())
		return nil
	}
	config := filepath.Join(home, ".energy")
	if !tools.IsExist(config) {
		// 创建
		cfg := command.EnergyConfig{
			Registry: consts.DomainGithub,
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
		if strings.TrimSpace(cfg.Registry) == "" {
			cfg.Registry = consts.DomainGithub
		}
		return &cfg
	}
}
