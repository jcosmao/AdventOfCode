#!/usr/bin/env python3


from collections import namedtuple
from queue import Queue
from time import perf_counter_ns

from pprint import pprint

Point = namedtuple('Point',['x_pos','y_pos'])

def get_neighbors(pos,max_rows,max_cols):
    directions = [(0,1),(1,0),(-1,0),(0,-1)]
    result = []

    for direction in directions:
        neighbor = Point(pos.x_pos + direction[0], pos.y_pos + direction[1])
        if 0 <= neighbor.x_pos < max_cols and 0 <= neighbor.y_pos < max_rows:
            result.append(neighbor)
    return result

def is_lower(pos,heights,max_rows,max_cols):
    adjacents = get_neighbors(pos, max_rows, max_cols)
    adjacent_heights = [heights[pos] for pos in adjacents]
    for height in adjacent_heights:
        if heights[pos] >= height:
            return False
    return True

def get_basin_size(pos, heights, nines, max_rows, max_cols):
    frontier = Queue()
    frontier.put(pos)
    reached = set()
    reached.add(pos)

    while not frontier.empty():
        current = frontier.get()
        for neighbor in get_neighbors(current,max_rows, max_cols):
            if neighbor not in reached and neighbor not in nines:
                frontier.put(neighbor)
                reached.add(neighbor)
    return len(reached)

def main():

    START = perf_counter_ns()

    with open('input.txt','r') as infile:
        heights = {Point(x_pos,y_pos):int(height) for y_pos,line in enumerate(infile.readlines()) for x_pos,height in enumerate(line.strip())}
        infile.seek(0)
        max_rows = len([line for line in infile.readlines()])
        infile.seek(0)
        max_cols = len(infile.readline().strip())

    low_points = set()
    nines = set()
    all_points = set(heights.keys())

    for pos, height in heights.items():
        if height == 9:
            nines.add(pos)
        elif is_lower(pos, heights, max_rows, max_cols):
            low_points.add(pos)

    basin_sizes = []
    for low_point in low_points:
        basin_size = get_basin_size(low_point, heights, nines, max_rows, max_cols)
        basin_sizes.append(basin_size)

    basin_sizes.sort(reverse=True)
    total = 1
    for i in (0,1,2):
        total *= basin_sizes[i]

    print(total)



    STOP = perf_counter_ns()
    print(f"Time taken: {STOP - START} nanoseconds")

if __name__ == '__main__':
    main()
