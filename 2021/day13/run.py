#!/usr/bin/env python3

import cProfile
import io
import pstats

import unittest
import logging
import re
from time import perf_counter_ns
import numpy

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


class Grid(object):

    """Docstring for Grid. """

    def __init__(self, points):
        """TODO: to be defined. """

        self.points = points
        self.grid = []
        self.max_x = max([p[0] for p in self.points])
        self.max_y = max([p[1] for p in self.points])
        self._build_grid()

    def __repr__(self):
        return '\n' + \
            '\n'.join([' '.join(['#' if c else '.' for c in row]) for row in self.grid]) + \
            '\n'

    def _build_grid(self):
        self.grid = [[False]*(self.max_x + 1) for i in range(self.max_y + 1)]

        for x, y in self.points:
            self.grid[y][x] = True

        # for y in range(self.max_y + 1):
        #     line = []
        #     for x in range(self.max_x + 1):
        #         if (x, y) in self.points:
        #             line.append(True)
        #         else:
        #             line.append(False)
        #     self.grid.append(line)

    def vfold(self, axis):
        self._fold(axis)

    def hfold(self, axis):
        # Transpose array before fold using numpy / revert after fold
        self.grid = numpy.array(self.grid).transpose().tolist()
        self._fold(axis)
        self.grid = numpy.array(self.grid).transpose().tolist()

    def _fold(self, axis):
        output = []
        for line in self.grid:
            # split in 2 array
            part1, part2 = line[0:axis], line[axis + 1:]
            # Then merge both
            for i, n in enumerate(part2):
                if n:
                    part1[-(i + 1)] = n
            output.append(part1)
        self.grid = output

    def count_dot(self):
        result = 0
        for line in self.grid:
            result += sum([1 for i in line if i])
        return result


# @profile
def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    points, folds = parse_input(lines)
    LOG.debug(f"points: {points} - folds: {folds}")

    g = Grid(points)
    LOG.debug(g)

    # Fold grid
    for fold in folds:
        if fold[0] == 'x':
            g.vfold(fold[1])

        elif fold[0] == 'y':
            g.hfold(fold[1])

        LOG.debug(f"fold: {fold} => {g}")

        if part == 1:
            # Stop after first grid fold
            break

    if part == 1:
        return g.count_dot()
    else:
        LOG.debug("Part 2 result")
        print(g)


def parse_input(lines):
    points = []
    folds = []
    for line in lines:
        if ',' in line:
            x, y = [int(i) for i in line.strip().split(',')]
            points.append((x, y))
        if 'fold' in line:
            i, v = re.findall(r"fold along (x|y)=(\d+)", line).pop()
            folds.append((i, int(v)))

    return points, folds


class Test(unittest.TestCase):
    def test_algo(self):
        example_result = algo('input_test.txt', part=1)
        self.assertEqual(example_result, 17)
        example_result = algo('input_test.txt', part=2)
        self.assertEqual(example_result, None)


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
