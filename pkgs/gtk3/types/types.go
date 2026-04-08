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

type StyleProviderPriority = uint

const (
	STYLE_PROVIDER_PRIORITY_FALLBACK    StyleProviderPriority = 1
	STYLE_PROVIDER_PRIORITY_THEME       StyleProviderPriority = 200
	STYLE_PROVIDER_PRIORITY_SETTINGS    StyleProviderPriority = 400
	STYLE_PROVIDER_PRIORITY_APPLICATION StyleProviderPriority = 600
	STYLE_PROVIDER_PRIORITY_USER        StyleProviderPriority = 800
)
