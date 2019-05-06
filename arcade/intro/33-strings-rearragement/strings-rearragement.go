package codesingal

/* stringsRearrangement
Given an array of equal-length strings, you'd like to know if it's possible to
rearrange the order of the elements in such a way that each consecutive pair of
strings differ by exactly one character. Return true if it's possible, and false
if not.

Note: You're only rearranging the order of the strings, not the order of the
letters within the strings!

Example

    For inputArray = ["aba", "bbb", "bab"], the output should be
    stringsRearrangement(inputArray) = false.

    There are 6 possible arrangements for these strings:
        ["aba", "bbb", "bab"]
        ["aba", "bab", "bbb"]
        ["bbb", "aba", "bab"]
        ["bbb", "bab", "aba"]
        ["bab", "bbb", "aba"]
        ["bab", "aba", "bbb"]

	None of these satisfy the condition of consecutive strings differing by 1
	character, so the answer is false.

    For inputArray = ["ab", "bb", "aa"], the output should be
    stringsRearrangement(inputArray) = true.

	It's possible to arrange these strings in a way that each consecutive pair
	of strings differ by 1 character (eg: "aa", "ab", "bb" or "bb", "ab", "aa"),
	so return true.

Input/Output

    [execution time limit] 5 seconds (ts)

    [input] array.string inputArray

    A non-empty array of strings of lowercase letters.

    Guaranteed constraints:
    2 ≤ inputArray.length ≤ 10,
    1 ≤ inputArray[i].length ≤ 15.

    [output] boolean
		Return true if the strings can be reordered in such a way that each
		neighbouring pair of strings differ by exactly one character (false
		otherwise).
*/

import (
	"fmt"
	"sort"
)

func stringsRearrangement(slice []string) bool {
	sort.Strings(slice)
	fmt.Println(slice)
	for i, s := range slice[1:] {
		if nDifferent(s, slice[i-1]) != 1 {
			return false
		}
	}
	return true
}

func nDifferent(s1, s2 string) (n int) {
	runes1, runes2 := []rune(s1), []rune(s2)
	for _, r1 := range runes1 {
		for i, r2 := range runes2 {
			if r1 == r2 {
				runes2[i] = 0
			} else {
				n++
			}
		}
	}
	return
}
