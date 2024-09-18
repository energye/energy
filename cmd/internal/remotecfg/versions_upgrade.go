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
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/tools"
)

type TVersionsUpgrade struct {
	Enable           int               `json:"enable"`
	Identical        string            `json:"identical"`
	DependenceModule TDependenceModule `json:"dependenceModule"`
}

type TDependenceModule struct {
	LCL map[string]string `json:"lcl"`
	CEF map[string]string `json:"cef"`
}

func VersionUpgradeList() (map[string]TVersionsUpgrade, error) {
	data, err := tools.Get(consts.VERSIONS_UPGRADE_URL)
	if err != nil {
		return nil, err
	}
	var vu map[string]TVersionsUpgrade
	err = json.Unmarshal(data, &vu)
	if err != nil {
		return nil, err
	}
	return vu, nil
}
