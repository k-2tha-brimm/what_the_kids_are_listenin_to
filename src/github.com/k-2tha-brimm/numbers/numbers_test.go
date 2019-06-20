package numbers

import (
	"testing"
)

func Test(t *testing.T) {
	value := sum(1, 2)
	if value != 3 {
		t.Fatalf("expected 3, got: %d", value)
	}
}
