package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"math"
	"unsafe"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, nil)
	var window = &WindowDemo{}
	//创建应用
	cefApp := cef.NewApplication(true)
	//
	cefApp.SetWindowlessRenderingEnabled(true)
	cefApp.SetExternalMessagePump(true)
	cefApp.SetMultiThreadedMessageLoop(false)
	// work
	cef.GlobalWorkSchedulerCreate(nil)
	cefApp.SetOnScheduleMessagePumpWork(nil)

	//指定一个URL地址，或本地html文件目录
	cefApp.StartMainProcess()
	lcl.RunApp(&window)
}

type WindowDemo struct {
	*lcl.TForm
}

func (m *WindowDemo) OnFormCreate(sender lcl.IObject) {

	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	fmt.Println("OnFormCreate")
	m.SetWidth(600)
	m.SetHeight(400)
	m.ScreenCenter()
	chromium := cef.NewChromium(m, nil)

	bufferPanel := cef.NewBufferPanel(m)
	bufferPanel.SetParent(m)
	bufferPanel.SetColor(colors.ClAqua)
	bufferPanel.SetWidth(800)
	bufferPanel.SetHeight(600)
	bufferPanel.SetAlign(types.AlClient)
	bufferPanel.SetOnClick(func(sender lcl.IObject) {
		fmt.Println("SetOnClick")
	})

	m.SetOnShow(func(sender lcl.IObject) {
		chromium.Initialized()
		cb := chromium.CreateBrowser(nil, "", nil, nil)
		fmt.Println("CreateBrowser:", cb)
		chromium.InitializeDragAndDrop(bufferPanel)
		chromium.LoadUrl("https://www.baidu.com")
	})

	var (
		popUpBitmap                  *lcl.TBitmap
		tempBitMap                   *lcl.TBitmap
		tempWidth, tempHeight        int32
		tempForcedResize             bool
		tempLineSize                 int
		tempSrcOffset, tempDstOffset int
		src, dst                     uintptr
	)
	chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("SetOnAfterCreated")
		chromium.LoadUrl("https://www.baidu.com")
	})
	chromium.SetOnPopupShow(func(sender lcl.IObject, browser *cef.ICefBrowser, show bool) {
		fmt.Println("PopupShow - show:", show)
		if chromium != nil {
			chromium.Invalidate(consts.PET_VIEW)
		}
	})
	chromium.SetOnPopupSize(func(sender lcl.IObject, browser *cef.ICefBrowser, rect *cef.TCefRect) {
		screenScale := bufferPanel.ScreenScale()
		fmt.Println("PopupSize - rect:", rect, "screenScale:", screenScale)
		cef.LogicalToDeviceRect(rect, float64(screenScale))
		fmt.Println("PopupSize - rect:", rect, "screenScale:", screenScale)
	})
	chromium.SetOnPaint(func(sender lcl.IObject, browser *cef.ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *cef.TCefRectArray, buffer uintptr, width, height int32) {
		if bufferPanel.BeginBufferDraw() {
			if kind == consts.PET_POPUP {
				if popUpBitmap == nil || popUpBitmap.Width() != width || popUpBitmap.Height() != height {
					if popUpBitmap != nil {
						popUpBitmap.Free()
					}
					popUpBitmap = lcl.NewBitmap()
					popUpBitmap.SetPixelFormat(types.Pf32bit)
					popUpBitmap.SetHandleType(types.BmDIB)
					popUpBitmap.SetWidth(width)
					popUpBitmap.SetHeight(height)
				}
				tempBitMap = popUpBitmap
				tempBitMap.BeginUpdate(false)
				tempWidth, tempHeight = popUpBitmap.Width(), popUpBitmap.Height()
			} else {
				tempForcedResize = bufferPanel.UpdateBufferDimensions(width, height) || bufferPanel.BufferIsResized(false)
				tempBitMap = bufferPanel.Buffer()
				tempBitMap.BeginUpdate(false)
				tempWidth = bufferPanel.BufferWidth()
				tempHeight = bufferPanel.BufferHeight()
			}
			//byteBufPtr := (*byte)(unsafe.Pointer(buffer))
			rgbSizeOf := int(unsafe.Sizeof(cef.TRGBQuad{}))
			srcStride := int(width) + rgbSizeOf
			fmt.Println("srcStride:", srcStride)
			for i := 0; i < dirtyRects.Count(); i++ {
				rect := dirtyRects.Get(i)
				tempLineSize = int(math.Min(float64(rect.Width), float64(tempWidth-rect.X))) * rgbSizeOf
				fmt.Println("tempLineSize:tempLineSize", tempLineSize)
				if tempLineSize > 0 {
					tempSrcOffset = int((rect.Y*width)+rect.X) * rgbSizeOf
					tempDstOffset = int(rect.X) * rgbSizeOf
					//src := @pbyte(buffer)[TempSrcOffset];
					src = uintptr(common.GetParamPtr(buffer, tempSrcOffset)) // 拿到src指针
					fmt.Println("src-dst-offset:", tempSrcOffset, tempDstOffset, src)
					j := int(math.Min(float64(rect.Height), float64(tempHeight-rect.Y)))
					fmt.Println("j:", j)
					for ii := 0; ii < j; ii++ {
						tempBufferBits := tempBitMap.ScanLine(rect.Y + int32(i))
						dst = uintptr(common.GetParamPtr(tempBufferBits, tempDstOffset)) //拿到dst指针
						fmt.Println("dst:", dst)
						rtl.Move(src, dst, tempLineSize)
						src = src + uintptr(srcStride)
					}
				}
			}
			tempBitMap.EndUpdate(false)
			//if FShowPopup and (FPopUpBitmap <> nil) then
			//begin
			//TempSrcRect := Rect(0, 0, min(FPopUpRect.Right - FPopUpRect.Left, FPopUpBitmap.Width), min(FPopUpRect.Bottom - FPopUpRect.Top, FPopUpBitmap.Height));
			//
			//Panel1.BufferDraw(FPopUpBitmap, TempSrcRect, FPopUpRect);
			//end;
			if popUpBitmap != nil {

			}

			bufferPanel.EndBufferDraw()
			ha := m.HandleAllocated()
			fmt.Println("ha:", ha)
			if ha {
				cef.QueueAsyncCall(func(id int) {
					bufferPanel.Invalidate()
				})
			}
		}
		fmt.Println("SetOnPaint", browser.Identifier(), kind, dirtyRects.Count(), dirtyRects.Get(0), buffer, width, height)
		fmt.Println(tempWidth, tempHeight, tempForcedResize)
	})
}
