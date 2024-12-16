//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package env

import (
	"encoding/json"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/energy/v2/cmd/internal/tools/homedir"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var GlobalDevEnvConfig *EnergyConfig

// 获取远程配置的Registry
func (m *EnergyConfig) RemoteURL(targetURL string) string {
	r, err := url.JoinPath(m.Registry, targetURL)
	if err != nil {
		term.Logger.Error(err.Error())
		return ""
	}
	return r
}

// InitDevEnvConfig 开发环境信息读取和修改
func InitDevEnvConfig(wd string) {
	if GlobalDevEnvConfig == nil {
		home, err := homedir.Dir()
		if err != nil {
			term.Section.Println(err.Error())
			return
		}
		config := filepath.Join(home, ".energy")
		if finfo, err := os.Stat(config); err == nil {
			if finfo.IsDir() {
				os.RemoveAll(config)
			}
		}
		if !tools.IsExist(config) {
			// 创建
			GlobalDevEnvConfig = &EnergyConfig{
				homedir:  config,
				Registry: consts.DomainGithub,
				Root:     wd,
			}
			GlobalDevEnvConfig.Update()
		} else {
			// 读取&更新
			cfgJSON, err := ioutil.ReadFile(config)
			if err != nil {
				term.Section.Println(err.Error())
				return
			}
			GlobalDevEnvConfig = &EnergyConfig{
				homedir: config,
			}
			err = json.Unmarshal(cfgJSON, GlobalDevEnvConfig)
			if err != nil {
				term.Section.Println(err.Error())
				return
			}
			if strings.TrimSpace(GlobalDevEnvConfig.Registry) == "" {
				GlobalDevEnvConfig.Registry = consts.DomainGithub
				GlobalDevEnvConfig.Update()
			}
			if strings.TrimSpace(GlobalDevEnvConfig.Root) == "" {
				GlobalDevEnvConfig.Root = wd
				GlobalDevEnvConfig.Update()
			}
		}
	}
}

type EnergyConfig struct {
	homedir   string `json:"-"`
	GoRoot    string `json:"goroot"`
	NSIS      string `json:"nsis"`
	Z7Z       string `json:"z7z"`
	UPX       string `json:"upx"`
	Root      string `json:"root"`
	Framework string `json:"framework"`
	Registry  string `json:"registry"`
	Proxy     string `json:"proxy"`
}

func (m *EnergyConfig) Update() {
	if m.homedir == "" {
		home, err := homedir.Dir()
		if err != nil {
			term.Section.Println(err.Error())
			return
		}
		m.homedir = filepath.Join(home, ".energy")
	}
	cfgJSON, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		term.Section.Println(err.Error())
		return
	}
	if err := ioutil.WriteFile(m.homedir, cfgJSON, 0644); err != nil {
		term.Section.Println(err.Error())
		return
	}
}

func (m *EnergyConfig) FrameworkPath() string {
	return filepath.Join(m.Root, m.Framework)
}

func (m *EnergyConfig) GoCMD() string {
	if tools.CommandExists("go") {
		return "go"
	}
	gocmd := filepath.Join(m.GoRoot, "bin", "go")
	if tools.IsExist(gocmd) {
		return gocmd
	}
	return ""
}

func (m *EnergyConfig) NSISCMD() string {
	if tools.CommandExists("makensis") {
		return "makensis"
	}
	nsis := filepath.Join(m.NSIS, "makensis")
	if tools.IsExist(nsis) {
		return nsis
	}
	return ""
}

func (m *EnergyConfig) Z7ZCMD() string {
	if tools.CommandExists("7z") {
		return "7z"
	}
	if tools.CommandExists("7za") {
		return "7za"
	}
	z7z := filepath.Join(m.Z7Z, "7z")
	if tools.IsExist(z7z) {
		return z7z
	}
	z7z = filepath.Join(m.Z7Z, "7za")
	if tools.IsExist(z7z) {
		return z7z
	}
	z7z = filepath.Join(m.Z7Z, "7zz")
	if tools.IsExist(z7z) {
		return z7z
	}
	return ""
}

func (m *EnergyConfig) UPXCMD() string {
	if tools.CommandExists("upx") {
		return "upx"
	}
	upx := filepath.Join(m.UPX, "upx")
	if tools.IsExist(upx) {
		return upx
	}
	return ""
}
