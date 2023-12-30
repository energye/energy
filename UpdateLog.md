#### Windows XP 版本支持

以下修改需同步到其它特定分枝和主分枝

增加了特定版本支持 API
liblcl
    CEFAppConfig_SpecificVersion

修改字符串返回, 使用TString API
liblcl
    CEFFrame_Name
    CEFFrame_Url

侯改 ChromiumEvent_OnBeforePopup
go
    SetOnBeforePopup 增加 settings *TCefBrowserSettings
liblcl
    beforePopupInfo 和 browserSettings 指针改为结构
    同步修改
    CefBrowserSettingsToGoBrowserSettings
    GoBrowserSettingsToCefBrowserSettings

    增加参数 TCefWindowInfo 
    增加参数 TCefPopupFeatures