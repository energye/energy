package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	t "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"math"
	"time"
	"unsafe"
)

//

func main() {
	println("必须将执行文件(osr_linux)和子进程执行文件(osr_subprocess)放置在同一目录")
	cef.GlobalInit(nil, nil)
	var window = &WindowForm{}
	//创建应用
	app := cef.CreateApplication()
	cef.SetApplication(app)
	app.SetBrowserSubprocessPath("osr_subprocess")
	// OSR 离屏渲染
	app.SetWindowlessRenderingEnabled(true)
	// 指定消息模式
	app.SetExternalMessagePump(true)
	app.SetMultiThreadedMessageLoop(false)
	app.SetSetCurrentDir(true)
	app.SetDisableZygote(true)
	app.SetOnScheduleMessagePumpWork(nil)

	global := cef.GlobalWorkSchedulerCreate(nil)
	fmt.Println("start main")
	start := app.StartMainProcess()
	fmt.Println("start main", start)
	if start {
		global.CreateThread()
		api.CustomWidgetSetInitialization()
		// 运行应用, 传入窗口
		lcl.RunApp(&window)
	}
}

// 窗口
type WindowForm struct {
	*lcl.TForm
	focusWorkaround *lcl.TEdit // linux 焦点获取替代
	controlPanel    *lcl.TPanel
	bufferPanel     *cef.TBufferPanel
	chromium        cef.IChromium
}

// 窗口创建时回调事件
func (m *WindowForm) OnFormCreate(sender lcl.IObject) {
	fmt.Println("window create")
	m.SetCaption("Energy - OSR")
	m.SetWidth(800)
	m.SetHeight(600)
	m.ScreenCenter()
	fmt.Println("window create 1")
	// 创建 chromium
	m.chromium = cef.NewChromium(m, nil)
	m.chromiumEvent() //注册 chromium 事件

	// 创建 地址栏 panel
	m.controlPanel = lcl.NewPanel(m)
	m.controlPanel.SetParent(m)
	m.controlPanel.SetAlign(types.AlTop)
	m.controlPanel.SetHeight(25)
	m.controlPanel.SetBevelOuter(types.BvNone)
	m.controlPanel.SetBevelInner(types.BvNone)
	m.controlPanelWidget() // 创建地址栏组件

	// 创建 bufferPanel
	m.bufferPanel = cef.NewBufferPanel(m)
	m.bufferPanel.SetParent(m)
	m.bufferPanel.SetColor(colors.ClAqua)
	m.bufferPanel.SetTop(50)
	m.bufferPanel.SetLeft(50)
	// 这里设置的宽高还未生效，chromium.SetOnGetViewRect 函数里设置生效
	//m.bufferPanel.SetWidth(600)
	//m.bufferPanel.SetHeight(400)
	m.bufferPanel.SetAlign(types.AlClient) // 宽高同步和主窗口一样大小
	m.bufferPanelEvent()                   //注册 bufferPanel 事件
	m.SetOnShow(func(sender lcl.IObject) { //显示窗口时回调
		fmt.Println("on show")
		// 在这里创建初始化和创建chromium
		m.chromium.Initialized()
		m.chromium.CreateBrowser(nil, "", nil, nil)
		m.chromium.Options().SetBackgroundColor(cef.CefColorSetARGB(0x00, 0x00, 0xff, 0xff)) // 可选, 随便设置个背景色
		// CreateIMEHandler 当Panel1具有有效句柄时
		// 需要在创建浏览器之前创建IME处理程序。
		// 如果用户不需要“输入法编辑器”，则可以跳过此操作
		m.bufferPanel.CreateIMEHandler()
		m.chromium.InitializeDragAndDrop(m.bufferPanel)
	})

}

