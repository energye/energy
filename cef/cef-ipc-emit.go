package cef

type IEmitTarget interface {
	GetBrowserId() int32
	GetFrameId() int64
}

//GoEmit相关事件的接收目标
type EmitTarget struct {
	BrowseId int32
	FrameId  int64
}

func (m *EmitTarget) GetBrowserId() int32 {
	return m.BrowseId
}

func (m *EmitTarget) GetFrameId() int64 {
	return m.FrameId
}
