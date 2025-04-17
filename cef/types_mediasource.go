package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// Represents a source from which media can be routed. Instances of this object
// are retrieved via ICefMediaRouter.GetSource. The functions of this
// interface may be called on any browser process thread unless otherwise
// indicated.
// <para><see cref="uCEFTypes|TCefMediaSource">Implements TCefMediaSource</see></para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_source_t)</see></para>
type ICefMediaSource struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// MediaSourceRef -> ICefMediaSource
var MediaSourceRef mediaSource

type mediaSource uintptr

func (*mediaSource) UnWrap(delegate *ICefMediaSource) *ICefMediaSource {
	var result uintptr
	imports.Proc(def.MediaSourceRef_UnWrap).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSource{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefMediaSource) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaSource) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaSource) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Returns the ID (media source URN or URL) for this source.
func (m *ICefMediaSource) GetId() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.MediaSource_GetId).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns true (1) if this source outputs its content via Cast.
func (m *ICefMediaSource) IsCastSource() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.MediaSource_IsCastSource).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns true (1) if this source outputs its content via DIAL.
func (m *ICefMediaSource) IsDialSource() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.MediaSource_IsDialSource).Call(m.Instance())
	return api.GoBool(r1)
}
