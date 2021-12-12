#!/usr/bin/env python3

import unittest
import logging
import numpy

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)

POINTS = {
    1: {')': 3, ']': 57, '}': 1197, '>': 25137},
    2: {')': 1, ']': 2,  '}': 3,    '>': 4},
}

MATCH = {
    '(': ')', '[': ']', '{': '}', '<': '>',
}

errors = {
    ')': 0, ']': 0, '}': 0, '>': 0,
}


def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    part2 = []

    for i, line in enumerate(lines):
        search_line = list(line.strip())
        try:
            parse_line(search_line)
            LOG.info(f"Line complete ? {search_line}")

            missing_complete = [MATCH[c] for c in reversed(list(filter(lambda x: x is not None, search_line)))]
            LOG.info(f"Complete: {missing_complete}")
            part2.append(missing_complete)

        except Exception as e:
            LOG.error(f"Line {i} not complete")
            LOG.exception(e)

    LOG.debug(errors)

    if part == 1:
        result = 0
        for item in errors.items():
            result += item[1] * POINTS[part][item[0]]
    else:
        result = 0
        scores = []
        for compl in part2:
            r = 0
            for c in compl:
                LOG.debug(f"{r} * 5 + {POINTS[part][c]}")
                r = (r * 5) + POINTS[part][c]
                LOG.debug(f"r = {r}")

            scores.append(r)

        LOG.info(scores)
        result = sorted(scores)[int(len(scores) / 2)]

    return result


def parse_line(search_line):
    for i, char in enumerate(search_line):
        for item in MATCH.items():
            if char == item[1]:
                found, index = find_match(item[0], search_line, i)
                if not found:
                    LOG.debug(f"{char} not found in line")
                    LOG.debug(search_line)
                    errors[char] += 1
                    raise ValueError("Line corrupt")
                else:
                    LOG.debug(f"{char} close {item[0]} found in index {index}")
                    search_line[index] = None
                    search_line[i] = None


def find_match(char, array, index):
    LOG.debug(f"Search {char} between [0:{index}]")
    for i, open_char in reversed(list(enumerate(array))[0:index]):
        if not open_char:
            continue
        elif char == open_char:
            return True, i
        else:
            return False, False


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file, part=1)
        self.assertEqual(example_result, 26397)
        example_result = algo(file, part=2)
        self.assertEqual(example_result, 288957)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    if not r.result.failures and not r.result.errors:
        res = algo('input.txt', part=1)
        LOG.info(f"RESULT part1 from input.txt = {res}")
        res = algo('input.txt', part=2)
        LOG.info(f"RESULT part2 from input.txt = {res}")
