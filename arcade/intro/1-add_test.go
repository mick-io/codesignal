package codesignal

import (
	"testing"
)

func Test_add(t *testing.T) {
	tests := []struct{ i1, i2, expect int }{
		{1, 2, 3},
		{50, 50, 100},
		{-100, 200, 100},
	}
	for _, test := range tests {
		output := add(test.i1, test.i2)
		if output != test.expect {
			msg := "\n[FAIL] add(%d, %d) == %d expect %d"
			t.Errorf(msg, test.i1, test.i2, output, test.expect)
		}
	}
}
