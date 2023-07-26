package main

import (
	"fmt"
	"github.com/energye/energy/v2/example/build-examples/syso"
	"github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const (
	isWindows = runtime.GOOS == "windows" //support
	isLinux   = runtime.GOOS == "linux"   //support
	isDarwin  = runtime.GOOS == "darwin"  //support
)

func main() {
	wd, _ := os.Getwd()
	println("current:", wd)
	dist := filepath.Join(wd, "example", "dist")
	if !isExist(dist) {
		os.MkdirAll(dist, os.ModePerm)
	}
	examples := []string{"context-menu", "control", "cookie", "cookie-manager", "custom-drag-window", "dev-tools",
		"dom-visitor", "download", "drag-file", "execute-dev-tools-method", "execute-javascript", "flash-for-cef-v87",
		"frameless", "i18n", "internal-http-server", "ipc-on-emit/go-composite-type", "ipc-on-emit/go-to-js", "ipc-on-emit/js-to-go",
		"key-event", "lib-checkupdate", "load-html-url",
		"msgbox", "osr/linux", "osr/windows", "popup-sub-window", "print-pdf", "process-type", "proxy", "response-filter",
		"scheme", "screen", "search-text", "simple", "sub-process/main-process", "sub-process/sub-process", "sys-dialog",
		"sys-menu", "tempdll", "view-source", "vue",
		"webkit-register", "zoom"}
	var (
		ext     string
		ldflags string
	)

	if isWindows {
		examples = append(examples, "windows/custom-browser-create", "windows/transparent")
		examples = append(examples, "tray/lclceftray", "tray/lcltray", "tray/systray", "tray/lclvftray")
		ext = ".exe"
		ldflags = `-H windowsgui -s -w`
	} else if isLinux || isDarwin {
		examples = append(examples, "tray/lcltray", "tray/systray")
		ldflags = `-s -w`
	}
	cmd := command.NewCMD()
	for i, example := range examples {
		dir := filepath.Join(wd, "example", example)
		if isExist(dir) {
			cmd.Dir = dir
			copySyso(dir)
			out := filepath.Join(dist, example+ext)
			println("build example", example, fmt.Sprintf("%d/%d", i+1, len(examples)), "\n\tbuild-dir:", dir, "\n\tout-dir:", out)
			cmd.Command("go", "build", "-ldflags", ldflags, "-o", out, `-tags=tempdll`)
			removeSyso(dir)
			println()
		} else {
			println("not found:", dir)
		}
	}
	println("build end.")
	if isWindows {
		print(`Build sample completed, do you want to use upx tool for compression
	Need to install upx compression tool: https://github.com/upx/upx/releases
Input 1 IS: `)
		var useUPX string
		fmt.Scan(&useUPX)
		if useUPX == "1" {
			for _, example := range examples {
				dir := filepath.Join(wd, "example", example)
				if isExist(dir) {
					cmd.Dir = dir
					out := filepath.Join(dist, example+ext)
					cmd.Command("upx", out) // upx 压缩
					println()
				}
			}
		} else {
			println("exit")
		}
	}
	cmd.Close()
}

func upx() {

}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func sysoname() string {
	if runtime.GOARCH == "amd64" {
		return "example_windows_amd64.syso"
	} else if runtime.GOARCH == "386" {
		return "example_windows_386.syso"
	}
	return ""
}

func copySyso(dir string) {
	if !syso.Syso {
		return
	}
	if sysoname() != "" {
		out := filepath.Join(dir, sysoname())
		if runtime.GOARCH == "amd64" {
			ioutil.WriteFile(out, syso.SysoBytesx64, 0666)
		} else if runtime.GOARCH == "386" {
			ioutil.WriteFile(out, syso.SysoBytes386, 0666)
		}
	}
}

func removeSyso(dir string) {
	if !syso.Syso {
		return
	}
	if sysoname() != "" {
		file := filepath.Join(dir, sysoname())
		if runtime.GOARCH == "amd64" {
			os.Remove(file)
		} else if runtime.GOARCH == "386" {
			os.Remove(file)
		}
	}
}
