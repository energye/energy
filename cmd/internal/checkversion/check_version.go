//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package checkversion

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/term"
	"io/ioutil"
	"net/http"
)

//Check 检查版本
func Check() {
	term.Section.Println(" current", fmt.Sprintf("%d.%d.%d", term.Build, term.Major, term.Minor))
	remoteVersion()
}

// 获取远程版本
func remoteVersion() {
	term.Section.Println(" latest ", fmt.Sprintf("%d.%d.%d", term.Build, term.Major, term.Minor))
}

type GeoInfo struct {
	Country string `json:"country"`
}

func area() {
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
