//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/bzip2"
	"encoding/json"
	"fmt"
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var CmdInstall = &Command{
	UsageLine: "install -p [path] -v [version] -n [name] -d [download] -c [cef]",
	Short:     "Automatically configure the CEF and Energy framework",
	Long: `
	-p Installation directory Default current directory
	-v Specifying a version number,Default latest
	-n Name of the frame after installation
	-d Download Source, 0:gitee or 1:github, Default empty
	-c Install system supports CEF version, provide 4 options, default empty
		default : Automatically select support for the latest version based on the current system.
		109 : CEF 109.1.18 is the last one to support Windows 7.
		106 : CEF 106.1.1 is the last default support for GTK2 in Linux.
		87  : CEF 87.1.14 is the last one to support Flash.
	.  Execute default command

Automatically configure the CEF and Energy framework.

During this process, CEF and Energy are downloaded.

Default framework name is "EnergyFramework".
`,
}

type downloadInfo struct {
	fileName      string
	frameworkPath string
	downloadPath  string
	url           string
	success       bool
	isSupport     bool
}

func init() {
	CmdInstall.Run = runInstall
}

const (
	GTK3 = iota + 1
	GTK2
)

// https://cef-builds.spotifycdn.com/cef_binary_107.1.11%2Bg26c0b5e%2Bchromium-107.0.5304.110_windows64.tar.bz2
// 运行安装
func runInstall(c *CommandConfig) error {
	// 获取提取文件配置
	extractData, err := httpRequestGET(DownloadExtractURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	var extractConfig map[string]any
	extractData = bytes.TrimPrefix(extractData, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(extractData, &extractConfig); err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	extractOSConfig := extractConfig[runtime.GOOS].(map[string]any)

	// 获取安装版本配置
	downloadJSON, err := httpRequestGET(DownloadVersionURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	var edv map[string]any
	downloadJSON = bytes.TrimPrefix(downloadJSON, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(downloadJSON, &edv); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	// -c cef args value
	// default(empty), windows7, gtk2, flash
	cef := strings.ToLower(c.Install.CEF)
	if cef != CefEmpty && cef != Cef109 && cef != Cef106 && cef != Cef87 {
		fmt.Fprint(os.Stderr, "-c [cef] Incorrect args value\n")
		os.Exit(1)
	}

	if c.Install.Path == "" {
		// current dir
		c.Install.Path = c.Wd
	}
	installPathName := filepath.Join(c.Install.Path, c.Install.Name)
	println("Install Path", installPathName)
	if c.Install.Version == "" {
		// latest
		c.Install.Version = "latest"
	}
	// 创建安装目录
	os.MkdirAll(c.Install.Path, fs.ModePerm)
	os.MkdirAll(installPathName, fs.ModePerm)
	os.MkdirAll(filepath.Join(c.Install.Path, frameworkCache), fs.ModePerm)
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
		println("Invalid version number:", c.Install.Version)
		os.Exit(1)
	}
	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefModuleName, liblclModuleName string
	)
	// 使用提供的特定版本号
	if cef == Cef106 {
		cefModuleName = "cef-106" // CEF 106.1.1
	} else if cef == Cef109 {
		cefModuleName = "cef-109" // CEF 109.1.18
	} else if cef == Cef87 {
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
		println("error: cef module", cefModuleName, "is not configured in the current version")
		os.Exit(1)
	}
	// 下载源选择
	var replaceSource = func(url, source string, sourceSelect int, module string) string {
		s := strings.Split(source, ",")
		// liblcl 如果自己选择下载源
		if module == "liblcl" && c.Install.Download != "" {
			sourceSelect = ToInt(c.Install.Download)
		}
		if len(s) > sourceSelect {
			return strings.ReplaceAll(url, "{source}", s[sourceSelect])
		}
		return url
	}
	// 下载集合
	var downloads = make(map[string]*downloadInfo)
	// 根据模块名拿到版本号
	cefVersion := ToRNilString(installVersion[cefModuleName], "")
	// 当前模块版本支持系统，如果支持返回下载地址
	libCEFOS, isSupport := cefOS(cefModule)
	downloadCefURL := ToString(cefModule["downloadUrl"])
	downloadCefURL = replaceSource(downloadCefURL, ToString(cefModule["downloadSource"]), ToInt(cefModule["downloadSourceSelect"]), "cef")
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloads[cefKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadCefURL), downloadPath: filepath.Join(c.Install.Path, frameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL}

	// liblcl
	// 如果选定的cef 106，在linux会指定liblcl gtk2 版本, 其它系统和版本以默认的形式区分
	// 最后根据模块名称来确定使用哪个liblcl
	liblclVersion := ToRNilString(installVersion[liblclModuleName], "")
	if liblclModule != nil {
		libEnergyOS, isSupport := liblclOS(cef, liblclVersion, liblclModule)
		downloadEnergyURL := ToString(liblclModule["downloadUrl"])
		downloadEnergyURL = replaceSource(downloadEnergyURL, ToString(liblclModule["downloadSource"]), ToInt(liblclModule["downloadSourceSelect"]), "liblcl")
		module := ToString(liblclModule["module"])
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", liblclVersion)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", module)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
		downloads[liblclKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, frameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL}
	}

	// 在线下载框架二进制包
	for key, dl := range downloads {
		fmt.Printf("Download %s: %s\n", key, dl.url)
		if !dl.isSupport {
			println("module is not built or configured 【", dl.fileName, "]")
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
			println("Download [", dl.fileName, "] error", err)
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	println("Unpack files")
	var removeFileList = make([]string, 0, 0)
	for key, di := range downloads {
		if !di.isSupport {
			println("energy command line does not support the system architecture.")
			continue
		}
		if di.success {
			if key == cefKey {
				bar := progressbar.NewBar(0)
				bar.SetNotice("Unpack file " + key + ": ")
				tarName := UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
					bar.PrintSizeBar(processLength)
				})
				bar.PrintEnd()
				ExtractFiles(key, tarName, di, extractOSConfig)
				removeFileList = append(removeFileList, tarName)
			} else if key == liblclKey {
				ExtractFiles(key, di.downloadPath, di, extractOSConfig)
			}
			println("Unpack file", key, "success\n")
		}
	}
	for _, rmFile := range removeFileList {
		println("Remove file", rmFile)
		os.Remove(rmFile)
	}
	setEnergyHomeEnv(EnergyHomeKey, installPathName)
	println()
	println(CmdInstall.Short, "SUCCESS \nInstalled version:", c.Install.Version, liblclVersion)
	if liblclModule == nil {
		println("hint: liblcl module", liblclModuleName, `is not configured in the current version, You need to use built-in binary build. [go build -tags="tempdll"]`)
	}
	return nil
}

func cefOS(module map[string]any) (string, bool) {
	buildSupportOSArch := ToString(module["buildSupportOSArch"])
	mod := ToString(module["module"])
	archs := strings.Split(buildSupportOSArch, ",")
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}
	if isWindows { // windows arm for 64 bit, windows for 32/64 bit
		if runtime.GOARCH == "arm64" {
			return "windowsarm64", isSupport(WindowsARM64)
		}
		if strconv.IntSize == 32 {
			return fmt.Sprintf("windows%d", strconv.IntSize), isSupport(Windows32)
		}
		return fmt.Sprintf("windows%d", strconv.IntSize), isSupport(Windows64)
	} else if isLinux { //linux for 64 bit
		if runtime.GOARCH == "arm64" {
			if mod == Cef106 {
				return "linuxarm64", isSupport(LinuxARM64GTK2)
			}
			return "linuxarm64", isSupport(LinuxARM64) || isSupport(LinuxARM64GTK3)
		} else if runtime.GOARCH == "amd64" {
			if mod == Cef106 {
				return "linux64", isSupport(Linux64GTK2)
			}
			return "linux64", isSupport(Linux64) || isSupport(Linux64GTK3)
		}
	} else if isDarwin { // macosx for 64 bit
		if runtime.GOARCH == "arm64" {
			return "macosarm64", isSupport(MacOSARM64)
		} else if runtime.GOARCH == "amd64" {
			return "macosx64", isSupport(MacOSX64)
		}
	}
	//not support
	return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
}

