#!/usr/bin/env python3

import unittest
import logging
import re
import numpy

LOG = logging.getLogger()
logging.basicConfig(encoding='utf-8', level=logging.DEBUG)


def algo(input):
    with open(input) as f:
        lines = f.readlines()
    gps = _parse_coordinate(lines)
    matrix = _init_matrix(gps)
    draw = _draw_lines(gps, matrix)
    dangerous = _most_dangerous_area(draw)

    return dangerous

def _parse_coordinate(lines):
    """return gps points as (x1,y1) (xx,xx) ... (x2,y2) - only lines"""
    res = []
    for line in lines:
        find = re.match(r"(\d+),(\d+) -> (\d+),(\d+)",line)
        if find:
            x1 = find.group(1)
            y1 = find.group(2)
            x2 = find.group(3)
            y2 = find.group(4)

            res.append([(int(x1),int(y1)),(int(x2),int(y2))])

    LOG.debug(f"GPS coordinate = {res}")
    return res

def _init_matrix(gps):
    """find max and draw . in a matrix"""
    matrix = []
    max_h = 0
    max_v = 0
    for coord in gps:
        for tuple in coord:
            max_h = max(max_h, tuple[0])
            max_v = max(max_v, tuple[1])

    row = [0 for i in range(0, max_h + 1)]
    for j in range(0, max_v + 1):
        matrix.append(row.copy())

    LOG.debug(numpy.array(matrix))
    return matrix

def _draw_lines(gps, matrix):
    """return matrix with lines representation"""

    for coord in gps:
        if coord[0][0] != coord[1][0] and coord[0][1] != coord[1][1]:
            x1 = coord[0][1]
            x2 = coord[1][1]
            y1 = coord[0][0]
            y2 = coord[1][0]
            x_list = [i for i in range(x1, x2 + 1)] if x1 < x2 else [i for i in range(x1, x2 -1, -1)]
            y_list = [i for i in range(y1, y2 + 1)] if y1 < y2 else [i for i in range(y1, y2 -1, -1)]
            if len(x_list) != len(y_list):
                LOG.debug(f"{coord} is not 45 degree diagonal line")
                continue

            for i in range(len(x_list)):
                matrix[x_list[i]][y_list[i]] += 1

        # horizontal line from coord[0][0] to coord[1][0]
        elif coord[0][0] != coord[1][0]:
            x = coord[0][1]
            y1 = min(coord[0][0], coord[1][0])
            y2 = max(coord[0][0], coord[1][0])
            y_list = [i for i in range(y1, y2 + 1)]
            LOG.debug(f"horizontal line = {x} - points = {y_list}")
            for y in y_list:
                matrix[x][y] += 1

        # vertical line from coord[0][1] to coord[1][1]
        elif coord[0][1] != coord[1][1]:
            y = coord[0][0]
            x1 = min(coord[0][1], coord[1][1])
            x2 = max(coord[0][1], coord[1][1])
            x_list = [i for i in range(x1, x2 + 1)]
            LOG.debug(f"vertical line = {y} - points = {x_list}")
            for x in x_list:
                matrix[x][y] += 1

        # single point
        else:
            LOG.debug(f"Single point: {coord}")
            matrix[coord[0][0]][coord[0][1]] += 1

    LOG.debug(numpy.array(matrix))
    return matrix

def _most_dangerous_area(matrix):
    danger = {}
    for line in matrix:
        for p in line:
            if not p in danger.keys():
                danger[p] = 0
            danger[p] += 1
    LOG.debug(f"Most dangerous: {danger}")

    dangerous = 0
    for k in danger.keys():
        if k >= 2:
            dangerous += danger[k]

    return dangerous


class Test(unittest.TestCase):
    def test_algo(self):
        file = 'input_test.txt'
        example_result = algo(file)
        self.assertEqual(example_result, 12)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    if not r.result.failures:
        res = algo('input.txt')
        LOG.info(f"RESULT from input.txt = {res}")
