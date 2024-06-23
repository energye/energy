//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package assets energy internal assets
package assets

// 说明: go:generate, 利用Go: go generate 命令执行 go-bindata 生成资源文件
// 前提条件：
//   获取 go-bindata 命令工具 -> go get -u github.com/go-bindata/go-bindata/...
// 参数说明:
//       -fs 生成文件实例用来获取文件, -o 输出的go文件字节资源文件, -pkg 生成的包名称 ./assets要生成的打包目录
// 生成命令:
//		 到指定目录下执行 go generate
//go:generate go-bindata -fs -o=iconassets.go -pkg=assets ./assets

// DefaultPNGICON energy app default icon.png
func DefaultPNGICON() []byte {
	if d, err := __assetsFile__.ReadFile("assets/icon.png"); err == nil {
		return d
	}
	return nil
}

// DefaultICOICON energy app default icon.ico
func DefaultICOICON() []byte {
	if d, err := __assetsFile__.ReadFile("assets/icon.ico"); err == nil {
		return d
	}
	return nil
}
