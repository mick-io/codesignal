package codesignal

import "math"

// Write a function that returns the sum of two numbers.

func add(n1 int, n2 int) int {
	return n1 + n2
}

// Given a year, return the century it is in. The first century spans from the
// year 1 up to and including the year 100, the second - from the year 101 up
// to and including the year 200, etc.

func centuryFromYear(year int) int {
	return int(math.Ceil(float64(year) / 100))
}

// Given the string, check if it is a palindrome.

func checkPalindrome(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}
