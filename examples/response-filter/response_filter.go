package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unsafe"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	var app = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/response-filter.html"
	cef.BrowserWindow.Config.Title = "Energy - response-filter"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 选中的radio按钮替换类型
		var check = 1
		// 监听刷新
		ipc.On("refresh", func() {
			fmt.Println("refresh-check:", check)
			window.Chromium().ReloadIgnoreCache()
		})
		// 监听 radio 按钮的替换类型
		// 1 html 文本替换
		// 2 icon 图片替换
		ipc.On("replace", func(t int) {
			check = t
			fmt.Println("replace-check:", check)
		})
		// 页面加载完成后 发送消息, 页面 radio 选中
		window.Chromium().SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			// 控制页面中默认选中的按钮
			ipc.Emit("replace", check)
		})

		// 要替换的图片
		wd := consts.CurrentExecuteDir
		fmt.Println("exePath", wd)
		replaceImageBuf, err := ioutil.ReadFile(filepath.Join(wd, "examples", "response-filter", "resources", "jupiter.png"))
		fmt.Println("err:", err)
		// 当前读取位置 和 图片总大小
		var position, size = 0, len(replaceImageBuf)
		// 页面开始加载时 重置读取位置
		window.Chromium().SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
			position = 0 // 重置位置 0
		})
		// 创建响应过滤, 在OnGetResourceResponseFilter函数中根据要过滤的url
		var createFilter = func() *cef.ICefResponseFilter {
			var replaceString = "在GO中使用JavaScript、HTML和CSS构建跨平台的桌面应用程序"
			var newReplaceString = "Building cross platform desktop applications using JavaScript, HTML and CSS in go"
			var filter = cef.ResponseFilterRef.New()
			// 初始化过滤器 默认返回 false
			filter.InitFilter(func() bool {
				fmt.Println("SetOnGetResourceResponseFilter-InitFilter")
				return true // 返回true使过滤器生效
			})
			// 过滤器
			filter.Filter(func(dataIn uintptr, dataInSize uint32, dataInRead *uint32, dataOut uintptr, dataOutSize uint32, dataOutWritten *uint32) (status consts.TCefResponseFilterStatus) {
				status = consts.RESPONSE_FILTER_DONE
				fmt.Println("SetOnGetResourceResponseFilter-Filter dataIn:", dataIn, "dataInSize:", dataInSize, "*dataInRead:", *dataInRead, "dataOut:", dataOut, "dataOutSize:", dataOutSize, "*dataOutWritten", *dataOutWritten)
				// status 默认值 RESPONSE_FILTER_DONE, 停止过滤
				// dataIn, 输入流内容, 内容指针地址, 为0时表示nil, 可以读取内容
				// dataOut, 输出流, 如果过滤器生效, 我们需要将内容通过该参数返回, 为0时表示nil,
				// 这个函数 可能会被多次调用
				if dataIn != 0 && check == 1 { // 替换文本
					// 读取输入流内容
					// 创建一个 dataInSize 大小的缓冲
					var contentBuf = make([]byte, dataInSize)
					var i uint32 = 0 // 使用 i 表示下一个字节指针
					// 循环读取
					for i < dataInSize {
						contentBuf[i] = *(*byte)(unsafe.Pointer(dataIn + uintptr(i)))
						i = i + 1 // 下一个字节指针
					}
					// 读取内容需要把当前 读取的大小 dataInSize 设置到 dataInRead
					*dataInRead = dataInSize
					//fmt.Println("content:", string(contentBuf))
					// 这里可以替换内容, 然后将替换后的内容返回到 dataOut, 通过指针操作
					contentBuf = []byte(strings.Replace(string(contentBuf), replaceString, newReplaceString, -1))
					// 将输入流设置到 dataOut
					i = 0 // 默认 0
					for i < dataInSize {
						*(*byte)(unsafe.Pointer(dataOut + uintptr(i))) = contentBuf[i]
						i = i + 1 // 下一个字节指针
					}
					*dataOutWritten = i
				} else if dataOut != 0 && check == 2 { // 替换图片
					status = consts.RESPONSE_FILTER_DONE
					*dataInRead = dataInSize
					var i uint32 = 0 // 默认 0
					// 循环把要替换的图片内容设置到输出上
					for i < dataOutSize && position < size {
						// 这里是通过指针地址将赋值, 图片二进制缓存数组
						// 描述: dataOut byte[i] = replaceImageBuf byte[position]
						*(*byte)(unsafe.Pointer(dataOut + uintptr(i))) = replaceImageBuf[position]
						position++ // 图片缓存数据的下一个位置
						i = i + 1  // 计数, 当前最后的读取&输出大小
					}
					*dataOutWritten = i
					if position < size {
						// 如果我们还有更多的内容没返回完, 在这里必须返回 RESPONSE_FILTER_NEED_MORE_DATA
						status = consts.RESPONSE_FILTER_NEED_MORE_DATA
					}
				}
				return
			})
			return filter
		}

		window.Chromium().SetOnGetResourceResponseFilter(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, response *cef.ICefResponse) (responseFilter *cef.ICefResponseFilter) {
			if strings.Index(request.URL(), "response-filter.html") != -1 && check == 1 {
				return createFilter()
			} else if strings.Index(request.URL(), "icon.png") != -1 && check == 2 {
				return createFilter()
			}
			return nil
		})
	})
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
