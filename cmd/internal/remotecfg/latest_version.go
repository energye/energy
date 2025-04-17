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
	"strings"
)

type TLatestVersion struct {
	Version string `json:"version"`
	Major   int32  `json:"-"`
	Minor   int32  `json:"-"`
	Build   int32  `json:"-"`
}

func LatestVersion() (*TLatestVersion, error) {
	data, err := tools.Get(env.GlobalDevEnvConfig.RemoteURL(consts.LATEST_VERSION_URL), env.GlobalDevEnvConfig.Proxy)
	if err != nil {
		return nil, err
	}
	var lv TLatestVersion
	err = json.Unmarshal(data, &lv)
	if err != nil {
		return nil, err
	}
	vs := strings.Split(lv.Version, ".")
	lv.Major = int32(tools.ToInt(vs[0]))
	lv.Minor = int32(tools.ToInt(vs[1]))
	lv.Build = int32(tools.ToInt(vs[2]))
	return &lv, nil
}
