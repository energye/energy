package env

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/golcl/energy/homedir"
	toolsCommand "github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func SetEnergyHomeEnv(key, value string) {
	println("\nSetting environment Variables [ENERGY_HOME] to", value)
	cmd := toolsCommand.NewCMD()
	cmd.MessageCallback = func(s []byte, e error) {
		fmt.Println("CMD:", s, " error:", e)
	}
	defer cmd.Close()
	if command.IsWindows {
		var args = []string{"/c", "setx", key, value}
		cmd.Command("cmd.exe", args...)
	} else {
		var envFiles []string
		var energyHomeKey = fmt.Sprintf("export %s", key)
		var energyHome = fmt.Sprintf("export %s=%s", key, value)
		if command.IsLinux {
			envFiles = []string{".profile", ".zshrc", ".bashrc"}
		} else if command.IsDarwin {
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
