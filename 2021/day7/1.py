#!/usr/bin/env python3

import unittest
import sys
import logging
import numpy

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)

def algo(input):
    with open(input) as f:
        lines = f.readlines()
    num = [int(i) for i in lines[0].split(',')]
    median = numpy.percentile(num, 50, interpolation='nearest')
    fuel = 0
    for i in num:
        fuel += max(i, median) - min(i, median)

    return fuel

def algo2(input):
    with open(input) as f:
        lines = f.readlines()
    num = [int(i) for i in lines[0].split(',')]

    res = {}
    for n in range(min(num), max(num)):
        fuel = _fuel_consumption2(num, n)
        LOG.debug(f"target {n} -> cost = {fuel}")
        res[fuel] = n

    LOG.debug(min(res.keys()))
    return min(res.keys())
    # check_range = range(min(num), max(num))
    # check = {}
    # for i in check_range:
    #     fuel = 0
    #     fuel = _fuel_consumption2(num, i)
    #     LOG.debug(f"median = {median} = {fuel}")
    #     check[fuel] = i
    # LOG.debug(check)

    # LOG.debug(res[min(res.keys())])
    # LOG.debug(min(res.keys()))
    # LOG.debug(check[min(check.keys())])
    # LOG.debug(min(check.keys()))

    # return min(res.keys())

def _partial_sum(n):
    """https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF"""
    return int(n * (n + 1) / 2)

def _fuel_consumption2(position, target):
    fuel = 0
    for i in position:
        move = max(i, target) - min(i, target)
        fuel += _partial_sum(move)
    return fuel

def _closest(list, nb):
    aux = []
    for val in list:
        aux.append(abs(nb - val))
    return aux.index(min(aux))

class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file)
        self.assertEqual(example_result, 37)
        example_result = algo2(file)
        self.assertEqual(example_result, 168)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    if not r.result.failures:
        res = algo('input.txt')
        LOG.info(f"RESULT from input.txt = {res}")

        res = algo2('input.txt')
        LOG.info(f"RESULT2 from input.txt = {res}")
