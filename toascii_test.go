package packeddecimal_test

import (
	"testing"

	"github.com/gomodul/packeddecimal"
)

func TestToASCII(t *testing.T) {
	enc, err := packeddecimal.Encode(10091993, 5)
	if err != nil {
		t.Fatal(err)
	}
	packeddecimal.ToASCII(enc)
}
