//go:build !windows

package install

func versionNumber() (majorVersion, minorVersion, buildNumber uint32) {
	return
}
