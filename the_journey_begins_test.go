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

func Test_centuryFromYear(t *testing.T) {
	tests := []struct{ input, expect int }{
		{1905, 20},
		{1700, 17},
		{200, 2},
		{45, 1},
	}
	for _, test := range tests {
		output := centuryFromYear(test.input)
		if output != test.expect {
			msg := "\n[FAIL] centuryFromYear(%d) == %d expect %d"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}

func Test_checkPalindrome(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"apple", false},
		{"ada", true},
		{"rotator", true},
	}
	for _, test := range tests {
		output := checkPalindrome(test.input)
		if output != test.expect {
			msg := "\n[FAIL] checkPalindrome(%q) == %t expect %t"
			t.Errorf(msg, test.input, output, test.expect)
		}
	}
}
