#!/usr/bin/env python3

import unittest
import sys
import logging

LOG = logging.getLogger()
logging.basicConfig(encoding='utf-8', level=logging.DEBUG)


def algo(input):
    with open(input) as f:
        lines = f.readlines()
    return


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file)
        self.assertEqual(example_result, 5)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    if not r.result.failures:
        res = algo('input.txt')
        LOG.info(f"RESULT from input.txt = {res}")
