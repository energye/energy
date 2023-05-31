package assetserve

import (
	"embed"
	"testing"
)

//go:embed assets
var assets embed.FS

func TestServer(t *testing.T) {
	server := NewAssetsHttpServer()
	server.AssetsFSName = "assets" //必须设置目录名
	server.Assets = &assets
	server.StartHttpServer() //go server.StartHttpServer()
}
