package cef

// Instance 实例
func (m *ICefBinaryValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}
