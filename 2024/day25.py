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
  lock_heights = []
  key_heights = []

  async for group in read_file_groups("day25input"):
    g = Grid.from_lines(group)
    lock = all(g.get_value(v[0], v[1]) == "#" for v in g.right(0, 0))
    if lock:
  
      heights = []
      for r,c in [(0, 0)] + g.right(0, 0):
        height = 0
        for r2, c2 in g.bottom(r, c):
          if g.get_value(r2, c2) == ".":
            break
          else:
            height += 1
            
        heights.append(height)
      lock_heights.append(heights)

    else:
      heights = []
      for r, c in [(g.height -1 , 0)] + g.right(g.height -1 , 0):
        height = 0
        for r2, c2 in g.top(r, c):
          if g.get_value(r2, c2) == ".":
            break
          else:
            height += 1
            
        heights.append(height)
      key_heights.append(heights)
        
  overlap = 0
  
  for key_height in key_heights:
    for lock_height in lock_heights:
      if not any(kh + lh > 5 for kh, lh in zip(key_height, lock_height)):
        overlap += 1

  print(overlap)


async def part2():
  async for line in read_file("day25input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
