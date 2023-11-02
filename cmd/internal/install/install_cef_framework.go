//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//

package install

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

// 安装CEF和liblcl框架
//  根据当前系统自动安装CEF
//  Windows
//    	系统版本号 > windows10: CEF > 109
//    	系统版本号 < windows10: CEF = 109
//  MacOS
//		CEF 最新版本
//  Linux
//    Gtk2: CEF = 106
//    Gtk3: CEF 最新版本
func installCEFFramework(c *command.Config) (string, func()) {
	if !c.Install.ICEF {
		return "", nil
	}
	pterm.Println()
	term.Section.Println("Install CEF")
	// 获取提取文件配置
	extractData, err := tools.HttpRequestGET(consts.DownloadExtractURL)
	if err != nil {
		term.Logger.Error(err.Error())
		return "", nil
	}
	var extractConfig map[string]any
	extractData = bytes.TrimPrefix(extractData, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(extractData, &extractConfig); err != nil {
		term.Logger.Error(err.Error())
		return "", nil
	}
	extractOSConfig := extractConfig[string(c.Install.OS)].(map[string]any)

	// 获取安装版本配置
	downloadJSON, err := tools.HttpRequestGET(consts.DownloadVersionURL)
	if err != nil {
		term.Logger.Error(err.Error())
		return "", nil
	}
	var edv map[string]any
	downloadJSON = bytes.TrimPrefix(downloadJSON, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(downloadJSON, &edv); err != nil {
		term.Logger.Error(err.Error())
		return "", nil
	}

	// -c cef args value
	// default(empty), windows7, gtk2, flash
	cef := strings.ToLower(c.Install.CEF)
	//if cef == consts.CefEmpty /*&& cef != consts.Cef109 && cef != consts.Cef106 && cef != consts.Cef87*/ {
	//	term.Logger.Error("-c [cef] Incorrect args value")
	//	return "", nil
	//}
	// 安装目录名称
	installPathName := cefInstallPathName(c)
	term.Section.Println("Install Path", installPathName)

	term.Section.Println("Start downloading CEF and Energy dependency")
	// 所有版本列表
	var versionList = edv["versionList"].(map[string]any)

	// 获取到当前安装版本
	var installVersion map[string]any
	if c.Install.Version == "latest" {
		// 获取最新版本号, latest=vx.x.x
		if v, ok := versionList[edv["latest"].(string)]; ok {
			installVersion = v.(map[string]any)
		}
	} else {
		// 自己选择版本
		if v, ok := versionList[c.Install.Version]; ok {
			installVersion = v.(map[string]any)
		}
	}
	term.Section.Println("Check version")
	if installVersion == nil || len(installVersion) == 0 {
		term.Logger.Error("Invalid version number " + c.Install.Version)
		return "", nil
	}
	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefModuleName, liblclModuleName string
	)
	if consts.IsWindows {
		// windows 版本小于10，使用CEF109，CEF 109是最后一个支持windows7的版本
		majorVersion, _, _ := versionNumber()
		if majorVersion < 10 {
			cef = consts.Cef109
		}
	}
	// 使用提供的特定版本号
	// cefModuleName 对应配置字段 module
	// liblclModuleName 对应配置字段
	if cef == consts.Cef106 {
		cefModuleName = "cef-106" // CEF 106.1.1, linux gtk2
	} else if cef == consts.Cef109 {
		cefModuleName = "cef-109"       // CEF 109.1.18
		liblclModuleName = "liblcl-109" // liblcl 109, windows7
	} else if cef == consts.Cef87 {
		// cef 87 要和 liblcl 87 配对
		cefModuleName = "cef-87"       // CEF 87.1.14
		liblclModuleName = "liblcl-87" // liblcl 87, flash
	}
	// 如未指定CEF参数、或参数不正确，选择当前CEF模块最（新）大的版本号
	if cefModuleName == "" {
		var cefDefault string
		var number int
		for module, _ := range installVersion {
			if strings.Index(module, "cef") == 0 {
				if s := strings.Split(module, "-"); len(s) == 2 {
					// module = "cef-xxx"
					n, _ := strconv.Atoi(s[1])
					if n >= number {
						number = n
						cefDefault = module
					}
				} else {
					// module = "cef"
					cefDefault = module
					break
				}
			}
		}
		cefModuleName = cefDefault
	}
	// liblcl, 在未指定flash版本时，它是空 ""
	if liblclModuleName == "" {
		liblclModuleName = "liblcl"
	}
	// 当前安装版本的所有模块
	var modules map[string]any
	if m, ok := installVersion["modules"]; ok {
		modules = m.(map[string]any)
	}
	// 根据模块名拿到对应的模块配置
	var (
		cefModule, liblclModule map[string]any
	)
	if module, ok := modules[cefModuleName]; ok {
		cefModule = module.(map[string]any)
	}
	if module, ok := modules[liblclModuleName]; ok {
		liblclModule = module.(map[string]any)
	}
	if cefModule == nil {
		term.Logger.Error("CEF module " + cefModuleName + " is not configured in the current version")
		return "", nil
	}
	// 下载源选择
	var replaceSource = func(url, source string, sourceSelect int, module string) string {
		s := strings.Split(source, ",")
		// liblcl 如果自己选择下载源
		if module == "liblcl" && c.Install.Download != "" {
			sourceSelect = tools.ToInt(c.Install.Download)
		}
		if len(s) > sourceSelect {
			return strings.ReplaceAll(url, "{source}", strings.TrimSpace(s[sourceSelect]))
		}
		return url
	}
	// 下载集合
	var downloads = make(map[string]*downloadInfo)
	// 根据模块名拿到版本号
	cefVersion := tools.ToRNilString(installVersion[cefModuleName], "")
	// 当前模块版本支持系统，如果支持返回下载地址
	libCEFOS, isSupport := cefOS(c, cefModule)
	downloadCefURL := tools.ToString(cefModule["downloadUrl"])
	downloadCefURL = replaceSource(downloadCefURL, tools.ToString(cefModule["downloadSource"]), tools.ToInt(cefModule["downloadSourceSelect"]), "cef")
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloads[consts.CefKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadCefURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL, module: cefModuleName}

	// liblcl
	// 如果选定的cef 106，在linux会指定liblcl gtk2 版本, 其它系统和版本以默认的形式区分
	// 最后根据模块名称来确定使用哪个liblcl
	liblclVersion := tools.ToRNilString(installVersion[liblclModuleName], "")
	if liblclModule != nil {
		libEnergyOS, isSupport := liblclOS(c, cef, liblclVersion, tools.ToString(liblclModule["buildSupportOSArch"]))
		downloadEnergyURL := tools.ToString(liblclModule["downloadUrl"])
		downloadEnergyURL = replaceSource(downloadEnergyURL, tools.ToString(liblclModule["downloadSource"]), tools.ToInt(liblclModule["defaultSourceSelect"]), "liblcl")
		module := tools.ToString(liblclModule["module"])
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", liblclVersion)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", module)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
		downloads[consts.LiblclKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}
	}

	// 在线下载框架二进制包
	for key, dl := range downloads {
		term.Section.Println("Download", key, ":", dl.url)
		if !dl.isSupport {
			term.Logger.Warn("Warn module is not built or configured [" + dl.module + "]")
			continue
		}
		err = DownloadFile(dl.url, dl.downloadPath, nil)
		if err != nil {
			term.Logger.Error("Download [" + dl.fileName + "] " + err.Error())
			return "", nil
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	term.Logger.Info("Unpack files")
	for key, di := range downloads {
		if !di.isSupport {
			term.Logger.Warn("module is not built or configured [" + di.module + "]")
			continue
		}
		if di.success {
			if key == consts.CefKey {
				processBar, err := pterm.DefaultProgressbar.WithShowCount(false).WithShowPercentage(false).WithMaxWidth(1).Start()
				if err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
				tarName, err := UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
					processBar.UpdateTitle(fmt.Sprintf("Unpack file %s, process: %d", key, processLength)) // Update the title of the progressbar.
				})
				processBar.Stop()
				if err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
				if err := ExtractFiles(key, tarName, di, extractOSConfig); err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
			} else if key == consts.LiblclKey {
				if err := ExtractFiles(key, di.downloadPath, di, extractOSConfig); err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
			}
			term.Section.Println("Unpack file", key, "success")
		}
	}
	return installPathName, func() {
		term.Logger.Info("CEF Installed Successfully", term.Logger.Args("Version", c.Install.Version, "liblcl", liblclVersion))
		if liblclModule == nil {
			term.Section.Println("hint: liblcl module", liblclModuleName, `is not configured in the current version, You need to use built-in binary build. [go build -tags="tempdll"]`)
		}
	}
}

