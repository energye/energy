package main

import (
	"fmt"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/tools/command"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	wd := consts.ExeDir
	exePath := filepath.Join(wd, "example", "app-icon-gen")     // 当前执行目录
	resourcePath := filepath.Join(exePath, "resources")         // 资源 icon.png 存放目录
	savePath := filepath.Join(resourcePath, "bytes.txt")        // icns 字节数组
	iconSetPath := filepath.Join(resourcePath, "icons.iconset") // 生成 icon 图标集合目录
	icnsPath := filepath.Join(resourcePath, "icon.icns")        // 生成 icon.icns 目录
	os.Remove(iconSetPath)
	os.Remove(savePath)
	os.Remove(icnsPath)
	os.MkdirAll(iconSetPath, fs.ModePerm)
	//sips
	sipsCmds := []string{
		"-z 16 16 icon.png -o icons.iconset/icon_16x16.png",
		"-z 32 32 icon.png -o icons.iconset/icon_16x16@2x.png",
		"-z 32 32 icon.png -o icons.iconset/icon_32x32.png",
		"-z 64 64 icon.png -o icons.iconset/icon_32x32@2x.png",
		"-z 128 128 icon.png -o icons.iconset/icon_128x128.png",
		"-z 256 256 icon.png -o icons.iconset/icon_128x128@2x.png",
		"-z 256 256 icon.png -o icons.iconset/icon_256x256.png",
		"-z 512 512 icon.png -o icons.iconset/icon_256x256@2x.png",
		"-z 512 512 icon.png -o icons.iconset/icon_512x512.png",
		"-z 1024 1024 icon.png -o icons.iconset/icon_512x512@2x.png",
	}
	cmd := command.NewCMD()
	cmd.Dir = resourcePath
	// 生成图标集合
	for _, arg := range sipsCmds {
		cmd.Command("sips", strings.Split(arg, " ")...)
	}
	// 生成 icns
	cmd.Command("iconutil", strings.Split("-c icns icons.iconset -o icon.icns", " ")...)

	// 保存 icns bytes
	saveFile, _ := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	data, _ := ioutil.ReadFile(icnsPath)
	for _, b := range data {
		saveFile.WriteString(fmt.Sprintf("0x%02x", b))
		saveFile.WriteString(", ")
	}
	saveFile.Close()
}
