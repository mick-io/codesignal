"""Let's call a list beautiful if its first element is equal to its last
element, or if a list is empty. Given a list a, your task is to chop off its
first and its last element until it becomes beautiful. Implement a function that
will make the given a beautiful as described, and return the resulting list as
an answer.

Hint: one of the features introduced in Python 3 called extended unpacking could
help here.

Example

    For a = [3, 4, 2, 4, 38, 4, 5, 3, 2], the output should be
    listBeautifier(a) = [4, 38, 4].

    Here's how the answer is obtained:
    [3, 4, 2, 4, 38, 4, 5, 3, 2] => [4, 2, 4, 38, 4, 5, 3] =>
    [2, 4, 38, 4, 5] => [4, 38, 4].

    For a = [1, 4, -5], the output should be
    listBeautifier(a) = [4].

Input/Output

    [execution time limit] 4 seconds (py3)

    [input] array.integer a

    A list of integers.

    Guaranteed constraints:
    0 ≤ a.length ≤ 50,
    1 ≤ a[i] ≤ 100.

    [output] array.integer
        A beautiful list obtained as described above.
"""


def listBeautifier(a):
    res = a[:]
    while res and res[0] != res[-1]:
        _, *res, _ = res
    return res


if __name__ == "__main__":
    assertions = [  # [(input, expect)...]
        ([3, 4, 2, 4, 38, 4, 5, 3, 2], [4, 38, 4]),
        ([1, 4, -5], [4]),
        ([], []),
        ([8], [8]),
        ([10, 2, 10, 9, 7, 3, 8, 5, 3, 2, 8, 7], []),
        (
            [8, 5, 1, 2, 3, 8, 1, 10, 5, 1, 4, 6, 9, 2, 8, 8, 9, 4, 10, 3],
            [8, 1, 10, 5, 1, 4, 6, 9, 2, 8]
        ),
        ([10, 8, 10], [10, 8, 10]),
        (
            [4, 9, 3, 5, 7, 6, 1, 8, 7, 6, 8, 9, 2],
            [9, 3, 5, 7, 6, 1, 8, 7, 6, 8, 9]
        ),
        (
            [2, 4, 2, 10, 4, 2, 6, 1, 5, 1, 10, 2, 10, 5, 2, 10],
            [10, 4, 2, 6, 1, 5, 1, 10, 2, 10]
        ),
        ([5, 3, 3, 7, 1, 1, 7, 9, 3, 1, 7, 1], []),
    ]
    for input, expect in assertions:
        out = listBeautifier(input)
        if input != expect:
            e = f"[FAIL] listBeautifier({input}) == {out} expect {expect}"
            raise AssertionError(e)
    print("[PASSED]")