var liblclFileNames = map[string]string{
	"windows32":       Windows32,
	"windows64":       Windows64,
	"windowsarm64":    WindowsARM64,
	"linuxarm64":      LinuxARM64,
	"linuxarm64gtk2":  LinuxARM64GTK2,
	"linux64":         Linux64,
	"linux64gtk2":     Linux64GTK2,
	"darwin64":        MacOSX64,
	"darwinarm64":     MacOSARM64,
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
		if isLinux && cef == Cef106 { // 只linux区别liblcl gtk2
			key = "linuxarm64gtk2"
		} else {
			key = fmt.Sprintf("%sarm64", runtime.GOOS)
		}
	} else if runtime.GOARCH == "amd64" {
		if isLinux && cef == Cef106 { // 只linux区别liblcl gtk2
			key = "linux64gtk2"
		} else {
			key = fmt.Sprintf("%s%d", runtime.GOOS, strconv.IntSize)
		}
	}
	if Compare("2.2.4", version) {
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
func liblclOS(cef, version string, module map[string]any) (string, bool) {
	buildSupportOSArch := ToString(module["buildSupportOSArch"])
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
	if keyName == cefKey {
		//tar
		ExtractUnTar(sourcePath, di.frameworkPath, files...)
	} else if keyName == liblclKey {
		//zip
		ExtractUnZip(sourcePath, di.frameworkPath, files...)
	}
}

func filePathInclude(compressPath string, files ...any) (string, bool) {
	for _, file := range files {
		f := file.(string)
		tIdx := strings.LastIndex(f, "/*")
		if tIdx != -1 {
			f = f[:tIdx]
			if f[0] == '/' {
				if strings.Index(compressPath, f[1:]) == 0 {
					return compressPath, true
				}
			}
			if strings.Index(compressPath, f) == 0 {
				return strings.Replace(compressPath, f, "", 1), true
			}
		} else {
			if f[0] == '/' {
				if compressPath == f[1:] {
					return f, true
				}
			}
			if compressPath == f {
				f = f[strings.Index(f, "/")+1:]
				return f, true
			}
		}
	}
	return "", false
}

func dir(path string) string {
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))
	lastSep := strings.LastIndex(path, string(filepath.Separator))
	return path[:lastSep]
}

