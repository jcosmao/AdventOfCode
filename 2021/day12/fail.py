#!/usr/bin/env python3

import itertools
import unittest
import logging
from igraph import *
from time import perf_counter_ns
from collections import Counter

LOG = logging.getLogger()
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)8s]  %(message)s'
)


def algo(input, part):
    with open(input) as f:
        lines = f.readlines()

    graph = Graph(directed=False)
    # g = {}
    vertices = set()
    edges = set()
    for line in lines:
        a, b = line.strip().split('-')
        vertices.add(a)
        vertices.add(b)
        edges.add((a, b))

    graph.add_vertices(list(vertices))
    graph.add_edges(edges)

    dest = []
    allowed_multi = []
    for vs in graph.vs:
        LOG.debug(vs)
        if vs['name'] not in ['start', 'end'] and vs['name'].islower():
            dest.append(vs)
        if vs['name'].isupper():
            allowed_multi.append(vs.index)


    LOG.debug(allowed_multi)
    LOG.debug(dest)
    # help(graph)
    # all = paths_from_to(graph, graph.vs.find(name='start'), graph.vs.find(name='end'))
    all_path = graph.get_all_simple_paths('start', 'end')

    for intermediate in dest:
        p1l = graph.get_all_simple_paths('start', intermediate)
        p2l = graph.get_all_simple_paths(intermediate, 'end')
        for p1 in p1l:
            for p2 in p2l:
                if p1[-1] == p2[0]:
                    new_path = p1[0:-1] + p2
                    all_path.append(new_path)

    LOG.debug(all_path)

    all_path.sort()
    all_path = list(all_path for all_path,_ in itertools.groupby(all_path))
    routes = all_path.copy()

    for path in routes:
        count = dict(filter(lambda x: x[1] > 1, Counter(path).items())).keys()
        for k in count:
            if k not in allowed_multi:
                LOG.debug(f"path {path} not allowed")
                if path in all_path:
                    all_path.remove(path)

    LOG.debug(all_path)
    return len(all_path)


    # visual_style = {}

    # out_name = "graph_coloured.png"

    # # Set bbox and margin
    # visual_style["bbox"] = (400, 400)
    # visual_style["margin"] = 27

    # # Set vertex colours
    # graph.vs["color"] = ["red", "green", "blue", "yellow", "orange"]

    # # Set vertex size
    # visual_style["vertex_size"] = 45

    # # Set vertex lable size
    # visual_style["vertex_label_size"] = 22

    # # Don't curve the edges
    # visual_style["edge_curved"] = False

    # # Set the layout
    # my_layout = graph.layout_lgl()
    # visual_style["layout"] = my_layout

    # # Plot the graph
    # plot(graph, out_name, **visual_style)


class Test(unittest.TestCase):
    def test_algo(self):
        # example_result = algo('input_test.txt', part=1)
        # self.assertEqual(example_result, 10)
        # example_result = algo('input_test2.txt', part=1)
        # self.assertEqual(example_result, 19)
        example_result = algo('input_test3.txt', part=1)
        self.assertEqual(example_result, 226)


if __name__ == "__main__":
    r = unittest.main(exit=False)

    # Test pass on example input, work on data set
    if not r.result.failures and not r.result.errors:
        LOG.setLevel(logging.INFO)

        START = perf_counter_ns()
        res = algo('input.txt', part=1)
        STOP = perf_counter_ns()
        LOG.info(
            f"RESULT part1 from input.txt = {res} - (TIME {STOP - START} ns)")

        START = perf_counter_ns()
        res = algo('input.txt', part=2)
        STOP = perf_counter_ns()
        LOG.info(
            f"RESULT part2 from input.txt = {res} - (TIME {STOP - START} ns)")
