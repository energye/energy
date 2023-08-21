//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import (
	"errors"
	"fmt"
	"github.com/energye/golcl/energy/homedir"
	"github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var CmdSetenv = &Command{
	UsageLine: "setenv -p [path]",
	Short:     "Set energy framework development environment",
	Long: `
	-p Set the Framework pointed to by the ENERGY_HOME development environment variable
	.  Execute default command
`,
}

func init() {
	CmdSetenv.Run = runSetenv
}

func runSetenv(c *CommandConfig) error {
	if c.Setenv.Path == "" {
		return errors.New("ERROR: ENERGY environment variable, command line argument [-p] directory to empty")
	}
	if !IsExist(c.Setenv.Path) {
		return errors.New("Directory [" + c.Setenv.Path + "] does not exist")
	}
	setEnergyHomeEnv(EnergyHomeKey, c.Setenv.Path)
	println("SUCCESS")
	return nil
}

func setEnergyHomeEnv(key, value string) {
	println("\nSetting environment Variables [ENERGY_HOME] to", value)
	cmd := command.NewCMD()
	cmd.MessageCallback = func(s []byte, e error) {
		fmt.Println("CMD:", s, " error:", e)
	}
	defer cmd.Close()
	if isWindows {
		var args = []string{"/c", "setx", key, value}
		cmd.Command("cmd.exe", args...)
	} else {
		var envFiles []string
		var energyHomeKey = fmt.Sprintf("export %s", key)
		var energyHome = fmt.Sprintf("export %s=%s", key, value)
		if isLinux {
			envFiles = []string{".profile", ".zshrc", ".bashrc"}
		} else if isDarwin {
			envFiles = []string{".profile", ".zshrc", ".bash_profile"}
		}
		homeDir, _ := homedir.Dir()
		for _, file := range envFiles {
			var fp = path.Join(homeDir, file)
			cmd.Command("touch", fp)
			f, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND, 0666)
			if err == nil {
				var oldContent string
				if contentBytes, err := ioutil.ReadAll(f); err == nil {
					content := string(contentBytes)
					oldContent = content
					var lines = strings.Split(content, "\n")
					var exist = false
					for i := 0; i < len(lines); i++ {
						line := lines[i]
						if strings.Index(line, energyHomeKey) == 0 {
							content = strings.ReplaceAll(content, line, energyHome)
							exist = true
						}
					}
					if exist {
						if err := f.Close(); err == nil {
							var oldWrite = func() {
								if f, err = os.OpenFile(fp, os.O_RDWR, 0666); err == nil {
									f.WriteString(oldContent)
									f.Close()
								}
							}
							if newOpenFile, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0666); err == nil {
								if _, err := newOpenFile.WriteString(content); err == nil {
									newOpenFile.Close()
								} else {
									newOpenFile.Close()
									oldWrite()
								}
							} else {
								oldWrite()
							}
						}
					} else {
						f.WriteString("\n")
						f.WriteString(energyHome)
						f.WriteString("\n")
					}
				} else {
					f.Close()
				}
			}
		}
	}
}
