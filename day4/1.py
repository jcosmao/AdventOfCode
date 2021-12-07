#!/usr/bin/env python3

# >>> import numpy as np
# >>> a=np.array([[1,2,3],[11,12,13],[21,22,23]])
# >>> a
# array([[ 1,  2,  3],
#        [11, 12, 13],
#        [21, 22, 23]])
# >>> a[:,0]
# array([ 1, 11, 21])

import numpy as np

f = open('input.txt')
lines = f.readlines()

bingo_num = [n for n in lines.pop(0).strip().split(',')]
bingo_matrix = []
matrix_size = 5

# print(bingo_num)

# create matrix
a = []
for l in lines:
    if l.strip():
        a.append([n for n in l.strip().split()])
        if len(a) == matrix_size:
            bingo_matrix.append(a)
            a = []

# print(bingo_matrix)

def mark_bingo(matrix, num):
    for card in matrix:
        for line in card:
            for i in range(len(line)):
                if line[i] == num:
                    line[i] = f"*{num}"

def check_winner(matrix):
    for card in matrix:
        for i in range(len(card)):
            # check by line
            res = list(filter(lambda x: '*' in x, card[i]))
            if len(res) == matrix_size:
                # GOT winner
                return card
            # check by col - use numpy
            res = list(filter(lambda x: '*' in x, card[:, i]))
            if len(res) == matrix_size:
                # GOT winner
                return card


for i in bingo_num:
    mark_bingo(bingo_matrix, i)
    wincard = check_winner(np.array(bingo_matrix))
    if wincard is not None:
        print(f"Got winner with num {i}: {wincard}")
        unmarked = []
        for line in wincard:
            for n in line:
                if '*' not in n:
                    unmarked.append(int(n))
        print(sum(unmarked)*int(i))
        break
