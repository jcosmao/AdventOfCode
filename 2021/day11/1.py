#!/usr/bin/env python3

import unittest
import logging
from time import perf_counter_ns

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)


class Point():
    '''
        single point in the grid
    '''

    def __init__(self, x, y, value, grid):
        self.x = x
        self.y = y
        self.value = value
        self.has_flash = False
        self.grid = grid

    def neighbors(self, filter_flashed=False):

        neighbors = [
            (self.x - 1, self.y),  # up
            (self.x + 1, self.y),  # down
            (self.x, self.y - 1),   # left
            (self.x, self.y + 1),   # right
            (self.x - 1, self.y - 1),  # up left
            (self.x - 1, self.y + 1),  # up right
            (self.x + 1, self.y - 1),  # down left
            (self.x + 1, self.y + 1),  # down right
        ]

        for x, y in neighbors:
            if x > -1 and y > -1 and x < self.grid.x_dim and y < self.grid.y_dim:
                if filter_flashed and self.has_flash:
                    continue
                yield self.grid.points[x][y]

    def flash(self, neighbor=False):

        if neighbor and not self.has_flash:
            self.value += 1

        if self.value > 9 and not self.has_flash:
            self.has_flash = True
            self.grid.flash()
            for n in self.neighbors():
                n.flash(neighbor=True)
            self.value = 0
            return True
        else:
            return False


class Grid():
    '''
     class for the grid, with methods to flash
    '''

    def __init__(self, data):
        self.data = data
        self.points = data
        for x, line in enumerate(data):
            for y, value in enumerate(line):
                self.points[x][y] = Point(x, y, value, grid=self)

        self.x_dim = len(self.points)
        self.y_dim = len(self.points[0])
        self.total_flash = 0
        self.iter_flash = 0

    def __repr__(self):
        display = '\n'
        for line_points in self.points:
            line = ''
            for p in line_points:
                if p.has_flash:
                    val = f"*{p.value}"
                else:
                    val = f" {p.value}"
                line = f"{line} | {val} |"

            display = f"{display}\n{line}"
        return display

    def yield_points(self):
        for x, line in enumerate(self.points):
            for y, point in enumerate(line):
                yield x, y, point

    def increase_all(self):
        for _, _, p in self.yield_points():
            p.value += 1

    def flash(self):
        self.total_flash += 1
        self.iter_flash += 1

    def flash_all(self):
        for _, _, p in self.yield_points():
            p.flash()

    def reset_flash(self):
        self.iter_flash = 0
        for _, _, p in self.yield_points():
            p.has_flash = False


def algo(input, part, iteration=100):
    with open(input) as f:
        lines = f.readlines()

    if part == 2:
        iteration = 1000000

    matrix = []
    for l in lines:
        matrix.append([int(i) for i in list(l.strip())])

    grid = Grid(matrix)

    for iter in range(1, iteration + 1):
        grid.increase_all()
        grid.flash_all()
        LOG.debug(f"ITER = {iter} - FLASH GRID = {grid.total_flash}")
        LOG.debug(f"\n{grid}")

        if part == 2:
            if grid.iter_flash == (grid.x_dim * grid.y_dim):
                return iter

            # if iter == 195:
            #     LOG.debug(grid.iter_flash)
            #     break

        grid.reset_flash()

    return grid.total_flash


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file, part=1, iteration=10)
        self.assertEqual(example_result, 204)
        example_result = algo(file, part=1, iteration=100)
        self.assertEqual(example_result, 1656)
        example_result = algo(file, part=2)
        self.assertEqual(example_result, 195)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    # Test pass on example input, work on data set
    if not r.result.failures and not r.result.errors:
        LOG.setLevel(logging.INFO)

        START = perf_counter_ns()
        res = algo('input.txt', part=1, iteration=100)
        STOP = perf_counter_ns()
        LOG.info(f"RESULT part1 from input.txt = {res} - (TIME {STOP - START} ns)")

        START = perf_counter_ns()
        res = algo('input.txt', part=2)
        STOP = perf_counter_ns()
        LOG.info(f"RESULT part2 from input.txt = {res} - (TIME {STOP - START} ns)")
