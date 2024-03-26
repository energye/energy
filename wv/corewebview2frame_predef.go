//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
)

// AddHostObjectToScriptWithOrigins
//
//	Add the provided host object to script running in the iframe with the
//	specified name for the list of the specified origins. The host object
//	will be accessible for this iframe only if the iframe's origin during
//	access matches one of the origins which are passed. The provided origins
//	will be normalized before comparing to the origin of the document.
//	So the scheme name is made lower case, the host will be punycode decoded
//	as appropriate, default port values will be removed, and so on.
//	This means the origin's host may be punycode encoded or not and will match
//	regardless. If list contains malformed origin the call will fail.
//	The method can be called multiple times in a row without calling
//	RemoveHostObjectFromScript for the same object name. It will replace
//	the previous object with the new object and new list of origins.
//	List of origins will be treated as following:
//	1. empty list - call will succeed and object will be added for the iframe
//	but it will not be exposed to any origin;
//	2. list with origins - during access to host object from iframe the
//	origin will be checked that it belongs to this list;
//	3. list with "*" element - host object will be available for iframe for
//	all origins. We suggest not to use this feature without understanding
//	security implications of giving access to host object from from iframes
//	with unknown origins.
//	4. list with "file://" element - host object will be available for iframes
//	loaded via file protocol.
//	Calling this method fails if it is called after the iframe is destroyed.
//	snippet ScenarioAddHostObject.cpp AddHostObjectToScriptWithOrigins
//	For more information about host objects navigate to
//	ICoreWebView2::AddHostObjectToScript.
func (m *TCoreWebView2Frame) AddHostObjectToScriptWithOrigins(name string, object OleVariant, originsCount uint32, origins string) bool {
	r1 := WVPreDef().SysCallN(-1, m.Instance())
	return GoBool(r1)
}
