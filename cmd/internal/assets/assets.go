//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package assets

import "embed"

//go:embed assets
var assets embed.FS

func ReadFile(name string) ([]byte, error) {
	return assets.ReadFile(name)
}