func cefOS(c *command.Config, module map[string]any) (string, bool) {
	buildSupportOSArch := tools.ToString(module["buildSupportOSArch"])
	mod := tools.ToString(module["module"])
	archs := strings.Split(buildSupportOSArch, ",")
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}

	if c.Install.OS.IsWindows() { // windows arm for 64 bit, windows for 32/64 bit
		if c.Install.Arch.IsARM64() {
			return "windowsarm64", isSupport(consts.WindowsARM64)
		}
		if c.Install.Arch.Is386() {
			return "windows32", isSupport(consts.Windows32)
		}
		return "windows64", isSupport(consts.Windows64)
	} else if c.Install.OS.IsLinux() { //linux for 64 bit
		if c.Install.Arch.IsARM64() {
			if mod == consts.Cef106 {
				return "linuxarm64", isSupport(consts.LinuxARM64GTK2)
			}
			return "linuxarm64", isSupport(consts.LinuxARM64) || isSupport(consts.LinuxARM64GTK3)
		} else if c.Install.Arch.IsAMD64() {
			if mod == consts.Cef106 {
				return "linux64", isSupport(consts.Linux64GTK2)
			}
			return "linux64", isSupport(consts.Linux64) || isSupport(consts.Linux64GTK3)
		}
	} else if c.Install.OS.IsDarwin() { // macosx for 64 bit
		//if runtime.GOARCH == "arm64" {
		//	return "macosarm64", isSupport(MacOSARM64)
		//} else if runtime.GOARCH == "amd64" {
		//	return "macosx64", isSupport(MacOSX64)
		//}
		// Mac amd64 m1 m2 架构目前使用amd64, m1,m2使用Rosetta2兼容
		return "macosx64", isSupport(consts.MacOSX64)
	}
	//not support
	return fmt.Sprintf("%v %v", c.Install.OS, c.Install.Arch), false
}

