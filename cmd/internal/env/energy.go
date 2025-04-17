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
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools/homedir"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
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
			term.Logger.Error(err.Error())
			return
		}
		config := filepath.Join(home, ".energy")
		if finfo, err := os.Stat(config); err == nil {
			if finfo.IsDir() {
				os.RemoveAll(config)
			}
		}
		if !tools.IsExist(config) {
			//var frameworkDir string
			//println("Use Current Directory install framework? (Y)")
			//println("Current Directory:", wd)
			//var input string
			//for {
			//	print("(Y) Or Other Directory: ")
			//	fmt.Scan(&input)
			//	println()
			//	input = strings.ToLower(strings.TrimSpace(input))
			//	if input == "y" || input == "yes" || input == "" {
			//		frameworkDir = wd
			//		break
			//	} else {
			//		if filepath.IsAbs(input) {
			//			frameworkDir = input
			//			break
			//		}
			//	}
			//}

			// 创建
			GlobalDevEnvConfig = &EnergyConfig{
				homedir:  config,
				Registry: consts.DomainYangHY,
				Root:     wd,
			}
			GlobalDevEnvConfig.Update()
		} else {
			// 读取&更新
			cfgJSON, err := ioutil.ReadFile(config)
			if err != nil {
				term.Logger.Error(err.Error())
				return
			}
			GlobalDevEnvConfig = &EnergyConfig{
				homedir: config,
			}
			err = json.Unmarshal(cfgJSON, GlobalDevEnvConfig)
			if err != nil {
				term.Logger.Error(err.Error())
				return
			}
			if strings.TrimSpace(GlobalDevEnvConfig.Registry) == "" {
				GlobalDevEnvConfig.Registry = consts.DomainYangHY
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
	GoRoot    string `json:"goroot"`    // Go根目录, 如果配置构建时使用
	Version   string `json:"version"`   // 全局默认版本号 vx.x.x
	NSIS      string `json:"nsis"`      // nsis 路径
	Z7Z       string `json:"z7z"`       // 7z 路径
	UPX       string `json:"upx"`       // upx 路径
	Root      string `json:"root"`      // framework 根目录， 不包括 "energy" 目录
	Framework string `json:"framework"` // CEF framework 名
	Registry  string `json:"registry"`  // 远程版本信息获取源
	Proxy     string `json:"proxy"`     // 代理
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
	return filepath.Join(m.Root, consts.ENERGY, m.Framework)
}

// 返回当前使用的 CEF Framework 版本号
// 规则 Framework: CEF-[VER]_[OS]_[ARCH]
// 返回 VER
func (m *EnergyConfig) CEFVersion() int {
	framework := strings.Split(m.Framework, "-")
	// CEF [VER]_[OS]_[ARCH]
	if len(framework) != 2 {
		return 0
	}
	// split: [VER]_[OS]_[ARCH]
	framework = strings.Split(framework[1], "_")
	if len(framework) != 3 {
		return 0
	}
	// [VER] [OS] [ARCH]
	ver, err := strconv.Atoi(framework[0])
	if err != nil {
		return 0
	}
	return ver
}

// 返回当前使用的 CEF Framework 系统 (windows, darwin, linux)
// 规则 Framework: CEF-[VER]_[OS]_[ARCH]
// 返回 OS
func (m *EnergyConfig) OS() string {
	framework := strings.Split(m.Framework, "-")
	// CEF [VER]_[OS]_[ARCH]
	if len(framework) != 2 {
		return ""
	}
	// split: [VER]_[OS]_[ARCH]
	framework = strings.Split(framework[1], "_")
	if len(framework) != 3 {
		return ""
	}
	// [VER] [OS] [ARCH]
	os := strings.ToLower(framework[1])
	return os
}

// 返回当前使用的 CEF Framework 架构 (386, amd64, arm, arm64, loong64)
// 规则 Framework: CEF-[VER]_[OS]_[ARCH]
// 返回 ARCH
func (m *EnergyConfig) Arch() string {
	framework := strings.Split(m.Framework, "-")
	// CEF [VER]_[OS]_[ARCH]
	if len(framework) != 2 {
		return ""
	}
	// split: [VER]_[OS]_[ARCH]
	framework = strings.Split(framework[1], "_")
	if len(framework) != 3 {
		return ""
	}
	// [VER] [OS] [ARCH]
	arch := strings.ToLower(framework[2])
	switch arch {
	case "386", "i386", "32":
		return "386"
	case "amd64", "64", "x64", "x86_64":
		return "amd64"
	case "arm", "arm64", "loong64":
		return arch
	}
	return ""
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
	//if tools.CommandExists("makensis") {
	//	return "makensis"
	//}

	nsis := filepath.Join(m.NSIS, tools.FixCMDName("makensis"))
	if tools.IsExist(nsis) {
		return nsis
	}
	return ""
}

func (m *EnergyConfig) Z7ZCMD() string {
	//if tools.CommandExists("7z") {
	//	return "7z"
	//}
	//if tools.CommandExists("7za") {
	//	return "7za"
	//}
	z7z := filepath.Join(m.Z7Z, tools.FixCMDName("7z"))
	if tools.IsExist(z7z) {
		return z7z
	}
	z7z = filepath.Join(m.Z7Z, tools.FixCMDName("7za"))
	if tools.IsExist(z7z) {
		return z7z
	}
	z7z = filepath.Join(m.Z7Z, tools.FixCMDName("7zz"))
	if tools.IsExist(z7z) {
		return z7z
	}
	return ""
}

func (m *EnergyConfig) UPXCMD() string {
	//if tools.CommandExists("upx") {
	//	return "upx"
	//}
	upx := filepath.Join(m.UPX, tools.FixCMDName("upx"))
	if tools.IsExist(upx) {
		return upx
	}
	return ""
}
