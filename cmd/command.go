package cmd

type CommandConfig struct {
	Index   int
	Wd      string
	Install Install `command:"install"`
	Package Package `command:"package"`
}

type Install struct {
	Path string `short:"m" long:"path" description:"Installation directory Default current directory"`
}

type Package struct {
}

type Command struct {
	Run                    func(c *CommandConfig) error
	UsageLine, Short, Long string
}
