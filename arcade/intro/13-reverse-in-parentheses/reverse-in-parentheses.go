package codesignal

/* reverseInParentheses

Write a function that reverses characters in (possibly nested) parentheses in
the input string.

Input strings will always be well-formed with matching ()s.

Example

    For inputString = "(bar)", the output should be
	reverseInParentheses(inputString) = "rab";

    For inputString = "foo(bar)baz", the output should be
	reverseInParentheses(inputString) = "foorabbaz";

    For inputString = "foo(bar)baz(blim)", the output should be
	reverseInParentheses(inputString) = "foorabbazmilb";

    For inputString = "foo(bar(baz))blim", the output should be
    reverseInParentheses(inputString) = "foobazrabblim".
    Because "foo(bar(baz))blim" becomes "foo(barzab)blim" & then "foobazrabblim".

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string inputString

	A string consisting of lowercase English letters and the characters ( and ).
	It is guaranteed that all parentheses in inputString form a regular
	bracket sequence.

    Guaranteed constraints:
    0 ≤ inputString.length ≤ 50.
*/

// Constants are resolved at compile time. They increase readability without
// hindering preformace.
const (
	openParen  = '('
	closeParen = ')'
)

func reverseInParentheses(s string) string {
	runes := []rune(s)
	locations := locateParenPairs(runes)
	for i := len(locations) - 1; i >= 0; i-- {
		// iterating backwards insures that nested sets are reversed first.
		runes = reverse(runes, locations[i][0], locations[i][1])
	}
	runes = filter(runes, func(r rune) bool {
		return r == openParen || r == closeParen
	})
	return string(runes)
}

// filter filters the elements of a slice based on passed function.
func filter(runes []rune, f func(rune) bool) []rune {
	output := make([]rune, 0)
	for _, r := range runes {
		if !f(r) {
			output = append(output, r)
		}
	}
	return output
}

// reverse reverses the order of runes slice between the start & end indices.
func reverse(runes []rune, start, end int) []rune {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

// locateParenPairs locates the indices of each pair of parentheses. Returns a
// slice of "tuples": [[openingIndex, closingIndex], [openingIndex, closingIndex]...]
func locateParenPairs(runes []rune) (locations [][]int) {
	searchIndex, length := 0, len(runes)
	for {
		op := findRune(runes, openParen, searchIndex, length)
		if op == -1 {
			break
		}
		searchIndex = op + 1
		cp := findMatchingClosingParen(runes, op, length)
		locations = append(locations, []int{op, cp})
	}
	return locations
}

// findMatchingClosingParen locates & returns the matching closing paren.
func findMatchingClosingParen(runes []rune, openLoc, end int) int {
	nestedTracker := 0
	for i := openLoc + 1; i < end; i++ {
		switch runes[i] {
		case openParen:
			nestedTracker++
		case closeParen:
			if nestedTracker > 0 {
				nestedTracker--
				continue
			}
			return i
		}
	}
	return -1
}

// findRune returns the index of the input paremeter 'r' within 'runes'. If no
// instances of 'r' are found, -1 is returned.
func findRune(runes []rune, r rune, start, end int) int {
	// I've omitted the error handing.
	for i := start; i < end; i++ {
		if runes[i] == r {
			return i
		}
	}
	return -1
}
