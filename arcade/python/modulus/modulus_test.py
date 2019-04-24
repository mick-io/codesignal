import pytest
from .modulus import modulus


def test_modulus():
    tests = [(15, 1), (23.12, -1), (0, 0), (232.2423, -1), (30, 0), (11, 1)]
    for input, expect in tests:
        assert modulus(input) == expect
