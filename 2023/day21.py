from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

async def part1():
  g = await Grid.grid_from_file("day21input")
  
  start = g.find('S')
  assert start is not None
  # TODO 64
  l = {start}
  for i in range(64):
    l2 = set()
    for curr in l:
      for n in g.neighbors(curr[0], curr[1]):
        v = g.get_value(n[0], n[1])
        if v == '.' or v == 'S':
          l2.add(n)
    l = l2

  print(len(l))

  async for line in read_file("day21input"):
    pass


async def part2():
  async for line in read_file("day21input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
