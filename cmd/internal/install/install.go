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
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type downloadInfo struct {
	fileName      string
	frameworkPath string
	downloadPath  string
	url           string
	success       bool
	isSupport     bool
	module        string
}
type softEnf struct {
	name      string
	desc      string
	installed bool
	yes       func()
}

func Install(c *command.Config) error {
	// 设置默认参数
	defaultInstallConfig(c)
	// 检查环境
	willInstall := checkInstallEnv(c)
	var (
		goRoot                      string
		goSuccessCallback           func()
		cefFrameworkRoot            string
		cefFrameworkSuccessCallback func()
		nsisRoot                    string
		nsisSuccessCallback         func()
		upxRoot                     string
		upxSuccessCallback          func()
		z7zRoot                     string
		z7zSuccessCallback          func()
	)
	if len(willInstall) > 0 {
		// energy 依赖的软件框架
		tableData := pterm.TableData{
			{"Software", "Installed", "Support Platform Description"},
		}
		var options []string
		for _, se := range willInstall {
			var installed string
			if se.installed {
				installed = "Yes"
			} else {
				installed = "No"
				options = append(options, se.name)
				//se.yes()
			}
			var data = []string{se.name, installed, se.desc}
			tableData = append(tableData, data)
		}
		term.Section.Println("Energy Development Environment Framework Dependencies")
		err := pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
		if err != nil {
			return err
		}
		// 选择
		if len(options) > 0 {
			printer := term.DefaultInteractiveMultiselect.WithOnInterruptFunc(func() {
				os.Exit(1)
			}).WithOptions(options)
			printer.CheckmarkANSI()
			printer.DefaultText = "Optional Installation"
			printer.Filter = false
			selectedOptions, err := printer.Show()
			if err != nil {
				return err
			}
			term.Section.Printfln("Selected : %s", pterm.Green(selectedOptions))
			for _, option := range selectedOptions {
				for _, wi := range willInstall {
					if option == wi.name {
						wi.yes()
						break
					}
				}
			}
			if len(selectedOptions) > 0 {
				// 初始配置和安装目录
				if err := initInstall(c); err != nil {
					return err
				}
			}
		}
	}
	// 安装Go开发环境
	goRoot, goSuccessCallback = installGolang(c)
	// 设置 go 环境变量
	if goRoot != "" {
		env.SetGoEnv(goRoot)
	}

	// 安装CEF二进制框架
	cefFrameworkRoot, cefFrameworkSuccessCallback = installCEFFramework(c)
	// 设置 energy cef 环境变量
	if cefFrameworkRoot != "" && c.Install.IsSame {
		env.SetEnergyHomeEnv(cefFrameworkRoot)
	}

	// 安装nsis安装包制作工具, 仅windows - amd64
	nsisRoot, nsisSuccessCallback = installNSIS(c)
	// 设置nsis环境变量
	if nsisRoot != "" {
		env.SetNSISEnv(nsisRoot)
	}

	// 安装upx, 内置, 仅windows, linux
	upxRoot, upxSuccessCallback = installUPX(c)
	// 设置upx环境变量
	if upxRoot != "" {
		env.SetUPXEnv(upxRoot)
	}

	// 安装7za
	z7zRoot, z7zSuccessCallback = install7z(c)
	// 设置7za环境变量
	if z7zRoot != "" {
		env.Set7zaEnv(z7zRoot)
	}
	env.SourceEnvFiles()

	// success 输出
	if nsisSuccessCallback != nil || goSuccessCallback != nil || upxSuccessCallback != nil || cefFrameworkSuccessCallback != nil || z7zSuccessCallback != nil {
		term.BoxPrintln("Hint: Reopen the cmd window for command to take effect.")
	}

	if nsisSuccessCallback != nil {
		nsisSuccessCallback()
	}
	if cefFrameworkSuccessCallback != nil {
		cefFrameworkSuccessCallback()
	}
	if goSuccessCallback != nil {
		goSuccessCallback()
	}
	if upxSuccessCallback != nil {
		upxSuccessCallback()
	}
	if z7zSuccessCallback != nil {
		z7zSuccessCallback()
	}
	copyEnergyCMD(goRoot)

	return nil
}

