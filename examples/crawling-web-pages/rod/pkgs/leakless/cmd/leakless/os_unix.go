//go:build !windows
// +build !windows

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func osSetupCmd(cmd *exec.Cmd) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return nil
}

func kill(p *os.Process) {
	_ = syscall.Kill(-p.Pid, syscall.SIGKILL)
}
