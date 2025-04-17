//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 将资源以二进制形式生成go资源文件

package internal

import (
	"github.com/cyber-xxm/energy/v2/cmd/internal/bindata"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
)

var CmdBindata = &command.Command{
	UsageLine: "bindata",
	Short:     "Use bindata to embed static resources",
	Long: `
	If the go version is less than 1.16, you can use bindata to embed static resources
Example golang code:
	package main  

	import (  
		"fmt"  
		// other imports...  
	)  
	  
	//go:generate energy bindata --fs --o=assets/assets.go --pkg=assets --paths=./assets
	  
	func main() {  
		// your code here...  
	}
Bash: Run the following command in the same directory
	go generate
`,
}

func init() {
	CmdBindata.Run = runBindata
}

func runBindata(c *command.Config) error {
	return bindata.Bindata(c)
}
