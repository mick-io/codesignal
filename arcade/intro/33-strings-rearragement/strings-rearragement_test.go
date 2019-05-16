package codesingal

import "testing"

func Test_stringsRearrangement(t *testing.T) {
	tests := []struct {
		in     []string
		expect bool
	}{
		{[]string{"aba", "bbb", "bab"}, false},
		{[]string{"ab", "bb", "aa"}, true},
		{[]string{"q", "q"}, false},
		{[]string{"zzzzab", "zzzzbb", "zzzzaa"}, true},
		{[]string{"ab", "ad", "ef", "eg"}, false},
		{[]string{"abc", "bef", "bcc", "bec", "bbc", "bdc"}, true},
		{[]string{"abc", "abx", "axx", "abc"}, false},
		{[]string{"abc", "abx", "axx", "abx", "abc"}, true},
		{[]string{"f", "g", "a", "h"}, true},
		{[]string{"ff", "gf", "af", "ar", "hf"}, true},
	}
	for _, test := range tests {
		original := make([]string, len(test.in))
		copy(original, test.in)
		out := stringsRearrangement(test.in)
		if out != test.expect {
			msg := "\n[FAIL] stringsRearrangement(%v) == %t expect %t"
			t.Errorf(msg, original, out, test.expect)
		}
	}
}
