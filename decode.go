package packeddecimal

import (
	"errors"
)

// Decode func(encoded []byte) (int, error).
func Decode(encoded []byte) (int, error) {
	var res = 0
	var length = len(encoded)

	for i := 0; i < length; i++ {
		currByte := encoded[i] & dropH0Sign

		digit := int(currByte >> 4)
		res = res*10 + digit

		if i == length-1 {
			sign := currByte & l0Sign
			if sign == minusSign {
				res = -res
			} else if sign != plusSign {
				return res, errors.New("invalid value")
			}
		} else {
			digit = int(currByte & l0Sign)
			res = res*10 + digit
		}
	}

	return res, nil
}
