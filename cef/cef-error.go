//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

func cefErrorMessage(code CEF_V8_EXCEPTION) string {
	switch code {
	case CVE_ERROR_OK:
		return empty
	case CVE_ERROR_NOT_FOUND_FIELD:
		return "field not found, or field undefined"
	case CVE_ERROR_NOT_FOUND_FUNC:
		return "function not found, or function undefined"
	case CVE_ERROR_TYPE_NOT_SUPPORTED:
		return "variable type is not supported. Only the variable type is supported [string int double bool null undefined]"
	case CVE_ERROR_TYPE_CANNOT_CHANGE:
		return "field type is common and cannot be changed to array、object、function"
	case CVE_ERROR_TYPE_INVALID:
		return "type is invalid"
	case CVE_ERROR_GET_STRING_FAIL:
		return "failed to get string"
	case CVE_ERROR_GET_INT_FAIL:
		return "fFailed to get int"
	case CVE_ERROR_GET_DOUBLE_FAIL:
		return "failed to get a double"
	case CVE_ERROR_GET_BOOL_FAIL:
		return "failed to get a bool"
	case CVE_ERROR_GET_NULL_FAIL:
		return "failed to get null"
	case CVE_ERROR_GET_UNDEFINED_FAIL:
		return "failed to get undefined"
	case CVE_ERROR_FUNC_INVALID_P_L_9:
		return "function is invalid, of an incorrect type, or has more than 9 arguments"
	case CVE_ERROR_FUNC_IN_PAM:
		return "input parameter type is incorrect. It can only be string int double Boolean"
	case CVE_ERROR_FUNC_OUT_PAM:
		return "parameter type is incorrect and can only be [string int double boolean]"
	case CVE_ERROR_FUNC_GET_IN_PAM_STRING_FAIL:
		return "failed to obtain a string value for the input parameter"
	case CVE_ERROR_FUNC_GET_IN_PAM_INT_FAIL:
		return "failed to obtain an int value for the input parameter"
	case CVE_ERROR_FUNC_GET_IN_PAM_DOUBLE_FAIL:
		return "entry failed to get a value of type double"
	case CVE_ERROR_FUNC_GET_IN_PAM_BOOLEAN_FAIL:
		return "failed to get a value of the Boolean type"
	case CVE_ERROR_FUNC_GET_OUT_PAM_STRING_FAIL:
		return "failed to obtain a string value for the output parameter"
	case CVE_ERROR_FUNC_GET_OUT_PAM_INT_FAIL:
		return "failed to obtain a value of type int"
	case CVE_ERROR_FUNC_GET_OUT_PAM_DOUBLE_FAIL:
		return "output parameter failed to get a value of type double"
	case CVE_ERROR_FUNC_GET_OUT_PAM_BOOLEAN_FAIL:
		return "failed to obtain a value of type Boolean"
	case CVE_ERROR_FUNC_GET_OUT_PAM_CEFERROR_FAIL:
		return "failed to obtain a value for the output parameter"
	case CVE_ERROR_IPC_GET_BIND_FIELD_VALUE_FAIL:
		return "IPC failed to get the binding value. Procedure"
	default:
		return "unknown error at runtime"
	}
}
