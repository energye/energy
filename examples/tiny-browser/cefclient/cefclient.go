package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/examples/tiny-browser/cefclient/assserv"
	. "github.com/energye/energy/v2/examples/tiny-browser/cefclient/browse"
	"github.com/energye/golcl/lcl/api"
	"os"
	"path/filepath"
)

var (
	app *cef.TCEFApplication
)

//go:embed assets
var assets embed.FS

func main() {
	Assets = assets
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	app = Application()
	app.SetFrameworkDirPath(os.Getenv("ENERGY_HOME"))
	app.SetMultiThreadedMessageLoop(false)
	app.SetExternalMessagePump(false)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetLocale(consts.LANGUAGE_zh_CN)
	app.SetTouchEvents(consts.STATE_ENABLED)
	//lp := common.FrameworkDir()
	//if lp != "" {
	//	app.SetFrameworkDirPath(lp)
	//}
	//app.SetDisableZygote(true)
	if common.IsDarwin() {
		app.AddCrDelegate()
		cef.GlobalWorkSchedulerCreate(nil)
		app.SetOnScheduleMessagePumpWork(nil)
	}
	//fmt.Println("libname.LibName:", libname.LibName)
	fmt.Println("WidgetUI:", api.WidgetUI(), "ChromeVersion:", app.ChromeVersion(), "LibCefVersion:", app.LibCefVersion())

	kPrefWindowRestore := "cefclient.window_restore"
	app.SetOnRegisterCustomPreferences(func(type_ consts.TCefPreferencesType, registrar *cef.TCefPreferenceRegistrarRef) {
		fmt.Println("OnRegisterCustomPreferences ProcessType:", process.Args.ProcessType())
		if type_ == consts.CEF_PREFERENCES_TYPE_GLOBAL {
			dict := cef.DictionaryValueRef.New()
			dict.SetInt(kPrefWindowRestore, int32(consts.CEF_SHOW_STATE_NORMAL))
			value := cef.ValueRef.New()
			value.SetDictionary(dict)
			registrar.AddPreference(kPrefWindowRestore, value)
		}
	})
	app.SetOnAlreadyRunningAppRelaunch(func(commandLine *cef.ICefCommandLine, currentDirectory string) bool {
		fmt.Println("OnAlreadyRunningAppRelaunch ProcessType:", process.Args.ProcessType())
		// 在此处创建一个新窗口
		// 重新启动处理好了
		return true
	})
	app.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized ProcessType:", process.Args.ProcessType())
		fmt.Println("  GetScreenDPI:", cef.GetScreenDPI(), "GetDeviceScaleFactor:", cef.GetDeviceScaleFactor())
	})
	if common.IsDarwin() && !process.Args.IsMain() {
		startSub := app.StartSubProcess()
		fmt.Println("start sub:", startSub)
		app.Free()
	} else {
		startMain := app.StartMainProcess()
		fmt.Println("start main:", startMain)
		if startMain {
			assserv.StartServer()
			// 创建窗口
			MainWindow()
		}
	}
}
