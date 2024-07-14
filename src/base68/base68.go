package base68

import (
	"bytes"
	"fmt"
)

var (
	base68Chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_~$!")
)

func IncrementBase68String(indexString string) string {
	index := []byte(indexString)
	if len(index) == 0 {
		return string([]byte{base68Chars[0]})
	}

	indices := make([]int, len(index))
	for i, char := range index {
		indices[i] = bytes.IndexByte(base68Chars, char)
	}

	for i := len(indices) - 1; i >= 0; i-- {

		newVal := indices[i] + 1

		indices[i] = newVal % len(base68Chars)
		fmt.Printf("%v, %v\n", newVal, len(base68Chars))
		if newVal < len(base68Chars) {
			break
		}

		if i == 0 {
			indices = append([]int{0}, indices...)
		}

	}

	result := make([]byte, len(indices))
	for i, idx := range indices {
		result[i] = base68Chars[idx]
	}
	return string(result)
}
