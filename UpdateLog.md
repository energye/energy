#### Windows XP 版本支持


go1.10 是最后一个支持WinXP的, 但是编译出的exe对CEF封装的不好
go1.11.4 和1.11.13 编译出的exe可以在WinXP SP3 运行, 测试赞未发现问题
windows xp go 1.11.13 

以下修改需同步到其它特定分枝和主分枝

增加了特定版本支持 API
liblcl
    CEFAppConfig_SpecificVersion

修改字符串返回, 使用TString API
liblcl
    CEFFrame_Name
    CEFFrame_Url

侯改 ChromiumEvent_OnBeforePopup
Go
    SetOnBeforePopup 增加 settings *TCefBrowserSettings
liblcl
    beforePopupInfo 和 browserSettings 指针改为结构
    同步修改
    CefBrowserSettingsToGoBrowserSettings
    GoBrowserSettingsToCefBrowserSettings

    增加参数 TCefWindowInfo 
    增加参数 TCefPopupFeatures



XP系统支持Go energy不支持go mod和embed.FS
go mod: 使用 gopath, 手动模块
embed.FS: 使用第三方编译 go-bindata https://zhuanlan.zhihu.com/p/458008381
golcl 增加 支持 go1.10的打开文件接口


Go
  移除 liblclbinres 在构建时生成二进制liblcl.go 