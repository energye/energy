//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package consts

func (m *TCefProcessTypeValue) IsBrowser() bool {
	return *m == PtvBrowser
}

func (m *TCefProcessTypeValue) IsRenderer() bool {
	return *m == PtvRenderer
}

func (m *TCefProcessTypeValue) IsZygote() bool {
	return *m == PtvZygote
}

func (m *TCefProcessTypeValue) IsGPU() bool {
	return *m == PtvGPU
}

func (m *TCefProcessTypeValue) IsUtility() bool {
	return *m == PtvUtility
}

func (m *TCefProcessTypeValue) IsBroker() bool {
	return *m == PtvBroker
}

func (m *TCefProcessTypeValue) IsCrashpad() bool {
	return *m == PtvCrashpad
}

func (m *TCefProcessTypeValue) IsOther() bool {
	return *m == PtvOther
}

func (m *TCefProcessType) IsPtBrowser() bool {
	return *m == PtBrowser
}

func (m *TCefProcessType) IsPtRenderer() bool {
	return *m == PtRenderer
}

func (m *TCefProcessType) IsPtZygote() bool {
	return *m == PtZygote
}

func (m *TCefProcessType) IsPtGPU() bool {
	return *m == PtGPU
}

func (m *TCefProcessType) IsPtUtility() bool {
	return *m == PtUtility
}

func (m *TCefProcessType) IsPtBroker() bool {
	return *m == PtBroker
}

func (m *TCefProcessType) IsPtCrashpad() bool {
	return *m == PtCrashpad
}

func (m *TCefProcessType) IsPtOther() bool {
	return *m == PtOther
}

func (m *TCefReturnValue) Cancel() {
	*m = RV_CANCEL
}

func (m *TCefReturnValue) Continue() {
	*m = RV_CONTINUE
}

func (m *TCefReturnValue) ContinueAsync() {
	*m = RV_CONTINUE_ASYNC
}
