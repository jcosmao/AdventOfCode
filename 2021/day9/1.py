#!/usr/bin/env python3

import unittest
import sys
import logging
import numpy as np

LOG = logging.getLogger()
logging.basicConfig(encoding='utf-8', level=logging.DEBUG)

class Point(object):
    def __init__(self, coord, val, up, down, left, right):
        self.coord = coord
        self.val = int(val)
        self.up = up
        self.down = down
        self.left = left
        self.right = right
        self.bassin = set()

    def find_adjacent_low_points(self, matrix, points):
        LOG.debug(f"WORKING on low point {self.coord}")
        self.bassin.add(self)

        if self in points:
            points.remove(self)

        for point in self.find_adjacent_points(points, bassin=self.bassin):
            LOG.debug(f"check point = {point.coord}")

            if point not in points:
                LOG.debug(f"Point already in bassin")
                continue


            is_low = False
            for p in point.find_adjacent_points(points, bassin=self.bassin):
                LOG.debug(f"check {p.val} = {p.coord} is > {point.val}")
                if p.val > point.val:
                    is_low = True
                    break
            if is_low:
                LOG.debug(f"FOUD LOW = {point.coord}")
                self.bassin.add(point)
                if point in points:
                    points.remove(point)
                bassin = point.find_adjacent_low_points(matrix, points)
                self.bassin.update(bassin)

        return self.bassin

    def find_adjacent_points(self, points, bassin):
        coords = list(filter(lambda x: x != None, [self.up, self.down, self.right, self.left]))
        adj = set(list(filter(lambda x: x.coord in coords, points))) - set(bassin)
        LOG.debug(f"ADJACENT COORDS = {[a.coord for a in adj]}")
        return list(adj)


def algo(input, ex=1):
    with open(input) as f:
        lines = f.readlines()

    matrix = []
    for l in lines:
        matrix.append(list(l.strip()))

    LOG.debug(np.array(matrix))

    low_points = []
    points = []
    low_point = 0
    for l in range(len(matrix)):
        for c in range(len(matrix[l])):
            current = int(matrix[l][c])

            if c == 0:
                left = 100
                cleft = None
            else:
                left = int(matrix[l][c - 1])
                cleft = (l, c - 1)

            if c == len(matrix[l]) - 1:
                right = 100
                cright = None
            else:
                right = int(matrix[l][c + 1])
                cright = (l, c + 1)

            if l == 0:
                up = 100
                cup = None
            else:
                up = int(matrix[l - 1][c])
                cup = (l - 1, c)

            if l == len(matrix) - 1:
                down = 100
                cdown = None
            else:
                down = int(matrix[l + 1][c])
                cdown = (l + 1, c)

            p = Point(coord=(l,c), val=current, up=cup, down=cdown, left=cleft, right=cright)
            points.append(p)

            LOG.debug(f"line = {l} - col = {c} - current = {current}")
            if current < left and current < right and current < up and current < down:
                LOG.debug(f"{left} {right} {up} {down}")
                low_point += current + 1
                low_points.append(p)

    LOG.debug(low_points)

    if ex == 2:
        bassin = []
        for p in low_points:
            p.find_adjacent_low_points(matrix, points)
            LOG.debug(f"BASSIN = {[pt.coord for pt in p.bassin]} - {len(p.bassin)}")
            bassin.append(len(p.bassin))

        return np.prod(sorted(bassin)[-3:])

    return low_point


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file)
        self.assertEqual(example_result, 15)
        example_result = algo(file, ex=2)
        self.assertEqual(example_result, 1134)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    # if not r.result.failures:
    #     res = algo('input.txt')
    #     LOG.info(f"RESULT from input.txt = {res}")
    #     res = algo('input.txt', ex=2)
    #     LOG.info(f"RESULT from input.txt = {res}")
