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

import "github.com/energye/energy/v2/cmd/internal/command"

var configInstance *TConfig

type TConfig struct {
	LatestVersion  *TLatestVersion  `json:"-"`
	ModeBaseConfig *TModeBaseConfig `json:"-"`
	ModelCEFConfig TModelCEFConfigs `json:"-"`
	ModelLCLConfig TModelLCLConfigs `json:"-"`
}

type TModelItem struct {
	DownloadSource     int    `json:"downloadSource"`
	DownloadSourceList []int  `json:"downloadSourceList"`
	SupportOSArch      string `json:"supportOSArch"`
	Identical          string `json:"identical"`
}

func BaseConfig(c command.EnergyConfig) (*TConfig, error) {
	if configInstance == nil {
		configInstance = new(TConfig)
		lv, err := LatestVersion(c)
		if err != nil {
			return nil, err
		}
		mdc, err := ModeBaseConfig(c)
		if err != nil {
			return nil, err
		}
		cefCfg, err := ModelCEFConfig(c)
		if err != nil {
			return nil, err
		}
		lclCfg, err := ModelLCLConfig(c)
		if err != nil {
			return nil, err
		}
		configInstance.LatestVersion = lv
		configInstance.ModeBaseConfig = mdc
		configInstance.ModelCEFConfig = cefCfg
		configInstance.ModelLCLConfig = lclCfg
	}
	return configInstance, nil
}