func ExtractUnTar(filePath, targetPath string, files ...any) {
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error: cannot read tar file, error=[%v]\n", err)
		return
	}
	defer reader.Close()

	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error: cannot read tar file, error=[%v]\n", err)
			os.Exit(1)
			return
		}
		compressPath := header.Name[strings.Index(header.Name, "/")+1:]
		includePath, isInclude := filePathInclude(compressPath, files...)
		if !isInclude {
			continue
		}
		info := header.FileInfo()
		targetFile := filepath.Join(targetPath, includePath)
		fmt.Println("compressPath:", compressPath, "-> targetFile:", targetFile)
		if info.IsDir() {
			if err = os.MkdirAll(targetFile, info.Mode()); err != nil {
				fmt.Printf("error: cannot mkdir file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
		} else {
			fDir := dir(targetFile)
			_, err = os.Stat(fDir)
			if os.IsNotExist(err) {
				if err = os.MkdirAll(fDir, info.Mode()); err != nil {
					fmt.Printf("error: cannot file mkdir file, error=[%v]\n", err)
					os.Exit(1)
					return
				}
			}
			file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
			if err != nil {
				fmt.Printf("error: cannot open file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
			defer file.Close()
			bar := progressbar.NewBar(100)
			bar.SetNotice("\t")
			bar.HideRatio()
			writeFile(tarReader, file, header.Size, func(totalLength, processLength int64) {
				bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
			})
			file.Sync()
			bar.PrintBar(100)
			bar.PrintEnd()
			if err != nil {
				fmt.Printf("error: cannot write file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
		}
	}
}

func ExtractUnZip(filePath, targetPath string, files ...any) {
	if rc, err := zip.OpenReader(filePath); err == nil {
		defer rc.Close()
		for i := 0; i < len(files); i++ {
			if f, err := rc.Open(files[i].(string)); err == nil {
				defer f.Close()
				st, _ := f.Stat()
				targetFileName := filepath.Join(targetPath, st.Name())
				if st.IsDir() {
					os.MkdirAll(targetFileName, st.Mode())
					continue
				}
				if targetFile, err := os.Create(targetFileName); err == nil {
					fmt.Println("extract file: ", st.Name())
					bar := progressbar.NewBar(100)
					bar.SetNotice("\t")
					bar.HideRatio()
					writeFile(f, targetFile, st.Size(), func(totalLength, processLength int64) {
						bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
					})
					bar.PrintBar(100)
					bar.PrintEnd()
					targetFile.Close()
				}
			} else {
				fmt.Printf("error: cannot open file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
		}
	} else {
		if err != nil {
			fmt.Printf("error: cannot read zip file, error=[%v]\n", err)
			os.Exit(1)
		}
	}
}

// 释放bz2文件到tar
func UnBz2ToTar(name string, callback func(totalLength, processLength int64)) string {
	fileBz2, err := os.Open(name)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
	defer fileBz2.Close()
	dirName := fileBz2.Name()
	dirName = dirName[:strings.LastIndex(dirName, ".")]
	r := bzip2.NewReader(fileBz2)
	w, err := os.Create(dirName)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
	defer w.Close()
	writeFile(r, w, 0, callback)
	return dirName
}

func writeFile(r io.Reader, w *os.File, totalLength int64, callback func(totalLength, processLength int64)) {
	buf := make([]byte, 1024*10)
	var written int64
	for {
		nr, err := r.Read(buf)
		if nr > 0 {
			nw, err := w.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(totalLength, written)
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if err != nil {
			break
		}
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

func isFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	os.Remove(filename)
	return false
}

// 下载文件
func downloadFile(url string, localPath string, callback func(totalLength, processLength int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 1024*10)
		written int64
	)
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("download-error=[%v]\n", err)
		os.Exit(1)
		return err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		fmt.Printf("download-error=[%v]\n", err)
		os.Exit(1)
		return err
	}
	if isFileExist(localPath, fsize) {
		println("File already exists")
		return nil
	}
	println("Save path: [", localPath, "] file size:", fsize)
	file, err := os.Create(tmpFilePath)
	if err != nil {
		fmt.Printf("download-error=[%v]\n", err)
		os.Exit(1)
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		fmt.Printf("Download-error=[body is null]\n")
		os.Exit(1)
		return nil
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, err := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(fsize, written)
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	if err == nil {
		file.Sync()
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
		if err != nil {
			return err
		}
	}
	return err
}
