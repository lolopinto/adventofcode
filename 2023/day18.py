from __future__ import annotations
from utils import read_file_groups, read_file, ints
import asyncio
from collections import defaultdict
from dataclasses import dataclass
import re
from grid import Grid
from sparse_grid import SparseGrid
import itertools
import math

delta = {
  'R': (0, 1),
  '0': (0, 1),
  'L': (0, -1),
  '2': (0, -1),
  'U': (-1, 0),
  '3': (-1, 0),
  'D': (1, 0),
  '1': (1, 0),
}

# gave up on this one and used shoelace formula + pick's theorem
# based on reddit. not 100% sure i understand what's going on here
# tried to use flood fill, didn't work. tried to work with ChatGPT (aa3ca93c-0946-492c-970a-1c3fd3e82ecb) to get 
# a working solution and that didn't work either
# did a bunch of stuff with sparse grid and that didn't work either.
# https://en.wikipedia.org/wiki/Shoelace_formula
# https://en.wikipedia.org/wiki/Pick's_theorem
async def part1():
  curr = (0, 0)
  area = 0
  
  b = 0
  async for line in read_file("day18input"):
    parts = line.split()
    assert len(parts) == 3
    
    d = delta[parts[0]]
    n = int(parts[1])
    next = (curr[0] + d[0] * n, curr[1] + d[1] * n)
    area += (next[0] * curr[1] - curr[0] * next[1])
    curr = next
    b += n

  # https://en.wikipedia.org/wiki/Shoelace_formula#Shoelace_formula
  # doesn't quite map
  # only gives me the first part...
  print(int(area/2 - b/2 + 1  + b))


async def part2():
  curr = (0, 0)
  area = 0
  
  b = 0
  async for line in read_file("day18input"):
    parts = line.split()
    assert len(parts) == 3
    
    # these 2 swapped for part 2
    n = int(parts[2][2:7], 16)
    d = delta[parts[2][7]]

    next = (curr[0] + d[0] * n, curr[1] + d[1] * n)
    area += (next[0] * curr[1] - curr[0] * next[1])
    curr = next
    b += n
        
  print(int(area/2 - b/2 + 1  + b))

if __name__ == "__main__":
    asyncio.run(part1())
    asyncio.run(part2())
