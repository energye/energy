//go:build windows
// +build windows

package form

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/ext/opengl"
	t "github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
	"github.com/go-gl/gl/v2.1/gl"
	"unsafe"
)

var App *cef.TCEFApplication

type WindowForm struct {
	*lcl.TForm
	openGL              *opengl.TOpenGLControl
	glInitialized       bool
	textureID           uint32
	textureWidth        int32
	textureHeight       int32
	isClose             bool
	updateRect          cef.TCefRect
	chromium            cef.IChromium
	chromiumCreateTimer *lcl.TTimer
}

func (m *WindowForm) OnFormCreate(sender lcl.IObject) {
	m.SetCaption("energy - opengl")
	m.SetWidth(1000)
	m.SetHeight(600)
	m.ScreenCenter()

	m.openGLInit()

	m.chromiumInit()

	m.SetOnShow(func(sender lcl.IObject) {
		if m.chromium.Initialized() {
			m.chromium.WasHidden(false)
			m.chromium.SetFocus(true)
		} else {
			//不透明的白色背景色
			m.chromium.Options().SetBackgroundColor(cef.CefColorSetARGB(0xff, 0xff, 0xff, 0xff))
			if m.chromium.CreateBrowser(nil, "", nil, nil) {
				m.chromium.InitializeDragAndDrop(m.openGL)
			} else {
				m.chromiumCreateTimer.SetEnabled(true)
			}
		}
	})
	m.SetOnHide(func(sender lcl.IObject) {
		m.chromium.SetFocus(false)
		m.chromium.WasHidden(true)
	})
	m.SetOnDestroy(func(sender lcl.IObject) {
		m.chromium.ShutdownDragAndDrop()
		if m.textureID != 0 {
			gl.DeleteTextures(1, &m.textureID)
			VerifyOpenGLErrors()
		}
	})

	m.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) {
		*canClose = m.isClose
		if !m.isClose {
			m.chromium.CloseBrowser(true)
		}
	})

	m.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		switch msg.Msg {
		case messages.WM_DPICHANGED:
			App.UpdateDeviceScaleFactor()
			if m.chromium != nil {
				m.chromium.NotifyScreenInfoChanged()
				m.chromium.WasResized()
			}
		}
	})
}

func (m *WindowForm) createTimer(sender lcl.IObject) {
	m.chromiumCreateTimer.SetEnabled(false)
	if m.chromium.CreateBrowser(nil, "", nil, nil) {
		m.chromium.InitializeDragAndDrop(m.openGL)
	} else {
		m.chromiumCreateTimer.SetEnabled(true)
	}
}

func (m *WindowForm) chromiumPaint(sender lcl.IObject, browser *cef.ICefBrowser, kind consts.TCefPaintElementType, dirtyRects *cef.TCefRectArray, buffer uintptr, width, height int32) {
	if m.initializeOpenGL() {
		gl.Enable(gl.TEXTURE_2D)
		VerifyOpenGLErrors()

		gl.BindTexture(gl.TEXTURE_2D, m.textureID)
		VerifyOpenGLErrors()

		switch kind {
		case consts.PET_VIEW:
			TempOldWidth := m.textureWidth
			TempOldHeight := m.textureHeight
			m.textureWidth = width
			m.textureHeight = height

			gl.PixelStorei(gl.UNPACK_ROW_LENGTH, m.textureWidth)
			VerifyOpenGLErrors()
			var rect *cef.TCefRect
			if dirtyRects.Count() == 1 {
				rect = dirtyRects.Get(0)
			}
			if (m.textureWidth != TempOldWidth) || (m.textureHeight != TempOldHeight) || ((dirtyRects.Count() == 1) &&
				(rect.X == 0) && (rect.Y == 0) && (rect.Width == m.textureWidth) && (rect.Height == m.textureHeight)) {
				// Update/resize the whole texture.
				gl.PixelStorei(gl.UNPACK_SKIP_PIXELS, 0)
				VerifyOpenGLErrors()

				gl.PixelStorei(gl.UNPACK_SKIP_ROWS, 0)
				VerifyOpenGLErrors()

				gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, m.textureWidth, m.textureHeight, 0, gl.BGRA, gl.UNSIGNED_INT_8_8_8_8_REV, unsafe.Pointer(buffer))
				VerifyOpenGLErrors()
			} else {
				for i := 0; i < dirtyRects.Count(); i++ {
					TempRect := dirtyRects.Get(i)
					if (TempRect.X+TempRect.Width <= m.textureWidth) && (TempRect.Y+TempRect.Height <= m.textureHeight) {
						gl.PixelStorei(gl.UNPACK_SKIP_PIXELS, TempRect.X)
						VerifyOpenGLErrors()

						gl.PixelStorei(gl.UNPACK_SKIP_ROWS, TempRect.Y)
						VerifyOpenGLErrors()

						gl.TexSubImage2D(gl.TEXTURE_2D, 0, TempRect.X, TempRect.Y, TempRect.Width, TempRect.Height,
							gl.BGRA, gl.UNSIGNED_INT_8_8_8_8_REV, unsafe.Pointer(buffer))
						VerifyOpenGLErrors()
					}
				}
			}

		case consts.PET_POPUP:
		}

		// Disable 2D textures.
		gl.Disable(gl.TEXTURE_2D)
		VerifyOpenGLErrors()

		m.openGL.Invalidate()
	}
}

