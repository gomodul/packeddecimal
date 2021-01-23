package packeddecimal_test

import (
	"testing"

	"github.com/gomodul/packeddecimal"
)

func TestDecode(t *testing.T) {
	const val = -1234544

	enc, err := packeddecimal.Encode(val, 10)
	if err != nil {
		t.Fatal(err)
	}

	res, err := packeddecimal.Decode(enc)
	if err != nil {
		t.Fatal(err)
	}

	if val != res {
		t.Fatal("unexpected", val, res)
	}
}
