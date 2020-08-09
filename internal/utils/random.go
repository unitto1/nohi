package utils

import (
	"io"
	"log"
)

func RandomBytes(rand io.Reader, b []byte) {
	_, err := io.ReadFull(rand, b)
	if err != nil {
		log.Fatal(err)
	}
}

const minBufferSize = 4
const maxBufferSize = 1024 * 10

func RandInt8(rand io.Reader, max uint8, count int) []uint8 {
	k := int(255/max+1) * count
	if k < minBufferSize {
		k = minBufferSize
	} else if k > maxBufferSize {
		k = maxBufferSize
	}

	b := make([]byte, k)
	RandomBytes(rand, b)

	res := make([]uint8, 0, count)

	for i := 0; len(res) < count; i++ {
		if b[i] < max {
			res = append(res, b[i])
		}

		if i == k-1 {
			RandomBytes(rand, b)

			i = 0
		}
	}

	return res
}
