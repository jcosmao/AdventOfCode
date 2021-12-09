#!/usr/bin/env python3

import unittest
import sys
import logging
import numpy
import string

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.INFO,
    format='[%(levelname)8s]  %(message)s'
)

digits = {
    0: ['a', 'b', 'c', 'e', 'f', 'g'],
    1: ['c', 'f'],
    2: ['a', 'c', 'd', 'e', 'g'],
    3: ['a', 'c', 'd', 'f', 'g'],
    4: ['b', 'c', 'd', 'f'],
    5: ['a', 'b', 'd', 'f', 'g'],
    6: ['a', 'b', 'd', 'e', 'f', 'g'],
    7: ['a', 'c', 'f'],
    8: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
    9: ['a', 'b', 'c', 'd', 'f', 'g'],
}

# 1
# 4 -> subset of 9 and 8
# 7
# 8
# 3 -> subset of 9
# 3 - 6 -> reste 2
# 5 -> subset of 6
# reste => 2


def algo(input):

    result = 0
    parts = _parse(input)
    LOG.debug(parts)
    for l in parts[2]:
        for i in l:
            if len(i) in [len(digits[1]), len(digits[4]), len(digits[7]), len(digits[8])]:
                LOG.debug(sorted(i))
                result += 1

    return result

def algo2(input):
    parts = _parse(input)
    total = 0
    for line in range(len(parts[1])):
        l = parts[1][line]
        output = parts[2][line]
        # Test _serachfor algo based on digits input
        # l = [digits[i] for i in digits.keys()]
        found = {}
        for i in range(10):
            for code in l:
                try:
                    _search_for(found, code)
                except Exception as e:
                    LOG.debug(e)

        LOG.debug(f"FOUND: {found}")

        code = _decode(found, output)
        LOG.debug(code)
        total += code

    LOG.debug(total)
    return total

def _decode(found, output):
    output_code = ''
    for code in output:
        for i in found.items():
            if set(i[1]) == set(code):
                output_code = output_code + str(i[0])

    return int(output_code)


def _search_for(found, code):
    LOG.debug(f" CODE = {code}")
    for i in [1,4,7,8]:
        if len(code) == len(digits[i]):
            LOG.debug(f"FOUND {i}")
            found[i] = code
            return

    if len(code) == len(digits[9]) and set(found[4]).issubset(set(code)):
        LOG.debug(f"FOUND 9")
        found[9] = code
        return
    elif len(code) == len(digits[3]) and set(code).issubset(found[9]) and len(set(code).difference(set(found[1]))) == 3:
        LOG.debug(f"FOUND 3")
        found[3] = code
        return
    elif len(code) == len(digits[0])  and len(set(code).difference(set(found[1]))) == 4:
        LOG.debug(f"FOUND 0")
        found[0] = code
        return
    elif len(code) == len(digits[6])  and len(set(code).difference(set(found[1]))) == 5:
        LOG.debug(f"FOUND 6")
        found[6] = code
        return
    elif len(code) == len(digits[5]) and set(code).issubset(set(found[6])):
        LOG.debug(f"FOUND 5")
        found[5] = code
        return
    elif len(code) == len(digits[5])  and set(code).issubset(set(found[9])):
        LOG.debug(f"FOUND 5")
        found[5] = code
        return
    elif len(found.keys()) == 9:
        LOG.debug(f"FOUND 2")
        found[2] = code
        return
    else:
        LOG.debug(f"CODE {code} NOT FOUND")


def _parse(input):

    with open(input) as f:
        lines = f.readlines()

    parsed = {1: [], 2: []}
    for l in lines:
        lpart = l.split(' | ')
        part1 = [list(i.strip()) for i in lpart[0].split(' ')]
        part2 = [list(i.strip()) for i in lpart[1].split(' ')]
        LOG.debug(part1)
        parsed[1].append(part1)
        parsed[2].append(part2)

    return parsed


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file)
        self.assertEqual(example_result, 26)
        example_result = algo2(file)
        self.assertEqual(example_result, 61229)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    if not r.result.failures:
        res = algo('input.txt')
        LOG.info(f"RESULT 1 from input.txt = {res}")
        res = algo2('input.txt')
        LOG.info(f"RESULT 2 from input.txt = {res}")