func (m *WindowForm) chromiumEvent() {
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
		m.bufferPanel.SetCursor(cef.CefCursorToWindowsCursor(cursorType))
		return true
	})
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		return true
	})
	m.chromium.SetOnTooltip(func(sender lcl.IObject, browser *cef.ICefBrowser, text *string) (result bool) {
		fmt.Println("SetOnTooltip", *text)
		result = true
		m.bufferPanel.SetHint(*text)
		m.bufferPanel.SetShowHint(len(*text) > 0)
		return
	})
	// 得到显示大小, 这样bufferPanel就显示实际大小
	m.chromium.SetOnGetViewRect(func(sender lcl.IObject, browser *cef.ICefBrowser) *cef.TCefRect {
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
		var scale = float64(m.bufferPanel.ScreenScale())
		var viewPoint = types.TPoint{}
		viewPoint.X = cef.LogicalToDeviceInt32(viewX, scale)
		viewPoint.Y = cef.LogicalToDeviceInt32(viewY, scale)
		var screenPoint = m.bufferPanel.ClientToScreen(viewPoint)
		result = true
		screenX = screenPoint.X
		screenY = screenPoint.Y
		return
	})
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		m.chromium.LoadUrl("https://www.baidu.com")
	})
	m.chromium.SetOnPopupShow(func(sender lcl.IObject, browser *cef.ICefBrowser, show bool) {
		if m.chromium != nil {
			m.chromium.Invalidate(consts.PET_VIEW)
		}
	})
	m.chromium.SetOnPopupSize(func(sender lcl.IObject, browser *cef.ICefBrowser, rect *cef.TCefRect) {
		screenScale := m.bufferPanel.ScreenScale()
		fmt.Println("PopupSize - rect:", rect, "screenScale:", screenScale)
		cef.LogicalToDeviceRect(*rect, float64(screenScale))
		fmt.Println("PopupSize - rect:", rect, "screenScale:", screenScale)
	})
	// windows IME
	m.chromium.SetOnIMECompositionRangeChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, selectedRange cef.TCefRange, characterBoundsCount uint32, characterBounds cef.TCefRect) {
		fmt.Println("SetOnIMECompositionRangeChanged", selectedRange, characterBoundsCount, characterBounds)
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
			if m.HandleAllocated() {
				m.bufferPanel.Invalidate()
			}
		}
	})
}

func (m *WindowForm) bufferPanelEvent() {
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
			keyEvent.WindowsKeyCode = int32(*key)
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
			keyEvent.WindowsKeyCode = int32(*key)
			keyEvent.NativeKeyCode = 0
			keyEvent.IsSystemKey = 0           // 0=false, 1=true
			keyEvent.Character = '0'           // #0
			keyEvent.UnmodifiedCharacter = '0' // #0
			keyEvent.FocusOnEditableField = 0  // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key = 0;
		}
	})
	// 中文输入不知到砸搞。
	// PGtkWidget
	//  g_signal_connect_data(aWidget, 'key-press-event', TGCallback(@GTKKeyPress), nil, nil, 0)
	//  g_signal_connect_data(aWidget, 'key-release-event', TGCallback(@GTKKeyPress), nil, nil, 0)
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
			keyEvent.WindowsKeyCode = int32(*key)
			keyEvent.NativeKeyCode = int32(*key)
			keyEvent.IsSystemKey = 0 // 0=false, 1=true
			keyEvent.Character = *key
			keyEvent.UnmodifiedCharacter = *key
			keyEvent.FocusOnEditableField = 0 // 0=false, 1=true
			m.chromium.SendKeyEvent(keyEvent)
			//if (Key in [VK_LEFT, VK_RIGHT, VK_UP, VK_DOWN, VK_TAB]) then Key := 0;
		}
	})

	m.bufferPanel.SetOnClick(func(sender lcl.IObject) {
		m.bufferPanel.SetFocus()
		m.focusWorkaround.SetFocus() // 焦点获取替代
	})
	m.bufferPanel.SetOnEnter(func(sender lcl.IObject) {
		m.chromium.SetFocus(true)
	})
	m.bufferPanel.SetOnExit(func(sender lcl.IObject) {
		m.chromium.SetFocus(false)
	})
	// panel Align 设置为 client 时， 如果调整窗口大小
	// 该函数被回调, 需要调用 WasResized 调整页面同步和主窗口一样
	m.bufferPanel.SetOnResize(func(sender lcl.IObject) {
		if m.bufferPanel.BufferIsResized(false) {
			m.chromium.Invalidate(consts.PET_VIEW)
		} else {
			m.chromium.WasResized()
		}
	})
	// 鼠标移动
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
	// 鼠标事件 点击按下
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
	// 鼠标事件 点击抬起
	m.bufferPanel.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		//fmt.Println("SetOnMouseUp:", clickTime, button, shift, x, y)
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseClickEvent(mouseEvent, getButton(button), true, clickCount)
	})
	// 鼠标滚轮事件
	m.bufferPanel.SetOnMouseWheel(func(sender lcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
		//fmt.Println("SetOnMouseWheel:", shift, wheelDelta, x, y)
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(m.bufferPanel.ScreenScale()))
		m.chromium.SendMouseWheelEvent(mouseEvent, 0, wheelDelta)
	})
}

func (m *WindowForm) controlPanelWidget() {
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
