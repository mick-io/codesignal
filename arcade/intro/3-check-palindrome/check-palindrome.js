const assert = require("assert");

/* checkPalindrome

Given the string, check if it is a palindrome.

Example

    For inputString = "aabaa", the output should be
    checkPalindrome(inputString) = true;
    For inputString = "abac", the output should be
    checkPalindrome(inputString) = false;
    For inputString = "a", the output should be
    checkPalindrome(inputString) = true.

Input/Output

    [execution time limit] 4 seconds (js)

    [input] string inputString

    A non-empty string consisting of lowercase characters.

    Guaranteed constraints:
    1 ≤ inputString.length ≤ 105.

    [output] boolean
        true if inputString is a palindrome, false otherwise.
*/

function checkPalindrome(s) {
  return s === s.split("").reverse().join("");
}

(function() {
  let tests = [
    {
      input: "aabaa",
      expect: true
    },
    {
      input: "abac",
      expect: false
    },
    {
      input: "a",
      expect: true
    },
    {
      input: "az",
      expect: false
    },
    {
      input: "abacaba",
      expect: true
    },
    {
      input: "z",
      expect: true
    },
    {
      input: "aaabaaaa",
      expect: false
    },
    {
      input: "zzzazzazz",
      expect: false
    },
    {
      input: "hlbeeykoqqqqokyeeblh",
      expect: true
    },
  ];
  tests.forEach(test => {
    let { input, expect } = test;
    let output = checkPalindrome(input);
    let msg = `[FAIL] checkPalindrome(${input}) == ${output} expect ${expect}`;
    assert.strictEqual(output, expect, msg);
  });
  console.log("\nPASSED");
})();
