# Upgrade Log

### 2.4.3

```text
- cli, add check version
- Fix: #32 The rendering process ipc listens for adding a frameId group.  Add a frame object to implement the ipc target.IWindow interface
- Fix: #33 After obtaining the window handle, pop-up window dragging failed
- Fix: #34 --webkit-app-region Invalid after refreshing the page
- Windows and MacOS optimize borderless window styles. Previously, there was no glass shadow when borderless, but now it is the same as the system default,
- Windows, MacOS windows drag and resize using JS-IPC implementation
```

### 2.4.2

```text
- A: CEFTask CefPostTask, CefPostDelayedTask, CefCurrentlyOn
- Fix: Issue of rendering process deadlock caused by IPC nested calls
- U: command-line update
- U: demo gifplay
```

### 2.4.1

```text
1. Modify all demo syso
1. U: Add JS ipc.emit to trigger Go event synchronization mode configuration option, default: `MSync`
1. U: Optimizing the conflict between fullscreen and maximized window.
1. A: demo headless
1. U: command-line manifest requestedExecutionLevel => asInvoker
1. U: command-line add gen windows > icon, syso cmd

Remarks: `ipc.On`

// go: Asynchronous listening mode
ipc.On("name", func(){
	// ...
}, ipcTypes.OnOptions{Mode: ipcTypes.MAsync})

```

### 2.4.0

```text
This version is incompatible with 2.3.x.

1. U: LCLBrowserWindow and ViewsFrameworkBrowserWindow, Add func ChromiumBrowser() ICEFChromiumBrowser
1. U: rename, ipc.emitSync => ipc.emitWait, This trigger has a timeout configuration
1. A: linux arm64 demo startup.sh
1. U: energy cmd version 1.0.6, install golang default version 1.19.13
1. A: Add gif play component
1. U: TCEFWindowComponent.SetOnGetTitleBarHeight param titleBarHeight => *float32
1. A: extension, misc_functions api
1. U: MacOS UI async thread run function
1. U: Logic when using RunOnMainThread to determine IsMessage Loop
1. Fix: vf tary Window state control
1. U: Remove MainFormOnTaskBar configuration and use Enabling MainWindow configuration when the taskbar is not displayed
1. U: all demo, windows import syso
1. Fix: Use VF Application init. RunOnMainThread VF Use ThreadSync UI. ChromiumBrowser LCLBrowserWindow nil bug.
1. U: Go execution IPC listening event changed to asynchronous execution
1. U: Condition judgment when the gate is empty
1. U: browserConfig > BrowserConfig
1. U: IPC NewTarget IWindow > Add Chromium
1. U: Chromium All Event Callback Parameters NativeUInt Type Pointer Passing
1. A: Chromium SendDevToolsMessage function, ExecuteDevToolsMethod Add Result messageId
1. A&U: examples
```

### 2.3.8

```text
1. Fix: Chromium event callback parameter pointer value
2. U: Adjust the timing of the main window settings
3. U: Adjusting the default implementation event to the chrrimbrowser structure
4. Fix: When customizing the layout of Chromium in the window, you cannot drag to change the window size
5. U: energy custom menu modify
6. U: command-line, windows build write [app].manifest to disk
7. Fix: Energy custom event, pop-up window event only triggers once issue
8. Fix: Window Min,Max Size Bug
9. Fix: cmd download cef-framework file name
10. U: Modify some examples

```

### 2.3.7

```text
1. Fix: Window Min,Max Size Bug
2. Fix: Energy custom event, pop-up window event only triggers once issue
```

### 2.3.6

