//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package common

import (
	"encoding/json"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

// 返回 CEF 支持的系统, 格式 [os][arch] 示例 windows32, windows64, macosx64 ...
func LibLCLOS(c *command.Config) string {
	ins := c.Install
	os := command.OS(runtime.GOOS)
	arch := command.Arch(runtime.GOARCH)
	if ins.OS != "" {
		os = ins.OS
	}
	if ins.Arch != "" {
		arch = ins.Arch
	}
	return OS(os, arch)
}

func OS(os command.OS, arch command.Arch) string {
	if os.IsMacOS() {
		os = "macos"
		if arch.IsAMD64() && os[len(os)-1] != 'x' {
			os += "x"
		}
	}
	if arch.Is386() {
		arch = "32"
	} else if arch.IsAMD64() {
		arch = "64"
	}
	var LibLCLFileNames = map[string]string{
		"windows32":    consts.Windows32,
		"windows64":    consts.Windows64,
		"windowsarm64": consts.WindowsARM64,
		"linuxarm":     consts.LinuxARM,
		"linuxarm64":   consts.LinuxARM64,
		"linux32":      consts.Linux32,
		"linux64":      consts.Linux64,
		"macosx64":     consts.MacOSX64,
		"macosarm64":   consts.MacOSARM64,
		"linuxloong64": consts.LinuxLoong64,
	}
	return LibLCLFileNames[fmt.Sprintf("%v%v", os, arch)]
}

// 返回 liblcl 在 linux 版本号大于 106 版本默认使用 GTK3
// 完整格式: liblcl[-ver][-GTK3]
// 入参:
// insOS: 传入的OS
// insWS: 传入的控件UI, 只Linux有效 GTK2
// liblclModuleName: 模块名, liblcl-109
// moduleVersion: 模块版本, 109
func LibLCLLinuxUseGTK3(insOS command.OS, insWS, liblclModuleName, moduleVersion string) string {
	// default os
	os := command.OS(runtime.GOOS)
	if insOS != "" {
		os = insOS
	}
	// Linux 从106版本以后默认使用 GTK3
	if os.IsLinux() && !tools.Equals(insWS, "GTK2") {
		isGT106 := false         // 标记大于106
		if moduleVersion == "" { // 空值表示 LibLCL 是支持 CEF 最新版本, 默认用GTK3
			isGT106 = true
		} else {
			if ver, err := strconv.Atoi(moduleVersion); err == nil && ver > 106 {
				isGT106 = true
			}
		}
		if isGT106 {
			// GTK3 的 liblcl 库压缩文件名 + -GTK3
			return liblclModuleName + "-GTK3"
		}
	}
	return liblclModuleName
}

// url文件名
func UrlName(downloadUrl string) string {
	downloadUrl = downloadUrl[strings.LastIndex(downloadUrl, "/")+1:]
	return downloadUrl
}

func Area() {
	resp, err := http.Get("http://ip-api.com/json/") // 使用公共API，注意实际使用时选择合适的服务
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	var geoInfo GeoInfo
	err = json.Unmarshal(body, &geoInfo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Country:", geoInfo.Country)
}

type GeoInfo struct {
	Country string `json:"country"`
}
