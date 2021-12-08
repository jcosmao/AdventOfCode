#!/usr/bin/env python3

import argparse
import unittest
import logging

LOG = logging.getLogger()


def algo(input, days):
    with open(input) as f:
        lines = f.readlines()
    data = [int(i) for i in lines[0].split(',')]
    result = _process_lanternfish(data, days)

    return result

def _process_lanternfish(data, days):
    count = [0,0,0,0,0,0,0,0,0]
    for d in data:
        count[d] += 1
    for day in range(1, days + 1):
        LOG.debug(f"DAY {day} = {count}")
        born = count.pop(0)
        if born:
            count[6] += born
            count.append(born)
        else:
            count.append(0)
    return sum(count)



class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file, days=18)
        self.assertEqual(example_result, 26)
        example_result = algo(file, days=80)
        self.assertEqual(example_result, 5934)
        example_result = algo('input.txt', days=80)
        self.assertEqual(example_result, 350605)


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.DEBUG,
        format='[%(levelname)8s]  %(message)s'
    )
    r = unittest.main(exit=False)

    if not r.result.failures:
        res = algo('input.txt', days=256)
        LOG.info(f"RESULT from input.txt = {res}")
