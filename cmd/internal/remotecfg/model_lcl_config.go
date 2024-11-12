//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package remotecfg

import (
	"encoding/json"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/tools"
)

type TModelLCLConfigs map[string]TModelLCLConfig

type TModelLCLConfig map[string]TModelItem

func ModelLCLConfig(c command.EnergyConfig) (TModelLCLConfigs, error) {
	data, err := tools.Get(consts.RemoteURL(c, consts.MODEL_LCL_URL))
	if err != nil {
		return nil, err
	}
	var config TModelLCLConfigs
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (m TModelLCLConfigs) Model(name string) TModelLCLConfig {
	return m[name]
}

func (m TModelLCLConfig) Item(name string) TModelItem {
	item, ok := m[name]
	if ok && item.Identical != "" {
		return m.Item(item.Identical)
	}
	return item
}
