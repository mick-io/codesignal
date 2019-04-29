package codesignal

import "testing"

func Test_deleteWhitespace(t *testing.T) {
	tests := []struct{ in, expect string }{
		{"a  b cd  e", "abcde"},
		{"a   b c  dc ", "abcdc"},
	}
	for _, test := range tests {
		out := deleteWhitespaces(test.in)
		if out != test.expect {
			msg := "\n[FAIL] deleteWhitespaces(%s) == %s expect %s"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
