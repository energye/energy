package install

import "testing"

func TestVersion(t *testing.T) {
	majorVersion, minorVersion, buildNumber := versionNumber()
	println(majorVersion, minorVersion, buildNumber)
}
