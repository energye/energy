package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var CmdVersion = &Command{
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

func runVersion(c *CommandConfig) error {
	downloadJSON, err := downloadConfig(download_version_config_url)
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
		for version, fver := range versionList {
			var ver = fver.(map[string]interface{})
			if c.Version.All {
				println(" ", version, fmt.Sprintf(`
    CEF: %s
    ENERGY: %s`, ver["cef"].(string), ver["energy"].(string)))
			} else {
				println(" ", version)
			}
		}
	}
	return nil
}
