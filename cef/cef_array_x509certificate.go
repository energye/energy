//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICefX509CertificateArray
//
//	[]ICefX509Certificate
type ICefX509CertificateArray interface {
	Instance() uintptr
	Get(index int) ICefX509Certificate
	Size() int
	Free()
	Add(value ICefX509Certificate)
	Set(value []ICefX509Certificate)
}

// TCefX509CertificateArray
//
//	[]ICefX509Certificate
type TCefX509CertificateArray struct {
	instance unsafePointer
	count    int
	values   []ICefX509Certificate
}

// X509CertificateArrayRef -> TCefX509CertificateArray
var X509CertificateArrayRef x509CertificateArray

// x509CertificateArray
type x509CertificateArray uintptr

func (*x509CertificateArray) New(count int, instance uintptr) ICefX509CertificateArray {
	return &TCefX509CertificateArray{
		count:    count,
		instance: unsafePointer(instance),
		values:   make([]ICefX509Certificate, count),
	}
}

func (m *TCefX509CertificateArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// Get 根据下标获取 ICefX509Certificate
func (m *TCefX509CertificateArray) Get(index int) ICefX509Certificate {
	if m == nil {
		return nil
	}
	if index < m.count {
		result := m.values[index]
		if result == nil {
			result = AsCefX509Certificate(getParamOf(index, m.Instance()))
			m.values[index] = result
		}
		return result
	}
	return nil
}

// Size 返回 ICefX509Certificate 数组长度
func (m *TCefX509CertificateArray) Size() int {
	if m == nil {
		return 0
	}
	return m.count
}

func (m *TCefX509CertificateArray) Free() {
	if m == nil {
		return
	}
	if m.values != nil {
		for i, v := range m.values {
			if v != nil && v.Instance() != 0 {
				v.Free()
				m.values[i] = nil
			}
		}
		m.values = nil
	}
	m.instance = nil
	m.count = 0
}

func (m *TCefX509CertificateArray) Add(value ICefX509Certificate) {
	m.values = append(m.values, value)
	m.count++
	m.instance = unsafePointer(m.values[0].Instance())
}

func (m *TCefX509CertificateArray) Set(value []ICefX509Certificate) {
	if m.values != nil {
		for i, v := range m.values {
			if v != nil && v.Instance() != 0 {
				v.Free()
				m.values[i] = nil
			}
		}
	}
	m.values = value
	m.count = len(value)
	m.instance = unsafePointer(m.values[0].Instance())
}
