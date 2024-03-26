//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/v2/api"

// GetDEREncodedIssuerChain
//
//	Returns the DER encoded data for the certificate issuer chain. If we
//	failed to encode a certificate in the chain it is still present in the
//	array but is an NULL string.
func (m *TCefX509Certificate) GetDEREncodedIssuerChain(chainCount *NativeUInt, chain *ICefBinaryValueArray) {
	var result1 uintptr
	api.CEFPreDef().SysCallN(5, m.Instance(), uintptr(unsafePointer(chainCount)), uintptr(unsafePointer(&result1)))
	if result1 > 0 {
		*chain = BinaryValueArrayRef.New(int(*chainCount), result1)
	}
}

// GetPEMEncodedIssuerChain
//
//	Returns the PEM encoded data for the certificate issuer chain. If we
//	failed to encode a certificate in the chain it is still present in the
//	array but is an NULL string.
func (m *TCefX509Certificate) GetPEMEncodedIssuerChain(chainCount *NativeUInt, chain *ICefBinaryValueArray) {
	var result1 uintptr
	api.CEFPreDef().SysCallN(6, m.Instance(), uintptr(unsafePointer(chainCount)), uintptr(unsafePointer(&result1)))
	if result1 > 0 {
		*chain = BinaryValueArrayRef.New(int(*chainCount), result1)
	}
}