var liblclFileNames = map[string]string{
	"windows386":         consts.Windows32,
	"windowsamd64":       consts.Windows64,
	"windowsarm64":       consts.WindowsARM64,
	"linuxarm64":         consts.LinuxARM64,
	"linuxarm64gtk2":     consts.LinuxARM64GTK2,
	"linuxamd64":         consts.Linux64,
	"linuxamd64gtk2":     consts.Linux64GTK2,
	"darwinamd64":        consts.MacOSX64,
	"darwinarm64":        consts.MacOSARM64,
	"windows386_old":     "Windows 32 bits",
	"windowsamd64_old":   "Windows 64 bits",
	"linuxamd64gtk2_old": "Linux GTK2 x86 64 bits",
	"linuxamd64_old":     "Linux x86 64 bits",
	"darwinamd64_old":    "MacOSX x86 64 bits",
}

func liblclName(c *command.Config, version, cef string) (string, bool) {
	var key string
	var isOld bool
	if c.Install.Arch.IsARM64() {
		if c.Install.OS.IsLinux() && cef == consts.Cef106 { // 只linux区别liblcl gtk2
			key = "linuxarm64gtk2"
		} else {
			if c.Install.OS.IsDarwin() {
				// Mac amd64 m1 m2 架构目前使用amd64, m1,m2使用Rosetta2兼容
				key = fmt.Sprintf("%samd64", c.Install.OS)
			} else {
				key = fmt.Sprintf("%sarm64", c.Install.OS)
			}
		}
	} else {
		if c.Install.OS.IsLinux() && cef == consts.Cef106 { // 只linux区别liblcl gtk2
			key = "linuxamd64gtk2"
		} else {
			key = fmt.Sprintf("%s%s", c.Install.OS, c.Install.Arch)
		}
	}
	if tools.Compare("2.2.4", version) {
		if key != "" {
			key += "_old"
			isOld = true
		}
	}
	if key != "" {
		return liblclFileNames[key], isOld
	}
	return "", false
}

// 命名规则 OS+[ARCH]+BIT+[GTK2]
//  ARCH: 非必需, ARM 时填写, AMD为空
//  GTK2: 非必需, GTK2(Linux CEF 106) 时填写, 非Linux或GTK3时为空
func liblclOS(c *command.Config, cef, version, buildSupportOSArch string) (string, bool) {
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

// 提取文件
func ExtractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig map[string]any) error {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	files := extractOSConfig[keyName].([]any)
	if keyName == consts.CefKey {
		//tar
		return ExtractUnTar(sourcePath, di.frameworkPath, files...)
	} else if keyName == consts.LiblclKey {
		//zip
		return ExtractUnZip(sourcePath, di.frameworkPath, false, files...)
	}
	return errors.New("not module")
}

// url文件名
func urlName(downloadUrl string) string {
	if u, err := url.QueryUnescape(downloadUrl); err != nil {
		return ""
	} else {
		u = u[strings.LastIndex(u, "/")+1:]
		return u
	}
}
