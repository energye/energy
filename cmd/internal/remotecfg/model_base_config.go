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
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"strconv"
)

type TExtracts map[string]TExtract
type TDownloadItems map[string]TDownloadItem

type TModeBaseConfig struct {
	DownloadSourceItem TDownloadSourceItem `json:"downloadSourceItem"` // 下载源集合配置
	Extract            TExtracts           `json:"extract"`            // 压缩包提取目录文件规则
}

type TDownloadSourceItem struct {
	LCL    TDownloadItems `json:"lcl"`
	CEF    TDownloadItems `json:"cef"`
	GoLang TDownloadItems `json:"golang"`
	NSIS   TDownloadItems `json:"nsis"`
	NSIS7Z TDownloadItems `json:"nsis7z"`
	Z7z    TDownloadItems `json:"7z"`
}

type TDownloadItem struct {
	Label   string `json:"label"`   // 显示的下载源名
	Url     string `json:"url"`     // 下载地址
	Md5     string `json:"md5"`     // 下载md5
	Version string `json:"version"` // 其它工具下载版本
}

type TExtract struct {
	CEF []string `json:"cef"`
	LCL []string `json:"lcl"`
}

func ModeBaseConfig() (*TModeBaseConfig, error) {
	data, err := tools.Get(env.GlobalDevEnvConfig.RemoteURL(consts.BASE_CONFIG_URL), env.GlobalDevEnvConfig.Proxy)
	if err != nil {
		return nil, err
	}
	var mbc TModeBaseConfig
	err = json.Unmarshal(data, &mbc)
	if err != nil {
		return nil, err
	}
	return &mbc, nil
}

func (m TExtracts) Item(os command.OS) TExtract {
	return m[os.Value()]
}

func (m TDownloadItems) Item(index int) TDownloadItem {
	return m[strconv.Itoa(index)]
}
