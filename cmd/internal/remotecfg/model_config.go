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

func BaseConfig() (*TConfig, error) {
	if configInstance == nil {
		configInstance = new(TConfig)
		lv, err := LatestVersion()
		if err != nil {
			return nil, err
		}
		mdc, err := ModeBaseConfig()
		if err != nil {
			return nil, err
		}
		cefCfg, err := ModelCEFConfig()
		if err != nil {
			return nil, err
		}
		lclCfg, err := ModelLCLConfig()
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
