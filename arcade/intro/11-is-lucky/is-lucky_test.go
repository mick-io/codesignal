package codesignal

import "testing"

func Test_isLucky(t *testing.T) {
	tests := []struct {
		in     int
		expect bool
	}{
		{1230, true},
		{239017, false},
		{134008, true},
		{10, false},
		{11, true},
		{1010, true},
		{261534, false},
		{100000, false},
		{999999, true},
		{123321, true},
	}
	for _, test := range tests {
		out := isLucky(test.in)
		if out != test.expect {
			msg := "\n[FAIL] isLucky(%d) == %t expect %t"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
