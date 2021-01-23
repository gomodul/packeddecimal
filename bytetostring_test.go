package packeddecimal_test

import (
	"fmt"
	"testing"

	"github.com/gomodul/packeddecimal"
)

func TestByteToString(t *testing.T) {
	const val = 1

	res := packeddecimal.ByteToString([]byte{val})
	if len(res) < 1 || res != fmt.Sprintf("0x%02x", val) {
		t.Fatal("unexpected", res)
	}
}
