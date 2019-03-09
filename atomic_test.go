package atomic_test

import "testing"
import va "github.com/vsrmars/atomic"

var v va.Value

func TestValue(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
