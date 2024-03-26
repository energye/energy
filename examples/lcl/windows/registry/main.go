package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/win"
)

// 注册表操作示例

func main() {
	inits.Init(nil, nil)
	// 64位下传入KEY_WOW64_64KEY
	//reg := lcl.NewRegistry(win.KEY_ALL_ACCESS|win.KEY_WOW64_64KEY)
	reg := lcl.NewRegistryAllAccess()
	defer reg.Free()
	reg.SetRootKey(win.HKEY_CURRENT_USER)
	if reg.OpenKeyReadOnly("SOFTWARE\\Google\\Chrome\\BLBeacon") {
		defer reg.CloseKey()
		fmt.Println("version:", reg.ReadString("version"))
		fmt.Println("state:", reg.ReadInteger("state"))
		fmt.Println("BLBeacon Exists:", reg.KeyExists("BLBeacon"))
		fmt.Println("failed_count Exists:", reg.ValueExists("failed_count"))
		//
		// reg.WriteBool()
	} else {
		fmt.Println("打开失败！")
	}
}
