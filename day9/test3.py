'''
tools for advent code day 9
'''
from collections import Counter
from operator import le


def load_data(text):
    '''
    load data from file
    '''
    data = text.strip().split('\n')
    data = [[int(i) for i in x] for x in data]

    return data


class Point():
    '''
        single point in the grid
    '''
    def __init__(self, x, y, value,grid=None):
        self.x = x
        self.y = y
        self.value = value
        self.lowest = None
        self.basin = None
        if grid:
            self.grid = grid
        else:
            self.grid = None

    def neighbors(self,grid=None):
        if self.grid==None:
            self.grid = grid
        for x, y in [[self.x - 1, self.y], [self.x + 1, self.y],
                     [self.x, self.y - 1], [self.x, self.y + 1]]:
            if x > -1 and y > -1 and x < self.grid.x_dim and y < self.grid.y_dim:
                yield self.grid.points[x][y]

    def is_lowest(self):
        if self.lowest == False:
            return False
        if self.value == 0:
            return True
        if self.value == 9:
            return False
        for other_point in self.neighbors():
            if other_point.value <= self.value:
                return False
            else:
                other_point.lowest = False
        return True

class Grid():
    '''
     class for the grid, with methods to find the basins and the lowest points
    '''
    def __init__(self, data):
        self.data = data
        self.points = data
        for x,line in enumerate(data):
            for y,value in enumerate(line):
                self.points[x][y] = Point(x, y, value,grid=self)

        self.x_dim = len(self.points)
        self.y_dim = len(self.points[0])

    def yield_points(self):
        for x, line in enumerate(self.points):
            for y, point in enumerate(line):
                yield x,y,point

    def find_low_points(self):
        lowests = []
        for i, line in enumerate(self.points):
            for j, _ in enumerate(line):
                point = self.points[i][j]
                if point.is_lowest():
                    lowests.append(point.value)
        return lowests

    def find_basins(self):
        '''
        find all basins
        '''
        basins = []
        basin = []
        neighbours = []
        for _,_,point in self.yield_points():
            neighbours.append(point)
            while neighbours:
                point = neighbours.pop()
                if point.value != 9 and point.basin is None:
                    basin.append(point)
                    point.basin = len(basins)
                    neighbours.extend(list(point.neighbors()))
            basins.append(basin[:])
            basin= []
        return basins


def calc_risk_levels(data):
    '''
    sum of all the low points
    '''
    return sum(data) + len(data)


def calc_largest_basins(basin_list):
    '''
    find largest basin
    '''
    basin_list_len = [len(x) for x in basin_list if x != []]
    basin_list_len.sort(reverse=True)
    return basin_list_len[0]*basin_list_len[1]*basin_list_len[2]


def part1(text):
    '''
     find lowest points
    '''
    data = load_data(text)
    grid = Grid(data)
    low_points = grid.find_low_points()
    return calc_risk_levels(low_points)


def part2(text):
    '''
    find largest basins'''
    data = load_data(text)
    grid = Grid(data)
    basins  = grid.find_basins()
    return calc_largest_basins(basins)

with open('input.txt') as f:
    print(part2(f.read()))
