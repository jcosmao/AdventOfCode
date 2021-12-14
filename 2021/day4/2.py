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
    winners = []
    for card in np.array(matrix):
        for i in range(len(card)):
            # check by line
            res = list(filter(lambda x: '*' in x, card[i]))
            if len(res) == matrix_size:
                # GOT winner
                winners.append(card.tolist())
            # check by col - use numpy
            res = list(filter(lambda x: '*' in x, card[:, i]))
            if len(res) == matrix_size:
                # GOT winner
                winners.append(card.tolist())
    print(f"WINNERS: {winners}")
    return winners

last_win = None
num_win = None

for i in bingo_num:
    mark_bingo(bingo_matrix, i)
    wincards = check_winner(bingo_matrix)
    if wincards is not None:
        for wincard in wincards:
            print(f"remove {wincard}")
            if wincard in bingo_matrix:
                bingo_matrix.remove(wincard)

            last_win = wincard
            num_win = i


print(f"Got winner with num {num_win}: {last_win}")
unmarked = []
for line in last_win:
    for n in line:
        if '*' not in n:
            unmarked.append(int(n))
print(sum(unmarked)*int(num_win))
