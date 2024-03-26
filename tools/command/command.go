//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package command

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

type CMD struct {
	HideWindow      bool
	IsPrint         bool
	Dir             string
	MessageCallback func([]byte, error)
	stdout          io.ReadCloser
}

func NewCMD() *CMD {
	return &CMD{IsPrint: true}
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func (m *CMD) Close() {
	if m.stdout != nil {
		m.stdout.Close()
	}
}

func (m *CMD) Command(name string, args ...string) {
	if m.IsPrint {
		fmt.Println("command name:", name, "args:", args)
	}
	cmd := exec.Command(name, args...)
	if m.Dir != "" {
		cmd.Dir = m.Dir
	}
	//隐藏调用外部命令窗口
	if m.HideWindow && IsWindows() {
		cmd.SysProcAttr = HideWindow(true)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		if m.MessageCallback != nil {
			m.MessageCallback(nil, err)
		}
		return
	}
	m.stdout = stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		if m.MessageCallback != nil {
			m.MessageCallback(nil, err)
		}
		return
	}
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			byt, b, err := reader.ReadLine()
			if err != nil || io.EOF == err {
				if m.MessageCallback != nil {
					m.MessageCallback([]byte("exit"), nil)
				}
				break
			}
			if m.MessageCallback != nil {
				m.MessageCallback(byt, nil)
			} else {
				if m.IsPrint {
					fmt.Println("line:", string(byt), b)
				}
			}
		}
		stdout.Close()
	}()
	err = cmd.Wait()
	if err != nil {
		if m.MessageCallback != nil {
			m.MessageCallback(nil, errors.New("wait end error "+err.Error()))
		}
	} else {
		if m.MessageCallback != nil {
			m.MessageCallback(nil, nil)
		}
	}
}
