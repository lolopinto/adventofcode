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

EXAMPLE = (7, 12)
PROD = (71, 1024)
async def part1():
  if True:
    length = PROD[0]
    steps = PROD[1]
  else:
    length = EXAMPLE[0]
    steps = EXAMPLE[1]
  
  g = Grid.square_grid(length)
  i = 0
  async for line in read_file("day18input"):
    p = ints(line, ",")
    print
    g.set(p[0], p[1], "#")
    
    i += 1
    if i == steps:
      break
    
  print(g.dijkstra2((0, 0), (g.width - 1, g.height - 1), "#"))


async def part2():
  async for line in read_file("day18input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
