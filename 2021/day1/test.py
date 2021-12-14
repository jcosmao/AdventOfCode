#!/usr/bin/python3

with open("input", 'r') as file:
    total = 0
    prev = None
    for i in [int(x) for x in file.read().splitlines()]:
        if prev and prev < i:
            total += 1
        prev = i

print(total)
