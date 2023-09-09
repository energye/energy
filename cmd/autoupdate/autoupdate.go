//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package autoupdate Energy lib-lcl check auto update
package autoupdate

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/liblclbinres"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type LiblclCallback func(model *Model, level int, canUpdate bool)

var (
	// 检查更新开关, 默认关闭
	isCheckUpdate   = false
	CanUpdateLiblcl LiblclCallback // 参数 model: 更新模块, Level: 更新版本级别, canUpdate: 是否有更新
)

// Model 模块
type Model struct {
	CurrentVersion string                 `json:"-"`        // 当前版本
	Latest         string                 `json:"latest"`   // 最新版本
	Versions       map[string]VersionInfo `json:"versions"` // 当前模块所有版本集合 key=版本, value=版本信息
}

// 版本信息
type VersionInfo struct {
	Content              string `json:"content"`              // 更新内容
	DownloadSource       string `json:"downloadSource"`       // 下载源 逗号分隔, 数组 ["gitee.com", "github.com"]
	DownloadSourceSelect int    `json:"downloadSourceSelect"` // 下载源 选择
	DownloadUrl          string `json:"downloadUrl"`          // 下载地址, https://{url}/energye/energy/releases/download/{version}/{OSARCH}.zip
	Module               string `json:"module"`               // 模块名
	BuildSupportOSArch   string `json:"supportOSArch"`        // 已提供构建支持的系统架构
	Version              string `json:"version"`              // 版本
}

// CheckUpdate
//	Check for updates, when isCheckUpdate is true
func CheckUpdate() {
	if isCheckUpdate {
		check()
	}
}

// IsCheckUpdate
//	Set whether to check for updates
func IsCheckUpdate(v bool) {
	isCheckUpdate = v
}

func check() {
	client := &http.Client{Timeout: 5 * time.Second, Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	request, err := http.NewRequest("GET", consts.CheckUpgradeURL, nil)
	if err != nil {
		println("energy check update http.NewRequest error:", err.Error())
		return
	}
	response, err := client.Do(request)
	if err != nil {
		println("energy check update client.Do error:", err.Error())
		return
	}
	defer response.Body.Close()
	if data, err := ioutil.ReadAll(response.Body); err == nil {
		var v map[string]*Model
		if err = json.Unmarshal(data, &v); err == nil {
			liblcl(v["liblcl"])
		} else {
			println("energy check update json.Unmarshal error", err.Error())
		}
	} else {
		println("energy check update ioutil.ReadAl(response.Body) error", err.Error())
	}
}

// liblcl Model
func liblcl(model *Model) {
	currentLib := liblclbinres.LibVersion()
	originLib := model.Latest
	if currentLib == "" {
		currentLib = "0.0.0"
	}
	if originLib == "" {
		originLib = "0.0.0"
	}
	model.CurrentVersion = currentLib
	can, level := compare(currentLib, originLib)
	if CanUpdateLiblcl != nil {
		CanUpdateLiblcl(model, level, can)
	}
}

func LibLCLName(version, buildSupportOSArch string) (string, bool) {
	return internal.LibLCLName(version, buildSupportOSArch)
}

// 命名规则 OS+[ARCH]+BIT+[GTK2]
//  ARCH: 非必需, ARM 时填写, AMD为空
//  GTK2: 非必需, GTK2(Linux CEF 106) 时填写, 非Linux或GTK3时为空
func liblclOS(version, buildSupportOSArch string) (string, bool) {
	archs := strings.Split(buildSupportOSArch, ",")
	var goarch string
	if c.Install.OS.IsWindows() && c.Install.Arch.Is386() {
		goarch = "32" // windows32 = > windows386
	} else {
		goarch = string(c.Install.Arch)
	}
	noSuport := fmt.Sprintf("%v %v", c.Install.OS, goarch)
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}
	if name, isOld := liblclName(c, version, cef); isOld {
		if name == "" {
			return noSuport, false
		}
		return name, true
	} else {
		return name, isSupport(name)
	}
}

// Version comparison, returns true if the current version is smaller than the remote version
func compare(current, origin string) (bool, int) {
	cmajor, cminor, crevision := versionConvert(current)
	omajor, ominor, orevision := versionConvert(origin)
	if omajor > cmajor {
		return true, 1 // major
	} else if ominor > cminor {
		return true, 2 // minor
	} else if orevision > crevision {
		return true, 3 // revision
	}
	return false, 0
}

// Version number conversion=>Major version, minor version, revised version
func versionConvert(ver string) (major, minor, revision int) {
	ver = strings.ToLower(ver)
	lastv := strings.LastIndex(ver, "v")
	if lastv != -1 {
		ver = ver[lastv+1:]
	}
	vers := strings.Split(strings.Split(ver, "-")[0], ".")
	if len(vers) >= 3 {
		major, _ = strconv.Atoi(vers[0])
		minor, _ = strconv.Atoi(vers[1])
		revision, _ = strconv.Atoi(vers[2])
	} else if len(vers) >= 2 {
		major, _ = strconv.Atoi(vers[0])
		minor, _ = strconv.Atoi(vers[1])
	} else if len(vers) >= 1 {
		major, _ = strconv.Atoi(vers[0])
	}
	return
}
