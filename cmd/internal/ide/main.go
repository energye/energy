package main

import (
	"embed"
	"github.com/energye/energy/v2/cmd/internal/ide/internal"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	inits.Init(nil, &resources)
	lcl.RunApp(&internal.Ide)
}
