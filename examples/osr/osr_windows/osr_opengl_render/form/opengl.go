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
	"github.com/go-gl/gl/v2.1/gl"
	"time"
	"unsafe"
)

func (m *WindowForm) openGLInit() {
	m.openGL = opengl.NewOpenGLControl(m)
	m.openGL.SetName("openGL")
	m.openGL.SetParent(m)

	//m.openGL.SetAlign(types.AlClient)

	m.openGL.SetTop(50)
	m.openGL.SetLeft(50)
	m.openGL.SetWidth(m.Width() - 100)
	m.openGL.SetHeight(m.Height() - 100)
	m.openGL.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkBottom, types.AkRight))
	m.openGL.SetOpenGLMajorVersion(2)
	m.openGL.SetOpenGLMinorVersion(1)

	m.openGLEvent()
}

func (m *WindowForm) initializeOpenGL() bool {
	if m.glInitialized {
		return true
	}
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Hint(gl.POLYGON_SMOOTH_HINT, gl.NICEST)
	VerifyOpenGLErrors()

	backgroundColor := m.chromium.Options().BackgroundColor()
	gl.ClearColor(float32(cef.CefColorGetR(backgroundColor)/255), float32(cef.CefColorGetG(backgroundColor)/255), float32(cef.CefColorGetB(backgroundColor)/255), 1)
	VerifyOpenGLErrors()

	// 非power-of-2纹理正确渲染所必需的
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	VerifyOpenGLErrors()

	// 创建纹理
	gl.GenTextures(1, &m.textureID)
	VerifyOpenGLErrors()
	if m.textureID == 0 {
		return false
	}

	gl.BindTexture(gl.TEXTURE_2D, m.textureID)
	VerifyOpenGLErrors()

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	VerifyOpenGLErrors()

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	VerifyOpenGLErrors()

	gl.TexEnvf(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	VerifyOpenGLErrors()

	m.glInitialized = true
	return true
}

func (m *WindowForm) openGLEvent() {
	m.openGL.SetOnResize(func(sender lcl.IObject) {
		m.chromium.WasResized()
	})

	m.openGL.SetOnPaint(m.openGLPaint)

	m.openGL.SetOnClick(func(sender lcl.IObject) {
		m.openGL.SetFocus()
	})
	m.openGL.SetOnEnter(func(sender lcl.IObject) {
		m.chromium.SetFocus(true)
	})
	m.openGL.SetOnExit(func(sender lcl.IObject) {
		m.chromium.SetFocus(false)
	})
	// 鼠标移动
	m.openGL.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(App.DeviceScaleFactor()))
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
	m.openGL.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
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
		cef.DeviceToLogicalMouse(mouseEvent, float64(App.DeviceScaleFactor()))
		m.chromium.SendMouseClickEvent(mouseEvent, getButton(button), false, clickCount)
	})
	// 鼠标事件 点击抬起
	m.openGL.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(App.DeviceScaleFactor()))
		m.chromium.SendMouseClickEvent(mouseEvent, getButton(button), true, clickCount)
	})
	// 鼠标滚轮事件
	m.openGL.SetOnMouseWheel(func(sender lcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
		mouseEvent := &cef.TCefMouseEvent{}
		mouseEvent.X = x
		mouseEvent.Y = y
		mouseEvent.Modifiers = getModifiers(shift)
		cef.DeviceToLogicalMouse(mouseEvent, float64(App.DeviceScaleFactor()))
		m.chromium.SendMouseWheelEvent(mouseEvent, 0, wheelDelta)
	})
	// 键盘事件 按下
	m.openGL.SetOnKeyDown(func(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
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
		}
	})
	// 键盘事件 抬起
	m.openGL.SetOnKeyUp(func(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
		//fmt.Println("SetOnOnKeyUp", *key, shift)
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
		}
	})
	// 键盘事件, 上字
	m.openGL.SetOnUTF8KeyPress(func(sender lcl.IObject, utf8key *types.TUTF8Char) {
		//fmt.Println("SetOnUTF8KeyPress", utf8key.ToString(), m.bufferPanel.Focused())
		if m.openGL.Focused() {
			if utf8key.Len > 0 {
				var asciiCode int
				fmt.Sscanf(utf8key.ToString(), "%c", &asciiCode)
				keyEvent := &cef.TCefKeyEvent{}
				keyEvent.Kind = consts.KEYEVENT_CHAR
				keyEvent.Modifiers = cef.GetCefKeyboardModifiers(t.WPARAM(asciiCode), 0)
				keyEvent.WindowsKeyCode = int32(asciiCode)
				keyEvent.NativeKeyCode = 0
				keyEvent.IsSystemKey = 0
				keyEvent.Character = '0'
				keyEvent.UnmodifiedCharacter = '0'
				keyEvent.FocusOnEditableField = 0
				m.chromium.SendKeyEvent(keyEvent)
			}
		}
	})
}

type TVertex struct {
	tu, tv  float32
	x, y, z float32
}

var vertices = []TVertex{
	{tu: 0.0, tv: 1.0, x: -1.0, y: -1.0, z: 0.0},
	{tu: 1.0, tv: 1.0, x: 1.0, y: -1.0, z: 0.0},
	{tu: 1.0, tv: 0.0, x: 1.0, y: 1.0, z: 0.0},
	{tu: 0.0, tv: 0.0, x: -1.0, y: 1.0, z: 0.0},
}

func (m *WindowForm) openGLPaint(sender lcl.IObject) {
	if m.textureWidth == 0 || m.textureHeight == 0 || !m.glInitialized {
		return
	}
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	VerifyOpenGLErrors()

	gl.MatrixMode(gl.MODELVIEW)
	VerifyOpenGLErrors()

	gl.LoadIdentity()
	VerifyOpenGLErrors()

	gl.Viewport(0, 0, m.textureWidth, m.textureHeight)
	VerifyOpenGLErrors()

	gl.MatrixMode(gl.PROJECTION)
	VerifyOpenGLErrors()

	gl.LoadIdentity()
	VerifyOpenGLErrors()

	// Draw the background gradient.
	gl.PushAttrib(gl.ALL_ATTRIB_BITS)
	VerifyOpenGLErrors()

	gl.Begin(gl.QUADS)
	gl.Color4f(1.0, 0.0, 0.0, 1.0) // red
	gl.Vertex2f(-1.0, -1.0)
	gl.Vertex2f(1.0, -1.0)
	gl.Color4f(0.0, 0.0, 1.0, 1.0) // blue
	gl.Vertex2f(1.0, 1.0)
	gl.Vertex2f(-1.0, 1.0)
	gl.End()
	VerifyOpenGLErrors()

	gl.PopAttrib()
	VerifyOpenGLErrors()

	// Enable 2D textures.
	gl.Enable(gl.TEXTURE_2D)
	VerifyOpenGLErrors()

	// Draw the facets with the texture.
	gl.BindTexture(gl.TEXTURE_2D, m.textureID)
	VerifyOpenGLErrors()

	gl.InterleavedArrays(gl.T2F_V3F, 0, unsafe.Pointer(&vertices[0]))
	VerifyOpenGLErrors()

	gl.DrawArrays(gl.QUADS, 0, 4)
	VerifyOpenGLErrors()

	// Disable 2D textures.
	gl.Disable(gl.TEXTURE_2D)
	VerifyOpenGLErrors()

	m.openGL.SwapBuffers()

}
