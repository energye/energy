package cmd

var CmdPackage = &Command{
	UsageLine: "package",
	Short:     "Making an Installation Package",
	Long:      `Making an Installation Package`,
}

func init() {
	CmdPackage.Run = runPackage
}

func runPackage(c *CommandConfig) error {
	return nil
}
