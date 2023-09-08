//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package assets

import (
	"embed"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed assets
var assets embed.FS

// AssetsPath 返回配置资源目录
func AssetsPath(projectData *project.Project, file string) string {
	return filepath.ToSlash(filepath.Join(projectData.AssetsDir, file))
}

// BuildOutPath 返回固定构建输出目录 $current/build
func BuildOutPath(projectData *project.Project) string {
	return filepath.Join(projectData.ProjectPath, "build")
}

// ReadFile
//  读取文件，根据项目配置先在本地目录读取，如果读取失败，则在内置资源目录读取
func ReadFile(projectData *project.Project, assetsFSPath, file string) ([]byte, error) {
	var (
		content []byte
		err     error
	)
	if projectData != nil {
		localFilePath := AssetsPath(projectData, file)
		content, err = os.ReadFile(localFilePath)
	}
	if errors.Is(err, fs.ErrNotExist) || content == nil {
		content, err = assets.ReadFile(assetsFSPath + file)
		if err != nil {
			return nil, err
		}
		return content, nil
	}

	return content, err
}

// WriteFile 写文件到本地目录
func WriteFile(projectData *project.Project, file string, content []byte) error {
	buildOutDir := BuildOutPath(projectData)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	targetPath := filepath.Join(buildOutDir, file)
	if !projectData.Clean {
		if tools.IsExist(targetPath) {
			return nil
		}
	}
	if err := os.WriteFile(targetPath, content, 0644); err != nil {
		return err
	}
	return nil
}
