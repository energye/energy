package main

import (
	"github.com/energye/golcl/tools/command"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	const (
		isWindows = runtime.GOOS == "windows" //support
		isLinux   = runtime.GOOS == "linux"   //support
		isDarwin  = runtime.GOOS == "darwin"  //support
	)
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
	for _, example := range examples {
		dir := filepath.Join(wd, "example", example)
		if isExist(dir) {
			cmd.Dir = dir
			out := filepath.Join(dist, example+ext)
			println("build example", example, "\n\tbuild-dir:", dir, "\n\tout-dir:", out)
			cmd.Command("go", "build", "-ldflags", ldflags, "-o", out, `-tags=tempdll`)
			println()
		} else {
			println("not found:", dir)
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