func copyEnergyCMD(goRoot string) {
	term.Logger.Info("Copy energy command-line to GOROOT/bin")
	if goRoot == "" {
		goRoot = os.Getenv("GOROOT")
		if goRoot == "" {
			term.Logger.Error("Install Copy energy command-line, GOROOT not found, Incorrect configuration or please restart term")
			return
		}
	}
	exe, err := os.Executable()
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	energyName := "energy"
	if consts.IsWindows {
		energyName += ".exe"
	}
	energyBin := filepath.Join(goRoot, "bin", energyName)
	if filepath.ToSlash(exe) == filepath.ToSlash(energyBin) {
		term.Logger.Info("current energy")
		return
	}
	if tools.IsExist(energyBin) {
		os.Remove(energyBin)
	}
	src, err := os.Open(exe)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(energyBin, os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
}

func cefInstallPathName(c *command.Config) string {
	if c.Install.IsSame {
		return filepath.Join(c.Install.Path, consts.ENERGY, c.Install.Name)
	} else {
		return filepath.Join(c.Install.Path, consts.ENERGY, fmt.Sprintf("%s_%s%s", c.Install.Name, c.Install.OS, c.Install.Arch))
	}
}

func goInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "go")
}

func nsisInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "nsis")
}

func upxInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "upx")
}

func z7zInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "7za")
}

func nsisCanInstall() bool {
	return consts.IsWindows && !consts.IsARM64
}

func upxCanInstall() bool {
	return (consts.IsWindows && !consts.IsARM64) || consts.IsLinux || consts.IsDarwin
}

func z7zCanInstall() bool {
	return consts.IsWindows && !consts.IsARM64
}

// 检查当前环境
//  golang, nsis, cef, upx
//  golang: all os
//  nsis: windows
//  cef: all os
//  upx: windows amd64, 386, linux amd64, arm64
func checkInstallEnv(c *command.Config) (result []*softEnf) {
	result = make([]*softEnf, 0)
	var check = func(chkInstall func() (string, bool), name string, yes func()) {
		desc, ok := chkInstall()
		result = append(result, &softEnf{name: name, desc: desc, installed: ok, yes: yes})
	}

	// go
	check(func() (string, bool) {
		return "All", tools.CommandExists("go")
	}, "Golang", func() {
		c.Install.IGolang = true //yes callback
	})

	// cef
	var cefName = "CEF Framework"
	if !c.Install.IsSame {
		cefName = fmt.Sprintf("CEF Framework %s%s", c.Install.OS, c.Install.Arch)
	}
	check(func() (string, bool) {
		if c.Install.IsSame {
			// 检查环境变量是否配置
			return "All", tools.CheckCEFDir()
		}
		// 非当月系统架构时检查一下目标安装路径是否已经存在
		var lib = func() string {
			if c.Install.OS.IsWindows() {
				return "libcef.dll"
			} else if c.Install.OS.IsLinux() {
				return "libcef.so"
			} else if c.Install.OS.IsDarwin() {
				return "cef_sandbox.a"
			}
			return ""
		}()
		s := filepath.Join(cefInstallPathName(c), lib)
		return "All", tools.IsExist(s)
	}, cefName, func() {
		c.Install.ICEF = true //yes callback
	})
	if nsisCanInstall() {
		// nsis
		check(func() (string, bool) {
			if nsisCanInstall() {
				return "Windows AMD", tools.CommandExists("makensis")
			} else {
				return "Non Windows AMD skipping NSIS.", true
			}
		}, "NSIS", func() {
			c.Install.INSIS = true //yes callback
		})
	}
	if upxCanInstall() {
		// upx
		check(func() (string, bool) {
			if upxCanInstall() {
				if consts.IsDarwin {
					if tools.CommandExists("upx") {
						return "All", true
					}
					return "Install: brew install upx", true
				}
				return "All", tools.CommandExists("upx")
			} else {
				return "Unsupported platform UPX.", true
			}
		}, "UPX", func() {
			c.Install.IUPX = true //yes callback
		})
	}
	if z7zCanInstall() {
		// 7za
		check(func() (string, bool) {
			if z7zCanInstall() {
				return "Windows AMD", tools.CommandExists("7za")
			} else {
				return "Non Windows skipping UPX.", true
			}
		}, "7za", func() {
			c.Install.I7za = true //yes callback
		})
	}
	if consts.IsLinux {
		// gtk2
		// gtk3
		// dpkg
	}
	if consts.IsDarwin {
		//
	}
	return
}

func defaultInstallConfig(c *command.Config) {
	if c.Install.Path == "" {
		// current dir
		c.Install.Path = c.Wd
	}
	if c.Install.Version == "" {
		// latest
		c.Install.Version = "latest"
	}
	if c.Install.OS == "" {
		c.Install.OS = command.OS(runtime.GOOS)
	}
	if c.Install.Arch == "" {
		c.Install.Arch = command.Arch(runtime.GOARCH)
	}
	if string(c.Install.OS) == runtime.GOOS && string(c.Install.Arch) == runtime.GOARCH {
		c.Install.IsSame = true
	}
}

