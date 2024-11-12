//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package consts

import (
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/term"
	"net/url"
)

// 获取远程配置的Registry
func RemoteURL(c command.EnergyConfig, targetURL string) string {
	r, err := url.JoinPath(c.Registry, targetURL)
	if err != nil {
		term.Logger.Error(err.Error())
		return ""
	}
	return r
}
