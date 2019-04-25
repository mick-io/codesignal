package codesignal

/* specialPolynomial

Given integers x and n, find the largest integer k such that x^0+x^1+x^2+...+x^k ≤ n.

Example

For x = 2 and n = 5, the output should be
specialPolynomial(x, n) = 1.

We have 2^0 + 2^1 < 5 and 2^0 + 2^1 + 2^2 > 5.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] integer x

    Guaranteed constraints:
    1 ≤ x ≤ 10.

    [input] integer n

    Guaranteed constraints:
    5 ≤ n < 109.

    [output] integer
*/

import "math"

func specialPolynomial(x int, n int) int {
	base, max := float64(x), float64(n)
	exponent, total := 0.0, 0.0
	for {
		switch total += math.Pow(base, exponent); {
		case total < max:
			exponent++
		case total > max:
			return int(exponent) - 1
		default:
			return int(exponent)
		}
	}
}
