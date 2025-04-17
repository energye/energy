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
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
)

type TModelCEFConfigs map[string]TModelCEFConfig

type TModelCEFConfig map[string]TModelItem

func ModelCEFConfig() (TModelCEFConfigs, error) {
	data, err := tools.Get(env.GlobalDevEnvConfig.RemoteURL(consts.MODEL_CEF_URL), env.GlobalDevEnvConfig.Proxy)
	if err != nil {
		return nil, err
	}
	var config TModelCEFConfigs
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (m TModelCEFConfigs) Model(name string) TModelCEFConfig {
	return m[name]
}

func (m TModelCEFConfig) Item(name string) TModelItem {
	item, ok := m[name]
	if ok && item.Identical != "" {
		return m.Item(item.Identical)
	}
	return item
}
