package main

import (
	"fmt"
	"github.com/energye/energy/v2/examples/lcl/grids/drawgrid/form"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"math/rand"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()

	mainForm := lcl.Application.CreateForm()
	mainForm.SetWidth(700)
	mainForm.SetHeight(500)
	mainForm.WorkAreaCenter()
	mainForm.SetCaption("表格自绘")
	mainForm.ScaleSelf()
	grid := form.NewPlayControl(mainForm)
	grid.SetParent(mainForm)
	grid.SetAlign(types.AlClient)
	for i := 1; i <= 100; i++ {
		grid.Add(form.TPlayListItem{fmt.Sprintf("标题%d", i), "张三", 100000 + rand.Int31n(100000), "", ""})

	}
	lcl.Application.Run()
}
