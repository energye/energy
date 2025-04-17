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
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
)

type TCMDVersion struct {
	Version     string `json:"-"`
	Major       int32  `json:"major"`
	Minor       int32  `json:"minor"`
	Build       int32  `json:"build"`
	DownloadURL string `json:"vDownloadUrl"`
}

func CMDVersion() (*TCMDVersion, error) {
	data, err := tools.Get(env.GlobalDevEnvConfig.RemoteURL(consts.CMD_VERSION_URL), env.GlobalDevEnvConfig.Proxy)
	if err != nil {
		return nil, err
	}
	var cv TCMDVersion
	err = json.Unmarshal(data, &cv)
	if err != nil {
		return nil, err
	}
	cv.Version = fmt.Sprintf("%v.%v.%v", cv.Major, cv.Minor, cv.Build)
	return &cv, nil
}
