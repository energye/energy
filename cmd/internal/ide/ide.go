package ide

import (
	"embed"
	"github.com/energye/energy/v2/cmd/internal/ide/internal"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
)

func Run(resources *embed.FS) {
	inits.Init(nil, resources)
	lcl.RunApp(&internal.Ide)
}