func (m *WindowForm) chromiumInit() {
	m.chromiumCreateTimer = lcl.NewTimer(m)
	m.chromiumCreateTimer.SetEnabled(false)
	m.chromiumCreateTimer.SetOnTimer(m.createTimer)

	m.chromium = cef.NewChromium(m, nil)
	m.chromium.SetDefaultURL("http://localhost:22022/index.html")

	m.chromium.SetOnGetScreenInfo(func(sender lcl.IObject, browser *cef.ICefBrowser) (screenInfo *cef.TCefScreenInfo, result bool) {
		scale := App.DeviceScaleFactor()
		rect := &cef.TCefRect{}
		screenInfo = new(cef.TCefScreenInfo)
		rect.Width = cef.DeviceToLogicalInt32(m.openGL.Width(), float64(scale))
		rect.Height = cef.DeviceToLogicalInt32(m.openGL.Height(), float64(scale))
		screenInfo.DeviceScaleFactor = t.Single(scale)
		screenInfo.Depth = 0
		screenInfo.DepthPerComponent = 0
		screenInfo.IsMonochrome = 0 // bool
		screenInfo.Rect = *rect
		screenInfo.AvailableRect = *rect
		result = true
		return
	})
	m.chromium.SetOnGetScreenPoint(func(sender lcl.IObject, browser *cef.ICefBrowser, viewX, viewY int32) (screenX, screenY int32, result bool) {
		scale := App.DeviceScaleFactor()
		viewPoint := types.TPoint{}
		viewPoint.X = cef.LogicalToDeviceInt32(viewX, float64(scale))
		viewPoint.Y = cef.LogicalToDeviceInt32(viewY, float64(scale))
		screenPoint := m.openGL.ClientToScreen(viewPoint)
		screenX = screenPoint.X
		screenY = screenPoint.Y
		result = true
		return
	})
	m.chromium.SetOnGetViewRect(func(sender lcl.IObject, browser *cef.ICefBrowser) *cef.TCefRect {
		scale := App.DeviceScaleFactor()
		rect := &cef.TCefRect{}
		rect.X = 0
		rect.Y = 0
		rect.Width = cef.DeviceToLogicalInt32(m.openGL.Width(), float64(scale))
		rect.Height = cef.DeviceToLogicalInt32(m.openGL.Height(), float64(scale))
		return rect
	})
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		m.isClose = true
		cef.RunOnMainThread(func() {
			m.Close()
		})
	})
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		return true // 阻止弹出窗口
	})

	m.chromium.SetOnCursorChange(func(sender lcl.IObject, browser *cef.ICefBrowser, cursor consts.TCefCursorHandle, cursorType consts.TCefCursorType, customCursorInfo *cef.TCefCursorInfo) bool {
		m.openGL.SetCursor(cef.CefCursorToWindowsCursor(cursorType))
		return true
	})
	m.chromium.SetOnPaint(m.chromiumPaint)
}

func VerifyOpenGLErrors(funcName ...string) {
	err := gl.GetError()
	if err != gl.NO_ERROR {
		if len(funcName) == 1 {
			fmt.Println("Open GL error, func:", funcName[0], "code:", err)
		} else {
			fmt.Println("Open GL error:", err)
		}
	}
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
