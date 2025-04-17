package app

import "github.com/cyber-xxm/energy/v2/cef"

var application *cef.TCEFApplication

// GetApplication
//
//	创建 Application 对象，保持主进程和helper进程配置一样
func GetApplication() *cef.TCEFApplication {
	if application == nil {
		application = cef.NewApplication()
		//.. 在这里统一配置
	}
	return application
}
