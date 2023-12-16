package common

import "embed"

//go:embed resources
var resources embed.FS

// ResourcesFS Static resource directory used by all examples
func ResourcesFS() *embed.FS {
	return &resources
}
