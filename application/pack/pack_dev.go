//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !prod

<<<<<<<< HEAD:cmd/internal/tools/version_other.go
package tools

func VersionNumber() (majorVersion, minorVersion, buildNumber uint32) {
	return
}
========
package pack

const IsDev = true
>>>>>>>> v3-alpha:application/pack/pack_dev.go
