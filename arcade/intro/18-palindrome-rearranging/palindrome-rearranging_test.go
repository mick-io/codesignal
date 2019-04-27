package codesignal

import "testing"

func Test_palindromeRearranging(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{in: "aabb", expect: true},
		{in: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc", expect: false},
		{in: "abbcabb", expect: true},
		{in: "zyyzzzzz", expect: true},
		{in: "z", expect: true},
		{in: "zaa", expect: true},
		{in: "abca", expect: false},
		{in: "abcad", expect: false},
		{in: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbccccaaaaaaaaaaaaa", expect: false},
		{in: "abdhuierf", expect: false},
	}
	for _, test := range tests {
		out := palindromeRearranging(test.in)
		if out != test.expect {
			msg := "\n[FAIL] palindromeRearranging(%s) == %t expect %t"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
