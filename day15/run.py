#!/usr/bin/env python3

import unittest
import logging
from time import perf_counter_ns
from collections import Counter
import numpy
import copy

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)


class Grid():
    '''
     class for the grid, with methods to flash
    '''

    def __init__(self, data):
        self.data = copy.deepcopy(data)
        self.points = data
        for x, line in enumerate(data):
            for y, value in enumerate(line):
                self.points[x][y] = Point(x, y, value, grid=self)

        self.x_dim = len(self.points)
        self.y_dim = len(self.points[0])

    def __repr__(self):
        display = '\n'
        for line_points in self.points:
            line = ''
            for p in line_points:
                val = f" {p.value}"
                # line = f"{line} | {val} |"
                line = f"{line}{val}"

            display = f"{display}\n{line}"
        return display

    def yield_points(self):
        for x, line in enumerate(self.points):
            for y, point in enumerate(line):
                yield x, y, point

    def min_cost_path(self):
        """ Work only if solution is right/down """

        self.data[0][0] = 0

        # For 1st column
        for i in range(1, self.x_dim):
            self.data[i][0] += self.data[i - 1][0]

        # For 1st row
        for j in range(1, self.y_dim):
            self.data[0][j] += self.data[0][j - 1]

        # For rest of the 2d matrix
        for i in range(1, self.x_dim):
            for j in range(1, self.y_dim):
                self.data[i][j] += min(self.data[i - 1][j], self.data[i][j - 1])

        # Returning the value in
        # last cell
        return self.data[self.x_dim - 1][self.y_dim - 1]


    def _get_min_cost(self, min_cost):
        min_point = list(min_cost)[0]
        for p in min_cost:
            if p.djikstra_dist < min_point.djikstra_dist:
                min_point = p
        return min_point

    def djikstra(self):
        """Djikstra implementation"""

        start_point = self.points[0][0]
        start_point.djikstra_dist = 0

        min_cost = set()
        min_cost.add(start_point)

        while min_cost:
            p = self._get_min_cost(min_cost)
            min_cost.remove(p)

            for n in p.neighbors():
                if self.points[n.x][n.y].djikstra_dist > p.djikstra_dist + n.value:
                    if self.points[n.x][n.y].djikstra_dist != float('inf'):
                        # Already in set
                        min_cost.remove(n)

                    self.points[n.x][n.y].djikstra_dist = p.djikstra_dist + n.value
                    n.djikstra_from = p

                    min_cost.add(n)

        # Display path
        p = self.points[self.x_dim - 1][self.y_dim -1]
        p.value = '\u23FA'
        while p != self.points[0][0]:
            p = p.djikstra_from
            # p.value =  f"\033[1m{p.value}\033[0m"
            p.value = '\u23FA'

        return self.points[self.x_dim - 1][self.y_dim -1].djikstra_dist


class Point():
    '''
        single point in the grid
    '''

    def __init__(self, x, y, value, grid):
        self.x = x
        self.y = y
        self.value = value
        self.grid = grid
        self.djikstra_dist = float('inf')
        self.djikstra_from = None

    def neighbors(self):

        neighbors = [
            (self.x - 1, self.y),  # up
            (self.x + 1, self.y),  # down
            (self.x, self.y - 1),   # left
            (self.x, self.y + 1),   # right
        ]

        for x, y in neighbors:
            if x > -1 and y > -1 and x < self.grid.x_dim and y < self.grid.y_dim:
                yield self.grid.points[x][y]



def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    matrix = parse_input(lines, part)
    LOG.debug(f"matrix: {matrix} ")

    g = Grid(matrix)
    # r = g.min_cost_path()
    r = g.djikstra()

    LOG.debug(r)

    with open('grid.txt', 'w') as f:
        f.write(str(g))

    return r


def parse_input(lines, part):
    matrix = []
    if part == 1:
        for line in lines:
            matrix.append(list(map(int, list(line.strip()))))
    if part == 2:
        for line in lines:
            chunk = list(map(int, list(line.strip())))
            newline = chunk
            for i in range(4):
                chunk = [i + 1 if i < 9 else 1 for i in chunk]
                newline = newline + chunk

            matrix.append(newline)

        chunk = matrix
        for i in range(4):
            newchunk = []
            for line in chunk:
                newline = [i + 1 if i < 9 else 1 for i in line]
                newchunk.append(newline)
            chunk = newchunk
            for l in chunk:
                matrix.append(l)

    return matrix


class Test(unittest.TestCase):
    def test_algo(self):
        example_result = algo('input_test.txt', part=1)
        self.assertEqual(example_result, 40)
        example_result = algo('input_test.txt', part=2)
        self.assertEqual(example_result, 315)


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
