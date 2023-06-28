package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/assetserve"
	t "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"math"
	"time"
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
	delayed := cef.GlobalWorkSchedulerCreate(nil)
	cefApp.SetOnScheduleMessagePumpWork(nil)

	//指定一个URL地址，或本地html文件目录
	cefApp.StartMainProcess()

	delayed.CreateThread()
	lclwidget.CustomWidgetSetInitialization()
	lcl.RunApp(&window)
}

type WindowDemo struct {
	*lcl.TForm
	focusWorkaround *lcl.TEdit // linux 焦点获取替代
	controlPanel    *lcl.TPanel
	bufferPanel     *cef.TBufferPanel
	chromium        cef.IChromium
}

func (m *WindowDemo) OnFormCreate(sender lcl.IObject) {
	m.SetCaption("Energy - OSR")
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
	m.SetWidth(1400)
	m.SetHeight(900)
	m.ScreenCenter()
	m.chromium = cef.NewChromium(m, nil)

	m.controlPanel = lcl.NewPanel(m)
	m.controlPanel.SetParent(m)
	m.controlPanel.SetAlign(types.AlTop)
	m.controlPanel.SetHeight(25)
	m.controlPanel.SetBevelOuter(types.BvNone)
	m.controlPanel.SetBevelInner(types.BvNone)
	m.controlPanelWidget()

	m.bufferPanel = cef.NewBufferPanel(m)
	m.bufferPanel.SetParent(m)
	m.bufferPanel.SetColor(colors.ClAqua)
	m.bufferPanel.SetTop(50)
	m.bufferPanel.SetLeft(50)
	// 这里设置的宽高还未生效，chromium.SetOnGetViewRect 函数里设置生效
	//m.bufferPanel.SetWidth(600)
	//m.bufferPanel.SetHeight(400)
	m.bufferPanel.SetAlign(types.AlClient) // 同步和客户端一样大小
	m.bufferPanelEvent()
	m.chromiumEvent()
	m.SetOnShow(func(sender lcl.IObject) {
		b := m.chromium.Initialized()
		fmt.Println("init:", b)
		cb := m.chromium.CreateBrowser(nil, "", nil, nil)
		fmt.Println("CreateBrowser:", cb)
		m.chromium.Options().SetBackgroundColor(cef.CefColorSetARGB(0x00, 0x00, 0xff, 0xff))
		m.bufferPanel.CreateIMEHandler()
		m.chromium.InitializeDragAndDrop(m.bufferPanel)
		//m.chromium.LoadUrl("https://www.baidu.com")
	})

}

func (m *WindowDemo) controlPanelWidget() {
	saveDialog := lcl.NewSaveDialog(m)
	saveDialog.SetTitle("OSR Save Page")
	saveDialog.SetFilter("Bitmap files (*.bmp)|*.BMP|Png files (*.png)|*.PNG")

	combox := lcl.NewComboBox(m)
	combox.SetParent(m.controlPanel)
	combox.SetText("https://energy.yanghy.cn")
	items := lcl.NewStringList()
	items.Add("https://energy.yanghy.cn")
	items.Add("https://www.baidu.com")
	combox.SetItems(items)
	combox.SetAlign(types.AlClient)

	btnPanel := lcl.NewPanel(m)
	btnPanel.SetParent(m.controlPanel)
	btnPanel.SetAlign(types.AlRight)
	btnPanel.SetBevelOuter(types.BvNone)
	btnPanel.SetBevelInner(types.BvNone)

	goBtn := lcl.NewButton(m)
	goBtn.SetParent(btnPanel)
	goBtn.SetCaption("GO")
	goBtn.SetAlign(types.AlLeft)
	goBtn.SetOnClick(func(sender lcl.IObject) {
		m.chromium.LoadUrl(combox.Text())
	})

	saveBtn := lcl.NewButton(m)
	saveBtn.SetParent(btnPanel)
	saveBtn.SetCaption("SavePage")
	saveBtn.SetAlign(types.AlRight)
	saveBtn.SetOnClick(func(sender lcl.IObject) {
		if saveDialog.Execute() {
			m.bufferPanel.SaveToFile(saveDialog.FileName())
		}
	})
}

