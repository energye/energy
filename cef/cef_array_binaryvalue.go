package cef

// ICefBinaryValueArray = array of ICefBinaryValue
type ICefBinaryValueArray interface {
	Instance() uintptr
	Get(index int) ICefBinaryValue
	Count() int
	Free()
}

// TCefBinaryValueArray = array of ICefBinaryValue
type TCefBinaryValueArray struct {
	instance unsafePointer
	count    int
	values   []ICefBinaryValue
}

// BinaryValueArrayRef -> TCefBinaryValueArray
var BinaryValueArrayRef binaryValueArray

// v8ValueArray
type binaryValueArray uintptr

func (*binaryValueArray) New(count int, instance uintptr) ICefBinaryValueArray {
	return &TCefBinaryValueArray{
		count:    count,
		instance: unsafePointer(instance),
		values:   make([]ICefBinaryValue, count),
	}
}
func (m *TCefBinaryValueArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefBinaryValueArray) Get(index int) ICefBinaryValue {
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefBinaryValue(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

func (m *TCefBinaryValueArray) Count() int {
	return m.count
}

func (m *TCefBinaryValueArray) Free() {
	if m.instance != nil {
		m.instance = nil
		m.count = 0
	}
}
