### 2.4.0
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

### 2.3.8
1. Fix: Chromium event callback parameter pointer value
1. U: Adjust the timing of the main window settings
1. U: Adjusting the default implementation event to the chrrimbrowser structure
1. Fix: When customizing the layout of Chromium in the window, you cannot drag to change the window size
1. U: energy custom menu modify
1. U: command-line, windows build write [app].manifest to disk
1. Fix: Energy custom event, pop-up window event only triggers once issue
1. Fix: Window Min,Max Size Bug
1. Fix: cmd download cef-framework file name
1. U: Modify some examples


### 2.3.7
1. Fix: Window Min,Max Size Bug
1. Fix: Energy custom event, pop-up window event only triggers once issue

### 2.3.6
1. MacOS 增加 touch bar 支持， 和touchbar示例
1. MacOS 无标题栏窗口状态控制
1. 删除示例下的多于icon资源文件
1. 增加 ipc 多窗口通信示例
1. 修改托盘示例
1. 增加LCL支持主窗口配置，关闭主窗口后，如果在多窗口时直到最后一个窗口关闭才退出应用
1. 增加WindowsXP SP3支持
1. 修改部分API字符串使用TString类
1. 为了支持Go的底版本将所有any类型改为interface类型
1. energy最底支持Go1.11版本
1. 修复chromium相关事件回调函数参数
1. 增加部分API判断, 对CEF API不支持CEF49, 未判断完全，但不影响，CEF49支持的API不如CEF新版本的多
1. 命令行工具优化，未增加对WindowsXP的安装，目前WindowsXP需要手动安装
1. 命令工具增加bindata命令，当Go版本小于1.16时，为支持Embed内嵌资源接口
1. 优化LCL托盘可以同时创建多个
1. 增加一些energy还未实现的CEF API
1. 优化预先创建下一个子弹出窗口
1. 修复一些错误，记录结构类型调用 API 时传递指针错误问题

