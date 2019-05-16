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

func stringsRearrangement(s []string) bool {
	rearraged, s := []string{s[0]}, s[1:]
	for len(s) != 0 {
		similar, i := findSimilar(s, rearraged[len(rearraged)-1])
		if i != -1 {
			s = append(s[:i], s[i+1:]...)
			rearraged = append(rearraged, similar)
			continue
		}
		similar, i = findSimilar(s, rearraged[0])
		if i != -1 {
			s = append(s[:i], s[i+1:]...)
			rearraged = append([]string{similar}, rearraged...)
			continue
		}
		ok := false
		for j, length := 0, len(rearraged); j < length-1; j++ {
			for k, chars := range s {
				if nDifferent(chars, rearraged[j]) == 1 && nDifferent(chars, rearraged[j+1]) == 1 {
					temp := append([]string{chars}, rearraged[k+1:]...)
					rearraged = append(rearraged[:j], temp...)
					s = append(s[:k], s[k+1:]...)
					ok = true
					break
				}
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

func findSimilar(slice []string, str string) (similar string, index int) {
	for i, s := range slice {
		if nDifferent(s, str) == 1 {
			return s, i
		}
	}
	return "", -1
}

func nDifferent(s1, s2 string) (differences int) {
	runes1, runes2 := []rune(s1), []rune(s2)
	for i, length := 0, len(runes1); i < length; i++ {
		if runes1[i] != runes2[i] {
			differences++
		}
	}
	return
}
