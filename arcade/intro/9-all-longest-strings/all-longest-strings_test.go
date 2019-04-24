package codesignal

import "testing"

func Test_allLongestStrings(t *testing.T) {
	tests := []struct{ in, expect []string }{
		{
			in:     []string{"aba", "aa", "ad", "vcd", "aba"},
			expect: []string{"aba", "vcd", "aba"},
		},
		{
			in:     []string{"aa"},
			expect: []string{"aa"},
		},
		{
			in:     []string{"abc", "eeee", "abcd", "dcd"},
			expect: []string{"eeee", "abcd"},
		},
		{
			in:     []string{"a", "abc", "cbd", "zzzzzz", "a", "abcdef", "asasa", "aaaaaa"},
			expect: []string{"zzzzzz", "abcdef", "aaaaaa"},
		},
		{
			in:     []string{"enyky", "benyky", "yely", "varennyky"},
			expect: []string{"varennyky"},
		},
		{
			in:     []string{"abacaba", "abacab", "abac", "xxxxxx"},
			expect: []string{"abacaba"},
		},
		{
			in:     []string{"young", "yooooooung", "hot", "or", "not", "come", "on", "fire", "water", "watermelon"},
			expect: []string{"yooooooung", "watermelon"},
		},
		{
			in:     []string{"onsfnib", "aokbcwthc", "jrfcw"},
			expect: []string{"aokbcwthc"},
		},
		{
			in:     []string{"lbgwyqkry"},
			expect: []string{"lbgwyqkry"},
		},
		{
			in:     []string{"i"},
			expect: []string{"i"},
		},
	}
	for _, test := range tests {
		out := allLongestStrings(test.in)
		if !slicesMatch(out, test.expect) {
			msg := "\n[FAIL] allLongestStrings(%v) == %v expect %v"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}

func slicesMatch(s1, s2 []string) bool {
	for i, str := range s1 {
		if str != s2[i] {
			return false
		}
	}
	return true
}
