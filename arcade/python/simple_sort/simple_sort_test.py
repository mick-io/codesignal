import pytest
from .simple_sort import simpleSort


def test_modulus():
    tests = [
        ([2, 4, 1, 5], [1, 2, 4, 5]),
        ([3, 6, 1, 5, 3, 6], [1, 3, 3, 5, 6, 6]),
        ([100], [100]),
        ([-1, -2, 0], [-2, -1, 0]),
        ([100, 100, 100], [100, 100, 100]),
        ([1], [1])
    ]
    for input, expect in tests:
        assert simpleSort(input) == expect
