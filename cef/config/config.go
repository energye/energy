//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package config

var config *Config

type Config struct {
	Root      string `json:"root"`
	Framework string `json:"framework"`
}

// 返回 config 环境
// mode: dev 使用 .energy
// prod 使用自定义或当前执行目录
func Get() *Config {
	return config
}
