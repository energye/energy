//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"archive/zip"
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/energye/golcl/energy/homedir"
	"hash/crc32"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const liblclVersion = "v2.2.4" // liblcl发布版本

const (
	liblclbinres = "/src/github.com/energye/liblclbinres"
	golcl        = "golcl"
)

func main() {
	gopaths := os.Getenv("GOPATH")
	if gopaths == "" {
		panic("GOPATH为空！")
	}
	sp := strings.Split(gopaths, ";")
	libLCLBinResDir := strings.TrimSpace(sp[0])
	for _, s := range sp {
		s = strings.TrimSpace(s) + liblclbinres
		if fileExists(s) {
			libLCLBinResDir = s
			break
		}
	}
	fmt.Println("找到路径:", libLCLBinResDir)
	if !fileExists(libLCLBinResDir) {
		if err := os.MkdirAll(libLCLBinResDir, 0666); err != nil {
			panic(err)
		}
	}
	// 用户目录
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	liblclPath := filepath.Join(dir, golcl, "liblcl")
	fmt.Println("用户目录:", dir)
	finfo, err := ioutil.ReadDir(liblclPath)
	if err != nil {
		panic(err)
	}

	for _, info := range finfo {
		zipPath := filepath.Join(liblclPath, info.Name())
		zz, err := zip.OpenReader(zipPath)
		if err != nil {
			panic(err)
		}
		defer zz.Close()
		var (
			file fs.File
		)
		name := strings.ToLower(info.Name())
		if strings.Contains(name, "windows") {
			file, err = zz.Open("liblcl.dll")
		} else if strings.Contains(name, "linux") {
			file, err = zz.Open("liblcl.so")
		} else if strings.Contains(name, "macos") {
			file, err = zz.Open("liblcl.dylib")
		}
		if err != nil {
			panic(err)
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		if strings.Contains(name, "windows32.zip") {
			genresByte(data, "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_386.go"))
		} else if strings.Contains(name, "windows64.zip") {
			genresByte(data, "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_amd64.go"))
		} else if strings.Contains(name, "windowsarm64.zip") {
			genresByte(data, "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_arm64.go"))
		} else if strings.Contains(name, "linux64.zip") {
			genresByte(data, "linux && gtk3", filepath.Join(libLCLBinResDir, "liblcl_gtk3_linux_amd64.go"))
		} else if strings.Contains(name, "linux64gtk2.zip") {
			genresByte(data, "linux && gtk2", filepath.Join(libLCLBinResDir, "liblcl_gtk2_linux_amd64.go"))
		} else if strings.Contains(name, "linuxarm64.zip") {
			genresByte(data, "linux && gtk3", filepath.Join(libLCLBinResDir, "liblcl_gtk3_linux_arm64.go"))
		} else if strings.Contains(name, "linuxarm64gtk2.zip") {
			genresByte(data, "linux && gtk2", filepath.Join(libLCLBinResDir, "liblcl_gtk2_linux_arm64.go"))
		} else if strings.Contains(name, "macosarm64.zip") {
			genresByte(data, "darwin", filepath.Join(libLCLBinResDir, "liblcl_darwin_arm64.go"))
		} else if strings.Contains(name, "macosx64.zip") {
			genresByte(data, "darwin", filepath.Join(libLCLBinResDir, "liblcl_darwin_amd64.go"))
		}
	}
	genresLiblclVersion(libLCLBinResDir, liblclVersion)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//  zlib压缩
func zlibCompress(input []byte) ([]byte, error) {
	var in bytes.Buffer
	w, err := zlib.NewWriterLevel(&in, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(input)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}

func genresByte(input []byte, tags, newFileName string) {
	fmt.Println("genFile: ", newFileName)
	if len(input) == 0 {
		fmt.Println("000000")
		return
	}
	crc32Val := crc32.ChecksumIEEE(input)
	//压缩
	bs, err := zlibCompress(input)
	if err != nil {
		panic(err)
	}
	code := bytes.NewBuffer(nil)
	code.WriteString("//go:build ")
	code.WriteString(tags)
	code.WriteString("\r\n\r\n")
	code.WriteString("package liblclbinres")
	code.WriteString("\r\n\r\n")
	code.WriteString(fmt.Sprintf("const CRC32Value uint32 = 0x%x\r\n\r\n", crc32Val))

	code.WriteString("var LCLBinRes = []byte(\"")
	for _, b := range bs {
		code.WriteString("\\x" + fmt.Sprintf("%.2x", b))
	}
	code.WriteString("\")\r\n")
	ioutil.WriteFile(newFileName, code.Bytes(), 0666)
}

func genresLiblclVersion(libLCLBinResDir, version string) {
	code := bytes.NewBuffer(nil)
	code.WriteString("package liblclbinres")
	code.WriteString("\r\n\r\n")
	code.WriteString(`const version = "` + version + `"`)
	code.WriteString("\r\n\r\n")
	code.WriteString("func LibVersion() string {")
	code.WriteString("\n\t")
	code.WriteString("return version")
	code.WriteString("\n}")
	ioutil.WriteFile(filepath.Join(libLCLBinResDir, "liblcl.go"), code.Bytes(), 0666)
}
