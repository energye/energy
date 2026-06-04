//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package pack

import (
	"encoding/json"
)

var (
	// JSON app id
	//  appName + appId + version + arch
	//  go build -ldflags "-X github.com/energye/energy/v3/application/pack.JSON="{}""
	JSON = "{}"
	Info = info{}
)

type info struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
}

func init() {
	_ = json.Unmarshal([]byte(JSON), &Info)
}
