#!/usr/bin/env python3

f = open("input.txt")

position = 0
depth = 0

for l in f.readlines():
    instruction = l.split()[0]
    num = int(l.split()[1])
    if instruction == 'forward':
        position += num
    elif instruction == 'up':
        depth -= num
    elif instruction == 'down':
        depth += num

print(position)
print(depth)

print(position*depth)
