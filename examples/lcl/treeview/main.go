package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/types"
	"runtime"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.SetIconResId(3)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(500)

	imglist := lcl.NewImageList(mainForm)

	icon := lcl.NewIcon()
	if runtime.GOOS == "windows" {
		icon.LoadFromResourceName(rtl.MainInstance(), "MAINICON")
	} else {
		icon.LoadFromFile("brown.ico")
	}
	imglist.AddIcon(icon)
	icon.Free()

	ico2 := lcl.NewIcon()
	ico2.LoadFromFile("brown.ico")
	imglist.AddIcon(ico2)
	ico2.Free()

	// -----------TreeView 不同Node弹出不同菜单，两个右键例程不同使用

	tvpm1 := lcl.NewPopupMenu(mainForm)
	mItem := lcl.NewMenuItem(mainForm)
	mItem.SetCaption("第一种")
	tvpm1.Items().Add(mItem)

	tvpm2 := lcl.NewPopupMenu(mainForm)
	mItem = lcl.NewMenuItem(mainForm)
	mItem.SetCaption("第二种")
	tvpm2.Items().Add(mItem)

	tv := lcl.NewTreeView(mainForm)
	tv.SetParent(mainForm)
	tv.SetAlign(types.AlClient)

	tv.SetImages(imglist)
	tv.SetStateImages(imglist)

	// 自动展开
	//tv.SetAutoExpand(true)

	tv.SetOnClick(func(lcl.IObject) {
		node := tv.Selected()
		if node != nil /*&& node.IsValid()*/ {
			fmt.Println("Text:", node.Text(), ", Level:", node.Level(), ", Index:", node.Index(), ", hasChild:", node.HasChildren())
		}
	})
	// 双击删除
	tv.SetOnDblClick(func(sender lcl.IObject) {
		sel := tv.Selected()
		if sel != nil /*&& sel.IsValid()*/ {
			sel.Delete()
		}
		// 或者
		//tv.Items().Delete(sel)
	})
	tv.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbRight {
			node := tv.GetNodeAt(x, y)
			if node != nil /*&& node.IsValid()*/ {
				// 自由决择是否选中
				node.SetSelected(true)
				// 根据Level来判断，这里只是做演示
				p := lcl.Mouse.CursorPos()
				switch node.Level() {
				case 0:
					tvpm1.PopUp1(p.X, p.Y)
				case 1:
					tvpm2.PopUp1(p.X, p.Y)
				}
				fmt.Println("node.Level():", node.Level(), ", text:", node.Text())
			}
		}
	})

	//	tv.Items().Clear()
	// 第一个节点
	node := tv.Items().AddChild(nil, "首个")

	// 批量添加最好使用BeginUpdate与EndUpdate组合
	tv.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		subnode := tv.Items().AddChild(node, fmt.Sprintf("Node%d", i))
		// 设置相关imagelist中的图标索引
		// 默认图标索引
		subnode.SetImageIndex(1)
		// 鼠标选中时索引
		subnode.SetSelectedIndex(1)
		// 节点展开时索引
		//subnode.SetExpandedImageIndex(1)
		// 状态图标索引
		//subnode.SetStateIndex(1)

	}
	tv.Items().EndUpdate()
	// 展开
	node.Expand(true)

	// 第二个节点
	node = tv.Items().AddChild(nil, "第二个节点")
	// 批量添加最好使用BeginUpdate与EndUpdate组合
	tv.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		tv.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	tv.Items().EndUpdate()

	lcl.Application.Run()
}
