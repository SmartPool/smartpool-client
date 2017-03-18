package mtree

import (
	"github.com/SmartPool/smartpool-client"
)

func conventionalWord(data smartpool.Word) ([]byte, []byte) {
	first := rev(data[:32])
	first = append(first, rev(data[32:64])...)
	second := rev(data[64:96])
	second = append(second, rev(data[96:128])...)
	return first, second
}

func rev(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func msbPadding(a []byte, size uint32) []byte {
	result := make([]byte, len(a))
	copy(result, a)
	for i := uint32(len(a)); i < size; i++ {
		result = append([]byte{0}, result...)
	}
	return result
}
