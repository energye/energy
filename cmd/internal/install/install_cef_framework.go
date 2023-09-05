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
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func installCEFFramework(c *command.Config) (string, func()) {
	if tools.CheckCEFDir() {
		println("CEF Framework installed")
		return "", nil
	}
	print("CEF Framework is not installed. Determine whether to install CEF Framework? Y/n: ")
	var s string
	if strings.ToLower(c.Install.All) != "y" {
		fmt.Scanln(&s)
		if strings.ToLower(s) != "y" {
			println("CEF Framework install exit")
			return "", nil
		}
	}
	// 获取提取文件配置
	extractData, err := tools.HttpRequestGET(consts.DownloadExtractURL)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
		return "", nil
	}
	var extractConfig map[string]any
	extractData = bytes.TrimPrefix(extractData, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(extractData, &extractConfig); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
		return "", nil
	}
	extractOSConfig := extractConfig[runtime.GOOS].(map[string]any)

	// 获取安装版本配置
	downloadJSON, err := tools.HttpRequestGET(consts.DownloadVersionURL)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
		return "", nil
	}
	var edv map[string]any
	downloadJSON = bytes.TrimPrefix(downloadJSON, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(downloadJSON, &edv); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
		return "", nil
	}

	// -c cef args value
	// default(empty), windows7, gtk2, flash
	cef := strings.ToLower(c.Install.CEF)
	if cef != consts.CefEmpty && cef != consts.Cef109 && cef != consts.Cef106 && cef != consts.Cef87 {
		fmt.Println("Error:", "-c [cef] Incorrect args value")
		os.Exit(1)
		return "", nil
	}
	installPathName := cefInstallPathName(c)
	println("Install Path", installPathName)

	println("Start downloading CEF and Energy dependency")
	// 所有版本列表
	var versionList = edv["versionList"].(map[string]any)

	// 当前安装版本
	var installVersion map[string]any
	if c.Install.Version == "latest" {
		// 默认最新版本
		if v, ok := versionList[edv["latest"].(string)]; ok {
			installVersion = v.(map[string]any)
		}
	} else {
		// 自己选择版本
		if v, ok := versionList[c.Install.Version]; ok {
			installVersion = v.(map[string]any)
		}
	}
	println("Check version")
	if installVersion == nil || len(installVersion) == 0 {
		fmt.Println("Error:", "Invalid version number ", c.Install.Version)
		os.Exit(1)
		return "", nil
	}
	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefModuleName, liblclModuleName string
	)
	// 使用提供的特定版本号
	if cef == consts.Cef106 {
		cefModuleName = "cef-106" // CEF 106.1.1
	} else if cef == consts.Cef109 {
		cefModuleName = "cef-109" // CEF 109.1.18
	} else if cef == consts.Cef87 {
		// cef 87 要和 liblcl 87 配对
		cefModuleName = "cef-87"       // CEF 87.1.14
		liblclModuleName = "liblcl-87" // liblcl 87
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
		fmt.Println("Error:", "CEF module", cefModuleName, "is not configured in the current version")
		os.Exit(1)
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
			return strings.ReplaceAll(url, "{source}", s[sourceSelect])
		}
		return url
	}
	// 下载集合
	var downloads = make(map[string]*downloadInfo)
	// 根据模块名拿到版本号
	cefVersion := tools.ToRNilString(installVersion[cefModuleName], "")
	// 当前模块版本支持系统，如果支持返回下载地址
	libCEFOS, isSupport := cefOS(cefModule)
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
		libEnergyOS, isSupport := liblclOS(cef, liblclVersion, tools.ToString(liblclModule["buildSupportOSArch"]))
		downloadEnergyURL := tools.ToString(liblclModule["downloadUrl"])
		downloadEnergyURL = replaceSource(downloadEnergyURL, tools.ToString(liblclModule["downloadSource"]), tools.ToInt(liblclModule["downloadSourceSelect"]), "liblcl")
		module := tools.ToString(liblclModule["module"])
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", liblclVersion)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", module)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
		downloads[consts.LiblclKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}
	}

	// 在线下载框架二进制包
	for key, dl := range downloads {
		fmt.Printf("Download %s: %s\n", key, dl.url)
		if !dl.isSupport {
			println("Warn module is not built or configured 【", dl.module, "】")
			continue
		}
		bar := progressbar.NewBar(100)
		bar.SetNotice("\t")
		bar.HideRatio()
		err = downloadFile(dl.url, dl.downloadPath, func(totalLength, processLength int64) {
			bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
		})
		bar.PrintEnd("Download [" + dl.fileName + "] success")
		if err != nil {
			fmt.Println("Error:", "Download [", dl.fileName, "]", err.Error())
			os.Exit(1)
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	println("Unpack files")
	for key, di := range downloads {
		if !di.isSupport {
			println("Warn module is not built or configured 【", di.module, "】")
			continue
		}
		if di.success {
			if key == consts.CefKey {
				bar := progressbar.NewBar(0)
				bar.SetNotice("Unpack file " + key + ": ")
				tarName := UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
					bar.PrintSizeBar(processLength)
				})
				bar.PrintEnd()
				ExtractFiles(key, tarName, di, extractOSConfig)
			} else if key == consts.LiblclKey {
				ExtractFiles(key, di.downloadPath, di, extractOSConfig)
			}
			println("Unpack file", key, "success\n")
		}
	}
	return installPathName, func() {
		println("\nCEF Installed Successfully \nInstalled version:", c.Install.Version, liblclVersion)
		if liblclModule == nil {
			println("hint: liblcl module", liblclModuleName, `is not configured in the current version, You need to use built-in binary build. [go build -tags="tempdll"]`)
		}
	}
}

