package main

import (
	"fmt"
	"github.com/energye/energy/v2/api"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)

type TMainForm struct {
	lcl.IForm
}

var (
	mainForm *TMainForm
)

func main() {
	lcl.DEBUG = true
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm)
}

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	fmt.Println("TMainForm OnFormCreate")
	m.SetCaption("LazVirtualDrawTree")
	m.WorkAreaCenter()
	m.SetWidth(600)
	m.SetHeight(600)

	// CreateBitmapFromResourceName()
	// CreateCheckImageList resname: LAZ_VT_CHECK_DARK
	// https://forum.lazarus.freepascal.org/index.php?board=31.0

	vDrawTree := lcl.NewLazVirtualDrawTree(m)
	vDrawTree.SetParent(m)
	vDrawTree.SetTop(50)
	vDrawTree.SetWidth(m.Width())
	vDrawTree.SetHeight(200)
	vDrawTree.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	dRoot := vDrawTree.AddChild(nil, api.PascalStr("Root"))
	child1 := vDrawTree.AddChild(dRoot, api.PascalStr("Child 1"))
	vDrawTree.AddChild(dRoot, api.PascalStr("Child 2"))
	vDrawTree.AddChild(dRoot, api.PascalStr("Child 3"))
	vDrawTree.AddChild(child1, api.PascalStr("Grandchild of Child 1"))

	vStrTree := lcl.NewLazVirtualStringTree(m)
	vStrTree.SetParent(m)
	vStrTree.SetTop(300)
	vStrTree.SetWidth(m.Width())
	vStrTree.SetHeight(200)
	vStrTree.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	vStrTree.SetDefaultText("测试")

	vStrTree.SetOnGetNodeDataSize(func(sender lcl.IBaseVirtualTree, nodeDataSize *int32) {
		fmt.Println("OnGetNodeDataSize", *nodeDataSize)
	})
	vStrTree.SetOnResize(func(sender lcl.IObject) {
		//fmt.Println("SetOnResize", vStrTree.Width(), vStrTree.Height())
	})
	sRoot := vStrTree.AddChild(nil, api.PascalStr("Root"))
	vStrTree.SetText(sRoot, 0, "Root")
	sChild1 := vStrTree.AddChild(sRoot, api.PascalStr("Child 1"))
	vStrTree.AddChild(sRoot, api.PascalStr("Child 2"))
	vStrTree.AddChild(sRoot, api.PascalStr("Child 3"))
	vStrTree.AddChild(sChild1, api.PascalStr("Grandchild of Child 1"))

	vsHeader := vStrTree.Header()
	vsHeader.SetHeight(19)
	vsHeader.SetDefaultHeight(19)
	columns := vsHeader.Columns()
	vSItem1 := lcl.AsVirtualTreeColumn(columns.Add())

	//vSItem1 := lcl.NewVirtualTreeColumn(columns)
	vSItem1.SetText("item?")
	vSItem1.SetWidth(50)
	//columns.SetItems(0, vSItem1)

	btn := lcl.NewButton(m)
	btn.SetParent(m)
	btn.SetCaption("点击一下？")
	btn.SetOnClick(func(sender lcl.IObject) {
		count := vDrawTree.RootNodeCount()
		fmt.Println("count:", count, dRoot.Index(), api.GoStr(dRoot.Data()))
		fmt.Println(vDrawTree.TotalCount())
		fmt.Println(vStrTree.TreeOptions().AutoOptions())
	})
}
