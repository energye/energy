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
	"strconv"
)

type TExtracts map[string]TExtract
type TDownloadItems map[string]TDownloadItem

type TModeBaseConfig struct {
	DownloadSourceItem TDownloadSourceItem `json:"downloadSourceItem"`
	Environment        TEnvironment        `json:"environment"`
	Extract            TExtracts           `json:"extract"`
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
	Label   string `json:"label"`
	Url     string `json:"url"`
	Md5     string `json:"md5"`
	Version string `json:"version"`
}

type TEnvironment struct {
	EnergyHomeKey string
	GolanHomeKey  string
	NSISHomeKey   string
	Z7ZHomeKey    string
}

type TExtract struct {
	CEF []string `json:"cef"`
	LCL []string `json:"lcl"`
}

func ModeBaseConfig() (*TModeBaseConfig, error) {
	data, err := tools.Get(consts.BASE_CONFIG_URL)
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

func (m TExtracts) Item(os string) TExtract {
	return m[os]
}

func (m TDownloadItems) Item(index int) TDownloadItem {
	return m[strconv.Itoa(index)]
}
