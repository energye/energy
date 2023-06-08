package main

import (
	"embed"
	"github.com/energye/energy/v2/cmd/internal/ide"
)

//go:embed resources
var resources embed.FS

func main() {
	ide.Run(&resources)
}
