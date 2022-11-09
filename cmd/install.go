package cmd

import "fmt"

var CmdInstall = &Command{
	UsageLine: "install [path]",
	Short:     "Automatically configure the CEF and Energy framework",
	Long:      `Automatically configure the CEF and Energy framework, During this process, CEF and Energy are downloaded`,
}

func init() {
	CmdInstall.Run = runInstall
}

func runInstall(c *CommandConfig) error {
	if c.Install.Path == "" {
		c.Install.Path = c.Wd
	}
	fmt.Println("开始安装CEF和Energy依赖")
	return nil
}
