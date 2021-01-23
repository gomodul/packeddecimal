package packeddecimal

import (
	"errors"
	"fmt"
	"strconv"
)

// Encode func(n int, size int) ([]byte, error).
func Encode(n int, size int) ([]byte, error) {
	res := make([]byte, size)

	s := fmt.Sprint(n)
	length := len(s)

	isNegative := n < 0
	if isNegative {
		s = s[1:]
		length--
	}

	if length > 2*size-1 {
		return res, errors.New("index out of range")
	}

	if length%2 == 0 {
		s = "0" + s
		length++
	}

	neededBytes := ((length + 1) / 2) + (1 / 2)
	extraBytes := neededBytes - size

	if extraBytes < 0 {
		for i := 0; i < -extraBytes; i++ {
			res[i] = 0x00
		}
	} else if extraBytes > 0 {
		s = s[extraBytes:]
		length -= extraBytes
		extraBytes = 0
	}

	var pos int
	for i := 0; i < length; i++ {
		digit, _ := strconv.Atoi(s[i : i+1])
		pos = (i / 2) - extraBytes

		if i%2 == 0 {
			res[pos] = byte(digit << 4)
		} else {
			res[pos] = res[pos] | (byte(digit) & 0x0f)
		}
	}

	if isNegative {
		res[size-1] |= minusSign
	} else {
		res[size-1] |= plusSign
	}

	return res, nil
}
