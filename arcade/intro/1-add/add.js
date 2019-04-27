const assert = require("assert");

/* add

Write a function that returns the sum of two numbers.

Example

For param1 = 1 and param2 = 2, the output should be
add(param1, param2) = 3.

Input/Output

    [execution time limit] 4 seconds (js)

    [input] integer param1

    Guaranteed constraints:
    -1000 ≤ param1 ≤ 1000.

    [input] integer param2

    Guaranteed constraints:
    -1000 ≤ param2 ≤ 1000.

    [output] integer
        The sum of the two inputs.

 */

function add(n1, n2) {
  return n1 + n2;
}


// Unit Test
(function() {
  let test = [
    {
      n1: 1,
      n2: 2,
      expect: 3
    },
    {
      n1: 0,
      n2: 1000,
      expect: 1000
    },
    {
      n1: 2,
      n2: -39,
      expect: -37
    },
    {
      n1: 99,
      n2: 100,
      expect: 199
    },
    {
      n1: -100,
      n2: 100,
      expect: 0
    },
    {
      n1: -1000,
      n2: -1000,
      expect: -2000
    }
  ];
  test.forEach(test => {
    let { n1, n2, expect } = test;
    let out = add(n1, n2);
    let msg = `[FAIL] add{${n1}, ${n2}} == ${out} expect ${expect}`;
    assert.strictEqual(out, expect, msg);
  });
  console.log("\n[PASSED]")
})();
