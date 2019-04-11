package codesignal

import "math"

func add(param1 int, param2 int) int {
	return param1 + param2
}

func centuryFromYear(year int) int {
	return int(math.Ceil(float64(year) / 100))
}

func checkPalindrome(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}
