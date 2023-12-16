package common

import "embed"

//go:embed resources
var resources embed.FS

func Resources() embed.FS {
	return resources
}
