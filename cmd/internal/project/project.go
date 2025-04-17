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
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

type AppType int

const (
	AtApp AppType = iota
	AtHelper
)

// Project holds the data related to a ENERGY project
type Project struct {
	AppType        AppType `json:"-"`              // app, helper
	Clean          bool    `json:"-"`              // 清空配置重新生成
	Name           string  `json:"name"`           // 应用名称
	ProjectPath    string  `json:"projectPath"`    // 项目目录
	FrameworkPath  string  `json:"frameworkPath"`  // CEF 框架目录
	AssetsDir      string  `json:"assetsDir"`      // 构建配置所在目录 未指定使用田默认内置配置
	OutputFilename string  `json:"outputFilename"` // 输出安装包文件名
	HelperFilePath string  `json:"helperFilePath"` // MacOS helper 进程二进制文件目录
	LibEMFS        string  `json:"libemfs"`        // 内置libs存放目录, 以项目目录根目录开始 ProjectPath + Libs = liblcl.dll 目录, 默认libs
	Info           Info    `json:"info"`           // 应用信息
	NSIS           NSIS    `json:"nsis"`           // windows nsis 安装包
	Dpkg           DPKG    `json:"dpkg"`           // linux dpkg 安装包
	PList          PList   `json:"plist"`          // darwin plist 安装包
	Author         Author  `json:"author"`         // 作者信息
}

