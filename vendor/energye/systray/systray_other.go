//go:build !darwin
// +build !darwin

package systray

// CreateMenu 如果菜单项是空，把菜单项添加到托盘
// 该法主动调用后 鼠标事件失效
//
// MacOSX平台
func CreateMenu() {
}

// SetMenuNil 如果菜单项不是空，把菜单项设置为null
// 该方法主动调用后 鼠标事件生效
//
// MacOSX平台
func SetMenuNil() {
}