func initInstall(c *command.Config) (err error) {
	// 创建安装目录
	err = os.MkdirAll(c.Install.Path, fs.ModePerm) // framework root
	if err != nil {
		return
	}
	if c.Install.ICEF {
		err = os.MkdirAll(cefInstallPathName(c), fs.ModePerm) //cef
		if err != nil {
			return
		}
	}
	if c.Install.IGolang {
		err = os.MkdirAll(goInstallPathName(c), fs.ModePerm) // go
		if err != nil {
			return
		}
	}
	if c.Install.INSIS {
		if nsisCanInstall() {
			err = os.MkdirAll(nsisInstallPathName(c), fs.ModePerm) // nsis
			if err != nil {
				return
			}
		}
	}
	if c.Install.IUPX {
		if upxCanInstall() {
			err = os.MkdirAll(upxInstallPathName(c), fs.ModePerm) //upx
			if err != nil {
				return
			}
		}
	}
	if c.Install.I7za {
		if z7zCanInstall() {
			err = os.MkdirAll(z7zInstallPathName(c), fs.ModePerm) //upx
			if err != nil {
				return
			}
		}
	}
	// framework download cache
	err = os.MkdirAll(filepath.Join(c.Install.Path, consts.FrameworkCache), fs.ModePerm)
	return err
}

func filePathInclude(compressPath string, files ...any) (string, bool) {
	if len(files) == 0 {
		return compressPath, true
	} else {
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
	}
	return "", false
}

func dir(path string) string {
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))
	lastSep := strings.LastIndex(path, string(filepath.Separator))
	return path[:lastSep]
}

func tarFileReader(filePath string) (*tar.Reader, func(), error) {
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error: cannot open file, error=[%v]\n", err)
		return nil, nil, err
	}
	if filepath.Ext(filePath) == ".gz" {
		gr, err := gzip.NewReader(reader)
		if err != nil {
			fmt.Printf("error: cannot open gzip file, error=[%v]\n", err)
			return nil, nil, err
		}
		return tar.NewReader(gr), func() {
			gr.Close()
			reader.Close()
		}, nil
	} else {
		return tar.NewReader(reader), func() {
			reader.Close()
		}, nil
	}
}

