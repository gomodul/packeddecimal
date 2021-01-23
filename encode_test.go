package packeddecimal_test

import (
	"testing"

	"github.com/gomodul/packeddecimal"
)

func TestEncode(t *testing.T) {
	_, err := packeddecimal.Encode(-11011921, 5)
	if err != nil {
		t.Fatal(err)
	}
}
