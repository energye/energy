//go:build !windows

package main

const (
	syso = false
)

var sysoBytesAMD64 []byte
var sysoBytes386 []byte
