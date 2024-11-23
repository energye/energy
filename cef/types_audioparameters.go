//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/v2/consts"

// TCefAudioParameters
// include/internal/cef_types.h (cef_audio_parameters_t)
type TCefAudioParameters struct {
	channelLayout   consts.TCefChannelLayout
	sampleRate      int32
	framesPerBuffer int32
}
