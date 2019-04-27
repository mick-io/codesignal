const assert = require("assert");

/* CenturyFromYear

Given a year, return the century it is in. The first century spans from the year
1 up to and including the year 100, the second - from the year 101 up to and
including the year 200, etc.

Example

    For year = 1905, the output should be
    centuryFromYear(year) = 20;
    For year = 1700, the output should be
    centuryFromYear(year) = 17.

Input/Output

    [execution time limit] 4 seconds (go)

    [input] integer year

    A positive integer, designating the year.

    Guaranteed constraints:
    1 ≤ year ≤ 2005.

    [output] integer
        The number of the century the year is in.
*/
function centuryFromYear(year) {
  return Math.ceil(year / 100);
}

// Unit Test
(function() {
  let test = [
    {
      year: 1905,
      expect: 20
    },
    {
      year: 1700,
      expect: 17
    },
    {
      year: 1988,
      expect: 20
    },
    {
      year: 2000,
      expect: 20
    },
    {
      year: 2001,
      expect: 21
    },
    {
      year: 200,
      expect: 2
    },
    {
      year: 374,
      expect: 4
    },
    {
      year: 45,
      expect: 1
    },
    {
      year: 8,
      expect: 1
    }
  ];
  test.forEach(test => {
    let { year, expect } = test;
    let out = centuryFromYear(year);
    let msg = `[FAIL] centuryFromYear(${year}) == ${out} expect ${expect}`;
    assert.strictEqual(out, expect, msg);
  });
  console.log("\nPASSED");
})();
