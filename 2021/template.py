#!/usr/bin/env python3

import unittest
import logging
from time import perf_counter_ns

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)


def algo(input, part):
    with open(input) as f:
        lines = f.readlines()


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file, part=1)
        self.assertEqual(example_result, 204)
        example_result = algo(file, part=2)
        self.assertEqual(example_result, 195)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    # Test pass on example input, work on data set
    if not r.result.failures and not r.result.errors:
        LOG.setLevel(logging.INFO)

        START = perf_counter_ns()
        res = algo('input.txt', part=1)
        STOP = perf_counter_ns()
        LOG.info(f"RESULT part1 from input.txt = {res} - (TIME {STOP - START} ns)")

        START = perf_counter_ns()
        res = algo('input.txt', part=2)
        STOP = perf_counter_ns()
        LOG.info(f"RESULT part2 from input.txt = {res} - (TIME {STOP - START} ns)")
