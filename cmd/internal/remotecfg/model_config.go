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
	"errors"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"path/filepath"
	"strconv"
	"strings"
)

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
		term.Section.Println("Getting the latest version information from remote service...")
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

// 获取当前安装的版本
func (m TConfig) GetInstallVersion(c *command.Config) (installVersion *TVersionsUpgrade, cefModuleName string, liblclModuleName string, retErr error) {
	releaseList, err := VersionUpgradeList()
	if err != nil {
		retErr = err
		return
	}

	if c.Install.Version == "latest" {
		// 最新版本, latest=x.x.x
		installVersion = releaseList.Item(m.LatestVersion.Version)
	} else {
		// 自己选择版本
		ver := c.Install.Version
		// 去除 v
		if ver[0] == 'v' {
			ver = ver[1:]
		}
		installVersion = releaseList.Item(ver)
	}
	if installVersion == nil {
		retErr = errors.New("Invalid Version Number: " + c.Install.Version)
		return
	}
	// 指定支持的固定 CEF 版本号
	cef := strings.ToLower(c.Install.CEF)
	if cef == "" {
		if consts.IsWindows {
			// 判断 windows 当小于 windows 10，默认使用 CEF 109
			majorVersion, _, _ := tools.VersionNumber()
			if majorVersion < 10 {
				cef = consts.CEF109 // windows7 默认
			}
		}
	}
	// 匹配固定的几个模块名
	if cef == consts.CEF109 {
		cefModuleName = "cef-109"
		liblclModuleName = "liblcl-109"
	} else if cef == consts.CEF106 {
		cefModuleName = "cef-106"
		liblclModuleName = "liblcl-106"
	} else if cef == consts.CEF101 {
		cefModuleName = "cef-101"
		liblclModuleName = "liblcl-101"
	} else if cef == consts.CEF87 {
		cefModuleName = "cef-87"
		liblclModuleName = "liblcl-87"
	} else if cef == consts.CEF89 {
		cefModuleName = "cef-89"
		liblclModuleName = "liblcl-89"
	}
	// 未匹配到, 找到当前安装 energy 版本所支持的最新 CEF 版本号, 规则为取版本号最大数字
	if cefModuleName == "" {
		var (
			cefDefault string
			number     int
		)
		for module, _ := range installVersion.DependenceModule.CEF {
			// module = "cef-xxx"
			if s := strings.Split(module, "-"); len(s) == 2 {
				n, _ := strconv.Atoi(s[1])
				if n >= number {
					number = n
					cefDefault = module
				}
			} else {
				// module = "cef"
				cefDefault = module
				break
			}
		}
		cefModuleName = cefDefault
		liblclModuleName = "liblcl" // 固定名前缀
	}
	return
}

// 获取安装目录, 使用 cef 模块 + 系统 + 架构
func (m TConfig) GetFrameworkInstallPath(c *command.Config) string {
	frameworkName := m.GetFrameworkName(c)
	if frameworkName != "" {
		path := c.Install.Path // cli 当前执行目录 或 用户指定目录
		path = filepath.Join(path, consts.ENERGY)
		path = filepath.Join(path, frameworkName)
		return path
	}
	return ""
}

// 获取 CEF 完整名 CEF_[VER]_[OS]_[ARCH]
func (m TConfig) GetFrameworkName(c *command.Config) string {
	_, cefModuleName, _, err := m.GetInstallVersion(c)
	if cefModuleName != "" && err == nil {
		frameworkName := fmt.Sprintf("%s_%s_%s", cefModuleName, c.Install.OS.Value(), c.Install.Arch.Value())
		return strings.ToUpper(frameworkName)
	}
	return ""
}