func (m *WindowDemo) chromiumEvent() {
	var (
		popUpBitmap                  *lcl.TBitmap
		tempBitMap                   *lcl.TBitmap
		tempWidth, tempHeight        int32
		tempLineSize                 int
		tempSrcOffset, tempDstOffset int
		src, dst                     uintptr
	)
	m.chromium.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
		fmt.Println("SetOnLoadStart", frame.Url())
	})
	m.chromium.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
		fmt.Println("SetOnLoadEnd", frame.Url())
	})
	m.chromium.SetOnCursorChange(func(sender lcl.IObject, browser *cef.ICefBrowser, cursor consts.TCefCursorHandle, cursorType consts.TCefCursorType, customCursorInfo *cef.TCefCursorInfo) bool {
		fmt.Println("SetOnCursorChange")
		m.bufferPanel.SetCursor(cef.CefCursorToWindowsCursor(cursorType))
		return true
	})
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, client *cef.ICefClient, noJavascriptAccess *bool) bool {
		return true // 阻止弹出窗口
	})
	// 得到显示大小, 这样bufferPanel就显示实际大小
	m.chromium.SetOnGetViewRect(func(sender lcl.IObject, browser *cef.ICefBrowser) *cef.TCefRect {
		//fmt.Println("SetOnGetViewRect")
		var scale = float64(m.bufferPanel.ScreenScale())
		var rect = &cef.TCefRect{}
		rect.X = 0
		rect.Y = 0
		rect.Width = cef.DeviceToLogicalInt32(m.bufferPanel.Width(), scale)
		rect.Height = cef.DeviceToLogicalInt32(m.bufferPanel.Height(), scale)
		return rect
	})
	// 获取设置屏幕信息
	m.chromium.SetOnGetScreenInfo(func(sender lcl.IObject, browser *cef.ICefBrowser) (screenInfo *cef.TCefScreenInfo, result bool) {
		//fmt.Println("SetOnGetScreenInfo")
		var scale = float64(m.bufferPanel.ScreenScale())
		var rect = &cef.TCefRect{}
		screenInfo = new(cef.TCefScreenInfo)
		rect.Width = cef.DeviceToLogicalInt32(m.bufferPanel.Width(), scale)
		rect.Height = cef.DeviceToLogicalInt32(m.bufferPanel.Height(), scale)
		screenInfo.DeviceScaleFactor = t.Single(scale)
		screenInfo.Depth = 0
		screenInfo.DepthPerComponent = 0
		screenInfo.IsMonochrome = 0
		screenInfo.Rect = *rect
		screenInfo.AvailableRect = *rect
		result = true
		return
	})
	// 获取设置屏幕点
	m.chromium.SetOnGetScreenPoint(func(sender lcl.IObject, browser *cef.ICefBrowser, viewX, viewY int32) (screenX, screenY int32, result bool) {
		//fmt.Println("SetOnGetScreenPoint")
		var scale = float64(m.bufferPanel.ScreenScale())
		var viewPoint = types.TPoint{}
		viewPoint.X = cef.LogicalToDeviceInt32(viewX, scale)
		viewPoint.Y = cef.LogicalToDeviceInt32(viewY, scale)
		var screenPoint = m.bufferPanel.ClientToScreen(viewPoint)
		result = true
		screenX = screenPoint.X
		screenY = screenPoint.Y
		//fmt.Println("SetOnGetScreenPoint result:", screenX, screenY)
		return
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
	m.chromium.SetOnIMECompositionRangeChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, selectedRange *cef.TCefRange, characterBoundsCount uint32, characterBounds *cef.TCefRect) {
		fmt.Println("SetOnIMECompositionRangeChanged", *selectedRange, characterBoundsCount, *characterBounds)
	})
	// 在Paint内展示内容到窗口中
	// 使用 bitmap
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
				m.bufferPanel.UpdateBufferDimensions(width, height)
				m.bufferPanel.BufferIsResized(false)
				tempBitMap = m.bufferPanel.Buffer()
				tempBitMap.BeginUpdate(false)
				tempWidth = m.bufferPanel.BufferWidth()
				tempHeight = m.bufferPanel.BufferHeight()
			}
			//byteBufPtr := (*byte)(unsafe.Pointer(buffer))
			rgbSizeOf := int(unsafe.Sizeof(cef.TRGBQuad{}))
			srcStride := int(width) * rgbSizeOf
			for i := 0; i < dirtyRects.Count(); i++ {
				rect := dirtyRects.Get(i)
				if rect.X >= 0 && rect.Y >= 0 {
					tempLineSize = int(math.Min(float64(rect.Width), float64(tempWidth-rect.X))) * rgbSizeOf
					if tempLineSize > 0 {
						tempSrcOffset = int((rect.Y*width)+rect.X) * rgbSizeOf
						tempDstOffset = int(rect.X) * rgbSizeOf
						//src := @pbyte(buffer)[TempSrcOffset];
						src = uintptr(common.GetParamPtr(buffer, tempSrcOffset)) // 拿到src指针, 实际是 byte 指针
						j := int(math.Min(float64(rect.Height), float64(tempHeight-rect.Y)))
						for ii := 0; ii < j; ii++ {
							tempBufferBits := tempBitMap.ScanLine(rect.Y + int32(ii))
							dst = uintptr(common.GetParamPtr(tempBufferBits, tempDstOffset)) //拿到dst指针, 实际是 byte 指针
							rtl.Move(src, dst, tempLineSize)                                 //  也可以自己实现字节复制
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
			//fmt.Println("ha:", ha)
			if ha {
				m.bufferPanel.Invalidate()
			}
		}
		//fmt.Println("SetOnPaint", browser.Identifier(), kind, dirtyRects.Count(), dirtyRects.Get(0), buffer, width, height)
		//fmt.Println(tempWidth, tempHeight, tempForcedResize)
	})
}

