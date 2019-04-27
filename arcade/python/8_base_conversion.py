""" Implement the missing code, denoted by ellipses. You may not modify the pre-
existing code.

Your university professor decided to have a little fun and asked the class to
implement a function that, given a number n and a base x, converts the number
from base x to base 16. To make things more interesting, he announced that the
first student to write the solution will have to answer fewer question than the
rest of the class during the final exam.

Laughing devilishly, you asked if it was okay to use a language of your choice,
and the unsuspecting professor answered "yes". It's settled then: Python is your
language of choice!

Now you're bound to win. Implement a function that, given an integer number n
and a base x, converts n from base x to base 16.

Example

For n = "1302" and x = 5, the output should be
	baseConversion(n, x) = "ca".

Here's why:
    13025 = 20210 = ca16.

Input/Output

[execution time limit] 4 seconds (py3)

[input] string n

	 A valid non-negative integer in base x. The string is guaranteed to consist of
	digits and lowercase English letters.

Guaranteed constraints:
	1 < n.length ≤ 10.

[input] integer x

	The base of n.

Guaranteed constraints:
	2 ≤ x ≤ 36.

[output] string
	The value of n in base 16. The string should contain only digits and lowercase
	English letters 'a' - 'f'.
"""


def baseConversion(n, x):
    return hex(int(n, x))[2:]


if __name__ == "__main__":
    assertions = [  # [(n, x, expect)...]
        ("1302", 5, "ca"),
        ("1010100101", 2, "2a5"),
        ("z", 36, "23"),
        ("30", 4, "c"),
        ("6", 7, "6"),
        ("337", 8, "df"),
        ("ab3f", 16, "ab3f"),
        ("0", 2, "0"),
        ("ci", 19, "f6"),
        ("8c4897", 13, "32b5cc"),
    ]
    for n, x, expect in assertions:
        out = baseConversion(n, x)
        if out != expect:
            e = f"[FAIL] baseConversion({n}, {x}) == {out} expect {expect}"
            raise AssertionError(e)
    print("[PASSED]")
