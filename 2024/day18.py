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

def get_consts():
  if True:
    return PROD
  else:
    return EXAMPLE
  
async def part1():
  length, steps = get_consts()
  
  g = Grid.square_grid(length)
  i = 0
  async for line in read_file("day18input"):
    p = ints(line, ",")
    g.set(p[0], p[1], "#")
    
    i += 1
    if i == steps:
      break
    
  print(g.dijkstra2((0, 0), (g.width - 1, g.height - 1), "#"))


async def part2():
  length, steps = get_consts()
  
  lines = [line async for line in read_file("day18input")]

  for i in range(steps, len(lines)):
    g = Grid.square_grid(length)
    
    for j in range(i):
      p = ints(lines[j], ",")
      g.set(p[0], p[1], "#")

    # 18.93s for this
    try:
      g.dijkstra2((0, 0), (g.width - 1, g.height - 1), "#")
    except AssertionError:
      print(lines[j])
      break

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
