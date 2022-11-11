package cmd

type CommandConfig struct {
	Index   int
	Wd      string
	Install Install `command:"install"`
	Package Package `command:"package"`
}

type Install struct {
	Path    string `short:"p" long:"path" description:"Installation directory Default current directory"`
	Version string `short:"v" long:"version" description:"Specifying a version number"`
	Name    string `short:"n" long:"name" description:"Name of the frame after installation" default:"EnergyFramework"`
}

type Package struct {
}

type Command struct {
	Run                    func(c *CommandConfig) error
	UsageLine, Short, Long string
}
