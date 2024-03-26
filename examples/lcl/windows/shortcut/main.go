//go:build windows
// +build windows

package main

import (
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/pkgs/win"
	"os"
)

func main() {
	inits.Init(nil, nil)
	rtl.CreateURLShortCut(win.GetDesktopPath(), "energy", "https://github.com/energye/energy")
	rtl.CreateShortCut(win.GetDesktopPath(), "shortcut", os.Args[0], "", "描述", "-b -c")
	lcl.ShowMessage("Hello!")
}
