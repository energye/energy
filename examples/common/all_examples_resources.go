package common

import (
	"embed"
	"github.com/energye/golcl/energy/emfs"
)

//go:embed resources
var resources embed.FS

// ResourcesFS Static resource directory used by all examples
func ResourcesFS() emfs.IEmbedFS {
	return resources
}

// Go版本小于1.16时使用
////go:generate energy bindata --fs --o=assets/assets.go --pkg=assets --paths=./resources/...
//func ResourcesFS() emfs.IEmbedFS {
//	return assets.AssetFile()
//}
