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
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const checkURL = "https://energy.yanghy.cn/autoconfig/update.json"

type LiblclCallback func(model *Model, level int)

var (
	// 检查更新开关, 默认关闭
	isCheckUpdate   = false
	CanUpdateLiblcl LiblclCallback
)

// Update 更新模块
type Update struct {
	Liblcl Model `json:"liblcl"`
	Energy Model `json:"energy"`
	CEF    Model `json:"cef"`
}

// Model 模块
type Model struct {
	Latest   string                 `json:"latest"`   // 最新版本
	Download Download               `json:"download"` // 下载源
	Enable   bool                   `json:"enable"`   // 是否开启该模块更新
	Versions map[string]VersionInfo `json:"versions"` // 当前模块所有版本集合 key=版本, value=版本信息
}

// 更新下载源
type Download struct {
	Url          string   `json:"url"`          // 下载地址模板 https://{url}/energye/energy/releases/download/{version}/{OSARCH}.zip
	Source       []string `json:"source"`       // 下载地址源 ["gitee.com", "github.com"]
	SourceSelect uint8    `json:"sourceSelect"` // 下载地址源选择, 根据下标(index)选择源(Source), 替换到(Url)模板
}

// 版本信息
type VersionInfo struct {
	Content       []string `json:"content"`       // 更新内容
	Forced        bool     `json:"forced"`        // 是否强制更新, 适用当前版本时才启作用
	EnergyVersion string   `json:"energyVersion"` // 所属 energy 版本
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
	request, err := http.NewRequest("GET", checkURL, nil)
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
		var v = Update{}
		if err = json.Unmarshal(data, &v); err == nil {
			liblcl(&v.Liblcl)
		} else {
			println("energy check update json.Unmarshal error", err.Error())
		}
	} else {
		println("energy check update ioutil.ReadAl(response.Body) error", err.Error())
	}
}

// liblcl Model
func liblcl(model *Model) {
	if model.Enable {
		r1, _, _ := imports.Proc(1).Call()
		currentLib := api.GoStr(r1)
		originLib := model.Latest
		if c, v := compare(currentLib, originLib); c {
			if CanUpdateLiblcl != nil {
				CanUpdateLiblcl(model, v)
			}
		}
	}
}

// Version comparison, returns true if the current version is smaller than the remote version
func compare(current, origin string) (bool, int) {
	if current == "" {
		current = "0.0.0"
	}
	if origin == "" {
		origin = "0.0.0"
	}
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
