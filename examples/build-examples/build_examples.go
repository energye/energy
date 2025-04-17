package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/build-examples/syso"
	"github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	isWindows = runtime.GOOS == "windows" //support
	isLinux   = runtime.GOOS == "linux"   //support
	isDarwin  = runtime.GOOS == "darwin"  //support
)

func main() {
	wd := consts.CurrentExecuteDir
	spDir := strings.Index(wd, "example")
	if spDir > 0 {
		wd = wd[:spDir]
	}
	println("current:", wd)
	dist := filepath.Join(wd, "example", "dist")
	if !isExist(dist) {
		os.MkdirAll(dist, os.ModePerm)
	}
	examples := []string{"audio-video", "clipbrd", "context-menu", "control-widget", "cookie", "cookie-manager", "custom-drag-window", "dev-tools",
		"dom-visitor", "download", "drag-file", "execute-dev-tools-method", "execute-javascript", "flash-for-cef-v87",
		"frameless", "html5test", "i18n", "internal-http-server", "ipc-on-emit/go-composite-type", "ipc-on-emit/go-to-go", "ipc-on-emit/go-to-js", "ipc-on-emit/js-to-go",
		"key-event", "lib-checkupdate", "load-html-url", "local-load", "mockevent",
		"msgbox", "popup-sub-window", "print-pdf", "process-type", "proxy", "response-filter",
		"scheme", "screen", "screenshot", "search-text", "simple", "simple-local-load", "helper-process", "helper-process/helper", "sys-dialog",
		"sys-menu", "tempdll", "view-source", "vue", "window/close-for-hidden", "window/close-for-minimize", "window/default-hidden-window", "window/new-window",
		"webkit-register", "window/close-for-hidden", "window/close-for-minimize", "window/default-hidden-window", "zoom"}
	var (
		ext     string
		ldflags string
		tags    = `-tags=tempdll`
	)

	if isWindows {
		examples = append(examples, "windows/custom-browser-create", "windows/notintaskbar", "windows/transparent", "osr/windows")
		examples = append(examples, "tray/lclceftray", "tray/lcltray", "tray/systray", "tray/lclvftray")
		ext = ".exe"
		ldflags = `-H windowsgui -s -w`
	} else if isLinux {
		examples = append(examples, "osr/linux")
		tags = `-tags=tempdll gtk3`
	}
	if isLinux || isDarwin {
		examples = append(examples, "tray/lcltray", "tray/systray")
	}
	cmd := command.NewCMD()
	for i, example := range examples {
		dir := filepath.Join(wd, "example", example)
		if isExist(dir) {
			cmd.Dir = dir
			copySyso(dir)
			out := filepath.Join(dist, example+ext)
			println("build example", example, fmt.Sprintf("%d/%d", i+1, len(examples)), "\n\tbuild-dir:", dir, "\n\tout-dir:", out)
			cmd.Command("go", "build", "-ldflags", ldflags, "-o", out, tags)
			removeSyso(dir)
			println()
		} else {
			println("error not found:", dir)
			println()
		}
	}
	cmd.Close()
	println("build end.")
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
