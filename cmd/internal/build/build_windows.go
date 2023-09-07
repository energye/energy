//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package build

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/project"
	"path/filepath"
)

// 构建windows执行程序
//  exe生成图标
//  编译go
//  upx
func build(c *command.Config) error {
	// 读取项目配置文件 energy.json 在main函数目录
	// test
	var path = filepath.ToSlash("E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\demo")
	var app, err = project.NewProject(path)
	if err != nil {
		return err
	}
	fmt.Println(app)
	// 生成exe图标

	return nil
}