### 2.3.5
1. 修改所有*.go文件名 中横线 -, 改为下划线 _
1. 修改独立子进程示例
1. 增加一窗口多Chromium示例
1. 升级 liblclbinres v2.3.5
1. 修复readme.me一些错误描述
1. 修改ipc, net socket 端口号默认随机获取, net socket 在Windows10 Build < 17063 版本开启， 原固定19878端口
1. 调整命令行工具编译命令
1. 编译内置dll调整
```go
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
1. 升级 118.7.1
1. 修改了Get和Set同一属性过程使用
1. 增加 GlobalCEFApp.ChromePolicyId
1. 移除 Browser, Chromium accept_language_list
1. 增加 TChromiumOptions.ChromeZoomBubble
1. 增加 TChromium.IncZoomCommand
1. 增加 TChromium.DecZoomCommand
1. 增加 TChromium.ResetZoomCommand
1. 增加 TChromium.DefaultZoomLevel
1. 增加 TChromium.CanIncZoom
1. 增加 TChromium.CanDecZoom
1. 增加 TChromium.CanResetZoom
1. 增加 TChromium.Fullscreen
1. 增加 TChromium.ExitFullscreen
1. 增加 ICefDragData.GetFilePaths


1. 添加 https://crbug.com/1500371 https://bitbucket.org/chromiumembedded/cef/commits/99817d2d3ebf5983ea4491f8770ef1e581554f91 解决方法
1. 在全屏窗口退出时更新 CSS（修复 #3597） https://bitbucket.org/chromiumembedded/cef/commits/9d1cdd020f4bc877cb9675afeed439c6e4749ec2
1. 在调整边框大小之前对 PiP 可拖动区域进行命中测试（请参阅问题 #3566） https://bitbucket.org/chromiumembedded/cef/commits/38848f1780ea59b8b8819e06250b25aacd5c45c6

### 2.3.3
优化和修复一些问题
1. LCL无边框窗口, 点击任务栏不能切换窗口问题
1. WndProc 回调函数，修改&增加winapi函数，增加 HDWP 类型
1. 增加 lcl 窗口配置函数: 扩展事件-SetOnWndProc, SetOnPaint, 边框Frameless,FramelessForLine, SetRoundRectRgn
1. windows, lcl 窗口调整屏幕缩放比拖拽区域计算位置不正确问题
1. 创建 Application 初始化配置增加默认开启GPU加速
1. 移除生成图标示例修改部分示例

### 2.3.2
1. LCL无边框窗口, 点击任务栏不能切换窗口问题
1. WndProc 回调函数，修改&增加winapi函数，增加 HDWP 类型
1. 增加 lcl 窗口配置函数: 扩展事件-SetOnWndProc, SetOnPaint, 边框Frameless,FramelessForLine, SetRoundRectRgn
1. windows, lcl 窗口调整屏幕缩放比拖拽区域计算位置不正确问题
1. 创建 Application 初始化配置增加默认开启GPU加速
1. 移除生成图标示例修改部分示例

### 2.3.1
1.  增加底层动态库异常捕获, 仅Windows, MacOS
1.  升级CEF从109直接跳到117, 110~116版本的liblcl构建跳过, 此时会增加和移除一些api
1.  升级命令行工具1.0.2,增加兼容Windows7 CEF109
1.  增加底层库windows, macos异常捕获
1.  energy 底层依赖库自动化构建和发布
1.  增加一些示例：屏幕截取，模拟事件，IPC Go to Go。
1.  修复一些已知问题

升级 liblcl v2.3.1


### 2.3.0
1. 主要：
1. 2.3.0 对部分回调函数做出调整，主要增加了 cef.IBrowserWindow 当前窗口参数

1. 增加 静态资源使用本地或内置资源加载, 暂时不能加载视频资源。
1. 本地或内置资源加载，xhr 代理请求配置支持ssl
1. 修复 linux(高版本) gtk3(默认)加载动态库错误问题。
1. 修复 linux gtk3 无法切换英文问题,
1. 优化 Mac开发环境 energy_env=dev > env=dev

其它：
1. 完善命令行工具
1. 优化 install 开发环境全自动安装:
1. 增加 init 应用项目初始化
1. 增加 build 构建&编译应用执行文件
1. 增加 package 制作应用安装包
1. 依赖库升级
1. golcl v1.0.7
1. liblcl v1.0.4

修复其它已知问题

### 2.2.4
1. 增加常用示例
1. 增加窗口焦点
1. 同步govcl库和liblcl库
1. 升级命令行工具
1. 增加 linux arm 架构二进制包

### 2.2.3
1. U: liblclbinres v1.0.2
1. U: demo main-browser-window
1. U: Return value(float32) method
1. A: demo Window IScreen
1. A: Window IScreen
1. A: displayRef proc api
1. Fix: Multiple display window centering issue
1. U: demo scheme
1. U: demo popup-sub-window elliptic
1. U: Optimize custom window drag and drop creation logic
1. Fix: potential problem, proc api return string error
1. Fix: v8value bug, string value error
1. U: demo frameless
1. A: window, full screen model, add common attributes
1. U: demo frameless
1. U: browser window, full screen
1. U: demo window state
1. U: demo frameless, fullscreen
1. U: window state
1. U: chromium context-menu-command callback
1. U: context-menu
1. A: open tab url callback event
1. U: demo liblcl autoupdate , linux lcl widget init
1. U: tempdll README.md
1. U: gen libbin
1. A: energy command line, Set energy framework development environment
1. U: energy command line, support linux select gtk2 or gtk3 framework
1. U: remove CustomWidgetSetFinalization
1. U: demo context-menu
1. U: LCL CloseBrowserWindow RunOnMainThread
1. A: on message paint
1. U: on message struct
1. U: demo test
1. A: TForm WM Message - > NotifyMoveOrResizeStarted
1. A: TForm WM Message
1. Fix: aux-viewsource, linux
1. U: demos, ui use gtk3, IconFS = xxx.png, other IconFS = xxx.ico
1. A: gtk2 support for CEF 106.1.1

### 2.2.2
1. 修改和优化已知问题
1. https://gitee.com/energye/energy/commits/v2.2.2

### 2.2.1
1. U: demo msgbox
1. A: windows demo custom-browser-create
1. U: .gitattributes
1. U: tempdll for mac load liblcl
1. A: demo tempdll
1. A: go build -tags="tempdll" open TempDLL, import liblclbinres.
1. U: window drag switch
1. U: window drag explanatory
1. U: viewsource
1. A: demo drag file
1. R: remove BrowserWindow.Config.EnableWebkitAppRegion
1. A: demo custom-drag-window
1. U: lcl customer drag
1. U: demo frameless for mac, no hide caption
1. U: lcl window close for mac, 1.show,2.hide
1. U: demo
1. U: remove const.IsMessageLoop, add application.IsMessageLoop()
1. U: message const
1. Merge branch 'main' into dev
1. U: win32 consts, cefwinapi
1. Update README.zh_CN.md
1. Update README.md
1. A: lcl custom window drag
1. A: FramelessForDefault
1. A: CEF FileDialog
1. Fix: Chromium & DialogHandler onFileDialog callback args:acceptFiltersList
1. A: 4 types systray demo
1. A: windows frameless Control Window Border
1. U: demo-windows transparent

### 2.2.0-beta
1. 增加 browser RunFileDialog 回调函数, 该功能是使用CEF打开选择文件弹窗
1. 增加 DownloadImage 回调函数
1. 增加 CEF的 MenuModel、View、Button、LabelButton、MenuButton、Panel、textfield 组件相关 Api
1. 重命名示例目录名
1. 增加 其它功能处理回调函数
1. 修改底层实现类变更接口
1. v2.2.0开始兼容老版本CEF 87 ~ 89, 新版本最低兼容到109
1. 优化energy框架中的window相关事件