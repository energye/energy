//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package project

import (
	"encoding/json"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/golcl/energy/tools"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Project holds the data related to a ENERGY project
type Project struct {
	Name           string `json:"name"`           // 应用名称
	ProjectPath    string `json:"projectPath"`    // 项目目录
	FrameworkPath  string `json:"frameworkPath"`  // 框架目录 未指定时使用环境变量 ENERGY_HOME
	BuildAssetsDir string `json:"buildAssetsDir"` // 构建配置所在目录 未指定使用田默认内置配置
	OutputFilename string `json:"outputFilename"` // 输出安装包文件名
	Info           Info   `json:"info"`           // 应用信息
	Author         Author `json:"author"`         // 作者信息
}

func (m *Project) setDefaults() {
	if m.Name == "" {
		m.Name = "energyapp"
	}
	if m.ProjectPath == "" {
		m.ProjectPath, _ = os.Getwd()
	}
	if m.FrameworkPath == "" {
		m.FrameworkPath = os.Getenv(consts.EnergyHomeKey)
	}
	if !tools.IsExist(m.FrameworkPath) {
		panic("energy framework directory does not exist: " + m.FrameworkPath)
	}
	if m.BuildAssetsDir == "" {
		m.BuildAssetsDir = "assets"
	}
	if m.OutputFilename == "" {
		m.OutputFilename = m.Name
	}
	if m.Author.Name == "" {
		m.Author.Name = "yanghy"
	}
	if m.Author.Email == "" {
		m.Author.Email = "snxamdf@126.com"
	}
	if m.Info.CompanyName == "" {
		m.Info.CompanyName = m.Name
	}
	if m.Info.ProductName == "" {
		m.Info.ProductName = m.Name
	}
	if m.Info.ProductVersion == "" {
		m.Info.ProductVersion = "1.0.0"
	}
	if m.Info.Copyright == nil {
		v := "Copyright........."
		m.Info.Copyright = &v
	}
	if m.Info.Comments == nil {
		v := "Built using ENERGY (https://github.com/energye/energy)"
		m.Info.Comments = &v
	}
	switch runtime.GOOS {
	case "windows":
		if !strings.HasSuffix(m.OutputFilename, ".exe") {
			m.OutputFilename += ".exe"
		}
	case "darwin", "linux":
		m.OutputFilename = strings.TrimSuffix(m.OutputFilename, ".exe")
	}

}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Info struct {
	Icon           string  `json:"icon"`           //应用图标
	InstallIcon    string  `json:"installIcon"`    //应用包装图标
	CompanyName    string  `json:"companyName"`    //公司名称
	ProductName    string  `json:"productName"`    //产品名称
	FiletVersion   string  `json:"filetVersion"`   //文件版本
	ProductVersion string  `json:"productVersion"` //产品版本
	Copyright      *string `json:"copyright"`      //版权
	Comments       *string `json:"comments"`       //描述
	License        string  `json:"license"`        //安装包授权信息,(license.txt)文件路径
	Language       string  `json:"language"`       //安装包语言, 中文: SimpChinese, 英文: English, 语言在 NSIS_HOME/Contrib/Language files
}

//  APP项目配置转换到Project
func parse(projectData []byte) (*Project, error) {
	m := &Project{}
	err := json.Unmarshal(projectData, m)
	if err != nil {
		return nil, err
	}
	m.setDefaults()
	return m, nil
}

// NewProject 创建项目对象, 根据energy.json配置
func NewProject(projectPath string) (*Project, error) {
	if projectPath == "" {
		projectPath, _ = os.Getwd()
	}
	config := filepath.Join(projectPath, consts.EnergyProjectConfig)
	rawBytes, err := os.ReadFile(config)
	if err != nil {
		return nil, err
	}
	m, err := parse(rawBytes)
	if err != nil {
		return nil, err
	}
	m.ProjectPath = projectPath
	return m, nil
}
