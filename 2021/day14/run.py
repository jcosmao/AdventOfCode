#!/usr/bin/env python3

import unittest
import logging
from time import perf_counter_ns
from collections import Counter

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)


def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    input, elt_map = parse_input(lines)
    LOG.debug(f"input: {input} - elt_map: {elt_map}")

    if part == 1:
        iteration = 10
    elif part == 2:
        iteration = 40

    # count each distinct element => needed to compute final result
    elt_count = dict(Counter(input))
    for elt in set(list(''.join(elt_map.keys()))):
        if elt not in elt_count.keys():
            elt_count[elt] = 0

    # count element combination on each iteration
    elt_iter = {k: 0 for k in elt_map.keys()}
    for i in range(len(input) - 1):
        k = ''.join(list(input)[i:i+2])
        elt_iter[k] += 1

    for iter in range(iteration):
        LOG.debug(f"Iteration {iter + 1}")
        LOG.debug(elt_iter)
        LOG.debug(elt_count)

        # will store new element combination creation per iteration
        elt_current = {k:0 for k in elt_map.keys()}

        # Loop over each element combination created on previous iteration
        # to generate new ones
        for elt in elt_iter.keys():
            counter = elt_iter[elt]
            if counter == 0:
                continue

            # element combination  AA create => B
            # increment counter for each new element created
            new = elt_map[elt]
            elt_count[new] += counter

            # each element combination create 2 new elements
            # AA => AB BA
            elt_current[f"{list(elt)[0]}{new}"] += counter
            elt_current[f"{new}{list(elt)[1]}"] += counter
            # increment counter of produced

        # store created combination for next iteration
        elt_iter = elt_current

    LOG.debug(elt_count)
    LOG.debug(elt_iter)

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
        example_result = algo('input_test.txt', part=2)
        self.assertEqual(example_result, 2188189693529)


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
