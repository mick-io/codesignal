package codesignal

import "testing"

func Test_longestDigitPrefix(t *testing.T) {
	tests := []struct{ in, expect string }{
		{"123aa1", "123"},
		{"0123456789", "0123456789"},
		{"  3) always check for whitespace", ""},
		{"12abc34", "12"},
		{"the output is 42", ""},
		{"aaaaaaa", ""},
	}
	for _, test := range tests {
		out := longestDigitsPrefix(test.in)
		if out != test.expect {
			msg := "\n[FAIL] longestDigitsPrefix(%s) == %s expect %s"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