```text
1. MacOS 增加 touch bar 支持， 和touchbar示例
2. MacOS 无标题栏窗口状态控制
3. 删除示例下的多于icon资源文件
4. 增加 ipc 多窗口通信示例
5. 修改托盘示例
5. 增加LCL支持主窗口配置，关闭主窗口后，如果在多窗口时直到最后一个窗口关闭才退出应用
6. 增加WindowsXP SP3支持
7. 修改部分API字符串使用TString类
8. 为了支持Go的底版本将所有any类型改为interface类型
9. energy最底支持Go1.11版本
10. 修复chromium相关事件回调函数参数
11. 增加部分API判断, 对CEF API不支持CEF49, 未判断完全，但不影响，CEF49支持的API不如CEF新版本的多
12. 命令行工具优化，未增加对WindowsXP的安装，目前WindowsXP需要手动安装
13. 命令工具增加bindata命令，当Go版本小于1.16时，为支持Embed内嵌资源接口
14. 优化LCL托盘可以同时创建多个
15. 增加一些energy还未实现的CEF API
16. 优化预先创建下一个子弹出窗口
17. 修复一些错误，记录结构类型调用 API 时传递指针错误问题
```

### 2.3.5

```text
修改所有*.go文件名 中横线 -, 改为下划线 _
修改独立子进程示例
增加一窗口多Chromium示例
升级 liblclbinres v2.3.5
修复readme.me一些错误描述
修改ipc, net socket 端口号默认随机获取, net socket 在Windows10 Build < 17063 版本开启， 原固定19878端口
调整命令行工具编译命令
    编译内置dll调整
     windows:
         386: -tags="tempdll latest"
         amd64: -tags="tempdll latest"
     windows(Windows 7, 8/8.1 and Windows Server 2012):
         386: -tags="tempdll 109"
         amd64: -tags="tempdll 109"
     linux(gtk3):
         amd64: -tags="tempdll latest"
         arm64: -tags="tempdll latest"
     linux(gtk2):
         amd64: -tags="tempdll 106"
         arm64: -tags="tempdll 106"
     macos:
         amd64: -tags="tempdll latest"
         arm64: -tags="tempdll latest"
```

### 2.3.4

```text
升级 118.7.1
修改了Get和Set同一属性过程使用
增加 GlobalCEFApp.ChromePolicyId
移除 Browser, Chromium accept_language_list
增加 TChromiumOptions.ChromeZoomBubble
增加 TChromium.IncZoomCommand
增加 TChromium.DecZoomCommand
增加 TChromium.ResetZoomCommand
增加 TChromium.DefaultZoomLevel
增加 TChromium.CanIncZoom
增加 TChromium.CanDecZoom
增加 TChromium.CanResetZoom
增加 TChromium.Fullscreen
增加 TChromium.ExitFullscreen
增加 ICefDragData.GetFilePaths


添加 https://crbug.com/1500371 https://bitbucket.org/chromiumembedded/cef/commits/99817d2d3ebf5983ea4491f8770ef1e581554f91 解决方法
在全屏窗口退出时更新 CSS（修复 #3597） https://bitbucket.org/chromiumembedded/cef/commits/9d1cdd020f4bc877cb9675afeed439c6e4749ec2
在调整边框大小之前对 PiP 可拖动区域进行命中测试（请参阅问题 #3566） https://bitbucket.org/chromiumembedded/cef/commits/38848f1780ea59b8b8819e06250b25aacd5c45c6
```

### 2.3.3

```text
优化和修复一些问题
1. LCL无边框窗口, 点击任务栏不能切换窗口问题
2. WndProc 回调函数，修改&增加winapi函数，增加 HDWP 类型
3. 增加 lcl 窗口配置函数: 扩展事件-SetOnWndProc, SetOnPaint, 边框Frameless,FramelessForLine, SetRoundRectRgn
4. windows, lcl 窗口调整屏幕缩放比拖拽区域计算位置不正确问题
5. 创建 Application 初始化配置增加默认开启GPU加速
6. 移除生成图标示例修改部分示例
```

### 2.3.1

```text
- 增加底层动态库异常捕获, 仅Windows, MacOS
- 升级CEF从109直接跳到117, 110~116版本的liblcl构建跳过, 此时会增加和移除一些api
- 升级命令行工具1.0.2,增加兼容Windows7 CEF109
- 增加底层库windows, macos异常捕获
- energy 底层依赖库自动化构建和发布
- 增加一些示例：屏幕截取，模拟事件，IPC Go to Go。
- 修复一些已知问题

升级 liblcl v2.3.1

```

