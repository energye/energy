//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package gen

import (
	"errors"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
)

func GeneraICON(iconFilePath, outPath string) (string, error) {
	return "", errors.New("Genera ICON only applicable to Windows")
}

func GeneraSYSO(exeName, iconFilePath, manifestFilePath, outPath, arch string, info project.Info) (string, error) {
	return "", errors.New("Genera SYSO only applicable to Windows")
}
