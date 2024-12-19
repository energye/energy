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

func GetConfig() *Config {
	return config
}
