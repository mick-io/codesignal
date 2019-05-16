package codesignal

import "testing"

func Test_digitDegrees(t *testing.T) {
	tests := []struct{ in, expect int }{
		{5, 0},
		{100, 1},
		{91, 2},
		{99, 2},
		{1000000000, 1},
		{9, 0},
		{73, 2},
		{877, 2},
		{777773, 3},
		{304, 1},
	}
	for _, test := range tests {
		out := digitDegree(test.in)
		if out != test.expect {
			msg := "\n[FAIL] digitDegree(%d) == %n expect %d"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
