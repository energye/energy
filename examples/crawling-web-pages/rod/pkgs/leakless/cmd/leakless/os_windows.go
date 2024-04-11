//go:build windows
// +build windows

package main

import (
	"os"
	"os/exec"
	"strconv"
)

func osSetupCmd(cmd *exec.Cmd) error {
	return nil
}

func kill(p *os.Process) {
	_ = exec.Command("taskkill", "/t", "/f", "/pid", strconv.Itoa(p.Pid)).Run()
}
