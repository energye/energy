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

// Current Branch Version
var (
	version         = "2.0.1-beta" // golang energy
	libVersion      = ""           // pascal lib-lcl
	libBuildVersion = ""           // pascal lib-lcl build
)

// Version return energy version
func Version() string {
	return version
}

// LibVersion return lib-lcl version
func LibVersion() string {
	return libVersion
}

// SetLibVersion Set pascal lib-lcl version
func SetLibVersion(version string) {
	if version == "" {
		libVersion = "0.0.0"
	} else {
		libVersion = version
	}
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
