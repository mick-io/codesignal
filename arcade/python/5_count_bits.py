"""Implement the missing code, denoted by ellipses. You may not modify the
pre-existing code.

Implement a function that, given an integer n, uses a specific method on it
and returns the number of bits in its binary representation.

Note: in this task and most of the following tasks you will be given a code
snippet with some part of it replaced by the ellipsis (...). Only this part is
allowed to be changed.

Example

For n = 50, the output should be
countBits(n) = 6.

5010 = 1100102, a number that consists of 6 digits.
Thus, the output should be 6.

Input/Output

    [execution time limit] 4 seconds (py)


    [input] integer n

    A positive integer.

    Guaranteed constraints:
    1 ≤ n ≤ 109.

    [output] integer
        The number of bits in binary representation of n.
"""


def countBits(n):
    return n.bit_length()


if __name__ == "__main__":
    assertions = [  # [(n, expect)...]
        (50, 6),
        (1, 1),
        (1000000000, 30),
        (237487384, 28),
        (278, 9)
    ]
    for n, expect in assertions:
        out = countBits(n)
        if out != expect:
            e = f"[FAIL] countBits({n}) == {out}, expect {expect}"
            raise AssertionError(e)
    print("[PASSED]")
