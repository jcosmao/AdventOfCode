#!/usr/bin/env python3

import cProfile
import io
import pstats

import unittest
import logging
import re
from time import perf_counter_ns
import numpy
from collections import Counter

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)


# https://towardsdatascience.com/how-to-profile-your-code-in-python-e70c834fad89
def profile(func):
    def wrapper(*args, **kwargs):
        pr = cProfile.Profile()
        pr.enable()
        retval = func(*args, **kwargs)
        pr.disable()
        s = io.StringIO()
        sortby = pstats.SortKey.CUMULATIVE  # 'cumulative'
        ps = pstats.Stats(pr, stream=s).sort_stats(sortby)
        ps.print_stats()
        print(s.getvalue())
        return retval

    return wrapper



# @profile
def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    input, elt_map = parse_input(lines)
    LOG.debug(f"input: {input} - elt_map: {elt_map}")

    if part == 1:
        iteration = 10
    elif part == 2:
        iteration = 40

    for iter in range(iteration):
        LOG.info(f"Iteration {iter +1}")
        new = ''
        for i in range(len(input)-1):
            ins = elt_map[''.join(list(input)[i:i+2])]
            new = new + f"{list(input)[i:i+2][0]}{ins}"
            if i == len(input) - 2:
                new = new + list(input)[i:i+2][1]

        # LOG.debug(new)
        input = new

        LOG.debug(len(input))
        elt_count = Counter(input)
        LOG.debug(elt_count)

    elt_count = Counter(input)
    max_elt = max(elt_count.values())
    min_elt = min(elt_count.values())

    return max_elt - min_elt

def parse_input(lines):
    input = lines[0].strip()
    elt_map = {}
    for line in lines:
        if '->' in line:
            k, v = line.strip().split(' -> ')
            elt_map[k] = v

    return input, elt_map


class Test(unittest.TestCase):
    def test_algo(self):
        example_result = algo('input_test.txt', part=1)
        self.assertEqual(example_result, 1588)
        # example_result = algo('input_test.txt', part=2)
        # self.assertEqual(example_result, 2188189693529)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    # Test pass on example input, work on data set
    if not r.result.failures and not r.result.errors:
        LOG.setLevel(logging.INFO)

        START = perf_counter_ns()
        res = algo('input.txt', part=1)
        STOP = perf_counter_ns()
        LOG.info(
            f"RESULT part1 from input.txt = {res} - (TIME {(STOP - START) / 1000000} ms)")

        START = perf_counter_ns()
        res = algo('input.txt', part=2)
        STOP = perf_counter_ns()
        LOG.info(
            f"RESULT part2 from input.txt = {res} - (TIME {(STOP - START) / 1000000} ms)")
