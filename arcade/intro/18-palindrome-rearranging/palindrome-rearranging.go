package codesignal

/* palindromeRearranging
Given a string, find out if its characters can be rearranged to form a palindrome.

Example

For inputString = "aabb", the output should be
palindromeRearranging(inputString) = true.

We can rearrange "aabb" to make "abba", which is a palindrome.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string inputString

    A string consisting of lowercase English letters.

    Guaranteed constraints:
    1 ≤ inputString.length ≤ 50.

    [output] boolean
        true if the characters of the inputString can be rearranged to form a palindrome, false otherwise.

*/

func palindromeRearranging(s string) bool {
	runes := []rune(s)
	isEven, soloFound := len(runes)%2 == 0, false
	for i, r := range runes {
		runes[i] = 0
		duplicate := find(r, runes)
		if duplicate != -1 {
			runes[duplicate] = 0
			continue // flat is better than nested.
		}
		if soloFound || isEven {
			return false
		}
		soloFound = true
	}
	return true
}

func find(search rune, runes []rune) int {
	for i, r := range runes {
		if r == search {
			return i
		}
	}
	return -1
}
