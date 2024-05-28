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
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
)

const (
	cliVersionURL = "https://energye.github.io/data/command-line-tools.json"
)

type cliVersion struct {
	Build       int    `json:"build"`
	Major       int    `json:"major"`
	Minor       int    `json:"minor"`
	DownloadURL string `json:"downloadUrl"`
}

// Check 检查版本
func Check() {
	term.Section.Println(" Current:", fmt.Sprintf("%d.%d.%d", term.Build, term.Major, term.Minor))
	cli, err := remoteVersion()
	if err != nil {
		term.Logger.Error("Check cli version failed to obtain remote information: " + err.Error())
		return
	}
	term.Section.Println(" Latest :", fmt.Sprintf("%d.%d.%d", cli.Build, cli.Major, cli.Minor))
	cv, err := strconv.Atoi(fmt.Sprintf("%d%d%d", term.Build, term.Major, term.Minor))
	if err != nil {
		term.Logger.Error("Check cli version failed: " + err.Error())
		return
	}
	rv, err := strconv.Atoi(fmt.Sprintf("%d%d%d", cli.Build, cli.Major, cli.Minor))
	if err != nil {
		term.Logger.Error("Check cli version failed: " + err.Error())
		return
	}
	if cv < rv {
		// 先这样，以后在规范名字
		cliName := "energy-" + runtime.GOOS
		if consts.IsARM64 {
			cliName += "arm"
		}
		if consts.IsWindows && consts.Is386 {
			cliName += "-32"
		} else {
			cliName += "-64"
		}
		cliName += ".zip"
		//term.Section.Println("There is a new version available, would you like to update?(y)")
		downloadURL, _ := url.JoinPath(cli.DownloadURL, cliName)
		term.Section.Println("There new version available. Download:", downloadURL)
	}
}

// 获取远程版本
func remoteVersion() (*cliVersion, error) {
	data, err := tools.Get(cliVersionURL)
	if err != nil {
		return nil, err
	}
	var cli cliVersion
	err = json.Unmarshal(data, &cli)
	if err != nil {
		return nil, err
	}
	return &cli, nil
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
