package codesignal

import "testing"

func Test_symbolsPermutation(t *testing.T) {
	tests := []struct {
		word1, word2 string
		expect       bool
	}{
		{
			word1:  "abc",
			word2:  "cab",
			expect: true,
		},
		{
			word1:  "aaaa",
			word2:  "aaa",
			expect: false,
		},
		{
			word1:  "sutr",
			word2:  "cybk",
			expect: false,
		},
		{
			word1:  "kscsa",
			word2:  "ncwxt",
			expect: false,
		},
		{
			word1:  "imazpsni",
			word2:  "kbyafemd",
			expect: false,
		},
		{
			word1:  "ekufzjmk",
			word2:  "chhmjxmy",
			expect: false,
		},
		{
			word1:  "seha",
			word2:  "zims",
			expect: false,
		},
		{
			word1:  "beicgzwj",
			word2:  "pazofnfl",
			expect: false,
		},
		{
			word1:  "nbimwm",
			word2:  "xwidkg",
			expect: false,
		},
		{
			word1:  "ryqa",
			word2:  "ayrq",
			expect: true,
		},
	}
	for _, test := range tests {
		out := symbolsPermutation(test.word1, test.word2)
		if out != test.expect {
			msg := "\n[FAIL] symbolsPermutation(%s, %s) == %t expect %t"
			t.Errorf(msg, test.word1, test.word2, out, test.expect)
		}
	}
}
