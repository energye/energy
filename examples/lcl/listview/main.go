package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/types"
	"math/rand"
	"runtime"
	"strings"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.SetOnException(func(sender lcl.IObject, e lcl.IException) {

	})
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(500)
	mainForm.SetHeight(600)
	// 双缓冲
	mainForm.SetDoubleBuffered(true)

	imgList := lcl.NewImageList(mainForm)
	if runtime.GOOS == "windows" {
		icon := lcl.NewIcon()
		icon.LoadFromResourceName(rtl.MainInstance(), "MAINICON")
		imgList.AddIcon(icon)
		icon.Free()
	}

	lv1 := lcl.NewListView(mainForm)
	lv1.SetParent(mainForm)
	lv1.SetAlign(types.AlTop)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetViewStyle(types.VsReport)
	lv1.SetGridLines(true)
	//lv1.SetColumnClick(false)
	lv1.SetHideSelection(false)

	col := lcl.AsListColumn(lv1.Columns().Add())
	col.SetCaption("序号")
	col.SetWidth(100)

	col = lcl.AsListColumn(lv1.Columns().Add())
	col.SetCaption("项目1")
	col.SetWidth(200)
	lv1.SetOnClick(func(lcl.IObject) {
		if lv1.ItemIndex() != -1 {
			item := lv1.Selected() // lv1.Items().Item(lv1.ItemIndex())
			fmt.Println(item.Caption(), ", ", item.SubItems().Strings(0))
		}
	})
	// 双击删除选中项
	lv1.SetOnDblClick(func(sender lcl.IObject) {
		if lv1.ItemIndex() != -1 {
			lv1.Items().Delete(lv1.ItemIndex())
		}
	})

	// 排序箭头
	lv1.SetAutoSortIndicator(true)
	//lv1.SetSortDirection(types.SdAscending) //  AES or DES
	lv1.SetSortType(types.StText) // 按文本排序

	lv1.SetOnColumnClick(func(sender lcl.IObject, column lcl.IListColumn) {
		fmt.Println("index:", column.Index())
		// 按柱头索引排序, lcl兼容版第二个参数永远为 column
		lv1.CustomSort(func(item1, item2 lcl.IListItem, optionalParam uint32) int32 {
			if optionalParam == 0 {
				return int32(strings.Compare(item1.Caption(), item2.Caption()))
			} else {
				return int32(strings.Compare(item1.SubItems().Strings(int32(optionalParam-1)), item2.SubItems().Strings(int32(optionalParam-1))))
			}
		}, uint32(column.Index()))
	})

	//lv1.SetOnCompare(func(sender lcl.IObject, item1, item2 lcl.IListItem, data int32, compare *int32) {
	//	fmt.Println("SetOnCompare", data)
	//	if data == 0 {
	//		*compare = int32(strings.Compare(item1.Caption(), item2.Caption()))
	//	} else {
	//		*compare = int32(strings.Compare(item1.SubItems().Strings(data-1), item2.SubItems().Strings(data-1)))
	//	}
	//})

	//	lv1.Clear()
	lv1.Items().BeginUpdate()
	for i := 1; i <= 1000; i++ {
		item := lv1.Items().Add()
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i+rand.Int()))
		item.SubItems().Add(fmt.Sprintf("值：%d", i+rand.Int()))
	}
	lv1.Items().EndUpdate()

	// icon样式

	lv2 := lcl.NewListView(mainForm)
	lv2.SetParent(mainForm)
	lv2.SetAlign(types.AlTop)
	//lv2.SetRowSelect(true)
	//lv2.SetReadOnly(true)
	lv2.SetViewStyle(types.VsIcon)
	//lv2.SetSmallImages(imgList)
	lv2.SetLargeImages(imgList)
	// 因为这时候还没有计算altop的，所以要手动设置宽度，这样添加的值才会根据当前宽度排列
	lv2.SetWidth(mainForm.Width())
	//lv2.SetStateImages(imgList)

	lv2.SetOnClick(func(lcl.IObject) {
		if lv2.ItemIndex() != -1 {
			item := lv2.Selected()
			fmt.Println(item.Caption())
		}
	})
	// 双击删除选中项
	lv2.SetOnDblClick(func(sender lcl.IObject) {
		if lv2.ItemIndex() != -1 {
			lv2.Items().Delete(lv2.ItemIndex())
		}
	})
	lv2.Items().BeginUpdate()
	for i := 1; i <= 10; i++ {
		item := lv2.Items().Add()
		item.SetImageIndex(0)
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
	}
	lv2.Items().EndUpdate()

	// lv3
	lv3 := lcl.NewListView(mainForm)
	lv3.SetParent(mainForm)
	lv3.SetAlign(types.AlClient)
	lv3.SetRowSelect(true)
	lv3.SetReadOnly(true)
	lv3.SetViewStyle(types.VsReport)
	lv3.SetGridLines(true)
	// 失去焦点不隐藏选择的
	lv3.SetHideSelection(false)

	col = lcl.AsListColumn(lv3.Columns().Add())
	col.SetCaption("序号")
	col.SetWidth(100)
	col = lcl.AsListColumn(lv3.Columns().Add())
	col.SetCaption("Sub1")
	col.SetWidth(100)

	lv3.SetOnClick(func(lcl.IObject) {
		if lv3.ItemIndex() != -1 {
			item := lv3.Selected()
			fmt.Println(item.Caption(), ", ", item.SubItems().Strings(0))
		}
	})
	lv3.Items().BeginUpdate()
	for i := 1; i <= 2; i++ {
		item := lv3.Items().Add()
		item.SetImageIndex(0)

		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("值：%d", i))
	}
	for i := 1; i <= 2; i++ {
		item := lv3.Items().Add()
		item.SetImageIndex(0)

		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("值：%d", i))
	}
	lv3.Items().EndUpdate()

	pnlbottom := lcl.NewPanel(mainForm)
	pnlbottom.SetParent(mainForm)
	pnlbottom.SetAlign(types.AlBottom)
	btnTest := lcl.NewButton(mainForm)
	btnTest.SetParent(pnlbottom)
	btnTest.SetCaption("SetSelected")
	btnTest.SetWidth(120)
	btnTest.SetTop(10)
	btnTest.SetLeft(10)
	btnTest.SetOnClick(func(sender lcl.IObject) {

		if lv1.Items().Count() > 5 {
			fmt.Println("click select")
			item := lv1.Items().Item(3) // 第四个
			lv1.SetSelected(item)
		}
	})

	btnTest2 := lcl.NewButton(mainForm)
	btnTest2.SetParent(pnlbottom)
	btnTest2.SetTop(10)
	btnTest2.SetLeft(btnTest.Left() + btnTest.Width() + 10)
	btnTest2.SetWidth(120)
	btnTest2.SetCaption("DeleteSelected")
	btnTest2.SetOnClick(func(sender lcl.IObject) {
		if lv1.SelCount() > 0 {
			fmt.Println("click delete")
			lv1.DeleteSelected()
		}
	})

	btnTest3 := lcl.NewButton(mainForm)
	btnTest3.SetParent(pnlbottom)
	btnTest3.SetTop(10)
	btnTest3.SetLeft(btnTest2.Left() + btnTest2.Width() + 10)
	btnTest3.SetWidth(120)
	btnTest3.SetCaption("Add Item")
	btnTest3.SetOnClick(func(sender lcl.IObject) {
		item := lv3.Items().Add()
		item.SetCaption("111")
		item.SubItems().Add("sub")

		// 总是显示
		item.MakeVisible(true)

		// 如果要选中的话
		item.SetSelected(true)

	})

	lcl.Application.Run()
}
