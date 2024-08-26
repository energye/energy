//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import "github.com/energye/lcl/lcl"

// LocalLoadResource Local or built-in resource loading
type LocalLoadResource struct {
	*LocalLoad
	streams map[string]lcl.IMemoryStream
}

func newLocalLoadResource(ll *LocalLoad) *LocalLoadResource {
	if ll != nil {
		newLocalLoad := *ll
		ret := &LocalLoadResource{
			LocalLoad: &newLocalLoad,
			streams:   make(map[string]lcl.IMemoryStream),
		}
		ret.LocalLoad.defaultInit()
		return ret
	}
	return nil
}
