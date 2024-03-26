package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
)

type TMainForm struct {
	*lcl.TForm
}

var mainForm *TMainForm

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(object lcl.IObject) {
	f.SetCaption("drop files")
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()
	f.EnabledMaximize(false)

	// allow drop file
	f.SetAllowDropFiles(true)

	// windows10没生效，有待研究
	windowsUACMessageFilter(f.Handle())
}

func (f *TMainForm) OnFormDropFiles(sender lcl.IObject, aFileNames []string) {
	fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
	for i, s := range aFileNames {
		fmt.Println("index:", i, ", filename:", s)
	}
}
