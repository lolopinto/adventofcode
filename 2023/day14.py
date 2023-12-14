from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
import itertools

async def part1():
  g = await Grid.grid_from_file("day14input")
  g.print()
  
  # for each column
  for c in range(g.width):
    for r in range(g.height):
      match g.get_value(r, c):
        # can't move anything. we're done
        case '#':
          continue
        # rock, keep going 
        case 'O':
          continue
        
      for r2, c2 in g.bottom(r, c):
        # look for the next rock
        match g.get_value(r2, c2):
          case '#':
            break
          case 'O':
            # found a rock, swap 
            g.set(r2, c2, '.')
            g.set(r, c, 'O')
            break
  
  g.print()

  sum = 0  
  for r in range(g.height):
    row = g.height - r
    for c in range(g.width):
      if g.get_value(r, c) == 'O':
        sum += row

  print(sum)

async def part2():
  async for line in read_file("day14input"):
    pass


if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
