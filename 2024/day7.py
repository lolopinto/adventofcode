from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
import enum
import math

async def part1():
  res = 0

  async for line in read_file("day7input"):
    left, right = line.split(": ")
    value = int(left)
    values = ints(right)
    
    assert len(values) >= 2

    def perform_op(op, left, right):
      if op == "*":
        return left * right
      elif op == "+":
        return left + right
      else:
        raise ValueError("Invalid operator")

    for v in itertools.product(["*", "+"], repeat=len(values) - 1):

      s = perform_op(v[0], values[0], values[1])
      idx = 1
      for j in range(2, len(values)):
        s = perform_op(v[idx], s, values[j])
        idx += 1

      if s == value:
        res += value
        break

  print(res)



async def part2():
  async for line in read_file("day3input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
