package codesignal

import (
	"testing"
)

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
