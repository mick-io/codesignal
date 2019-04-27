import pytest
from .base_conversion import baseConversion


def test_base_conversion():
    tests = [
        {
            "n": "1302",
            "x": 5,
            "expect": "ca",
        },
        {
            "n": "1010100101",
            "x": 2,
            "expect": "2a5",
        },
        {
            "n": "z",
            "x": 36,
            "expect": "23",
        },
        {
            "n": "30",
            "x": 4,
            "expect": "c",
        },
        {
            "n": "6",
            "x": 7,
            "expect": "6",
        },
        {
            "n": "337",
            "x": 8,
            "expect": "df",
        },
        {
            "n": "ab3f",
            "x": 16,
            "expect": "ab3f",
        },
        {
            "n": "0",
            "x": 2,
            "expect": "0",
        },
        {
            "n": "ci",
            "x": 19,
            "expect": "f6",
        },
        {
            "n": "8c4897",
            "x": 13,
            "expect": "32b5cc",
        },
    ]
    for test in tests:
        assert baseConversion(test["n"], test["x"]) == test["expect"]
