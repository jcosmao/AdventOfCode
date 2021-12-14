#!/usr/bin/env python3

from collections import defaultdict
import itertools
import unittest
import logging
from time import perf_counter_ns
import sys

print(sys.getrecursionlimit())
sys.setrecursionlimit(15000)

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)

"""
DFS implementation
https://www.geeksforgeeks.org/find-paths-given-source-destination/?ref=lbp
"""

class Graph:

    def __init__(self, vertices, part):
        self.V = vertices
        self.graph = defaultdict(list)
        self.part = part
        self.visited = dict()

    def add_edge(self, u, v):
        self.graph[u].append(v)

    def _can_visit(self, u):

        nb_visit = self.visited[u]
        if self.part == 1:
            if u.islower() and nb_visit >= 1:
                return False
            else:
                return True

        elif self.part == 2:
            if u == 'start' and nb_visit >= 1:
                return False
            elif u.islower() and nb_visit >= 2:
                return False
            elif u.islower() and nb_visit == 1:
                for small_cave in [c for c in self.graph.keys() if c.islower() and c != u]:
                    if self.visited[small_cave] == 2:
                        return False
                return True
            else:
                return True


    def _init_visit(self):
        self.visited = dict()
        for i in self.graph.keys():
            self.visited[i] = 0

    def _get_all_paths(self, u, d, path):
        ''' A recursive function to print all paths from 'u' to 'd'.
            visited[] keeps track of vertices in current path.
            path[] stores actual vertices and path_index is current
            index in path[] '''

        path.append(u)
        # Mark the current node as visited and store in path
        # we allow passing multiple time through big cave (uppercase vertex)
        self.visited[u] += 1

        # If current vertex is same as destination, then print
        # current path[]
        if u == d:
            yield path
        else:
            # print(f"{u} - {path}")
            # If current vertex is not destination
            # Recur for all the vertices adjacent to this vertex
            for i in self.graph[u]:
                if self._can_visit(i):
                    yield from self._get_all_paths(i, d, path)

        # Remove current vertex from path[] and mark it as unvisited
        path.pop()
        self.visited[u] -= 1

    def get_all_paths(self, s, d):

        self._init_visit()

        # Create an array to store paths
        path = []

        # Call the recursive helper function to print all paths
        return self._get_all_paths(s, d, path)


def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    vertices = set()
    edges = set()
    for line in lines:
        a, b = line.strip().split('-')
        vertices.add(a)
        vertices.add(b)
        edges.add((a, b))

    g = Graph(len(vertices), part)

    for line in lines:
        a, b = line.strip().split('-')
        g.add_edge(a, b)
        g.add_edge(b, a)

    paths = g.get_all_paths('start', 'end')

    return len([p for p in paths])


class Test(unittest.TestCase):
    def test_algo(self):
        example_result = algo('input_test.txt', part=1)
        self.assertEqual(example_result, 10)
        example_result = algo('input_test2.txt', part=1)
        self.assertEqual(example_result, 19)
        example_result = algo('input_test3.txt', part=1)
        self.assertEqual(example_result, 226)
        example_result = algo('input_test.txt', part=2)
        self.assertEqual(example_result, 36)
        example_result = algo('input_test2.txt', part=2)
        self.assertEqual(example_result, 103)
        example_result = algo('input_test3.txt', part=2)
        self.assertEqual(example_result, 3509)


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
