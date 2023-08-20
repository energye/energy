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
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

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
	liblclPath := filepath.Join(dir, golcl)
	fmt.Println("用户目录:", dir)
	if len(os.Args) > 1 {
		zipFileName := os.Args[1]
		fmt.Println("指定压缩文件目录:", zipFileName)
		if strings.ToLower(path.Ext(zipFileName)) != ".zip" {
			panic("输入正确的zip包，如：“liblcl-2.2.2.zip”")
		}
		zz, err := zip.OpenReader(zipFileName)
		if err != nil {
			panic(err)
		}
		defer zz.Close()
		for _, ff := range zz.File {
			//fmt.Println(ff.Name)
			switch ff.Name {
			case "win32/liblcl.dll":
				genresByte(readZipData(ff), "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_386.go"))
			case "win64/liblcl.dll":
				genresByte(readZipData(ff), "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_amd64.go"))
			case "macos64-cocoa/liblcl.dylib":
				genresByte(readZipData(ff), "darwin", filepath.Join(libLCLBinResDir, "liblcl_darwin_amd64.go"))
			case "linux64-gtk3/liblcl.so":
				genresByte(readZipData(ff), "linux && gtk3", filepath.Join(libLCLBinResDir, "liblcl_gtk3_linux_amd64.go"))
			case "linux64-gtk2/liblcl.so":
				genresByte(readZipData(ff), "linux && gtk2", filepath.Join(libLCLBinResDir, "liblcl_gtk2_linux_amd64.go"))
			}
		}
	} else {
		// windows 32
		genresFile(filepath.Join(liblclPath, "win32", "liblcl.dll"), "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_386.go"))
		// windows 64
		genresFile(filepath.Join(liblclPath, "win64", "liblcl.dll"), "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_amd64.go"))
		// windows arm 64
		//genresFile(filepath.Join(liblclPath, "winarm64", "liblcl.dll"), "windows", filepath.Join(libLCLBinResDir, "liblcl_windows_arm64.go"))
		// macos cocoa
		genresFile(filepath.Join(liblclPath, "macos64-cocoa", "liblcl.dylib"), "darwin", filepath.Join(libLCLBinResDir, "liblcl_darwin_amd64.go"))
		// macos arm cocoa
		//genresFile(filepath.Join(liblclPath, "macosarm64-cocoa", "liblcl.dylib"), "darwin", filepath.Join(libLCLBinResDir, "liblcl_darwin_arm64.go"))
		// linux 64 gtk3
		genresFile(filepath.Join(liblclPath, "linux64-gtk3", "liblcl.so"), "linux && gtk3", filepath.Join(libLCLBinResDir, "liblcl_gtk3_linux_amd64.go"))
		// linux 64 gtk2
		genresFile(filepath.Join(liblclPath, "linux64-gtk2", "liblcl.so"), "linux && gtk2", filepath.Join(libLCLBinResDir, "liblcl_gtk2_linux_amd64.go"))
		// linux arm 64 gtk3
		//genresFile(filepath.Join(liblclPath, "linuxarm64-gtk3", "liblcl.so"), "linux && gtk3", filepath.Join(libLCLBinResDir, "liblcl_gtk3_linux_arm64.go"))
		// linux arm 64 gtk2
		//genresFile(filepath.Join(liblclPath, "linuxarm64-gtk2", "liblcl.so"), "linux && gtk2", filepath.Join(libLCLBinResDir, "liblcl_gtk2_linux_arm64.go"))
	}
}

// 生成字节的单元
func genresFile(fileName, tags, newFileName string) {
	bs, err := ioutil.ReadFile(fileName)
	if err == nil {
		genresByte(bs, tags, newFileName)
	} else {
		fmt.Println("生成字节Go文件:", newFileName, "Error:", err)
	}
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

func readZipData(ff *zip.File) []byte {
	if rr, err := ff.Open(); err == nil {
		defer rr.Close()
		bs, err := ioutil.ReadAll(rr)
		if err != nil {
			return nil
		}
		return bs
	}
	return nil
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
