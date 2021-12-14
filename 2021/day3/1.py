#!/usr/bin/env python3

f = open("input_test.txt")
lines = f.readlines()

gamma = ''
epsylon = ''
global columns
columns = []

def get_max(lines, i):
    columns.append({'0': 0, '1': 0})
    for l in lines:
        nb = list(l.strip())[i]
        columns[i][str(nb)] += 1

l = lines
for i in range(len(lines[0].strip())):
    print(i)
    if i > 0:
        l = list(filter(lambda x: list(x)[i-1] == list(gamma)[i-1], filtered_lines))

    if len(l) == 1:
        gamma = l[0].strip()
        break

    filtered_lines = l
    print(filtered_lines)
    get_max(filtered_lines, i)

    if columns[i]['0'] > columns[i]['1']:
        gamma = gamma + '0'
    else:
        gamma = gamma + '1'


    print(columns)
    print(f"gamma = {gamma}")

columns = []
l = lines
for i in range(len(lines[0].strip())):
    print(i)
    if i > 0:
        l = list(filter(lambda x: list(x)[i-1] == list(epsylon)[i-1], filtered_lines))

    if len(l) == 1:
        epsylon = l[0].strip()
        break

    filtered_lines = l
    print(filtered_lines)
    get_max(filtered_lines, i)

    if columns[i]['0'] > columns[i]['1']:
        epsylon = epsylon + '1'
    else:
        epsylon = epsylon + '0'

    print(columns)
    print(f"epsylon = {epsylon}")

print(columns)
print(int(gamma, base=2)*int(epsylon, base=2))
