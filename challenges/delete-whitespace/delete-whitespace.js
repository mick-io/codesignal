/*
Given a string, delete whitespaces from it.

Example

For inputStr = "a b cd e", the output should be
deleteWhitespaces(inputStr) = "abcde".

Input/Output

    [execution time limit] 4 seconds (go)

    [input] string inputStr

    String consisting of lowercase English letters and whitespace characters.

    Guaranteed constraints:
    5 ≤ inputStr.length ≤ 15.

    [output] string
        inputString without whitespace characters.
*/

function deleteWhitespaces(s) {
  return s.split("").filter(char => char != " ").join("");
}

(function() {
  let assertions = [
    {
      input: "a  b cd  e",
      expect: "abcde"
    },
    {
      input: "a   b c  dc ",
      expect: "abcdc"
    }
  ];
  for (let { input, expect } of assertions) {
    let output = deleteWhitespaces(input);
    if (output != expect) {
      let e = `[FAIL] deleteWhitespace(${input}) == ${output} expect ${expect}`;
      throw new Error(e);
    }
  }
  console.info("[SUCCESS]")
})();