func (m *WindowDemo) bufferPanelEvent() {
	m.focusWorkaround = lcl.NewEdit(m)
	m.focusWorkaround.SetParent(m.bufferPanel)
	m.focusWorkaround.SetAutoSize(true)
	m.focusWorkaround.SetAutoSelect(true)
	m.focusWorkaround.SetParentDoubleBuffered(true)
	m.focusWorkaround.SetHideSelection(true)
	m.focusWorkaround.SetParentFont(true)
	m.focusWorkaround.SetParentShowHint(true)
	m.focusWorkaround.SetReadOnly(true)
	m.focusWorkaround.SetTabStop(true)
	m.focusWorkaround.SetTop(-9999) // 控件移动到后面
	m.focusWorkaround.SetOnKeyDown(func(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
		fmt.Println("SetOnOnKeyDown", *key, shift)
		keyEvent := &cef.TCefKeyEvent{}
		if *key != 0 {
			keyEvent.Kind = consts.KEYEVENT_RAW_KEYDOWN
			keyEvent.Modifiers = getModifiers(shift)
			keyEvent.WindowsKeyCode = t.Int32(*key)
			keyEvent.NativeKeyCode = 0
			keyEvent.IsSystemKey = 0           // 0=false, 1=true
			keyEvent.Character = '0'           // #0
			keyEvent.UnmodifiedCharacter = '0' // '#0`
			keyEvent.FocusOnEditableField = 0  // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key = 0;
		}
	})
	m.focusWorkaround.SetOnKeyUp(func(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
		fmt.Println("SetOnOnKeyUp", *key, shift)
		keyEvent := &cef.TCefKeyEvent{}
		if *key != 0 {
			keyEvent.Kind = consts.KEYEVENT_KEYUP
			keyEvent.Modifiers = getModifiers(shift)
			keyEvent.WindowsKeyCode = t.Int32(*key)
			keyEvent.NativeKeyCode = 0
			keyEvent.IsSystemKey = 0           // 0=false, 1=true
			keyEvent.Character = '0'           // #0
			keyEvent.UnmodifiedCharacter = '0' // #0
			keyEvent.FocusOnEditableField = 0  // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key = 0;
		}
	})
	m.focusWorkaround.SetOnKeyPress(func(sender lcl.IObject, key *types.Char) {
		fmt.Println("key:", *key, m.focusWorkaround.Focused())
		//aCefEvent = TCEFKEYEVENT (KIND = KEYEVENT_RAWKEYDOWN;
		//MODIFIERS = 0;
		//WINDOWS_KEY_CODE = 8;
		//NATIVE_KEY_CODE = 22;
		//IS_SYSTEM_KEY = 0;
		//CHARACTER = #$08;
		//UNMODIFIED_CHARACTER = #$08;
		//FOCUS_ON_EDITABLE_FIELD = 0)
		//aCefEvent = TCEFKEYEVENT (KIND = KEYEVENT_RAWKEYDOWN;
		//MODIFIERS = 0;
		//WINDOWS_KEY_CODE = 65;
		//NATIVE_KEY_CODE = 38;
		//IS_SYSTEM_KEY = 0;
		//CHARACTER = 'a';
		//UNMODIFIED_CHARACTER = 'a';
		//FOCUS_ON_EDITABLE_FIELD = 0)

		if m.focusWorkaround.Focused() {
			keyEvent := &cef.TCefKeyEvent{}
			keyEvent.Kind = consts.KEYEVENT_CHAR
			//keyEvent.Modifiers = cef.GetCefKeyboardModifiers(t.WPARAM(asciiCode), 0) // windows
			keyEvent.Modifiers = consts.EVENTFLAG_NONE // windows
			keyEvent.WindowsKeyCode = t.Int32(*key)
			keyEvent.NativeKeyCode = t.Int32(*key)
			keyEvent.IsSystemKey = 0 // 0=false, 1=true
			keyEvent.Character = t.UInt16(*key)
			keyEvent.UnmodifiedCharacter = t.UInt16(*key)
			keyEvent.FocusOnEditableField = 0 // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key := 0;
		}
	})

	m.bufferPanel.SetOnClick(func(sender lcl.IObject) {
		m.bufferPanel.SetFocus()
		m.focusWorkaround.SetFocus()
	})
	m.bufferPanel.SetOnEnter(func(sender lcl.IObject) {
		m.chromium.SetFocus(true)
	})
	m.bufferPanel.SetOnExit(func(sender lcl.IObject) {
		m.chromium.SetFocus(false)
	})
	m.bufferPanel.SetOnResize(func(sender lcl.IObject) {
		if m.bufferPanel.BufferIsResized(false) {
			m.chromium.Invalidate(consts.PET_VIEW)
		} else {
			m.chromium.WasResized()
		}
	})
	m.bufferPanel.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseMoveEvent(mouseEvent, false)
	})
	var (
		// 自己简单处理一下单击·双击·时间和点击次数控制
		// 一搬使用系统的消息时间
		clickTime  int64 = 300 // N 毫秒内连续点击 = 双击
		preTime    int64 = 0
		clickCount int32
	)
	m.bufferPanel.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		//fmt.Println("OnMouseDown:", clickTime, button, shift, x, y)
		if (time.Now().UnixMilli() - preTime) > clickTime {
			clickCount = 1
		} else if clickCount == 2 {
			clickCount = 1 //连续双击 > 恢复单击
		} else {
			clickCount = 2
		}
		preTime = time.Now().UnixMilli()
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseClickEvent(mouseEvent, getButton(button), false, clickCount)
	})
	m.bufferPanel.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		//fmt.Println("SetOnMouseUp:", clickTime, button, shift, x, y)
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseClickEvent(mouseEvent, getButton(button), true, clickCount)
	})
	m.bufferPanel.SetOnMouseWheel(func(sender lcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
		//fmt.Println("SetOnMouseWheel:", shift, wheelDelta, x, y)
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseWheelEvent(mouseEvent, 0, wheelDelta)
	})
	m.bufferPanel.SetOnOnKeyDown(func(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
		//fmt.Println("SetOnOnKeyDown", *key, shift)
		keyEvent := &cef.TCefKeyEvent{}
		if *key != 0 {
			keyEvent.Kind = consts.KEYEVENT_RAW_KEYDOWN
			keyEvent.Modifiers = getModifiers(shift)
			keyEvent.WindowsKeyCode = t.Int32(*key)
			keyEvent.NativeKeyCode = 0
			keyEvent.IsSystemKey = 0           // 0=false, 1=true
			keyEvent.Character = '0'           // #0
			keyEvent.UnmodifiedCharacter = '0' // '#0`
			keyEvent.FocusOnEditableField = 0  // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key = 0;
		}
	})
	m.bufferPanel.SetOnOnKeyUp(func(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
		//fmt.Println("SetOnOnKeyUp", *key, shift)
		keyEvent := &cef.TCefKeyEvent{}
		if *key != 0 {
			keyEvent.Kind = consts.KEYEVENT_KEYUP
			keyEvent.Modifiers = getModifiers(shift)
			keyEvent.WindowsKeyCode = t.Int32(*key)
			keyEvent.NativeKeyCode = 0
			keyEvent.IsSystemKey = 0           // 0=false, 1=true
			keyEvent.Character = '0'           // #0
			keyEvent.UnmodifiedCharacter = '0' // #0
			keyEvent.FocusOnEditableField = 0  // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key = 0;
		}
	})
	m.bufferPanel.SetOnUTF8KeyPress(func(sender lcl.IObject, utf8key *types.TUTF8Char) {
		//fmt.Println("SetOnUTF8KeyPress", utf8key.ToString(), m.bufferPanel.Focused())
		if m.bufferPanel.Focused() {
			if utf8key.Len > 0 {
				var asciiCode int
				fmt.Sscanf(utf8key.ToString(), "%c", &asciiCode)
				keyEvent := &cef.TCefKeyEvent{}
				keyEvent.Kind = consts.KEYEVENT_CHAR
				//keyEvent.Modifiers = cef.GetCefKeyboardModifiers(t.WPARAM(asciiCode), 0) // windows
				keyEvent.WindowsKeyCode = t.Int32(asciiCode)
				keyEvent.NativeKeyCode = 0
				keyEvent.IsSystemKey = 0
				keyEvent.Character = '0'
				keyEvent.UnmodifiedCharacter = '0'
				keyEvent.FocusOnEditableField = 0
				m.chromium.SendKeyEvent(keyEvent)
				//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key := 0;
			}
		}
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

func getButton(Button types.TMouseButton) (result consts.TCefMouseButtonType) {
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
