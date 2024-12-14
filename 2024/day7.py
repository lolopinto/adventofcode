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

async def solve(ops: list[str]) -> int:
  res = 0
  async for line in read_file("day7input"):
    left, right = line.split(": ")
    value = int(left)
    values = ints(right)
    
    assert len(values) >= 2

    def perform_op(op, left, right):
      match op:
        case "*":
          return left * right
        case "+":
          return left + right
        case "||":
          return int(str(left) + str(right))
        case "_":
          raise ValueError("Invalid operator")

    for v in itertools.product(ops, repeat=len(values) - 1):
      s = perform_op(v[0], values[0], values[1])

      idx = 1
      for j in range(2, len(values)):
        s = perform_op(v[idx], s, values[j])
        idx += 1

      if s == value:
        res += value
        break
      
  return res

async def part1():
  res = await solve(["*", "+"])
  print(res)


async def part2():
  res = await solve(["*", "+", "||"])
  print(res)

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
