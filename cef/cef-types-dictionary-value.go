package cef

// Instance 实例
func (m *ICefDictionaryValue) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}
