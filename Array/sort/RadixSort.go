package sort

import (
	"unsafe"
)

// 基数排序，不依赖比较操作，具有稳定性
// 复杂度为 O((w/m)N) & O(N+2^m)
func RadixSort(list []Unit) {
	const base = -int((^uint(0))>>1) - 1
	size := len(list)
	for i := 0; i < size; i++ {
		list[i].val += base
	}

	shadow := make([]Unit, size)
	book := new([256]uint)

	const UINT_LEN = uint(unsafe.Sizeof(uint(0))) * 8
	for step := uint(0); step < UINT_LEN; step += 8 {
		for i := 0; i < 256; i++ {
			book[i] = 0
		}
		for i := 0; i < size; i++ {
			radix := uint8((list[i].val >> step) & 0xFF)
			book[radix]++
		}
		line := uint(0)
		for i := 0; i < 256; i++ {
			book[i], line = line, line+book[i]
		}
		for i := 0; i < size; i++ {
			radix := uint8((list[i].val >> step) & 0xFF)
			shadow[book[radix]] = list[i]
			book[radix]++
		}
		list, shadow = shadow, list
	}

	//if UINT_LEN%16 != 0 {
	//	copy(list, shadow)
	//}
	for i := 0; i < size; i++ {
		list[i].val -= base
	}
}
