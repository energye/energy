// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package ipc

import _ "embed"

// JSIPC js
//
//go:embed embed_ipc.js
var JSIPC []byte

// JSDrag js
//
//go:embed embed_drag.js
var JSDrag []byte
