package packeddecimal

import (
	"strconv"
	"strings"
)

// ByteToString func(b []byte) string.
func ByteToString(b []byte) string {
	var res string
	for i := 0; i < len(b); i++ {
		currByte := b[i] & 0xFF

		res += "0x"
		// if currByte <= 15
		if currByte <= 0x0F {
			res += "0"
		}

		res += strconv.FormatInt(int64(currByte), 16) + " "
	}
	return strings.TrimSpace(res)
}
