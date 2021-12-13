#!/usr/bin/env python3


def counter() -> int:

    with open('input') as data:
        lines = [x.strip() for x in data.readlines()]

    count = 0

    lst = []
    for i in lines:
        lst.append(int(i))

    triplet_list = [lst[i:i+3] for i in range(0, len(lst), 1)]

    sum_list = []

    for i in triplet_list:
        sum_list.append(sum(i))

    for i, j in enumerate(sum_list[:-1]):
        if j < sum_list[i+1]:
            count += 1

    print(count)

counter()
