package codesignal

import (
	"testing"
)

func Test_shapeArea(t *testing.T) {
	tests := []struct{ input, expect int }{
		{2, 5},
		{3, 13},
		{1, 1},
		{5, 41},
		{7000, 97986001},
	}
	for _, test := range tests {
		output := shapeArea(test.input)
		if output != test.expect {
			msg := "\n[FAIL] shapeArea(%d) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}