func tarFileCount(filePath string) (int, error) {
	tarReader, clos, err := tarFileReader(filePath)
	if err != nil {
		return 0, err
	}
	defer clos()
	var count int
	for {
		_, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

func ExtractUnTar(filePath, targetPath string, files ...any) error {
	term.Logger.Info("Read Files Number")
	fileCount, err := tarFileCount(filePath)
	println(fileCount)
	if err != nil {
		return err
	}
	tarReader, clos, err := tarFileReader(filePath)
	if err != nil {
		return err
	}
	defer clos()
	multi := pterm.DefaultMultiPrinter
	defer multi.Stop()
	fileTotalProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
	writeFileProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
	extractFilesProcessBar, err := fileTotalProcessBar.WithTotal(fileCount).Start("Extract File")
	if err != nil {
		return err
	}
	multi.Start()
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		extractFilesProcessBar.Increment()
		// 去除压缩包内的一级目录
		compressPath := filepath.Clean(header.Name[strings.Index(header.Name, "/")+1:])
		includePath, isInclude := filePathInclude(compressPath, files...)
		if !isInclude {
			continue
		}
		info := header.FileInfo()
		targetFile := filepath.Join(targetPath, includePath)
		if info.IsDir() {
			if err = os.MkdirAll(targetFile, info.Mode()); err != nil {
				return err
			}
		} else {
			fDir := dir(targetFile)
			_, err = os.Stat(fDir)
			if os.IsNotExist(err) {
				if err = os.MkdirAll(fDir, info.Mode()); err != nil {
					return err
				}
			}
			file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
			if err != nil {
				return err
			}
			var (
				total = 100
				c     int
				cn    int
			)
			_, fileName := filepath.Split(targetFile)
			wfpb, err := writeFileProcessBar.WithCurrent(0).WithTotal(total).Start("Write File " + fileName)
			if err != nil {
				return err
			}
			writeFile(tarReader, file, header.Size, func(totalLength, processLength int64) {
				process := int((float64(processLength) / float64(totalLength)) * 100)
				if process > c {
					c = process
					wfpb.Add(process)
					cn++
				}
			})
			if cn < total {
				wfpb.Add(total - cn)
			}
			wfpb.Stop()
			file.Sync()
			file.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ExtractUnZip(filePath, targetPath string, rmRootDir bool, files ...any) error {
	if rc, err := zip.OpenReader(filePath); err == nil {
		defer rc.Close()
		multi := pterm.DefaultMultiPrinter
		defer multi.Stop()
		fileTotalProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
		writeFileProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
		var createWriteFile = func(info fs.FileInfo, path string, file io.Reader) error {
			targetFileName := filepath.Join(targetPath, path)
			if info.IsDir() {
				os.MkdirAll(targetFileName, info.Mode())
				return nil
			}
			fDir, name := filepath.Split(targetFileName)
			if !tools.IsExist(fDir) {
				os.MkdirAll(fDir, 0755)
			}
			if targetFile, err := os.Create(targetFileName); err == nil {
				defer targetFile.Close()
				var (
					total = 100
					c     int
					cn    int
				)
				wfpb, err := writeFileProcessBar.WithCurrent(0).WithTotal(total).Start("Write File " + name)
				if err != nil {
					return err
				}
				writeFile(file, targetFile, info.Size(), func(totalLength, processLength int64) {
					process := int((float64(processLength) / float64(totalLength)) * 100)
					if process > c {
						c = process
						wfpb.Add(process)
						cn++
					}
				})
				if cn < total {
					wfpb.Add(total - cn)
				}
				wfpb.Stop()
				return nil
			} else {
				return err
			}
		}
		// 所有文件
		if len(files) == 0 {
			zipFiles := rc.File
			extractFilesProcessBar, err := fileTotalProcessBar.WithTotal(len(zipFiles)).Start("Extract File")
			if err != nil {
				return err
			}
			multi.Start()
			for _, f := range zipFiles {
				extractFilesProcessBar.Increment() // +1
				r, _ := f.Open()
				var name string
				if rmRootDir {
					// 移除压缩包内的根文件夹名
					name = filepath.Clean(f.Name[strings.Index(f.Name, "/")+1:])
				} else {
					name = filepath.Clean(f.Name)
				}
				if err := createWriteFile(f.FileInfo(), name, r.(io.Reader)); err != nil {
					return err
				}
				_ = r.Close()
			}
			return nil
		} else {
			extractFilesProcessBar, err := fileTotalProcessBar.WithTotal(len(files)).Start("Extract File")
			if err != nil {
				return err
			}
			multi.Start()
			// 指定名字的文件
			for i := 0; i < len(files); i++ {
				extractFilesProcessBar.Increment() // +1
				if f, err := rc.Open(files[i].(string)); err == nil {
					info, _ := f.Stat()
					if err := createWriteFile(info, files[i].(string), f); err != nil {
						return err
					}
					_ = f.Close()
				} else {
					return err
				}
			}
			return nil
		}
	} else {
		return err
	}
}

// 释放bz2文件到tar
func UnBz2ToTar(name string, callback func(totalLength, processLength int64)) (string, error) {
	fileBz2, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer fileBz2.Close()
	dirName := fileBz2.Name()
	dirName = dirName[:strings.LastIndex(dirName, ".")]
	if !tools.IsExist(dirName) {
		r := bzip2.NewReader(fileBz2)
		w, err := os.Create(dirName)
		if err != nil {
			return "", err
		}
		defer w.Close()
		writeFile(r, w, 0, callback)
	} else {
		term.Section.Println("File already exists")
	}
	return dirName, nil
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

// DownloadFile 下载文件
//  如果文件存在大小一样不再下载
func DownloadFile(url string, localPath string, callback func(totalLength, processLength int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 1024*10)
		written int64
	)
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return err
	}
	if isFileExist(localPath, fsize) {
		term.Section.Println("File already exists")
		return nil
	}
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	total := 100
	_, fileName := filepath.Split(localPath)
	if len(fileName) > 70 {
		c := len(fileName) - 70
		for i := 0; i < c; i++ {

		}
	}
	p, err := pterm.DefaultProgressbar.WithMaxWidth(80).WithTotal(total).WithTitle("Download " + fileName).Start()
	if err != nil {
		return err
	}
	var stop = func() {
		if p != nil {
			p.Stop()
			p = nil
		}
	}
	defer stop()
	var (
		count int
		cn    int
	)
	var nw int
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, err = file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if callback != nil {
				callback(fsize, written)
			} else {
				process := int((float64(written) / float64(fsize)) * 100)
				if process > count {
					count = process
					p.Add(1)
					cn++
				}
			}
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
	if cn < total {
		p.Add(total - cn)
	}
	stop()
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
