//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

type TEvaluateScriptCallbackEvent func(result string, err string)

type TOpenURLsEvent func(urls []string)
type TUniversalLinkEvent func(link string)
