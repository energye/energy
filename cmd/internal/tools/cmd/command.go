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
	"strings"
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
		fmt.Println(name, strings.Join(args, " "))
	}
	cmd := exec.Command(name, args...)
	if m.Dir != "" {
		cmd.Dir = m.Dir
	}
	//隐藏调用外部命令窗口
	if m.HideWindow && IsWindows() {
		cmd.SysProcAttr = HideWindow(true)
	}
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
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
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
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