### 2.3.0

```text
主要：
  2.3.0 对部分回调函数做出调整，主要增加了 cef.IBrowserWindow 当前窗口参数

  增加 静态资源使用本地或内置资源加载, 暂时不能加载视频资源。
   本地或内置资源加载，xhr 代理请求配置支持ssl
  修复 linux(高版本) gtk3(默认)加载动态库错误问题。
  修复 linux gtk3 无法切换英文问题, 
  优化 Mac开发环境 energy_env=dev > env=dev

其它：
 完善命令行工具
  优化 install 开发环境全自动安装:
  增加 init 应用项目初始化
  增加 build 构建&编译应用执行文件
  增加 package 制作应用安装包
 依赖库升级
  golcl v1.0.7
  liblcl v1.0.4

 修复其它已知问题
```

### 2.2.4

```text
1. 增加常用示例
2. 增加窗口焦点
3. 同步govcl库和liblcl库
4. 升级命令行工具
5. 增加 linux arm 架构二进制包

```

### 2.2.3

```text
U: liblclbinres v1.0.2
U: demo main-browser-window
U: Return value(float32) method
A: demo Window IScreen
A: Window IScreen
A: displayRef proc api
Fix: Multiple display window centering issue
U: demo scheme
U: demo popup-sub-window elliptic
U: Optimize custom window drag and drop creation logic
Fix: potential problem, proc api return string error
Fix: v8value bug, string value error
U: demo frameless
A: window, full screen model, add common attributes
U: demo frameless
U: browser window, full screen
U: demo window state
U: demo frameless, fullscreen
U: window state
U: chromium context-menu-command callback
U: context-menu
A: open tab url callback event
U: demo liblcl autoupdate , linux lcl widget init
U: tempdll README.md
U: gen libbin
A: energy command line, Set energy framework development environment
U: energy command line, support linux select gtk2 or gtk3 framework
U: remove CustomWidgetSetFinalization
U: demo context-menu
U: LCL CloseBrowserWindow RunOnMainThread
A: on message paint
U: on message struct
U: demo test
A: TForm WM Message - > NotifyMoveOrResizeStarted
A: TForm WM Message
Fix: aux-viewsource, linux
U: demos, ui use gtk3, IconFS = xxx.png, other IconFS = xxx.ico
A: gtk2 support for CEF 106.1.1
```

### 2.2.2

```text
修改和优化已知问题
https://gitee.com/energye/energy/commits/v2.2.2
```

### 2.2.1

```text
U: demo msgbox
A: windows demo custom-browser-create
U: .gitattributes
U: tempdll for mac load liblcl
A: demo tempdll
A: go build -tags="tempdll" open TempDLL, import liblclbinres.
U: window drag switch
U: window drag explanatory
U: viewsource
A: demo drag file
R: remove BrowserWindow.Config.EnableWebkitAppRegion
A: demo custom-drag-window
U: lcl customer drag
U: demo frameless for mac, no hide caption
U: lcl window close for mac, 1.show,2.hide
U: demo
U: remove const.IsMessageLoop, add application.IsMessageLoop()
U: message const
Merge branch 'main' into dev
U: win32 consts, cefwinapi
Update README.zh_CN.md
Update README.md
A: lcl custom window drag
A: FramelessForDefault
A: CEF FileDialog
Fix: Chromium & DialogHandler onFileDialog callback args:acceptFiltersList
A: 4 types systray demo
A: windows frameless Control Window Border
U: demo-windows transparent
```

### 2.2.0-beta

```text
1. 增加 browser RunFileDialog 回调函数, 该功能是使用CEF打开选择文件弹窗
2. 增加 DownloadImage 回调函数
3. 增加 CEF的 MenuModel、View、Button、LabelButton、MenuButton、Panel、textfield 组件相关 Api
4. 重命名示例目录名
5. 增加 其它功能处理回调函数
6. 修改底层实现类变更接口
7. v2.2.0开始兼容老版本CEF 87 ~ 89, 新版本最低兼容到109
8. 优化energy框架中的window相关事件
```

