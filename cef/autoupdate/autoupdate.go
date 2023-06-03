//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package autoupdate Energy check auto update
package autoupdate

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cef/autoupdate/internal"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	// 检查更新开关, 默认关闭
	isCheckUpdate = false
)

// 更新模块
type update struct {
	Liblcl model `json:"liblcl"`
	Energy model `json:"energy"`
	CEF    model `json:"cef"`
}

// 模块
type model struct {
	Latest   string                 `json:"latest"`   // 最新版本
	Download download               `json:"download"` // 下载源
	Enable   bool                   `json:"enable"`   // 是否开启该模块更新
	Versions map[string]versionInfo `json:"versions"` // 当前模块所有版本集合 key=版本, value=版本信息
}

// 更新下载源
type download struct {
	Url          string   `json:"url"`          // 下载地址模板 https://{url}/energye/energy/releases/download/{version}/{OSARCH}.zip
	Source       []string `json:"source"`       // 下载地址源  ["gitee.com", "github.com"]
	SourceSelect uint8    `json:"sourceSelect"` // 下载地址源选择, 根据下标(index)选择源(Source), 替换到(Url)模板
}

// 版本信息
type versionInfo struct {
	Content []string `json:"content"` // 更新内容
	Forced  bool     `json:"forced"`  // 强制更新, 适用当前版本的启作用
}

// CheckUpdate
//	检查更新, isCheckUpdate 为true时
func CheckUpdate() {
	if isCheckUpdate {
		go check()
	}
}

// IsCheckUpdate
//	设置是否检查更新
func IsCheckUpdate(v bool) {
	isCheckUpdate = v
}

func UpdateLog() []string {
	return nil
}

func check() {
	client := &http.Client{Timeout: 5 * time.Second, Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	request, err := http.NewRequest("GET", internal.CheckURL, nil)
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
		//var v map[string]any
		var v = update{}
		if err = json.Unmarshal(data, &v); err == nil {
			fmt.Printf("data: %+v\n", v)
		} else {
			println("energy check update json.Unmarshal error", err.Error())
		}
	} else {
		println("energy check update ioutil.ReadAl(response.Body) error", err.Error())
	}
}

func updatePrompt() {

}

func updateDownload() {

}
