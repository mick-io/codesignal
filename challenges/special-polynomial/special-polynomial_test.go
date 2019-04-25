package codesignal

import "testing"

func Test_specialPolynomial(t *testing.T) {
	tests := []struct{ x, n, expect int }{
		{2, 5, 1},
		{10, 111111110, 7},
		{1, 100, 99},
		{3, 140, 4},
	}
	for _, test := range tests {
		out := specialPolynomial(test.x, test.n)
		if out != test.expect {
			msg := "\n[FAIL] specialPolynomial(%d, %d) == %d expect %d"
			t.Errorf(msg, test.x, test.n, out, test.expect)
		}
	}
}
