package cmd

type CommandConfig struct {
	Install Install `command:"install"`
	Package Package `command:"package"`
}

type Install struct {
	Path string `short:"m" long:"path" description:"Installation directory Default current directory"`
}

type Package struct {
}
