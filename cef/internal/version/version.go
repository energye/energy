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
	version         = "2.0.1-beta" // energy
	libVersion      = "2.0.1"      // lib-lcl
	cefVersion      = "1.109.18"   // cef framework
	libBuildVersion = ""           // pascal lib build
)

// Version return energy version
func Version() string {
	return version
}

// LibVersion return lib-lcl version
func LibVersion() string {
	return libVersion
}

// CEFVersion return cef framework version
func CEFVersion() string {
	return cefVersion
}

// LibBuildVersion return pascal lib build version
func LibBuildVersion() string {
	return libBuildVersion
}

// SetLibBuildVersion Set pascal lib build version
func SetLibBuildVersion(version string) {
	libBuildVersion = version
}
