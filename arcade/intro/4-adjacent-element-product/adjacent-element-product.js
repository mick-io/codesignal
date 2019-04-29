/* adjacentElementsProduct
Given an array of integers, find the pair of adjacent elements that has the
largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

Input/Output

    [execution time limit] 4 seconds (js)

    [input] array.integer inputArray

    An array of integers containing at least two elements.

    Guaranteed constraints:
    2 ≤ inputArray.length ≤ 10,
    -1000 ≤ inputArray[i] ≤ 1000.

    [output] integer
        The largest product of adjacent elements.

 */
function adjacentElementsProduct(arr) {
  let products = arr.slice(1).map((n, i) => n * arr[i]);
  return Math.max(...products);
}

(function() {
  let assertions = [
    { arr: [3, 6, -2, -5, 7, 3], expect: 21 },
    { arr: [-1, -2], expect: 2 },
    { arr: [5, 1, 2, 3, 1, 4], expect: 6 },
    { arr: [1, 2, 3, 0], expect: 6 },
    { arr: [9, 5, 10, 2, 24, -1, -48], expect: 50 }
  ];
  for (let { arr, expect } of assertions) {
    let out = adjacentElementsProduct(arr);
    if (out != expect) {
      let e = `[FAIL] adjacentElementsProduct(${arr}) == ${out} expect ${expect}`;
      throw new Error(e);
    }
  }
  console.info("[PASSED]");
})();
