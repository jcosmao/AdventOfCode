#!/usr/bin/env python3

f = open("input.txt")

position = 0
depth = 0
aim = 0

for l in f.readlines():
    instruction = l.split()[0]
    num = int(l.split()[1])
    if instruction == 'forward':
        position += num
        depth += aim * num
    elif instruction == 'up':
        aim -= num
    elif instruction == 'down':
        aim += num

print(position)
print(depth)

print(position*depth)
