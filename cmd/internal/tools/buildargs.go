//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package tools

import "os"

// 返回 buildargs 参数索引 + 1
// 在 build 时，如果设置 go 的构建参数, 需要设置 --buildargs 标记，并且让其在 cli 命令最一个有效参数位置
// 其之后参数都将做为 go build [args] 传递
func GetBuildArgsFlagIndex() int {
	for i, arg := range os.Args {
		if arg == "--buildargs" {
			return i + 1
		}
	}
	return 0
}
