import pytest
from .count_bits import countBits


def test_count_bits():
    tests = [(50, 6), (1, 1), (1000000000, 30), (237487384, 28), (278, 9)]
    for input, expect in tests:
        assert countBits(input) == expect