func (m *Project) setDefaults() {
	if m.Name == "" {
		m.Name = "energyapp"
	}
	if m.ProjectPath == "" || m.ProjectPath[0] == '@' {
		// 设置当前执行目录为项目目录
		m.ProjectPath = tools.CurrentExecuteDir()
	}
	if m.FrameworkPath == "" {
		m.FrameworkPath = filepath.ToSlash(env.GlobalDevEnvConfig.FrameworkPath())
	}
	if !tools.IsExist(m.FrameworkPath) {
		panic("energy framework directory does not exist: " + m.FrameworkPath)
	}
	if m.AssetsDir == "" {
		m.AssetsDir = "assets"
	}
	if m.OutputFilename == "" {
		m.OutputFilename = m.Name
	}
	if m.Author.Name == "" {
		m.Author.Name = "Your name"
	}
	if m.Author.Email == "" {
		m.Author.Email = "YourEmail@xxx.com"
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
	if m.Info.FileDescription == nil {
		v := "Built using ENERGY (https://github.com/energye/energy)"
		m.Info.FileDescription = &v
	}
	if m.LibEMFS == "" {
		m.LibEMFS = "Libs"
	}
	switch runtime.GOOS {
	case "windows":
		if !strings.HasSuffix(m.OutputFilename, ".exe") {
			m.OutputFilename += ".exe"
		}
	case "darwin", "linux":
		m.OutputFilename = strings.TrimSuffix(m.OutputFilename, ".exe")
	}
	var replaceRootPath = func(fieldValue string) string {
		if fieldValue == "" || fieldValue[0] != '@' {
			return fieldValue
		}
		if fieldValue[0] == '@' {
			fieldValue = fieldValue[1:]
			return filepath.Join(m.ProjectPath, fieldValue)
		}
		return fieldValue
	}
	switch runtime.GOOS {
	case "windows":
		m.Info.Manifest = replaceRootPath(m.Info.Manifest)
		m.Info.Icon = replaceRootPath(m.Info.Icon)
		m.NSIS.Icon = replaceRootPath(m.NSIS.Icon)
		m.NSIS.UnIcon = replaceRootPath(m.NSIS.UnIcon)
	case "darwin":
		m.PList.Icon = replaceRootPath(m.PList.Icon)
	case "linux":
		m.Info.Icon = replaceRootPath(m.Info.Icon)
	}
}

type Info struct {
	Manifest        string  `json:"manifest"`        //应用 manifest
	Title           string  `json:"title"`           //标题 desktop-name
	WMClass         string  `json:"wmClass"`         //VMClass desktop-VMClass
	Icon            string  `json:"icon"`            //应用图标
	CompanyName     string  `json:"companyName"`     //公司名称
	ProductName     string  `json:"productName"`     //产品名称
	FileVersion     string  `json:"FileVersion"`     //文件版本
	ProductVersion  string  `json:"productVersion"`  //产品版本
	Copyright       *string `json:"copyright"`       //版权
	Comments        *string `json:"comments"`        //exe详情描述
	FileDescription *string `json:"fileDescription"` //描述
}

func (m *Info) FromSlash() {
	m.Icon = filepath.FromSlash(m.Icon)
}

func (m *Info) ToSlash() {
	m.Icon = filepath.ToSlash(m.Icon)
}

func (m *NSIS) FromSlash() {
	m.Icon = filepath.FromSlash(m.Icon)
	m.UnIcon = filepath.FromSlash(m.UnIcon)
	m.License = filepath.FromSlash(m.License)
	for i, as := range m.Include {
		m.Include[i] = filepath.FromSlash(as)
	}
}

func (m *NSIS) ToSlash() {
	m.Icon = filepath.ToSlash(m.Icon)
	m.UnIcon = filepath.ToSlash(m.UnIcon)
	m.License = filepath.ToSlash(m.License)
	for i, as := range m.Include {
		m.Include[i] = filepath.ToSlash(as)
	}
}

// NSIS windows NSIS
type NSIS struct {
	Icon                  string   `json:"icon"`                  //安装包图标
	Include               []string `json:"include"`               //打包资源目录、或文件 ["/to/path/file.txt", "/to/dir/*.*", "/to/dir"]
	Exclude               []string `json:"exclude"`               //打包排除资源目录、或文件 ["/to/path/file.txt", "/to/dir/*.*", "/to/dir"]
	ShortCutName          string   `json:"shortCutName"`          //快捷方式名称
	UnIcon                string   `json:"unIcon"`                //安装包卸载图标
	License               string   `json:"license"`               //安装包授权信息,(license.txt)文件路径
	Language              string   `json:"language"`              //安装包语言, 中文: SimpChinese, 英文: English, 语言在 NSIS_HOME/Contrib/Language files
	RequestExecutionLevel string   `json:"requestExecutionLevel"` //admin or ""
	Compress              string   `json:"compress"`              //压缩CEF, 当前仅支持7z/a压缩，""(空)时不启用压缩 默认: 7za
	CompressName          string   `json:"compressName"`          //压缩CEF后的7z包名称
	UseCompress           bool     `json:"-"`                     //如果支持配置的, true=使用压缩
	CompressFile          string   `json:"-"`                     //压缩后的文件完全目录
}

type DPKG struct {
	Include      []string `json:"include"` //打包资源目录、或文件 ["/to/path/file.txt", "/to/dir/*.*", "/to/dir"]
	Exclude      []string `json:"exclude"` //打包排除资源目录、或文件 ["/to/path/file.txt", "/to/dir/*.*", "/to/dir"]
	Package      string   `json:"package"`
	Homepage     string   `json:"homepage"`
	Compress     string   `json:"compress"` //压纹CEF, 当前仅支持7z/a压缩，""(空)时不启用压缩 默认: 7za
	UseCompress  bool     `json:"-"`        //如果支持配置的, true=使用压缩
	CompressFile string   `json:"-"`        //压缩后的文件完全目录
}

type PList struct {
	Include                  []string `json:"include"`                  //打包资源目录、或文件 ["/to/path/file.txt", "/to/dir/*.*", "/to/dir"]
	Exclude                  []string `json:"exclude"`                  //打包排除资源目录、或文件 ["/to/path/file.txt", "/to/dir/*.*", "/to/dir"]
	Icon                     string   `json:"icon"`                     //应用图标, png 或 icns, 如果指定png则生成icns, 如果指定icns则直接使用
	BundleName               string   `json:"bundleName"`               //Application bundleName
	BundleDisplayName        string   `json:"bundleDisplayName"`        //Application bundleDisplayName
	BundleIdentifier         string   `json:"bundleIdentifier"`         //Application bsundleIdentifier
	FileVersion              string   `json:"fileVersion"`              //文件版本
	Locals                   []string `json:"locals"`                   //语言
	BundleVersion            string   `json:"bundleVersion"`            //内部版本
	BundleShortVersionString string   `json:"bundleShortVersionString"` //发布版本号版本
	Copyright                *string  `json:"copyright"`                //版权
	Comments                 *string  `json:"comments"`                 //exe详情描述
	LSUIElement              bool     `json:"-"`                        //UI
	Pkgbuild                 bool     `json:"-"`                        // 生成pkg安装包
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// APP项目配置转换到Project
func parse(projectData []byte) (*Project, error) {
	m := &Project{}
	err := json.Unmarshal(projectData, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// NewProject 创建项目对象, 根据energy.json配置
func NewProject(argsProjectPath string) (*Project, error) {
	if argsProjectPath == "" {
		// 设置当前执行目录为项目目录
		argsProjectPath = tools.CurrentExecuteDir()
	}
	config := filepath.Join(argsProjectPath, PlatformConfigFile(""))
	rawBytes, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}
	m, err := parse(rawBytes)
	if err != nil {
		return nil, err
	}
	if m.ProjectPath == "" || m.ProjectPath != argsProjectPath {
		m.ProjectPath = argsProjectPath
	}
	m.setDefaults()
	m.ProjectPath = filepath.FromSlash(m.ProjectPath)
	return m, nil
}

// PlatformConfigFile 返回指定平台的 energy 项目配置文件目录
func PlatformConfigFile(fileName string) string {
	if strings.TrimSpace(fileName) == "" {
		fileName = fmt.Sprintf(consts.ConfigFile, runtime.GOOS)
	}
	return filepath.Join("config", fileName)
}
