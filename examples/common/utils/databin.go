// 示例工具

package utils

import (
	"unsafe"
)

// ReadData 读取指针数据到 []byte
func ReadData(data uintptr, count uint32) []byte {
	result := make([]byte, count, count)
	// 表示下一个字节指针
	var n uint32 = 0
	for n < count {
		// 取出值
		result[n] = *(*byte)(unsafe.Pointer(data + uintptr(n)))
		// 下一个字节指针
		n = n + 1
	}
	return result
}
