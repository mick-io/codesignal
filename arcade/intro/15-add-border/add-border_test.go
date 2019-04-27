package codesignal

import "testing"

func Test_addBorder(t *testing.T) {
	tests := []struct{ in, expect []string }{
		{
			in:     []string{"abc", "ded"},
			expect: []string{"*****", "*abc*", "*ded*", "*****"},
		},
		{
			in:     []string{"a"},
			expect: []string{"***", "*a*", "***"},
		},
		{
			in:     []string{"aa", "**", "zz"},
			expect: []string{"****", "*aa*", "****", "*zz*", "****"},
		},
		{
			in:     []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy"},
			expect: []string{"*******", "*abcde*", "*fghij*", "*klmno*", "*pqrst*", "*uvwxy*", "*******"},
		},
		{
			in:     []string{"wzy**"},
			expect: []string{"*******", "*wzy***", "*******"},
		},
	}
	slicesMatch := func(s1, s2 []string) bool {
		for i, s := range s1 {
			if s2[i] != s {
				return false
			}
		}
		return true
	}
	for _, test := range tests {
		out := addBorder(test.in)
		if !slicesMatch(test.expect, out) {
			msg := "\n[FAIL] addBorder(%v) == %v expect %v"
			t.Errorf(msg, test.in, out, test.expect)
		}
	}
}
