package codesignal

/* areSimilar
Two arrays are called similar if one can be obtained from another by swapping
at most one pair of elements in one of the arrays.

Given two arrays a and b, check whether they are similar.

Example

For a = [1, 2, 3] and b = [1, 2, 3], the output should be
areSimilar(a, b) = true.

The arrays are equal, no need to swap any elements.

For a = [1, 2, 3] and b = [2, 1, 3], the output should be
areSimilar(a, b) = true.

We can obtain b from a by swapping 2 and 1 in b.

For a = [1, 2, 2] and b = [2, 1, 1], the output should be
areSimilar(a, b) = false.

Any swap of any two elements either in a or in b won't make a and b equal.

Input/Output

[execution time limit] 4 seconds (go)

[input] array.integer a

Array of integers.

Guaranteed constraints:
3 ≤ a.length ≤ 105,
1 ≤ a[i] ≤ 1000.

[input] array.integer b

Array of integers of the same length as a.

Guaranteed constraints:
b.length = a.length,
1 ≤ b[i] ≤ 1000.

[output] boolean
true if a and b are similar, false otherwise.

*/

func areSimilar(a []int, b []int) bool {
	var swapA, swapB int
	cached, swapped := false, false
	for i, length := 0, len(a); i < length; i++ { // caching len(var) saves CPU cycles.
		switch intA, intB := a[i], b[i]; {
		case intA == intB:
			continue
		case !cached:
			swapA, swapB = intA, intB
			cached = true
		case !swapped:
			swapped = true
			if intA != swapB || intB != swapA {
				return false
			}
		default:
			return false
		}
	}
	return true
}
