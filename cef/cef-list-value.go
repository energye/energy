package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
)

func (m *ICefListValue) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefListValue) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefListValue_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}
