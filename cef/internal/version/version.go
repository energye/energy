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

// Current Branch Version
var (
	libBuildVersion = "" // pascal lib-lcl build
)

// LibVersion return lib-lcl version
func LibVersion() string {
	return liblclbinres.LibVersion()
}

// LibBuildVersion return pascal lib build version
func LibBuildVersion() string {
	return libBuildVersion
}

// SetLibBuildVersion Set pascal lib build version
func SetLibBuildVersion(version string) {
	if version == "" {
		libBuildVersion = "0.0.0"
	} else {
		libBuildVersion = version
	}
}
