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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"os"
	"sort"
	"strings"
)

var CmdVersion = &command.Command{
	UsageLine: "package -a [all]",
	Short:     "Get version list",
	Long: `
	-a show all details
	.  Execute default command
`,
}

func init() {
	CmdVersion.Run = runVersion
}

func runVersion(c *command.Config) error {
	downloadJSON, err := tools.HttpRequestGET(consts.DownloadVersionURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	var edv map[string]interface{}
	downloadJSON = bytes.TrimPrefix(downloadJSON, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(downloadJSON, &edv); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	if versionList, ok := edv["versionList"].(map[string]interface{}); ok {
		var keys []string
		for k, _ := range versionList {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		println("Latest:", edv["latest"].(string))
		println("Version list")
		for i := len(keys) - 1; i >= 0; i-- {
			var version = keys[i]
			var ver = versionList[version].(map[string]interface{})
			if c.Version.All {
				wrt := &bytes.Buffer{}
				wrt.WriteString("  ")
				wrt.WriteString(version)
				wrt.WriteString("\r\n")
				for key, value := range ver {
					key = strings.ToUpper(key)
					if key == "MODULES" {
						continue
					}
					wrt.WriteString(fmt.Sprintf("\t%s: %s", strings.ToUpper(key), value))
					wrt.WriteString("\r\n")
				}
				println(wrt.String())
			} else {
				println("  ", version)
			}
		}
	}
	return nil
}