func cefOS(module map[string]any) (string, bool) {
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
	if consts.IsWindows { // windows arm for 64 bit, windows for 32/64 bit
		if runtime.GOARCH == "arm64" {
			return "windowsarm64", isSupport(consts.WindowsARM64)
		}
		if strconv.IntSize == 32 {
			return fmt.Sprintf("windows%d", strconv.IntSize), isSupport(consts.Windows32)
		}
		return fmt.Sprintf("windows%d", strconv.IntSize), isSupport(consts.Windows64)
	} else if consts.IsLinux { //linux for 64 bit
		if runtime.GOARCH == "arm64" {
			if mod == consts.Cef106 {
				return "linuxarm64", isSupport(consts.LinuxARM64GTK2)
			}
			return "linuxarm64", isSupport(consts.LinuxARM64) || isSupport(consts.LinuxARM64GTK3)
		} else if runtime.GOARCH == "amd64" {
			if mod == consts.Cef106 {
				return "linux64", isSupport(consts.Linux64GTK2)
			}
			return "linux64", isSupport(consts.Linux64) || isSupport(consts.Linux64GTK3)
		}
	} else if consts.IsDarwin { // macosx for 64 bit
		//if runtime.GOARCH == "arm64" {
		//	return "macosarm64", isSupport(MacOSARM64)
		//} else if runtime.GOARCH == "amd64" {
		//	return "macosx64", isSupport(MacOSX64)
		//}
		// Mac amd64 m1 m2 架构目前使用amd64, m1,m2使用Rosetta2兼容
		return "macosx64", isSupport(consts.MacOSX64)
	}
	//not support
	return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
}

var liblclFileNames = map[string]string{
	"windows32":       consts.Windows32,
	"windows64":       consts.Windows64,
	"windowsarm64":    consts.WindowsARM64,
	"linuxarm64":      consts.LinuxARM64,
	"linuxarm64gtk2":  consts.LinuxARM64GTK2,
	"linux64":         consts.Linux64,
	"linux64gtk2":     consts.Linux64GTK2,
	"darwin64":        consts.MacOSX64,
	"darwinarm64":     consts.MacOSARM64,
	"windows32_old":   "Windows 32 bits",
	"windows64_old":   "Windows 64 bits",
	"linux64gtk2_old": "Linux GTK2 x86 64 bits",
	"linux64_old":     "Linux x86 64 bits",
	"darwin64_old":    "MacOSX x86 64 bits",
}

func liblclName(version, cef string) (string, bool) {
	var key string
	var isOld bool
	if runtime.GOARCH == "arm64" {
		if consts.IsLinux && cef == consts.Cef106 { // 只linux区别liblcl gtk2
			key = "linuxarm64gtk2"
		} else {
			if consts.IsDarwin {
				// Mac amd64 m1 m2 架构目前使用amd64, m1,m2使用Rosetta2兼容
				key = fmt.Sprintf("%samd64", runtime.GOOS)
			} else {
				key = fmt.Sprintf("%sarm64", runtime.GOOS)
			}
		}
	} else if runtime.GOARCH == "amd64" {
		if consts.IsLinux && cef == consts.Cef106 { // 只linux区别liblcl gtk2
			key = "linux64gtk2"
		} else {
			key = fmt.Sprintf("%s%d", runtime.GOOS, strconv.IntSize)
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
func liblclOS(cef, version, buildSupportOSArch string) (string, bool) {
	archs := strings.Split(buildSupportOSArch, ",")
	noSuport := fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH)
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}
	if name, isOld := liblclName(version, cef); isOld {
		if name == "" {
			return noSuport, false
		}
		return name, true
	} else {
		return name, isSupport(name)
	}
}

// 提取文件
func ExtractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig map[string]any) {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	files := extractOSConfig[keyName].([]any)
	if keyName == consts.CefKey {
		//tar
		ExtractUnTar(sourcePath, di.frameworkPath, files...)
	} else if keyName == consts.LiblclKey {
		//zip
		ExtractUnZip(sourcePath, di.frameworkPath, false, files...)
	}
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
