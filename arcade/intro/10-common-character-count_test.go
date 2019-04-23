package codesignal

import (
	"testing"
)

func Test_commonCharacterCount(t *testing.T) {
	tests := []struct {
		s1, s2 string
		expect int
	}{
		{s1: "aabcc", s2: "adcaa", expect: 3},
		{s1: "zzzz", s2: "zzzzzzz", expect: 4},
		{s1: "abca", s2: "xyzbac", expect: 3},
		{s1: "a", s2: "b", expect: 0},
		{s1: "a", s2: "aaa", expect: 1},
	}
	for _, test := range tests {
		out := commonCharacterCount(test.s1, test.s2)
		if out != test.expect {
			msg := "\n[FAIL] commonCharacterCount(%q, %q) == %d expect %d"
			t.Errorf(msg, test.s1, test.s2, out, test.expect)
		}
	}
}
