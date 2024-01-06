
### 2.4.0

#### Windows XP 版本支持

```text
go1.10 是最后一个支持WinXP的, 但是编译出的exe对CEF封装的不好
go1.11.4 和1.11.13 编译出的exe可以在WinXP SP3 运行, 测试赞未发现问题
windows xp go 1.11.13 

以下修改liblcl需同步到其它特定分枝和主分枝

增加了特定版本支持 API
liblcl
    CEFAppConfig_SpecificVersion  OK

修改字符串返回, 使用TString API  OK
liblcl
    CEFFrame_Name
    CEFFrame_Url

侯改 ChromiumEvent_OnBeforePopup  OK
Go
    SetOnBeforePopup 增加 settings *TCefBrowserSettings
liblcl
    beforePopupInfo 和 browserSettings 指针改为结构
    同步修改
    CefBrowserSettingsToGoBrowserSettings
    GoBrowserSettingsToCefBrowserSettings

    增加参数 TCefWindowInfo 
    增加参数 TCefPopupFeatures
  
    修复-增加: Chromium OnExtension XXXX 相关函数没返回 Sender
    其它带有 TCefBrowserSettings 回调函数, uCEF_LCL_BrowserViewDelegateRef
        OnGetDelegateForPopupBrowserView

    ICefLifeSpanHandler onBeforePopup



XP系统支持Go energy不支持go mod和embed.FS   ok
embed.FS: 使用第三方编译 go-bindata https://zhuanlan.zhihu.com/p/458008381
golcl 增加 支持 go1.11的打开文件接口 emfs


Go
  移除 liblclbinres 在构建时生成二进制liblcl.go , 将go-bindata集成到energy命令行工具中
  命令行工具 
     集成 go-bindata
     build 增加自定义扩展参数  

移除 liblclbinres, 使用libs编译时的内置动态链接库

Go 
  修改了本地资源加载顺序，需要验证Linux MacOS
  
  
liblcl
    增加
    CEF_LCL_WindowDelegate.inc      {$I CEF_LCL_WindowDelegate.inc}
    uCEF_LCL_WindowDelegateRef.pas  uCEF_LCL_WindowDelegateRef
```