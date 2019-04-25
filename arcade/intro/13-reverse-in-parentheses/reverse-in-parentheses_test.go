package codesignal

import "testing"

func Test_reverseInParentheses(t *testing.T) {
	tests := []struct{ in, expect string }{
		{"(bar)", "rab"},
		{"foo(bar)baz", "foorabbaz"},
		{"foo(bar)baz(blim)", "foorabbazmilb"},
		{"foo(bar(baz))blim", "foobazrabblim"},
		{"", ""},
		{"()", ""},
		{"(abc)d(efg)", "cbadgfe"},
	}
	for _, test := range tests {
		out := reverseInParentheses(test.in)
		if out != test.expect {
			msg := "\n[FAIL] reverseInParentheses(%s) == %s expect %s"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
