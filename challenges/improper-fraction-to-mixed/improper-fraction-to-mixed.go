package codesignal

/* improperFractionToMixed

Given a reduced improper fraction output it as a reduced mixed fraction.

Example

For a = [7, 2], the output should be
improperFractionToMixed(a) = [3, 1, 2].

Explanation: 7/2 = 3 + 1/2.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.integer a

	Array of two positive integers representing a reduced improper fraction
	a[0] / a[1].

    Guaranteed constraints:
    2 ≤ a[i] ≤ 25.

    [output] array.integer
		Array of three integers representing the reduced mixed fraction equal
		to a in the form b[0] + b[1] / b[2].

*/

import "math"

func improperFractionToMixed(a []int) []int {
	numerator, dominator := float64(a[0]), float64(a[1])
	whole := math.Floor(numerator / dominator)
	numerator -= whole * dominator
	return []int{int(whole), int(numerator), int(dominator)}
}
