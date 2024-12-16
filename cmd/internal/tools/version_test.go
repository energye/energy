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

import "testing"

func TestVersion(t *testing.T) {
	majorVersion, minorVersion, buildNumber := VersionNumber()
	println(majorVersion, minorVersion, buildNumber)
}
