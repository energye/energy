//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package version Energy framework version config
package version

import "github.com/energye/liblclbinres"

// LibVersion return lib-lcl version
func LibVersion() string {
	return liblclbinres.LibVersion()
}
