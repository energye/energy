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
	bufferPanel *cef.TBufferPanel
	chromium    cef.IChromium
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
	m.SetWidth(800)
	m.SetHeight(600)
	m.ScreenCenter()
	m.chromium = cef.NewChromium(m, nil)

	m.bufferPanel = cef.NewBufferPanel(m)
	m.bufferPanel.SetParent(m)
	m.bufferPanel.SetColor(colors.ClAqua)
	m.bufferPanel.SetWidth(800)
	m.bufferPanel.SetHeight(600)
	//bufferPanel.SetAlign(types.AlClient)
	m.bufferPanelEvent()
	m.chromiumEvent()
	m.SetOnShow(func(sender lcl.IObject) {
		b := m.chromium.Initialized()
		fmt.Println("init:", b)
		cb := m.chromium.CreateBrowser(nil, "", nil, nil)
		fmt.Println("CreateBrowser:", cb)
		m.chromium.Options().SetBackgroundColor(cef.CefColorSetARGB(0, 0, 0xff, 0xff))
		m.bufferPanel.CreateIMEHandler()
		m.chromium.InitializeDragAndDrop(m.bufferPanel)
		m.chromium.LoadUrl("https://www.baidu.com")
	})

}
func (m *WindowDemo) chromiumEvent() {

	var (
		popUpBitmap                  *lcl.TBitmap
		tempBitMap                   *lcl.TBitmap
		tempWidth, tempHeight        int32
		tempForcedResize             bool
		tempLineSize                 int
		tempSrcOffset, tempDstOffset int
		src, dst                     uintptr
	)
	m.chromium.SetOnCursorChange(func(sender lcl.IObject, browser *cef.ICefBrowser, cursor consts.TCefCursorHandle, cursorType consts.TCefCursorType, customCursorInfo *cef.TCefCursorInfo) bool {
		fmt.Println("SetOnCursorChange")
		m.bufferPanel.SetCursor(cef.CefCursorToWindowsCursor(cursorType))
		return true
	})
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("SetOnAfterCreated")
		m.chromium.LoadUrl("https://www.baidu.com")
	})
	m.chromium.SetOnPopupShow(func(sender lcl.IObject, browser *cef.ICefBrowser, show bool) {
		fmt.Println("PopupShow - show:", show)
		if m.chromium != nil {
			m.chromium.Invalidate(consts.PET_VIEW)
		}
	})
	m.chromium.SetOnPopupSize(func(sender lcl.IObject, browser *cef.ICefBrowser, rect *cef.TCefRect) {
		screenScale := m.bufferPanel.ScreenScale()
		fmt.Println("PopupSize - rect:", rect, "screenScale:", screenScale)
		cef.LogicalToDeviceRect(rect, float64(screenScale))
		fmt.Println("PopupSize - rect:", rect, "screenScale:", screenScale)
	})

	// 在Paint内展示内容到窗口中
	m.chromium.SetOnPaint(func(sender lcl.IObject, browser *cef.ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *cef.TCefRectArray, buffer uintptr, width, height int32) {
		if m.bufferPanel.BeginBufferDraw() {
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
				tempForcedResize = m.bufferPanel.UpdateBufferDimensions(width, height) || m.bufferPanel.BufferIsResized(false)
				tempBitMap = m.bufferPanel.Buffer()
				tempBitMap.BeginUpdate(false)
				tempWidth = m.bufferPanel.BufferWidth()
				tempHeight = m.bufferPanel.BufferHeight()
			}
			fmt.Println("tempWidth:", tempWidth, "tempHeight:", tempHeight)
			//byteBufPtr := (*byte)(unsafe.Pointer(buffer))
			rgbSizeOf := int(unsafe.Sizeof(cef.TRGBQuad{}))
			fmt.Println("SizeOf(TRGBQuad):", rgbSizeOf)
			srcStride := int(width) * rgbSizeOf
			fmt.Println("srcStride:", srcStride)
			for i := 0; i < dirtyRects.Count(); i++ {
				rect := dirtyRects.Get(i)
				if rect.X >= 0 && rect.Y >= 0 {
					tempLineSize = int(math.Min(float64(rect.Width), float64(tempWidth-rect.X))) * rgbSizeOf
					fmt.Println("tempLineSize:tempLineSize", tempLineSize)
					if tempLineSize > 0 {
						tempSrcOffset = int((rect.Y*width)+rect.X) * rgbSizeOf
						tempDstOffset = int(rect.X) * rgbSizeOf
						//src := @pbyte(buffer)[TempSrcOffset];
						src = uintptr(common.GetParamPtr(buffer, tempSrcOffset)) // 拿到src指针, 实际是 byte 指针
						fmt.Println("src-dst-offset:", tempSrcOffset, tempDstOffset, src)
						j := int(math.Min(float64(rect.Height), float64(tempHeight-rect.Y)))
						//fmt.Println("j:", j)
						for ii := 0; ii < j; ii++ {
							tempBufferBits := tempBitMap.ScanLine(rect.Y + int32(ii))
							dst = uintptr(common.GetParamPtr(tempBufferBits, tempDstOffset)) //拿到dst指针, 实际是 byte 指针
							//fmt.Println("dst:", dst)
							rtl.Move(src, dst, tempLineSize)
							src = src + uintptr(srcStride)
						}
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

			m.bufferPanel.EndBufferDraw()
			ha := m.HandleAllocated()
			fmt.Println("ha:", ha)
			if ha {
				m.bufferPanel.Invalidate()
			}
		}
		fmt.Println("SetOnPaint", browser.Identifier(), kind, dirtyRects.Count(), dirtyRects.Get(0), buffer, width, height)
		fmt.Println(tempWidth, tempHeight, tempForcedResize)
	})
}
func getModifiers(shift types.TShiftState) consts.TCefEventFlags {
	var result = consts.EVENTFLAG_NONE
	if shift.In(types.SsShift) {
		result = result | consts.EVENTFLAG_SHIFT_DOWN
	} else if shift.In(types.SsAlt) {
		result = result | consts.EVENTFLAG_ALT_DOWN
	} else if shift.In(types.SsCtrl) {
		result = result | consts.EVENTFLAG_CONTROL_DOWN
	} else if shift.In(types.SsLeft) {
		result = result | consts.EVENTFLAG_LEFT_MOUSE_BUTTON
	} else if shift.In(types.SsRight) {
		result = result | consts.EVENTFLAG_RIGHT_MOUSE_BUTTON
	} else if shift.In(types.SsMiddle) {
		result = result | consts.EVENTFLAG_MIDDLE_MOUSE_BUTTON
	}
	return result
}
func GetButton(Button types.TMouseButton) (result consts.TCefMouseButtonType) {
	switch Button {
	case types.MbRight:
		result = consts.MBT_RIGHT
	case types.MbMiddle:
		result = consts.MBT_MIDDLE
	default:
		result = consts.MBT_LEFT
	}
	return
}

func (m *WindowDemo) bufferPanelEvent() {
	m.bufferPanel.SetOnClick(func(sender lcl.IObject) {
		fmt.Println("SetOnClick")
	})
	m.bufferPanel.SetOnEnter(func(sender lcl.IObject) {
		fmt.Println("SetOnEnter")
		m.chromium.SetFocus(true)
	})
	m.bufferPanel.SetOnExit(func(sender lcl.IObject) {
		fmt.Println("SetOnExit")
		m.chromium.SetFocus(false)
	})
	m.bufferPanel.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		fmt.Println("SetOnMouseMove", shift, x, y)
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseMoveEvent(mouseEvent, false)
	})
}
