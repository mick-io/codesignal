package codesignal

/* adjacentElementsProduct

Given an array of integers, find the pair of adjacent elements that has the
largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.integer inputArray

    An array of integers containing at least two elements.

    Guaranteed constraints:
    2 ≤ inputArray.length ≤ 10,
    -1000 ≤ inputArray[i] ≤ 1000.

    [output] integer
		The largest product of adjacent elements.
*/

func adjacentElementsProduct(input []int) (largest int) {
	largest = input[0] * input[1]
	for i, length := 2, len(input); i < length; i++ {
		product := input[i-1] * input[i]
		if product > largest {
			largest = product
		}
	}
	return largest
}

/* shapeArea

Below we will define an n-interesting polygon. Your task is to find the area
of a polygon for a given n.

A 1-interesting polygon is just a square with a side of length 1. An
n-interesting polygon is obtained by taking the n - 1-interesting polygon and
appending 1-interesting polygons to its rim, side by side. You can see the 1-,
2-, 3- and 4-interesting polygons in the picture below.

Example

    For n = 2, the output should be
    shapeArea(n) = 5;
    For n = 3, the output should be
    shapeArea(n) = 13.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] integer n

    Guaranteed constraints:
    1 ≤ n < 104.

    [output] integer
        The area of the n-interesting polygon.
*/

func shapeArea(n int) int {
	return n*n + (n-1)*(n-1)
}

/* Make Array Consecutive 2

Ratiorg got statues of different sizes as a present from CodeMaster for his
birthday, each statue having an non-negative integer size. Since he likes to
make things perfect, he wants to arrange them from smallest to largest so that
each statue will be bigger than the previous one exactly by 1. He may need some
additional statues to be able to accomplish that. Help him figure out the
minimum number of additional statues needed.

Example

For statues = [6, 2, 3, 8], the output should be
makeArrayConsecutive2(statues) = 3.

Ratiorg needs statues of sizes 4, 5 and 7.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] array.integer statues

    An array of distinct non-negative integers.

    Guaranteed constraints:
    1 ≤ statues.length ≤ 10,
    0 ≤ statues[i] ≤ 20.

    [output] integer
		The minimal number of statues that need to be added to existing statues
		such that it contains every integer size from an interval [L, R]
		(for some L, R) and no other sizes.
*/
