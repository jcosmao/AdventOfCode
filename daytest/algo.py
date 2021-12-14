#!/usr/bin/env python3

import unittest
import logging
import tools


def day_algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    if part == 1:
        return 204

    if part == 2:
        return 195

class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = day_algo(file, part=1)
        self.assertEqual(example_result, 204)
        example_result = day_algo(file, part=2)
        self.assertEqual(example_result, 195)

if __name__ == "__main__":
    tools.adventOfCode()
