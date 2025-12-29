//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package application

import "github.com/energye/lcl/lcl"

// LocalLoadResource Local or built-in resource loading
type LocalLoadResource struct {
	*LocalLoad
	streams map[string]lcl.IMemoryStream
}

func NewLocalLoadResource(localLoad *LocalLoad) *LocalLoadResource {
	if localLoad != nil {
		newLocalLoad := *localLoad
		ret := &LocalLoadResource{
			LocalLoad: &newLocalLoad,
			streams:   make(map[string]lcl.IMemoryStream),
		}
		ret.LocalLoad.initDefault()
		return ret

	}
	return nil
}
