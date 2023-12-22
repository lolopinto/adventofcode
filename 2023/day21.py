from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools
from infinite_grid import InfiniteGrid

async def part1():
  g = await Grid.grid_from_file("day21input")
  
  start = g.find('S')
  assert start is not None
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

async def part2():
  temp = await Grid.grid_from_file("day21input")
  
  start = temp.find('S')
  assert start is not None
  
  g = InfiniteGrid.from_grid(temp)
  l = {start}
  # if we can't even do this on its own before infinite...
  
  # cycle in first 1000 before infinite, now we need to do infinite
  # 15453 too low  based on seeing the wrap in first 1000
  for i in range(500):
    l2 = set()
    for curr in l:
      for n in g.neighbors(curr):
        v = g.get_valuex(n)
        if v == '.' or v == 'S':
          l2.add(n)
        
    print(i+1 ,len(l2))
    
    # can't see any tren
    l = l2

  print(len(l))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
