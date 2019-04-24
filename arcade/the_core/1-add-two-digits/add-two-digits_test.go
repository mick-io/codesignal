package codesignal

import "testing"

func Test_addTwoDigits(t *testing.T) {
	tests := []struct{ in, expect int }{
		{29, 11},
		{48, 12},
		{10, 1},
		{25, 7},
		{52, 7},
		{99, 18},
		{44, 8},
		{50, 5},
		{39, 12},
		{26, 8},
	}
	for _, test := range tests {
		out := addTwoDigits(test.in)
		if out != test.expect {
			msg := "[FAIL] addTwoDigits(%d) == %d, expect %d"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
